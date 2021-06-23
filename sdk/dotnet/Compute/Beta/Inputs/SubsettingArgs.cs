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
    /// Subsetting configuration for this BackendService. Currently this is applicable only for Internal TCP/UDP load balancing, Internal HTTP(S) load balancing and Traffic Director.
    /// </summary>
    public sealed class SubsettingArgs : Pulumi.ResourceArgs
    {
        [Input("policy")]
        public Input<Pulumi.GoogleNative.Compute.Beta.SubsettingPolicy>? Policy { get; set; }

        public SubsettingArgs()
        {
        }
    }
}
