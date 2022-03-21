// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this method to get details about a private connectivity configuration.
func LookupPrivateConnection(ctx *pulumi.Context, args *LookupPrivateConnectionArgs, opts ...pulumi.InvokeOption) (*LookupPrivateConnectionResult, error) {
	var rv LookupPrivateConnectionResult
	err := ctx.Invoke("google-native:datastream/v1:getPrivateConnection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPrivateConnectionArgs struct {
	Location            string  `pulumi:"location"`
	PrivateConnectionId string  `pulumi:"privateConnectionId"`
	Project             *string `pulumi:"project"`
}

type LookupPrivateConnectionResult struct {
	// The create time of the resource.
	CreateTime string `pulumi:"createTime"`
	// Display name.
	DisplayName string `pulumi:"displayName"`
	// In case of error, the details of the error in a user-friendly format.
	Error ErrorResponse `pulumi:"error"`
	// Labels.
	Labels map[string]string `pulumi:"labels"`
	// The resource's name.
	Name string `pulumi:"name"`
	// The state of the Private Connection.
	State string `pulumi:"state"`
	// The update time of the resource.
	UpdateTime string `pulumi:"updateTime"`
	// VPC Peering Config.
	VpcPeeringConfig VpcPeeringConfigResponse `pulumi:"vpcPeeringConfig"`
}

func LookupPrivateConnectionOutput(ctx *pulumi.Context, args LookupPrivateConnectionOutputArgs, opts ...pulumi.InvokeOption) LookupPrivateConnectionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPrivateConnectionResult, error) {
			args := v.(LookupPrivateConnectionArgs)
			r, err := LookupPrivateConnection(ctx, &args, opts...)
			return *r, err
		}).(LookupPrivateConnectionResultOutput)
}

type LookupPrivateConnectionOutputArgs struct {
	Location            pulumi.StringInput    `pulumi:"location"`
	PrivateConnectionId pulumi.StringInput    `pulumi:"privateConnectionId"`
	Project             pulumi.StringPtrInput `pulumi:"project"`
}

func (LookupPrivateConnectionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPrivateConnectionArgs)(nil)).Elem()
}

type LookupPrivateConnectionResultOutput struct{ *pulumi.OutputState }

func (LookupPrivateConnectionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPrivateConnectionResult)(nil)).Elem()
}

func (o LookupPrivateConnectionResultOutput) ToLookupPrivateConnectionResultOutput() LookupPrivateConnectionResultOutput {
	return o
}

func (o LookupPrivateConnectionResultOutput) ToLookupPrivateConnectionResultOutputWithContext(ctx context.Context) LookupPrivateConnectionResultOutput {
	return o
}

// The create time of the resource.
func (o LookupPrivateConnectionResultOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateConnectionResult) string { return v.CreateTime }).(pulumi.StringOutput)
}

// Display name.
func (o LookupPrivateConnectionResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateConnectionResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

// In case of error, the details of the error in a user-friendly format.
func (o LookupPrivateConnectionResultOutput) Error() ErrorResponseOutput {
	return o.ApplyT(func(v LookupPrivateConnectionResult) ErrorResponse { return v.Error }).(ErrorResponseOutput)
}

// Labels.
func (o LookupPrivateConnectionResultOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupPrivateConnectionResult) map[string]string { return v.Labels }).(pulumi.StringMapOutput)
}

// The resource's name.
func (o LookupPrivateConnectionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateConnectionResult) string { return v.Name }).(pulumi.StringOutput)
}

// The state of the Private Connection.
func (o LookupPrivateConnectionResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateConnectionResult) string { return v.State }).(pulumi.StringOutput)
}

// The update time of the resource.
func (o LookupPrivateConnectionResultOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateConnectionResult) string { return v.UpdateTime }).(pulumi.StringOutput)
}

// VPC Peering Config.
func (o LookupPrivateConnectionResultOutput) VpcPeeringConfig() VpcPeeringConfigResponseOutput {
	return o.ApplyT(func(v LookupPrivateConnectionResult) VpcPeeringConfigResponse { return v.VpcPeeringConfig }).(VpcPeeringConfigResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPrivateConnectionResultOutput{})
}
