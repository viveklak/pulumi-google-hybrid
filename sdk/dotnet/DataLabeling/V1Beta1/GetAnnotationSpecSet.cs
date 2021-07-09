// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.DataLabeling.V1Beta1
{
    public static class GetAnnotationSpecSet
    {
        /// <summary>
        /// Gets an annotation spec set by resource name.
        /// </summary>
        public static Task<GetAnnotationSpecSetResult> InvokeAsync(GetAnnotationSpecSetArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetAnnotationSpecSetResult>("google-native:datalabeling/v1beta1:getAnnotationSpecSet", args ?? new GetAnnotationSpecSetArgs(), options.WithVersion());
    }


    public sealed class GetAnnotationSpecSetArgs : Pulumi.InvokeArgs
    {
        [Input("annotationSpecSetId", required: true)]
        public string AnnotationSpecSetId { get; set; } = null!;

        [Input("project", required: true)]
        public string Project { get; set; } = null!;

        public GetAnnotationSpecSetArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetAnnotationSpecSetResult
    {
        /// <summary>
        /// The array of AnnotationSpecs that you define when you create the AnnotationSpecSet. These are the possible labels for the labeling task.
        /// </summary>
        public readonly ImmutableArray<Outputs.GoogleCloudDatalabelingV1beta1AnnotationSpecResponse> AnnotationSpecs;
        /// <summary>
        /// The names of any related resources that are blocking changes to the annotation spec set.
        /// </summary>
        public readonly ImmutableArray<string> BlockingResources;
        /// <summary>
        /// Optional. User-provided description of the annotation specification set. The description can be up to 10,000 characters long.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The display name for AnnotationSpecSet that you define when you create it. Maximum of 64 characters.
        /// </summary>
        public readonly string DisplayName;
        /// <summary>
        /// The AnnotationSpecSet resource name in the following format: "projects/{project_id}/annotationSpecSets/{annotation_spec_set_id}"
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private GetAnnotationSpecSetResult(
            ImmutableArray<Outputs.GoogleCloudDatalabelingV1beta1AnnotationSpecResponse> annotationSpecs,

            ImmutableArray<string> blockingResources,

            string description,

            string displayName,

            string name)
        {
            AnnotationSpecs = annotationSpecs;
            BlockingResources = blockingResources;
            Description = description;
            DisplayName = displayName;
            Name = name;
        }
    }
}
