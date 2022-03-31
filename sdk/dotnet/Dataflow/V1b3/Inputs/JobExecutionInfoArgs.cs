// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dataflow.V1b3.Inputs
{

    /// <summary>
    /// Additional information about how a Cloud Dataflow job will be executed that isn't contained in the submitted job.
    /// </summary>
    public sealed class JobExecutionInfoArgs : Pulumi.ResourceArgs
    {
        [Input("stages")]
        private InputMap<string>? _stages;

        /// <summary>
        /// A mapping from each stage to the information about that stage.
        /// </summary>
        public InputMap<string> Stages
        {
            get => _stages ?? (_stages = new InputMap<string>());
            set => _stages = value;
        }

        public JobExecutionInfoArgs()
        {
        }
    }
}
