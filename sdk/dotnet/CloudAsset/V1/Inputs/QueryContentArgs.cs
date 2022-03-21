// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.CloudAsset.V1.Inputs
{

    /// <summary>
    /// The query content.
    /// </summary>
    public sealed class QueryContentArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// An IAM Policy Analysis query, which could be used in the AssetService.AnalyzeIamPolicy rpc or the AssetService.AnalyzeIamPolicyLongrunning rpc.
        /// </summary>
        [Input("iamPolicyAnalysisQuery")]
        public Input<Inputs.IamPolicyAnalysisQueryArgs>? IamPolicyAnalysisQuery { get; set; }

        public QueryContentArgs()
        {
        }
    }
}
