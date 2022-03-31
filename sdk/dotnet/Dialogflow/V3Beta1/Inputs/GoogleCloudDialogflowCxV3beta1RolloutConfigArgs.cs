// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dialogflow.V3Beta1.Inputs
{

    /// <summary>
    /// The configuration for auto rollout.
    /// </summary>
    public sealed class GoogleCloudDialogflowCxV3beta1RolloutConfigArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The conditions that are used to evaluate the failure of a rollout step. If not specified, no rollout steps will fail. E.g. "containment_rate &lt; 10% OR average_turn_count &lt; 3". See the [conditions reference](https://cloud.google.com/dialogflow/cx/docs/reference/condition).
        /// </summary>
        [Input("failureCondition")]
        public Input<string>? FailureCondition { get; set; }

        /// <summary>
        /// The conditions that are used to evaluate the success of a rollout step. If not specified, all rollout steps will proceed to the next one unless failure conditions are met. E.g. "containment_rate &gt; 60% AND callback_rate &lt; 20%". See the [conditions reference](https://cloud.google.com/dialogflow/cx/docs/reference/condition).
        /// </summary>
        [Input("rolloutCondition")]
        public Input<string>? RolloutCondition { get; set; }

        [Input("rolloutSteps")]
        private InputList<Inputs.GoogleCloudDialogflowCxV3beta1RolloutConfigRolloutStepArgs>? _rolloutSteps;

        /// <summary>
        /// Steps to roll out a flow version. Steps should be sorted by percentage in ascending order.
        /// </summary>
        public InputList<Inputs.GoogleCloudDialogflowCxV3beta1RolloutConfigRolloutStepArgs> RolloutSteps
        {
            get => _rolloutSteps ?? (_rolloutSteps = new InputList<Inputs.GoogleCloudDialogflowCxV3beta1RolloutConfigRolloutStepArgs>());
            set => _rolloutSteps = value;
        }

        public GoogleCloudDialogflowCxV3beta1RolloutConfigArgs()
        {
        }
    }
}
