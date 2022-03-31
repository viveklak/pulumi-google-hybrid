// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.ContainerAnalysis.V1Alpha1.Outputs
{

    /// <summary>
    /// RelationshipNote represents an SPDX Relationship section: https://spdx.github.io/spdx-spec/7-relationships-between-SPDX-elements/
    /// </summary>
    [OutputType]
    public sealed class RelationshipNoteResponse
    {
        /// <summary>
        /// The type of relationship between the source and target SPDX elements
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private RelationshipNoteResponse(string type)
        {
            Type = type;
        }
    }
}
