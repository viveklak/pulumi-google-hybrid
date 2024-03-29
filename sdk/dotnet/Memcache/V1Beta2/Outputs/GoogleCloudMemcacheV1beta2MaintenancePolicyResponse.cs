// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Memcache.V1Beta2.Outputs
{

    /// <summary>
    /// Maintenance policy per instance.
    /// </summary>
    [OutputType]
    public sealed class GoogleCloudMemcacheV1beta2MaintenancePolicyResponse
    {
        /// <summary>
        /// The time when the policy was created.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// Description of what this policy is for. Create/Update methods return INVALID_ARGUMENT if the length is greater than 512.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The time when the policy was updated.
        /// </summary>
        public readonly string UpdateTime;
        /// <summary>
        /// Maintenance window that is applied to resources covered by this policy. Minimum 1. For the current version, the maximum number of weekly_maintenance_windows is expected to be one.
        /// </summary>
        public readonly ImmutableArray<Outputs.WeeklyMaintenanceWindowResponse> WeeklyMaintenanceWindow;

        [OutputConstructor]
        private GoogleCloudMemcacheV1beta2MaintenancePolicyResponse(
            string createTime,

            string description,

            string updateTime,

            ImmutableArray<Outputs.WeeklyMaintenanceWindowResponse> weeklyMaintenanceWindow)
        {
            CreateTime = createTime;
            Description = description;
            UpdateTime = updateTime;
            WeeklyMaintenanceWindow = weeklyMaintenanceWindow;
        }
    }
}
