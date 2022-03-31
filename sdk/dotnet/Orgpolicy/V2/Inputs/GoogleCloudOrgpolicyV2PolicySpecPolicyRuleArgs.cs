// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Orgpolicy.V2.Inputs
{

    /// <summary>
    /// A rule used to express this policy.
    /// </summary>
    public sealed class GoogleCloudOrgpolicyV2PolicySpecPolicyRuleArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Setting this to true means that all values are allowed. This field can be set only in Policies for list constraints.
        /// </summary>
        [Input("allowAll")]
        public Input<bool>? AllowAll { get; set; }

        /// <summary>
        /// A condition which determines whether this rule is used in the evaluation of the policy. When set, the `expression` field in the `Expr' must include from 1 to 10 subexpressions, joined by the "||" or "&amp;&amp;" operators. Each subexpression must be of the form "resource.matchTag('/tag_key_short_name, 'tag_value_short_name')". or "resource.matchTagId('tagKeys/key_id', 'tagValues/value_id')". where key_name and value_name are the resource names for Label Keys and Values. These names are available from the Tag Manager Service. An example expression is: "resource.matchTag('123456789/environment, 'prod')". or "resource.matchTagId('tagKeys/123', 'tagValues/456')".
        /// </summary>
        [Input("condition")]
        public Input<Inputs.GoogleTypeExprArgs>? Condition { get; set; }

        /// <summary>
        /// Setting this to true means that all values are denied. This field can be set only in Policies for list constraints.
        /// </summary>
        [Input("denyAll")]
        public Input<bool>? DenyAll { get; set; }

        /// <summary>
        /// If `true`, then the `Policy` is enforced. If `false`, then any configuration is acceptable. This field can be set only in Policies for boolean constraints.
        /// </summary>
        [Input("enforce")]
        public Input<bool>? Enforce { get; set; }

        /// <summary>
        /// List of values to be used for this PolicyRule. This field can be set only in Policies for list constraints.
        /// </summary>
        [Input("values")]
        public Input<Inputs.GoogleCloudOrgpolicyV2PolicySpecPolicyRuleStringValuesArgs>? Values { get; set; }

        public GoogleCloudOrgpolicyV2PolicySpecPolicyRuleArgs()
        {
        }
    }
}
