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
    /// Configuration of a specific billing destination (Currently only support bill against consumer project).
    /// </summary>
    public sealed class BillingDestinationArgs : Pulumi.ResourceArgs
    {
        [Input("metrics")]
        private InputList<string>? _metrics;

        /// <summary>
        /// Names of the metrics to report to this billing destination. Each name must be defined in Service.metrics section.
        /// </summary>
        public InputList<string> Metrics
        {
            get => _metrics ?? (_metrics = new InputList<string>());
            set => _metrics = value;
        }

        /// <summary>
        /// The monitored resource type. The type must be defined in Service.monitored_resources section.
        /// </summary>
        [Input("monitoredResource")]
        public Input<string>? MonitoredResource { get; set; }

        public BillingDestinationArgs()
        {
        }
    }
}
