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
    /// Sets the access control policy on the specified resource. Replaces any existing policy.
    /// Note - this resource's API doesn't support deletion. When deleted, the resource will persist
    /// on Google Cloud even though it will be deleted from Pulumi state.
    /// </summary>
    [GoogleNativeResourceType("google-native:compute/beta:BackendBucketIamPolicy")]
    public partial class BackendBucketIamPolicy : Pulumi.CustomResource
    {
        /// <summary>
        /// Specifies cloud audit logging configuration for this policy.
        /// </summary>
        [Output("auditConfigs")]
        public Output<ImmutableArray<Outputs.AuditConfigResponse>> AuditConfigs { get; private set; } = null!;

        /// <summary>
        /// Associates a list of `members`, or principals, with a `role`. Optionally, may specify a `condition` that determines how and when the `bindings` are applied. Each of the `bindings` must contain at least one principal. The `bindings` in a `Policy` can refer to up to 1,500 principals; up to 250 of these principals can be Google groups. Each occurrence of a principal counts towards these limits. For example, if the `bindings` grant 50 different roles to `user:alice@example.com`, and not to any other principal, then you can add another 1,450 principals to the `bindings` in the `Policy`.
        /// </summary>
        [Output("bindings")]
        public Output<ImmutableArray<Outputs.BindingResponse>> Bindings { get; private set; } = null!;

        /// <summary>
        /// `etag` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other. It is strongly suggested that systems make use of the `etag` in the read-modify-write cycle to perform policy updates in order to avoid race conditions: An `etag` is returned in the response to `getIamPolicy`, and systems are expected to put that etag in the request to `setIamPolicy` to ensure that their change will be applied to the same version of the policy. **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost.
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        /// <summary>
        /// This is deprecated and has no effect. Do not use.
        /// </summary>
        [Output("rules")]
        public Output<ImmutableArray<Outputs.RuleResponse>> Rules { get; private set; } = null!;

        /// <summary>
        /// Specifies the format of the policy. Valid values are `0`, `1`, and `3`. Requests that specify an invalid value are rejected. Any operation that affects conditional role bindings must specify version `3`. This requirement applies to the following operations: * Getting a policy that includes a conditional role binding * Adding a conditional role binding to a policy * Changing a conditional role binding in a policy * Removing any role binding, with or without a condition, from a policy that includes conditions **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost. If a policy does not include any conditions, operations on that policy may specify any valid version or leave the field unset. To learn which resources support conditions in their IAM policies, see the [IAM documentation](https://cloud.google.com/iam/help/conditions/resource-policies).
        /// </summary>
        [Output("version")]
        public Output<int> Version { get; private set; } = null!;


        /// <summary>
        /// Create a BackendBucketIamPolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public BackendBucketIamPolicy(string name, BackendBucketIamPolicyArgs args, CustomResourceOptions? options = null)
            : base("google-native:compute/beta:BackendBucketIamPolicy", name, args ?? new BackendBucketIamPolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private BackendBucketIamPolicy(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:compute/beta:BackendBucketIamPolicy", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing BackendBucketIamPolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static BackendBucketIamPolicy Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new BackendBucketIamPolicy(name, id, options);
        }
    }

    public sealed class BackendBucketIamPolicyArgs : Pulumi.ResourceArgs
    {
        [Input("auditConfigs")]
        private InputList<Inputs.AuditConfigArgs>? _auditConfigs;

        /// <summary>
        /// Specifies cloud audit logging configuration for this policy.
        /// </summary>
        public InputList<Inputs.AuditConfigArgs> AuditConfigs
        {
            get => _auditConfigs ?? (_auditConfigs = new InputList<Inputs.AuditConfigArgs>());
            set => _auditConfigs = value;
        }

        [Input("bindings")]
        private InputList<Inputs.BindingArgs>? _bindings;

        /// <summary>
        /// Associates a list of `members`, or principals, with a `role`. Optionally, may specify a `condition` that determines how and when the `bindings` are applied. Each of the `bindings` must contain at least one principal. The `bindings` in a `Policy` can refer to up to 1,500 principals; up to 250 of these principals can be Google groups. Each occurrence of a principal counts towards these limits. For example, if the `bindings` grant 50 different roles to `user:alice@example.com`, and not to any other principal, then you can add another 1,450 principals to the `bindings` in the `Policy`.
        /// </summary>
        public InputList<Inputs.BindingArgs> Bindings
        {
            get => _bindings ?? (_bindings = new InputList<Inputs.BindingArgs>());
            set => _bindings = value;
        }

        /// <summary>
        /// `etag` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other. It is strongly suggested that systems make use of the `etag` in the read-modify-write cycle to perform policy updates in order to avoid race conditions: An `etag` is returned in the response to `getIamPolicy`, and systems are expected to put that etag in the request to `setIamPolicy` to ensure that their change will be applied to the same version of the policy. **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost.
        /// </summary>
        [Input("etag")]
        public Input<string>? Etag { get; set; }

        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("resource", required: true)]
        public Input<string> Resource { get; set; } = null!;

        [Input("rules")]
        private InputList<Inputs.RuleArgs>? _rules;

        /// <summary>
        /// This is deprecated and has no effect. Do not use.
        /// </summary>
        public InputList<Inputs.RuleArgs> Rules
        {
            get => _rules ?? (_rules = new InputList<Inputs.RuleArgs>());
            set => _rules = value;
        }

        /// <summary>
        /// Specifies the format of the policy. Valid values are `0`, `1`, and `3`. Requests that specify an invalid value are rejected. Any operation that affects conditional role bindings must specify version `3`. This requirement applies to the following operations: * Getting a policy that includes a conditional role binding * Adding a conditional role binding to a policy * Changing a conditional role binding in a policy * Removing any role binding, with or without a condition, from a policy that includes conditions **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost. If a policy does not include any conditions, operations on that policy may specify any valid version or leave the field unset. To learn which resources support conditions in their IAM policies, see the [IAM documentation](https://cloud.google.com/iam/help/conditions/resource-policies).
        /// </summary>
        [Input("version")]
        public Input<int>? Version { get; set; }

        public BackendBucketIamPolicyArgs()
        {
        }
    }
}
