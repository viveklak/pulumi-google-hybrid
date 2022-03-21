// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Beta.Outputs
{

    /// <summary>
    /// An InstanceSchedulePolicy specifies when and how frequent certain operations are performed on the instance.
    /// </summary>
    [OutputType]
    public sealed class ResourcePolicyInstanceSchedulePolicyResponse
    {
        /// <summary>
        /// The expiration time of the schedule. The timestamp is an RFC3339 string.
        /// </summary>
        public readonly string ExpirationTime;
        /// <summary>
        /// The start time of the schedule. The timestamp is an RFC3339 string.
        /// </summary>
        public readonly string StartTime;
        /// <summary>
        /// Specifies the time zone to be used in interpreting Schedule.schedule. The value of this field must be a time zone name from the tz database: http://en.wikipedia.org/wiki/Tz_database.
        /// </summary>
        public readonly string TimeZone;
        /// <summary>
        /// Specifies the schedule for starting instances.
        /// </summary>
        public readonly Outputs.ResourcePolicyInstanceSchedulePolicyScheduleResponse VmStartSchedule;
        /// <summary>
        /// Specifies the schedule for stopping instances.
        /// </summary>
        public readonly Outputs.ResourcePolicyInstanceSchedulePolicyScheduleResponse VmStopSchedule;

        [OutputConstructor]
        private ResourcePolicyInstanceSchedulePolicyResponse(
            string expirationTime,

            string startTime,

            string timeZone,

            Outputs.ResourcePolicyInstanceSchedulePolicyScheduleResponse vmStartSchedule,

            Outputs.ResourcePolicyInstanceSchedulePolicyScheduleResponse vmStopSchedule)
        {
            ExpirationTime = expirationTime;
            StartTime = startTime;
            TimeZone = timeZone;
            VmStartSchedule = vmStartSchedule;
            VmStopSchedule = vmStopSchedule;
        }
    }
}
