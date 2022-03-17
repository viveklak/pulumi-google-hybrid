// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.ToolResults.V1Beta3.Inputs
{

    /// <summary>
    /// Details when multiple steps are run with the same configuration as a group.
    /// </summary>
    public sealed class MultiStepArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Unique int given to each step. Ranges from 0(inclusive) to total number of steps(exclusive). The primary step is 0.
        /// </summary>
        [Input("multistepNumber")]
        public Input<int>? MultistepNumber { get; set; }

        /// <summary>
        /// Present if it is a primary (original) step.
        /// </summary>
        [Input("primaryStep")]
        public Input<Inputs.PrimaryStepArgs>? PrimaryStep { get; set; }

        /// <summary>
        /// Step Id of the primary (original) step, which might be this step.
        /// </summary>
        [Input("primaryStepId")]
        public Input<string>? PrimaryStepId { get; set; }

        public MultiStepArgs()
        {
        }
    }
}
