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
    /// Creates a `Membership`.
    /// Auto-naming is currently not supported for this resource.
    /// </summary>
    [GoogleNativeResourceType("google-native:cloudidentity/v1beta1:Membership")]
    public partial class Membership : Pulumi.CustomResource
    {
        /// <summary>
        /// The time when the `Membership` was created.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// Immutable. The `EntityKey` of the member. Either `member_key` or `preferred_member_key` must be set when calling MembershipsService.CreateMembership but not both; both shall be set when returned.
        /// </summary>
        [Output("memberKey")]
        public Output<Outputs.EntityKeyResponse> MemberKey { get; private set; } = null!;

        /// <summary>
        /// The [resource name](https://cloud.google.com/apis/design/resource_names) of the `Membership`. Shall be of the form `groups/{group_id}/memberships/{membership_id}`.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Immutable. The `EntityKey` of the member. Either `member_key` or `preferred_member_key` must be set when calling MembershipsService.CreateMembership but not both; both shall be set when returned.
        /// </summary>
        [Output("preferredMemberKey")]
        public Output<Outputs.EntityKeyResponse> PreferredMemberKey { get; private set; } = null!;

        /// <summary>
        /// The `MembershipRole`s that apply to the `Membership`. If unspecified, defaults to a single `MembershipRole` with `name` `MEMBER`. Must not contain duplicate `MembershipRole`s with the same `name`.
        /// </summary>
        [Output("roles")]
        public Output<ImmutableArray<Outputs.MembershipRoleResponse>> Roles { get; private set; } = null!;

        /// <summary>
        /// The type of the membership.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// The time when the `Membership` was last updated.
        /// </summary>
        [Output("updateTime")]
        public Output<string> UpdateTime { get; private set; } = null!;


        /// <summary>
        /// Create a Membership resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Membership(string name, MembershipArgs args, CustomResourceOptions? options = null)
            : base("google-native:cloudidentity/v1beta1:Membership", name, args ?? new MembershipArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Membership(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:cloudidentity/v1beta1:Membership", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing Membership resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Membership Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Membership(name, id, options);
        }
    }

    public sealed class MembershipArgs : Pulumi.ResourceArgs
    {
        [Input("groupId", required: true)]
        public Input<string> GroupId { get; set; } = null!;

        /// <summary>
        /// Immutable. The `EntityKey` of the member. Either `member_key` or `preferred_member_key` must be set when calling MembershipsService.CreateMembership but not both; both shall be set when returned.
        /// </summary>
        [Input("memberKey")]
        public Input<Inputs.EntityKeyArgs>? MemberKey { get; set; }

        /// <summary>
        /// Immutable. The `EntityKey` of the member. Either `member_key` or `preferred_member_key` must be set when calling MembershipsService.CreateMembership but not both; both shall be set when returned.
        /// </summary>
        [Input("preferredMemberKey", required: true)]
        public Input<Inputs.EntityKeyArgs> PreferredMemberKey { get; set; } = null!;

        [Input("roles")]
        private InputList<Inputs.MembershipRoleArgs>? _roles;

        /// <summary>
        /// The `MembershipRole`s that apply to the `Membership`. If unspecified, defaults to a single `MembershipRole` with `name` `MEMBER`. Must not contain duplicate `MembershipRole`s with the same `name`.
        /// </summary>
        public InputList<Inputs.MembershipRoleArgs> Roles
        {
            get => _roles ?? (_roles = new InputList<Inputs.MembershipRoleArgs>());
            set => _roles = value;
        }

        public MembershipArgs()
        {
        }
    }
}
