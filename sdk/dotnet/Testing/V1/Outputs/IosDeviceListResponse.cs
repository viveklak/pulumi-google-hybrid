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
    /// A list of iOS device configurations in which the test is to be executed.
    /// </summary>
    [OutputType]
    public sealed class IosDeviceListResponse
    {
        /// <summary>
        /// A list of iOS devices.
        /// </summary>
        public readonly ImmutableArray<Outputs.IosDeviceResponse> IosDevices;

        [OutputConstructor]
        private IosDeviceListResponse(ImmutableArray<Outputs.IosDeviceResponse> iosDevices)
        {
            IosDevices = iosDevices;
        }
    }
}
