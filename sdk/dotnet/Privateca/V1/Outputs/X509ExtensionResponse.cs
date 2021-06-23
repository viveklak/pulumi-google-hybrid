// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Privateca.V1.Outputs
{

    [OutputType]
    public sealed class X509ExtensionResponse
    {
        /// <summary>
        /// Optional. Indicates whether or not this extension is critical (i.e., if the client does not know how to handle this extension, the client should consider this to be an error).
        /// </summary>
        public readonly bool Critical;
        /// <summary>
        /// Required. The OID for this X.509 extension.
        /// </summary>
        public readonly Outputs.ObjectIdResponse ObjectId;
        /// <summary>
        /// Required. The value of this X.509 extension.
        /// </summary>
        public readonly string Value;

        [OutputConstructor]
        private X509ExtensionResponse(
            bool critical,

            Outputs.ObjectIdResponse objectId,

            string value)
        {
            Critical = critical;
            ObjectId = objectId;
            Value = value;
        }
    }
}
