// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.DeploymentManager.V2.Inputs
{

    public sealed class TargetConfigurationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The configuration to use for this deployment.
        /// </summary>
        [Input("config")]
        public Input<Inputs.ConfigFileArgs>? Config { get; set; }

        [Input("imports")]
        private InputList<Inputs.ImportFileArgs>? _imports;

        /// <summary>
        /// Specifies any files to import for this configuration. This can be used to import templates or other files. For example, you might import a text file in order to use the file in a template.
        /// </summary>
        public InputList<Inputs.ImportFileArgs> Imports
        {
            get => _imports ?? (_imports = new InputList<Inputs.ImportFileArgs>());
            set => _imports = value;
        }

        public TargetConfigurationArgs()
        {
        }
    }
}
