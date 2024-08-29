package main

import (
	"context"
	"fmt"
	"time"

	"net"

	protoFile "mysqlbadconnection/proto/gen/proto"

	"github.com/oklog/ulid/v2"
	"google.golang.org/grpc"
)

type MyServer struct {
	protoFile.UnimplementedMysqlIssueDemoServer
	db *MySQL
}

func (s *MyServer) IsReady(ctx context.Context) (bool, error) {
	return true, nil
}
func (s *MyServer) Close() {
	s.db.Close()
}

func main() {
}

func runServer(ctx context.Context, grpcPort int, dbUri string) error {
	mysqlDb, err := New(dbUri, &Config{
		MaxIdleConns: 10,
		MaxOpenConns: 30,
	})
	if err != nil {
		return err
	}

	service := &MyServer{
		db: mysqlDb,
	}
	grpcServer := grpc.NewServer()
	protoFile.RegisterMysqlIssueDemoServer(grpcServer, service)
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", grpcPort))
	if err != nil {
		return err
	}

	go func() {
		fmt.Println(fmt.Sprintf("grpc server listening on %d", grpcPort))
		if err := grpcServer.Serve(lis); err != nil {
			fmt.Println("failed to start gRPC server: %w", err)
		}
	}()
	// wait for goroutine above to start
	time.Sleep(1 * time.Second)

	<-ctx.Done()

	fmt.Println("attempting to shutdown gracefully...")

	grpcServer.GracefulStop()
	service.Close()

	fmt.Println("server exited. goodbye 👋")
	return nil
}

type Store struct {
	Id   string
	Name string
}

func (s *MyServer) GetAllStores(ctx context.Context, _ *protoFile.GetAllStoresRequest) (*protoFile.GetAllStoresResponse, error) {
	all, err := s.db.ReadAll(ctx)
	if err != nil {
		return nil, err
	}
	var res []string
	for _, store := range all {
		res = append(res, store.Id)
	}
	return &protoFile.GetAllStoresResponse{StoreIds: res}, nil
}

func (s *MyServer) Write(ctx context.Context, _ *protoFile.WriteRequest) (*protoFile.WriteResponse, error) {
	randomStore := ulid.Make().String()
	_, err := s.db.Write(ctx, randomStore)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
