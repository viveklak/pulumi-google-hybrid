// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.V1.Outputs
{

    [OutputType]
    public sealed class FirewallPolicyRuleSecureTagResponse
    {
        /// <summary>
        /// Name of the secure tag, created with TagManager's TagValue API.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// State of the secure tag, either `EFFECTIVE` or `INEFFECTIVE`. A secure tag is `INEFFECTIVE` when it is deleted or its network is deleted.
        /// </summary>
        public readonly string State;

        [OutputConstructor]
        private FirewallPolicyRuleSecureTagResponse(
            string name,

            string state)
        {
            Name = name;
            State = state;
        }
    }
}
