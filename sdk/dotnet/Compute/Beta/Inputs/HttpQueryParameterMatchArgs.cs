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
    /// HttpRouteRuleMatch criteria for a request's query parameter.
    /// </summary>
    public sealed class HttpQueryParameterMatchArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The queryParameterMatch matches if the value of the parameter exactly matches the contents of exactMatch. Only one of presentMatch, exactMatch or regexMatch must be set.
        /// </summary>
        [Input("exactMatch")]
        public Input<string>? ExactMatch { get; set; }

        /// <summary>
        /// The name of the query parameter to match. The query parameter must exist in the request, in the absence of which the request match fails.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Specifies that the queryParameterMatch matches if the request contains the query parameter, irrespective of whether the parameter has a value or not. Only one of presentMatch, exactMatch or regexMatch must be set.
        /// </summary>
        [Input("presentMatch")]
        public Input<bool>? PresentMatch { get; set; }

        /// <summary>
        /// The queryParameterMatch matches if the value of the parameter matches the regular expression specified by regexMatch. For the regular expression grammar, please see github.com/google/re2/wiki/Syntax Only one of presentMatch, exactMatch or regexMatch must be set. Note that regexMatch only applies when the loadBalancingScheme is set to INTERNAL_SELF_MANAGED.
        /// </summary>
        [Input("regexMatch")]
        public Input<string>? RegexMatch { get; set; }

        public HttpQueryParameterMatchArgs()
        {
        }
    }
}
