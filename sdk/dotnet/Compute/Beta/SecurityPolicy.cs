// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Beta
{
    /// <summary>
    /// Creates a new policy in the specified project using the data included in the request.
    /// </summary>
    [GoogleNativeResourceType("google-native:compute/beta:SecurityPolicy")]
    public partial class SecurityPolicy : Pulumi.CustomResource
    {
        [Output("adaptiveProtectionConfig")]
        public Output<Outputs.SecurityPolicyAdaptiveProtectionConfigResponse> AdaptiveProtectionConfig { get; private set; } = null!;

        [Output("advancedOptionsConfig")]
        public Output<Outputs.SecurityPolicyAdvancedOptionsConfigResponse> AdvancedOptionsConfig { get; private set; } = null!;

        /// <summary>
        /// A list of associations that belong to this policy.
        /// </summary>
        [Output("associations")]
        public Output<ImmutableArray<Outputs.SecurityPolicyAssociationResponse>> Associations { get; private set; } = null!;

        /// <summary>
        /// [Output Only] Creation timestamp in RFC3339 text format.
        /// </summary>
        [Output("creationTimestamp")]
        public Output<string> CreationTimestamp { get; private set; } = null!;

        /// <summary>
        /// An optional description of this resource. Provide this property when you create the resource.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// User-provided name of the Organization security plicy. The name should be unique in the organization in which the security policy is created. This should only be used when SecurityPolicyType is FIREWALL. The name must be 1-63 characters long, and comply with https://www.ietf.org/rfc/rfc1035.txt. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        /// </summary>
        [Output("displayName")]
        public Output<string> DisplayName { get; private set; } = null!;

        /// <summary>
        /// Specifies a fingerprint for this resource, which is essentially a hash of the metadata's contents and used for optimistic locking. The fingerprint is initially generated by Compute Engine and changes after every request to modify or update metadata. You must always provide an up-to-date fingerprint hash in order to update or change metadata, otherwise the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make get() request to the security policy.
        /// </summary>
        [Output("fingerprint")]
        public Output<string> Fingerprint { get; private set; } = null!;

        /// <summary>
        /// [Output only] Type of the resource. Always compute#securityPolicyfor security policies
        /// </summary>
        [Output("kind")]
        public Output<string> Kind { get; private set; } = null!;

        /// <summary>
        /// A fingerprint for the labels being applied to this security policy, which is essentially a hash of the labels set used for optimistic locking. The fingerprint is initially generated by Compute Engine and changes after every request to modify or update labels. You must always provide an up-to-date fingerprint hash in order to update or change labels. To see the latest fingerprint, make get() request to the security policy.
        /// </summary>
        [Output("labelFingerprint")]
        public Output<string> LabelFingerprint { get; private set; } = null!;

        /// <summary>
        /// Labels for this resource. These can only be added or modified by the setLabels method. Each label key/value pair must comply with RFC1035. Label values may be empty.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>> Labels { get; private set; } = null!;

        /// <summary>
        /// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// [Output Only] The parent of the security policy.
        /// </summary>
        [Output("parent")]
        public Output<string> Parent { get; private set; } = null!;

        /// <summary>
        /// [Output Only] Total count of all security policy rule tuples. A security policy can not exceed a set number of tuples.
        /// </summary>
        [Output("ruleTupleCount")]
        public Output<int> RuleTupleCount { get; private set; } = null!;

        /// <summary>
        /// A list of rules that belong to this policy. There must always be a default rule (rule with priority 2147483647 and match "*"). If no rules are provided when creating a security policy, a default rule with action "allow" will be added.
        /// </summary>
        [Output("rules")]
        public Output<ImmutableArray<Outputs.SecurityPolicyRuleResponse>> Rules { get; private set; } = null!;

        /// <summary>
        /// [Output Only] Server-defined URL for the resource.
        /// </summary>
        [Output("selfLink")]
        public Output<string> SelfLink { get; private set; } = null!;

        /// <summary>
        /// [Output Only] Server-defined URL for this resource with the resource id.
        /// </summary>
        [Output("selfLinkWithId")]
        public Output<string> SelfLinkWithId { get; private set; } = null!;

        /// <summary>
        /// The type indicates the intended use of the security policy. CLOUD_ARMOR policies apply to backend services. FIREWALL policies apply to organizations.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a SecurityPolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SecurityPolicy(string name, SecurityPolicyArgs args, CustomResourceOptions? options = null)
            : base("google-native:compute/beta:SecurityPolicy", name, args ?? new SecurityPolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SecurityPolicy(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:compute/beta:SecurityPolicy", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing SecurityPolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SecurityPolicy Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new SecurityPolicy(name, id, options);
        }
    }

    public sealed class SecurityPolicyArgs : Pulumi.ResourceArgs
    {
        [Input("adaptiveProtectionConfig")]
        public Input<Inputs.SecurityPolicyAdaptiveProtectionConfigArgs>? AdaptiveProtectionConfig { get; set; }

        [Input("advancedOptionsConfig")]
        public Input<Inputs.SecurityPolicyAdvancedOptionsConfigArgs>? AdvancedOptionsConfig { get; set; }

        [Input("associations")]
        private InputList<Inputs.SecurityPolicyAssociationArgs>? _associations;

        /// <summary>
        /// A list of associations that belong to this policy.
        /// </summary>
        public InputList<Inputs.SecurityPolicyAssociationArgs> Associations
        {
            get => _associations ?? (_associations = new InputList<Inputs.SecurityPolicyAssociationArgs>());
            set => _associations = value;
        }

        /// <summary>
        /// [Output Only] Creation timestamp in RFC3339 text format.
        /// </summary>
        [Input("creationTimestamp")]
        public Input<string>? CreationTimestamp { get; set; }

        /// <summary>
        /// An optional description of this resource. Provide this property when you create the resource.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// User-provided name of the Organization security plicy. The name should be unique in the organization in which the security policy is created. This should only be used when SecurityPolicyType is FIREWALL. The name must be 1-63 characters long, and comply with https://www.ietf.org/rfc/rfc1035.txt. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// Specifies a fingerprint for this resource, which is essentially a hash of the metadata's contents and used for optimistic locking. The fingerprint is initially generated by Compute Engine and changes after every request to modify or update metadata. You must always provide an up-to-date fingerprint hash in order to update or change metadata, otherwise the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make get() request to the security policy.
        /// </summary>
        [Input("fingerprint")]
        public Input<string>? Fingerprint { get; set; }

        /// <summary>
        /// [Output Only] The unique identifier for the resource. This identifier is defined by the server.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// [Output only] Type of the resource. Always compute#securityPolicyfor security policies
        /// </summary>
        [Input("kind")]
        public Input<string>? Kind { get; set; }

        /// <summary>
        /// A fingerprint for the labels being applied to this security policy, which is essentially a hash of the labels set used for optimistic locking. The fingerprint is initially generated by Compute Engine and changes after every request to modify or update labels. You must always provide an up-to-date fingerprint hash in order to update or change labels. To see the latest fingerprint, make get() request to the security policy.
        /// </summary>
        [Input("labelFingerprint")]
        public Input<string>? LabelFingerprint { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// Labels for this resource. These can only be added or modified by the setLabels method. Each label key/value pair must comply with RFC1035. Label values may be empty.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// [Output Only] The parent of the security policy.
        /// </summary>
        [Input("parent")]
        public Input<string>? Parent { get; set; }

        [Input("project", required: true)]
        public Input<string> Project { get; set; } = null!;

        [Input("requestId")]
        public Input<string>? RequestId { get; set; }

        /// <summary>
        /// [Output Only] Total count of all security policy rule tuples. A security policy can not exceed a set number of tuples.
        /// </summary>
        [Input("ruleTupleCount")]
        public Input<int>? RuleTupleCount { get; set; }

        [Input("rules")]
        private InputList<Inputs.SecurityPolicyRuleArgs>? _rules;

        /// <summary>
        /// A list of rules that belong to this policy. There must always be a default rule (rule with priority 2147483647 and match "*"). If no rules are provided when creating a security policy, a default rule with action "allow" will be added.
        /// </summary>
        public InputList<Inputs.SecurityPolicyRuleArgs> Rules
        {
            get => _rules ?? (_rules = new InputList<Inputs.SecurityPolicyRuleArgs>());
            set => _rules = value;
        }

        /// <summary>
        /// [Output Only] Server-defined URL for the resource.
        /// </summary>
        [Input("selfLink")]
        public Input<string>? SelfLink { get; set; }

        /// <summary>
        /// [Output Only] Server-defined URL for this resource with the resource id.
        /// </summary>
        [Input("selfLinkWithId")]
        public Input<string>? SelfLinkWithId { get; set; }

        /// <summary>
        /// The type indicates the intended use of the security policy. CLOUD_ARMOR policies apply to backend services. FIREWALL policies apply to organizations.
        /// </summary>
        [Input("type")]
        public Input<Pulumi.GoogleNative.Compute.Beta.SecurityPolicyType>? Type { get; set; }

        [Input("validateOnly")]
        public Input<string>? ValidateOnly { get; set; }

        public SecurityPolicyArgs()
        {
        }
    }
}
