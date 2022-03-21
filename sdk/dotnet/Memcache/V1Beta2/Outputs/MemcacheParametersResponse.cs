// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Memcache.V1Beta2.Outputs
{

    [OutputType]
    public sealed class MemcacheParametersResponse
    {
        /// <summary>
        /// User defined set of parameters to use in the memcached process.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Params;

        [OutputConstructor]
        private MemcacheParametersResponse(ImmutableDictionary<string, string> @params)
        {
            Params = @params;
        }
    }
}
