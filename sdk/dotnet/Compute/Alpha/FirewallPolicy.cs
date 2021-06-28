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
    /// Creates a new policy in the specified project using the data included in the request.
    /// </summary>
    [GoogleNativeResourceType("google-native:compute/alpha:FirewallPolicy")]
    public partial class FirewallPolicy : Pulumi.CustomResource
    {
        /// <summary>
        /// A list of associations that belong to this firewall policy.
        /// </summary>
        [Output("associations")]
        public Output<ImmutableArray<Outputs.FirewallPolicyAssociationResponse>> Associations { get; private set; } = null!;

        /// <summary>
        /// Creation timestamp in RFC3339 text format.
        /// </summary>
        [Output("creationTimestamp")]
        public Output<string> CreationTimestamp { get; private set; } = null!;

        /// <summary>
        /// An optional description of this resource. Provide this property when you create the resource.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// Depreacted, please use short name instead. User-provided name of the Organization firewall plicy. The name should be unique in the organization in which the firewall policy is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        /// </summary>
        [Output("displayName")]
        public Output<string> DisplayName { get; private set; } = null!;

        /// <summary>
        /// Specifies a fingerprint for this resource, which is essentially a hash of the metadata's contents and used for optimistic locking. The fingerprint is initially generated by Compute Engine and changes after every request to modify or update metadata. You must always provide an up-to-date fingerprint hash in order to update or change metadata, otherwise the request will fail with error 412 conditionNotMet.
        /// 
        /// To see the latest fingerprint, make get() request to the firewall policy.
        /// </summary>
        [Output("fingerprint")]
        public Output<string> Fingerprint { get; private set; } = null!;

        /// <summary>
        /// [Output only] Type of the resource. Always compute#firewallPolicyfor firewall policies
        /// </summary>
        [Output("kind")]
        public Output<string> Kind { get; private set; } = null!;

        /// <summary>
        /// Name of the resource. It is a numeric ID allocated by GCP which uniquely identifies the Firewall Policy.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The parent of the firewall policy.
        /// </summary>
        [Output("parent")]
        public Output<string> Parent { get; private set; } = null!;

        /// <summary>
        /// URL of the region where the regional firewall policy resides. This field is not applicable to global firewall policies. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// Total count of all firewall policy rule tuples. A firewall policy can not exceed a set number of tuples.
        /// </summary>
        [Output("ruleTupleCount")]
        public Output<int> RuleTupleCount { get; private set; } = null!;

        /// <summary>
        /// A list of rules that belong to this policy. There must always be a default rule (rule with priority 2147483647 and match "*"). If no rules are provided when creating a firewall policy, a default rule with action "allow" will be added.
        /// </summary>
        [Output("rules")]
        public Output<ImmutableArray<Outputs.FirewallPolicyRuleResponse>> Rules { get; private set; } = null!;

        /// <summary>
        /// Server-defined URL for the resource.
        /// </summary>
        [Output("selfLink")]
        public Output<string> SelfLink { get; private set; } = null!;

        /// <summary>
        /// Server-defined URL for this resource with the resource id.
        /// </summary>
        [Output("selfLinkWithId")]
        public Output<string> SelfLinkWithId { get; private set; } = null!;

        /// <summary>
        /// User-provided name of the Organization firewall plicy. The name should be unique in the organization in which the firewall policy is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        /// </summary>
        [Output("shortName")]
        public Output<string> ShortName { get; private set; } = null!;


        /// <summary>
        /// Create a FirewallPolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public FirewallPolicy(string name, FirewallPolicyArgs? args = null, CustomResourceOptions? options = null)
            : base("google-native:compute/alpha:FirewallPolicy", name, args ?? new FirewallPolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private FirewallPolicy(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:compute/alpha:FirewallPolicy", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing FirewallPolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static FirewallPolicy Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new FirewallPolicy(name, id, options);
        }
    }

    public sealed class FirewallPolicyArgs : Pulumi.ResourceArgs
    {
        [Input("associations")]
        private InputList<Inputs.FirewallPolicyAssociationArgs>? _associations;

        /// <summary>
        /// A list of associations that belong to this firewall policy.
        /// </summary>
        public InputList<Inputs.FirewallPolicyAssociationArgs> Associations
        {
            get => _associations ?? (_associations = new InputList<Inputs.FirewallPolicyAssociationArgs>());
            set => _associations = value;
        }

        /// <summary>
        /// An optional description of this resource. Provide this property when you create the resource.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Depreacted, please use short name instead. User-provided name of the Organization firewall plicy. The name should be unique in the organization in which the firewall policy is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        [Input("parentId")]
        public Input<string>? ParentId { get; set; }

        [Input("requestId")]
        public Input<string>? RequestId { get; set; }

        [Input("rules")]
        private InputList<Inputs.FirewallPolicyRuleArgs>? _rules;

        /// <summary>
        /// A list of rules that belong to this policy. There must always be a default rule (rule with priority 2147483647 and match "*"). If no rules are provided when creating a firewall policy, a default rule with action "allow" will be added.
        /// </summary>
        public InputList<Inputs.FirewallPolicyRuleArgs> Rules
        {
            get => _rules ?? (_rules = new InputList<Inputs.FirewallPolicyRuleArgs>());
            set => _rules = value;
        }

        /// <summary>
        /// User-provided name of the Organization firewall plicy. The name should be unique in the organization in which the firewall policy is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        /// </summary>
        [Input("shortName")]
        public Input<string>? ShortName { get; set; }

        public FirewallPolicyArgs()
        {
        }
    }
}
