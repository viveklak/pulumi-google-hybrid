// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.SQLAdmin.V1.Inputs
{

    /// <summary>
    /// Read-replica configuration specific to MySQL databases.
    /// </summary>
    public sealed class MySqlReplicaConfigurationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// PEM representation of the trusted CA's x509 certificate.
        /// </summary>
        [Input("caCertificate")]
        public Input<string>? CaCertificate { get; set; }

        /// <summary>
        /// PEM representation of the replica's x509 certificate.
        /// </summary>
        [Input("clientCertificate")]
        public Input<string>? ClientCertificate { get; set; }

        /// <summary>
        /// PEM representation of the replica's private key. The corresponsing public key is encoded in the client's certificate.
        /// </summary>
        [Input("clientKey")]
        public Input<string>? ClientKey { get; set; }

        /// <summary>
        /// Seconds to wait between connect retries. MySQL's default is 60 seconds.
        /// </summary>
        [Input("connectRetryInterval")]
        public Input<int>? ConnectRetryInterval { get; set; }

        /// <summary>
        /// Path to a SQL dump file in Google Cloud Storage from which the replica instance is to be created. The URI is in the form gs://bucketName/fileName. Compressed gzip files (.gz) are also supported. Dumps have the binlog co-ordinates from which replication begins. This can be accomplished by setting --master-data to 1 when using mysqldump.
        /// </summary>
        [Input("dumpFilePath")]
        public Input<string>? DumpFilePath { get; set; }

        /// <summary>
        /// This is always `sql#mysqlReplicaConfiguration`.
        /// </summary>
        [Input("kind")]
        public Input<string>? Kind { get; set; }

        /// <summary>
        /// Interval in milliseconds between replication heartbeats.
        /// </summary>
        [Input("masterHeartbeatPeriod")]
        public Input<string>? MasterHeartbeatPeriod { get; set; }

        /// <summary>
        /// The password for the replication connection.
        /// </summary>
        [Input("password")]
        public Input<string>? Password { get; set; }

        /// <summary>
        /// A list of permissible ciphers to use for SSL encryption.
        /// </summary>
        [Input("sslCipher")]
        public Input<string>? SslCipher { get; set; }

        /// <summary>
        /// The username for the replication connection.
        /// </summary>
        [Input("username")]
        public Input<string>? Username { get; set; }

        /// <summary>
        /// Whether or not to check the primary instance's Common Name value in the certificate that it sends during the SSL handshake.
        /// </summary>
        [Input("verifyServerCertificate")]
        public Input<bool>? VerifyServerCertificate { get; set; }

        public MySqlReplicaConfigurationArgs()
        {
        }
    }
}
