package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"testing"

	protoFile "mysqlbadconnection/proto/gen/proto"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	grpcbackoff "google.golang.org/grpc/backoff"
	"google.golang.org/grpc/credentials/insecure"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestMySQLConcurrency(t *testing.T) {
	t.Parallel()
	mysqlDocker := RunMySQLTestContainer(t)
	mysqlUri := mysqlDocker.GetConnectionURI(true)

	ctx, cancel := context.WithCancel(context.Background())
	t.Cleanup(cancel)

	// pick a port where the server will run
	grpcPort, grpcPortReleaser := TCPRandomPort()
	grpcPortReleaser()

	go func() {
		err := runServer(ctx, grpcPort, mysqlUri)
		if err != nil {
			t.Log(err)
		}
	}()

	// create client of the server
	serverConn, err := grpc.Dial(fmt.Sprintf("0.0.0.0:%d", grpcPort),
		grpc.WithBlock(),
		grpc.WithConnectParams(grpc.ConnectParams{Backoff: grpcbackoff.DefaultConfig}),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	require.NoError(t, err)
	serverClient := protoFile.NewMysqlIssueDemoClient(serverConn)

	for i := 0; i < 2600; i++ {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			t.Parallel()

			_, err := serverClient.Write(ctx, nil)
			require.NoError(t, err)

			_, err = serverClient.GetAllStores(ctx, nil)
			require.NoError(t, err)

			t.Log("Passed", i)
		})
	}
}

func TCPRandomPort() (int, func()) {
	l, err := net.Listen("tcp", "")
	if err != nil {
		panic(err)
	}
	return l.Addr().(*net.TCPAddr).Port, func() {
		l.Close()
	}
}
