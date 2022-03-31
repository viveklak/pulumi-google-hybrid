// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets details of a single Mesh.
func LookupMesh(ctx *pulumi.Context, args *LookupMeshArgs, opts ...pulumi.InvokeOption) (*LookupMeshResult, error) {
	var rv LookupMeshResult
	err := ctx.Invoke("google-native:networkservices/v1beta1:getMesh", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupMeshArgs struct {
	Location string  `pulumi:"location"`
	MeshId   string  `pulumi:"meshId"`
	Project  *string `pulumi:"project"`
}

type LookupMeshResult struct {
	// The timestamp when the resource was created.
	CreateTime string `pulumi:"createTime"`
	// Optional. A free-text description of the resource. Max length 1024 characters.
	Description string `pulumi:"description"`
	// Optional. If set to a valid TCP port (1-65535), instructs the SIDECAR proxy to listen on the specified port of localhost (127.0.0.1) address. The SIDECAR proxy will expect all traffic to be redirected to this port regardless of its actual ip:port destination. If unset, a port '15001' is used as the interception port. This will is applicable only for sidecar proxy deployments.
	InterceptionPort int `pulumi:"interceptionPort"`
	// Optional. Set of label tags associated with the Mesh resource.
	Labels map[string]string `pulumi:"labels"`
	// Name of the Mesh resource. It matches pattern `projects/*/locations/global/meshes/`.
	Name string `pulumi:"name"`
	// Server-defined URL of this resource
	SelfLink string `pulumi:"selfLink"`
	// The timestamp when the resource was updated.
	UpdateTime string `pulumi:"updateTime"`
}

func LookupMeshOutput(ctx *pulumi.Context, args LookupMeshOutputArgs, opts ...pulumi.InvokeOption) LookupMeshResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupMeshResult, error) {
			args := v.(LookupMeshArgs)
			r, err := LookupMesh(ctx, &args, opts...)
			return *r, err
		}).(LookupMeshResultOutput)
}

type LookupMeshOutputArgs struct {
	Location pulumi.StringInput    `pulumi:"location"`
	MeshId   pulumi.StringInput    `pulumi:"meshId"`
	Project  pulumi.StringPtrInput `pulumi:"project"`
}

func (LookupMeshOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMeshArgs)(nil)).Elem()
}

type LookupMeshResultOutput struct{ *pulumi.OutputState }

func (LookupMeshResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMeshResult)(nil)).Elem()
}

func (o LookupMeshResultOutput) ToLookupMeshResultOutput() LookupMeshResultOutput {
	return o
}

func (o LookupMeshResultOutput) ToLookupMeshResultOutputWithContext(ctx context.Context) LookupMeshResultOutput {
	return o
}

// The timestamp when the resource was created.
func (o LookupMeshResultOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMeshResult) string { return v.CreateTime }).(pulumi.StringOutput)
}

// Optional. A free-text description of the resource. Max length 1024 characters.
func (o LookupMeshResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMeshResult) string { return v.Description }).(pulumi.StringOutput)
}

// Optional. If set to a valid TCP port (1-65535), instructs the SIDECAR proxy to listen on the specified port of localhost (127.0.0.1) address. The SIDECAR proxy will expect all traffic to be redirected to this port regardless of its actual ip:port destination. If unset, a port '15001' is used as the interception port. This will is applicable only for sidecar proxy deployments.
func (o LookupMeshResultOutput) InterceptionPort() pulumi.IntOutput {
	return o.ApplyT(func(v LookupMeshResult) int { return v.InterceptionPort }).(pulumi.IntOutput)
}

// Optional. Set of label tags associated with the Mesh resource.
func (o LookupMeshResultOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupMeshResult) map[string]string { return v.Labels }).(pulumi.StringMapOutput)
}

// Name of the Mesh resource. It matches pattern `projects/*/locations/global/meshes/`.
func (o LookupMeshResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMeshResult) string { return v.Name }).(pulumi.StringOutput)
}

// Server-defined URL of this resource
func (o LookupMeshResultOutput) SelfLink() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMeshResult) string { return v.SelfLink }).(pulumi.StringOutput)
}

// The timestamp when the resource was updated.
func (o LookupMeshResultOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMeshResult) string { return v.UpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupMeshResultOutput{})
}
