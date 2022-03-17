// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.SQLAdmin.V1Beta4.Inputs
{

    /// <summary>
    /// SQL Server specific audit configuration.
    /// </summary>
    public sealed class SqlServerAuditConfigArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the destination bucket (e.g., gs://mybucket).
        /// </summary>
        [Input("bucket")]
        public Input<string>? Bucket { get; set; }

        /// <summary>
        /// This is always sql#sqlServerAuditConfig
        /// </summary>
        [Input("kind")]
        public Input<string>? Kind { get; set; }

        public SqlServerAuditConfigArgs()
        {
        }
    }
}
