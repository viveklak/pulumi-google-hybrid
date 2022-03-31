// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets details of a single GrpcRoute.
func LookupGrpcRoute(ctx *pulumi.Context, args *LookupGrpcRouteArgs, opts ...pulumi.InvokeOption) (*LookupGrpcRouteResult, error) {
	var rv LookupGrpcRouteResult
	err := ctx.Invoke("google-native:networkservices/v1beta1:getGrpcRoute", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupGrpcRouteArgs struct {
	GrpcRouteId string  `pulumi:"grpcRouteId"`
	Location    string  `pulumi:"location"`
	Project     *string `pulumi:"project"`
}

type LookupGrpcRouteResult struct {
	// The timestamp when the resource was created.
	CreateTime string `pulumi:"createTime"`
	// Optional. A free-text description of the resource. Max length 1024 characters.
	Description string `pulumi:"description"`
	// Optional. Gateways defines a list of gateways this GrpcRoute is attached to, as one of the routing rules to route the requests served by the gateway. Each gateway reference should match the pattern: `projects/*/locations/global/gateways/`
	Gateways []string `pulumi:"gateways"`
	// Service hostnames with an optional port for which this route describes traffic. Format: [:] Hostname is the fully qualified domain name of a network host. This matches the RFC 1123 definition of a hostname with 2 notable exceptions: - IPs are not allowed. - A hostname may be prefixed with a wildcard label (*.). The wildcard label must appear by itself as the first label. Hostname can be "precise" which is a domain name without the terminating dot of a network host (e.g. "foo.example.com") or "wildcard", which is a domain name prefixed with a single wildcard label (e.g. *.example.com). Note that as per RFC1035 and RFC1123, a label must consist of lower case alphanumeric characters or '-', and must start and end with an alphanumeric character. No other punctuation is allowed. The routes associated with a Mesh or Gateway must have unique hostnames. If you attempt to attach multiple routes with conflicting hostnames, the configuration will be rejected. For example, while it is acceptable for routes for the hostnames "*.foo.bar.com" and "*.bar.com" to be associated with the same route, it is not possible to associate two routes both with "*.bar.com" or both with "bar.com". If a port is specified, then gRPC clients must use the channel URI with the port to match this rule (i.e. "xds:///service:123"), otherwise they must supply the URI without a port (i.e. "xds:///service").
	Hostnames []string `pulumi:"hostnames"`
	// Optional. Set of label tags associated with the GrpcRoute resource.
	Labels map[string]string `pulumi:"labels"`
	// Optional. Meshes defines a list of meshes this GrpcRoute is attached to, as one of the routing rules to route the requests served by the mesh. Each mesh reference should match the pattern: `projects/*/locations/global/meshes/`
	Meshes []string `pulumi:"meshes"`
	// Name of the GrpcRoute resource. It matches pattern `projects/*/locations/global/grpcRoutes/`
	Name string `pulumi:"name"`
	// A list of detailed rules defining how to route traffic. Within a single GrpcRoute, the GrpcRoute.RouteAction associated with the first matching GrpcRoute.RouteRule will be executed. At least one rule must be supplied.
	Rules []GrpcRouteRouteRuleResponse `pulumi:"rules"`
	// Server-defined URL of this resource
	SelfLink string `pulumi:"selfLink"`
	// The timestamp when the resource was updated.
	UpdateTime string `pulumi:"updateTime"`
}

func LookupGrpcRouteOutput(ctx *pulumi.Context, args LookupGrpcRouteOutputArgs, opts ...pulumi.InvokeOption) LookupGrpcRouteResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupGrpcRouteResult, error) {
			args := v.(LookupGrpcRouteArgs)
			r, err := LookupGrpcRoute(ctx, &args, opts...)
			return *r, err
		}).(LookupGrpcRouteResultOutput)
}

type LookupGrpcRouteOutputArgs struct {
	GrpcRouteId pulumi.StringInput    `pulumi:"grpcRouteId"`
	Location    pulumi.StringInput    `pulumi:"location"`
	Project     pulumi.StringPtrInput `pulumi:"project"`
}

func (LookupGrpcRouteOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGrpcRouteArgs)(nil)).Elem()
}

type LookupGrpcRouteResultOutput struct{ *pulumi.OutputState }

func (LookupGrpcRouteResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGrpcRouteResult)(nil)).Elem()
}

func (o LookupGrpcRouteResultOutput) ToLookupGrpcRouteResultOutput() LookupGrpcRouteResultOutput {
	return o
}

func (o LookupGrpcRouteResultOutput) ToLookupGrpcRouteResultOutputWithContext(ctx context.Context) LookupGrpcRouteResultOutput {
	return o
}

// The timestamp when the resource was created.
func (o LookupGrpcRouteResultOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGrpcRouteResult) string { return v.CreateTime }).(pulumi.StringOutput)
}

// Optional. A free-text description of the resource. Max length 1024 characters.
func (o LookupGrpcRouteResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGrpcRouteResult) string { return v.Description }).(pulumi.StringOutput)
}

// Optional. Gateways defines a list of gateways this GrpcRoute is attached to, as one of the routing rules to route the requests served by the gateway. Each gateway reference should match the pattern: `projects/*/locations/global/gateways/`
func (o LookupGrpcRouteResultOutput) Gateways() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupGrpcRouteResult) []string { return v.Gateways }).(pulumi.StringArrayOutput)
}

// Service hostnames with an optional port for which this route describes traffic. Format: [:] Hostname is the fully qualified domain name of a network host. This matches the RFC 1123 definition of a hostname with 2 notable exceptions: - IPs are not allowed. - A hostname may be prefixed with a wildcard label (*.). The wildcard label must appear by itself as the first label. Hostname can be "precise" which is a domain name without the terminating dot of a network host (e.g. "foo.example.com") or "wildcard", which is a domain name prefixed with a single wildcard label (e.g. *.example.com). Note that as per RFC1035 and RFC1123, a label must consist of lower case alphanumeric characters or '-', and must start and end with an alphanumeric character. No other punctuation is allowed. The routes associated with a Mesh or Gateway must have unique hostnames. If you attempt to attach multiple routes with conflicting hostnames, the configuration will be rejected. For example, while it is acceptable for routes for the hostnames "*.foo.bar.com" and "*.bar.com" to be associated with the same route, it is not possible to associate two routes both with "*.bar.com" or both with "bar.com". If a port is specified, then gRPC clients must use the channel URI with the port to match this rule (i.e. "xds:///service:123"), otherwise they must supply the URI without a port (i.e. "xds:///service").
func (o LookupGrpcRouteResultOutput) Hostnames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupGrpcRouteResult) []string { return v.Hostnames }).(pulumi.StringArrayOutput)
}

// Optional. Set of label tags associated with the GrpcRoute resource.
func (o LookupGrpcRouteResultOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupGrpcRouteResult) map[string]string { return v.Labels }).(pulumi.StringMapOutput)
}

// Optional. Meshes defines a list of meshes this GrpcRoute is attached to, as one of the routing rules to route the requests served by the mesh. Each mesh reference should match the pattern: `projects/*/locations/global/meshes/`
func (o LookupGrpcRouteResultOutput) Meshes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupGrpcRouteResult) []string { return v.Meshes }).(pulumi.StringArrayOutput)
}

// Name of the GrpcRoute resource. It matches pattern `projects/*/locations/global/grpcRoutes/`
func (o LookupGrpcRouteResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGrpcRouteResult) string { return v.Name }).(pulumi.StringOutput)
}

// A list of detailed rules defining how to route traffic. Within a single GrpcRoute, the GrpcRoute.RouteAction associated with the first matching GrpcRoute.RouteRule will be executed. At least one rule must be supplied.
func (o LookupGrpcRouteResultOutput) Rules() GrpcRouteRouteRuleResponseArrayOutput {
	return o.ApplyT(func(v LookupGrpcRouteResult) []GrpcRouteRouteRuleResponse { return v.Rules }).(GrpcRouteRouteRuleResponseArrayOutput)
}

// Server-defined URL of this resource
func (o LookupGrpcRouteResultOutput) SelfLink() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGrpcRouteResult) string { return v.SelfLink }).(pulumi.StringOutput)
}

// The timestamp when the resource was updated.
func (o LookupGrpcRouteResultOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGrpcRouteResult) string { return v.UpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupGrpcRouteResultOutput{})
}
