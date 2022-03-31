// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.ToolResults.V1Beta3
{
    /// <summary>
    /// Creates an Execution. The returned Execution will have the id set. May return any of the following canonical error codes: - PERMISSION_DENIED - if the user is not authorized to write to project - INVALID_ARGUMENT - if the request is malformed - NOT_FOUND - if the containing History does not exist
    /// Auto-naming is currently not supported for this resource.
    /// Note - this resource's API doesn't support deletion. When deleted, the resource will persist
    /// on Google Cloud even though it will be deleted from Pulumi state.
    /// </summary>
    [GoogleNativeResourceType("google-native:toolresults/v1beta3:Execution")]
    public partial class Execution : Pulumi.CustomResource
    {
        /// <summary>
        /// The time when the Execution status transitioned to COMPLETE. This value will be set automatically when state transitions to COMPLETE. - In response: set if the execution state is COMPLETE. - In create/update request: never set
        /// </summary>
        [Output("completionTime")]
        public Output<Outputs.TimestampResponse> CompletionTime { get; private set; } = null!;

        /// <summary>
        /// The time when the Execution was created. This value will be set automatically when CreateExecution is called. - In response: always set - In create/update request: never set
        /// </summary>
        [Output("creationTime")]
        public Output<Outputs.TimestampResponse> CreationTime { get; private set; } = null!;

        /// <summary>
        /// The dimensions along which different steps in this execution may vary. This must remain fixed over the life of the execution. Returns INVALID_ARGUMENT if this field is set in an update request. Returns INVALID_ARGUMENT if the same name occurs in more than one dimension_definition. Returns INVALID_ARGUMENT if the size of the list is over 100. - In response: present if set by create - In create request: optional - In update request: never set
        /// </summary>
        [Output("dimensionDefinitions")]
        public Output<ImmutableArray<Outputs.MatrixDimensionDefinitionResponse>> DimensionDefinitions { get; private set; } = null!;

        /// <summary>
        /// A unique identifier within a History for this Execution. Returns INVALID_ARGUMENT if this field is set or overwritten by the caller. - In response always set - In create/update request: never set
        /// </summary>
        [Output("executionId")]
        public Output<string> ExecutionId { get; private set; } = null!;

        /// <summary>
        /// Classify the result, for example into SUCCESS or FAILURE - In response: present if set by create/update request - In create/update request: optional
        /// </summary>
        [Output("outcome")]
        public Output<Outputs.OutcomeResponse> Outcome { get; private set; } = null!;

        /// <summary>
        /// Lightweight information about execution request. - In response: present if set by create - In create: optional - In update: optional
        /// </summary>
        [Output("specification")]
        public Output<Outputs.SpecificationResponse> Specification { get; private set; } = null!;

        /// <summary>
        /// The initial state is IN_PROGRESS. The only legal state transitions is from IN_PROGRESS to COMPLETE. A PRECONDITION_FAILED will be returned if an invalid transition is requested. The state can only be set to COMPLETE once. A FAILED_PRECONDITION will be returned if the state is set to COMPLETE multiple times. If the state is set to COMPLETE, all the in-progress steps within the execution will be set as COMPLETE. If the outcome of the step is not set, the outcome will be set to INCONCLUSIVE. - In response always set - In create/update request: optional
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;

        /// <summary>
        /// TestExecution Matrix ID that the TestExecutionService uses. - In response: present if set by create - In create: optional - In update: never set
        /// </summary>
        [Output("testExecutionMatrixId")]
        public Output<string> TestExecutionMatrixId { get; private set; } = null!;


        /// <summary>
        /// Create a Execution resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Execution(string name, ExecutionArgs args, CustomResourceOptions? options = null)
            : base("google-native:toolresults/v1beta3:Execution", name, args ?? new ExecutionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Execution(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:toolresults/v1beta3:Execution", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing Execution resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Execution Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Execution(name, id, options);
        }
    }

    public sealed class ExecutionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The time when the Execution status transitioned to COMPLETE. This value will be set automatically when state transitions to COMPLETE. - In response: set if the execution state is COMPLETE. - In create/update request: never set
        /// </summary>
        [Input("completionTime")]
        public Input<Inputs.TimestampArgs>? CompletionTime { get; set; }

        /// <summary>
        /// The time when the Execution was created. This value will be set automatically when CreateExecution is called. - In response: always set - In create/update request: never set
        /// </summary>
        [Input("creationTime")]
        public Input<Inputs.TimestampArgs>? CreationTime { get; set; }

        [Input("dimensionDefinitions")]
        private InputList<Inputs.MatrixDimensionDefinitionArgs>? _dimensionDefinitions;

        /// <summary>
        /// The dimensions along which different steps in this execution may vary. This must remain fixed over the life of the execution. Returns INVALID_ARGUMENT if this field is set in an update request. Returns INVALID_ARGUMENT if the same name occurs in more than one dimension_definition. Returns INVALID_ARGUMENT if the size of the list is over 100. - In response: present if set by create - In create request: optional - In update request: never set
        /// </summary>
        public InputList<Inputs.MatrixDimensionDefinitionArgs> DimensionDefinitions
        {
            get => _dimensionDefinitions ?? (_dimensionDefinitions = new InputList<Inputs.MatrixDimensionDefinitionArgs>());
            set => _dimensionDefinitions = value;
        }

        /// <summary>
        /// A unique identifier within a History for this Execution. Returns INVALID_ARGUMENT if this field is set or overwritten by the caller. - In response always set - In create/update request: never set
        /// </summary>
        [Input("executionId")]
        public Input<string>? ExecutionId { get; set; }

        [Input("historyId", required: true)]
        public Input<string> HistoryId { get; set; } = null!;

        /// <summary>
        /// Classify the result, for example into SUCCESS or FAILURE - In response: present if set by create/update request - In create/update request: optional
        /// </summary>
        [Input("outcome")]
        public Input<Inputs.OutcomeArgs>? Outcome { get; set; }

        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// A unique request ID for server to detect duplicated requests. For example, a UUID. Optional, but strongly recommended.
        /// </summary>
        [Input("requestId")]
        public Input<string>? RequestId { get; set; }

        /// <summary>
        /// Lightweight information about execution request. - In response: present if set by create - In create: optional - In update: optional
        /// </summary>
        [Input("specification")]
        public Input<Inputs.SpecificationArgs>? Specification { get; set; }

        /// <summary>
        /// The initial state is IN_PROGRESS. The only legal state transitions is from IN_PROGRESS to COMPLETE. A PRECONDITION_FAILED will be returned if an invalid transition is requested. The state can only be set to COMPLETE once. A FAILED_PRECONDITION will be returned if the state is set to COMPLETE multiple times. If the state is set to COMPLETE, all the in-progress steps within the execution will be set as COMPLETE. If the outcome of the step is not set, the outcome will be set to INCONCLUSIVE. - In response always set - In create/update request: optional
        /// </summary>
        [Input("state")]
        public Input<Pulumi.GoogleNative.ToolResults.V1Beta3.ExecutionState>? State { get; set; }

        /// <summary>
        /// TestExecution Matrix ID that the TestExecutionService uses. - In response: present if set by create - In create: optional - In update: never set
        /// </summary>
        [Input("testExecutionMatrixId")]
        public Input<string>? TestExecutionMatrixId { get; set; }

        public ExecutionArgs()
        {
        }
    }
}
