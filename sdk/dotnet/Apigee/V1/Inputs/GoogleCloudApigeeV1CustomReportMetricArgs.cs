// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Apigee.V1.Inputs
{

    /// <summary>
    /// This encapsulates a metric property of the form sum(message_count) where name is message_count and function is sum
    /// </summary>
    public sealed class GoogleCloudApigeeV1CustomReportMetricArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// aggregate function
        /// </summary>
        [Input("function")]
        public Input<string>? Function { get; set; }

        /// <summary>
        /// name of the metric
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public GoogleCloudApigeeV1CustomReportMetricArgs()
        {
        }
    }
}
