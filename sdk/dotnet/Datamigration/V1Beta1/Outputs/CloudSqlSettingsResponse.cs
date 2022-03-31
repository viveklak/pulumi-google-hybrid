// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Datamigration.V1Beta1.Outputs
{

    /// <summary>
    /// Settings for creating a Cloud SQL database instance.
    /// </summary>
    [OutputType]
    public sealed class CloudSqlSettingsResponse
    {
        /// <summary>
        /// The activation policy specifies when the instance is activated; it is applicable only when the instance state is 'RUNNABLE'. Valid values: 'ALWAYS': The instance is on, and remains so even in the absence of connection requests. `NEVER`: The instance is off; it is not activated, even if a connection request arrives.
        /// </summary>
        public readonly string ActivationPolicy;
        /// <summary>
        /// [default: ON] If you enable this setting, Cloud SQL checks your available storage every 30 seconds. If the available storage falls below a threshold size, Cloud SQL automatically adds additional storage capacity. If the available storage repeatedly falls below the threshold size, Cloud SQL continues to add storage until it reaches the maximum of 30 TB.
        /// </summary>
        public readonly bool AutoStorageIncrease;
        /// <summary>
        /// The storage capacity available to the database, in GB. The minimum (and default) size is 10GB.
        /// </summary>
        public readonly string DataDiskSizeGb;
        /// <summary>
        /// The type of storage: `PD_SSD` (default) or `PD_HDD`.
        /// </summary>
        public readonly string DataDiskType;
        /// <summary>
        /// The database flags passed to the Cloud SQL instance at startup. An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
        /// </summary>
        public readonly ImmutableDictionary<string, string> DatabaseFlags;
        /// <summary>
        /// The database engine type and version.
        /// </summary>
        public readonly string DatabaseVersion;
        /// <summary>
        /// The settings for IP Management. This allows to enable or disable the instance IP and manage which external networks can connect to the instance. The IPv4 address cannot be disabled.
        /// </summary>
        public readonly Outputs.SqlIpConfigResponse IpConfig;
        /// <summary>
        /// Input only. Initial root password.
        /// </summary>
        public readonly string RootPassword;
        /// <summary>
        /// Indicates If this connection profile root password is stored.
        /// </summary>
        public readonly bool RootPasswordSet;
        /// <summary>
        /// The Database Migration Service source connection profile ID, in the format: `projects/my_project_name/locations/us-central1/connectionProfiles/connection_profile_ID`
        /// </summary>
        public readonly string SourceId;
        /// <summary>
        /// The maximum size to which storage capacity can be automatically increased. The default value is 0, which specifies that there is no limit.
        /// </summary>
        public readonly string StorageAutoResizeLimit;
        /// <summary>
        /// The tier (or machine type) for this instance, for example: `db-n1-standard-1` (MySQL instances). For more information, see [Cloud SQL Instance Settings](https://cloud.google.com/sql/docs/mysql/instance-settings).
        /// </summary>
        public readonly string Tier;
        /// <summary>
        /// The resource labels for a Cloud SQL instance to use to annotate any related underlying resources such as Compute Engine VMs. An object containing a list of "key": "value" pairs. Example: `{ "name": "wrench", "mass": "18kg", "count": "3" }`.
        /// </summary>
        public readonly ImmutableDictionary<string, string> UserLabels;
        /// <summary>
        /// The Google Cloud Platform zone where your Cloud SQL datdabse instance is located.
        /// </summary>
        public readonly string Zone;

        [OutputConstructor]
        private CloudSqlSettingsResponse(
            string activationPolicy,

            bool autoStorageIncrease,

            string dataDiskSizeGb,

            string dataDiskType,

            ImmutableDictionary<string, string> databaseFlags,

            string databaseVersion,

            Outputs.SqlIpConfigResponse ipConfig,

            string rootPassword,

            bool rootPasswordSet,

            string sourceId,

            string storageAutoResizeLimit,

            string tier,

            ImmutableDictionary<string, string> userLabels,

            string zone)
        {
            ActivationPolicy = activationPolicy;
            AutoStorageIncrease = autoStorageIncrease;
            DataDiskSizeGb = dataDiskSizeGb;
            DataDiskType = dataDiskType;
            DatabaseFlags = databaseFlags;
            DatabaseVersion = databaseVersion;
            IpConfig = ipConfig;
            RootPassword = rootPassword;
            RootPasswordSet = rootPasswordSet;
            SourceId = sourceId;
            StorageAutoResizeLimit = storageAutoResizeLimit;
            Tier = tier;
            UserLabels = userLabels;
            Zone = zone;
        }
    }
}
