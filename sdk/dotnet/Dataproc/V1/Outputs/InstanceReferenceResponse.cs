// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dataproc.V1.Outputs
{

    [OutputType]
    public sealed class InstanceReferenceResponse
    {
        /// <summary>
        /// The unique identifier of the Compute Engine instance.
        /// </summary>
        public readonly string InstanceId;
        /// <summary>
        /// The user-friendly name of the Compute Engine instance.
        /// </summary>
        public readonly string InstanceName;
        /// <summary>
        /// The public ECIES key used for sharing data with this instance.
        /// </summary>
        public readonly string PublicEciesKey;
        /// <summary>
        /// The public RSA key used for sharing data with this instance.
        /// </summary>
        public readonly string PublicKey;

        [OutputConstructor]
        private InstanceReferenceResponse(
            string instanceId,

            string instanceName,

            string publicEciesKey,

            string publicKey)
        {
            InstanceId = instanceId;
            InstanceName = instanceName;
            PublicEciesKey = publicEciesKey;
            PublicKey = publicKey;
        }
    }
}
