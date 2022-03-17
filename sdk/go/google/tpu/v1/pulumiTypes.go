// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A network endpoint over which a TPU worker can be reached.
type NetworkEndpointResponse struct {
	// The IP address of this network endpoint.
	IpAddress string `pulumi:"ipAddress"`
	// The port of this network endpoint.
	Port int `pulumi:"port"`
}

// A network endpoint over which a TPU worker can be reached.
type NetworkEndpointResponseOutput struct{ *pulumi.OutputState }

func (NetworkEndpointResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkEndpointResponse)(nil)).Elem()
}

func (o NetworkEndpointResponseOutput) ToNetworkEndpointResponseOutput() NetworkEndpointResponseOutput {
	return o
}

func (o NetworkEndpointResponseOutput) ToNetworkEndpointResponseOutputWithContext(ctx context.Context) NetworkEndpointResponseOutput {
	return o
}

// The IP address of this network endpoint.
func (o NetworkEndpointResponseOutput) IpAddress() pulumi.StringOutput {
	return o.ApplyT(func(v NetworkEndpointResponse) string { return v.IpAddress }).(pulumi.StringOutput)
}

// The port of this network endpoint.
func (o NetworkEndpointResponseOutput) Port() pulumi.IntOutput {
	return o.ApplyT(func(v NetworkEndpointResponse) int { return v.Port }).(pulumi.IntOutput)
}

type NetworkEndpointResponseArrayOutput struct{ *pulumi.OutputState }

func (NetworkEndpointResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NetworkEndpointResponse)(nil)).Elem()
}

func (o NetworkEndpointResponseArrayOutput) ToNetworkEndpointResponseArrayOutput() NetworkEndpointResponseArrayOutput {
	return o
}

func (o NetworkEndpointResponseArrayOutput) ToNetworkEndpointResponseArrayOutputWithContext(ctx context.Context) NetworkEndpointResponseArrayOutput {
	return o
}

func (o NetworkEndpointResponseArrayOutput) Index(i pulumi.IntInput) NetworkEndpointResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NetworkEndpointResponse {
		return vs[0].([]NetworkEndpointResponse)[vs[1].(int)]
	}).(NetworkEndpointResponseOutput)
}

// Sets the scheduling options for this node.
type SchedulingConfig struct {
	// Defines whether the node is preemptible.
	Preemptible *bool `pulumi:"preemptible"`
	// Whether the node is created under a reservation.
	Reserved *bool `pulumi:"reserved"`
}

// SchedulingConfigInput is an input type that accepts SchedulingConfigArgs and SchedulingConfigOutput values.
// You can construct a concrete instance of `SchedulingConfigInput` via:
//
//          SchedulingConfigArgs{...}
type SchedulingConfigInput interface {
	pulumi.Input

	ToSchedulingConfigOutput() SchedulingConfigOutput
	ToSchedulingConfigOutputWithContext(context.Context) SchedulingConfigOutput
}

// Sets the scheduling options for this node.
type SchedulingConfigArgs struct {
	// Defines whether the node is preemptible.
	Preemptible pulumi.BoolPtrInput `pulumi:"preemptible"`
	// Whether the node is created under a reservation.
	Reserved pulumi.BoolPtrInput `pulumi:"reserved"`
}

func (SchedulingConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SchedulingConfig)(nil)).Elem()
}

func (i SchedulingConfigArgs) ToSchedulingConfigOutput() SchedulingConfigOutput {
	return i.ToSchedulingConfigOutputWithContext(context.Background())
}

func (i SchedulingConfigArgs) ToSchedulingConfigOutputWithContext(ctx context.Context) SchedulingConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SchedulingConfigOutput)
}

func (i SchedulingConfigArgs) ToSchedulingConfigPtrOutput() SchedulingConfigPtrOutput {
	return i.ToSchedulingConfigPtrOutputWithContext(context.Background())
}

func (i SchedulingConfigArgs) ToSchedulingConfigPtrOutputWithContext(ctx context.Context) SchedulingConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SchedulingConfigOutput).ToSchedulingConfigPtrOutputWithContext(ctx)
}

// SchedulingConfigPtrInput is an input type that accepts SchedulingConfigArgs, SchedulingConfigPtr and SchedulingConfigPtrOutput values.
// You can construct a concrete instance of `SchedulingConfigPtrInput` via:
//
//          SchedulingConfigArgs{...}
//
//  or:
//
//          nil
type SchedulingConfigPtrInput interface {
	pulumi.Input

	ToSchedulingConfigPtrOutput() SchedulingConfigPtrOutput
	ToSchedulingConfigPtrOutputWithContext(context.Context) SchedulingConfigPtrOutput
}

type schedulingConfigPtrType SchedulingConfigArgs

func SchedulingConfigPtr(v *SchedulingConfigArgs) SchedulingConfigPtrInput {
	return (*schedulingConfigPtrType)(v)
}

func (*schedulingConfigPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SchedulingConfig)(nil)).Elem()
}

func (i *schedulingConfigPtrType) ToSchedulingConfigPtrOutput() SchedulingConfigPtrOutput {
	return i.ToSchedulingConfigPtrOutputWithContext(context.Background())
}

func (i *schedulingConfigPtrType) ToSchedulingConfigPtrOutputWithContext(ctx context.Context) SchedulingConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SchedulingConfigPtrOutput)
}

// Sets the scheduling options for this node.
type SchedulingConfigOutput struct{ *pulumi.OutputState }

func (SchedulingConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SchedulingConfig)(nil)).Elem()
}

func (o SchedulingConfigOutput) ToSchedulingConfigOutput() SchedulingConfigOutput {
	return o
}

func (o SchedulingConfigOutput) ToSchedulingConfigOutputWithContext(ctx context.Context) SchedulingConfigOutput {
	return o
}

func (o SchedulingConfigOutput) ToSchedulingConfigPtrOutput() SchedulingConfigPtrOutput {
	return o.ToSchedulingConfigPtrOutputWithContext(context.Background())
}

func (o SchedulingConfigOutput) ToSchedulingConfigPtrOutputWithContext(ctx context.Context) SchedulingConfigPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SchedulingConfig) *SchedulingConfig {
		return &v
	}).(SchedulingConfigPtrOutput)
}

// Defines whether the node is preemptible.
func (o SchedulingConfigOutput) Preemptible() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SchedulingConfig) *bool { return v.Preemptible }).(pulumi.BoolPtrOutput)
}

// Whether the node is created under a reservation.
func (o SchedulingConfigOutput) Reserved() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SchedulingConfig) *bool { return v.Reserved }).(pulumi.BoolPtrOutput)
}

type SchedulingConfigPtrOutput struct{ *pulumi.OutputState }

func (SchedulingConfigPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SchedulingConfig)(nil)).Elem()
}

func (o SchedulingConfigPtrOutput) ToSchedulingConfigPtrOutput() SchedulingConfigPtrOutput {
	return o
}

func (o SchedulingConfigPtrOutput) ToSchedulingConfigPtrOutputWithContext(ctx context.Context) SchedulingConfigPtrOutput {
	return o
}

func (o SchedulingConfigPtrOutput) Elem() SchedulingConfigOutput {
	return o.ApplyT(func(v *SchedulingConfig) SchedulingConfig {
		if v != nil {
			return *v
		}
		var ret SchedulingConfig
		return ret
	}).(SchedulingConfigOutput)
}

// Defines whether the node is preemptible.
func (o SchedulingConfigPtrOutput) Preemptible() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SchedulingConfig) *bool {
		if v == nil {
			return nil
		}
		return v.Preemptible
	}).(pulumi.BoolPtrOutput)
}

// Whether the node is created under a reservation.
func (o SchedulingConfigPtrOutput) Reserved() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SchedulingConfig) *bool {
		if v == nil {
			return nil
		}
		return v.Reserved
	}).(pulumi.BoolPtrOutput)
}

// Sets the scheduling options for this node.
type SchedulingConfigResponse struct {
	// Defines whether the node is preemptible.
	Preemptible bool `pulumi:"preemptible"`
	// Whether the node is created under a reservation.
	Reserved bool `pulumi:"reserved"`
}

// Sets the scheduling options for this node.
type SchedulingConfigResponseOutput struct{ *pulumi.OutputState }

func (SchedulingConfigResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SchedulingConfigResponse)(nil)).Elem()
}

func (o SchedulingConfigResponseOutput) ToSchedulingConfigResponseOutput() SchedulingConfigResponseOutput {
	return o
}

func (o SchedulingConfigResponseOutput) ToSchedulingConfigResponseOutputWithContext(ctx context.Context) SchedulingConfigResponseOutput {
	return o
}

// Defines whether the node is preemptible.
func (o SchedulingConfigResponseOutput) Preemptible() pulumi.BoolOutput {
	return o.ApplyT(func(v SchedulingConfigResponse) bool { return v.Preemptible }).(pulumi.BoolOutput)
}

// Whether the node is created under a reservation.
func (o SchedulingConfigResponseOutput) Reserved() pulumi.BoolOutput {
	return o.ApplyT(func(v SchedulingConfigResponse) bool { return v.Reserved }).(pulumi.BoolOutput)
}

// A Symptom instance.
type SymptomResponse struct {
	// Timestamp when the Symptom is created.
	CreateTime string `pulumi:"createTime"`
	// Detailed information of the current Symptom.
	Details string `pulumi:"details"`
	// Type of the Symptom.
	SymptomType string `pulumi:"symptomType"`
	// A string used to uniquely distinguish a worker within a TPU node.
	WorkerId string `pulumi:"workerId"`
}

// A Symptom instance.
type SymptomResponseOutput struct{ *pulumi.OutputState }

func (SymptomResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SymptomResponse)(nil)).Elem()
}

func (o SymptomResponseOutput) ToSymptomResponseOutput() SymptomResponseOutput {
	return o
}

func (o SymptomResponseOutput) ToSymptomResponseOutputWithContext(ctx context.Context) SymptomResponseOutput {
	return o
}

// Timestamp when the Symptom is created.
func (o SymptomResponseOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v SymptomResponse) string { return v.CreateTime }).(pulumi.StringOutput)
}

// Detailed information of the current Symptom.
func (o SymptomResponseOutput) Details() pulumi.StringOutput {
	return o.ApplyT(func(v SymptomResponse) string { return v.Details }).(pulumi.StringOutput)
}

// Type of the Symptom.
func (o SymptomResponseOutput) SymptomType() pulumi.StringOutput {
	return o.ApplyT(func(v SymptomResponse) string { return v.SymptomType }).(pulumi.StringOutput)
}

// A string used to uniquely distinguish a worker within a TPU node.
func (o SymptomResponseOutput) WorkerId() pulumi.StringOutput {
	return o.ApplyT(func(v SymptomResponse) string { return v.WorkerId }).(pulumi.StringOutput)
}

type SymptomResponseArrayOutput struct{ *pulumi.OutputState }

func (SymptomResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SymptomResponse)(nil)).Elem()
}

func (o SymptomResponseArrayOutput) ToSymptomResponseArrayOutput() SymptomResponseArrayOutput {
	return o
}

func (o SymptomResponseArrayOutput) ToSymptomResponseArrayOutputWithContext(ctx context.Context) SymptomResponseArrayOutput {
	return o
}

func (o SymptomResponseArrayOutput) Index(i pulumi.IntInput) SymptomResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SymptomResponse {
		return vs[0].([]SymptomResponse)[vs[1].(int)]
	}).(SymptomResponseOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SchedulingConfigInput)(nil)).Elem(), SchedulingConfigArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SchedulingConfigPtrInput)(nil)).Elem(), SchedulingConfigArgs{})
	pulumi.RegisterOutputType(NetworkEndpointResponseOutput{})
	pulumi.RegisterOutputType(NetworkEndpointResponseArrayOutput{})
	pulumi.RegisterOutputType(SchedulingConfigOutput{})
	pulumi.RegisterOutputType(SchedulingConfigPtrOutput{})
	pulumi.RegisterOutputType(SchedulingConfigResponseOutput{})
	pulumi.RegisterOutputType(SymptomResponseOutput{})
	pulumi.RegisterOutputType(SymptomResponseArrayOutput{})
}
