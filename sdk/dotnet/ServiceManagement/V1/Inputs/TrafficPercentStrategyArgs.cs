// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.ServiceManagement.V1.Inputs
{

    /// <summary>
    /// Strategy that specifies how clients of Google Service Controller want to send traffic to use different config versions. This is generally used by API proxy to split traffic based on your configured percentage for each config version. One example of how to gradually rollout a new service configuration using this strategy: Day 1 Rollout { id: "example.googleapis.com/rollout_20160206" traffic_percent_strategy { percentages: { "example.googleapis.com/20160201": 70.00 "example.googleapis.com/20160206": 30.00 } } } Day 2 Rollout { id: "example.googleapis.com/rollout_20160207" traffic_percent_strategy: { percentages: { "example.googleapis.com/20160206": 100.00 } } }
    /// </summary>
    public sealed class TrafficPercentStrategyArgs : Pulumi.ResourceArgs
    {
        [Input("percentages")]
        private InputMap<string>? _percentages;

        /// <summary>
        /// Maps service configuration IDs to their corresponding traffic percentage. Key is the service configuration ID, Value is the traffic percentage which must be greater than 0.0 and the sum must equal to 100.0.
        /// </summary>
        public InputMap<string> Percentages
        {
            get => _percentages ?? (_percentages = new InputMap<string>());
            set => _percentages = value;
        }

        public TrafficPercentStrategyArgs()
        {
        }
    }
}
