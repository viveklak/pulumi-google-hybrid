// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Alpha.Inputs
{

    /// <summary>
    /// This is deprecated and has no effect. Do not use.
    /// </summary>
    public sealed class ConditionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// This is deprecated and has no effect. Do not use.
        /// </summary>
        [Input("iam")]
        public Input<Pulumi.GoogleNative.Compute.Alpha.ConditionIam>? Iam { get; set; }

        /// <summary>
        /// This is deprecated and has no effect. Do not use.
        /// </summary>
        [Input("op")]
        public Input<Pulumi.GoogleNative.Compute.Alpha.ConditionOp>? Op { get; set; }

        /// <summary>
        /// This is deprecated and has no effect. Do not use.
        /// </summary>
        [Input("svc")]
        public Input<string>? Svc { get; set; }

        /// <summary>
        /// This is deprecated and has no effect. Do not use.
        /// </summary>
        [Input("sys")]
        public Input<Pulumi.GoogleNative.Compute.Alpha.ConditionSys>? Sys { get; set; }

        [Input("values")]
        private InputList<string>? _values;

        /// <summary>
        /// This is deprecated and has no effect. Do not use.
        /// </summary>
        public InputList<string> Values
        {
            get => _values ?? (_values = new InputList<string>());
            set => _values = value;
        }

        public ConditionArgs()
        {
        }
    }
}
