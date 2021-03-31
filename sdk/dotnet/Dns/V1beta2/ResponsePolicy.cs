// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleCloud.Dns.V1beta2
{
    /// <summary>
    /// Creates a new Response Policy
    /// </summary>
    [GoogleCloudResourceType("google-cloud:dns/v1beta2:ResponsePolicy")]
    public partial class ResponsePolicy : Pulumi.CustomResource
    {
        /// <summary>
        /// Create a ResponsePolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ResponsePolicy(string name, ResponsePolicyArgs args, CustomResourceOptions? options = null)
            : base("google-cloud:dns/v1beta2:ResponsePolicy", name, args ?? new ResponsePolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ResponsePolicy(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-cloud:dns/v1beta2:ResponsePolicy", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing ResponsePolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ResponsePolicy Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ResponsePolicy(name, id, options);
        }
    }

    public sealed class ResponsePolicyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// User-provided description for this Response Policy.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Unique identifier for the resource; defined by the server (output only).
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        [Input("kind")]
        public Input<string>? Kind { get; set; }

        [Input("networks")]
        private InputList<Inputs.ResponsePolicyNetworkArgs>? _networks;

        /// <summary>
        /// List of network names specifying networks to which this policy is applied.
        /// </summary>
        public InputList<Inputs.ResponsePolicyNetworkArgs> Networks
        {
            get => _networks ?? (_networks = new InputList<Inputs.ResponsePolicyNetworkArgs>());
            set => _networks = value;
        }

        [Input("project", required: true)]
        public Input<string> Project { get; set; } = null!;

        [Input("responsePolicy", required: true)]
        public Input<string> ResponsePolicy { get; set; } = null!;

        /// <summary>
        /// User assigned name for this Response Policy.
        /// </summary>
        [Input("responsePolicyName")]
        public Input<string>? ResponsePolicyName { get; set; }

        public ResponsePolicyArgs()
        {
        }
    }
}
