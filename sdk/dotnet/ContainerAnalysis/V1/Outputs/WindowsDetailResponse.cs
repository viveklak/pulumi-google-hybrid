// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.ContainerAnalysis.V1.Outputs
{

    [OutputType]
    public sealed class WindowsDetailResponse
    {
        /// <summary>
        /// The [CPE URI](https://cpe.mitre.org/specification/) this vulnerability affects.
        /// </summary>
        public readonly string CpeUri;
        /// <summary>
        /// The description of this vulnerability.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The names of the KBs which have hotfixes to mitigate this vulnerability. Note that there may be multiple hotfixes (and thus multiple KBs) that mitigate a given vulnerability. Currently any listed KBs presence is considered a fix.
        /// </summary>
        public readonly ImmutableArray<Outputs.KnowledgeBaseResponse> FixingKbs;
        /// <summary>
        /// The name of this vulnerability.
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private WindowsDetailResponse(
            string cpeUri,

            string description,

            ImmutableArray<Outputs.KnowledgeBaseResponse> fixingKbs,

            string name)
        {
            CpeUri = cpeUri;
            Description = description;
            FixingKbs = fixingKbs;
            Name = name;
        }
    }
}
