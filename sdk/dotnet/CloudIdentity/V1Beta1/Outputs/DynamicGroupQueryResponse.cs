// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.CloudIdentity.V1Beta1.Outputs
{

    /// <summary>
    /// Defines a query on a resource.
    /// </summary>
    [OutputType]
    public sealed class DynamicGroupQueryResponse
    {
        /// <summary>
        /// Query that determines the memberships of the dynamic group. Examples: All users with at least one `organizations.department` of engineering. `user.organizations.exists(org, org.department=='engineering')` All users with at least one location that has `area` of `foo` and `building_id` of `bar`. `user.locations.exists(loc, loc.area=='foo' &amp;&amp; loc.building_id=='bar')` All users with any variation of the name John Doe (case-insensitive queries add `equalsIgnoreCase()` to the value being queried). `user.name.value.equalsIgnoreCase('jOhn DoE')`
        /// </summary>
        public readonly string Query;
        public readonly string ResourceType;

        [OutputConstructor]
        private DynamicGroupQueryResponse(
            string query,

            string resourceType)
        {
            Query = query;
            ResourceType = resourceType;
        }
    }
}
