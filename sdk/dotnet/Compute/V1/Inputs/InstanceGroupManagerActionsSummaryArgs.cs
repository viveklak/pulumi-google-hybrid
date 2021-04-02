// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleCloud.Compute.V1.Inputs
{

    public sealed class InstanceGroupManagerActionsSummaryArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// [Output Only] The total number of instances in the managed instance group that are scheduled to be abandoned. Abandoning an instance removes it from the managed instance group without deleting it.
        /// </summary>
        [Input("abandoning")]
        public Input<int>? Abandoning { get; set; }

        /// <summary>
        /// [Output Only] The number of instances in the managed instance group that are scheduled to be created or are currently being created. If the group fails to create any of these instances, it tries again until it creates the instance successfully.
        /// 
        /// If you have disabled creation retries, this field will not be populated; instead, the creatingWithoutRetries field will be populated.
        /// </summary>
        [Input("creating")]
        public Input<int>? Creating { get; set; }

        /// <summary>
        /// [Output Only] The number of instances that the managed instance group will attempt to create. The group attempts to create each instance only once. If the group fails to create any of these instances, it decreases the group's targetSize value accordingly.
        /// </summary>
        [Input("creatingWithoutRetries")]
        public Input<int>? CreatingWithoutRetries { get; set; }

        /// <summary>
        /// [Output Only] The number of instances in the managed instance group that are scheduled to be deleted or are currently being deleted.
        /// </summary>
        [Input("deleting")]
        public Input<int>? Deleting { get; set; }

        /// <summary>
        /// [Output Only] The number of instances in the managed instance group that are running and have no scheduled actions.
        /// </summary>
        [Input("none")]
        public Input<int>? None { get; set; }

        /// <summary>
        /// [Output Only] The number of instances in the managed instance group that are scheduled to be recreated or are currently being being recreated. Recreating an instance deletes the existing root persistent disk and creates a new disk from the image that is defined in the instance template.
        /// </summary>
        [Input("recreating")]
        public Input<int>? Recreating { get; set; }

        /// <summary>
        /// [Output Only] The number of instances in the managed instance group that are being reconfigured with properties that do not require a restart or a recreate action. For example, setting or removing target pools for the instance.
        /// </summary>
        [Input("refreshing")]
        public Input<int>? Refreshing { get; set; }

        /// <summary>
        /// [Output Only] The number of instances in the managed instance group that are scheduled to be restarted or are currently being restarted.
        /// </summary>
        [Input("restarting")]
        public Input<int>? Restarting { get; set; }

        /// <summary>
        /// [Output Only] The number of instances in the managed instance group that are being verified. See the managedInstances[].currentAction property in the listManagedInstances method documentation.
        /// </summary>
        [Input("verifying")]
        public Input<int>? Verifying { get; set; }

        public InstanceGroupManagerActionsSummaryArgs()
        {
        }
    }
}