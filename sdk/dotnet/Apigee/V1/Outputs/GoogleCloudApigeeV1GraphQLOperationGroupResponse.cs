// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Apigee.V1.Outputs
{

    [OutputType]
    public sealed class GoogleCloudApigeeV1GraphQLOperationGroupResponse
    {
        /// <summary>
        /// Flag that specifes whether the configuration is for Apigee API proxy or a remote service. Valid values are `proxy` or `remoteservice`. Defaults to `proxy`. Set to `proxy` when Apigee API proxies are associated with the API product. Set to `remoteservice` when non-Apigee proxies like Istio-Envoy are associated with the API product.
        /// </summary>
        public readonly string OperationConfigType;
        /// <summary>
        /// List of operation configurations for either Apigee API proxies or other remote services that are associated with this API product.
        /// </summary>
        public readonly ImmutableArray<Outputs.GoogleCloudApigeeV1GraphQLOperationConfigResponse> OperationConfigs;

        [OutputConstructor]
        private GoogleCloudApigeeV1GraphQLOperationGroupResponse(
            string operationConfigType,

            ImmutableArray<Outputs.GoogleCloudApigeeV1GraphQLOperationConfigResponse> operationConfigs)
        {
            OperationConfigType = operationConfigType;
            OperationConfigs = operationConfigs;
        }
    }
}
