// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Beta.Outputs
{

    /// <summary>
    /// The configuration for a built-in load balancing policy.
    /// </summary>
    [OutputType]
    public sealed class BackendServiceLocalityLoadBalancingPolicyConfigPolicyResponse
    {
        /// <summary>
        /// The name of a locality load balancer policy to be used. The value should be one of the predefined ones as supported by localityLbPolicy, although at the moment only ROUND_ROBIN is supported. This field should only be populated when the customPolicy field is not used. Note that specifying the same policy more than once for a backend is not a valid configuration and will be rejected.
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private BackendServiceLocalityLoadBalancingPolicyConfigPolicyResponse(string name)
        {
            Name = name;
        }
    }
}
