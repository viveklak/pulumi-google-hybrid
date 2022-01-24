// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi.Utilities;

namespace Pulumi.GoogleNative.VMMigration.V1Alpha1
{
    public static class GetTargetProject
    {
        /// <summary>
        /// Gets details of a single TargetProject. NOTE: TargetProject is a global resource; hence the only supported value for location is `global`.
        /// </summary>
        public static Task<GetTargetProjectResult> InvokeAsync(GetTargetProjectArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetTargetProjectResult>("google-native:vmmigration/v1alpha1:getTargetProject", args ?? new GetTargetProjectArgs(), options.WithDefaults());

        /// <summary>
        /// Gets details of a single TargetProject. NOTE: TargetProject is a global resource; hence the only supported value for location is `global`.
        /// </summary>
        public static Output<GetTargetProjectResult> Invoke(GetTargetProjectInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetTargetProjectResult>("google-native:vmmigration/v1alpha1:getTargetProject", args ?? new GetTargetProjectInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetTargetProjectArgs : Pulumi.InvokeArgs
    {
        [Input("location", required: true)]
        public string Location { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        [Input("targetProjectId", required: true)]
        public string TargetProjectId { get; set; } = null!;

        public GetTargetProjectArgs()
        {
        }
    }

    public sealed class GetTargetProjectInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("targetProjectId", required: true)]
        public Input<string> TargetProjectId { get; set; } = null!;

        public GetTargetProjectInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetTargetProjectResult
    {
        /// <summary>
        /// The time this target project resource was created (not related to when the Compute Engine project it points to was created).
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// The target project's description.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The name of the target project.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The target project ID (number) or project name.
        /// </summary>
        public readonly string Project;
        /// <summary>
        /// The last time the target project resource was updated.
        /// </summary>
        public readonly string UpdateTime;

        [OutputConstructor]
        private GetTargetProjectResult(
            string createTime,

            string description,

            string name,

            string project,

            string updateTime)
        {
            CreateTime = createTime;
            Description = description;
            Name = name;
            Project = project;
            UpdateTime = updateTime;
        }
    }
}
