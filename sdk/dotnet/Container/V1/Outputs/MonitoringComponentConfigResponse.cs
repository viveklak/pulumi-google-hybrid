// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Container.V1.Outputs
{

    /// <summary>
    /// MonitoringComponentConfig is cluster monitoring component configuration.
    /// </summary>
    [OutputType]
    public sealed class MonitoringComponentConfigResponse
    {
        /// <summary>
        /// Select components to collect metrics. An empty set would disable all monitoring.
        /// </summary>
        public readonly ImmutableArray<string> EnableComponents;

        [OutputConstructor]
        private MonitoringComponentConfigResponse(ImmutableArray<string> enableComponents)
        {
            EnableComponents = enableComponents;
        }
    }
}
