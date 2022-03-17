// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1alpha1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this method to get details about a stream.
func LookupStream(ctx *pulumi.Context, args *LookupStreamArgs, opts ...pulumi.InvokeOption) (*LookupStreamResult, error) {
	var rv LookupStreamResult
	err := ctx.Invoke("google-native:datastream/v1alpha1:getStream", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupStreamArgs struct {
	Location string  `pulumi:"location"`
	Project  *string `pulumi:"project"`
	StreamId string  `pulumi:"streamId"`
}

type LookupStreamResult struct {
	// Automatically backfill objects included in the stream source configuration. Specific objects can be excluded.
	BackfillAll BackfillAllStrategyResponse `pulumi:"backfillAll"`
	// Do not automatically backfill any objects.
	BackfillNone BackfillNoneStrategyResponse `pulumi:"backfillNone"`
	// The creation time of the stream.
	CreateTime string `pulumi:"createTime"`
	// Immutable. A reference to a KMS encryption key. If provided, it will be used to encrypt the data. If left blank, data will be encrypted using an internal Stream-specific encryption key provisioned through KMS.
	CustomerManagedEncryptionKey string `pulumi:"customerManagedEncryptionKey"`
	// Destination connection profile configuration.
	DestinationConfig DestinationConfigResponse `pulumi:"destinationConfig"`
	// Display name.
	DisplayName string `pulumi:"displayName"`
	// Errors on the Stream.
	Errors []ErrorResponse `pulumi:"errors"`
	// Labels.
	Labels map[string]string `pulumi:"labels"`
	// The stream's name.
	Name string `pulumi:"name"`
	// Source connection profile configuration.
	SourceConfig SourceConfigResponse `pulumi:"sourceConfig"`
	// The state of the stream.
	State string `pulumi:"state"`
	// The last update time of the stream.
	UpdateTime string `pulumi:"updateTime"`
}

func LookupStreamOutput(ctx *pulumi.Context, args LookupStreamOutputArgs, opts ...pulumi.InvokeOption) LookupStreamResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupStreamResult, error) {
			args := v.(LookupStreamArgs)
			r, err := LookupStream(ctx, &args, opts...)
			return *r, err
		}).(LookupStreamResultOutput)
}

type LookupStreamOutputArgs struct {
	Location pulumi.StringInput    `pulumi:"location"`
	Project  pulumi.StringPtrInput `pulumi:"project"`
	StreamId pulumi.StringInput    `pulumi:"streamId"`
}

func (LookupStreamOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStreamArgs)(nil)).Elem()
}

type LookupStreamResultOutput struct{ *pulumi.OutputState }

func (LookupStreamResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStreamResult)(nil)).Elem()
}

func (o LookupStreamResultOutput) ToLookupStreamResultOutput() LookupStreamResultOutput {
	return o
}

func (o LookupStreamResultOutput) ToLookupStreamResultOutputWithContext(ctx context.Context) LookupStreamResultOutput {
	return o
}

// Automatically backfill objects included in the stream source configuration. Specific objects can be excluded.
func (o LookupStreamResultOutput) BackfillAll() BackfillAllStrategyResponseOutput {
	return o.ApplyT(func(v LookupStreamResult) BackfillAllStrategyResponse { return v.BackfillAll }).(BackfillAllStrategyResponseOutput)
}

// Do not automatically backfill any objects.
func (o LookupStreamResultOutput) BackfillNone() BackfillNoneStrategyResponseOutput {
	return o.ApplyT(func(v LookupStreamResult) BackfillNoneStrategyResponse { return v.BackfillNone }).(BackfillNoneStrategyResponseOutput)
}

// The creation time of the stream.
func (o LookupStreamResultOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStreamResult) string { return v.CreateTime }).(pulumi.StringOutput)
}

// Immutable. A reference to a KMS encryption key. If provided, it will be used to encrypt the data. If left blank, data will be encrypted using an internal Stream-specific encryption key provisioned through KMS.
func (o LookupStreamResultOutput) CustomerManagedEncryptionKey() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStreamResult) string { return v.CustomerManagedEncryptionKey }).(pulumi.StringOutput)
}

// Destination connection profile configuration.
func (o LookupStreamResultOutput) DestinationConfig() DestinationConfigResponseOutput {
	return o.ApplyT(func(v LookupStreamResult) DestinationConfigResponse { return v.DestinationConfig }).(DestinationConfigResponseOutput)
}

// Display name.
func (o LookupStreamResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStreamResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

// Errors on the Stream.
func (o LookupStreamResultOutput) Errors() ErrorResponseArrayOutput {
	return o.ApplyT(func(v LookupStreamResult) []ErrorResponse { return v.Errors }).(ErrorResponseArrayOutput)
}

// Labels.
func (o LookupStreamResultOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupStreamResult) map[string]string { return v.Labels }).(pulumi.StringMapOutput)
}

// The stream's name.
func (o LookupStreamResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStreamResult) string { return v.Name }).(pulumi.StringOutput)
}

// Source connection profile configuration.
func (o LookupStreamResultOutput) SourceConfig() SourceConfigResponseOutput {
	return o.ApplyT(func(v LookupStreamResult) SourceConfigResponse { return v.SourceConfig }).(SourceConfigResponseOutput)
}

// The state of the stream.
func (o LookupStreamResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStreamResult) string { return v.State }).(pulumi.StringOutput)
}

// The last update time of the stream.
func (o LookupStreamResultOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStreamResult) string { return v.UpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupStreamResultOutput{})
}
