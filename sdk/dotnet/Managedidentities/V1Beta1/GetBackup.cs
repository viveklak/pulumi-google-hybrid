// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi.Utilities;

namespace Pulumi.GoogleNative.Managedidentities.V1Beta1
{
    public static class GetBackup
    {
        /// <summary>
        /// Gets details of a single Backup.
        /// </summary>
        public static Task<GetBackupResult> InvokeAsync(GetBackupArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetBackupResult>("google-native:managedidentities/v1beta1:getBackup", args ?? new GetBackupArgs(), options.WithDefaults());

        /// <summary>
        /// Gets details of a single Backup.
        /// </summary>
        public static Output<GetBackupResult> Invoke(GetBackupInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetBackupResult>("google-native:managedidentities/v1beta1:getBackup", args ?? new GetBackupInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetBackupArgs : Pulumi.InvokeArgs
    {
        [Input("backupId", required: true)]
        public string BackupId { get; set; } = null!;

        [Input("domainId", required: true)]
        public string DomainId { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        public GetBackupArgs()
        {
        }
    }

    public sealed class GetBackupInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("backupId", required: true)]
        public Input<string> BackupId { get; set; } = null!;

        [Input("domainId", required: true)]
        public Input<string> DomainId { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        public GetBackupInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetBackupResult
    {
        /// <summary>
        /// The time the backups was created.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// Optional. Resource labels to represent user provided metadata.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Labels;
        /// <summary>
        /// The unique name of the Backup in the form of projects/{project_id}/locations/global/domains/{domain_name}/backups/{name}
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The current state of the backup.
        /// </summary>
        public readonly string State;
        /// <summary>
        /// Additional information about the current status of this backup, if available.
        /// </summary>
        public readonly string StatusMessage;
        /// <summary>
        /// Indicates whether it’s an on-demand backup or scheduled.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// Last update time.
        /// </summary>
        public readonly string UpdateTime;

        [OutputConstructor]
        private GetBackupResult(
            string createTime,

            ImmutableDictionary<string, string> labels,

            string name,

            string state,

            string statusMessage,

            string type,

            string updateTime)
        {
            CreateTime = createTime;
            Labels = labels;
            Name = name;
            State = state;
            StatusMessage = statusMessage;
            Type = type;
            UpdateTime = updateTime;
        }
    }
}
