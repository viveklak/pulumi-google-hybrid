// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.FirebaseML.V1Beta2
{
    public static class GetModel
    {
        /// <summary>
        /// Gets a model resource.
        /// </summary>
        public static Task<GetModelResult> InvokeAsync(GetModelArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetModelResult>("google-native:firebaseml/v1beta2:getModel", args ?? new GetModelArgs(), options.WithDefaults());

        /// <summary>
        /// Gets a model resource.
        /// </summary>
        public static Output<GetModelResult> Invoke(GetModelInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetModelResult>("google-native:firebaseml/v1beta2:getModel", args ?? new GetModelInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetModelArgs : Pulumi.InvokeArgs
    {
        [Input("modelId", required: true)]
        public string ModelId { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        public GetModelArgs()
        {
        }
    }

    public sealed class GetModelInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("modelId", required: true)]
        public Input<string> ModelId { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        public GetModelInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetModelResult
    {
        /// <summary>
        /// Lists operation ids associated with this model whose status is NOT done.
        /// </summary>
        public readonly ImmutableArray<Outputs.OperationResponse> ActiveOperations;
        /// <summary>
        /// Timestamp when this model was created in Firebase ML.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// The name of the model to create. The name can be up to 32 characters long and can consist only of ASCII Latin letters A-Z and a-z, underscores(_) and ASCII digits 0-9. It must start with a letter.
        /// </summary>
        public readonly string DisplayName;
        /// <summary>
        /// See RFC7232 https://tools.ietf.org/html/rfc7232#section-2.3
        /// </summary>
        public readonly string Etag;
        /// <summary>
        /// The model_hash will change if a new file is available for download.
        /// </summary>
        public readonly string ModelHash;
        /// <summary>
        /// The resource name of the Model. Model names have the form `projects/{project_id}/models/{model_id}` The name is ignored when creating a model.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// State common to all model types. Includes publishing and validation information.
        /// </summary>
        public readonly Outputs.ModelStateResponse State;
        /// <summary>
        /// User defined tags which can be used to group/filter models during listing
        /// </summary>
        public readonly ImmutableArray<string> Tags;
        /// <summary>
        /// A TFLite Model
        /// </summary>
        public readonly Outputs.TfLiteModelResponse TfliteModel;
        /// <summary>
        /// Timestamp when this model was updated in Firebase ML.
        /// </summary>
        public readonly string UpdateTime;

        [OutputConstructor]
        private GetModelResult(
            ImmutableArray<Outputs.OperationResponse> activeOperations,

            string createTime,

            string displayName,

            string etag,

            string modelHash,

            string name,

            Outputs.ModelStateResponse state,

            ImmutableArray<string> tags,

            Outputs.TfLiteModelResponse tfliteModel,

            string updateTime)
        {
            ActiveOperations = activeOperations;
            CreateTime = createTime;
            DisplayName = displayName;
            Etag = etag;
            ModelHash = modelHash;
            Name = name;
            State = state;
            Tags = tags;
            TfliteModel = tfliteModel;
            UpdateTime = updateTime;
        }
    }
}
