// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.NetworkSecurity.V1Beta1
{
    /// <summary>
    /// Creates a new ServerTlsPolicy in a given project and location.
    /// </summary>
    [GoogleNativeResourceType("google-native:networksecurity/v1beta1:ServerTlsPolicy")]
    public partial class ServerTlsPolicy : Pulumi.CustomResource
    {
        /// <summary>
        ///  Determines if server allows plaintext connections. If set to true, server allows plain text connections. By default, it is set to false. This setting is not exclusive of other encryption modes. For example, if `allow_open` and `mtls_policy` are set, server allows both plain text and mTLS connections. See documentation of other encryption modes to confirm compatibility.
        /// </summary>
        [Output("allowOpen")]
        public Output<bool> AllowOpen { get; private set; } = null!;

        /// <summary>
        /// The timestamp when the resource was created.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// Free-text description of the resource.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// Set of label tags associated with the resource.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>> Labels { get; private set; } = null!;

        /// <summary>
        ///  Defines a mechanism to provision peer validation certificates for peer to peer authentication (Mutual TLS - mTLS). If not specified, client certificate will not be requested. The connection is treated as TLS and not mTLS. If `allow_open` and `mtls_policy` are set, server allows both plain text and mTLS connections.
        /// </summary>
        [Output("mtlsPolicy")]
        public Output<Outputs.MTLSPolicyResponse> MtlsPolicy { get; private set; } = null!;

        /// <summary>
        /// Name of the ServerTlsPolicy resource. It matches the pattern `projects/*/locations/{location}/serverTlsPolicies/{server_tls_policy}`
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        ///  Defines a mechanism to provision server identity (public and private keys). Cannot be combined with `allow_open` as a permissive mode that allows both plain text and TLS is not supported.
        /// </summary>
        [Output("serverCertificate")]
        public Output<Outputs.GoogleCloudNetworksecurityV1beta1CertificateProviderResponse> ServerCertificate { get; private set; } = null!;

        /// <summary>
        /// The timestamp when the resource was updated.
        /// </summary>
        [Output("updateTime")]
        public Output<string> UpdateTime { get; private set; } = null!;


        /// <summary>
        /// Create a ServerTlsPolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ServerTlsPolicy(string name, ServerTlsPolicyArgs args, CustomResourceOptions? options = null)
            : base("google-native:networksecurity/v1beta1:ServerTlsPolicy", name, args ?? new ServerTlsPolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ServerTlsPolicy(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:networksecurity/v1beta1:ServerTlsPolicy", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing ServerTlsPolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ServerTlsPolicy Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ServerTlsPolicy(name, id, options);
        }
    }

    public sealed class ServerTlsPolicyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        ///  Determines if server allows plaintext connections. If set to true, server allows plain text connections. By default, it is set to false. This setting is not exclusive of other encryption modes. For example, if `allow_open` and `mtls_policy` are set, server allows both plain text and mTLS connections. See documentation of other encryption modes to confirm compatibility.
        /// </summary>
        [Input("allowOpen")]
        public Input<bool>? AllowOpen { get; set; }

        /// <summary>
        /// Free-text description of the resource.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// Set of label tags associated with the resource.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        ///  Defines a mechanism to provision peer validation certificates for peer to peer authentication (Mutual TLS - mTLS). If not specified, client certificate will not be requested. The connection is treated as TLS and not mTLS. If `allow_open` and `mtls_policy` are set, server allows both plain text and mTLS connections.
        /// </summary>
        [Input("mtlsPolicy")]
        public Input<Inputs.MTLSPolicyArgs>? MtlsPolicy { get; set; }

        /// <summary>
        /// Name of the ServerTlsPolicy resource. It matches the pattern `projects/*/locations/{location}/serverTlsPolicies/{server_tls_policy}`
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        ///  Defines a mechanism to provision server identity (public and private keys). Cannot be combined with `allow_open` as a permissive mode that allows both plain text and TLS is not supported.
        /// </summary>
        [Input("serverCertificate")]
        public Input<Inputs.GoogleCloudNetworksecurityV1beta1CertificateProviderArgs>? ServerCertificate { get; set; }

        /// <summary>
        /// Required. Short name of the ServerTlsPolicy resource to be created. This value should be 1-63 characters long, containing only letters, numbers, hyphens, and underscores, and should not start with a number. E.g. "server_mtls_policy".
        /// </summary>
        [Input("serverTlsPolicyId", required: true)]
        public Input<string> ServerTlsPolicyId { get; set; } = null!;

        public ServerTlsPolicyArgs()
        {
        }
    }
}
