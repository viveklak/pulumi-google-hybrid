// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.ContainerAnalysis.V1Beta1.Inputs
{

    /// <summary>
    /// RelationshipNote represents an SPDX Relationship section: https://spdx.github.io/spdx-spec/7-relationships-between-SPDX-elements/
    /// </summary>
    public sealed class RelationshipNoteArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The type of relationship between the source and target SPDX elements
        /// </summary>
        [Input("type")]
        public Input<Pulumi.GoogleNative.ContainerAnalysis.V1Beta1.RelationshipNoteType>? Type { get; set; }

        public RelationshipNoteArgs()
        {
        }
    }
}
