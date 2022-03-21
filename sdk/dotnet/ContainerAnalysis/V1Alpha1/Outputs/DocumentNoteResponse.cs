// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.ContainerAnalysis.V1Alpha1.Outputs
{

    /// <summary>
    /// DocumentNote represents an SPDX Document Creation Infromation section: https://spdx.github.io/spdx-spec/2-document-creation-information/
    /// </summary>
    [OutputType]
    public sealed class DocumentNoteResponse
    {
        /// <summary>
        /// Compliance with the SPDX specification includes populating the SPDX fields therein with data related to such fields ("SPDX-Metadata")
        /// </summary>
        public readonly string DataLicence;
        /// <summary>
        /// Provide a reference number that can be used to understand how to parse and interpret the rest of the file
        /// </summary>
        public readonly string SpdxVersion;

        [OutputConstructor]
        private DocumentNoteResponse(
            string dataLicence,

            string spdxVersion)
        {
            DataLicence = dataLicence;
            SpdxVersion = spdxVersion;
        }
    }
}
