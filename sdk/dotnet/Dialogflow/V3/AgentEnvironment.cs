// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleCloud.Dialogflow.V3
{
    /// <summary>
    /// Creates an Environment in the specified Agent.
    /// </summary>
    [GoogleCloudResourceType("google-cloud:dialogflow/v3:AgentEnvironment")]
    public partial class AgentEnvironment : Pulumi.CustomResource
    {
        /// <summary>
        /// Create a AgentEnvironment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AgentEnvironment(string name, AgentEnvironmentArgs args, CustomResourceOptions? options = null)
            : base("google-cloud:dialogflow/v3:AgentEnvironment", name, args ?? new AgentEnvironmentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AgentEnvironment(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-cloud:dialogflow/v3:AgentEnvironment", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing AgentEnvironment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AgentEnvironment Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new AgentEnvironment(name, id, options);
        }
    }

    public sealed class AgentEnvironmentArgs : Pulumi.ResourceArgs
    {
        [Input("agentsId", required: true)]
        public Input<string> AgentsId { get; set; } = null!;

        /// <summary>
        /// The human-readable description of the environment. The maximum length is 500 characters. If exceeded, the request is rejected.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Required. The human-readable name of the environment (unique in an agent). Limit of 64 characters.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        [Input("environmentsId", required: true)]
        public Input<string> EnvironmentsId { get; set; } = null!;

        [Input("locationsId", required: true)]
        public Input<string> LocationsId { get; set; } = null!;

        /// <summary>
        /// The name of the environment. Format: `projects//locations//agents//environments/`.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("projectsId", required: true)]
        public Input<string> ProjectsId { get; set; } = null!;

        /// <summary>
        /// Output only. Update time of this environment.
        /// </summary>
        [Input("updateTime")]
        public Input<string>? UpdateTime { get; set; }

        [Input("versionConfigs")]
        private InputList<Inputs.GoogleCloudDialogflowCxV3EnvironmentVersionConfigArgs>? _versionConfigs;

        /// <summary>
        /// Required. A list of configurations for flow versions. You should include version configs for all flows that are reachable from `Start Flow` in the agent. Otherwise, an error will be returned.
        /// </summary>
        public InputList<Inputs.GoogleCloudDialogflowCxV3EnvironmentVersionConfigArgs> VersionConfigs
        {
            get => _versionConfigs ?? (_versionConfigs = new InputList<Inputs.GoogleCloudDialogflowCxV3EnvironmentVersionConfigArgs>());
            set => _versionConfigs = value;
        }

        public AgentEnvironmentArgs()
        {
        }
    }
}
