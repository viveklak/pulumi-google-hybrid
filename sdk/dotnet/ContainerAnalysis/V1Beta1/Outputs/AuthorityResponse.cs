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
    /// Note kind that represents a logical attestation "role" or "authority". For example, an organization might have one `Authority` for "QA" and one for "build". This note is intended to act strictly as a grouping mechanism for the attached occurrences (Attestations). This grouping mechanism also provides a security boundary, since IAM ACLs gate the ability for a principle to attach an occurrence to a given note. It also provides a single point of lookup to find all attached attestation occurrences, even if they don't all live in the same project.
    /// </summary>
    [OutputType]
    public sealed class AuthorityResponse
    {
        /// <summary>
        /// Hint hints at the purpose of the attestation authority.
        /// </summary>
        public readonly Outputs.HintResponse Hint;

        [OutputConstructor]
        private AuthorityResponse(Outputs.HintResponse hint)
        {
            Hint = hint;
        }
    }
}
