// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: proto/mysqlIssueDemo.proto

/*
Package mysql_issue_demo is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package mysql_issue_demo

import (
	"context"
	"io"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/v2/utilities"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

// Suppress "imported and not used" errors
var _ codes.Code
var _ io.Reader
var _ status.Status
var _ = runtime.String
var _ = utilities.NewDoubleArray
var _ = metadata.Join

func request_MysqlIssueDemo_Write_0(ctx context.Context, marshaler runtime.Marshaler, client MysqlIssueDemoClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq WriteRequest
	var metadata runtime.ServerMetadata

	msg, err := client.Write(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_MysqlIssueDemo_Write_0(ctx context.Context, marshaler runtime.Marshaler, server MysqlIssueDemoServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq WriteRequest
	var metadata runtime.ServerMetadata

	msg, err := server.Write(ctx, &protoReq)
	return msg, metadata, err

}

func request_MysqlIssueDemo_GetAllStores_0(ctx context.Context, marshaler runtime.Marshaler, client MysqlIssueDemoClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq GetAllStoresRequest
	var metadata runtime.ServerMetadata

	msg, err := client.GetAllStores(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_MysqlIssueDemo_GetAllStores_0(ctx context.Context, marshaler runtime.Marshaler, server MysqlIssueDemoServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq GetAllStoresRequest
	var metadata runtime.ServerMetadata

	msg, err := server.GetAllStores(ctx, &protoReq)
	return msg, metadata, err

}

// RegisterMysqlIssueDemoHandlerServer registers the http handlers for service MysqlIssueDemo to "mux".
// UnaryRPC     :call MysqlIssueDemoServer directly.
// StreamingRPC :currently unsupported pending https://github.com/grpc/grpc-go/issues/906.
// Note that using this registration option will cause many gRPC library features to stop working. Consider using RegisterMysqlIssueDemoHandlerFromEndpoint instead.
func RegisterMysqlIssueDemoHandlerServer(ctx context.Context, mux *runtime.ServeMux, server MysqlIssueDemoServer) error {

	mux.Handle("POST", pattern_MysqlIssueDemo_Write_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateIncomingContext(ctx, mux, req, "/mysql_issue_demo.MysqlIssueDemo/Write", runtime.WithHTTPPathPattern("/"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_MysqlIssueDemo_Write_0(annotatedContext, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_MysqlIssueDemo_Write_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_MysqlIssueDemo_GetAllStores_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateIncomingContext(ctx, mux, req, "/mysql_issue_demo.MysqlIssueDemo/GetAllStores", runtime.WithHTTPPathPattern("/"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_MysqlIssueDemo_GetAllStores_0(annotatedContext, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_MysqlIssueDemo_GetAllStores_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

// RegisterMysqlIssueDemoHandlerFromEndpoint is same as RegisterMysqlIssueDemoHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterMysqlIssueDemoHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.DialContext(ctx, endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()

	return RegisterMysqlIssueDemoHandler(ctx, mux, conn)
}

// RegisterMysqlIssueDemoHandler registers the http handlers for service MysqlIssueDemo to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterMysqlIssueDemoHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterMysqlIssueDemoHandlerClient(ctx, mux, NewMysqlIssueDemoClient(conn))
}

// RegisterMysqlIssueDemoHandlerClient registers the http handlers for service MysqlIssueDemo
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "MysqlIssueDemoClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "MysqlIssueDemoClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "MysqlIssueDemoClient" to call the correct interceptors.
func RegisterMysqlIssueDemoHandlerClient(ctx context.Context, mux *runtime.ServeMux, client MysqlIssueDemoClient) error {

	mux.Handle("POST", pattern_MysqlIssueDemo_Write_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateContext(ctx, mux, req, "/mysql_issue_demo.MysqlIssueDemo/Write", runtime.WithHTTPPathPattern("/"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_MysqlIssueDemo_Write_0(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_MysqlIssueDemo_Write_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_MysqlIssueDemo_GetAllStores_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateContext(ctx, mux, req, "/mysql_issue_demo.MysqlIssueDemo/GetAllStores", runtime.WithHTTPPathPattern("/"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_MysqlIssueDemo_GetAllStores_0(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_MysqlIssueDemo_GetAllStores_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_MysqlIssueDemo_Write_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0}, []string{""}, ""))

	pattern_MysqlIssueDemo_GetAllStores_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0}, []string{""}, ""))
)

var (
	forward_MysqlIssueDemo_Write_0 = runtime.ForwardResponseMessage

	forward_MysqlIssueDemo_GetAllStores_0 = runtime.ForwardResponseMessage
)
