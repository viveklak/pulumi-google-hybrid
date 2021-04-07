// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GcpNative.Metastore.V1Beta.Outputs
{

    [OutputType]
    public sealed class RestoreResponse
    {
        /// <summary>
        /// The relative resource name of the metastore service backup to restore from, in the following form:projects/{project_id}/locations/{location_id}/services/{service_id}/backups/{backup_id}.
        /// </summary>
        public readonly string Backup;
        /// <summary>
        /// The restore details containing the revision of the service to be restored to, in format of JSON.
        /// </summary>
        public readonly string Details;
        /// <summary>
        /// The time when the restore ended.
        /// </summary>
        public readonly string EndTime;
        /// <summary>
        /// The time when the restore started.
        /// </summary>
        public readonly string StartTime;
        /// <summary>
        /// The current state of the restore.
        /// </summary>
        public readonly string State;
        /// <summary>
        /// The type of restore.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private RestoreResponse(
            string backup,

            string details,

            string endTime,

            string startTime,

            string state,

            string type)
        {
            Backup = backup;
            Details = details;
            EndTime = endTime;
            StartTime = startTime;
            State = state;
            Type = type;
        }
    }
}