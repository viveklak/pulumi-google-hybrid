// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleCloud.Dlp.V2.Inputs
{

    /// <summary>
    /// Redact a given value. For example, if used with an `InfoTypeTransformation` transforming PHONE_NUMBER, and input 'My phone number is 206-555-0123', the output would be 'My phone number is '.
    /// </summary>
    public sealed class GooglePrivacyDlpV2RedactConfigArgs : Pulumi.ResourceArgs
    {
        public GooglePrivacyDlpV2RedactConfigArgs()
        {
        }
    }
}