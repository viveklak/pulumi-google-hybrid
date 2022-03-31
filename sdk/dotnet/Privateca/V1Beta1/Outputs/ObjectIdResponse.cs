// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Privateca.V1Beta1.Outputs
{

    /// <summary>
    /// An ObjectId specifies an object identifier (OID). These provide context and describe types in ASN.1 messages.
    /// </summary>
    [OutputType]
    public sealed class ObjectIdResponse
    {
        /// <summary>
        /// The parts of an OID path. The most significant parts of the path come first.
        /// </summary>
        public readonly ImmutableArray<int> ObjectIdPath;

        [OutputConstructor]
        private ObjectIdResponse(ImmutableArray<int> objectIdPath)
        {
            ObjectIdPath = objectIdPath;
        }
    }
}
