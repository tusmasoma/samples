package grpc

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	"google.golang.org/grpc"
)

const defaultPort = 9000

func NewClientWithDial[T any](ctx context.Context, newClient func(grpc.ClientConnInterface) T, opts ...grpc.DialOption) (T, error) {
	serviceName := ServiceNameFromType[T]()
	target := fmt.Sprintf("%s-server:%d", serviceName, defaultPort)
	conn, err := grpc.DialContext(ctx, target, opts...)
	if err != nil {
		return *new(T), fmt.Errorf("grpc: failed to dial %s: %w", target, err)
	}
	return newClient(conn), nil
}

func ServiceNameFromType[T any]() string {
	t := reflect.TypeFor[T]()
	typeName := t.String()
	packageName := strings.Split(typeName, ".")[0]
	return packageName
}
