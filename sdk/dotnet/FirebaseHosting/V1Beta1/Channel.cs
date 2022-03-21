// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.FirebaseHosting.V1Beta1
{
    /// <summary>
    /// Creates a new channel in the specified site.
    /// </summary>
    [GoogleNativeResourceType("google-native:firebasehosting/v1beta1:Channel")]
    public partial class Channel : Pulumi.CustomResource
    {
        /// <summary>
        /// The time at which the channel was created.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// The time at which the channel will be automatically deleted. If null, the channel will not be automatically deleted. This field is present in the output whether it's set directly or via the `ttl` field.
        /// </summary>
        [Output("expireTime")]
        public Output<string> ExpireTime { get; private set; } = null!;

        /// <summary>
        /// Text labels used for extra metadata and/or filtering.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>> Labels { get; private set; } = null!;

        /// <summary>
        /// The fully-qualified resource name for the channel, in the format: sites/ SITE_ID/channels/CHANNEL_ID
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The current release for the channel, if any.
        /// </summary>
        [Output("release")]
        public Output<Outputs.ReleaseResponse> Release { get; private set; } = null!;

        /// <summary>
        /// The number of previous releases to retain on the channel for rollback or other purposes. Must be a number between 1-100. Defaults to 10 for new channels.
        /// </summary>
        [Output("retainedReleaseCount")]
        public Output<int> RetainedReleaseCount { get; private set; } = null!;

        /// <summary>
        /// Input only. A time-to-live for this channel. Sets `expire_time` to the provided duration past the time of the request.
        /// </summary>
        [Output("ttl")]
        public Output<string> Ttl { get; private set; } = null!;

        /// <summary>
        /// The time at which the channel was last updated.
        /// </summary>
        [Output("updateTime")]
        public Output<string> UpdateTime { get; private set; } = null!;

        /// <summary>
        /// The URL at which the content of this channel's current release can be viewed. This URL is a Firebase-provided subdomain of `web.app`. The content of this channel's current release can also be viewed at the Firebase-provided subdomain of `firebaseapp.com`. If this channel is the `live` channel for the Hosting site, then the content of this channel's current release can also be viewed at any connected custom domains.
        /// </summary>
        [Output("url")]
        public Output<string> Url { get; private set; } = null!;


        /// <summary>
        /// Create a Channel resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Channel(string name, ChannelArgs args, CustomResourceOptions? options = null)
            : base("google-native:firebasehosting/v1beta1:Channel", name, args ?? new ChannelArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Channel(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:firebasehosting/v1beta1:Channel", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing Channel resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Channel Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Channel(name, id, options);
        }
    }

    public sealed class ChannelArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Required. Immutable. A unique ID within the site that identifies the channel.
        /// </summary>
        [Input("channelId", required: true)]
        public Input<string> ChannelId { get; set; } = null!;

        /// <summary>
        /// The time at which the channel will be automatically deleted. If null, the channel will not be automatically deleted. This field is present in the output whether it's set directly or via the `ttl` field.
        /// </summary>
        [Input("expireTime")]
        public Input<string>? ExpireTime { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// Text labels used for extra metadata and/or filtering.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// The fully-qualified resource name for the channel, in the format: sites/ SITE_ID/channels/CHANNEL_ID
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The number of previous releases to retain on the channel for rollback or other purposes. Must be a number between 1-100. Defaults to 10 for new channels.
        /// </summary>
        [Input("retainedReleaseCount")]
        public Input<int>? RetainedReleaseCount { get; set; }

        [Input("siteId", required: true)]
        public Input<string> SiteId { get; set; } = null!;

        /// <summary>
        /// Input only. A time-to-live for this channel. Sets `expire_time` to the provided duration past the time of the request.
        /// </summary>
        [Input("ttl")]
        public Input<string>? Ttl { get; set; }

        public ChannelArgs()
        {
        }
    }
}
