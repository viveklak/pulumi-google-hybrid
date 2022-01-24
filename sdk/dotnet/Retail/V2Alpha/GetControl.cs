// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi.Utilities;

namespace Pulumi.GoogleNative.Retail.V2Alpha
{
    public static class GetControl
    {
        /// <summary>
        /// Gets a Control.
        /// </summary>
        public static Task<GetControlResult> InvokeAsync(GetControlArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetControlResult>("google-native:retail/v2alpha:getControl", args ?? new GetControlArgs(), options.WithDefaults());

        /// <summary>
        /// Gets a Control.
        /// </summary>
        public static Output<GetControlResult> Invoke(GetControlInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetControlResult>("google-native:retail/v2alpha:getControl", args ?? new GetControlInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetControlArgs : Pulumi.InvokeArgs
    {
        [Input("catalogId", required: true)]
        public string CatalogId { get; set; } = null!;

        [Input("controlId", required: true)]
        public string ControlId { get; set; } = null!;

        [Input("location", required: true)]
        public string Location { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        public GetControlArgs()
        {
        }
    }

    public sealed class GetControlInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("catalogId", required: true)]
        public Input<string> CatalogId { get; set; } = null!;

        [Input("controlId", required: true)]
        public Input<string> ControlId { get; set; } = null!;

        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        public GetControlInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetControlResult
    {
        /// <summary>
        /// List of serving configuration ids that that are associated with this control. Note the association is managed via the ServingConfig, this is an output only denormalizeed view. Assumed to be in the same catalog.
        /// </summary>
        public readonly ImmutableArray<string> AssociatedServingConfigIds;
        /// <summary>
        /// The human readable control display name. Used in Retail UI. This field must be a UTF-8 encoded string with a length limit of 128 characters. Otherwise, an INVALID_ARGUMENT error is thrown.
        /// </summary>
        public readonly string DisplayName;
        /// <summary>
        /// A facet specification to perform faceted search.
        /// </summary>
        public readonly Outputs.GoogleCloudRetailV2alphaSearchRequestFacetSpecResponse FacetSpec;
        /// <summary>
        /// Immutable. Fully qualified name projects/*/locations/global/catalogs/*/controls/*
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// A rule control - a condition-action pair. Enacts a set action when the condition is triggered. For example: Boost "gShoe" when query full matches "Running Shoes".
        /// </summary>
        public readonly Outputs.GoogleCloudRetailV2alphaRuleResponse Rule;
        /// <summary>
        /// Immutable. The solution types that the serving config is used for. Currently we support setting only one type of solution at creation time. Only `SOLUTION_TYPE_SEARCH` value is supported at the moment.
        /// </summary>
        public readonly ImmutableArray<string> SolutionTypes;

        [OutputConstructor]
        private GetControlResult(
            ImmutableArray<string> associatedServingConfigIds,

            string displayName,

            Outputs.GoogleCloudRetailV2alphaSearchRequestFacetSpecResponse facetSpec,

            string name,

            Outputs.GoogleCloudRetailV2alphaRuleResponse rule,

            ImmutableArray<string> solutionTypes)
        {
            AssociatedServingConfigIds = associatedServingConfigIds;
            DisplayName = displayName;
            FacetSpec = facetSpec;
            Name = name;
            Rule = rule;
            SolutionTypes = solutionTypes;
        }
    }
}
