// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.APIGateway.V1Beta.Outputs
{

    /// <summary>
    /// Configuration settings for Gateways.
    /// </summary>
    [OutputType]
    public sealed class ApigatewayGatewayConfigResponse
    {
        /// <summary>
        /// Backend settings that are applied to all backends of the Gateway.
        /// </summary>
        public readonly Outputs.ApigatewayBackendConfigResponse BackendConfig;

        [OutputConstructor]
        private ApigatewayGatewayConfigResponse(Outputs.ApigatewayBackendConfigResponse backendConfig)
        {
            BackendConfig = backendConfig;
        }
    }
}
