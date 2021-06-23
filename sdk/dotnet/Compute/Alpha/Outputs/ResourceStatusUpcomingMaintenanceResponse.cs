// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Alpha.Outputs
{

    [OutputType]
    public sealed class ResourceStatusUpcomingMaintenanceResponse
    {
        /// <summary>
        /// Indicates if the maintenance can be customer triggered. See go/sf-ctm-design for more details
        /// </summary>
        public readonly bool CanReschedule;

        [OutputConstructor]
        private ResourceStatusUpcomingMaintenanceResponse(bool canReschedule)
        {
            CanReschedule = canReschedule;
        }
    }
}
