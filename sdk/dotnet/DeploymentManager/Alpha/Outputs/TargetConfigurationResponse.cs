// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.DeploymentManager.Alpha.Outputs
{

    [OutputType]
    public sealed class TargetConfigurationResponse
    {
        /// <summary>
        /// The configuration to use for this deployment.
        /// </summary>
        public readonly Outputs.ConfigFileResponse Config;
        /// <summary>
        /// Specifies any files to import for this configuration. This can be used to import templates or other files. For example, you might import a text file in order to use the file in a template.
        /// </summary>
        public readonly ImmutableArray<Outputs.ImportFileResponse> Imports;

        [OutputConstructor]
        private TargetConfigurationResponse(
            Outputs.ConfigFileResponse config,

            ImmutableArray<Outputs.ImportFileResponse> imports)
        {
            Config = config;
            Imports = imports;
        }
    }
}
