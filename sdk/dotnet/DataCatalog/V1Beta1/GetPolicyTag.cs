// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi.Utilities;

namespace Pulumi.GoogleNative.DataCatalog.V1Beta1
{
    public static class GetPolicyTag
    {
        /// <summary>
        /// Gets a policy tag.
        /// </summary>
        public static Task<GetPolicyTagResult> InvokeAsync(GetPolicyTagArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetPolicyTagResult>("google-native:datacatalog/v1beta1:getPolicyTag", args ?? new GetPolicyTagArgs(), options.WithDefaults());

        /// <summary>
        /// Gets a policy tag.
        /// </summary>
        public static Output<GetPolicyTagResult> Invoke(GetPolicyTagInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetPolicyTagResult>("google-native:datacatalog/v1beta1:getPolicyTag", args ?? new GetPolicyTagInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetPolicyTagArgs : Pulumi.InvokeArgs
    {
        [Input("location", required: true)]
        public string Location { get; set; } = null!;

        [Input("policyTagId", required: true)]
        public string PolicyTagId { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        [Input("taxonomyId", required: true)]
        public string TaxonomyId { get; set; } = null!;

        public GetPolicyTagArgs()
        {
        }
    }

    public sealed class GetPolicyTagInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("policyTagId", required: true)]
        public Input<string> PolicyTagId { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("taxonomyId", required: true)]
        public Input<string> TaxonomyId { get; set; } = null!;

        public GetPolicyTagInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetPolicyTagResult
    {
        /// <summary>
        /// Resource names of child policy tags of this policy tag.
        /// </summary>
        public readonly ImmutableArray<string> ChildPolicyTags;
        /// <summary>
        /// Description of this policy tag. It must: contain only unicode characters, tabs, newlines, carriage returns and page breaks; and be at most 2000 bytes long when encoded in UTF-8. If not set, defaults to an empty description. If not set, defaults to an empty description.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// User defined name of this policy tag. It must: be unique within the parent taxonomy; contain only unicode letters, numbers, underscores, dashes and spaces; not start or end with spaces; and be at most 200 bytes long when encoded in UTF-8.
        /// </summary>
        public readonly string DisplayName;
        /// <summary>
        /// Resource name of this policy tag, whose format is: "projects/{project_number}/locations/{location_id}/taxonomies/{taxonomy_id}/policyTags/{id}".
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Resource name of this policy tag's parent policy tag (e.g. for the "LatLong" policy tag in the example above, this field contains the resource name of the "Geolocation" policy tag). If empty, it means this policy tag is a top level policy tag (e.g. this field is empty for the "Geolocation" policy tag in the example above). If not set, defaults to an empty string.
        /// </summary>
        public readonly string ParentPolicyTag;

        [OutputConstructor]
        private GetPolicyTagResult(
            ImmutableArray<string> childPolicyTags,

            string description,

            string displayName,

            string name,

            string parentPolicyTag)
        {
            ChildPolicyTags = childPolicyTags;
            Description = description;
            DisplayName = displayName;
            Name = name;
            ParentPolicyTag = parentPolicyTag;
        }
    }
}
