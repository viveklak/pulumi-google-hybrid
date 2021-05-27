// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dialogflow.V2Beta1
{
    /// <summary>
    /// Creates a session entity type. If the specified session entity type already exists, overrides the session entity type. This method doesn't work with Google Assistant integration. Contact Dialogflow support if you need to use session entities with Google Assistant integration.
    /// </summary>
    [GoogleNativeResourceType("google-native:dialogflow/v2beta1:AgentEnvironmentUserSessionEntityType")]
    public partial class AgentEnvironmentUserSessionEntityType : Pulumi.CustomResource
    {
        /// <summary>
        /// Required. The collection of entities associated with this session entity type.
        /// </summary>
        [Output("entities")]
        public Output<ImmutableArray<Outputs.GoogleCloudDialogflowV2beta1EntityTypeEntityResponse>> Entities { get; private set; } = null!;

        /// <summary>
        /// Required. Indicates whether the additional data should override or supplement the custom entity type definition.
        /// </summary>
        [Output("entityOverrideMode")]
        public Output<string> EntityOverrideMode { get; private set; } = null!;

        /// <summary>
        /// Required. The unique identifier of this session entity type. Supported formats: - `projects//agent/sessions//entityTypes/` - `projects//locations//agent/sessions//entityTypes/` - `projects//agent/environments//users//sessions//entityTypes/` - `projects//locations//agent/environments/ /users//sessions//entityTypes/` If `Location ID` is not specified we assume default 'us' location. If `Environment ID` is not specified, we assume default 'draft' environment. If `User ID` is not specified, we assume default '-' user. `` must be the display name of an existing entity type in the same agent that will be overridden or supplemented.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;


        /// <summary>
        /// Create a AgentEnvironmentUserSessionEntityType resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AgentEnvironmentUserSessionEntityType(string name, AgentEnvironmentUserSessionEntityTypeArgs args, CustomResourceOptions? options = null)
            : base("google-native:dialogflow/v2beta1:AgentEnvironmentUserSessionEntityType", name, args ?? new AgentEnvironmentUserSessionEntityTypeArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AgentEnvironmentUserSessionEntityType(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:dialogflow/v2beta1:AgentEnvironmentUserSessionEntityType", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing AgentEnvironmentUserSessionEntityType resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AgentEnvironmentUserSessionEntityType Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new AgentEnvironmentUserSessionEntityType(name, id, options);
        }
    }

    public sealed class AgentEnvironmentUserSessionEntityTypeArgs : Pulumi.ResourceArgs
    {
        [Input("entities")]
        private InputList<Inputs.GoogleCloudDialogflowV2beta1EntityTypeEntityArgs>? _entities;

        /// <summary>
        /// Required. The collection of entities associated with this session entity type.
        /// </summary>
        public InputList<Inputs.GoogleCloudDialogflowV2beta1EntityTypeEntityArgs> Entities
        {
            get => _entities ?? (_entities = new InputList<Inputs.GoogleCloudDialogflowV2beta1EntityTypeEntityArgs>());
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
        /// Required. The unique identifier of this session entity type. Supported formats: - `projects//agent/sessions//entityTypes/` - `projects//locations//agent/sessions//entityTypes/` - `projects//agent/environments//users//sessions//entityTypes/` - `projects//locations//agent/environments/ /users//sessions//entityTypes/` If `Location ID` is not specified we assume default 'us' location. If `Environment ID` is not specified, we assume default 'draft' environment. If `User ID` is not specified, we assume default '-' user. `` must be the display name of an existing entity type in the same agent that will be overridden or supplemented.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("project", required: true)]
        public Input<string> Project { get; set; } = null!;

        [Input("sessionId", required: true)]
        public Input<string> SessionId { get; set; } = null!;

        [Input("userId", required: true)]
        public Input<string> UserId { get; set; } = null!;

        public AgentEnvironmentUserSessionEntityTypeArgs()
        {
        }
    }
}
