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
    public static class GetTaxonomy
    {
        /// <summary>
        /// Gets a taxonomy.
        /// </summary>
        public static Task<GetTaxonomyResult> InvokeAsync(GetTaxonomyArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetTaxonomyResult>("google-native:datacatalog/v1:getTaxonomy", args ?? new GetTaxonomyArgs(), options.WithDefaults());

        /// <summary>
        /// Gets a taxonomy.
        /// </summary>
        public static Output<GetTaxonomyResult> Invoke(GetTaxonomyInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetTaxonomyResult>("google-native:datacatalog/v1:getTaxonomy", args ?? new GetTaxonomyInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetTaxonomyArgs : Pulumi.InvokeArgs
    {
        [Input("location", required: true)]
        public string Location { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        [Input("taxonomyId", required: true)]
        public string TaxonomyId { get; set; } = null!;

        public GetTaxonomyArgs()
        {
        }
    }

    public sealed class GetTaxonomyInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("taxonomyId", required: true)]
        public Input<string> TaxonomyId { get; set; } = null!;

        public GetTaxonomyInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetTaxonomyResult
    {
        /// <summary>
        /// Optional. A list of policy types that are activated for this taxonomy. If not set, defaults to an empty list.
        /// </summary>
        public readonly ImmutableArray<string> ActivatedPolicyTypes;
        /// <summary>
        /// Optional. Description of this taxonomy. If not set, defaults to empty. The description must contain only Unicode characters, tabs, newlines, carriage returns, and page breaks, and be at most 2000 bytes long when encoded in UTF-8.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// User-defined name of this taxonomy. The name can't start or end with spaces, must contain only Unicode letters, numbers, underscores, dashes, and spaces, and be at most 200 bytes long when encoded in UTF-8.
        /// </summary>
        public readonly string DisplayName;
        /// <summary>
        /// Resource name of this taxonomy in URL format. Note: Policy tag manager generates unique taxonomy IDs.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Number of policy tags in this taxonomy.
        /// </summary>
        public readonly int PolicyTagCount;
        /// <summary>
        /// Creation and modification timestamps of this taxonomy.
        /// </summary>
        public readonly Outputs.GoogleCloudDatacatalogV1SystemTimestampsResponse TaxonomyTimestamps;

        [OutputConstructor]
        private GetTaxonomyResult(
            ImmutableArray<string> activatedPolicyTypes,

            string description,

            string displayName,

            string name,

            int policyTagCount,

            Outputs.GoogleCloudDatacatalogV1SystemTimestampsResponse taxonomyTimestamps)
        {
            ActivatedPolicyTypes = activatedPolicyTypes;
            Description = description;
            DisplayName = displayName;
            Name = name;
            PolicyTagCount = policyTagCount;
            TaxonomyTimestamps = taxonomyTimestamps;
        }
    }
}
