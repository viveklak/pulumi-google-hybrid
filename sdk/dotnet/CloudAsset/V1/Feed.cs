// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.CloudAsset.V1
{
    /// <summary>
    /// Creates a feed in a parent project/folder/organization to listen to its asset updates.
    /// Auto-naming is currently not supported for this resource.
    /// </summary>
    [GoogleNativeResourceType("google-native:cloudasset/v1:Feed")]
    public partial class Feed : Pulumi.CustomResource
    {
        /// <summary>
        /// A list of the full names of the assets to receive updates. You must specify either or both of asset_names and asset_types. Only asset updates matching specified asset_names or asset_types are exported to the feed. Example: `//compute.googleapis.com/projects/my_project_123/zones/zone1/instances/instance1`. For a list of the full names for supported asset types, see [Resource name format](/asset-inventory/docs/resource-name-format).
        /// </summary>
        [Output("assetNames")]
        public Output<ImmutableArray<string>> AssetNames { get; private set; } = null!;

        /// <summary>
        /// A list of types of the assets to receive updates. You must specify either or both of asset_names and asset_types. Only asset updates matching specified asset_names or asset_types are exported to the feed. Example: `"compute.googleapis.com/Disk"` For a list of all supported asset types, see [Supported asset types](/asset-inventory/docs/supported-asset-types).
        /// </summary>
        [Output("assetTypes")]
        public Output<ImmutableArray<string>> AssetTypes { get; private set; } = null!;

        /// <summary>
        /// A condition which determines whether an asset update should be published. If specified, an asset will be returned only when the expression evaluates to true. When set, `expression` field in the `Expr` must be a valid [CEL expression] (https://github.com/google/cel-spec) on a TemporalAsset with name `temporal_asset`. Example: a Feed with expression ("temporal_asset.deleted == true") will only publish Asset deletions. Other fields of `Expr` are optional. See our [user guide](https://cloud.google.com/asset-inventory/docs/monitoring-asset-changes-with-condition) for detailed instructions.
        /// </summary>
        [Output("condition")]
        public Output<Outputs.ExprResponse> Condition { get; private set; } = null!;

        /// <summary>
        /// Asset content type. If not specified, no content but the asset name and type will be returned.
        /// </summary>
        [Output("contentType")]
        public Output<string> ContentType { get; private set; } = null!;

        /// <summary>
        /// Feed output configuration defining where the asset updates are published to.
        /// </summary>
        [Output("feedOutputConfig")]
        public Output<Outputs.FeedOutputConfigResponse> FeedOutputConfig { get; private set; } = null!;

        /// <summary>
        /// The format will be projects/{project_number}/feeds/{client-assigned_feed_identifier} or folders/{folder_number}/feeds/{client-assigned_feed_identifier} or organizations/{organization_number}/feeds/{client-assigned_feed_identifier} The client-assigned feed identifier must be unique within the parent project/folder/organization.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// A list of relationship types to output, for example: `INSTANCE_TO_INSTANCEGROUP`. This field should only be specified if content_type=RELATIONSHIP. * If specified: it outputs specified relationship updates on the [asset_names] or the [asset_types]. It returns an error if any of the [relationship_types] doesn't belong to the supported relationship types of the [asset_names] or [asset_types], or any of the [asset_names] or the [asset_types] doesn't belong to the source types of the [relationship_types]. * Otherwise: it outputs the supported relationships of the types of [asset_names] and [asset_types] or returns an error if any of the [asset_names] or the [asset_types] has no replationship support. See [Introduction to Cloud Asset Inventory](https://cloud.google.com/asset-inventory/docs/overview) for all supported asset types and relationship types.
        /// </summary>
        [Output("relationshipTypes")]
        public Output<ImmutableArray<string>> RelationshipTypes { get; private set; } = null!;


        /// <summary>
        /// Create a Feed resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Feed(string name, FeedArgs args, CustomResourceOptions? options = null)
            : base("google-native:cloudasset/v1:Feed", name, args ?? new FeedArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Feed(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:cloudasset/v1:Feed", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Feed resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Feed Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Feed(name, id, options);
        }
    }

    public sealed class FeedArgs : Pulumi.ResourceArgs
    {
        [Input("assetNames")]
        private InputList<string>? _assetNames;

        /// <summary>
        /// A list of the full names of the assets to receive updates. You must specify either or both of asset_names and asset_types. Only asset updates matching specified asset_names or asset_types are exported to the feed. Example: `//compute.googleapis.com/projects/my_project_123/zones/zone1/instances/instance1`. For a list of the full names for supported asset types, see [Resource name format](/asset-inventory/docs/resource-name-format).
        /// </summary>
        public InputList<string> AssetNames
        {
            get => _assetNames ?? (_assetNames = new InputList<string>());
            set => _assetNames = value;
        }

        [Input("assetTypes")]
        private InputList<string>? _assetTypes;

        /// <summary>
        /// A list of types of the assets to receive updates. You must specify either or both of asset_names and asset_types. Only asset updates matching specified asset_names or asset_types are exported to the feed. Example: `"compute.googleapis.com/Disk"` For a list of all supported asset types, see [Supported asset types](/asset-inventory/docs/supported-asset-types).
        /// </summary>
        public InputList<string> AssetTypes
        {
            get => _assetTypes ?? (_assetTypes = new InputList<string>());
            set => _assetTypes = value;
        }

        /// <summary>
        /// A condition which determines whether an asset update should be published. If specified, an asset will be returned only when the expression evaluates to true. When set, `expression` field in the `Expr` must be a valid [CEL expression] (https://github.com/google/cel-spec) on a TemporalAsset with name `temporal_asset`. Example: a Feed with expression ("temporal_asset.deleted == true") will only publish Asset deletions. Other fields of `Expr` are optional. See our [user guide](https://cloud.google.com/asset-inventory/docs/monitoring-asset-changes-with-condition) for detailed instructions.
        /// </summary>
        [Input("condition")]
        public Input<Inputs.ExprArgs>? Condition { get; set; }

        /// <summary>
        /// Asset content type. If not specified, no content but the asset name and type will be returned.
        /// </summary>
        [Input("contentType")]
        public Input<Pulumi.GoogleNative.CloudAsset.V1.FeedContentType>? ContentType { get; set; }

        /// <summary>
        /// This is the client-assigned asset feed identifier and it needs to be unique under a specific parent project/folder/organization.
        /// </summary>
        [Input("feedId", required: true)]
        public Input<string> FeedId { get; set; } = null!;

        /// <summary>
        /// Feed output configuration defining where the asset updates are published to.
        /// </summary>
        [Input("feedOutputConfig", required: true)]
        public Input<Inputs.FeedOutputConfigArgs> FeedOutputConfig { get; set; } = null!;

        /// <summary>
        /// The format will be projects/{project_number}/feeds/{client-assigned_feed_identifier} or folders/{folder_number}/feeds/{client-assigned_feed_identifier} or organizations/{organization_number}/feeds/{client-assigned_feed_identifier} The client-assigned feed identifier must be unique within the parent project/folder/organization.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("relationshipTypes")]
        private InputList<string>? _relationshipTypes;

        /// <summary>
        /// A list of relationship types to output, for example: `INSTANCE_TO_INSTANCEGROUP`. This field should only be specified if content_type=RELATIONSHIP. * If specified: it outputs specified relationship updates on the [asset_names] or the [asset_types]. It returns an error if any of the [relationship_types] doesn't belong to the supported relationship types of the [asset_names] or [asset_types], or any of the [asset_names] or the [asset_types] doesn't belong to the source types of the [relationship_types]. * Otherwise: it outputs the supported relationships of the types of [asset_names] and [asset_types] or returns an error if any of the [asset_names] or the [asset_types] has no replationship support. See [Introduction to Cloud Asset Inventory](https://cloud.google.com/asset-inventory/docs/overview) for all supported asset types and relationship types.
        /// </summary>
        public InputList<string> RelationshipTypes
        {
            get => _relationshipTypes ?? (_relationshipTypes = new InputList<string>());
            set => _relationshipTypes = value;
        }

        [Input("v1Id", required: true)]
        public Input<string> V1Id { get; set; } = null!;

        [Input("v1Id1", required: true)]
        public Input<string> V1Id1 { get; set; } = null!;

        public FeedArgs()
        {
        }
    }
}
