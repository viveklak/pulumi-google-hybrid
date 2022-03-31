// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dataproc.V1
{
    public static class GetWorkflowTemplate
    {
        /// <summary>
        /// Retrieves the latest workflow template.Can retrieve previously instantiated template by specifying optional version parameter.
        /// </summary>
        public static Task<GetWorkflowTemplateResult> InvokeAsync(GetWorkflowTemplateArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetWorkflowTemplateResult>("google-native:dataproc/v1:getWorkflowTemplate", args ?? new GetWorkflowTemplateArgs(), options.WithDefaults());

        /// <summary>
        /// Retrieves the latest workflow template.Can retrieve previously instantiated template by specifying optional version parameter.
        /// </summary>
        public static Output<GetWorkflowTemplateResult> Invoke(GetWorkflowTemplateInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetWorkflowTemplateResult>("google-native:dataproc/v1:getWorkflowTemplate", args ?? new GetWorkflowTemplateInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetWorkflowTemplateArgs : Pulumi.InvokeArgs
    {
        [Input("location", required: true)]
        public string Location { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        [Input("version")]
        public string? Version { get; set; }

        [Input("workflowTemplateId", required: true)]
        public string WorkflowTemplateId { get; set; } = null!;

        public GetWorkflowTemplateArgs()
        {
        }
    }

    public sealed class GetWorkflowTemplateInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("version")]
        public Input<string>? Version { get; set; }

        [Input("workflowTemplateId", required: true)]
        public Input<string> WorkflowTemplateId { get; set; } = null!;

        public GetWorkflowTemplateInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetWorkflowTemplateResult
    {
        /// <summary>
        /// The time template was created.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// Optional. Timeout duration for the DAG of jobs, expressed in seconds (see JSON representation of duration (https://developers.google.com/protocol-buffers/docs/proto3#json)). The timeout duration must be from 10 minutes ("600s") to 24 hours ("86400s"). The timer begins when the first job is submitted. If the workflow is running at the end of the timeout period, any remaining jobs are cancelled, the workflow is ended, and if the workflow was running on a managed cluster, the cluster is deleted.
        /// </summary>
        public readonly string DagTimeout;
        /// <summary>
        /// The Directed Acyclic Graph of Jobs to submit.
        /// </summary>
        public readonly ImmutableArray<Outputs.OrderedJobResponse> Jobs;
        /// <summary>
        /// Optional. The labels to associate with this template. These labels will be propagated to all jobs and clusters created by the workflow instance.Label keys must contain 1 to 63 characters, and must conform to RFC 1035 (https://www.ietf.org/rfc/rfc1035.txt).Label values may be empty, but, if present, must contain 1 to 63 characters, and must conform to RFC 1035 (https://www.ietf.org/rfc/rfc1035.txt).No more than 32 labels can be associated with a template.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Labels;
        /// <summary>
        /// The resource name of the workflow template, as described in https://cloud.google.com/apis/design/resource_names. For projects.regions.workflowTemplates, the resource name of the template has the following format: projects/{project_id}/regions/{region}/workflowTemplates/{template_id} For projects.locations.workflowTemplates, the resource name of the template has the following format: projects/{project_id}/locations/{location}/workflowTemplates/{template_id}
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Optional. Template parameters whose values are substituted into the template. Values for parameters must be provided when the template is instantiated.
        /// </summary>
        public readonly ImmutableArray<Outputs.TemplateParameterResponse> Parameters;
        /// <summary>
        /// WorkflowTemplate scheduling information.
        /// </summary>
        public readonly Outputs.WorkflowTemplatePlacementResponse Placement;
        /// <summary>
        /// The time template was last updated.
        /// </summary>
        public readonly string UpdateTime;
        /// <summary>
        /// Optional. Used to perform a consistent read-modify-write.This field should be left blank for a CreateWorkflowTemplate request. It is required for an UpdateWorkflowTemplate request, and must match the current server version. A typical update template flow would fetch the current template with a GetWorkflowTemplate request, which will return the current template with the version field filled in with the current server version. The user updates other fields in the template, then returns it as part of the UpdateWorkflowTemplate request.
        /// </summary>
        public readonly int Version;

        [OutputConstructor]
        private GetWorkflowTemplateResult(
            string createTime,

            string dagTimeout,

            ImmutableArray<Outputs.OrderedJobResponse> jobs,

            ImmutableDictionary<string, string> labels,

            string name,

            ImmutableArray<Outputs.TemplateParameterResponse> parameters,

            Outputs.WorkflowTemplatePlacementResponse placement,

            string updateTime,

            int version)
        {
            CreateTime = createTime;
            DagTimeout = dagTimeout;
            Jobs = jobs;
            Labels = labels;
            Name = name;
            Parameters = parameters;
            Placement = placement;
            UpdateTime = updateTime;
            Version = version;
        }
    }
}
