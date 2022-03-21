// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.V1.Inputs
{

    /// <summary>
    /// Deprecation status for a public resource.
    /// </summary>
    public sealed class DeprecationStatusArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// An optional RFC3339 timestamp on or after which the state of this resource is intended to change to DELETED. This is only informational and the status will not change unless the client explicitly changes it.
        /// </summary>
        [Input("deleted")]
        public Input<string>? Deleted { get; set; }

        /// <summary>
        /// An optional RFC3339 timestamp on or after which the state of this resource is intended to change to DEPRECATED. This is only informational and the status will not change unless the client explicitly changes it.
        /// </summary>
        [Input("deprecated")]
        public Input<string>? Deprecated { get; set; }

        /// <summary>
        /// An optional RFC3339 timestamp on or after which the state of this resource is intended to change to OBSOLETE. This is only informational and the status will not change unless the client explicitly changes it.
        /// </summary>
        [Input("obsolete")]
        public Input<string>? Obsolete { get; set; }

        /// <summary>
        /// The URL of the suggested replacement for a deprecated resource. The suggested replacement resource must be the same kind of resource as the deprecated resource.
        /// </summary>
        [Input("replacement")]
        public Input<string>? Replacement { get; set; }

        /// <summary>
        /// The deprecation state of this resource. This can be ACTIVE, DEPRECATED, OBSOLETE, or DELETED. Operations which communicate the end of life date for an image, can use ACTIVE. Operations which create a new resource using a DEPRECATED resource will return successfully, but with a warning indicating the deprecated resource and recommending its replacement. Operations which use OBSOLETE or DELETED resources will be rejected and result in an error.
        /// </summary>
        [Input("state")]
        public Input<Pulumi.GoogleNative.Compute.V1.DeprecationStatusState>? State { get; set; }

        public DeprecationStatusArgs()
        {
        }
    }
}
