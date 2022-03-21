// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.AppEngine.V1Beta.Outputs
{

    /// <summary>
    /// VPC access connector specification.
    /// </summary>
    [OutputType]
    public sealed class VpcAccessConnectorResponse
    {
        /// <summary>
        /// The egress setting for the connector, controlling what traffic is diverted through it.
        /// </summary>
        public readonly string EgressSetting;
        /// <summary>
        /// Full Serverless VPC Access Connector name e.g. /projects/my-project/locations/us-central1/connectors/c1.
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private VpcAccessConnectorResponse(
            string egressSetting,

            string name)
        {
            EgressSetting = egressSetting;
            Name = name;
        }
    }
}
