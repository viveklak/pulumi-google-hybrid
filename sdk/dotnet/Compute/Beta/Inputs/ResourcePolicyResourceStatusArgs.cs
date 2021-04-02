// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleCloud.Compute.Beta.Inputs
{

    /// <summary>
    /// Contains output only fields. Use this sub-message for all output fields set on ResourcePolicy. The internal structure of this "status" field should mimic the structure of ResourcePolicy proto specification.
    /// </summary>
    public sealed class ResourcePolicyResourceStatusArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// [Output Only] Specifies a set of output values reffering to the instance_schedule_policy system status. This field should have the same name as corresponding policy field.
        /// </summary>
        [Input("instanceSchedulePolicy")]
        public Input<Inputs.ResourcePolicyResourceStatusInstanceSchedulePolicyStatusArgs>? InstanceSchedulePolicy { get; set; }

        public ResourcePolicyResourceStatusArgs()
        {
        }
    }
}