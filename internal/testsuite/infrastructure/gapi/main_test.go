package gapi_test

import (
	"context"
	"log"
	"net"
	"os"
	"testing"

	"govent/internal/domain/types"
	"govent/internal/infrastructure/configuration"
	"govent/internal/infrastructure/gapi"
	"govent/internal/testsuite/infrastructure/testdb"
	internal_test "govent/internal/testsuite/internal"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener
var grpcServer *grpc.Server
var grpcClient gapi.GoldenserviceClient
var ctx context.Context

func TestMain(m *testing.M) {
	setup()

	code := m.Run()

	grpcServer.Stop()
	os.Exit(code)
}

func setup() {
	lis = bufconn.Listen(bufSize)

	grpcServer = grpc.NewServer()

	config := &configuration.GoldenConfiguration{}
	server := gapi.NewServer(":8080", testdb.GetRepository(), *config)
	gapi.RegisterGoldenserviceServer(grpcServer, server)

	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("['.']:> Error serving gRPC server: %v", err)
		}
	}()

	conn, err := grpc.NewClient(
		"passthrough://localhost/bufnet",
		grpc.WithContextDialer(bufDialer),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("['.']:> Failed to dial bufnet: %v", err)
	}

	grpcClient = gapi.NewGoldenserviceClient(conn)
	ctx = context.Background()
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func createPersistedRandomGolden() *types.Golden {
	golden := internal_test.NewRandomGolden()
	_ = testdb.GetRepository().Create(golden)

	return golden
}

func deletePersistedRandomGolden(goldenId string) {
	id, err := types.NewGoldenId(goldenId)
	if err != nil {
		return
	}

	_ = testdb.GetRepository().Delete(id)
}
