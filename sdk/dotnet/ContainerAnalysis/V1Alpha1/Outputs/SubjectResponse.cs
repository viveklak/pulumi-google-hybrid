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
    /// Subject refers to the subject of the intoto statement
    /// </summary>
    [OutputType]
    public sealed class SubjectResponse
    {
        /// <summary>
        /// "": "" Algorithms can be e.g. sha256, sha512 See https://github.com/in-toto/attestation/blob/main/spec/field_types.md#DigestSet
        /// </summary>
        public readonly ImmutableDictionary<string, string> Digest;
        /// <summary>
        /// name is the name of the Subject used here
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private SubjectResponse(
            ImmutableDictionary<string, string> digest,

            string name)
        {
            Digest = digest;
            Name = name;
        }
    }
}
