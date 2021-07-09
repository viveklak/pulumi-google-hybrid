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
        /// GraphQL operation name, along with operation type which will be used to associate quotas with. If no name is specified, the quota will be applied to all graphQL operations irrespective of their operation names in the payload.
        /// </summary>
        public readonly string Operation;
        /// <summary>
        /// `query`, `mutation` and `subscription` are the three operation types offered by graphQL. Currently we support only `query` and `mutation`.
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
