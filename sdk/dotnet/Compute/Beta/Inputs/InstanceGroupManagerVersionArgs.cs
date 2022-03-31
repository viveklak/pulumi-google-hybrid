// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Beta.Inputs
{

    public sealed class InstanceGroupManagerVersionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The URL of the instance template that is specified for this managed instance group. The group uses this template to create new instances in the managed instance group until the `targetSize` for this version is reached. The templates for existing instances in the group do not change unless you run recreateInstances, run applyUpdatesToInstances, or set the group's updatePolicy.type to PROACTIVE; in those cases, existing instances are updated until the `targetSize` for this version is reached.
        /// </summary>
        [Input("instanceTemplate")]
        public Input<string>? InstanceTemplate { get; set; }

        /// <summary>
        /// Name of the version. Unique among all versions in the scope of this managed instance group.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Specifies the intended number of instances to be created from the instanceTemplate. The final number of instances created from the template will be equal to: - If expressed as a fixed number, the minimum of either targetSize.fixed or instanceGroupManager.targetSize is used. - if expressed as a percent, the targetSize would be (targetSize.percent/100 * InstanceGroupManager.targetSize) If there is a remainder, the number is rounded. If unset, this version will update any remaining instances not updated by another version. Read Starting a canary update for more information.
        /// </summary>
        [Input("targetSize")]
        public Input<Inputs.FixedOrPercentArgs>? TargetSize { get; set; }

        public InstanceGroupManagerVersionArgs()
        {
        }
    }
}
