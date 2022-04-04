// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Run.V2.Outputs
{

    /// <summary>
    /// Volume represents a named volume in a container.
    /// </summary>
    [OutputType]
    public sealed class GoogleCloudRunV2VolumeResponse
    {
        /// <summary>
        /// For Cloud SQL volumes, contains the specific instances that should be mounted. Visit https://cloud.google.com/sql/docs/mysql/connect-run for more information on how to connect Cloud SQL and Cloud Run.
        /// </summary>
        public readonly Outputs.GoogleCloudRunV2CloudSqlInstanceResponse CloudSqlInstance;
        /// <summary>
        /// Volume's name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Secret represents a secret that should populate this volume. More info: https://kubernetes.io/docs/concepts/storage/volumes#secret
        /// </summary>
        public readonly Outputs.GoogleCloudRunV2SecretVolumeSourceResponse Secret;

        [OutputConstructor]
        private GoogleCloudRunV2VolumeResponse(
            Outputs.GoogleCloudRunV2CloudSqlInstanceResponse cloudSqlInstance,

            string name,

            Outputs.GoogleCloudRunV2SecretVolumeSourceResponse secret)
        {
            CloudSqlInstance = cloudSqlInstance;
            Name = name;
            Secret = secret;
        }
    }
}