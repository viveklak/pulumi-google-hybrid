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
    /// List of graphQL operation configuration details associated with Apigee API proxies or remote services. Remote services are non-Apigee proxies, such as Istio-Envoy.
    /// </summary>
    public sealed class GoogleCloudApigeeV1GraphQLOperationGroupArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Flag that specifies whether the configuration is for Apigee API proxy or a remote service. Valid values include `proxy` or `remoteservice`. Defaults to `proxy`. Set to `proxy` when Apigee API proxies are associated with the API product. Set to `remoteservice` when non-Apigee proxies like Istio-Envoy are associated with the API product.
        /// </summary>
        [Input("operationConfigType")]
        public Input<string>? OperationConfigType { get; set; }

        [Input("operationConfigs", required: true)]
        private InputList<Inputs.GoogleCloudApigeeV1GraphQLOperationConfigArgs>? _operationConfigs;

        /// <summary>
        /// List of operation configurations for either Apigee API proxies or other remote services that are associated with this API product.
        /// </summary>
        public InputList<Inputs.GoogleCloudApigeeV1GraphQLOperationConfigArgs> OperationConfigs
        {
            get => _operationConfigs ?? (_operationConfigs = new InputList<Inputs.GoogleCloudApigeeV1GraphQLOperationConfigArgs>());
            set => _operationConfigs = value;
        }

        public GoogleCloudApigeeV1GraphQLOperationGroupArgs()
        {
        }
    }
}
