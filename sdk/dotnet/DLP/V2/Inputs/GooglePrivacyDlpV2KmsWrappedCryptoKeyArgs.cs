// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.DLP.V2.Inputs
{

    /// <summary>
    /// Include to use an existing data crypto key wrapped by KMS. The wrapped key must be a 128-, 192-, or 256-bit key. Authorization requires the following IAM permissions when sending a request to perform a crypto transformation using a KMS-wrapped crypto key: dlp.kms.encrypt For more information, see [Creating a wrapped key] (https://cloud.google.com/dlp/docs/create-wrapped-key).
    /// </summary>
    public sealed class GooglePrivacyDlpV2KmsWrappedCryptoKeyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Required. The resource name of the KMS CryptoKey to use for unwrapping.
        /// </summary>
        [Input("cryptoKeyName")]
        public Input<string>? CryptoKeyName { get; set; }

        /// <summary>
        /// Required. The wrapped data crypto key.
        /// </summary>
        [Input("wrappedKey")]
        public Input<string>? WrappedKey { get; set; }

        public GooglePrivacyDlpV2KmsWrappedCryptoKeyArgs()
        {
        }
    }
}
