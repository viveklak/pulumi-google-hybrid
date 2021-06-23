// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Container.V1Beta1.Outputs
{

    [OutputType]
    public sealed class AutoprovisioningNodePoolDefaultsResponse
    {
        /// <summary>
        ///  The Customer Managed Encryption Key used to encrypt the boot disk attached to each node in the node pool. This should be of the form projects/[KEY_PROJECT_ID]/locations/[LOCATION]/keyRings/[RING_NAME]/cryptoKeys/[KEY_NAME]. For more information about protecting resources with Cloud KMS Keys please see: https://cloud.google.com/compute/docs/disks/customer-managed-encryption
        /// </summary>
        public readonly string BootDiskKmsKey;
        /// <summary>
        /// Size of the disk attached to each node, specified in GB. The smallest allowed disk size is 10GB. If unspecified, the default disk size is 100GB.
        /// </summary>
        public readonly int DiskSizeGb;
        /// <summary>
        /// Type of the disk attached to each node (e.g. 'pd-standard', 'pd-ssd' or 'pd-balanced') If unspecified, the default disk type is 'pd-standard'
        /// </summary>
        public readonly string DiskType;
        /// <summary>
        /// The image type to use for NAP created node.
        /// </summary>
        public readonly string ImageType;
        /// <summary>
        /// NodeManagement configuration for this NodePool.
        /// </summary>
        public readonly Outputs.NodeManagementResponse Management;
        /// <summary>
        /// Minimum CPU platform to be used by this instance. The instance may be scheduled on the specified or newer CPU platform. Applicable values are the friendly names of CPU platforms, such as `minCpuPlatform: "Intel Haswell"` or `minCpuPlatform: "Intel Sandy Bridge"`. For more information, read [how to specify min CPU platform](https://cloud.google.com/compute/docs/instances/specify-min-cpu-platform) To unset the min cpu platform field pass "automatic" as field value.
        /// </summary>
        public readonly string MinCpuPlatform;
        /// <summary>
        /// The set of Google API scopes to be made available on all of the node VMs under the "default" service account. The following scopes are recommended, but not required, and by default are not included: * `https://www.googleapis.com/auth/compute` is required for mounting persistent storage on your nodes. * `https://www.googleapis.com/auth/devstorage.read_only` is required for communicating with **gcr.io** (the [Google Container Registry](https://cloud.google.com/container-registry/)). If unspecified, no scopes are added, unless Cloud Logging or Cloud Monitoring are enabled, in which case their required scopes will be added.
        /// </summary>
        public readonly ImmutableArray<string> OauthScopes;
        /// <summary>
        /// The Google Cloud Platform Service Account to be used by the node VMs. Specify the email address of the Service Account; otherwise, if no Service Account is specified, the "default" service account is used.
        /// </summary>
        public readonly string ServiceAccount;
        /// <summary>
        /// Shielded Instance options.
        /// </summary>
        public readonly Outputs.ShieldedInstanceConfigResponse ShieldedInstanceConfig;
        /// <summary>
        /// Upgrade settings control disruption and speed of the upgrade.
        /// </summary>
        public readonly Outputs.UpgradeSettingsResponse UpgradeSettings;

        [OutputConstructor]
        private AutoprovisioningNodePoolDefaultsResponse(
            string bootDiskKmsKey,

            int diskSizeGb,

            string diskType,

            string imageType,

            Outputs.NodeManagementResponse management,

            string minCpuPlatform,

            ImmutableArray<string> oauthScopes,

            string serviceAccount,

            Outputs.ShieldedInstanceConfigResponse shieldedInstanceConfig,

            Outputs.UpgradeSettingsResponse upgradeSettings)
        {
            BootDiskKmsKey = bootDiskKmsKey;
            DiskSizeGb = diskSizeGb;
            DiskType = diskType;
            ImageType = imageType;
            Management = management;
            MinCpuPlatform = minCpuPlatform;
            OauthScopes = oauthScopes;
            ServiceAccount = serviceAccount;
            ShieldedInstanceConfig = shieldedInstanceConfig;
            UpgradeSettings = upgradeSettings;
        }
    }
}
