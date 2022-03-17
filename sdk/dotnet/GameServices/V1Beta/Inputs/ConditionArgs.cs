// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.GameServices.V1Beta.Inputs
{

    /// <summary>
    /// A condition to be met.
    /// </summary>
    public sealed class ConditionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Trusted attributes supplied by the IAM system.
        /// </summary>
        [Input("iam")]
        public Input<Pulumi.GoogleNative.GameServices.V1Beta.ConditionIam>? Iam { get; set; }

        /// <summary>
        /// An operator to apply the subject with.
        /// </summary>
        [Input("op")]
        public Input<Pulumi.GoogleNative.GameServices.V1Beta.ConditionOp>? Op { get; set; }

        /// <summary>
        /// Trusted attributes discharged by the service.
        /// </summary>
        [Input("svc")]
        public Input<string>? Svc { get; set; }

        /// <summary>
        /// Trusted attributes supplied by any service that owns resources and uses the IAM system for access control.
        /// </summary>
        [Input("sys")]
        public Input<Pulumi.GoogleNative.GameServices.V1Beta.ConditionSys>? Sys { get; set; }

        [Input("values")]
        private InputList<string>? _values;

        /// <summary>
        /// The objects of the condition.
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
