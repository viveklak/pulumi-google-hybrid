// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.CloudResourceManager.V1
{
    public static class GetProject
    {
        /// <summary>
        /// Retrieves the Project identified by the specified `project_id` (for example, `my-project-123`). The caller must have read permissions for this Project.
        /// </summary>
        public static Task<GetProjectResult> InvokeAsync(GetProjectArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetProjectResult>("google-native:cloudresourcemanager/v1:getProject", args ?? new GetProjectArgs(), options.WithDefaults());

        /// <summary>
        /// Retrieves the Project identified by the specified `project_id` (for example, `my-project-123`). The caller must have read permissions for this Project.
        /// </summary>
        public static Output<GetProjectResult> Invoke(GetProjectInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetProjectResult>("google-native:cloudresourcemanager/v1:getProject", args ?? new GetProjectInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetProjectArgs : Pulumi.InvokeArgs
    {
        [Input("project")]
        public string? Project { get; set; }

        public GetProjectArgs()
        {
        }
    }

    public sealed class GetProjectInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("project")]
        public Input<string>? Project { get; set; }

        public GetProjectInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetProjectResult
    {
        /// <summary>
        /// Creation time. Read-only.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// The labels associated with this Project. Label keys must be between 1 and 63 characters long and must conform to the following regular expression: a-z{0,62}. Label values must be between 0 and 63 characters long and must conform to the regular expression [a-z0-9_-]{0,63}. A label value can be empty. No more than 256 labels can be associated with a given resource. Clients should store labels in a representation such as JSON that does not depend on specific characters being disallowed. Example: "environment" : "dev" Read-write.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Labels;
        /// <summary>
        /// The Project lifecycle state. Read-only.
        /// </summary>
        public readonly string LifecycleState;
        /// <summary>
        /// The optional user-assigned display name of the Project. When present it must be between 4 to 30 characters. Allowed characters are: lowercase and uppercase letters, numbers, hyphen, single-quote, double-quote, space, and exclamation point. Example: `My Project` Read-write.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// An optional reference to a parent Resource. Supported parent types include "organization" and "folder". Once set, the parent cannot be cleared. The `parent` can be set on creation or using the `UpdateProject` method; the end user must have the `resourcemanager.projects.create` permission on the parent.
        /// </summary>
        public readonly Outputs.ResourceIdResponse Parent;
        /// <summary>
        /// The unique, user-assigned ID of the Project. It must be 6 to 30 lowercase letters, digits, or hyphens. It must start with a letter. Trailing hyphens are prohibited. Example: `tokyo-rain-123` Read-only after creation.
        /// </summary>
        public readonly string ProjectId;
        /// <summary>
        /// The number uniquely identifying the project. Example: `415104041262` Read-only.
        /// </summary>
        public readonly string ProjectNumber;

        [OutputConstructor]
        private GetProjectResult(
            string createTime,

            ImmutableDictionary<string, string> labels,

            string lifecycleState,

            string name,

            Outputs.ResourceIdResponse parent,

            string projectId,

            string projectNumber)
        {
            CreateTime = createTime;
            Labels = labels;
            LifecycleState = lifecycleState;
            Name = name;
            Parent = parent;
            ProjectId = projectId;
            ProjectNumber = projectNumber;
        }
    }
}
