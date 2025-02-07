package sequencer

import (
	"context"
	"encoding/json"
	"net"
	"os"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/gogoproto/proto"
	"github.com/rotisserie/eris"
	zerolog "github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	namespacetypes "pkg.world.dev/world-engine/evm/x/namespace/types"
	"pkg.world.dev/world-engine/evm/x/shard/keeper"
	"pkg.world.dev/world-engine/evm/x/shard/types"
	"pkg.world.dev/world-engine/rift/credentials"
	shard "pkg.world.dev/world-engine/rift/shard/v2"
)

const (
	defaultPort = "9601"
)

var (
	Name                                = "shard_sequencer"
	_    shard.TransactionHandlerServer = &Sequencer{}
)

// Sequencer handles sequencing game shard transactions.
type Sequencer struct {
	shard.UnimplementedTransactionHandlerServer
	moduleAddr     sdk.AccAddress
	tq             *TxQueue
	queryCtxGetter GetQueryCtxFn
	shardKeeper    *keeper.Keeper

	// opts
	routerKey string
}

type GetQueryCtxFn func(height int64, prove bool) (sdk.Context, error)

// New returns a new game shard sequencer server. It runs on a default port of 9601,
// unless the SHARD_SEQUENCER_PORT environment variable is set.
//
// The sequencer exposes a single gRPC endpoint, SubmitShardTx, which will take in transactions from game shards,
// indexed by namespace. At every block, the sequencer tx queue is flushed, and processed in the storage shard storage
// module, persisting the data to the blockchain.
func New(shardKeeper *keeper.Keeper, queryCtxGetter GetQueryCtxFn, opts ...Option) *Sequencer {
	s := &Sequencer{
		moduleAddr:     shardKeeper.AuthorityAddress(),
		tq:             NewTxQueue(shardKeeper.AuthorityAddress().String()),
		queryCtxGetter: queryCtxGetter,
		shardKeeper:    shardKeeper,
	}
	for _, opt := range opts {
		opt(s)
	}
	return s
}

// Serve serves the server in a new go routine.
func (s *Sequencer) Serve() {
	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(s.serverCallInterceptor))
	shard.RegisterTransactionHandlerServer(grpcServer, s)
	port := defaultPort
	// check if a custom port was set
	if setPort := os.Getenv("SHARD_SEQUENCER_PORT"); setPort != "" {
		if _, err := strconv.Atoi(setPort); err == nil {
			port = setPort
		}
	}
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		zerolog.Fatal().Err(err).Msg("game shard sequencer failed to open listener")
	}
	go func() {
		err = grpcServer.Serve(listener)
		if err != nil {
			zerolog.Fatal().Err(err).Msg("game shard sequencer failed to serve grpc server")
		}
	}()
}

// FlushMessages empties and returns all messages stored in the queue.
func (s *Sequencer) FlushMessages() ([]*types.SubmitShardTxRequest, []*namespacetypes.UpdateNamespaceRequest) {
	return s.tq.FlushTxQueue(), s.tq.FlushInitQueue()
}

// Submit appends the game shard tx submission to the tx queue.
func (s *Sequencer) Submit(_ context.Context, req *shard.SubmitTransactionsRequest) (
	*shard.SubmitTransactionsResponse, error,
) {
	txIDs := sortMapKeys(req.GetTransactions())
	for _, txID := range txIDs {
		txs := req.GetTransactions()[txID].GetTxs()
		for _, tx := range txs {
			bz, err := proto.Marshal(tx)
			if err != nil {
				return nil, eris.Wrap(err, "failed to marshal transaction")
			}
			s.tq.AddTx(req.GetNamespace(), req.GetEpoch(), req.GetUnixTimestamp(), txID, bz)
		}
	}
	return &shard.SubmitTransactionsResponse{}, nil
}

func (s *Sequencer) RegisterGameShard(
	_ context.Context,
	req *shard.RegisterGameShardRequest,
) (*shard.RegisterGameShardResponse, error) {
	s.tq.AddInitMsg(req.GetNamespace(), req.GetRouterAddress())
	return &shard.RegisterGameShardResponse{}, nil
}

func (s *Sequencer) QueryTransactions(
	_ context.Context,
	request *shard.QueryTransactionsRequest,
) (*shard.QueryTransactionsResponse, error) {
	cosmosCtx, err := s.queryCtxGetter(0, false)
	if err != nil {
		return nil, eris.Wrap(err, "failed to get query context")
	}

	convertedQueryType := types.QueryTransactionsRequest{
		Namespace: request.GetNamespace(),
		Page: &types.PageRequest{
			Key:   request.GetPage().GetKey(),
			Limit: request.GetPage().GetLimit(),
		},
	}
	res, err := s.shardKeeper.Transactions(cosmosCtx, &convertedQueryType)
	if err != nil {
		return nil, eris.Wrap(err, "failed to query transactions")
	}
	bz, err := json.Marshal(res)
	if err != nil {
		return nil, eris.Wrap(err, "failed to marshal response")
	}
	convertedResponse := new(shard.QueryTransactionsResponse)
	err = json.Unmarshal(bz, convertedResponse)
	if err != nil {
		return nil, eris.Wrap(err, "failed to unmarshal evm type response into rift type response")
	}
	return convertedResponse, nil
}

// serverCallInterceptor catches calls to handlers and ensures they have the right secret routerKey.
func (s *Sequencer) serverCallInterceptor(
	ctx context.Context,
	req any,
	_ *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (resp any, err error) {
	rtrKey, err := credentials.TokenFromIncomingContext(ctx)
	if err != nil {
		return nil, err
	}

	if rtrKey != s.routerKey {
		return nil, status.Errorf(codes.Unauthenticated, "invalid %s", credentials.TokenKey)
	}

	return handler(ctx, req)
}
