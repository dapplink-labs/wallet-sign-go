package rpc

import (
	"context"
	"fmt"
	"net"
	"sync/atomic"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/ethereum/go-ethereum/log"

	"github.com/dapplink-labs/wallet-sign-go/leveldb"
	"github.com/dapplink-labs/wallet-sign-go/protobuf/wallet"
)

const MaxRecvMessageSize = 1024 * 1024 * 30000

type RpcServerConfig struct {
	GrpcHostname string
	GrpcPort     int
}

type RpcServer struct {
	*RpcServerConfig
	db *leveldb.Keys

	wallet.UnimplementedWalletServiceServer
	stopped atomic.Bool
}

func (s *RpcServer) Stop(ctx context.Context) error {
	s.stopped.Store(true)
	return nil
}

func (s *RpcServer) Stopped() bool {
	return s.stopped.Load()
}

func NewRpcServer(db *leveldb.Keys, config *RpcServerConfig) (*RpcServer, error) {
	return &RpcServer{
		RpcServerConfig: config,
		db:              db,
	}, nil
}

func (s *RpcServer) Start(ctx context.Context) error {
	go func(s *RpcServer) {
		addr := fmt.Sprintf("%s:%d", s.GrpcHostname, s.GrpcPort)
		log.Info("start rpc services", "addr", addr)
		listener, err := net.Listen("tcp", addr)
		if err != nil {
			log.Error("Could not start tcp listener. ")
		}

		opt := grpc.MaxRecvMsgSize(MaxRecvMessageSize)

		gs := grpc.NewServer(
			opt,
			grpc.ChainUnaryInterceptor(
				nil,
			),
		)
		reflection.Register(gs)

		wallet.RegisterWalletServiceServer(gs, s)

		log.Info("Grpc info", "port", s.GrpcPort, "address", listener.Addr())
		if err := gs.Serve(listener); err != nil {
			log.Error("Could not GRPC services")
		}
	}(s)
	return nil
}