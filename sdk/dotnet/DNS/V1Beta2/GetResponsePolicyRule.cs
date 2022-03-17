// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.DNS.V1Beta2
{
    public static class GetResponsePolicyRule
    {
        /// <summary>
        /// Fetches the representation of an existing Response Policy Rule.
        /// </summary>
        public static Task<GetResponsePolicyRuleResult> InvokeAsync(GetResponsePolicyRuleArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetResponsePolicyRuleResult>("google-native:dns/v1beta2:getResponsePolicyRule", args ?? new GetResponsePolicyRuleArgs(), options.WithDefaults());

        /// <summary>
        /// Fetches the representation of an existing Response Policy Rule.
        /// </summary>
        public static Output<GetResponsePolicyRuleResult> Invoke(GetResponsePolicyRuleInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetResponsePolicyRuleResult>("google-native:dns/v1beta2:getResponsePolicyRule", args ?? new GetResponsePolicyRuleInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetResponsePolicyRuleArgs : Pulumi.InvokeArgs
    {
        [Input("clientOperationId")]
        public string? ClientOperationId { get; set; }

        [Input("project")]
        public string? Project { get; set; }

        [Input("responsePolicy", required: true)]
        public string ResponsePolicy { get; set; } = null!;

        [Input("responsePolicyRule", required: true)]
        public string ResponsePolicyRule { get; set; } = null!;

        public GetResponsePolicyRuleArgs()
        {
        }
    }

    public sealed class GetResponsePolicyRuleInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("clientOperationId")]
        public Input<string>? ClientOperationId { get; set; }

        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("responsePolicy", required: true)]
        public Input<string> ResponsePolicy { get; set; } = null!;

        [Input("responsePolicyRule", required: true)]
        public Input<string> ResponsePolicyRule { get; set; } = null!;

        public GetResponsePolicyRuleInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetResponsePolicyRuleResult
    {
        /// <summary>
        /// Answer this query with a behavior rather than DNS data.
        /// </summary>
        public readonly string Behavior;
        /// <summary>
        /// The DNS name (wildcard or exact) to apply this rule to. Must be unique within the Response Policy Rule.
        /// </summary>
        public readonly string DnsName;
        public readonly string Kind;
        /// <summary>
        /// Answer this query directly with DNS data. These ResourceRecordSets override any other DNS behavior for the matched name; in particular they override private zones, the public internet, and GCP internal DNS. No SOA nor NS types are allowed.
        /// </summary>
        public readonly Outputs.ResponsePolicyRuleLocalDataResponse LocalData;
        /// <summary>
        /// An identifier for this rule. Must be unique with the ResponsePolicy.
        /// </summary>
        public readonly string RuleName;

        [OutputConstructor]
        private GetResponsePolicyRuleResult(
            string behavior,

            string dnsName,

            string kind,

            Outputs.ResponsePolicyRuleLocalDataResponse localData,

            string ruleName)
        {
            Behavior = behavior;
            DnsName = dnsName;
            Kind = kind;
            LocalData = localData;
            RuleName = ruleName;
        }
    }
}
