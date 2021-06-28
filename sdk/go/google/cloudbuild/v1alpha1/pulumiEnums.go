// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1alpha1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type WorkerPoolRegionsItem pulumi.String

const (
	// no region
	WorkerPoolRegionsItemRegionUnspecified = WorkerPoolRegionsItem("REGION_UNSPECIFIED")
	// us-central1 region
	WorkerPoolRegionsItemUsCentral1 = WorkerPoolRegionsItem("us-central1")
	// us-west1 region
	WorkerPoolRegionsItemUsWest1 = WorkerPoolRegionsItem("us-west1")
	// us-east1 region
	WorkerPoolRegionsItemUsEast1 = WorkerPoolRegionsItem("us-east1")
	// us-east4 region
	WorkerPoolRegionsItemUsEast4 = WorkerPoolRegionsItem("us-east4")
)

func (WorkerPoolRegionsItem) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e WorkerPoolRegionsItem) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e WorkerPoolRegionsItem) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e WorkerPoolRegionsItem) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e WorkerPoolRegionsItem) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

// WorkerPoolRegionsItemArrayInput is an input type that accepts WorkerPoolRegionsItemArray and WorkerPoolRegionsItemArrayOutput values.
// You can construct a concrete instance of `WorkerPoolRegionsItemArrayInput` via:
//
//          WorkerPoolRegionsItemArray{ WorkerPoolRegionsItemArgs{...} }
type WorkerPoolRegionsItemArrayInput interface {
	pulumi.Input

	ToWorkerPoolRegionsItemArrayOutput() WorkerPoolRegionsItemArrayOutput
	ToWorkerPoolRegionsItemArrayOutputWithContext(context.Context) WorkerPoolRegionsItemArrayOutput
}

type WorkerPoolRegionsItemArray []WorkerPoolRegionsItem

func (WorkerPoolRegionsItemArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]WorkerPoolRegionsItem)(nil)).Elem()
}

func (i WorkerPoolRegionsItemArray) ToWorkerPoolRegionsItemArrayOutput() WorkerPoolRegionsItemArrayOutput {
	return i.ToWorkerPoolRegionsItemArrayOutputWithContext(context.Background())
}

func (i WorkerPoolRegionsItemArray) ToWorkerPoolRegionsItemArrayOutputWithContext(ctx context.Context) WorkerPoolRegionsItemArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkerPoolRegionsItemArrayOutput)
}

type WorkerPoolRegionsItemArrayOutput struct{ *pulumi.OutputState }

func (WorkerPoolRegionsItemArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]WorkerPoolRegionsItem)(nil)).Elem()
}

func (o WorkerPoolRegionsItemArrayOutput) ToWorkerPoolRegionsItemArrayOutput() WorkerPoolRegionsItemArrayOutput {
	return o
}

func (o WorkerPoolRegionsItemArrayOutput) ToWorkerPoolRegionsItemArrayOutputWithContext(ctx context.Context) WorkerPoolRegionsItemArrayOutput {
	return o
}

func (o WorkerPoolRegionsItemArrayOutput) Index(i pulumi.IntInput) pulumi.StringOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) pulumi.StringOutput {
		return vs[0].([]WorkerPoolRegionsItem)[vs[1].(int)].ToStringOutput()
	}).(pulumi.StringOutput)
}
