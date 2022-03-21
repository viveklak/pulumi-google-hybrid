// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Datamigration.V1.Inputs
{

    /// <summary>
    /// The details of the VPC where the source database is located in Google Cloud. We will use this information to set up the VPC peering connection between Cloud SQL and this VPC.
    /// </summary>
    public sealed class VpcPeeringConnectivityArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the VPC network to peer with the Cloud SQL private network.
        /// </summary>
        [Input("vpc")]
        public Input<string>? Vpc { get; set; }

        public VpcPeeringConnectivityArgs()
        {
        }
    }
}
