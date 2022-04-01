package provider

import (
	"context"
	"errors"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/pulumi/pulumi-google-hybrid/provider/pkg/gen"
	"github.com/pulumi/pulumi-google-hybrid/provider/pkg/resources"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	rpc "github.com/pulumi/pulumi/sdk/v3/proto/go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type demuxProvider struct {
	schemaBytes     []byte
	version         string
	metadata        *resources.CloudAPIMetadata
	bridgedProvider rpc.ResourceProviderServer
	nativeProvider  rpc.ResourceProviderServer
}

func (f demuxProvider) Call(ctx context.Context, request *rpc.CallRequest) (*rpc.CallResponse, error) {
	return nil, status.Error(codes.Unimplemented, "Call is not yet implemented")
}

func (f demuxProvider) Cancel(ctx context.Context, e *empty.Empty) (*empty.Empty, error) {
	return &empty.Empty{}, nil
}

func (f demuxProvider) GetSchema(ctx context.Context, request *rpc.GetSchemaRequest) (*rpc.GetSchemaResponse, error) {
	decompressed, err := gen.DecompressSchema(f.schemaBytes)
	if err != nil {
		return nil, errors.New("failure loading schema")
	}
	return &rpc.GetSchemaResponse{Schema: string(decompressed)}, nil
}

func (f demuxProvider) CheckConfig(ctx context.Context, request *rpc.CheckRequest) (*rpc.CheckResponse, error) {
	return &rpc.CheckResponse{Inputs: request.GetNews()}, nil
}

func (f demuxProvider) DiffConfig(ctx context.Context, request *rpc.DiffRequest) (*rpc.DiffResponse, error) {
	return &rpc.DiffResponse{
		Changes:             0,
		Replaces:            []string{},
		Stables:             []string{},
		DeleteBeforeReplace: false,
	}, nil
}

func (f demuxProvider) Configure(ctx context.Context, request *rpc.ConfigureRequest) (*rpc.ConfigureResponse, error) {
	var myResp rpc.ConfigureResponse
	for _, prov := range []rpc.ResourceProviderServer{f.bridgedProvider, f.nativeProvider} {
		resp, err := prov.Configure(ctx, request)
		if err != nil {
			return nil, err
		}
		myResp = *resp
	}
	return &myResp, nil
}

func (f demuxProvider) Invoke(ctx context.Context, request *rpc.InvokeRequest) (*rpc.InvokeResponse, error) {
	if _, ok := f.metadata.Functions[request.Tok]; ok {
		return f.nativeProvider.Invoke(ctx, request)
	}
	return f.bridgedProvider.Invoke(ctx, request)
}

func (f demuxProvider) StreamInvoke(request *rpc.InvokeRequest, server rpc.ResourceProvider_StreamInvokeServer) error {
	return status.Error(codes.Unimplemented, "StreamInvoke is not yet implemented")
}

func (f demuxProvider) Check(ctx context.Context, request *rpc.CheckRequest) (*rpc.CheckResponse, error) {
	urn := resource.URN(request.GetUrn())
	tok := urn.Type().String()
	if _, isNative := f.metadata.Resources[tok]; isNative {
		return f.nativeProvider.Check(ctx, request)
	}
	return f.bridgedProvider.Check(ctx, request)
}

func (f demuxProvider) Diff(ctx context.Context, request *rpc.DiffRequest) (*rpc.DiffResponse, error) {
	urn := resource.URN(request.GetUrn())
	tok := urn.Type().String()
	if _, isNative := f.metadata.Resources[tok]; isNative {
		return f.nativeProvider.Diff(ctx, request)
	}
	return f.bridgedProvider.Diff(ctx, request)
}

func (f demuxProvider) Create(ctx context.Context, request *rpc.CreateRequest) (*rpc.CreateResponse, error) {
	urn := resource.URN(request.GetUrn())
	tok := urn.Type().String()
	if _, isNative := f.metadata.Resources[tok]; isNative {
		return f.nativeProvider.Create(ctx, request)
	}
	return f.bridgedProvider.Create(ctx, request)
}

func (f demuxProvider) Read(ctx context.Context, request *rpc.ReadRequest) (*rpc.ReadResponse, error) {
	urn := resource.URN(request.GetUrn())
	tok := urn.Type().String()
	if _, isNative := f.metadata.Resources[tok]; isNative {
		return f.nativeProvider.Read(ctx, request)
	}
	return f.bridgedProvider.Read(ctx, request)
}

func (f demuxProvider) Update(ctx context.Context, request *rpc.UpdateRequest) (*rpc.UpdateResponse, error) {
	urn := resource.URN(request.GetUrn())
	tok := urn.Type().String()
	if _, isNative := f.metadata.Resources[tok]; isNative {
		return f.nativeProvider.Update(ctx, request)
	}
	return f.bridgedProvider.Update(ctx, request)
}

func (f demuxProvider) Delete(ctx context.Context, request *rpc.DeleteRequest) (*empty.Empty, error) {
	urn := resource.URN(request.GetUrn())
	tok := urn.Type().String()
	if _, isNative := f.metadata.Resources[tok]; isNative {
		return f.nativeProvider.Delete(ctx, request)
	}
	return f.bridgedProvider.Delete(ctx, request)
}

func (f demuxProvider) Construct(ctx context.Context, request *rpc.ConstructRequest) (*rpc.ConstructResponse, error) {
	return nil, status.Error(codes.Unimplemented, "Construct is not yet implemented")
}

func (f demuxProvider) GetPluginInfo(ctx context.Context, empty *empty.Empty) (*rpc.PluginInfo, error) {
	return &rpc.PluginInfo{
		Version: f.version,
	}, nil
}
