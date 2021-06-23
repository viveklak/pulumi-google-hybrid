// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Alpha.Outputs
{

    [OutputType]
    public sealed class CallCredentialsResponse
    {
        /// <summary>
        /// The type of call credentials to use for GRPC requests to the SDS server. This field can be set to one of the following: - GCE_VM: The local GCE VM service account credentials are used to access the SDS server. - FROM_PLUGIN: Custom authenticator credentials are used to access the SDS server.
        /// </summary>
        public readonly string CallCredentialType;
        /// <summary>
        /// Custom authenticator credentials. Valid if callCredentialType is FROM_PLUGIN.
        /// </summary>
        public readonly Outputs.MetadataCredentialsFromPluginResponse FromPlugin;

        [OutputConstructor]
        private CallCredentialsResponse(
            string callCredentialType,

            Outputs.MetadataCredentialsFromPluginResponse fromPlugin)
        {
            CallCredentialType = callCredentialType;
            FromPlugin = fromPlugin;
        }
    }
}
