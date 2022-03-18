// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.SQLAdmin.V1
{
    /// <summary>
    /// Creates a new backup run on demand.
    /// Auto-naming is currently not supported for this resource.
    /// </summary>
    [GoogleNativeResourceType("google-native:sqladmin/v1:BackupRun")]
    public partial class BackupRun : Pulumi.CustomResource
    {
        /// <summary>
        /// Specifies the kind of backup, PHYSICAL or DEFAULT_SNAPSHOT.
        /// </summary>
        [Output("backupKind")]
        public Output<string> BackupKind { get; private set; } = null!;

        /// <summary>
        /// The description of this run, only applicable to on-demand backups.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// Encryption configuration specific to a backup.
        /// </summary>
        [Output("diskEncryptionConfiguration")]
        public Output<Outputs.DiskEncryptionConfigurationResponse> DiskEncryptionConfiguration { get; private set; } = null!;

        /// <summary>
        /// Encryption status specific to a backup.
        /// </summary>
        [Output("diskEncryptionStatus")]
        public Output<Outputs.DiskEncryptionStatusResponse> DiskEncryptionStatus { get; private set; } = null!;

        /// <summary>
        /// The time the backup operation completed in UTC timezone in [RFC 3339](https://tools.ietf.org/html/rfc3339) format, for example `2012-11-15T16:19:00.094Z`.
        /// </summary>
        [Output("endTime")]
        public Output<string> EndTime { get; private set; } = null!;

        /// <summary>
        /// The time the run was enqueued in UTC timezone in [RFC 3339](https://tools.ietf.org/html/rfc3339) format, for example `2012-11-15T16:19:00.094Z`.
        /// </summary>
        [Output("enqueuedTime")]
        public Output<string> EnqueuedTime { get; private set; } = null!;

        /// <summary>
        /// Information about why the backup operation failed. This is only present if the run has the FAILED status.
        /// </summary>
        [Output("error")]
        public Output<Outputs.OperationErrorResponse> Error { get; private set; } = null!;

        /// <summary>
        /// Name of the database instance.
        /// </summary>
        [Output("instance")]
        public Output<string> Instance { get; private set; } = null!;

        /// <summary>
        /// This is always `sql#backupRun`.
        /// </summary>
        [Output("kind")]
        public Output<string> Kind { get; private set; } = null!;

        /// <summary>
        /// Location of the backups.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// The URI of this resource.
        /// </summary>
        [Output("selfLink")]
        public Output<string> SelfLink { get; private set; } = null!;

        /// <summary>
        /// The time the backup operation actually started in UTC timezone in [RFC 3339](https://tools.ietf.org/html/rfc3339) format, for example `2012-11-15T16:19:00.094Z`.
        /// </summary>
        [Output("startTime")]
        public Output<string> StartTime { get; private set; } = null!;

        /// <summary>
        /// The status of this run.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// The type of this run; can be either "AUTOMATED" or "ON_DEMAND". This field defaults to "ON_DEMAND" and is ignored, when specified for insert requests.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// The start time of the backup window during which this the backup was attempted in [RFC 3339](https://tools.ietf.org/html/rfc3339) format, for example `2012-11-15T16:19:00.094Z`.
        /// </summary>
        [Output("windowStartTime")]
        public Output<string> WindowStartTime { get; private set; } = null!;


        /// <summary>
        /// Create a BackupRun resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public BackupRun(string name, BackupRunArgs args, CustomResourceOptions? options = null)
            : base("google-native:sqladmin/v1:BackupRun", name, args ?? new BackupRunArgs(), MakeResourceOptions(options, ""))
        {
        }

        private BackupRun(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:sqladmin/v1:BackupRun", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing BackupRun resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static BackupRun Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new BackupRun(name, id, options);
        }
    }

    public sealed class BackupRunArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the kind of backup, PHYSICAL or DEFAULT_SNAPSHOT.
        /// </summary>
        [Input("backupKind")]
        public Input<Pulumi.GoogleNative.SQLAdmin.V1.BackupRunBackupKind>? BackupKind { get; set; }

        /// <summary>
        /// The description of this run, only applicable to on-demand backups.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Encryption configuration specific to a backup.
        /// </summary>
        [Input("diskEncryptionConfiguration")]
        public Input<Inputs.DiskEncryptionConfigurationArgs>? DiskEncryptionConfiguration { get; set; }

        /// <summary>
        /// Encryption status specific to a backup.
        /// </summary>
        [Input("diskEncryptionStatus")]
        public Input<Inputs.DiskEncryptionStatusArgs>? DiskEncryptionStatus { get; set; }

        /// <summary>
        /// The time the backup operation completed in UTC timezone in [RFC 3339](https://tools.ietf.org/html/rfc3339) format, for example `2012-11-15T16:19:00.094Z`.
        /// </summary>
        [Input("endTime")]
        public Input<string>? EndTime { get; set; }

        /// <summary>
        /// The time the run was enqueued in UTC timezone in [RFC 3339](https://tools.ietf.org/html/rfc3339) format, for example `2012-11-15T16:19:00.094Z`.
        /// </summary>
        [Input("enqueuedTime")]
        public Input<string>? EnqueuedTime { get; set; }

        /// <summary>
        /// Information about why the backup operation failed. This is only present if the run has the FAILED status.
        /// </summary>
        [Input("error")]
        public Input<Inputs.OperationErrorArgs>? Error { get; set; }

        /// <summary>
        /// The identifier for this backup run. Unique only for a specific Cloud SQL instance.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// Name of the database instance.
        /// </summary>
        [Input("instance", required: true)]
        public Input<string> Instance { get; set; } = null!;

        /// <summary>
        /// This is always `sql#backupRun`.
        /// </summary>
        [Input("kind")]
        public Input<string>? Kind { get; set; }

        /// <summary>
        /// Location of the backups.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The URI of this resource.
        /// </summary>
        [Input("selfLink")]
        public Input<string>? SelfLink { get; set; }

        /// <summary>
        /// The time the backup operation actually started in UTC timezone in [RFC 3339](https://tools.ietf.org/html/rfc3339) format, for example `2012-11-15T16:19:00.094Z`.
        /// </summary>
        [Input("startTime")]
        public Input<string>? StartTime { get; set; }

        /// <summary>
        /// The type of this run; can be either "AUTOMATED" or "ON_DEMAND". This field defaults to "ON_DEMAND" and is ignored, when specified for insert requests.
        /// </summary>
        [Input("type")]
        public Input<Pulumi.GoogleNative.SQLAdmin.V1.BackupRunType>? Type { get; set; }

        /// <summary>
        /// The start time of the backup window during which this the backup was attempted in [RFC 3339](https://tools.ietf.org/html/rfc3339) format, for example `2012-11-15T16:19:00.094Z`.
        /// </summary>
        [Input("windowStartTime")]
        public Input<string>? WindowStartTime { get; set; }

        public BackupRunArgs()
        {
        }
    }
}
