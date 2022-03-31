// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Datastream.V1Alpha1.Outputs
{

    /// <summary>
    /// MySQL database profile.
    /// </summary>
    [OutputType]
    public sealed class MysqlProfileResponse
    {
        /// <summary>
        /// Hostname for the MySQL connection.
        /// </summary>
        public readonly string Hostname;
        /// <summary>
        /// Input only. Password for the MySQL connection.
        /// </summary>
        public readonly string Password;
        /// <summary>
        /// Port for the MySQL connection, default value is 3306.
        /// </summary>
        public readonly int Port;
        /// <summary>
        /// SSL configuration for the MySQL connection.
        /// </summary>
        public readonly Outputs.MysqlSslConfigResponse SslConfig;
        /// <summary>
        /// Username for the MySQL connection.
        /// </summary>
        public readonly string Username;

        [OutputConstructor]
        private MysqlProfileResponse(
            string hostname,

            string password,

            int port,

            Outputs.MysqlSslConfigResponse sslConfig,

            string username)
        {
            Hostname = hostname;
            Password = password;
            Port = port;
            SslConfig = sslConfig;
            Username = username;
        }
    }
}
