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
    public sealed class GoogleCloudApigeeV1GraphQLOperationResponse
    {
        /// <summary>
        /// GraphQL operation name. The name and operation type will be used to apply quotas. If no name is specified, the quota will be applied to all GraphQL operations irrespective of their operation names in the payload.
        /// </summary>
        public readonly string Operation;
        /// <summary>
        /// Required. GraphQL operation types. Valid values include `query` or `mutation`. **Note**: Apigee does not currently support `subscription` types.
        /// </summary>
        public readonly ImmutableArray<string> OperationTypes;

        [OutputConstructor]
        private GoogleCloudApigeeV1GraphQLOperationResponse(
            string operation,

            ImmutableArray<string> operationTypes)
        {
            Operation = operation;
            OperationTypes = operationTypes;
        }
    }
}
