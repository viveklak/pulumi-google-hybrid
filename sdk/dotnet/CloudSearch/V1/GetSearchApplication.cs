// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.CloudSearch.V1
{
    public static class GetSearchApplication
    {
        /// <summary>
        /// Gets the specified search application. **Note:** This API requires an admin account to execute.
        /// </summary>
        public static Task<GetSearchApplicationResult> InvokeAsync(GetSearchApplicationArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetSearchApplicationResult>("google-native:cloudsearch/v1:getSearchApplication", args ?? new GetSearchApplicationArgs(), options.WithDefaults());

        /// <summary>
        /// Gets the specified search application. **Note:** This API requires an admin account to execute.
        /// </summary>
        public static Output<GetSearchApplicationResult> Invoke(GetSearchApplicationInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetSearchApplicationResult>("google-native:cloudsearch/v1:getSearchApplication", args ?? new GetSearchApplicationInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSearchApplicationArgs : Pulumi.InvokeArgs
    {
        [Input("debugOptionsEnableDebugging")]
        public string? DebugOptionsEnableDebugging { get; set; }

        [Input("searchapplicationId", required: true)]
        public string SearchapplicationId { get; set; } = null!;

        public GetSearchApplicationArgs()
        {
        }
    }

    public sealed class GetSearchApplicationInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("debugOptionsEnableDebugging")]
        public Input<string>? DebugOptionsEnableDebugging { get; set; }

        [Input("searchapplicationId", required: true)]
        public Input<string> SearchapplicationId { get; set; } = null!;

        public GetSearchApplicationInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetSearchApplicationResult
    {
        /// <summary>
        /// Retrictions applied to the configurations. The maximum number of elements is 10.
        /// </summary>
        public readonly ImmutableArray<Outputs.DataSourceRestrictionResponse> DataSourceRestrictions;
        /// <summary>
        /// The default fields for returning facet results. The sources specified here also have been included in data_source_restrictions above.
        /// </summary>
        public readonly ImmutableArray<Outputs.FacetOptionsResponse> DefaultFacetOptions;
        /// <summary>
        /// The default options for sorting the search results
        /// </summary>
        public readonly Outputs.SortOptionsResponse DefaultSortOptions;
        /// <summary>
        /// Display name of the Search Application. The maximum length is 300 characters.
        /// </summary>
        public readonly string DisplayName;
        /// <summary>
        /// Indicates whether audit logging is on/off for requests made for the search application in query APIs.
        /// </summary>
        public readonly bool EnableAuditLog;
        /// <summary>
        /// Name of the Search Application. Format: searchapplications/{application_id}.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// IDs of the Long Running Operations (LROs) currently running for this schema. Output only field.
        /// </summary>
        public readonly ImmutableArray<string> OperationIds;
        /// <summary>
        /// The default options for query interpretation
        /// </summary>
        public readonly Outputs.QueryInterpretationConfigResponse QueryInterpretationConfig;
        /// <summary>
        /// With each result we should return the URI for its thumbnail (when applicable)
        /// </summary>
        public readonly bool ReturnResultThumbnailUrls;
        /// <summary>
        /// Configuration for ranking results.
        /// </summary>
        public readonly Outputs.ScoringConfigResponse ScoringConfig;
        /// <summary>
        /// Configuration for a sources specified in data_source_restrictions.
        /// </summary>
        public readonly ImmutableArray<Outputs.SourceConfigResponse> SourceConfig;

        [OutputConstructor]
        private GetSearchApplicationResult(
            ImmutableArray<Outputs.DataSourceRestrictionResponse> dataSourceRestrictions,

            ImmutableArray<Outputs.FacetOptionsResponse> defaultFacetOptions,

            Outputs.SortOptionsResponse defaultSortOptions,

            string displayName,

            bool enableAuditLog,

            string name,

            ImmutableArray<string> operationIds,

            Outputs.QueryInterpretationConfigResponse queryInterpretationConfig,

            bool returnResultThumbnailUrls,

            Outputs.ScoringConfigResponse scoringConfig,

            ImmutableArray<Outputs.SourceConfigResponse> sourceConfig)
        {
            DataSourceRestrictions = dataSourceRestrictions;
            DefaultFacetOptions = defaultFacetOptions;
            DefaultSortOptions = defaultSortOptions;
            DisplayName = displayName;
            EnableAuditLog = enableAuditLog;
            Name = name;
            OperationIds = operationIds;
            QueryInterpretationConfig = queryInterpretationConfig;
            ReturnResultThumbnailUrls = returnResultThumbnailUrls;
            ScoringConfig = scoringConfig;
            SourceConfig = sourceConfig;
        }
    }
}
