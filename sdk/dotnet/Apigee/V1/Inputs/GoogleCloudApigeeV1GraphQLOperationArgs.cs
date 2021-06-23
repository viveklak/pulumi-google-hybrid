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
    /// Represents the pairing of GraphQL operation types and the GraphQL operation name.
    /// </summary>
    public sealed class GoogleCloudApigeeV1GraphQLOperationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// GraphQL operation name. The name and operation type will be used to apply quotas. If no name is specified, the quota will be applied to all GraphQL operations irrespective of their operation names in the payload.
        /// </summary>
        [Input("operation")]
        public Input<string>? Operation { get; set; }

        [Input("operationTypes")]
        private InputList<string>? _operationTypes;

        /// <summary>
        /// Required. GraphQL operation types. Valid values include `query` or `mutation`. **Note**: Apigee does not currently support `subscription` types.
        /// </summary>
        public InputList<string> OperationTypes
        {
            get => _operationTypes ?? (_operationTypes = new InputList<string>());
            set => _operationTypes = value;
        }

        public GoogleCloudApigeeV1GraphQLOperationArgs()
        {
        }
    }
}
