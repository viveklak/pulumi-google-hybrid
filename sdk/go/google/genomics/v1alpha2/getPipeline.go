// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1alpha2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieves a pipeline based on ID. Caller must have READ permission to the project.
func LookupPipeline(ctx *pulumi.Context, args *LookupPipelineArgs, opts ...pulumi.InvokeOption) (*LookupPipelineResult, error) {
	var rv LookupPipelineResult
	err := ctx.Invoke("google-native:genomics/v1alpha2:getPipeline", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPipelineArgs struct {
	PipelineId string `pulumi:"pipelineId"`
}

type LookupPipelineResult struct {
	// User-specified description.
	Description string `pulumi:"description"`
	// Specifies the docker run information.
	Docker DockerExecutorResponse `pulumi:"docker"`
	// Input parameters of the pipeline.
	InputParameters []PipelineParameterResponse `pulumi:"inputParameters"`
	// A user specified pipeline name that does not have to be unique. This name can be used for filtering Pipelines in ListPipelines.
	Name string `pulumi:"name"`
	// Output parameters of the pipeline.
	OutputParameters []PipelineParameterResponse `pulumi:"outputParameters"`
	// Unique pipeline id that is generated by the service when CreatePipeline is called. Cannot be specified in the Pipeline used in the CreatePipelineRequest, and will be populated in the response to CreatePipeline and all subsequent Get and List calls. Indicates that the service has registered this pipeline.
	PipelineId string `pulumi:"pipelineId"`
	// The project in which to create the pipeline. The caller must have WRITE access.
	Project string `pulumi:"project"`
	// Specifies resource requirements for the pipeline run. Required fields: * minimumCpuCores * minimumRamGb
	Resources PipelineResourcesResponse `pulumi:"resources"`
}

func LookupPipelineOutput(ctx *pulumi.Context, args LookupPipelineOutputArgs, opts ...pulumi.InvokeOption) LookupPipelineResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPipelineResult, error) {
			args := v.(LookupPipelineArgs)
			r, err := LookupPipeline(ctx, &args, opts...)
			return *r, err
		}).(LookupPipelineResultOutput)
}

type LookupPipelineOutputArgs struct {
	PipelineId pulumi.StringInput `pulumi:"pipelineId"`
}

func (LookupPipelineOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPipelineArgs)(nil)).Elem()
}

type LookupPipelineResultOutput struct{ *pulumi.OutputState }

func (LookupPipelineResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPipelineResult)(nil)).Elem()
}

func (o LookupPipelineResultOutput) ToLookupPipelineResultOutput() LookupPipelineResultOutput {
	return o
}

func (o LookupPipelineResultOutput) ToLookupPipelineResultOutputWithContext(ctx context.Context) LookupPipelineResultOutput {
	return o
}

// User-specified description.
func (o LookupPipelineResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPipelineResult) string { return v.Description }).(pulumi.StringOutput)
}

// Specifies the docker run information.
func (o LookupPipelineResultOutput) Docker() DockerExecutorResponseOutput {
	return o.ApplyT(func(v LookupPipelineResult) DockerExecutorResponse { return v.Docker }).(DockerExecutorResponseOutput)
}

// Input parameters of the pipeline.
func (o LookupPipelineResultOutput) InputParameters() PipelineParameterResponseArrayOutput {
	return o.ApplyT(func(v LookupPipelineResult) []PipelineParameterResponse { return v.InputParameters }).(PipelineParameterResponseArrayOutput)
}

// A user specified pipeline name that does not have to be unique. This name can be used for filtering Pipelines in ListPipelines.
func (o LookupPipelineResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPipelineResult) string { return v.Name }).(pulumi.StringOutput)
}

// Output parameters of the pipeline.
func (o LookupPipelineResultOutput) OutputParameters() PipelineParameterResponseArrayOutput {
	return o.ApplyT(func(v LookupPipelineResult) []PipelineParameterResponse { return v.OutputParameters }).(PipelineParameterResponseArrayOutput)
}

// Unique pipeline id that is generated by the service when CreatePipeline is called. Cannot be specified in the Pipeline used in the CreatePipelineRequest, and will be populated in the response to CreatePipeline and all subsequent Get and List calls. Indicates that the service has registered this pipeline.
func (o LookupPipelineResultOutput) PipelineId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPipelineResult) string { return v.PipelineId }).(pulumi.StringOutput)
}

// The project in which to create the pipeline. The caller must have WRITE access.
func (o LookupPipelineResultOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPipelineResult) string { return v.Project }).(pulumi.StringOutput)
}

// Specifies resource requirements for the pipeline run. Required fields: * minimumCpuCores * minimumRamGb
func (o LookupPipelineResultOutput) Resources() PipelineResourcesResponseOutput {
	return o.ApplyT(func(v LookupPipelineResult) PipelineResourcesResponse { return v.Resources }).(PipelineResourcesResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPipelineResultOutput{})
}
