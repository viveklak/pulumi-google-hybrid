// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi.Utilities;

namespace Pulumi.GoogleNative.DataCatalog.V1
{
    public static class GetPolicyTag
    {
        /// <summary>
        /// Gets a policy tag.
        /// </summary>
        public static Task<GetPolicyTagResult> InvokeAsync(GetPolicyTagArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetPolicyTagResult>("google-native:datacatalog/v1:getPolicyTag", args ?? new GetPolicyTagArgs(), options.WithDefaults());

        /// <summary>
        /// Gets a policy tag.
        /// </summary>
        public static Output<GetPolicyTagResult> Invoke(GetPolicyTagInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetPolicyTagResult>("google-native:datacatalog/v1:getPolicyTag", args ?? new GetPolicyTagInvokeArgs(), options.WithDefaults());
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
        /// Description of this policy tag. If not set, defaults to empty. The description must contain only Unicode characters, tabs, newlines, carriage returns and page breaks, and be at most 2000 bytes long when encoded in UTF-8.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// User-defined name of this policy tag. The name can't start or end with spaces and must be unique within the parent taxonomy, contain only Unicode letters, numbers, underscores, dashes and spaces, and be at most 200 bytes long when encoded in UTF-8.
        /// </summary>
        public readonly string DisplayName;
        /// <summary>
        /// Resource name of this policy tag in the URL format. The policy tag manager generates unique taxonomy IDs and policy tag IDs.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Resource name of this policy tag's parent policy tag. If empty, this is a top level tag. If not set, defaults to an empty string. For example, for the "LatLong" policy tag in the example above, this field contains the resource name of the "Geolocation" policy tag, and, for "Geolocation", this field is empty.
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
