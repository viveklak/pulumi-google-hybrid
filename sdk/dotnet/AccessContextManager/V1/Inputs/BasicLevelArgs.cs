// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.AccessContextManager.V1.Inputs
{

    /// <summary>
    /// `BasicLevel` is an `AccessLevel` using a set of recommended features.
    /// </summary>
    public sealed class BasicLevelArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// How the `conditions` list should be combined to determine if a request is granted this `AccessLevel`. If AND is used, each `Condition` in `conditions` must be satisfied for the `AccessLevel` to be applied. If OR is used, at least one `Condition` in `conditions` must be satisfied for the `AccessLevel` to be applied. Default behavior is AND.
        /// </summary>
        [Input("combiningFunction")]
        public Input<Pulumi.GoogleNative.AccessContextManager.V1.BasicLevelCombiningFunction>? CombiningFunction { get; set; }

        [Input("conditions", required: true)]
        private InputList<Inputs.ConditionArgs>? _conditions;

        /// <summary>
        /// A list of requirements for the `AccessLevel` to be granted.
        /// </summary>
        public InputList<Inputs.ConditionArgs> Conditions
        {
            get => _conditions ?? (_conditions = new InputList<Inputs.ConditionArgs>());
            set => _conditions = value;
        }

        public BasicLevelArgs()
        {
        }
    }
}
