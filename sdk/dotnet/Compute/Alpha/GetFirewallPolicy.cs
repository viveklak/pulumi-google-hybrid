// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi.Utilities;

namespace Pulumi.GoogleNative.Compute.Alpha
{
    public static class GetFirewallPolicy
    {
        /// <summary>
        /// Returns the specified firewall policy.
        /// </summary>
        public static Task<GetFirewallPolicyResult> InvokeAsync(GetFirewallPolicyArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetFirewallPolicyResult>("google-native:compute/alpha:getFirewallPolicy", args ?? new GetFirewallPolicyArgs(), options.WithDefaults());

        /// <summary>
        /// Returns the specified firewall policy.
        /// </summary>
        public static Output<GetFirewallPolicyResult> Invoke(GetFirewallPolicyInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetFirewallPolicyResult>("google-native:compute/alpha:getFirewallPolicy", args ?? new GetFirewallPolicyInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetFirewallPolicyArgs : Pulumi.InvokeArgs
    {
        [Input("firewallPolicy", required: true)]
        public string FirewallPolicy { get; set; } = null!;

        public GetFirewallPolicyArgs()
        {
        }
    }

    public sealed class GetFirewallPolicyInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("firewallPolicy", required: true)]
        public Input<string> FirewallPolicy { get; set; } = null!;

        public GetFirewallPolicyInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetFirewallPolicyResult
    {
        /// <summary>
        /// A list of associations that belong to this firewall policy.
        /// </summary>
        public readonly ImmutableArray<Outputs.FirewallPolicyAssociationResponse> Associations;
        /// <summary>
        /// Creation timestamp in RFC3339 text format.
        /// </summary>
        public readonly string CreationTimestamp;
        /// <summary>
        /// An optional description of this resource. Provide this property when you create the resource.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// Specifies a fingerprint for this resource, which is essentially a hash of the metadata's contents and used for optimistic locking. The fingerprint is initially generated by Compute Engine and changes after every request to modify or update metadata. You must always provide an up-to-date fingerprint hash in order to update or change metadata, otherwise the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make get() request to the firewall policy.
        /// </summary>
        public readonly string Fingerprint;
        /// <summary>
        /// [Output only] Type of the resource. Always compute#firewallPolicyfor firewall policies
        /// </summary>
        public readonly string Kind;
        /// <summary>
        /// Name of the resource. It is a numeric ID allocated by GCP which uniquely identifies the Firewall Policy.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The parent of the firewall policy.
        /// </summary>
        public readonly string Parent;
        /// <summary>
        /// URL of the region where the regional firewall policy resides. This field is not applicable to global firewall policies. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
        /// </summary>
        public readonly string Region;
        /// <summary>
        /// Total count of all firewall policy rule tuples. A firewall policy can not exceed a set number of tuples.
        /// </summary>
        public readonly int RuleTupleCount;
        /// <summary>
        /// A list of rules that belong to this policy. There must always be a default rule (rule with priority 2147483647 and match "*"). If no rules are provided when creating a firewall policy, a default rule with action "allow" will be added.
        /// </summary>
        public readonly ImmutableArray<Outputs.FirewallPolicyRuleResponse> Rules;
        /// <summary>
        /// Server-defined URL for the resource.
        /// </summary>
        public readonly string SelfLink;
        /// <summary>
        /// Server-defined URL for this resource with the resource id.
        /// </summary>
        public readonly string SelfLinkWithId;
        /// <summary>
        /// User-provided name of the Organization firewall plicy. The name should be unique in the organization in which the firewall policy is created. This name must be set on creation and cannot be changed. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        /// </summary>
        public readonly string ShortName;

        [OutputConstructor]
        private GetFirewallPolicyResult(
            ImmutableArray<Outputs.FirewallPolicyAssociationResponse> associations,

            string creationTimestamp,

            string description,

            string fingerprint,

            string kind,

            string name,

            string parent,

            string region,

            int ruleTupleCount,

            ImmutableArray<Outputs.FirewallPolicyRuleResponse> rules,

            string selfLink,

            string selfLinkWithId,

            string shortName)
        {
            Associations = associations;
            CreationTimestamp = creationTimestamp;
            Description = description;
            Fingerprint = fingerprint;
            Kind = kind;
            Name = name;
            Parent = parent;
            Region = region;
            RuleTupleCount = ruleTupleCount;
            Rules = rules;
            SelfLink = selfLink;
            SelfLinkWithId = selfLinkWithId;
            ShortName = shortName;
        }
    }
}
