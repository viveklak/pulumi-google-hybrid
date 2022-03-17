// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.SQLAdmin.V1
{
    public static class GetBackupRun
    {
        /// <summary>
        /// Retrieves a resource containing information about a backup run.
        /// </summary>
        public static Task<GetBackupRunResult> InvokeAsync(GetBackupRunArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetBackupRunResult>("google-native:sqladmin/v1:getBackupRun", args ?? new GetBackupRunArgs(), options.WithDefaults());

        /// <summary>
        /// Retrieves a resource containing information about a backup run.
        /// </summary>
        public static Output<GetBackupRunResult> Invoke(GetBackupRunInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetBackupRunResult>("google-native:sqladmin/v1:getBackupRun", args ?? new GetBackupRunInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetBackupRunArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        [Input("instance", required: true)]
        public string Instance { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        public GetBackupRunArgs()
        {
        }
    }

    public sealed class GetBackupRunInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        [Input("instance", required: true)]
        public Input<string> Instance { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        public GetBackupRunInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetBackupRunResult
    {
        /// <summary>
        /// Specifies the kind of backup, PHYSICAL or DEFAULT_SNAPSHOT.
        /// </summary>
        public readonly string BackupKind;
        /// <summary>
        /// The description of this run, only applicable to on-demand backups.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// Encryption configuration specific to a backup.
        /// </summary>
        public readonly Outputs.DiskEncryptionConfigurationResponse DiskEncryptionConfiguration;
        /// <summary>
        /// Encryption status specific to a backup.
        /// </summary>
        public readonly Outputs.DiskEncryptionStatusResponse DiskEncryptionStatus;
        /// <summary>
        /// The time the backup operation completed in UTC timezone in [RFC 3339](https://tools.ietf.org/html/rfc3339) format, for example `2012-11-15T16:19:00.094Z`.
        /// </summary>
        public readonly string EndTime;
        /// <summary>
        /// The time the run was enqueued in UTC timezone in [RFC 3339](https://tools.ietf.org/html/rfc3339) format, for example `2012-11-15T16:19:00.094Z`.
        /// </summary>
        public readonly string EnqueuedTime;
        /// <summary>
        /// Information about why the backup operation failed. This is only present if the run has the FAILED status.
        /// </summary>
        public readonly Outputs.OperationErrorResponse Error;
        /// <summary>
        /// Name of the database instance.
        /// </summary>
        public readonly string Instance;
        /// <summary>
        /// This is always `sql#backupRun`.
        /// </summary>
        public readonly string Kind;
        /// <summary>
        /// Location of the backups.
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// The URI of this resource.
        /// </summary>
        public readonly string SelfLink;
        /// <summary>
        /// The time the backup operation actually started in UTC timezone in [RFC 3339](https://tools.ietf.org/html/rfc3339) format, for example `2012-11-15T16:19:00.094Z`.
        /// </summary>
        public readonly string StartTime;
        /// <summary>
        /// The status of this run.
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// The type of this run; can be either "AUTOMATED" or "ON_DEMAND". This field defaults to "ON_DEMAND" and is ignored, when specified for insert requests.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// The start time of the backup window during which this the backup was attempted in [RFC 3339](https://tools.ietf.org/html/rfc3339) format, for example `2012-11-15T16:19:00.094Z`.
        /// </summary>
        public readonly string WindowStartTime;

        [OutputConstructor]
        private GetBackupRunResult(
            string backupKind,

            string description,

            Outputs.DiskEncryptionConfigurationResponse diskEncryptionConfiguration,

            Outputs.DiskEncryptionStatusResponse diskEncryptionStatus,

            string endTime,

            string enqueuedTime,

            Outputs.OperationErrorResponse error,

            string instance,

            string kind,

            string location,

            string selfLink,

            string startTime,

            string status,

            string type,

            string windowStartTime)
        {
            BackupKind = backupKind;
            Description = description;
            DiskEncryptionConfiguration = diskEncryptionConfiguration;
            DiskEncryptionStatus = diskEncryptionStatus;
            EndTime = endTime;
            EnqueuedTime = enqueuedTime;
            Error = error;
            Instance = instance;
            Kind = kind;
            Location = location;
            SelfLink = selfLink;
            StartTime = startTime;
            Status = status;
            Type = type;
            WindowStartTime = windowStartTime;
        }
    }
}
