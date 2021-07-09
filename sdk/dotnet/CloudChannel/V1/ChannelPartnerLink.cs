// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.CloudChannel.V1
{
    /// <summary>
    /// Initiates a channel partner link between a distributor and a reseller, or between resellers in an n-tier reseller channel. Invited partners need to follow the invite_link_uri provided in the response to accept. After accepting the invitation, a link is set up between the two parties. You must be a distributor to call this method. Possible error codes: * PERMISSION_DENIED: The reseller account making the request is different from the reseller account in the API request. * INVALID_ARGUMENT: Required request parameters are missing or invalid. * ALREADY_EXISTS: The ChannelPartnerLink sent in the request already exists. * NOT_FOUND: No Cloud Identity customer exists for provided domain. * INTERNAL: Any non-user error related to a technical issue in the backend. Contact Cloud Channel support. * UNKNOWN: Any non-user error related to a technical issue in the backend. Contact Cloud Channel support. Return value: The new ChannelPartnerLink resource.
    /// </summary>
    [GoogleNativeResourceType("google-native:cloudchannel/v1:ChannelPartnerLink")]
    public partial class ChannelPartnerLink : Pulumi.CustomResource
    {
        /// <summary>
        /// Cloud Identity info of the channel partner (IR).
        /// </summary>
        [Output("channelPartnerCloudIdentityInfo")]
        public Output<Outputs.GoogleCloudChannelV1CloudIdentityInfoResponse> ChannelPartnerCloudIdentityInfo { get; private set; } = null!;

        /// <summary>
        /// Timestamp of when the channel partner link is created.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// URI of the web page where partner accepts the link invitation.
        /// </summary>
        [Output("inviteLinkUri")]
        public Output<string> InviteLinkUri { get; private set; } = null!;

        /// <summary>
        /// State of the channel partner link.
        /// </summary>
        [Output("linkState")]
        public Output<string> LinkState { get; private set; } = null!;

        /// <summary>
        /// Resource name for the channel partner link, in the format accounts/{account_id}/channelPartnerLinks/{id}.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Public identifier that a customer must use to generate a transfer token to move to this distributor-reseller combination.
        /// </summary>
        [Output("publicId")]
        public Output<string> PublicId { get; private set; } = null!;

        /// <summary>
        /// Cloud Identity ID of the linked reseller.
        /// </summary>
        [Output("resellerCloudIdentityId")]
        public Output<string> ResellerCloudIdentityId { get; private set; } = null!;

        /// <summary>
        /// Timestamp of when the channel partner link is updated.
        /// </summary>
        [Output("updateTime")]
        public Output<string> UpdateTime { get; private set; } = null!;


        /// <summary>
        /// Create a ChannelPartnerLink resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ChannelPartnerLink(string name, ChannelPartnerLinkArgs args, CustomResourceOptions? options = null)
            : base("google-native:cloudchannel/v1:ChannelPartnerLink", name, args ?? new ChannelPartnerLinkArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ChannelPartnerLink(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:cloudchannel/v1:ChannelPartnerLink", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing ChannelPartnerLink resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ChannelPartnerLink Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ChannelPartnerLink(name, id, options);
        }
    }

    public sealed class ChannelPartnerLinkArgs : Pulumi.ResourceArgs
    {
        [Input("accountId", required: true)]
        public Input<string> AccountId { get; set; } = null!;

        /// <summary>
        /// State of the channel partner link.
        /// </summary>
        [Input("linkState", required: true)]
        public Input<Pulumi.GoogleNative.CloudChannel.V1.ChannelPartnerLinkLinkState> LinkState { get; set; } = null!;

        /// <summary>
        /// Cloud Identity ID of the linked reseller.
        /// </summary>
        [Input("resellerCloudIdentityId", required: true)]
        public Input<string> ResellerCloudIdentityId { get; set; } = null!;

        public ChannelPartnerLinkArgs()
        {
        }
    }
}
