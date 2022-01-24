// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi.Utilities;

namespace Pulumi.GoogleNative.NetworkSecurity.V1
{
    public static class GetAuthorizationPolicy
    {
        /// <summary>
        /// Gets details of a single AuthorizationPolicy.
        /// </summary>
        public static Task<GetAuthorizationPolicyResult> InvokeAsync(GetAuthorizationPolicyArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetAuthorizationPolicyResult>("google-native:networksecurity/v1:getAuthorizationPolicy", args ?? new GetAuthorizationPolicyArgs(), options.WithDefaults());

        /// <summary>
        /// Gets details of a single AuthorizationPolicy.
        /// </summary>
        public static Output<GetAuthorizationPolicyResult> Invoke(GetAuthorizationPolicyInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetAuthorizationPolicyResult>("google-native:networksecurity/v1:getAuthorizationPolicy", args ?? new GetAuthorizationPolicyInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAuthorizationPolicyArgs : Pulumi.InvokeArgs
    {
        [Input("authorizationPolicyId", required: true)]
        public string AuthorizationPolicyId { get; set; } = null!;

        [Input("location", required: true)]
        public string Location { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        public GetAuthorizationPolicyArgs()
        {
        }
    }

    public sealed class GetAuthorizationPolicyInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("authorizationPolicyId", required: true)]
        public Input<string> AuthorizationPolicyId { get; set; } = null!;

        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        public GetAuthorizationPolicyInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetAuthorizationPolicyResult
    {
        /// <summary>
        /// The action to take when a rule match is found. Possible values are "ALLOW" or "DENY".
        /// </summary>
        public readonly string Action;
        /// <summary>
        /// The timestamp when the resource was created.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// Optional. Free-text description of the resource.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// Optional. Set of label tags associated with the AuthorizationPolicy resource.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Labels;
        /// <summary>
        /// Name of the AuthorizationPolicy resource. It matches pattern `projects/{project}/locations/{location}/authorizationPolicies/`.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Optional. List of rules to match. Note that at least one of the rules must match in order for the action specified in the 'action' field to be taken. A rule is a match if there is a matching source and destination. If left blank, the action specified in the `action` field will be applied on every request.
        /// </summary>
        public readonly ImmutableArray<Outputs.RuleResponse> Rules;
        /// <summary>
        /// The timestamp when the resource was updated.
        /// </summary>
        public readonly string UpdateTime;

        [OutputConstructor]
        private GetAuthorizationPolicyResult(
            string action,

            string createTime,

            string description,

            ImmutableDictionary<string, string> labels,

            string name,

            ImmutableArray<Outputs.RuleResponse> rules,

            string updateTime)
        {
            Action = action;
            CreateTime = createTime;
            Description = description;
            Labels = labels;
            Name = name;
            Rules = rules;
            UpdateTime = updateTime;
        }
    }
}
