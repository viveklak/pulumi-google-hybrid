// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Datamigration.V1Beta1.Inputs
{

    /// <summary>
    /// Settings for creating a Cloud SQL database instance.
    /// </summary>
    public sealed class CloudSqlSettingsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The activation policy specifies when the instance is activated; it is applicable only when the instance state is 'RUNNABLE'. Valid values: 'ALWAYS': The instance is on, and remains so even in the absence of connection requests. `NEVER`: The instance is off; it is not activated, even if a connection request arrives.
        /// </summary>
        [Input("activationPolicy")]
        public Input<Pulumi.GoogleNative.Datamigration.V1Beta1.CloudSqlSettingsActivationPolicy>? ActivationPolicy { get; set; }

        /// <summary>
        /// [default: ON] If you enable this setting, Cloud SQL checks your available storage every 30 seconds. If the available storage falls below a threshold size, Cloud SQL automatically adds additional storage capacity. If the available storage repeatedly falls below the threshold size, Cloud SQL continues to add storage until it reaches the maximum of 30 TB.
        /// </summary>
        [Input("autoStorageIncrease")]
        public Input<bool>? AutoStorageIncrease { get; set; }

        /// <summary>
        /// The storage capacity available to the database, in GB. The minimum (and default) size is 10GB.
        /// </summary>
        [Input("dataDiskSizeGb")]
        public Input<string>? DataDiskSizeGb { get; set; }

        /// <summary>
        /// The type of storage: `PD_SSD` (default) or `PD_HDD`.
        /// </summary>
        [Input("dataDiskType")]
        public Input<Pulumi.GoogleNative.Datamigration.V1Beta1.CloudSqlSettingsDataDiskType>? DataDiskType { get; set; }

        [Input("databaseFlags")]
        private InputMap<string>? _databaseFlags;

        /// <summary>
        /// The database flags passed to the Cloud SQL instance at startup. An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
        /// </summary>
        public InputMap<string> DatabaseFlags
        {
            get => _databaseFlags ?? (_databaseFlags = new InputMap<string>());
            set => _databaseFlags = value;
        }

        /// <summary>
        /// The database engine type and version.
        /// </summary>
        [Input("databaseVersion")]
        public Input<Pulumi.GoogleNative.Datamigration.V1Beta1.CloudSqlSettingsDatabaseVersion>? DatabaseVersion { get; set; }

        /// <summary>
        /// The settings for IP Management. This allows to enable or disable the instance IP and manage which external networks can connect to the instance. The IPv4 address cannot be disabled.
        /// </summary>
        [Input("ipConfig")]
        public Input<Inputs.SqlIpConfigArgs>? IpConfig { get; set; }

        /// <summary>
        /// Input only. Initial root password.
        /// </summary>
        [Input("rootPassword")]
        public Input<string>? RootPassword { get; set; }

        /// <summary>
        /// The Database Migration Service source connection profile ID, in the format: `projects/my_project_name/locations/us-central1/connectionProfiles/connection_profile_ID`
        /// </summary>
        [Input("sourceId")]
        public Input<string>? SourceId { get; set; }

        /// <summary>
        /// The maximum size to which storage capacity can be automatically increased. The default value is 0, which specifies that there is no limit.
        /// </summary>
        [Input("storageAutoResizeLimit")]
        public Input<string>? StorageAutoResizeLimit { get; set; }

        /// <summary>
        /// The tier (or machine type) for this instance, for example: `db-n1-standard-1` (MySQL instances). For more information, see [Cloud SQL Instance Settings](https://cloud.google.com/sql/docs/mysql/instance-settings).
        /// </summary>
        [Input("tier")]
        public Input<string>? Tier { get; set; }

        [Input("userLabels")]
        private InputMap<string>? _userLabels;

        /// <summary>
        /// The resource labels for a Cloud SQL instance to use to annotate any related underlying resources such as Compute Engine VMs. An object containing a list of "key": "value" pairs. Example: `{ "name": "wrench", "mass": "18kg", "count": "3" }`.
        /// </summary>
        public InputMap<string> UserLabels
        {
            get => _userLabels ?? (_userLabels = new InputMap<string>());
            set => _userLabels = value;
        }

        /// <summary>
        /// The Google Cloud Platform zone where your Cloud SQL datdabse instance is located.
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public CloudSqlSettingsArgs()
        {
        }
    }
}
