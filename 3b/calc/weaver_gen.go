// Code generated by "weaver generate". DO NOT EDIT.
//go:build !ignoreWeaverGen

package calc

import (
	"context"
	"errors"
	"github.com/ServiceWeaver/weaver"
	"github.com/ServiceWeaver/weaver/runtime/codegen"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"reflect"
)

func init() {
	codegen.Register(codegen.Registration{
		Name:  "github.com/mwhittaker/calculator/3b/calc/Adder",
		Iface: reflect.TypeOf((*Adder)(nil)).Elem(),
		Impl:  reflect.TypeOf(adder{}),
		LocalStubFn: func(impl any, caller string, tracer trace.Tracer) any {
			return adder_local_stub{impl: impl.(Adder), tracer: tracer, addMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "github.com/mwhittaker/calculator/3b/calc/Adder", Method: "Add", Remote: false})}
		},
		ClientStubFn: func(stub codegen.Stub, caller string) any {
			return adder_client_stub{stub: stub, addMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "github.com/mwhittaker/calculator/3b/calc/Adder", Method: "Add", Remote: true})}
		},
		ServerStubFn: func(impl any, addLoad func(uint64, float64)) codegen.Server {
			return adder_server_stub{impl: impl.(Adder), addLoad: addLoad}
		},
		ReflectStubFn: func(caller func(string, context.Context, []any, []any) error) any {
			return adder_reflect_stub{caller: caller}
		},
		RefData: "",
	})
	codegen.Register(codegen.Registration{
		Name:      "github.com/ServiceWeaver/weaver/Main",
		Iface:     reflect.TypeOf((*weaver.Main)(nil)).Elem(),
		Impl:      reflect.TypeOf(Server{}),
		Listeners: []string{"calc"},
		LocalStubFn: func(impl any, caller string, tracer trace.Tracer) any {
			return main_local_stub{impl: impl.(weaver.Main), tracer: tracer}
		},
		ClientStubFn: func(stub codegen.Stub, caller string) any { return main_client_stub{stub: stub} },
		ServerStubFn: func(impl any, addLoad func(uint64, float64)) codegen.Server {
			return main_server_stub{impl: impl.(weaver.Main), addLoad: addLoad}
		},
		ReflectStubFn: func(caller func(string, context.Context, []any, []any) error) any {
			return main_reflect_stub{caller: caller}
		},
		RefData: "⟦55739722:wEaVeReDgE:github.com/ServiceWeaver/weaver/Main→github.com/mwhittaker/calculator/3b/calc/Adder⟧\n⟦41990618:wEaVeRlIsTeNeRs:github.com/ServiceWeaver/weaver/Main→calc⟧\n",
	})
}

// weaver.InstanceOf checks.
var _ weaver.InstanceOf[Adder] = (*adder)(nil)
var _ weaver.InstanceOf[weaver.Main] = (*Server)(nil)

// weaver.Router checks.
var _ weaver.Unrouted = (*adder)(nil)
var _ weaver.Unrouted = (*Server)(nil)

// Local stub implementations.

type adder_local_stub struct {
	impl       Adder
	tracer     trace.Tracer
	addMetrics *codegen.MethodMetrics
}

// Check that adder_local_stub implements the Adder interface.
var _ Adder = (*adder_local_stub)(nil)

func (s adder_local_stub) Add(ctx context.Context, a0 int, a1 int) (r0 int, err error) {
	// Update metrics.
	begin := s.addMetrics.Begin()
	defer func() { s.addMetrics.End(begin, err != nil, 0, 0) }()
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.tracer.Start(ctx, "calc.Adder.Add", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.Add(ctx, a0, a1)
}

type main_local_stub struct {
	impl   weaver.Main
	tracer trace.Tracer
}

// Check that main_local_stub implements the weaver.Main interface.
var _ weaver.Main = (*main_local_stub)(nil)

// Client stub implementations.

type adder_client_stub struct {
	stub       codegen.Stub
	addMetrics *codegen.MethodMetrics
}

// Check that adder_client_stub implements the Adder interface.
var _ Adder = (*adder_client_stub)(nil)

func (s adder_client_stub) Add(ctx context.Context, a0 int, a1 int) (r0 int, err error) {
	// Update metrics.
	var requestBytes, replyBytes int
	begin := s.addMetrics.Begin()
	defer func() { s.addMetrics.End(begin, err != nil, requestBytes, replyBytes) }()

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.stub.Tracer().Start(ctx, "calc.Adder.Add", trace.WithSpanKind(trace.SpanKindClient))
	}

	defer func() {
		// Catch and return any panics detected during encoding/decoding/rpc.
		if err == nil {
			err = codegen.CatchPanics(recover())
			if err != nil {
				err = errors.Join(weaver.RemoteCallError, err)
			}
		}

		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
		}
		span.End()

	}()

	// Preallocate a buffer of the right size.
	size := 0
	size += 8
	size += 8
	enc := codegen.NewEncoder()
	enc.Reset(size)

	// Encode arguments.
	enc.Int(a0)
	enc.Int(a1)
	var shardKey uint64

	// Call the remote method.
	requestBytes = len(enc.Data())
	var results []byte
	results, err = s.stub.Run(ctx, 0, enc.Data(), shardKey)
	replyBytes = len(results)
	if err != nil {
		err = errors.Join(weaver.RemoteCallError, err)
		return
	}

	// Decode the results.
	dec := codegen.NewDecoder(results)
	r0 = dec.Int()
	err = dec.Error()
	return
}

type main_client_stub struct {
	stub codegen.Stub
}

// Check that main_client_stub implements the weaver.Main interface.
var _ weaver.Main = (*main_client_stub)(nil)

// Note that "weaver generate" will always generate the error message below.
// Everything is okay. The error message is only relevant if you see it when
// you run "go build" or "go run".
var _ codegen.LatestVersion = codegen.Version[[0][20]struct{}](`

ERROR: You generated this file with 'weaver generate' v0.21.3-0.20231005224231-991d5ec5bce4 (codegen
version v0.20.0). The generated code is incompatible with the version of the
github.com/ServiceWeaver/weaver module that you're using. The weaver module
version can be found in your go.mod file or by running the following command.

    go list -m github.com/ServiceWeaver/weaver

We recommend updating the weaver module and the 'weaver generate' command by
running the following.

    go get github.com/ServiceWeaver/weaver@latest
    go install github.com/ServiceWeaver/weaver/cmd/weaver@latest

Then, re-run 'weaver generate' and re-build your code. If the problem persists,
please file an issue at https://github.com/ServiceWeaver/weaver/issues.

`)

// Server stub implementations.

type adder_server_stub struct {
	impl    Adder
	addLoad func(key uint64, load float64)
}

// Check that adder_server_stub implements the codegen.Server interface.
var _ codegen.Server = (*adder_server_stub)(nil)

// GetStubFn implements the codegen.Server interface.
func (s adder_server_stub) GetStubFn(method string) func(ctx context.Context, args []byte) ([]byte, error) {
	switch method {
	case "Add":
		return s.add
	default:
		return nil
	}
}

func (s adder_server_stub) add(ctx context.Context, args []byte) (res []byte, err error) {
	// Catch and return any panics detected during encoding/decoding/rpc.
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	// Decode arguments.
	dec := codegen.NewDecoder(args)
	var a0 int
	a0 = dec.Int()
	var a1 int
	a1 = dec.Int()

	// TODO(rgrandl): The deferred function above will recover from panics in the
	// user code: fix this.
	// Call the local method.
	r0, appErr := s.impl.Add(ctx, a0, a1)

	// Encode the results.
	enc := codegen.NewEncoder()
	enc.Int(r0)
	enc.Error(appErr)
	return enc.Data(), nil
}

type main_server_stub struct {
	impl    weaver.Main
	addLoad func(key uint64, load float64)
}

// Check that main_server_stub implements the codegen.Server interface.
var _ codegen.Server = (*main_server_stub)(nil)

// GetStubFn implements the codegen.Server interface.
func (s main_server_stub) GetStubFn(method string) func(ctx context.Context, args []byte) ([]byte, error) {
	switch method {
	default:
		return nil
	}
}

// Reflect stub implementations.

type adder_reflect_stub struct {
	caller func(string, context.Context, []any, []any) error
}

// Check that adder_reflect_stub implements the Adder interface.
var _ Adder = (*adder_reflect_stub)(nil)

func (s adder_reflect_stub) Add(ctx context.Context, a0 int, a1 int) (r0 int, err error) {
	err = s.caller("Add", ctx, []any{a0, a1}, []any{&r0})
	return
}

type main_reflect_stub struct {
	caller func(string, context.Context, []any, []any) error
}

// Check that main_reflect_stub implements the weaver.Main interface.
var _ weaver.Main = (*main_reflect_stub)(nil)

