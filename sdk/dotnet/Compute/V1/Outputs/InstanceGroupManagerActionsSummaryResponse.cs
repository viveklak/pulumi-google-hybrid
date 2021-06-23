// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.V1.Outputs
{

    [OutputType]
    public sealed class InstanceGroupManagerActionsSummaryResponse
    {
        /// <summary>
        /// [Output Only] The total number of instances in the managed instance group that are scheduled to be abandoned. Abandoning an instance removes it from the managed instance group without deleting it.
        /// </summary>
        public readonly int Abandoning;
        /// <summary>
        /// [Output Only] The number of instances in the managed instance group that are scheduled to be created or are currently being created. If the group fails to create any of these instances, it tries again until it creates the instance successfully. If you have disabled creation retries, this field will not be populated; instead, the creatingWithoutRetries field will be populated.
        /// </summary>
        public readonly int Creating;
        /// <summary>
        /// [Output Only] The number of instances that the managed instance group will attempt to create. The group attempts to create each instance only once. If the group fails to create any of these instances, it decreases the group's targetSize value accordingly.
        /// </summary>
        public readonly int CreatingWithoutRetries;
        /// <summary>
        /// [Output Only] The number of instances in the managed instance group that are scheduled to be deleted or are currently being deleted.
        /// </summary>
        public readonly int Deleting;
        /// <summary>
        /// [Output Only] The number of instances in the managed instance group that are running and have no scheduled actions.
        /// </summary>
        public readonly int None;
        /// <summary>
        /// [Output Only] The number of instances in the managed instance group that are scheduled to be recreated or are currently being being recreated. Recreating an instance deletes the existing root persistent disk and creates a new disk from the image that is defined in the instance template.
        /// </summary>
        public readonly int Recreating;
        /// <summary>
        /// [Output Only] The number of instances in the managed instance group that are being reconfigured with properties that do not require a restart or a recreate action. For example, setting or removing target pools for the instance.
        /// </summary>
        public readonly int Refreshing;
        /// <summary>
        /// [Output Only] The number of instances in the managed instance group that are scheduled to be restarted or are currently being restarted.
        /// </summary>
        public readonly int Restarting;
        /// <summary>
        /// [Output Only] The number of instances in the managed instance group that are being verified. See the managedInstances[].currentAction property in the listManagedInstances method documentation.
        /// </summary>
        public readonly int Verifying;

        [OutputConstructor]
        private InstanceGroupManagerActionsSummaryResponse(
            int abandoning,

            int creating,

            int creatingWithoutRetries,

            int deleting,

            int none,

            int recreating,

            int refreshing,

            int restarting,

            int verifying)
        {
            Abandoning = abandoning;
            Creating = creating;
            CreatingWithoutRetries = creatingWithoutRetries;
            Deleting = deleting;
            None = none;
            Recreating = recreating;
            Refreshing = refreshing;
            Restarting = restarting;
            Verifying = verifying;
        }
    }
}
