// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets details of a single DeliveryPipeline.
func LookupDeliveryPipeline(ctx *pulumi.Context, args *LookupDeliveryPipelineArgs, opts ...pulumi.InvokeOption) (*LookupDeliveryPipelineResult, error) {
	var rv LookupDeliveryPipelineResult
	err := ctx.Invoke("google-native:clouddeploy/v1:getDeliveryPipeline", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDeliveryPipelineArgs struct {
	DeliveryPipelineId string  `pulumi:"deliveryPipelineId"`
	Location           string  `pulumi:"location"`
	Project            *string `pulumi:"project"`
}

type LookupDeliveryPipelineResult struct {
	// User annotations. These attributes can only be set and used by the user, and not by Google Cloud Deploy. See https://google.aip.dev/128#annotations for more details such as format and size limitations.
	Annotations map[string]string `pulumi:"annotations"`
	// Information around the state of the Delivery Pipeline.
	Condition PipelineConditionResponse `pulumi:"condition"`
	// Time at which the pipeline was created.
	CreateTime string `pulumi:"createTime"`
	// Description of the `DeliveryPipeline`. Max length is 255 characters.
	Description string `pulumi:"description"`
	// This checksum is computed by the server based on the value of other fields, and may be sent on update and delete requests to ensure the client has an up-to-date value before proceeding.
	Etag string `pulumi:"etag"`
	// Labels are attributes that can be set and used by both the user and by Google Cloud Deploy. Labels must meet the following constraints: * Keys and values can contain only lowercase letters, numeric characters, underscores, and dashes. * All characters must use UTF-8 encoding, and international characters are allowed. * Keys must start with a lowercase letter or international character. * Each resource is limited to a maximum of 64 labels. Both keys and values are additionally constrained to be <= 128 bytes.
	Labels map[string]string `pulumi:"labels"`
	// Optional. Name of the `DeliveryPipeline`. Format is projects/{project}/ locations/{location}/deliveryPipelines/a-z{0,62}.
	Name string `pulumi:"name"`
	// SerialPipeline defines a sequential set of stages for a `DeliveryPipeline`.
	SerialPipeline SerialPipelineResponse `pulumi:"serialPipeline"`
	// Unique identifier of the `DeliveryPipeline`.
	Uid string `pulumi:"uid"`
	// Most recent time at which the pipeline was updated.
	UpdateTime string `pulumi:"updateTime"`
}

func LookupDeliveryPipelineOutput(ctx *pulumi.Context, args LookupDeliveryPipelineOutputArgs, opts ...pulumi.InvokeOption) LookupDeliveryPipelineResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDeliveryPipelineResult, error) {
			args := v.(LookupDeliveryPipelineArgs)
			r, err := LookupDeliveryPipeline(ctx, &args, opts...)
			return *r, err
		}).(LookupDeliveryPipelineResultOutput)
}

type LookupDeliveryPipelineOutputArgs struct {
	DeliveryPipelineId pulumi.StringInput    `pulumi:"deliveryPipelineId"`
	Location           pulumi.StringInput    `pulumi:"location"`
	Project            pulumi.StringPtrInput `pulumi:"project"`
}

func (LookupDeliveryPipelineOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDeliveryPipelineArgs)(nil)).Elem()
}

type LookupDeliveryPipelineResultOutput struct{ *pulumi.OutputState }

func (LookupDeliveryPipelineResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDeliveryPipelineResult)(nil)).Elem()
}

func (o LookupDeliveryPipelineResultOutput) ToLookupDeliveryPipelineResultOutput() LookupDeliveryPipelineResultOutput {
	return o
}

func (o LookupDeliveryPipelineResultOutput) ToLookupDeliveryPipelineResultOutputWithContext(ctx context.Context) LookupDeliveryPipelineResultOutput {
	return o
}

// User annotations. These attributes can only be set and used by the user, and not by Google Cloud Deploy. See https://google.aip.dev/128#annotations for more details such as format and size limitations.
func (o LookupDeliveryPipelineResultOutput) Annotations() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupDeliveryPipelineResult) map[string]string { return v.Annotations }).(pulumi.StringMapOutput)
}

// Information around the state of the Delivery Pipeline.
func (o LookupDeliveryPipelineResultOutput) Condition() PipelineConditionResponseOutput {
	return o.ApplyT(func(v LookupDeliveryPipelineResult) PipelineConditionResponse { return v.Condition }).(PipelineConditionResponseOutput)
}

// Time at which the pipeline was created.
func (o LookupDeliveryPipelineResultOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeliveryPipelineResult) string { return v.CreateTime }).(pulumi.StringOutput)
}

// Description of the `DeliveryPipeline`. Max length is 255 characters.
func (o LookupDeliveryPipelineResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeliveryPipelineResult) string { return v.Description }).(pulumi.StringOutput)
}

// This checksum is computed by the server based on the value of other fields, and may be sent on update and delete requests to ensure the client has an up-to-date value before proceeding.
func (o LookupDeliveryPipelineResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeliveryPipelineResult) string { return v.Etag }).(pulumi.StringOutput)
}

// Labels are attributes that can be set and used by both the user and by Google Cloud Deploy. Labels must meet the following constraints: * Keys and values can contain only lowercase letters, numeric characters, underscores, and dashes. * All characters must use UTF-8 encoding, and international characters are allowed. * Keys must start with a lowercase letter or international character. * Each resource is limited to a maximum of 64 labels. Both keys and values are additionally constrained to be <= 128 bytes.
func (o LookupDeliveryPipelineResultOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupDeliveryPipelineResult) map[string]string { return v.Labels }).(pulumi.StringMapOutput)
}

// Optional. Name of the `DeliveryPipeline`. Format is projects/{project}/ locations/{location}/deliveryPipelines/a-z{0,62}.
func (o LookupDeliveryPipelineResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeliveryPipelineResult) string { return v.Name }).(pulumi.StringOutput)
}

// SerialPipeline defines a sequential set of stages for a `DeliveryPipeline`.
func (o LookupDeliveryPipelineResultOutput) SerialPipeline() SerialPipelineResponseOutput {
	return o.ApplyT(func(v LookupDeliveryPipelineResult) SerialPipelineResponse { return v.SerialPipeline }).(SerialPipelineResponseOutput)
}

// Unique identifier of the `DeliveryPipeline`.
func (o LookupDeliveryPipelineResultOutput) Uid() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeliveryPipelineResult) string { return v.Uid }).(pulumi.StringOutput)
}

// Most recent time at which the pipeline was updated.
func (o LookupDeliveryPipelineResultOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeliveryPipelineResult) string { return v.UpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDeliveryPipelineResultOutput{})
}
