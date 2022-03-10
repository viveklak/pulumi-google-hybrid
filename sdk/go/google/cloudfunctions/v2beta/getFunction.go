// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v2beta

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Returns a function with the given name from the requested project.
func LookupFunction(ctx *pulumi.Context, args *LookupFunctionArgs, opts ...pulumi.InvokeOption) (*LookupFunctionResult, error) {
	var rv LookupFunctionResult
	err := ctx.Invoke("google-native:cloudfunctions/v2beta:getFunction", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupFunctionArgs struct {
	FunctionId string  `pulumi:"functionId"`
	Location   string  `pulumi:"location"`
	Project    *string `pulumi:"project"`
}

type LookupFunctionResult struct {
	// Describes the Build step of the function that builds a container from the given source.
	BuildConfig BuildConfigResponse `pulumi:"buildConfig"`
	// User-provided description of a function.
	Description string `pulumi:"description"`
	// Describe whether the function is gen1 or gen2.
	Environment string `pulumi:"environment"`
	// An Eventarc trigger managed by Google Cloud Functions that fires events in response to a condition in another service.
	EventTrigger EventTriggerResponse `pulumi:"eventTrigger"`
	// Labels associated with this Cloud Function.
	Labels map[string]string `pulumi:"labels"`
	// A user-defined name of the function. Function names must be unique globally and match pattern `projects/*/locations/*/functions/*`
	Name string `pulumi:"name"`
	// Describes the Service being deployed. Currently deploys services to Cloud Run (fully managed).
	ServiceConfig ServiceConfigResponse `pulumi:"serviceConfig"`
	// State of the function.
	State string `pulumi:"state"`
	// State Messages for this Cloud Function.
	StateMessages []GoogleCloudFunctionsV2betaStateMessageResponse `pulumi:"stateMessages"`
	// The last update timestamp of a Cloud Function.
	UpdateTime string `pulumi:"updateTime"`
}

func LookupFunctionOutput(ctx *pulumi.Context, args LookupFunctionOutputArgs, opts ...pulumi.InvokeOption) LookupFunctionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupFunctionResult, error) {
			args := v.(LookupFunctionArgs)
			r, err := LookupFunction(ctx, &args, opts...)
			return *r, err
		}).(LookupFunctionResultOutput)
}

type LookupFunctionOutputArgs struct {
	FunctionId pulumi.StringInput    `pulumi:"functionId"`
	Location   pulumi.StringInput    `pulumi:"location"`
	Project    pulumi.StringPtrInput `pulumi:"project"`
}

func (LookupFunctionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFunctionArgs)(nil)).Elem()
}

type LookupFunctionResultOutput struct{ *pulumi.OutputState }

func (LookupFunctionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFunctionResult)(nil)).Elem()
}

func (o LookupFunctionResultOutput) ToLookupFunctionResultOutput() LookupFunctionResultOutput {
	return o
}

func (o LookupFunctionResultOutput) ToLookupFunctionResultOutputWithContext(ctx context.Context) LookupFunctionResultOutput {
	return o
}

// Describes the Build step of the function that builds a container from the given source.
func (o LookupFunctionResultOutput) BuildConfig() BuildConfigResponseOutput {
	return o.ApplyT(func(v LookupFunctionResult) BuildConfigResponse { return v.BuildConfig }).(BuildConfigResponseOutput)
}

// User-provided description of a function.
func (o LookupFunctionResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFunctionResult) string { return v.Description }).(pulumi.StringOutput)
}

// Describe whether the function is gen1 or gen2.
func (o LookupFunctionResultOutput) Environment() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFunctionResult) string { return v.Environment }).(pulumi.StringOutput)
}

// An Eventarc trigger managed by Google Cloud Functions that fires events in response to a condition in another service.
func (o LookupFunctionResultOutput) EventTrigger() EventTriggerResponseOutput {
	return o.ApplyT(func(v LookupFunctionResult) EventTriggerResponse { return v.EventTrigger }).(EventTriggerResponseOutput)
}

// Labels associated with this Cloud Function.
func (o LookupFunctionResultOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupFunctionResult) map[string]string { return v.Labels }).(pulumi.StringMapOutput)
}

// A user-defined name of the function. Function names must be unique globally and match pattern `projects/*/locations/*/functions/*`
func (o LookupFunctionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFunctionResult) string { return v.Name }).(pulumi.StringOutput)
}

// Describes the Service being deployed. Currently deploys services to Cloud Run (fully managed).
func (o LookupFunctionResultOutput) ServiceConfig() ServiceConfigResponseOutput {
	return o.ApplyT(func(v LookupFunctionResult) ServiceConfigResponse { return v.ServiceConfig }).(ServiceConfigResponseOutput)
}

// State of the function.
func (o LookupFunctionResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFunctionResult) string { return v.State }).(pulumi.StringOutput)
}

// State Messages for this Cloud Function.
func (o LookupFunctionResultOutput) StateMessages() GoogleCloudFunctionsV2betaStateMessageResponseArrayOutput {
	return o.ApplyT(func(v LookupFunctionResult) []GoogleCloudFunctionsV2betaStateMessageResponse { return v.StateMessages }).(GoogleCloudFunctionsV2betaStateMessageResponseArrayOutput)
}

// The last update timestamp of a Cloud Function.
func (o LookupFunctionResultOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFunctionResult) string { return v.UpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFunctionResultOutput{})
}