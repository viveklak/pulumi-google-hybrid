// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Alpha
{
    /// <summary>
    /// Creates a new resource policy.
    /// </summary>
    [GoogleNativeResourceType("google-native:compute/alpha:ResourcePolicy")]
    public partial class ResourcePolicy : Pulumi.CustomResource
    {
        /// <summary>
        /// [Output Only] Creation timestamp in RFC3339 text format.
        /// </summary>
        [Output("creationTimestamp")]
        public Output<string> CreationTimestamp { get; private set; } = null!;

        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// Resource policy for instances for placement configuration.
        /// </summary>
        [Output("groupPlacementPolicy")]
        public Output<Outputs.ResourcePolicyGroupPlacementPolicyResponse> GroupPlacementPolicy { get; private set; } = null!;

        /// <summary>
        /// Resource policy for scheduling instance operations.
        /// </summary>
        [Output("instanceSchedulePolicy")]
        public Output<Outputs.ResourcePolicyInstanceSchedulePolicyResponse> InstanceSchedulePolicy { get; private set; } = null!;

        /// <summary>
        /// [Output Only] Type of the resource. Always compute#resource_policies for resource policies.
        /// </summary>
        [Output("kind")]
        public Output<string> Kind { get; private set; } = null!;

        /// <summary>
        /// The name of the resource, provided by the client when initially creating the resource. The resource name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// [Output Only] The system status of the resource policy.
        /// </summary>
        [Output("resourceStatus")]
        public Output<Outputs.ResourcePolicyResourceStatusResponse> ResourceStatus { get; private set; } = null!;

        /// <summary>
        /// [Output Only] Server-defined fully-qualified URL for this resource.
        /// </summary>
        [Output("selfLink")]
        public Output<string> SelfLink { get; private set; } = null!;

        /// <summary>
        /// [Output Only] Server-defined URL for this resource with the resource id.
        /// </summary>
        [Output("selfLinkWithId")]
        public Output<string> SelfLinkWithId { get; private set; } = null!;

        /// <summary>
        /// Resource policy for persistent disks for creating snapshots.
        /// </summary>
        [Output("snapshotSchedulePolicy")]
        public Output<Outputs.ResourcePolicySnapshotSchedulePolicyResponse> SnapshotSchedulePolicy { get; private set; } = null!;

        /// <summary>
        /// [Output Only] The status of resource policy creation.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// Resource policy applicable to VMs for infrastructure maintenance.
        /// </summary>
        [Output("vmMaintenancePolicy")]
        public Output<Outputs.ResourcePolicyVmMaintenancePolicyResponse> VmMaintenancePolicy { get; private set; } = null!;


        /// <summary>
        /// Create a ResourcePolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ResourcePolicy(string name, ResourcePolicyArgs args, CustomResourceOptions? options = null)
            : base("google-native:compute/alpha:ResourcePolicy", name, args ?? new ResourcePolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ResourcePolicy(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:compute/alpha:ResourcePolicy", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ResourcePolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ResourcePolicy Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ResourcePolicy(name, id, options);
        }
    }

    public sealed class ResourcePolicyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// [Output Only] Creation timestamp in RFC3339 text format.
        /// </summary>
        [Input("creationTimestamp")]
        public Input<string>? CreationTimestamp { get; set; }

        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Resource policy for instances for placement configuration.
        /// </summary>
        [Input("groupPlacementPolicy")]
        public Input<Inputs.ResourcePolicyGroupPlacementPolicyArgs>? GroupPlacementPolicy { get; set; }

        /// <summary>
        /// [Output Only] The unique identifier for the resource. This identifier is defined by the server.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// Resource policy for scheduling instance operations.
        /// </summary>
        [Input("instanceSchedulePolicy")]
        public Input<Inputs.ResourcePolicyInstanceSchedulePolicyArgs>? InstanceSchedulePolicy { get; set; }

        /// <summary>
        /// [Output Only] Type of the resource. Always compute#resource_policies for resource policies.
        /// </summary>
        [Input("kind")]
        public Input<string>? Kind { get; set; }

        /// <summary>
        /// The name of the resource, provided by the client when initially creating the resource. The resource name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("project", required: true)]
        public Input<string> Project { get; set; } = null!;

        [Input("region", required: true)]
        public Input<string> Region { get; set; } = null!;

        [Input("requestId")]
        public Input<string>? RequestId { get; set; }

        /// <summary>
        /// [Output Only] The system status of the resource policy.
        /// </summary>
        [Input("resourceStatus")]
        public Input<Inputs.ResourcePolicyResourceStatusArgs>? ResourceStatus { get; set; }

        /// <summary>
        /// [Output Only] Server-defined fully-qualified URL for this resource.
        /// </summary>
        [Input("selfLink")]
        public Input<string>? SelfLink { get; set; }

        /// <summary>
        /// [Output Only] Server-defined URL for this resource with the resource id.
        /// </summary>
        [Input("selfLinkWithId")]
        public Input<string>? SelfLinkWithId { get; set; }

        /// <summary>
        /// Resource policy for persistent disks for creating snapshots.
        /// </summary>
        [Input("snapshotSchedulePolicy")]
        public Input<Inputs.ResourcePolicySnapshotSchedulePolicyArgs>? SnapshotSchedulePolicy { get; set; }

        /// <summary>
        /// [Output Only] The status of resource policy creation.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Resource policy applicable to VMs for infrastructure maintenance.
        /// </summary>
        [Input("vmMaintenancePolicy")]
        public Input<Inputs.ResourcePolicyVmMaintenancePolicyArgs>? VmMaintenancePolicy { get; set; }

        public ResourcePolicyArgs()
        {
        }
    }
}
