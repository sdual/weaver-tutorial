// Code generated by "weaver generate". DO NOT EDIT.
//go:build !ignoreWeaverGen

package main

import (
	"context"
	"errors"
	"github.com/ServiceWeaver/weaver"
	"github.com/ServiceWeaver/weaver/runtime/codegen"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"reflect"
)

var _ codegen.LatestVersion = codegen.Version[[0][17]struct{}](`

ERROR: You generated this file with 'weaver generate' v0.17.0 (codegen
version v0.17.0). The generated code is incompatible with the version of the
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

func init() {
	codegen.Register(codegen.Registration{
		Name:      "github.com/ServiceWeaver/weaver/Main",
		Iface:     reflect.TypeOf((*weaver.Main)(nil)).Elem(),
		Impl:      reflect.TypeOf(app{}),
		Listeners: []string{"hello"},
		LocalStubFn: func(impl any, caller string, tracer trace.Tracer) any {
			return main_local_stub{impl: impl.(weaver.Main), tracer: tracer}
		},
		ClientStubFn: func(stub codegen.Stub, caller string) any { return main_client_stub{stub: stub} },
		ServerStubFn: func(impl any, addLoad func(uint64, float64)) codegen.Server {
			return main_server_stub{impl: impl.(weaver.Main), addLoad: addLoad}
		},
		RefData: "⟦6a8c9a04:wEaVeReDgE:github.com/ServiceWeaver/weaver/Main→hello/Reverser⟧\n⟦17f36ff9:wEaVeRlIsTeNeRs:github.com/ServiceWeaver/weaver/Main→hello⟧\n",
	})
	codegen.Register(codegen.Registration{
		Name:  "hello/Reverser",
		Iface: reflect.TypeOf((*Reverser)(nil)).Elem(),
		Impl:  reflect.TypeOf(reverser{}),
		LocalStubFn: func(impl any, caller string, tracer trace.Tracer) any {
			return reverser_local_stub{impl: impl.(Reverser), tracer: tracer, reverseMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "hello/Reverser", Method: "Reverse", Remote: false})}
		},
		ClientStubFn: func(stub codegen.Stub, caller string) any {
			return reverser_client_stub{stub: stub, reverseMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "hello/Reverser", Method: "Reverse", Remote: true})}
		},
		ServerStubFn: func(impl any, addLoad func(uint64, float64)) codegen.Server {
			return reverser_server_stub{impl: impl.(Reverser), addLoad: addLoad}
		},
		RefData: "",
	})
}

// weaver.Instance checks.
var _ weaver.InstanceOf[weaver.Main] = (*app)(nil)
var _ weaver.InstanceOf[Reverser] = (*reverser)(nil)

// weaver.Router checks.
var _ weaver.Unrouted = (*app)(nil)
var _ weaver.Unrouted = (*reverser)(nil)

// Local stub implementations.

type main_local_stub struct {
	impl   weaver.Main
	tracer trace.Tracer
}

// Check that main_local_stub implements the weaver.Main interface.
var _ weaver.Main = (*main_local_stub)(nil)

type reverser_local_stub struct {
	impl           Reverser
	tracer         trace.Tracer
	reverseMetrics *codegen.MethodMetrics
}

// Check that reverser_local_stub implements the Reverser interface.
var _ Reverser = (*reverser_local_stub)(nil)

func (s reverser_local_stub) Reverse(ctx context.Context, a0 string) (r0 string, err error) {
	// Update metrics.
	begin := s.reverseMetrics.Begin()
	defer func() { s.reverseMetrics.End(begin, err != nil, 0, 0) }()
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.tracer.Start(ctx, "main.Reverser.Reverse", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.Reverse(ctx, a0)
}

// Client stub implementations.

type main_client_stub struct {
	stub codegen.Stub
}

// Check that main_client_stub implements the weaver.Main interface.
var _ weaver.Main = (*main_client_stub)(nil)

type reverser_client_stub struct {
	stub           codegen.Stub
	reverseMetrics *codegen.MethodMetrics
}

// Check that reverser_client_stub implements the Reverser interface.
var _ Reverser = (*reverser_client_stub)(nil)

func (s reverser_client_stub) Reverse(ctx context.Context, a0 string) (r0 string, err error) {
	// Update metrics.
	var requestBytes, replyBytes int
	begin := s.reverseMetrics.Begin()
	defer func() { s.reverseMetrics.End(begin, err != nil, requestBytes, replyBytes) }()

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.stub.Tracer().Start(ctx, "main.Reverser.Reverse", trace.WithSpanKind(trace.SpanKindClient))
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
	size += (4 + len(a0))
	enc := codegen.NewEncoder()
	enc.Reset(size)

	// Encode arguments.
	enc.String(a0)
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
	r0 = dec.String()
	err = dec.Error()
	return
}

// Server stub implementations.

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

type reverser_server_stub struct {
	impl    Reverser
	addLoad func(key uint64, load float64)
}

// Check that reverser_server_stub implements the codegen.Server interface.
var _ codegen.Server = (*reverser_server_stub)(nil)

// GetStubFn implements the codegen.Server interface.
func (s reverser_server_stub) GetStubFn(method string) func(ctx context.Context, args []byte) ([]byte, error) {
	switch method {
	case "Reverse":
		return s.reverse
	default:
		return nil
	}
}

func (s reverser_server_stub) reverse(ctx context.Context, args []byte) (res []byte, err error) {
	// Catch and return any panics detected during encoding/decoding/rpc.
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	// Decode arguments.
	dec := codegen.NewDecoder(args)
	var a0 string
	a0 = dec.String()

	// TODO(rgrandl): The deferred function above will recover from panics in the
	// user code: fix this.
	// Call the local method.
	r0, appErr := s.impl.Reverse(ctx, a0)

	// Encode the results.
	enc := codegen.NewEncoder()
	enc.String(r0)
	enc.Error(appErr)
	return enc.Data(), nil
}
