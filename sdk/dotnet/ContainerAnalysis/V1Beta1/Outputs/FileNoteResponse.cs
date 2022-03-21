// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.ContainerAnalysis.V1Beta1.Outputs
{

    /// <summary>
    /// FileNote represents an SPDX File Information section: https://spdx.github.io/spdx-spec/4-file-information/
    /// </summary>
    [OutputType]
    public sealed class FileNoteResponse
    {
        /// <summary>
        /// Provide a unique identifier to match analysis information on each specific file in a package
        /// </summary>
        public readonly ImmutableArray<string> Checksum;
        /// <summary>
        /// This field provides information about the type of file identified
        /// </summary>
        public readonly string FileType;
        /// <summary>
        /// Identify the full path and filename that corresponds to the file information in this section
        /// </summary>
        public readonly string Title;

        [OutputConstructor]
        private FileNoteResponse(
            ImmutableArray<string> checksum,

            string fileType,

            string title)
        {
            Checksum = checksum;
            FileType = fileType;
            Title = title;
        }
    }
}
