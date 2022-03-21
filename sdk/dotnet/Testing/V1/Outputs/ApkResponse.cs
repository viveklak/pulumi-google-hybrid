// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Testing.V1.Outputs
{

    /// <summary>
    /// An Android package file to install.
    /// </summary>
    [OutputType]
    public sealed class ApkResponse
    {
        /// <summary>
        /// The path to an APK to be installed on the device before the test begins.
        /// </summary>
        public readonly Outputs.FileReferenceResponse Location;
        /// <summary>
        /// The java package for the APK to be installed. Value is determined by examining the application's manifest.
        /// </summary>
        public readonly string PackageName;

        [OutputConstructor]
        private ApkResponse(
            Outputs.FileReferenceResponse location,

            string packageName)
        {
            Location = location;
            PackageName = packageName;
        }
    }
}
