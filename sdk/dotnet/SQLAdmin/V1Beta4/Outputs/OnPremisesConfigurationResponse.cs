// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.SQLAdmin.V1Beta4.Outputs
{

    /// <summary>
    /// On-premises instance configuration.
    /// </summary>
    [OutputType]
    public sealed class OnPremisesConfigurationResponse
    {
        /// <summary>
        /// PEM representation of the trusted CA's x509 certificate.
        /// </summary>
        public readonly string CaCertificate;
        /// <summary>
        /// PEM representation of the replica's x509 certificate.
        /// </summary>
        public readonly string ClientCertificate;
        /// <summary>
        /// PEM representation of the replica's private key. The corresponsing public key is encoded in the client's certificate.
        /// </summary>
        public readonly string ClientKey;
        /// <summary>
        /// The dump file to create the Cloud SQL replica.
        /// </summary>
        public readonly string DumpFilePath;
        /// <summary>
        /// The host and port of the on-premises instance in host:port format
        /// </summary>
        public readonly string HostPort;
        /// <summary>
        /// This is always `sql#onPremisesConfiguration`.
        /// </summary>
        public readonly string Kind;
        /// <summary>
        /// The password for connecting to on-premises instance.
        /// </summary>
        public readonly string Password;
        /// <summary>
        /// The reference to Cloud SQL instance if the source is Cloud SQL.
        /// </summary>
        public readonly Outputs.InstanceReferenceResponse SourceInstance;
        /// <summary>
        /// The username for connecting to on-premises instance.
        /// </summary>
        public readonly string Username;

        [OutputConstructor]
        private OnPremisesConfigurationResponse(
            string caCertificate,

            string clientCertificate,

            string clientKey,

            string dumpFilePath,

            string hostPort,

            string kind,

            string password,

            Outputs.InstanceReferenceResponse sourceInstance,

            string username)
        {
            CaCertificate = caCertificate;
            ClientCertificate = clientCertificate;
            ClientKey = clientKey;
            DumpFilePath = dumpFilePath;
            HostPort = hostPort;
            Kind = kind;
            Password = password;
            SourceInstance = sourceInstance;
            Username = username;
        }
    }
}
