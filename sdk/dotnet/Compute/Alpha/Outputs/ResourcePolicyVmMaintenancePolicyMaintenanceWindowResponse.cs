// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Alpha.Outputs
{

    /// <summary>
    /// A maintenance window for VMs. When set, we restrict our maintenance operations to this window.
    /// </summary>
    [OutputType]
    public sealed class ResourcePolicyVmMaintenancePolicyMaintenanceWindowResponse
    {
        public readonly Outputs.ResourcePolicyDailyCycleResponse DailyMaintenanceWindow;

        [OutputConstructor]
        private ResourcePolicyVmMaintenancePolicyMaintenanceWindowResponse(Outputs.ResourcePolicyDailyCycleResponse dailyMaintenanceWindow)
        {
            DailyMaintenanceWindow = dailyMaintenanceWindow;
        }
    }
}
