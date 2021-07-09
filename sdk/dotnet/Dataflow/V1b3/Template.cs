// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dataflow.V1b3
{
    /// <summary>
    /// Creates a Cloud Dataflow job from a template.
    /// </summary>
    [GoogleNativeResourceType("google-native:dataflow/v1b3:Template")]
    public partial class Template : Pulumi.CustomResource
    {
        /// <summary>
        /// The template metadata describing the template name, available parameters, etc.
        /// </summary>
        [Output("metadata")]
        public Output<Outputs.TemplateMetadataResponse> Metadata { get; private set; } = null!;

        /// <summary>
        /// Describes the runtime metadata with SDKInfo and available parameters.
        /// </summary>
        [Output("runtimeMetadata")]
        public Output<Outputs.RuntimeMetadataResponse> RuntimeMetadata { get; private set; } = null!;

        /// <summary>
        /// The status of the get template request. Any problems with the request will be indicated in the error_details.
        /// </summary>
        [Output("status")]
        public Output<Outputs.StatusResponse> Status { get; private set; } = null!;

        /// <summary>
        /// Template Type.
        /// </summary>
        [Output("templateType")]
        public Output<string> TemplateType { get; private set; } = null!;


        /// <summary>
        /// Create a Template resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Template(string name, TemplateArgs args, CustomResourceOptions? options = null)
            : base("google-native:dataflow/v1b3:Template", name, args ?? new TemplateArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Template(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:dataflow/v1b3:Template", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Template resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Template Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Template(name, id, options);
        }
    }

    public sealed class TemplateArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The runtime environment for the job.
        /// </summary>
        [Input("environment")]
        public Input<Inputs.RuntimeEnvironmentArgs>? Environment { get; set; }

        /// <summary>
        /// A Cloud Storage path to the template from which to create the job. Must be a valid Cloud Storage URL, beginning with `gs://`.
        /// </summary>
        [Input("gcsPath", required: true)]
        public Input<string> GcsPath { get; set; } = null!;

        /// <summary>
        /// The job name to use for the created job.
        /// </summary>
        [Input("jobName", required: true)]
        public Input<string> JobName { get; set; } = null!;

        /// <summary>
        /// The [regional endpoint] (https://cloud.google.com/dataflow/docs/concepts/regional-endpoints) to which to direct the request.
        /// </summary>
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("parameters")]
        private InputMap<string>? _parameters;

        /// <summary>
        /// The runtime parameters to pass to the job.
        /// </summary>
        public InputMap<string> Parameters
        {
            get => _parameters ?? (_parameters = new InputMap<string>());
            set => _parameters = value;
        }

        [Input("project", required: true)]
        public Input<string> Project { get; set; } = null!;

        public TemplateArgs()
        {
        }
    }
}
