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
    /// Creates a new policy in the specified project using the data included in the request.
    /// </summary>
    [GoogleNativeResourceType("google-native:compute/v1:RegionSecurityPolicy")]
    public partial class RegionSecurityPolicy : Pulumi.CustomResource
    {
        [Output("adaptiveProtectionConfig")]
        public Output<Outputs.SecurityPolicyAdaptiveProtectionConfigResponse> AdaptiveProtectionConfig { get; private set; } = null!;

        [Output("advancedOptionsConfig")]
        public Output<Outputs.SecurityPolicyAdvancedOptionsConfigResponse> AdvancedOptionsConfig { get; private set; } = null!;

        /// <summary>
        /// Creation timestamp in RFC3339 text format.
        /// </summary>
        [Output("creationTimestamp")]
        public Output<string> CreationTimestamp { get; private set; } = null!;

        [Output("ddosProtectionConfig")]
        public Output<Outputs.SecurityPolicyDdosProtectionConfigResponse> DdosProtectionConfig { get; private set; } = null!;

        /// <summary>
        /// An optional description of this resource. Provide this property when you create the resource.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

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
        /// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("recaptchaOptionsConfig")]
        public Output<Outputs.SecurityPolicyRecaptchaOptionsConfigResponse> RecaptchaOptionsConfig { get; private set; } = null!;

        /// <summary>
        /// URL of the region where the regional security policy resides. This field is not applicable to global security policies.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// A list of rules that belong to this policy. There must always be a default rule (rule with priority 2147483647 and match "*"). If no rules are provided when creating a security policy, a default rule with action "allow" will be added.
        /// </summary>
        [Output("rules")]
        public Output<ImmutableArray<Outputs.SecurityPolicyRuleResponse>> Rules { get; private set; } = null!;

        /// <summary>
        /// Server-defined URL for the resource.
        /// </summary>
        [Output("selfLink")]
        public Output<string> SelfLink { get; private set; } = null!;

        /// <summary>
        /// The type indicates the intended use of the security policy. - CLOUD_ARMOR: Cloud Armor backend security policies can be configured to filter incoming HTTP requests targeting backend services. They filter requests before they hit the origin servers. - CLOUD_ARMOR_EDGE: Cloud Armor edge security policies can be configured to filter incoming HTTP requests targeting backend services (including Cloud CDN-enabled) as well as backend buckets (Cloud Storage). They filter requests before the request is served from Google's cache. - CLOUD_ARMOR_INTERNAL_SERVICE: Cloud Armor internal service policies can be configured to filter HTTP requests targeting services managed by Traffic Director in a service mesh. They filter requests before the request is served from the application. This field can be set only at resource creation time.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a RegionSecurityPolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RegionSecurityPolicy(string name, RegionSecurityPolicyArgs args, CustomResourceOptions? options = null)
            : base("google-native:compute/v1:RegionSecurityPolicy", name, args ?? new RegionSecurityPolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private RegionSecurityPolicy(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:compute/v1:RegionSecurityPolicy", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing RegionSecurityPolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RegionSecurityPolicy Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new RegionSecurityPolicy(name, id, options);
        }
    }

    public sealed class RegionSecurityPolicyArgs : Pulumi.ResourceArgs
    {
        [Input("adaptiveProtectionConfig")]
        public Input<Inputs.SecurityPolicyAdaptiveProtectionConfigArgs>? AdaptiveProtectionConfig { get; set; }

        [Input("advancedOptionsConfig")]
        public Input<Inputs.SecurityPolicyAdvancedOptionsConfigArgs>? AdvancedOptionsConfig { get; set; }

        [Input("ddosProtectionConfig")]
        public Input<Inputs.SecurityPolicyDdosProtectionConfigArgs>? DdosProtectionConfig { get; set; }

        /// <summary>
        /// An optional description of this resource. Provide this property when you create the resource.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("recaptchaOptionsConfig")]
        public Input<Inputs.SecurityPolicyRecaptchaOptionsConfigArgs>? RecaptchaOptionsConfig { get; set; }

        [Input("region", required: true)]
        public Input<string> Region { get; set; } = null!;

        /// <summary>
        /// An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
        /// </summary>
        [Input("requestId")]
        public Input<string>? RequestId { get; set; }

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
        /// The type indicates the intended use of the security policy. - CLOUD_ARMOR: Cloud Armor backend security policies can be configured to filter incoming HTTP requests targeting backend services. They filter requests before they hit the origin servers. - CLOUD_ARMOR_EDGE: Cloud Armor edge security policies can be configured to filter incoming HTTP requests targeting backend services (including Cloud CDN-enabled) as well as backend buckets (Cloud Storage). They filter requests before the request is served from Google's cache. - CLOUD_ARMOR_INTERNAL_SERVICE: Cloud Armor internal service policies can be configured to filter HTTP requests targeting services managed by Traffic Director in a service mesh. They filter requests before the request is served from the application. This field can be set only at resource creation time.
        /// </summary>
        [Input("type")]
        public Input<Pulumi.GoogleNative.Compute.V1.RegionSecurityPolicyType>? Type { get; set; }

        /// <summary>
        /// If true, the request will not be committed.
        /// </summary>
        [Input("validateOnly")]
        public Input<string>? ValidateOnly { get; set; }

        public RegionSecurityPolicyArgs()
        {
        }
    }
}
