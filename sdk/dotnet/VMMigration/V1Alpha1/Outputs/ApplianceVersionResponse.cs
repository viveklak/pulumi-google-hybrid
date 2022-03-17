// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.VMMigration.V1Alpha1.Outputs
{

    /// <summary>
    /// Describes an appliance version.
    /// </summary>
    [OutputType]
    public sealed class ApplianceVersionResponse
    {
        /// <summary>
        /// Determine whether it's critical to upgrade the appliance to this version.
        /// </summary>
        public readonly bool Critical;
        /// <summary>
        /// Link to a page that contains the version release notes.
        /// </summary>
        public readonly string ReleaseNotesUri;
        /// <summary>
        /// A link for downloading the version.
        /// </summary>
        public readonly string Uri;
        /// <summary>
        /// The appliance version.
        /// </summary>
        public readonly string Version;

        [OutputConstructor]
        private ApplianceVersionResponse(
            bool critical,

            string releaseNotesUri,

            string uri,

            string version)
        {
            Critical = critical;
            ReleaseNotesUri = releaseNotesUri;
            Uri = uri;
            Version = version;
        }
    }
}
