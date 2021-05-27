// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dialogflow.V3Beta1
{
    /// <summary>
    /// Creates a session entity type.
    /// </summary>
    [GoogleNativeResourceType("google-native:dialogflow/v3beta1:AgentEnvironmentSessionEntityType")]
    public partial class AgentEnvironmentSessionEntityType : Pulumi.CustomResource
    {
        /// <summary>
        /// Required. The collection of entities to override or supplement the custom entity type.
        /// </summary>
        [Output("entities")]
        public Output<ImmutableArray<Outputs.GoogleCloudDialogflowCxV3beta1EntityTypeEntityResponse>> Entities { get; private set; } = null!;

        /// <summary>
        /// Required. Indicates whether the additional data should override or supplement the custom entity type definition.
        /// </summary>
        [Output("entityOverrideMode")]
        public Output<string> EntityOverrideMode { get; private set; } = null!;

        /// <summary>
        /// Required. The unique identifier of the session entity type. Format: `projects//locations//agents//sessions//entityTypes/` or `projects//locations//agents//environments//sessions//entityTypes/`. If `Environment ID` is not specified, we assume default 'draft' environment.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;


        /// <summary>
        /// Create a AgentEnvironmentSessionEntityType resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AgentEnvironmentSessionEntityType(string name, AgentEnvironmentSessionEntityTypeArgs args, CustomResourceOptions? options = null)
            : base("google-native:dialogflow/v3beta1:AgentEnvironmentSessionEntityType", name, args ?? new AgentEnvironmentSessionEntityTypeArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AgentEnvironmentSessionEntityType(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:dialogflow/v3beta1:AgentEnvironmentSessionEntityType", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing AgentEnvironmentSessionEntityType resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AgentEnvironmentSessionEntityType Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new AgentEnvironmentSessionEntityType(name, id, options);
        }
    }

    public sealed class AgentEnvironmentSessionEntityTypeArgs : Pulumi.ResourceArgs
    {
        [Input("agentId", required: true)]
        public Input<string> AgentId { get; set; } = null!;

        [Input("entities")]
        private InputList<Inputs.GoogleCloudDialogflowCxV3beta1EntityTypeEntityArgs>? _entities;

        /// <summary>
        /// Required. The collection of entities to override or supplement the custom entity type.
        /// </summary>
        public InputList<Inputs.GoogleCloudDialogflowCxV3beta1EntityTypeEntityArgs> Entities
        {
            get => _entities ?? (_entities = new InputList<Inputs.GoogleCloudDialogflowCxV3beta1EntityTypeEntityArgs>());
            set => _entities = value;
        }

        /// <summary>
        /// Required. Indicates whether the additional data should override or supplement the custom entity type definition.
        /// </summary>
        [Input("entityOverrideMode")]
        public Input<string>? EntityOverrideMode { get; set; }

        [Input("environmentId", required: true)]
        public Input<string> EnvironmentId { get; set; } = null!;

        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        /// <summary>
        /// Required. The unique identifier of the session entity type. Format: `projects//locations//agents//sessions//entityTypes/` or `projects//locations//agents//environments//sessions//entityTypes/`. If `Environment ID` is not specified, we assume default 'draft' environment.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("project", required: true)]
        public Input<string> Project { get; set; } = null!;

        [Input("sessionId", required: true)]
        public Input<string> SessionId { get; set; } = null!;

        public AgentEnvironmentSessionEntityTypeArgs()
        {
        }
    }
}
