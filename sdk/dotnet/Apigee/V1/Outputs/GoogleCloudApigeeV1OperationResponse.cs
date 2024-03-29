// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Apigee.V1.Outputs
{

    /// <summary>
    /// Represents the pairing of REST resource path and the actions (verbs) allowed on the resource path.
    /// </summary>
    [OutputType]
    public sealed class GoogleCloudApigeeV1OperationResponse
    {
        /// <summary>
        /// methods refers to the REST verbs as in https://www.w3.org/Protocols/rfc2616/rfc2616-sec9.html. When none specified, all verb types are allowed.
        /// </summary>
        public readonly ImmutableArray<string> Methods;
        /// <summary>
        /// REST resource path associated with the API proxy or remote service.
        /// </summary>
        public readonly string Resource;

        [OutputConstructor]
        private GoogleCloudApigeeV1OperationResponse(
            ImmutableArray<string> methods,

            string resource)
        {
            Methods = methods;
            Resource = resource;
        }
    }
}
