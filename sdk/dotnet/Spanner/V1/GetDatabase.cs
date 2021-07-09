// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Spanner.V1
{
    public static class GetDatabase
    {
        /// <summary>
        /// Gets the state of a Cloud Spanner database.
        /// </summary>
        public static Task<GetDatabaseResult> InvokeAsync(GetDatabaseArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetDatabaseResult>("google-native:spanner/v1:getDatabase", args ?? new GetDatabaseArgs(), options.WithVersion());
    }


    public sealed class GetDatabaseArgs : Pulumi.InvokeArgs
    {
        [Input("databaseId", required: true)]
        public string DatabaseId { get; set; } = null!;

        [Input("instanceId", required: true)]
        public string InstanceId { get; set; } = null!;

        [Input("project", required: true)]
        public string Project { get; set; } = null!;

        public GetDatabaseArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetDatabaseResult
    {
        /// <summary>
        /// If exists, the time at which the database creation started.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// Earliest timestamp at which older versions of the data can be read. This value is continuously updated by Cloud Spanner and becomes stale the moment it is queried. If you are using this value to recover data, make sure to account for the time from the moment when the value is queried to the moment when you initiate the recovery.
        /// </summary>
        public readonly string EarliestVersionTime;
        /// <summary>
        /// For databases that are using customer managed encryption, this field contains the encryption configuration for the database. For databases that are using Google default or other types of encryption, this field is empty.
        /// </summary>
        public readonly Outputs.EncryptionConfigResponse EncryptionConfig;
        /// <summary>
        /// For databases that are using customer managed encryption, this field contains the encryption information for the database, such as encryption state and the Cloud KMS key versions that are in use. For databases that are using Google default or other types of encryption, this field is empty. This field is propagated lazily from the backend. There might be a delay from when a key version is being used and when it appears in this field.
        /// </summary>
        public readonly ImmutableArray<Outputs.EncryptionInfoResponse> EncryptionInfo;
        /// <summary>
        /// The name of the database. Values are of the form `projects//instances//databases/`, where `` is as specified in the `CREATE DATABASE` statement. This name can be passed to other API methods to identify the database.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Applicable only for restored databases. Contains information about the restore source.
        /// </summary>
        public readonly Outputs.RestoreInfoResponse RestoreInfo;
        /// <summary>
        /// The current database state.
        /// </summary>
        public readonly string State;
        /// <summary>
        /// The period in which Cloud Spanner retains all versions of data for the database. This is the same as the value of version_retention_period database option set using UpdateDatabaseDdl. Defaults to 1 hour, if not set.
        /// </summary>
        public readonly string VersionRetentionPeriod;

        [OutputConstructor]
        private GetDatabaseResult(
            string createTime,

            string earliestVersionTime,

            Outputs.EncryptionConfigResponse encryptionConfig,

            ImmutableArray<Outputs.EncryptionInfoResponse> encryptionInfo,

            string name,

            Outputs.RestoreInfoResponse restoreInfo,

            string state,

            string versionRetentionPeriod)
        {
            CreateTime = createTime;
            EarliestVersionTime = earliestVersionTime;
            EncryptionConfig = encryptionConfig;
            EncryptionInfo = encryptionInfo;
            Name = name;
            RestoreInfo = restoreInfo;
            State = state;
            VersionRetentionPeriod = versionRetentionPeriod;
        }
    }
}
