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
    public sealed class NodeTemplateNodeTypeFlexibilityResponse
    {
        public readonly string Cpus;
        public readonly string LocalSsd;
        public readonly string Memory;

        [OutputConstructor]
        private NodeTemplateNodeTypeFlexibilityResponse(
            string cpus,

            string localSsd,

            string memory)
        {
            Cpus = cpus;
            LocalSsd = localSsd;
            Memory = memory;
        }
    }
}
