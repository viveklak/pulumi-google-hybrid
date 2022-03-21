// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.V1
{
    /// <summary>
    /// Creates a managed instance group using the information that you specify in the request. After the group is created, instances in the group are created using the specified instance template. This operation is marked as DONE when the group is created even if the instances in the group have not yet been created. You must separately verify the status of the individual instances with the listmanagedinstances method. A managed instance group can have up to 1000 VM instances per group. Please contact Cloud Support if you need an increase in this limit.
    /// </summary>
    [GoogleNativeResourceType("google-native:compute/v1:InstanceGroupManager")]
    public partial class InstanceGroupManager : Pulumi.CustomResource
    {
        /// <summary>
        /// The autohealing policy for this managed instance group. You can specify only one value.
        /// </summary>
        [Output("autoHealingPolicies")]
        public Output<ImmutableArray<Outputs.InstanceGroupManagerAutoHealingPolicyResponse>> AutoHealingPolicies { get; private set; } = null!;

        /// <summary>
        /// The base instance name to use for instances in this group. The value must be 1-58 characters long. Instances are named by appending a hyphen and a random four-character string to the base instance name. The base instance name must comply with RFC1035.
        /// </summary>
        [Output("baseInstanceName")]
        public Output<string> BaseInstanceName { get; private set; } = null!;

        /// <summary>
        /// The creation timestamp for this managed instance group in RFC3339 text format.
        /// </summary>
        [Output("creationTimestamp")]
        public Output<string> CreationTimestamp { get; private set; } = null!;

        /// <summary>
        /// The list of instance actions and the number of instances in this managed instance group that are scheduled for each of those actions.
        /// </summary>
        [Output("currentActions")]
        public Output<Outputs.InstanceGroupManagerActionsSummaryResponse> CurrentActions { get; private set; } = null!;

        /// <summary>
        /// An optional description of this resource.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// Policy specifying the intended distribution of managed instances across zones in a regional managed instance group.
        /// </summary>
        [Output("distributionPolicy")]
        public Output<Outputs.DistributionPolicyResponse> DistributionPolicy { get; private set; } = null!;

        /// <summary>
        /// Fingerprint of this resource. This field may be used in optimistic locking. It will be ignored when inserting an InstanceGroupManager. An up-to-date fingerprint must be provided in order to update the InstanceGroupManager, otherwise the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make a get() request to retrieve an InstanceGroupManager.
        /// </summary>
        [Output("fingerprint")]
        public Output<string> Fingerprint { get; private set; } = null!;

        /// <summary>
        /// The URL of the Instance Group resource.
        /// </summary>
        [Output("instanceGroup")]
        public Output<string> InstanceGroup { get; private set; } = null!;

        /// <summary>
        /// The URL of the instance template that is specified for this managed instance group. The group uses this template to create all new instances in the managed instance group. The templates for existing instances in the group do not change unless you run recreateInstances, run applyUpdatesToInstances, or set the group's updatePolicy.type to PROACTIVE.
        /// </summary>
        [Output("instanceTemplate")]
        public Output<string> InstanceTemplate { get; private set; } = null!;

        /// <summary>
        /// The resource type, which is always compute#instanceGroupManager for managed instance groups.
        /// </summary>
        [Output("kind")]
        public Output<string> Kind { get; private set; } = null!;

        /// <summary>
        /// The name of the managed instance group. The name must be 1-63 characters long, and comply with RFC1035.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Named ports configured for the Instance Groups complementary to this Instance Group Manager.
        /// </summary>
        [Output("namedPorts")]
        public Output<ImmutableArray<Outputs.NamedPortResponse>> NamedPorts { get; private set; } = null!;

        /// <summary>
        /// The URL of the region where the managed instance group resides (for regional resources).
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// The URL for this managed instance group. The server defines this URL.
        /// </summary>
        [Output("selfLink")]
        public Output<string> SelfLink { get; private set; } = null!;

        /// <summary>
        /// Stateful configuration for this Instanced Group Manager
        /// </summary>
        [Output("statefulPolicy")]
        public Output<Outputs.StatefulPolicyResponse> StatefulPolicy { get; private set; } = null!;

        /// <summary>
        /// The status of this managed instance group.
        /// </summary>
        [Output("status")]
        public Output<Outputs.InstanceGroupManagerStatusResponse> Status { get; private set; } = null!;

        /// <summary>
        /// The URLs for all TargetPool resources to which instances in the instanceGroup field are added. The target pools automatically apply to all of the instances in the managed instance group.
        /// </summary>
        [Output("targetPools")]
        public Output<ImmutableArray<string>> TargetPools { get; private set; } = null!;

        /// <summary>
        /// The target number of running instances for this managed instance group. You can reduce this number by using the instanceGroupManager deleteInstances or abandonInstances methods. Resizing the group also changes this number.
        /// </summary>
        [Output("targetSize")]
        public Output<int> TargetSize { get; private set; } = null!;

        /// <summary>
        /// The update policy for this managed instance group.
        /// </summary>
        [Output("updatePolicy")]
        public Output<Outputs.InstanceGroupManagerUpdatePolicyResponse> UpdatePolicy { get; private set; } = null!;

        /// <summary>
        /// Specifies the instance templates used by this managed instance group to create instances. Each version is defined by an instanceTemplate and a name. Every version can appear at most once per instance group. This field overrides the top-level instanceTemplate field. Read more about the relationships between these fields. Exactly one version must leave the targetSize field unset. That version will be applied to all remaining instances. For more information, read about canary updates.
        /// </summary>
        [Output("versions")]
        public Output<ImmutableArray<Outputs.InstanceGroupManagerVersionResponse>> Versions { get; private set; } = null!;

        /// <summary>
        /// The URL of a zone where the managed instance group is located (for zonal resources).
        /// </summary>
        [Output("zone")]
        public Output<string> Zone { get; private set; } = null!;


        /// <summary>
        /// Create a InstanceGroupManager resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public InstanceGroupManager(string name, InstanceGroupManagerArgs? args = null, CustomResourceOptions? options = null)
            : base("google-native:compute/v1:InstanceGroupManager", name, args ?? new InstanceGroupManagerArgs(), MakeResourceOptions(options, ""))
        {
        }

        private InstanceGroupManager(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:compute/v1:InstanceGroupManager", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing InstanceGroupManager resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static InstanceGroupManager Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new InstanceGroupManager(name, id, options);
        }
    }

    public sealed class InstanceGroupManagerArgs : Pulumi.ResourceArgs
    {
        [Input("autoHealingPolicies")]
        private InputList<Inputs.InstanceGroupManagerAutoHealingPolicyArgs>? _autoHealingPolicies;

        /// <summary>
        /// The autohealing policy for this managed instance group. You can specify only one value.
        /// </summary>
        public InputList<Inputs.InstanceGroupManagerAutoHealingPolicyArgs> AutoHealingPolicies
        {
            get => _autoHealingPolicies ?? (_autoHealingPolicies = new InputList<Inputs.InstanceGroupManagerAutoHealingPolicyArgs>());
            set => _autoHealingPolicies = value;
        }

        /// <summary>
        /// The base instance name to use for instances in this group. The value must be 1-58 characters long. Instances are named by appending a hyphen and a random four-character string to the base instance name. The base instance name must comply with RFC1035.
        /// </summary>
        [Input("baseInstanceName")]
        public Input<string>? BaseInstanceName { get; set; }

        /// <summary>
        /// An optional description of this resource.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Policy specifying the intended distribution of managed instances across zones in a regional managed instance group.
        /// </summary>
        [Input("distributionPolicy")]
        public Input<Inputs.DistributionPolicyArgs>? DistributionPolicy { get; set; }

        /// <summary>
        /// The URL of the instance template that is specified for this managed instance group. The group uses this template to create all new instances in the managed instance group. The templates for existing instances in the group do not change unless you run recreateInstances, run applyUpdatesToInstances, or set the group's updatePolicy.type to PROACTIVE.
        /// </summary>
        [Input("instanceTemplate")]
        public Input<string>? InstanceTemplate { get; set; }

        /// <summary>
        /// The name of the managed instance group. The name must be 1-63 characters long, and comply with RFC1035.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("namedPorts")]
        private InputList<Inputs.NamedPortArgs>? _namedPorts;

        /// <summary>
        /// Named ports configured for the Instance Groups complementary to this Instance Group Manager.
        /// </summary>
        public InputList<Inputs.NamedPortArgs> NamedPorts
        {
            get => _namedPorts ?? (_namedPorts = new InputList<Inputs.NamedPortArgs>());
            set => _namedPorts = value;
        }

        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
        /// </summary>
        [Input("requestId")]
        public Input<string>? RequestId { get; set; }

        /// <summary>
        /// Stateful configuration for this Instanced Group Manager
        /// </summary>
        [Input("statefulPolicy")]
        public Input<Inputs.StatefulPolicyArgs>? StatefulPolicy { get; set; }

        [Input("targetPools")]
        private InputList<string>? _targetPools;

        /// <summary>
        /// The URLs for all TargetPool resources to which instances in the instanceGroup field are added. The target pools automatically apply to all of the instances in the managed instance group.
        /// </summary>
        public InputList<string> TargetPools
        {
            get => _targetPools ?? (_targetPools = new InputList<string>());
            set => _targetPools = value;
        }

        /// <summary>
        /// The target number of running instances for this managed instance group. You can reduce this number by using the instanceGroupManager deleteInstances or abandonInstances methods. Resizing the group also changes this number.
        /// </summary>
        [Input("targetSize")]
        public Input<int>? TargetSize { get; set; }

        /// <summary>
        /// The update policy for this managed instance group.
        /// </summary>
        [Input("updatePolicy")]
        public Input<Inputs.InstanceGroupManagerUpdatePolicyArgs>? UpdatePolicy { get; set; }

        [Input("versions")]
        private InputList<Inputs.InstanceGroupManagerVersionArgs>? _versions;

        /// <summary>
        /// Specifies the instance templates used by this managed instance group to create instances. Each version is defined by an instanceTemplate and a name. Every version can appear at most once per instance group. This field overrides the top-level instanceTemplate field. Read more about the relationships between these fields. Exactly one version must leave the targetSize field unset. That version will be applied to all remaining instances. For more information, read about canary updates.
        /// </summary>
        public InputList<Inputs.InstanceGroupManagerVersionArgs> Versions
        {
            get => _versions ?? (_versions = new InputList<Inputs.InstanceGroupManagerVersionArgs>());
            set => _versions = value;
        }

        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public InstanceGroupManagerArgs()
        {
        }
    }
}
