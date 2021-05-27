// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.FirebaseML.V1Beta2
{
    /// <summary>
    /// Creates a model in Firebase ML. The longrunning operation will eventually return a Model
    /// </summary>
    [GoogleNativeResourceType("google-native:firebaseml/v1beta2:Model")]
    public partial class Model : Pulumi.CustomResource
    {
        /// <summary>
        /// Lists operation ids associated with this model whose status is NOT done.
        /// </summary>
        [Output("activeOperations")]
        public Output<ImmutableArray<Outputs.OperationResponse>> ActiveOperations { get; private set; } = null!;

        /// <summary>
        /// Timestamp when this model was created in Firebase ML.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// Required. The name of the model to create. The name can be up to 32 characters long and can consist only of ASCII Latin letters A-Z and a-z, underscores(_) and ASCII digits 0-9. It must start with a letter.
        /// </summary>
        [Output("displayName")]
        public Output<string> DisplayName { get; private set; } = null!;

        /// <summary>
        /// See RFC7232 https://tools.ietf.org/html/rfc7232#section-2.3
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        /// <summary>
        /// The model_hash will change if a new file is available for download.
        /// </summary>
        [Output("modelHash")]
        public Output<string> ModelHash { get; private set; } = null!;

        /// <summary>
        /// The resource name of the Model. Model names have the form `projects/{project_id}/models/{model_id}` The name is ignored when creating a model.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// State common to all model types. Includes publishing and validation information.
        /// </summary>
        [Output("state")]
        public Output<Outputs.ModelStateResponse> State { get; private set; } = null!;

        /// <summary>
        /// User defined tags which can be used to group/filter models during listing
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<string>> Tags { get; private set; } = null!;

        /// <summary>
        /// A TFLite Model
        /// </summary>
        [Output("tfliteModel")]
        public Output<Outputs.TfLiteModelResponse> TfliteModel { get; private set; } = null!;

        /// <summary>
        /// Timestamp when this model was updated in Firebase ML.
        /// </summary>
        [Output("updateTime")]
        public Output<string> UpdateTime { get; private set; } = null!;


        /// <summary>
        /// Create a Model resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Model(string name, ModelArgs args, CustomResourceOptions? options = null)
            : base("google-native:firebaseml/v1beta2:Model", name, args ?? new ModelArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Model(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:firebaseml/v1beta2:Model", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing Model resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Model Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Model(name, id, options);
        }
    }

    public sealed class ModelArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Required. The name of the model to create. The name can be up to 32 characters long and can consist only of ASCII Latin letters A-Z and a-z, underscores(_) and ASCII digits 0-9. It must start with a letter.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// The resource name of the Model. Model names have the form `projects/{project_id}/models/{model_id}` The name is ignored when creating a model.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("project", required: true)]
        public Input<string> Project { get; set; } = null!;

        /// <summary>
        /// State common to all model types. Includes publishing and validation information.
        /// </summary>
        [Input("state")]
        public Input<Inputs.ModelStateArgs>? State { get; set; }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// User defined tags which can be used to group/filter models during listing
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        /// <summary>
        /// A TFLite Model
        /// </summary>
        [Input("tfliteModel")]
        public Input<Inputs.TfLiteModelArgs>? TfliteModel { get; set; }

        public ModelArgs()
        {
        }
    }
}
