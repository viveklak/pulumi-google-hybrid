// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dataproc.V1.Outputs
{

    /// <summary>
    /// The config settings for Compute Engine resources in an instance group, such as a master or worker group.
    /// </summary>
    [OutputType]
    public sealed class InstanceGroupConfigResponse
    {
        /// <summary>
        /// Optional. The Compute Engine accelerator configuration for these instances.
        /// </summary>
        public readonly ImmutableArray<Outputs.AcceleratorConfigResponse> Accelerators;
        /// <summary>
        /// Optional. Disk option config settings.
        /// </summary>
        public readonly Outputs.DiskConfigResponse DiskConfig;
        /// <summary>
        /// Optional. The Compute Engine image resource used for cluster instances.The URI can represent an image or image family.Image examples: https://www.googleapis.com/compute/beta/projects/[project_id]/global/images/[image-id] projects/[project_id]/global/images/[image-id] image-idImage family examples. Dataproc will use the most recent image from the family: https://www.googleapis.com/compute/beta/projects/[project_id]/global/images/family/[custom-image-family-name] projects/[project_id]/global/images/family/[custom-image-family-name]If the URI is unspecified, it will be inferred from SoftwareConfig.image_version or the system default.
        /// </summary>
        public readonly string ImageUri;
        /// <summary>
        /// The list of instance names. Dataproc derives the names from cluster_name, num_instances, and the instance group.
        /// </summary>
        public readonly ImmutableArray<string> InstanceNames;
        /// <summary>
        /// List of references to Compute Engine instances.
        /// </summary>
        public readonly ImmutableArray<Outputs.InstanceReferenceResponse> InstanceReferences;
        /// <summary>
        /// Specifies that this instance group contains preemptible instances.
        /// </summary>
        public readonly bool IsPreemptible;
        /// <summary>
        /// Optional. The Compute Engine machine type used for cluster instances.A full URL, partial URI, or short name are valid. Examples: https://www.googleapis.com/compute/v1/projects/[project_id]/zones/us-east1-a/machineTypes/n1-standard-2 projects/[project_id]/zones/us-east1-a/machineTypes/n1-standard-2 n1-standard-2Auto Zone Exception: If you are using the Dataproc Auto Zone Placement (https://cloud.google.com/dataproc/docs/concepts/configuring-clusters/auto-zone#using_auto_zone_placement) feature, you must use the short name of the machine type resource, for example, n1-standard-2.
        /// </summary>
        public readonly string MachineTypeUri;
        /// <summary>
        /// The config for Compute Engine Instance Group Manager that manages this group. This is only used for preemptible instance groups.
        /// </summary>
        public readonly Outputs.ManagedGroupConfigResponse ManagedGroupConfig;
        /// <summary>
        /// Optional. Specifies the minimum cpu platform for the Instance Group. See Dataproc -&gt; Minimum CPU Platform (https://cloud.google.com/dataproc/docs/concepts/compute/dataproc-min-cpu).
        /// </summary>
        public readonly string MinCpuPlatform;
        /// <summary>
        /// Optional. The number of VM instances in the instance group. For HA cluster master_config groups, must be set to 3. For standard cluster master_config groups, must be set to 1.
        /// </summary>
        public readonly int NumInstances;
        /// <summary>
        /// Optional. Specifies the preemptibility of the instance group.The default value for master and worker groups is NON_PREEMPTIBLE. This default cannot be changed.The default value for secondary instances is PREEMPTIBLE.
        /// </summary>
        public readonly string Preemptibility;

        [OutputConstructor]
        private InstanceGroupConfigResponse(
            ImmutableArray<Outputs.AcceleratorConfigResponse> accelerators,

            Outputs.DiskConfigResponse diskConfig,

            string imageUri,

            ImmutableArray<string> instanceNames,

            ImmutableArray<Outputs.InstanceReferenceResponse> instanceReferences,

            bool isPreemptible,

            string machineTypeUri,

            Outputs.ManagedGroupConfigResponse managedGroupConfig,

            string minCpuPlatform,

            int numInstances,

            string preemptibility)
        {
            Accelerators = accelerators;
            DiskConfig = diskConfig;
            ImageUri = imageUri;
            InstanceNames = instanceNames;
            InstanceReferences = instanceReferences;
            IsPreemptible = isPreemptible;
            MachineTypeUri = machineTypeUri;
            ManagedGroupConfig = managedGroupConfig;
            MinCpuPlatform = minCpuPlatform;
            NumInstances = numInstances;
            Preemptibility = preemptibility;
        }
    }
}
