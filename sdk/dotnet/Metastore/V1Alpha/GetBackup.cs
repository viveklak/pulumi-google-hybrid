// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi.Utilities;

namespace Pulumi.GoogleNative.Metastore.V1Alpha
{
    public static class GetBackup
    {
        /// <summary>
        /// Gets details of a single backup.
        /// </summary>
        public static Task<GetBackupResult> InvokeAsync(GetBackupArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetBackupResult>("google-native:metastore/v1alpha:getBackup", args ?? new GetBackupArgs(), options.WithDefaults());

        /// <summary>
        /// Gets details of a single backup.
        /// </summary>
        public static Output<GetBackupResult> Invoke(GetBackupInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetBackupResult>("google-native:metastore/v1alpha:getBackup", args ?? new GetBackupInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetBackupArgs : Pulumi.InvokeArgs
    {
        [Input("backupId", required: true)]
        public string BackupId { get; set; } = null!;

        [Input("location", required: true)]
        public string Location { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        [Input("serviceId", required: true)]
        public string ServiceId { get; set; } = null!;

        public GetBackupArgs()
        {
        }
    }

    public sealed class GetBackupInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("backupId", required: true)]
        public Input<string> BackupId { get; set; } = null!;

        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("serviceId", required: true)]
        public Input<string> ServiceId { get; set; } = null!;

        public GetBackupInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetBackupResult
    {
        /// <summary>
        /// The time when the backup was started.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// The description of the backup.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The time when the backup finished creating.
        /// </summary>
        public readonly string EndTime;
        /// <summary>
        /// Immutable. The relative resource name of the backup, in the following form:projects/{project_number}/locations/{location_id}/services/{service_id}/backups/{backup_id}
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Services that are restoring from the backup.
        /// </summary>
        public readonly ImmutableArray<string> RestoringServices;
        /// <summary>
        /// The revision of the service at the time of backup.
        /// </summary>
        public readonly Outputs.ServiceResponse ServiceRevision;
        /// <summary>
        /// The current state of the backup.
        /// </summary>
        public readonly string State;

        [OutputConstructor]
        private GetBackupResult(
            string createTime,

            string description,

            string endTime,

            string name,

            ImmutableArray<string> restoringServices,

            Outputs.ServiceResponse serviceRevision,

            string state)
        {
            CreateTime = createTime;
            Description = description;
            EndTime = endTime;
            Name = name;
            RestoringServices = restoringServices;
            ServiceRevision = serviceRevision;
            State = state;
        }
    }
}
