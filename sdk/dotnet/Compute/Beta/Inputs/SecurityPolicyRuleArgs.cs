// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Beta.Inputs
{

    /// <summary>
    /// Represents a rule that describes one or more match conditions along with the action to be taken when traffic matches this condition (allow or deny).
    /// </summary>
    public sealed class SecurityPolicyRuleArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Action to perform when the client connection triggers the rule. Can currently be either "allow" or "deny()" where valid values for status are 403, 404, and 502.
        /// </summary>
        [Input("action")]
        public Input<string>? Action { get; set; }

        /// <summary>
        /// An optional description of this resource. Provide this property when you create the resource.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The direction in which this rule applies. This field may only be specified when versioned_expr is set to FIREWALL.
        /// </summary>
        [Input("direction")]
        public Input<Pulumi.GoogleNative.Compute.Beta.SecurityPolicyRuleDirection>? Direction { get; set; }

        /// <summary>
        /// Denotes whether to enable logging for a particular rule. If logging is enabled, logs will be exported to the configured export destination in Stackdriver. Logs may be exported to BigQuery or Pub/Sub. Note: you cannot enable logging on "goto_next" rules. This field may only be specified when the versioned_expr is set to FIREWALL.
        /// </summary>
        [Input("enableLogging")]
        public Input<bool>? EnableLogging { get; set; }

        /// <summary>
        /// [Output only] Type of the resource. Always compute#securityPolicyRule for security policy rules
        /// </summary>
        [Input("kind")]
        public Input<string>? Kind { get; set; }

        /// <summary>
        /// A match condition that incoming traffic is evaluated against. If it evaluates to true, the corresponding 'action' is enforced.
        /// </summary>
        [Input("match")]
        public Input<Inputs.SecurityPolicyRuleMatcherArgs>? Match { get; set; }

        /// <summary>
        /// If set to true, the specified action is not enforced.
        /// </summary>
        [Input("preview")]
        public Input<bool>? Preview { get; set; }

        /// <summary>
        /// An integer indicating the priority of a rule in the list. The priority must be a positive value between 0 and 2147483647. Rules are evaluated from highest to lowest priority where 0 is the highest priority and 2147483647 is the lowest priority.
        /// </summary>
        [Input("priority")]
        public Input<int>? Priority { get; set; }

        /// <summary>
        /// Identifier for the rule. This is only unique within the given security policy. This can only be set during rule creation, if rule number is not specified it will be generated by the server.
        /// </summary>
        [Input("ruleNumber")]
        public Input<string>? RuleNumber { get; set; }

        /// <summary>
        /// [Output Only] Calculation of the complexity of a single firewall security policy rule.
        /// </summary>
        [Input("ruleTupleCount")]
        public Input<int>? RuleTupleCount { get; set; }

        [Input("targetResources")]
        private InputList<string>? _targetResources;

        /// <summary>
        /// A list of network resource URLs to which this rule applies. This field allows you to control which network's VMs get this rule. If this field is left blank, all VMs within the organization will receive the rule. This field may only be specified when versioned_expr is set to FIREWALL.
        /// </summary>
        public InputList<string> TargetResources
        {
            get => _targetResources ?? (_targetResources = new InputList<string>());
            set => _targetResources = value;
        }

        [Input("targetServiceAccounts")]
        private InputList<string>? _targetServiceAccounts;

        /// <summary>
        /// A list of service accounts indicating the sets of instances that are applied with this rule.
        /// </summary>
        public InputList<string> TargetServiceAccounts
        {
            get => _targetServiceAccounts ?? (_targetServiceAccounts = new InputList<string>());
            set => _targetServiceAccounts = value;
        }

        public SecurityPolicyRuleArgs()
        {
        }
    }
}
