// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.StorageTransfer.V1
{
    /// <summary>
    /// Creates an agent pool resource.
    /// </summary>
    [GoogleNativeResourceType("google-native:storagetransfer/v1:AgentPool")]
    public partial class AgentPool : Pulumi.CustomResource
    {
        /// <summary>
        /// Specifies the bandwidth limit details. If this field is unspecified, the default value is set as 'No Limit'.
        /// </summary>
        [Output("bandwidthLimit")]
        public Output<Outputs.BandwidthLimitResponse> BandwidthLimit { get; private set; } = null!;

        /// <summary>
        /// Specifies the client-specified AgentPool description.
        /// </summary>
        [Output("displayName")]
        public Output<string> DisplayName { get; private set; } = null!;

        /// <summary>
        /// Specifies a unique string that identifies the agent pool. Format: `projects/{project_id}/agentPools/{agent_pool_id}`
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Specifies the state of the AgentPool.
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;


        /// <summary>
        /// Create a AgentPool resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AgentPool(string name, AgentPoolArgs args, CustomResourceOptions? options = null)
            : base("google-native:storagetransfer/v1:AgentPool", name, args ?? new AgentPoolArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AgentPool(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:storagetransfer/v1:AgentPool", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing AgentPool resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AgentPool Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new AgentPool(name, id, options);
        }
    }

    public sealed class AgentPoolArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Required. The ID of the agent pool to create. The `agent_pool_id` must meet the following requirements: * Length of 128 characters or less. * Not start with the string `goog`. * Start with a lowercase ASCII character, followed by: * Zero or more: lowercase Latin alphabet characters, numerals, hyphens (`-`), periods (`.`), underscores (`_`), or tildes (`~`). * One or more numerals or lowercase ASCII characters. As expressed by the regular expression: `^(?!goog)[a-z]([a-z0-9-._~]*[a-z0-9])?$`.
        /// </summary>
        [Input("agentPoolId", required: true)]
        public Input<string> AgentPoolId { get; set; } = null!;

        /// <summary>
        /// Specifies the bandwidth limit details. If this field is unspecified, the default value is set as 'No Limit'.
        /// </summary>
        [Input("bandwidthLimit")]
        public Input<Inputs.BandwidthLimitArgs>? BandwidthLimit { get; set; }

        /// <summary>
        /// Specifies the client-specified AgentPool description.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// Specifies a unique string that identifies the agent pool. Format: `projects/{project_id}/agentPools/{agent_pool_id}`
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("project")]
        public Input<string>? Project { get; set; }

        public AgentPoolArgs()
        {
        }
    }
}