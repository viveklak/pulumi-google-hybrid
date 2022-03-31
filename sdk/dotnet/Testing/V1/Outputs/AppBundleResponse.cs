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
    /// An Android App Bundle file format, containing a BundleConfig.pb file, a base module directory, zero or more dynamic feature module directories. See https://developer.android.com/guide/app-bundle/build for guidance on building App Bundles.
    /// </summary>
    [OutputType]
    public sealed class AppBundleResponse
    {
        /// <summary>
        /// .aab file representing the app bundle under test.
        /// </summary>
        public readonly Outputs.FileReferenceResponse BundleLocation;

        [OutputConstructor]
        private AppBundleResponse(Outputs.FileReferenceResponse bundleLocation)
        {
            BundleLocation = bundleLocation;
        }
    }
}
