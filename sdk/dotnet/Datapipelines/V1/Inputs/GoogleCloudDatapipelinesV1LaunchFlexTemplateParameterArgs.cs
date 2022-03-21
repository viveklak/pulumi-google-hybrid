// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Datapipelines.V1.Inputs
{

    /// <summary>
    /// Launch Flex Template parameter.
    /// </summary>
    public sealed class GoogleCloudDatapipelinesV1LaunchFlexTemplateParameterArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Cloud Storage path to a file with a JSON-serialized ContainerSpec as content.
        /// </summary>
        [Input("containerSpecGcsPath")]
        public Input<string>? ContainerSpecGcsPath { get; set; }

        /// <summary>
        /// The runtime environment for the Flex Template job.
        /// </summary>
        [Input("environment")]
        public Input<Inputs.GoogleCloudDatapipelinesV1FlexTemplateRuntimeEnvironmentArgs>? Environment { get; set; }

        /// <summary>
        /// The job name to use for the created job. For an update job request, the job name should be the same as the existing running job.
        /// </summary>
        [Input("jobName", required: true)]
        public Input<string> JobName { get; set; } = null!;

        [Input("launchOptions")]
        private InputMap<string>? _launchOptions;

        /// <summary>
        /// Launch options for this Flex Template job. This is a common set of options across languages and templates. This should not be used to pass job parameters.
        /// </summary>
        public InputMap<string> LaunchOptions
        {
            get => _launchOptions ?? (_launchOptions = new InputMap<string>());
            set => _launchOptions = value;
        }

        [Input("parameters")]
        private InputMap<string>? _parameters;

        /// <summary>
        /// The parameters for the Flex Template. Example: `{"num_workers":"5"}`
        /// </summary>
        public InputMap<string> Parameters
        {
            get => _parameters ?? (_parameters = new InputMap<string>());
            set => _parameters = value;
        }

        [Input("transformNameMappings")]
        private InputMap<string>? _transformNameMappings;

        /// <summary>
        /// Use this to pass transform name mappings for streaming update jobs. Example: `{"oldTransformName":"newTransformName",...}`
        /// </summary>
        public InputMap<string> TransformNameMappings
        {
            get => _transformNameMappings ?? (_transformNameMappings = new InputMap<string>());
            set => _transformNameMappings = value;
        }

        /// <summary>
        /// Set this to true if you are sending a request to update a running streaming job. When set, the job name should be the same as the running job.
        /// </summary>
        [Input("update")]
        public Input<bool>? Update { get; set; }

        public GoogleCloudDatapipelinesV1LaunchFlexTemplateParameterArgs()
        {
        }
    }
}
