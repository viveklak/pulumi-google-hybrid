// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.V1.Outputs
{

    [OutputType]
    public sealed class ConditionResponse
    {
        /// <summary>
        /// This is deprecated and has no effect. Do not use.
        /// </summary>
        public readonly string Iam;
        /// <summary>
        /// This is deprecated and has no effect. Do not use.
        /// </summary>
        public readonly string Op;
        /// <summary>
        /// This is deprecated and has no effect. Do not use.
        /// </summary>
        public readonly string Svc;
        /// <summary>
        /// This is deprecated and has no effect. Do not use.
        /// </summary>
        public readonly string Sys;
        /// <summary>
        /// This is deprecated and has no effect. Do not use.
        /// </summary>
        public readonly ImmutableArray<string> Values;

        [OutputConstructor]
        private ConditionResponse(
            string iam,

            string op,

            string svc,

            string sys,

            ImmutableArray<string> values)
        {
            Iam = iam;
            Op = op;
            Svc = svc;
            Sys = sys;
            Values = values;
        }
    }
}
