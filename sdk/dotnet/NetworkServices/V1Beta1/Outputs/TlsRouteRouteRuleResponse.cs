// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.NetworkServices.V1Beta1.Outputs
{

    /// <summary>
    /// Specifies how to match traffic and how to route traffic when traffic is matched.
    /// </summary>
    [OutputType]
    public sealed class TlsRouteRouteRuleResponse
    {
        /// <summary>
        /// The detailed rule defining how to route matched traffic.
        /// </summary>
        public readonly Outputs.TlsRouteRouteActionResponse Action;
        /// <summary>
        /// RouteMatch defines the predicate used to match requests to a given action. Multiple match types are “OR”ed for evaluation.
        /// </summary>
        public readonly ImmutableArray<Outputs.TlsRouteRouteMatchResponse> Matches;

        [OutputConstructor]
        private TlsRouteRouteRuleResponse(
            Outputs.TlsRouteRouteActionResponse action,

            ImmutableArray<Outputs.TlsRouteRouteMatchResponse> matches)
        {
            Action = action;
            Matches = matches;
        }
    }
}
