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
    /// An InstanceSchedulePolicy specifies when and how frequent certain operations are performed on the instance.
    /// </summary>
    public sealed class ResourcePolicyInstanceSchedulePolicyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The expiration time of the schedule. The timestamp is an RFC3339 string.
        /// </summary>
        [Input("expirationTime")]
        public Input<string>? ExpirationTime { get; set; }

        /// <summary>
        /// The start time of the schedule. The timestamp is an RFC3339 string.
        /// </summary>
        [Input("startTime")]
        public Input<string>? StartTime { get; set; }

        /// <summary>
        /// Specifies the time zone to be used in interpreting Schedule.schedule. The value of this field must be a time zone name from the tz database: http://en.wikipedia.org/wiki/Tz_database.
        /// </summary>
        [Input("timeZone")]
        public Input<string>? TimeZone { get; set; }

        /// <summary>
        /// Specifies the schedule for starting instances.
        /// </summary>
        [Input("vmStartSchedule")]
        public Input<Inputs.ResourcePolicyInstanceSchedulePolicyScheduleArgs>? VmStartSchedule { get; set; }

        /// <summary>
        /// Specifies the schedule for stopping instances.
        /// </summary>
        [Input("vmStopSchedule")]
        public Input<Inputs.ResourcePolicyInstanceSchedulePolicyScheduleArgs>? VmStopSchedule { get; set; }

        public ResourcePolicyInstanceSchedulePolicyArgs()
        {
        }
    }
}
