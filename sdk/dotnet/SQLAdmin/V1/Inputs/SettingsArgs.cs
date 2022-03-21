// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.SQLAdmin.V1.Inputs
{

    /// <summary>
    /// Database instance settings.
    /// </summary>
    public sealed class SettingsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The activation policy specifies when the instance is activated; it is applicable only when the instance state is RUNNABLE. Valid values: * `ALWAYS`: The instance is on, and remains so even in the absence of connection requests. * `NEVER`: The instance is off; it is not activated, even if a connection request arrives.
        /// </summary>
        [Input("activationPolicy")]
        public Input<Pulumi.GoogleNative.SQLAdmin.V1.SettingsActivationPolicy>? ActivationPolicy { get; set; }

        /// <summary>
        /// Active Directory configuration, relevant only for Cloud SQL for SQL Server.
        /// </summary>
        [Input("activeDirectoryConfig")]
        public Input<Inputs.SqlActiveDirectoryConfigArgs>? ActiveDirectoryConfig { get; set; }

        [Input("authorizedGaeApplications")]
        private InputList<string>? _authorizedGaeApplications;

        /// <summary>
        /// The App Engine app IDs that can access this instance. (Deprecated) Applied to First Generation instances only.
        /// </summary>
        [Obsolete(@"The App Engine app IDs that can access this instance. (Deprecated) Applied to First Generation instances only.")]
        public InputList<string> AuthorizedGaeApplications
        {
            get => _authorizedGaeApplications ?? (_authorizedGaeApplications = new InputList<string>());
            set => _authorizedGaeApplications = value;
        }

        /// <summary>
        /// Availability type. Potential values: * `ZONAL`: The instance serves data from only one zone. Outages in that zone affect data accessibility. * `REGIONAL`: The instance can serve data from more than one zone in a region (it is highly available)./ For more information, see [Overview of the High Availability Configuration](https://cloud.google.com/sql/docs/mysql/high-availability).
        /// </summary>
        [Input("availabilityType")]
        public Input<Pulumi.GoogleNative.SQLAdmin.V1.SettingsAvailabilityType>? AvailabilityType { get; set; }

        /// <summary>
        /// The daily backup configuration for the instance.
        /// </summary>
        [Input("backupConfiguration")]
        public Input<Inputs.BackupConfigurationArgs>? BackupConfiguration { get; set; }

        /// <summary>
        /// The name of server Instance collation.
        /// </summary>
        [Input("collation")]
        public Input<string>? Collation { get; set; }

        /// <summary>
        /// Configuration specific to read replica instances. Indicates whether database flags for crash-safe replication are enabled. This property was only applicable to First Generation instances.
        /// </summary>
        [Input("crashSafeReplicationEnabled")]
        public Input<bool>? CrashSafeReplicationEnabled { get; set; }

        /// <summary>
        /// The size of data disk, in GB. The data disk size minimum is 10GB.
        /// </summary>
        [Input("dataDiskSizeGb")]
        public Input<string>? DataDiskSizeGb { get; set; }

        /// <summary>
        /// The type of data disk: `PD_SSD` (default) or `PD_HDD`. Not used for First Generation instances.
        /// </summary>
        [Input("dataDiskType")]
        public Input<Pulumi.GoogleNative.SQLAdmin.V1.SettingsDataDiskType>? DataDiskType { get; set; }

        [Input("databaseFlags")]
        private InputList<Inputs.DatabaseFlagsArgs>? _databaseFlags;

        /// <summary>
        /// The database flags passed to the instance at startup.
        /// </summary>
        public InputList<Inputs.DatabaseFlagsArgs> DatabaseFlags
        {
            get => _databaseFlags ?? (_databaseFlags = new InputList<Inputs.DatabaseFlagsArgs>());
            set => _databaseFlags = value;
        }

        /// <summary>
        /// Configuration specific to read replica instances. Indicates whether replication is enabled or not. WARNING: Changing this restarts the instance.
        /// </summary>
        [Input("databaseReplicationEnabled")]
        public Input<bool>? DatabaseReplicationEnabled { get; set; }

        [Input("denyMaintenancePeriods")]
        private InputList<Inputs.DenyMaintenancePeriodArgs>? _denyMaintenancePeriods;

        /// <summary>
        /// Deny maintenance periods
        /// </summary>
        public InputList<Inputs.DenyMaintenancePeriodArgs> DenyMaintenancePeriods
        {
            get => _denyMaintenancePeriods ?? (_denyMaintenancePeriods = new InputList<Inputs.DenyMaintenancePeriodArgs>());
            set => _denyMaintenancePeriods = value;
        }

        /// <summary>
        /// Insights configuration, for now relevant only for Postgres.
        /// </summary>
        [Input("insightsConfig")]
        public Input<Inputs.InsightsConfigArgs>? InsightsConfig { get; set; }

        /// <summary>
        /// The settings for IP Management. This allows to enable or disable the instance IP and manage which external networks can connect to the instance. The IPv4 address cannot be disabled for Second Generation instances.
        /// </summary>
        [Input("ipConfiguration")]
        public Input<Inputs.IpConfigurationArgs>? IpConfiguration { get; set; }

        /// <summary>
        /// This is always `sql#settings`.
        /// </summary>
        [Input("kind")]
        public Input<string>? Kind { get; set; }

        /// <summary>
        /// The location preference settings. This allows the instance to be located as near as possible to either an App Engine app or Compute Engine zone for better performance. App Engine co-location was only applicable to First Generation instances.
        /// </summary>
        [Input("locationPreference")]
        public Input<Inputs.LocationPreferenceArgs>? LocationPreference { get; set; }

        /// <summary>
        /// The maintenance window for this instance. This specifies when the instance can be restarted for maintenance purposes.
        /// </summary>
        [Input("maintenanceWindow")]
        public Input<Inputs.MaintenanceWindowArgs>? MaintenanceWindow { get; set; }

        /// <summary>
        /// The local user password validation policy of the instance.
        /// </summary>
        [Input("passwordValidationPolicy")]
        public Input<Inputs.PasswordValidationPolicyArgs>? PasswordValidationPolicy { get; set; }

        /// <summary>
        /// The pricing plan for this instance. This can be either `PER_USE` or `PACKAGE`. Only `PER_USE` is supported for Second Generation instances.
        /// </summary>
        [Input("pricingPlan")]
        public Input<Pulumi.GoogleNative.SQLAdmin.V1.SettingsPricingPlan>? PricingPlan { get; set; }

        /// <summary>
        /// The type of replication this instance uses. This can be either `ASYNCHRONOUS` or `SYNCHRONOUS`. (Deprecated) This property was only applicable to First Generation instances.
        /// </summary>
        [Input("replicationType")]
        public Input<Pulumi.GoogleNative.SQLAdmin.V1.SettingsReplicationType>? ReplicationType { get; set; }

        /// <summary>
        /// The version of instance settings. This is a required field for update method to make sure concurrent updates are handled properly. During update, use the most recent settingsVersion value for this instance and do not try to update this value.
        /// </summary>
        [Input("settingsVersion")]
        public Input<string>? SettingsVersion { get; set; }

        /// <summary>
        /// SQL Server specific audit configuration.
        /// </summary>
        [Input("sqlServerAuditConfig")]
        public Input<Inputs.SqlServerAuditConfigArgs>? SqlServerAuditConfig { get; set; }

        /// <summary>
        /// Configuration to increase storage size automatically. The default value is true.
        /// </summary>
        [Input("storageAutoResize")]
        public Input<bool>? StorageAutoResize { get; set; }

        /// <summary>
        /// The maximum size to which storage capacity can be automatically increased. The default value is 0, which specifies that there is no limit.
        /// </summary>
        [Input("storageAutoResizeLimit")]
        public Input<string>? StorageAutoResizeLimit { get; set; }

        /// <summary>
        /// The tier (or machine type) for this instance, for example `db-custom-1-3840`. WARNING: Changing this restarts the instance.
        /// </summary>
        [Input("tier")]
        public Input<string>? Tier { get; set; }

        [Input("userLabels")]
        private InputMap<string>? _userLabels;

        /// <summary>
        /// User-provided labels, represented as a dictionary where each label is a single key value pair.
        /// </summary>
        public InputMap<string> UserLabels
        {
            get => _userLabels ?? (_userLabels = new InputMap<string>());
            set => _userLabels = value;
        }

        public SettingsArgs()
        {
        }
    }
}
