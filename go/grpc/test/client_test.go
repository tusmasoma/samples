package test

import (
	"context"
	"net"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tusmasoma/samples/go/grpc"
	user "github.com/tusmasoma/samples/go/grpc/test/proto"
	g "google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

func TestNewClientWithDial_Success(t *testing.T) {
	ctx := context.Background()
	lis = bufconn.Listen(bufSize)
	s := g.NewServer()
	go func() {
		_ = s.Serve(lis)
	}()
	opts := []g.DialOption{
		g.WithContextDialer(bufDialer),
		g.WithInsecure(),
	}
	client, err := grpc.NewClientWithDial[user.UserServiceClient](ctx, user.NewUserServiceClient, opts...)
	require.NoError(t, err)
	require.NotNil(t, client)
}

func TestServiceNameFromType(t *testing.T) {
	t.Run("TEST: user.UserServiceClient", func(t *testing.T) {
		got := grpc.ServiceNameFromType[user.UserServiceClient]()
		require.Equal(t, "user", got)
	})
}

const bufSize = 1024 * 1024

var lis *bufconn.Listener

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}
