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
    /// Binds the resources in a proxy or remote service with the GraphQL operation and its associated quota enforcement.
    /// </summary>
    public sealed class GoogleCloudApigeeV1GraphQLOperationConfigArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the API proxy endpoint or remote service with which the GraphQL operation and quota are associated.
        /// </summary>
        [Input("apiSource", required: true)]
        public Input<string> ApiSource { get; set; } = null!;

        [Input("attributes")]
        private InputList<Inputs.GoogleCloudApigeeV1AttributeArgs>? _attributes;

        /// <summary>
        /// Custom attributes associated with the operation.
        /// </summary>
        public InputList<Inputs.GoogleCloudApigeeV1AttributeArgs> Attributes
        {
            get => _attributes ?? (_attributes = new InputList<Inputs.GoogleCloudApigeeV1AttributeArgs>());
            set => _attributes = value;
        }

        [Input("operations", required: true)]
        private InputList<Inputs.GoogleCloudApigeeV1GraphQLOperationArgs>? _operations;

        /// <summary>
        /// List of GraphQL name/operation type pairs for the proxy or remote service to which quota will be applied. If only operation types are specified, the quota will be applied to all GraphQL requests irrespective of the GraphQL name. **Note**: Currently, you can specify only a single GraphQLOperation. Specifying more than one will cause the operation to fail.
        /// </summary>
        public InputList<Inputs.GoogleCloudApigeeV1GraphQLOperationArgs> Operations
        {
            get => _operations ?? (_operations = new InputList<Inputs.GoogleCloudApigeeV1GraphQLOperationArgs>());
            set => _operations = value;
        }

        /// <summary>
        /// Quota parameters to be enforced for the resources, methods, and API source combination. If none are specified, quota enforcement will not be done.
        /// </summary>
        [Input("quota")]
        public Input<Inputs.GoogleCloudApigeeV1QuotaArgs>? Quota { get; set; }

        public GoogleCloudApigeeV1GraphQLOperationConfigArgs()
        {
        }
    }
}
