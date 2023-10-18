package calc

import (
	"context"
	"fmt"
	"path/filepath"
	"reflect"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
)

// Local returns a ClientConnInterface that routes traffic to the provided
// server.
//
// NOTE(mwhittaker): Right now, this is implemented with reflection, but beaver
// could generate code for this.
func Local[T any](server T) grpc.ClientConnInterface {
	return &local[T]{server}
}

type local[T any] struct {
	server T
}

func (l *local[T]) Invoke(ctx context.Context, method string, args any, reply any, opts ...grpc.CallOption) error {
	v := reflect.ValueOf(l.server)
	m := v.MethodByName(filepath.Base(method))
	returns := m.Call([]reflect.Value{reflect.ValueOf(ctx), reflect.ValueOf(args)})
	if len(returns) != 2 {
		panic(fmt.Errorf("bad returns %v: got %d, want 2", returns, len(returns)))
	}
	proto.Merge(reply.(proto.Message), returns[0].Interface().(proto.Message))
	if x := returns[1].Interface(); x != nil {
		return x.(error)
	}
	return nil
}

func (l *local[T]) NewStream(ctx context.Context, desc *grpc.StreamDesc, method string, opts ...grpc.CallOption) (grpc.ClientStream, error) {
	panic(fmt.Errorf("NewStream unimplemented"))
}
