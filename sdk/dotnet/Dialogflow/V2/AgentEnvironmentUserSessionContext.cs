// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dialogflow.V2
{
    /// <summary>
    /// Creates a context. If the specified context already exists, overrides the context.
    /// </summary>
    [GoogleNativeResourceType("google-native:dialogflow/v2:AgentEnvironmentUserSessionContext")]
    public partial class AgentEnvironmentUserSessionContext : Pulumi.CustomResource
    {
        /// <summary>
        /// Optional. The number of conversational query requests after which the context expires. The default is `0`. If set to `0`, the context expires immediately. Contexts expire automatically after 20 minutes if there are no matching queries.
        /// </summary>
        [Output("lifespanCount")]
        public Output<int> LifespanCount { get; private set; } = null!;

        /// <summary>
        /// Required. The unique identifier of the context. Format: `projects//agent/sessions//contexts/`, or `projects//agent/environments//users//sessions//contexts/`. The `Context ID` is always converted to lowercase, may only contain characters in a-zA-Z0-9_-% and may be at most 250 bytes long. If `Environment ID` is not specified, we assume default 'draft' environment. If `User ID` is not specified, we assume default '-' user. The following context names are reserved for internal use by Dialogflow. You should not use these contexts or create contexts with these names: * `__system_counters__` * `*_id_dialog_context` * `*_dialog_params_size`
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Optional. The collection of parameters associated with this context. Depending on your protocol or client library language, this is a map, associative array, symbol table, dictionary, or JSON object composed of a collection of (MapKey, MapValue) pairs: - MapKey type: string - MapKey value: parameter name - MapValue type: - If parameter's entity type is a composite entity: map - Else: depending on parameter value type, could be one of string, number, boolean, null, list or map - MapValue value: - If parameter's entity type is a composite entity: map from composite entity property names to property values - Else: parameter value
        /// </summary>
        [Output("parameters")]
        public Output<ImmutableDictionary<string, string>> Parameters { get; private set; } = null!;


        /// <summary>
        /// Create a AgentEnvironmentUserSessionContext resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AgentEnvironmentUserSessionContext(string name, AgentEnvironmentUserSessionContextArgs args, CustomResourceOptions? options = null)
            : base("google-native:dialogflow/v2:AgentEnvironmentUserSessionContext", name, args ?? new AgentEnvironmentUserSessionContextArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AgentEnvironmentUserSessionContext(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:dialogflow/v2:AgentEnvironmentUserSessionContext", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing AgentEnvironmentUserSessionContext resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AgentEnvironmentUserSessionContext Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new AgentEnvironmentUserSessionContext(name, id, options);
        }
    }

    public sealed class AgentEnvironmentUserSessionContextArgs : Pulumi.ResourceArgs
    {
        [Input("environmentId", required: true)]
        public Input<string> EnvironmentId { get; set; } = null!;

        /// <summary>
        /// Optional. The number of conversational query requests after which the context expires. The default is `0`. If set to `0`, the context expires immediately. Contexts expire automatically after 20 minutes if there are no matching queries.
        /// </summary>
        [Input("lifespanCount")]
        public Input<int>? LifespanCount { get; set; }

        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        /// <summary>
        /// Required. The unique identifier of the context. Format: `projects//agent/sessions//contexts/`, or `projects//agent/environments//users//sessions//contexts/`. The `Context ID` is always converted to lowercase, may only contain characters in a-zA-Z0-9_-% and may be at most 250 bytes long. If `Environment ID` is not specified, we assume default 'draft' environment. If `User ID` is not specified, we assume default '-' user. The following context names are reserved for internal use by Dialogflow. You should not use these contexts or create contexts with these names: * `__system_counters__` * `*_id_dialog_context` * `*_dialog_params_size`
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("parameters")]
        private InputMap<string>? _parameters;

        /// <summary>
        /// Optional. The collection of parameters associated with this context. Depending on your protocol or client library language, this is a map, associative array, symbol table, dictionary, or JSON object composed of a collection of (MapKey, MapValue) pairs: - MapKey type: string - MapKey value: parameter name - MapValue type: - If parameter's entity type is a composite entity: map - Else: depending on parameter value type, could be one of string, number, boolean, null, list or map - MapValue value: - If parameter's entity type is a composite entity: map from composite entity property names to property values - Else: parameter value
        /// </summary>
        public InputMap<string> Parameters
        {
            get => _parameters ?? (_parameters = new InputMap<string>());
            set => _parameters = value;
        }

        [Input("project", required: true)]
        public Input<string> Project { get; set; } = null!;

        [Input("sessionId", required: true)]
        public Input<string> SessionId { get; set; } = null!;

        [Input("userId", required: true)]
        public Input<string> UserId { get; set; } = null!;

        public AgentEnvironmentUserSessionContextArgs()
        {
        }
    }
}
