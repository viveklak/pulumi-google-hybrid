// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.CloudAsset.V1
{
    public static class GetFeed
    {
        /// <summary>
        /// Gets details about an asset feed.
        /// </summary>
        public static Task<GetFeedResult> InvokeAsync(GetFeedArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetFeedResult>("google-native:cloudasset/v1:getFeed", args ?? new GetFeedArgs(), options.WithDefaults());

        /// <summary>
        /// Gets details about an asset feed.
        /// </summary>
        public static Output<GetFeedResult> Invoke(GetFeedInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetFeedResult>("google-native:cloudasset/v1:getFeed", args ?? new GetFeedInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetFeedArgs : Pulumi.InvokeArgs
    {
        [Input("feedId", required: true)]
        public string FeedId { get; set; } = null!;

        [Input("v1Id", required: true)]
        public string V1Id { get; set; } = null!;

        [Input("v1Id1", required: true)]
        public string V1Id1 { get; set; } = null!;

        public GetFeedArgs()
        {
        }
    }

    public sealed class GetFeedInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("feedId", required: true)]
        public Input<string> FeedId { get; set; } = null!;

        [Input("v1Id", required: true)]
        public Input<string> V1Id { get; set; } = null!;

        [Input("v1Id1", required: true)]
        public Input<string> V1Id1 { get; set; } = null!;

        public GetFeedInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetFeedResult
    {
        /// <summary>
        /// A list of the full names of the assets to receive updates. You must specify either or both of asset_names and asset_types. Only asset updates matching specified asset_names or asset_types are exported to the feed. Example: `//compute.googleapis.com/projects/my_project_123/zones/zone1/instances/instance1`. For a list of the full names for supported asset types, see [Resource name format](/asset-inventory/docs/resource-name-format).
        /// </summary>
        public readonly ImmutableArray<string> AssetNames;
        /// <summary>
        /// A list of types of the assets to receive updates. You must specify either or both of asset_names and asset_types. Only asset updates matching specified asset_names or asset_types are exported to the feed. Example: `"compute.googleapis.com/Disk"` For a list of all supported asset types, see [Supported asset types](/asset-inventory/docs/supported-asset-types).
        /// </summary>
        public readonly ImmutableArray<string> AssetTypes;
        /// <summary>
        /// A condition which determines whether an asset update should be published. If specified, an asset will be returned only when the expression evaluates to true. When set, `expression` field in the `Expr` must be a valid [CEL expression] (https://github.com/google/cel-spec) on a TemporalAsset with name `temporal_asset`. Example: a Feed with expression ("temporal_asset.deleted == true") will only publish Asset deletions. Other fields of `Expr` are optional. See our [user guide](https://cloud.google.com/asset-inventory/docs/monitoring-asset-changes-with-condition) for detailed instructions.
        /// </summary>
        public readonly Outputs.ExprResponse Condition;
        /// <summary>
        /// Asset content type. If not specified, no content but the asset name and type will be returned.
        /// </summary>
        public readonly string ContentType;
        /// <summary>
        /// Feed output configuration defining where the asset updates are published to.
        /// </summary>
        public readonly Outputs.FeedOutputConfigResponse FeedOutputConfig;
        /// <summary>
        /// The format will be projects/{project_number}/feeds/{client-assigned_feed_identifier} or folders/{folder_number}/feeds/{client-assigned_feed_identifier} or organizations/{organization_number}/feeds/{client-assigned_feed_identifier} The client-assigned feed identifier must be unique within the parent project/folder/organization.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// A list of relationship types to output, for example: `INSTANCE_TO_INSTANCEGROUP`. This field should only be specified if content_type=RELATIONSHIP. * If specified: it outputs specified relationship updates on the [asset_names] or the [asset_types]. It returns an error if any of the [relationship_types] doesn't belong to the supported relationship types of the [asset_names] or [asset_types], or any of the [asset_names] or the [asset_types] doesn't belong to the source types of the [relationship_types]. * Otherwise: it outputs the supported relationships of the types of [asset_names] and [asset_types] or returns an error if any of the [asset_names] or the [asset_types] has no replationship support. See [Introduction to Cloud Asset Inventory](https://cloud.google.com/asset-inventory/docs/overview) for all supported asset types and relationship types.
        /// </summary>
        public readonly ImmutableArray<string> RelationshipTypes;

        [OutputConstructor]
        private GetFeedResult(
            ImmutableArray<string> assetNames,

            ImmutableArray<string> assetTypes,

            Outputs.ExprResponse condition,

            string contentType,

            Outputs.FeedOutputConfigResponse feedOutputConfig,

            string name,

            ImmutableArray<string> relationshipTypes)
        {
            AssetNames = assetNames;
            AssetTypes = assetTypes;
            Condition = condition;
            ContentType = contentType;
            FeedOutputConfig = feedOutputConfig;
            Name = name;
            RelationshipTypes = relationshipTypes;
        }
    }
}
