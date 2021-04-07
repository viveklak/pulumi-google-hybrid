// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GcpNative.Transcoder.V1Beta1.Outputs
{

    [OutputType]
    public sealed class EncryptionResponse
    {
        /// <summary>
        /// Configuration for AES-128 encryption.
        /// </summary>
        public readonly Outputs.Aes128EncryptionResponse Aes128;
        /// <summary>
        /// Required. 128 bit Initialization Vector (IV) represented as lowercase hexadecimal digits.
        /// </summary>
        public readonly string Iv;
        /// <summary>
        /// Required. 128 bit encryption key represented as lowercase hexadecimal digits.
        /// </summary>
        public readonly string Key;
        /// <summary>
        /// Configuration for MPEG Common Encryption (MPEG-CENC).
        /// </summary>
        public readonly Outputs.MpegCommonEncryptionResponse MpegCenc;
        /// <summary>
        /// Configuration for SAMPLE-AES encryption.
        /// </summary>
        public readonly Outputs.SampleAesEncryptionResponse SampleAes;

        [OutputConstructor]
        private EncryptionResponse(
            Outputs.Aes128EncryptionResponse aes128,

            string iv,

            string key,

            Outputs.MpegCommonEncryptionResponse mpegCenc,

            Outputs.SampleAesEncryptionResponse sampleAes)
        {
            Aes128 = aes128;
            Iv = iv;
            Key = key;
            MpegCenc = mpegCenc;
            SampleAes = sampleAes;
        }
    }
}