// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.V1.Inputs
{

    public sealed class DistributionPolicyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The distribution shape to which the group converges either proactively or on resize events (depending on the value set in updatePolicy.instanceRedistributionType).
        /// </summary>
        [Input("targetShape")]
        public Input<Pulumi.GoogleNative.Compute.V1.DistributionPolicyTargetShape>? TargetShape { get; set; }

        [Input("zones")]
        private InputList<Inputs.DistributionPolicyZoneConfigurationArgs>? _zones;

        /// <summary>
        /// Zones where the regional managed instance group will create and manage its instances.
        /// </summary>
        public InputList<Inputs.DistributionPolicyZoneConfigurationArgs> Zones
        {
            get => _zones ?? (_zones = new InputList<Inputs.DistributionPolicyZoneConfigurationArgs>());
            set => _zones = value;
        }

        public DistributionPolicyArgs()
        {
        }
    }
}
