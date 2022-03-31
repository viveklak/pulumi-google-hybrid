// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Alpha.Outputs
{

    /// <summary>
    /// This is deprecated and has no effect. Do not use.
    /// </summary>
    [OutputType]
    public sealed class RuleResponse
    {
        /// <summary>
        /// This is deprecated and has no effect. Do not use.
        /// </summary>
        public readonly string Action;
        /// <summary>
        /// This is deprecated and has no effect. Do not use.
        /// </summary>
        public readonly ImmutableArray<Outputs.ConditionResponse> Conditions;
        /// <summary>
        /// This is deprecated and has no effect. Do not use.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// This is deprecated and has no effect. Do not use.
        /// </summary>
        public readonly ImmutableArray<string> Ins;
        /// <summary>
        /// This is deprecated and has no effect. Do not use.
        /// </summary>
        public readonly ImmutableArray<Outputs.LogConfigResponse> LogConfigs;
        /// <summary>
        /// This is deprecated and has no effect. Do not use.
        /// </summary>
        public readonly ImmutableArray<string> NotIns;
        /// <summary>
        /// This is deprecated and has no effect. Do not use.
        /// </summary>
        public readonly ImmutableArray<string> Permissions;

        [OutputConstructor]
        private RuleResponse(
            string action,

            ImmutableArray<Outputs.ConditionResponse> conditions,

            string description,

            ImmutableArray<string> ins,

            ImmutableArray<Outputs.LogConfigResponse> logConfigs,

            ImmutableArray<string> notIns,

            ImmutableArray<string> permissions)
        {
            Action = action;
            Conditions = conditions;
            Description = description;
            Ins = ins;
            LogConfigs = logConfigs;
            NotIns = notIns;
            Permissions = permissions;
        }
    }
}
