// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.AccessContextManager.V1Beta.Outputs
{

    /// <summary>
    /// `CustomLevel` is an `AccessLevel` using the Cloud Common Expression Language to represent the necessary conditions for the level to apply to a request. See CEL spec at: https://github.com/google/cel-spec
    /// </summary>
    [OutputType]
    public sealed class CustomLevelResponse
    {
        /// <summary>
        /// A Cloud CEL expression evaluating to a boolean.
        /// </summary>
        public readonly Outputs.ExprResponse Expr;

        [OutputConstructor]
        private CustomLevelResponse(Outputs.ExprResponse expr)
        {
            Expr = expr;
        }
    }
}
