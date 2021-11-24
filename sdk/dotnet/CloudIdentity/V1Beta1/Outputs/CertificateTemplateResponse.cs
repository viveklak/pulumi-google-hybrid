// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.CloudIdentity.V1Beta1.Outputs
{

    /// <summary>
    /// CertificateTemplate (v3 Extension in X.509).
    /// </summary>
    [OutputType]
    public sealed class CertificateTemplateResponse
    {
        /// <summary>
        /// The Major version of the template. Example: 100.
        /// </summary>
        public readonly int MajorVersion;
        /// <summary>
        /// The minor version of the template. Example: 12.
        /// </summary>
        public readonly int MinorVersion;

        [OutputConstructor]
        private CertificateTemplateResponse(
            int majorVersion,

            int minorVersion)
        {
            MajorVersion = majorVersion;
            MinorVersion = minorVersion;
        }
    }
}