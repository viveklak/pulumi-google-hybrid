// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Testing.V1.Outputs
{

    [OutputType]
    public sealed class AndroidDeviceListResponse
    {
        /// <summary>
        /// A list of Android devices.
        /// </summary>
        public readonly ImmutableArray<Outputs.AndroidDeviceResponse> AndroidDevices;

        [OutputConstructor]
        private AndroidDeviceListResponse(ImmutableArray<Outputs.AndroidDeviceResponse> androidDevices)
        {
            AndroidDevices = androidDevices;
        }
    }
}
