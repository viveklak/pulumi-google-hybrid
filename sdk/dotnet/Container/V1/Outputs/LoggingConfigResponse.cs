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
    /// LoggingConfig is cluster logging configuration.
    /// </summary>
    [OutputType]
    public sealed class LoggingConfigResponse
    {
        /// <summary>
        /// Logging components configuration
        /// </summary>
        public readonly Outputs.LoggingComponentConfigResponse ComponentConfig;

        [OutputConstructor]
        private LoggingConfigResponse(Outputs.LoggingComponentConfigResponse componentConfig)
        {
            ComponentConfig = componentConfig;
        }
    }
}
