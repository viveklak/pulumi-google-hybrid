// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Beta
{
    public static class GetResourcePolicy
    {
        /// <summary>
        /// Retrieves all information of the specified resource policy.
        /// </summary>
        public static Task<GetResourcePolicyResult> InvokeAsync(GetResourcePolicyArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetResourcePolicyResult>("google-native:compute/beta:getResourcePolicy", args ?? new GetResourcePolicyArgs(), options.WithVersion());
    }


    public sealed class GetResourcePolicyArgs : Pulumi.InvokeArgs
    {
        [Input("project", required: true)]
        public string Project { get; set; } = null!;

        [Input("region", required: true)]
        public string Region { get; set; } = null!;

        [Input("resourcePolicy", required: true)]
        public string ResourcePolicy { get; set; } = null!;

        public GetResourcePolicyArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetResourcePolicyResult
    {
        /// <summary>
        /// Creation timestamp in RFC3339 text format.
        /// </summary>
        public readonly string CreationTimestamp;
        public readonly string Description;
        /// <summary>
        /// Resource policy for instances for placement configuration.
        /// </summary>
        public readonly Outputs.ResourcePolicyGroupPlacementPolicyResponse GroupPlacementPolicy;
        /// <summary>
        /// Resource policy for scheduling instance operations.
        /// </summary>
        public readonly Outputs.ResourcePolicyInstanceSchedulePolicyResponse InstanceSchedulePolicy;
        /// <summary>
        /// Type of the resource. Always compute#resource_policies for resource policies.
        /// </summary>
        public readonly string Kind;
        /// <summary>
        /// The name of the resource, provided by the client when initially creating the resource. The resource name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        /// </summary>
        public readonly string Name;
        public readonly string Region;
        /// <summary>
        /// The system status of the resource policy.
        /// </summary>
        public readonly Outputs.ResourcePolicyResourceStatusResponse ResourceStatus;
        /// <summary>
        /// Server-defined fully-qualified URL for this resource.
        /// </summary>
        public readonly string SelfLink;
        /// <summary>
        /// Resource policy for persistent disks for creating snapshots.
        /// </summary>
        public readonly Outputs.ResourcePolicySnapshotSchedulePolicyResponse SnapshotSchedulePolicy;
        /// <summary>
        /// The status of resource policy creation.
        /// </summary>
        public readonly string Status;

        [OutputConstructor]
        private GetResourcePolicyResult(
            string creationTimestamp,

            string description,

            Outputs.ResourcePolicyGroupPlacementPolicyResponse groupPlacementPolicy,

            Outputs.ResourcePolicyInstanceSchedulePolicyResponse instanceSchedulePolicy,

            string kind,

            string name,

            string region,

            Outputs.ResourcePolicyResourceStatusResponse resourceStatus,

            string selfLink,

            Outputs.ResourcePolicySnapshotSchedulePolicyResponse snapshotSchedulePolicy,

            string status)
        {
            CreationTimestamp = creationTimestamp;
            Description = description;
            GroupPlacementPolicy = groupPlacementPolicy;
            InstanceSchedulePolicy = instanceSchedulePolicy;
            Kind = kind;
            Name = name;
            Region = region;
            ResourceStatus = resourceStatus;
            SelfLink = selfLink;
            SnapshotSchedulePolicy = snapshotSchedulePolicy;
            Status = status;
        }
    }
}
