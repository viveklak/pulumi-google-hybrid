// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.CloudIdentity.V1Beta1
{
    /// <summary>
    /// Creates a `Group`.
    /// </summary>
    [GoogleNativeResourceType("google-native:cloudidentity/v1beta1:Group")]
    public partial class Group : Pulumi.CustomResource
    {
        /// <summary>
        /// Additional entity key aliases for a Group.
        /// </summary>
        [Output("additionalGroupKeys")]
        public Output<ImmutableArray<Outputs.EntityKeyResponse>> AdditionalGroupKeys { get; private set; } = null!;

        /// <summary>
        /// The time when the `Group` was created.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// An extended description to help users determine the purpose of a `Group`. Must not be longer than 4,096 characters.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// The display name of the `Group`.
        /// </summary>
        [Output("displayName")]
        public Output<string> DisplayName { get; private set; } = null!;

        /// <summary>
        /// Optional. Dynamic group metadata like queries and status.
        /// </summary>
        [Output("dynamicGroupMetadata")]
        public Output<Outputs.DynamicGroupMetadataResponse> DynamicGroupMetadata { get; private set; } = null!;

        /// <summary>
        /// Immutable. The `EntityKey` of the `Group`.
        /// </summary>
        [Output("groupKey")]
        public Output<Outputs.EntityKeyResponse> GroupKey { get; private set; } = null!;

        /// <summary>
        /// One or more label entries that apply to the Group. Currently supported labels contain a key with an empty value. Google Groups are the default type of group and have a label with a key of `cloudidentity.googleapis.com/groups.discussion_forum` and an empty value. Existing Google Groups can have an additional label with a key of `cloudidentity.googleapis.com/groups.security` and an empty value added to them. **This is an immutable change and the security label cannot be removed once added.** Dynamic groups have a label with a key of `cloudidentity.googleapis.com/groups.dynamic`. Identity-mapped groups for Cloud Search have a label with a key of `system/groups/external` and an empty value. Examples: {"cloudidentity.googleapis.com/groups.discussion_forum": ""} or {"system/groups/external": ""}.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>> Labels { get; private set; } = null!;

        /// <summary>
        /// The [resource name](https://cloud.google.com/apis/design/resource_names) of the `Group`. Shall be of the form `groups/{group_id}`.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Immutable. The resource name of the entity under which this `Group` resides in the Cloud Identity resource hierarchy. Must be of the form `identitysources/{identity_source_id}` for external- identity-mapped groups or `customers/{customer_id}` for Google Groups.
        /// </summary>
        [Output("parent")]
        public Output<string> Parent { get; private set; } = null!;

        /// <summary>
        /// The time when the `Group` was last updated.
        /// </summary>
        [Output("updateTime")]
        public Output<string> UpdateTime { get; private set; } = null!;


        /// <summary>
        /// Create a Group resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Group(string name, GroupArgs args, CustomResourceOptions? options = null)
            : base("google-native:cloudidentity/v1beta1:Group", name, args ?? new GroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Group(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:cloudidentity/v1beta1:Group", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing Group resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Group Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Group(name, id, options);
        }
    }

    public sealed class GroupArgs : Pulumi.ResourceArgs
    {
        [Input("additionalGroupKeys")]
        private InputList<Inputs.EntityKeyArgs>? _additionalGroupKeys;

        /// <summary>
        /// Additional entity key aliases for a Group.
        /// </summary>
        public InputList<Inputs.EntityKeyArgs> AdditionalGroupKeys
        {
            get => _additionalGroupKeys ?? (_additionalGroupKeys = new InputList<Inputs.EntityKeyArgs>());
            set => _additionalGroupKeys = value;
        }

        /// <summary>
        /// An extended description to help users determine the purpose of a `Group`. Must not be longer than 4,096 characters.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The display name of the `Group`.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// Optional. Dynamic group metadata like queries and status.
        /// </summary>
        [Input("dynamicGroupMetadata")]
        public Input<Inputs.DynamicGroupMetadataArgs>? DynamicGroupMetadata { get; set; }

        /// <summary>
        /// Immutable. The `EntityKey` of the `Group`.
        /// </summary>
        [Input("groupKey", required: true)]
        public Input<Inputs.EntityKeyArgs> GroupKey { get; set; } = null!;

        [Input("initialGroupConfig", required: true)]
        public Input<string> InitialGroupConfig { get; set; } = null!;

        [Input("labels", required: true)]
        private InputMap<string>? _labels;

        /// <summary>
        /// One or more label entries that apply to the Group. Currently supported labels contain a key with an empty value. Google Groups are the default type of group and have a label with a key of `cloudidentity.googleapis.com/groups.discussion_forum` and an empty value. Existing Google Groups can have an additional label with a key of `cloudidentity.googleapis.com/groups.security` and an empty value added to them. **This is an immutable change and the security label cannot be removed once added.** Dynamic groups have a label with a key of `cloudidentity.googleapis.com/groups.dynamic`. Identity-mapped groups for Cloud Search have a label with a key of `system/groups/external` and an empty value. Examples: {"cloudidentity.googleapis.com/groups.discussion_forum": ""} or {"system/groups/external": ""}.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// Immutable. The resource name of the entity under which this `Group` resides in the Cloud Identity resource hierarchy. Must be of the form `identitysources/{identity_source_id}` for external- identity-mapped groups or `customers/{customer_id}` for Google Groups.
        /// </summary>
        [Input("parent", required: true)]
        public Input<string> Parent { get; set; } = null!;

        public GroupArgs()
        {
        }
    }
}
