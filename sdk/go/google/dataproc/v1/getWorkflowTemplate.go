// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieves the latest workflow template.Can retrieve previously instantiated template by specifying optional version parameter.
func LookupWorkflowTemplate(ctx *pulumi.Context, args *LookupWorkflowTemplateArgs, opts ...pulumi.InvokeOption) (*LookupWorkflowTemplateResult, error) {
	var rv LookupWorkflowTemplateResult
	err := ctx.Invoke("google-native:dataproc/v1:getWorkflowTemplate", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWorkflowTemplateArgs struct {
	Location           string  `pulumi:"location"`
	Project            *string `pulumi:"project"`
	Version            *string `pulumi:"version"`
	WorkflowTemplateId string  `pulumi:"workflowTemplateId"`
}

type LookupWorkflowTemplateResult struct {
	// The time template was created.
	CreateTime string `pulumi:"createTime"`
	// Optional. Timeout duration for the DAG of jobs, expressed in seconds (see JSON representation of duration (https://developers.google.com/protocol-buffers/docs/proto3#json)). The timeout duration must be from 10 minutes ("600s") to 24 hours ("86400s"). The timer begins when the first job is submitted. If the workflow is running at the end of the timeout period, any remaining jobs are cancelled, the workflow is ended, and if the workflow was running on a managed cluster, the cluster is deleted.
	DagTimeout string `pulumi:"dagTimeout"`
	// The Directed Acyclic Graph of Jobs to submit.
	Jobs []OrderedJobResponse `pulumi:"jobs"`
	// Optional. The labels to associate with this template. These labels will be propagated to all jobs and clusters created by the workflow instance.Label keys must contain 1 to 63 characters, and must conform to RFC 1035 (https://www.ietf.org/rfc/rfc1035.txt).Label values may be empty, but, if present, must contain 1 to 63 characters, and must conform to RFC 1035 (https://www.ietf.org/rfc/rfc1035.txt).No more than 32 labels can be associated with a template.
	Labels map[string]string `pulumi:"labels"`
	// The resource name of the workflow template, as described in https://cloud.google.com/apis/design/resource_names. For projects.regions.workflowTemplates, the resource name of the template has the following format: projects/{project_id}/regions/{region}/workflowTemplates/{template_id} For projects.locations.workflowTemplates, the resource name of the template has the following format: projects/{project_id}/locations/{location}/workflowTemplates/{template_id}
	Name string `pulumi:"name"`
	// Optional. Template parameters whose values are substituted into the template. Values for parameters must be provided when the template is instantiated.
	Parameters []TemplateParameterResponse `pulumi:"parameters"`
	// WorkflowTemplate scheduling information.
	Placement WorkflowTemplatePlacementResponse `pulumi:"placement"`
	// The time template was last updated.
	UpdateTime string `pulumi:"updateTime"`
	// Optional. Used to perform a consistent read-modify-write.This field should be left blank for a CreateWorkflowTemplate request. It is required for an UpdateWorkflowTemplate request, and must match the current server version. A typical update template flow would fetch the current template with a GetWorkflowTemplate request, which will return the current template with the version field filled in with the current server version. The user updates other fields in the template, then returns it as part of the UpdateWorkflowTemplate request.
	Version int `pulumi:"version"`
}

func LookupWorkflowTemplateOutput(ctx *pulumi.Context, args LookupWorkflowTemplateOutputArgs, opts ...pulumi.InvokeOption) LookupWorkflowTemplateResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupWorkflowTemplateResult, error) {
			args := v.(LookupWorkflowTemplateArgs)
			r, err := LookupWorkflowTemplate(ctx, &args, opts...)
			return *r, err
		}).(LookupWorkflowTemplateResultOutput)
}

type LookupWorkflowTemplateOutputArgs struct {
	Location           pulumi.StringInput    `pulumi:"location"`
	Project            pulumi.StringPtrInput `pulumi:"project"`
	Version            pulumi.StringPtrInput `pulumi:"version"`
	WorkflowTemplateId pulumi.StringInput    `pulumi:"workflowTemplateId"`
}

func (LookupWorkflowTemplateOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWorkflowTemplateArgs)(nil)).Elem()
}

type LookupWorkflowTemplateResultOutput struct{ *pulumi.OutputState }

func (LookupWorkflowTemplateResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWorkflowTemplateResult)(nil)).Elem()
}

func (o LookupWorkflowTemplateResultOutput) ToLookupWorkflowTemplateResultOutput() LookupWorkflowTemplateResultOutput {
	return o
}

func (o LookupWorkflowTemplateResultOutput) ToLookupWorkflowTemplateResultOutputWithContext(ctx context.Context) LookupWorkflowTemplateResultOutput {
	return o
}

// The time template was created.
func (o LookupWorkflowTemplateResultOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkflowTemplateResult) string { return v.CreateTime }).(pulumi.StringOutput)
}

// Optional. Timeout duration for the DAG of jobs, expressed in seconds (see JSON representation of duration (https://developers.google.com/protocol-buffers/docs/proto3#json)). The timeout duration must be from 10 minutes ("600s") to 24 hours ("86400s"). The timer begins when the first job is submitted. If the workflow is running at the end of the timeout period, any remaining jobs are cancelled, the workflow is ended, and if the workflow was running on a managed cluster, the cluster is deleted.
func (o LookupWorkflowTemplateResultOutput) DagTimeout() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkflowTemplateResult) string { return v.DagTimeout }).(pulumi.StringOutput)
}

// The Directed Acyclic Graph of Jobs to submit.
func (o LookupWorkflowTemplateResultOutput) Jobs() OrderedJobResponseArrayOutput {
	return o.ApplyT(func(v LookupWorkflowTemplateResult) []OrderedJobResponse { return v.Jobs }).(OrderedJobResponseArrayOutput)
}

// Optional. The labels to associate with this template. These labels will be propagated to all jobs and clusters created by the workflow instance.Label keys must contain 1 to 63 characters, and must conform to RFC 1035 (https://www.ietf.org/rfc/rfc1035.txt).Label values may be empty, but, if present, must contain 1 to 63 characters, and must conform to RFC 1035 (https://www.ietf.org/rfc/rfc1035.txt).No more than 32 labels can be associated with a template.
func (o LookupWorkflowTemplateResultOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupWorkflowTemplateResult) map[string]string { return v.Labels }).(pulumi.StringMapOutput)
}

// The resource name of the workflow template, as described in https://cloud.google.com/apis/design/resource_names. For projects.regions.workflowTemplates, the resource name of the template has the following format: projects/{project_id}/regions/{region}/workflowTemplates/{template_id} For projects.locations.workflowTemplates, the resource name of the template has the following format: projects/{project_id}/locations/{location}/workflowTemplates/{template_id}
func (o LookupWorkflowTemplateResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkflowTemplateResult) string { return v.Name }).(pulumi.StringOutput)
}

// Optional. Template parameters whose values are substituted into the template. Values for parameters must be provided when the template is instantiated.
func (o LookupWorkflowTemplateResultOutput) Parameters() TemplateParameterResponseArrayOutput {
	return o.ApplyT(func(v LookupWorkflowTemplateResult) []TemplateParameterResponse { return v.Parameters }).(TemplateParameterResponseArrayOutput)
}

// WorkflowTemplate scheduling information.
func (o LookupWorkflowTemplateResultOutput) Placement() WorkflowTemplatePlacementResponseOutput {
	return o.ApplyT(func(v LookupWorkflowTemplateResult) WorkflowTemplatePlacementResponse { return v.Placement }).(WorkflowTemplatePlacementResponseOutput)
}

// The time template was last updated.
func (o LookupWorkflowTemplateResultOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkflowTemplateResult) string { return v.UpdateTime }).(pulumi.StringOutput)
}

// Optional. Used to perform a consistent read-modify-write.This field should be left blank for a CreateWorkflowTemplate request. It is required for an UpdateWorkflowTemplate request, and must match the current server version. A typical update template flow would fetch the current template with a GetWorkflowTemplate request, which will return the current template with the version field filled in with the current server version. The user updates other fields in the template, then returns it as part of the UpdateWorkflowTemplate request.
func (o LookupWorkflowTemplateResultOutput) Version() pulumi.IntOutput {
	return o.ApplyT(func(v LookupWorkflowTemplateResult) int { return v.Version }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWorkflowTemplateResultOutput{})
}
