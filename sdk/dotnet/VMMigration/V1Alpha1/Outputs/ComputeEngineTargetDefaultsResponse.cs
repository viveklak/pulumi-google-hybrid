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
    /// ComputeEngineTargetDefaults is a collection of details for creating a VM in a target Compute Engine project.
    /// </summary>
    [OutputType]
    public sealed class ComputeEngineTargetDefaultsResponse
    {
        /// <summary>
        /// Additional licenses to assign to the VM.
        /// </summary>
        public readonly ImmutableArray<string> AdditionalLicenses;
        /// <summary>
        /// The OS license returned from the adaptation module report.
        /// </summary>
        public readonly Outputs.AppliedLicenseResponse AppliedLicense;
        /// <summary>
        /// The VM Boot Option, as set in the source vm.
        /// </summary>
        public readonly string BootOption;
        /// <summary>
        /// Compute instance scheduling information (if empty default is used).
        /// </summary>
        public readonly Outputs.ComputeSchedulingResponse ComputeScheduling;
        /// <summary>
        /// The disk type to use in the VM.
        /// </summary>
        public readonly string DiskType;
        /// <summary>
        /// A map of labels to associate with the VM.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Labels;
        /// <summary>
        /// The license type to use in OS adaptation.
        /// </summary>
        public readonly string LicenseType;
        /// <summary>
        /// The machine type to create the VM with.
        /// </summary>
        public readonly string MachineType;
        /// <summary>
        /// The machine type series to create the VM with.
        /// </summary>
        public readonly string MachineTypeSeries;
        /// <summary>
        /// The metadata key/value pairs to assign to the VM.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Metadata;
        /// <summary>
        /// List of NICs connected to this VM.
        /// </summary>
        public readonly ImmutableArray<Outputs.NetworkInterfaceResponse> NetworkInterfaces;
        /// <summary>
        /// A map of network tags to associate with the VM.
        /// </summary>
        public readonly ImmutableArray<string> NetworkTags;
        /// <summary>
        /// Defines whether the instance has Secure Boot enabled. This can be set to true only if the vm boot option is EFI.
        /// </summary>
        public readonly bool SecureBoot;
        /// <summary>
        /// The service account to associate the VM with.
        /// </summary>
        public readonly string ServiceAccount;
        /// <summary>
        /// The full path of the resource of type TargetProject which represents the Compute Engine project in which to create this VM.
        /// </summary>
        public readonly string TargetProject;
        /// <summary>
        /// The name of the VM to create.
        /// </summary>
        public readonly string VmName;
        /// <summary>
        /// The zone in which to create the VM.
        /// </summary>
        public readonly string Zone;

        [OutputConstructor]
        private ComputeEngineTargetDefaultsResponse(
            ImmutableArray<string> additionalLicenses,

            Outputs.AppliedLicenseResponse appliedLicense,

            string bootOption,

            Outputs.ComputeSchedulingResponse computeScheduling,

            string diskType,

            ImmutableDictionary<string, string> labels,

            string licenseType,

            string machineType,

            string machineTypeSeries,

            ImmutableDictionary<string, string> metadata,

            ImmutableArray<Outputs.NetworkInterfaceResponse> networkInterfaces,

            ImmutableArray<string> networkTags,

            bool secureBoot,

            string serviceAccount,

            string targetProject,

            string vmName,

            string zone)
        {
            AdditionalLicenses = additionalLicenses;
            AppliedLicense = appliedLicense;
            BootOption = bootOption;
            ComputeScheduling = computeScheduling;
            DiskType = diskType;
            Labels = labels;
            LicenseType = licenseType;
            MachineType = machineType;
            MachineTypeSeries = machineTypeSeries;
            Metadata = metadata;
            NetworkInterfaces = networkInterfaces;
            NetworkTags = networkTags;
            SecureBoot = secureBoot;
            ServiceAccount = serviceAccount;
            TargetProject = targetProject;
            VmName = vmName;
            Zone = zone;
        }
    }
}
