// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Datamigration.V1Beta1.Inputs
{

    /// <summary>
    /// An entry for an Access Control list.
    /// </summary>
    public sealed class SqlAclEntryArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The time when this access control entry expires in [RFC 3339](https://tools.ietf.org/html/rfc3339) format, for example: `2012-11-15T16:19:00.094Z`.
        /// </summary>
        [Input("expireTime")]
        public Input<string>? ExpireTime { get; set; }

        /// <summary>
        /// A label to identify this entry.
        /// </summary>
        [Input("label")]
        public Input<string>? Label { get; set; }

        /// <summary>
        /// Input only. The time-to-leave of this access control entry.
        /// </summary>
        [Input("ttl")]
        public Input<string>? Ttl { get; set; }

        /// <summary>
        /// The allowlisted value for the access control list.
        /// </summary>
        [Input("value")]
        public Input<string>? Value { get; set; }

        public SqlAclEntryArgs()
        {
        }
    }
}
