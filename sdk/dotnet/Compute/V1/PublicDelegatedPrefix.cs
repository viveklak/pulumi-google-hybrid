// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.V1
{
    /// <summary>
    /// Creates a PublicDelegatedPrefix in the specified project in the given region using the parameters that are included in the request.
    /// </summary>
    [GoogleNativeResourceType("google-native:compute/v1:PublicDelegatedPrefix")]
    public partial class PublicDelegatedPrefix : Pulumi.CustomResource
    {
        /// <summary>
        /// [Output Only] Creation timestamp in RFC3339 text format.
        /// </summary>
        [Output("creationTimestamp")]
        public Output<string> CreationTimestamp { get; private set; } = null!;

        /// <summary>
        /// An optional description of this resource. Provide this property when you create the resource.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking. This field will be ignored when inserting a new PublicDelegatedPrefix. An up-to-date fingerprint must be provided in order to update the PublicDelegatedPrefix, otherwise the request will fail with error 412 conditionNotMet.
        /// 
        /// To see the latest fingerprint, make a get() request to retrieve a PublicDelegatedPrefix.
        /// </summary>
        [Output("fingerprint")]
        public Output<string> Fingerprint { get; private set; } = null!;

        /// <summary>
        /// The IPv4 address range, in CIDR format, represented by this public delegated prefix.
        /// </summary>
        [Output("ipCidrRange")]
        public Output<string> IpCidrRange { get; private set; } = null!;

        /// <summary>
        /// If true, the prefix will be live migrated.
        /// </summary>
        [Output("isLiveMigration")]
        public Output<bool> IsLiveMigration { get; private set; } = null!;

        /// <summary>
        /// [Output Only] Type of the resource. Always compute#publicDelegatedPrefix for public delegated prefixes.
        /// </summary>
        [Output("kind")]
        public Output<string> Kind { get; private set; } = null!;

        /// <summary>
        /// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The URL of parent prefix. Either PublicAdvertisedPrefix or PublicDelegatedPrefix.
        /// </summary>
        [Output("parentPrefix")]
        public Output<string> ParentPrefix { get; private set; } = null!;

        /// <summary>
        /// The list of sub public delegated prefixes that exist for this public delegated prefix.
        /// </summary>
        [Output("publicDelegatedSubPrefixs")]
        public Output<ImmutableArray<Outputs.PublicDelegatedPrefixPublicDelegatedSubPrefixResponse>> PublicDelegatedSubPrefixs { get; private set; } = null!;

        /// <summary>
        /// [Output Only] URL of the region where the public delegated prefix resides. This field applies only to the region resource. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// [Output Only] Server-defined URL for the resource.
        /// </summary>
        [Output("selfLink")]
        public Output<string> SelfLink { get; private set; } = null!;

        /// <summary>
        /// [Output Only] The status of the public delegated prefix.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;


        /// <summary>
        /// Create a PublicDelegatedPrefix resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public PublicDelegatedPrefix(string name, PublicDelegatedPrefixArgs args, CustomResourceOptions? options = null)
            : base("google-native:compute/v1:PublicDelegatedPrefix", name, args ?? new PublicDelegatedPrefixArgs(), MakeResourceOptions(options, ""))
        {
        }

        private PublicDelegatedPrefix(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:compute/v1:PublicDelegatedPrefix", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing PublicDelegatedPrefix resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static PublicDelegatedPrefix Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new PublicDelegatedPrefix(name, id, options);
        }
    }

    public sealed class PublicDelegatedPrefixArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// [Output Only] Creation timestamp in RFC3339 text format.
        /// </summary>
        [Input("creationTimestamp")]
        public Input<string>? CreationTimestamp { get; set; }

        /// <summary>
        /// An optional description of this resource. Provide this property when you create the resource.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking. This field will be ignored when inserting a new PublicDelegatedPrefix. An up-to-date fingerprint must be provided in order to update the PublicDelegatedPrefix, otherwise the request will fail with error 412 conditionNotMet.
        /// 
        /// To see the latest fingerprint, make a get() request to retrieve a PublicDelegatedPrefix.
        /// </summary>
        [Input("fingerprint")]
        public Input<string>? Fingerprint { get; set; }

        /// <summary>
        /// [Output Only] The unique identifier for the resource type. The server generates this identifier.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// The IPv4 address range, in CIDR format, represented by this public delegated prefix.
        /// </summary>
        [Input("ipCidrRange")]
        public Input<string>? IpCidrRange { get; set; }

        /// <summary>
        /// If true, the prefix will be live migrated.
        /// </summary>
        [Input("isLiveMigration")]
        public Input<bool>? IsLiveMigration { get; set; }

        /// <summary>
        /// [Output Only] Type of the resource. Always compute#publicDelegatedPrefix for public delegated prefixes.
        /// </summary>
        [Input("kind")]
        public Input<string>? Kind { get; set; }

        /// <summary>
        /// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The URL of parent prefix. Either PublicAdvertisedPrefix or PublicDelegatedPrefix.
        /// </summary>
        [Input("parentPrefix")]
        public Input<string>? ParentPrefix { get; set; }

        [Input("project", required: true)]
        public Input<string> Project { get; set; } = null!;

        [Input("publicDelegatedSubPrefixs")]
        private InputList<Inputs.PublicDelegatedPrefixPublicDelegatedSubPrefixArgs>? _publicDelegatedSubPrefixs;

        /// <summary>
        /// The list of sub public delegated prefixes that exist for this public delegated prefix.
        /// </summary>
        public InputList<Inputs.PublicDelegatedPrefixPublicDelegatedSubPrefixArgs> PublicDelegatedSubPrefixs
        {
            get => _publicDelegatedSubPrefixs ?? (_publicDelegatedSubPrefixs = new InputList<Inputs.PublicDelegatedPrefixPublicDelegatedSubPrefixArgs>());
            set => _publicDelegatedSubPrefixs = value;
        }

        /// <summary>
        /// [Output Only] URL of the region where the public delegated prefix resides. This field applies only to the region resource. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
        /// </summary>
        [Input("region", required: true)]
        public Input<string> Region { get; set; } = null!;

        [Input("requestId")]
        public Input<string>? RequestId { get; set; }

        /// <summary>
        /// [Output Only] Server-defined URL for the resource.
        /// </summary>
        [Input("selfLink")]
        public Input<string>? SelfLink { get; set; }

        /// <summary>
        /// [Output Only] The status of the public delegated prefix.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public PublicDelegatedPrefixArgs()
        {
        }
    }
}
