// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Cloudkms.V1.Inputs
{

    /// <summary>
    /// ExternalProtectionLevelOptions stores a group of additional fields for configuring a CryptoKeyVersion that are specific to the EXTERNAL protection level and EXTERNAL_VPC protection levels.
    /// </summary>
    public sealed class ExternalProtectionLevelOptionsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The path to the external key material on the EKM when using EkmConnection e.g., "v0/my/key". Set this field instead of external_key_uri when using an EkmConnection.
        /// </summary>
        [Input("ekmConnectionKeyPath")]
        public Input<string>? EkmConnectionKeyPath { get; set; }

        /// <summary>
        /// The URI for an external resource that this CryptoKeyVersion represents.
        /// </summary>
        [Input("externalKeyUri")]
        public Input<string>? ExternalKeyUri { get; set; }

        public ExternalProtectionLevelOptionsArgs()
        {
        }
    }
}
