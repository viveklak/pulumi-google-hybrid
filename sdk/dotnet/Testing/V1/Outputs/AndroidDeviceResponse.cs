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
    /// A single Android device.
    /// </summary>
    [OutputType]
    public sealed class AndroidDeviceResponse
    {
        /// <summary>
        /// The id of the Android device to be used. Use the TestEnvironmentDiscoveryService to get supported options.
        /// </summary>
        public readonly string AndroidModelId;
        /// <summary>
        /// The id of the Android OS version to be used. Use the TestEnvironmentDiscoveryService to get supported options.
        /// </summary>
        public readonly string AndroidVersionId;
        /// <summary>
        /// The locale the test device used for testing. Use the TestEnvironmentDiscoveryService to get supported options.
        /// </summary>
        public readonly string Locale;
        /// <summary>
        /// How the device is oriented during the test. Use the TestEnvironmentDiscoveryService to get supported options.
        /// </summary>
        public readonly string Orientation;

        [OutputConstructor]
        private AndroidDeviceResponse(
            string androidModelId,

            string androidVersionId,

            string locale,

            string orientation)
        {
            AndroidModelId = androidModelId;
            AndroidVersionId = androidVersionId;
            Locale = locale;
            Orientation = orientation;
        }
    }
}
