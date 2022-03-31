// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.GameServices.V1.Inputs
{

    /// <summary>
    /// Autoscaling config for an Agones fleet.
    /// </summary>
    public sealed class ScalingConfigArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Agones fleet autoscaler spec (see [example spec](https://agones.dev/site/docs/reference/fleetautoscaler/)).
        /// </summary>
        [Input("fleetAutoscalerSpec", required: true)]
        public Input<string> FleetAutoscalerSpec { get; set; } = null!;

        /// <summary>
        /// The name of the scaling config.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("schedules")]
        private InputList<Inputs.ScheduleArgs>? _schedules;

        /// <summary>
        /// The schedules to which this scaling config applies.
        /// </summary>
        public InputList<Inputs.ScheduleArgs> Schedules
        {
            get => _schedules ?? (_schedules = new InputList<Inputs.ScheduleArgs>());
            set => _schedules = value;
        }

        [Input("selectors")]
        private InputList<Inputs.LabelSelectorArgs>? _selectors;

        /// <summary>
        /// Labels used to identify the game server clusters to which this Agones scaling config applies. A game server cluster is subject to this Agones scaling config if its labels match any of the selector entries.
        /// </summary>
        public InputList<Inputs.LabelSelectorArgs> Selectors
        {
            get => _selectors ?? (_selectors = new InputList<Inputs.LabelSelectorArgs>());
            set => _selectors = value;
        }

        public ScalingConfigArgs()
        {
        }
    }
}
