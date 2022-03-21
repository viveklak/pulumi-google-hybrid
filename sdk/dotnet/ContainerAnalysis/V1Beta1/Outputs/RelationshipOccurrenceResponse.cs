// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.ContainerAnalysis.V1Beta1.Outputs
{

    /// <summary>
    /// RelationshipOccurrence represents an SPDX Relationship section: https://spdx.github.io/spdx-spec/7-relationships-between-SPDX-elements/
    /// </summary>
    [OutputType]
    public sealed class RelationshipOccurrenceResponse
    {
        /// <summary>
        /// A place for the SPDX file creator to record any general comments about the relationship
        /// </summary>
        public readonly string Comment;
        /// <summary>
        /// Also referred to as SPDXRef-A The source SPDX element (file, package, etc)
        /// </summary>
        public readonly string Source;
        /// <summary>
        /// Also referred to as SPDXRef-B The target SPDC element (file, package, etc) In cases where there are "known unknowns", the use of the keyword NOASSERTION can be used The keywords NONE can be used to indicate that an SPDX element (package/file/snippet) has no other elements connected by some relationship to it
        /// </summary>
        public readonly string Target;
        /// <summary>
        /// The type of relationship between the source and target SPDX elements
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private RelationshipOccurrenceResponse(
            string comment,

            string source,

            string target,

            string type)
        {
            Comment = comment;
            Source = source;
            Target = target;
            Type = type;
        }
    }
}
