// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.NetworkManagement.V1Beta1.Outputs
{

    [OutputType]
    public sealed class CloudSQLInstanceInfoResponse
    {
        /// <summary>
        /// Name of a Cloud SQL instance.
        /// </summary>
        public readonly string DisplayName;
        /// <summary>
        /// External IP address of a Cloud SQL instance.
        /// </summary>
        public readonly string ExternalIp;
        /// <summary>
        /// Internal IP address of a Cloud SQL instance.
        /// </summary>
        public readonly string InternalIp;
        /// <summary>
        /// URI of a Cloud SQL instance network or empty string if the instance does not have one.
        /// </summary>
        public readonly string NetworkUri;
        /// <summary>
        /// Region in which the Cloud SQL instance is running.
        /// </summary>
        public readonly string Region;
        /// <summary>
        /// URI of a Cloud SQL instance.
        /// </summary>
        public readonly string Uri;

        [OutputConstructor]
        private CloudSQLInstanceInfoResponse(
            string displayName,

            string externalIp,

            string internalIp,

            string networkUri,

            string region,

            string uri)
        {
            DisplayName = displayName;
            ExternalIp = externalIp;
            InternalIp = internalIp;
            NetworkUri = networkUri;
            Region = region;
            Uri = uri;
        }
    }
}
