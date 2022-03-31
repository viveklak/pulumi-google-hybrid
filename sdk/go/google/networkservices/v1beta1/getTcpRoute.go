// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets details of a single TcpRoute.
func LookupTcpRoute(ctx *pulumi.Context, args *LookupTcpRouteArgs, opts ...pulumi.InvokeOption) (*LookupTcpRouteResult, error) {
	var rv LookupTcpRouteResult
	err := ctx.Invoke("google-native:networkservices/v1beta1:getTcpRoute", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupTcpRouteArgs struct {
	Location   string  `pulumi:"location"`
	Project    *string `pulumi:"project"`
	TcpRouteId string  `pulumi:"tcpRouteId"`
}

type LookupTcpRouteResult struct {
	// The timestamp when the resource was created.
	CreateTime string `pulumi:"createTime"`
	// Optional. A free-text description of the resource. Max length 1024 characters.
	Description string `pulumi:"description"`
	// Optional. Gateways defines a list of gateways this TcpRoute is attached to, as one of the routing rules to route the requests served by the gateway. Each gateway reference should match the pattern: `projects/*/locations/global/gateways/`
	Gateways []string `pulumi:"gateways"`
	// Optional. Set of label tags associated with the TcpRoute resource.
	Labels map[string]string `pulumi:"labels"`
	// Optional. Meshes defines a list of meshes this TcpRoute is attached to, as one of the routing rules to route the requests served by the mesh. Each mesh reference should match the pattern: `projects/*/locations/global/meshes/` The attached Mesh should be of a type SIDECAR
	Meshes []string `pulumi:"meshes"`
	// Name of the TcpRoute resource. It matches pattern `projects/*/locations/global/tcpRoutes/tcp_route_name>`.
	Name string `pulumi:"name"`
	// Rules that define how traffic is routed and handled. At least one RouteRule must be supplied. If there are multiple rules then the action taken will be the first rule to match.
	Rules []TcpRouteRouteRuleResponse `pulumi:"rules"`
	// Server-defined URL of this resource
	SelfLink string `pulumi:"selfLink"`
	// The timestamp when the resource was updated.
	UpdateTime string `pulumi:"updateTime"`
}

func LookupTcpRouteOutput(ctx *pulumi.Context, args LookupTcpRouteOutputArgs, opts ...pulumi.InvokeOption) LookupTcpRouteResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupTcpRouteResult, error) {
			args := v.(LookupTcpRouteArgs)
			r, err := LookupTcpRoute(ctx, &args, opts...)
			return *r, err
		}).(LookupTcpRouteResultOutput)
}

type LookupTcpRouteOutputArgs struct {
	Location   pulumi.StringInput    `pulumi:"location"`
	Project    pulumi.StringPtrInput `pulumi:"project"`
	TcpRouteId pulumi.StringInput    `pulumi:"tcpRouteId"`
}

func (LookupTcpRouteOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTcpRouteArgs)(nil)).Elem()
}

type LookupTcpRouteResultOutput struct{ *pulumi.OutputState }

func (LookupTcpRouteResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTcpRouteResult)(nil)).Elem()
}

func (o LookupTcpRouteResultOutput) ToLookupTcpRouteResultOutput() LookupTcpRouteResultOutput {
	return o
}

func (o LookupTcpRouteResultOutput) ToLookupTcpRouteResultOutputWithContext(ctx context.Context) LookupTcpRouteResultOutput {
	return o
}

// The timestamp when the resource was created.
func (o LookupTcpRouteResultOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTcpRouteResult) string { return v.CreateTime }).(pulumi.StringOutput)
}

// Optional. A free-text description of the resource. Max length 1024 characters.
func (o LookupTcpRouteResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTcpRouteResult) string { return v.Description }).(pulumi.StringOutput)
}

// Optional. Gateways defines a list of gateways this TcpRoute is attached to, as one of the routing rules to route the requests served by the gateway. Each gateway reference should match the pattern: `projects/*/locations/global/gateways/`
func (o LookupTcpRouteResultOutput) Gateways() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupTcpRouteResult) []string { return v.Gateways }).(pulumi.StringArrayOutput)
}

// Optional. Set of label tags associated with the TcpRoute resource.
func (o LookupTcpRouteResultOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupTcpRouteResult) map[string]string { return v.Labels }).(pulumi.StringMapOutput)
}

// Optional. Meshes defines a list of meshes this TcpRoute is attached to, as one of the routing rules to route the requests served by the mesh. Each mesh reference should match the pattern: `projects/*/locations/global/meshes/` The attached Mesh should be of a type SIDECAR
func (o LookupTcpRouteResultOutput) Meshes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupTcpRouteResult) []string { return v.Meshes }).(pulumi.StringArrayOutput)
}

// Name of the TcpRoute resource. It matches pattern `projects/*/locations/global/tcpRoutes/tcp_route_name>`.
func (o LookupTcpRouteResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTcpRouteResult) string { return v.Name }).(pulumi.StringOutput)
}

// Rules that define how traffic is routed and handled. At least one RouteRule must be supplied. If there are multiple rules then the action taken will be the first rule to match.
func (o LookupTcpRouteResultOutput) Rules() TcpRouteRouteRuleResponseArrayOutput {
	return o.ApplyT(func(v LookupTcpRouteResult) []TcpRouteRouteRuleResponse { return v.Rules }).(TcpRouteRouteRuleResponseArrayOutput)
}

// Server-defined URL of this resource
func (o LookupTcpRouteResultOutput) SelfLink() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTcpRouteResult) string { return v.SelfLink }).(pulumi.StringOutput)
}

// The timestamp when the resource was updated.
func (o LookupTcpRouteResultOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTcpRouteResult) string { return v.UpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupTcpRouteResultOutput{})
}
