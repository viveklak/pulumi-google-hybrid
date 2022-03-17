// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.AppEngine.V1.Outputs
{

    /// <summary>
    /// The feature specific settings to be used in the application. These define behaviors that are user configurable.
    /// </summary>
    [OutputType]
    public sealed class FeatureSettingsResponse
    {
        /// <summary>
        /// Boolean value indicating if split health checks should be used instead of the legacy health checks. At an app.yaml level, this means defaulting to 'readiness_check' and 'liveness_check' values instead of 'health_check' ones. Once the legacy 'health_check' behavior is deprecated, and this value is always true, this setting can be removed.
        /// </summary>
        public readonly bool SplitHealthChecks;
        /// <summary>
        /// If true, use Container-Optimized OS (https://cloud.google.com/container-optimized-os/) base image for VMs, rather than a base Debian image.
        /// </summary>
        public readonly bool UseContainerOptimizedOs;

        [OutputConstructor]
        private FeatureSettingsResponse(
            bool splitHealthChecks,

            bool useContainerOptimizedOs)
        {
            SplitHealthChecks = splitHealthChecks;
            UseContainerOptimizedOs = useContainerOptimizedOs;
        }
    }
}
