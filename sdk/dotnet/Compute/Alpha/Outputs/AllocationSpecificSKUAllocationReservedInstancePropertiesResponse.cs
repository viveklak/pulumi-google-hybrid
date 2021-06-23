// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Alpha.Outputs
{

    [OutputType]
    public sealed class AllocationSpecificSKUAllocationReservedInstancePropertiesResponse
    {
        /// <summary>
        /// Specifies accelerator type and count.
        /// </summary>
        public readonly ImmutableArray<Outputs.AcceleratorConfigResponse> GuestAccelerators;
        /// <summary>
        /// Specifies amount of local ssd to reserve with each instance. The type of disk is local-ssd.
        /// </summary>
        public readonly ImmutableArray<Outputs.AllocationSpecificSKUAllocationAllocatedInstancePropertiesReservedDiskResponse> LocalSsds;
        /// <summary>
        /// An opaque location hint used to place the allocation close to other resources. This field is for use by internal tools that use the public API.
        /// </summary>
        public readonly string LocationHint;
        /// <summary>
        /// Specifies type of machine (name only) which has fixed number of vCPUs and fixed amount of memory. This also includes specifying custom machine type following custom-NUMBER_OF_CPUS-AMOUNT_OF_MEMORY pattern.
        /// </summary>
        public readonly string MachineType;
        /// <summary>
        /// Specifies the number of hours after reservation creation where instances using the reservation won't be scheduled for maintenance.
        /// </summary>
        public readonly int MaintenanceFreezeDurationHours;
        /// <summary>
        /// For more information about maintenance intervals, see Setting maintenance intervals.
        /// </summary>
        public readonly string MaintenanceInterval;
        /// <summary>
        /// Minimum cpu platform the reservation.
        /// </summary>
        public readonly string MinCpuPlatform;

        [OutputConstructor]
        private AllocationSpecificSKUAllocationReservedInstancePropertiesResponse(
            ImmutableArray<Outputs.AcceleratorConfigResponse> guestAccelerators,

            ImmutableArray<Outputs.AllocationSpecificSKUAllocationAllocatedInstancePropertiesReservedDiskResponse> localSsds,

            string locationHint,

            string machineType,

            int maintenanceFreezeDurationHours,

            string maintenanceInterval,

            string minCpuPlatform)
        {
            GuestAccelerators = guestAccelerators;
            LocalSsds = localSsds;
            LocationHint = locationHint;
            MachineType = machineType;
            MaintenanceFreezeDurationHours = maintenanceFreezeDurationHours;
            MaintenanceInterval = maintenanceInterval;
            MinCpuPlatform = minCpuPlatform;
        }
    }
}
