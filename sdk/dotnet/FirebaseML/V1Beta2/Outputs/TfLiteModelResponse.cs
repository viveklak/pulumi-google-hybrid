// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.FirebaseML.V1Beta2.Outputs
{

    /// <summary>
    /// Information that is specific to TfLite models.
    /// </summary>
    [OutputType]
    public sealed class TfLiteModelResponse
    {
        /// <summary>
        /// The AutoML model id referencing a model you created with the AutoML API. The name should have format 'projects//locations//models/' (This is the model resource name returned from the AutoML API)
        /// </summary>
        public readonly string AutomlModel;
        /// <summary>
        /// The TfLite file containing the model. (Stored in Google Cloud). The gcs_tflite_uri should have form: gs://some-bucket/some-model.tflite Note: If you update the file in the original location, it is necessary to call UpdateModel for ML to pick up and validate the updated file.
        /// </summary>
        public readonly string GcsTfliteUri;
        /// <summary>
        /// The size of the TFLite model
        /// </summary>
        public readonly string SizeBytes;

        [OutputConstructor]
        private TfLiteModelResponse(
            string automlModel,

            string gcsTfliteUri,

            string sizeBytes)
        {
            AutomlModel = automlModel;
            GcsTfliteUri = gcsTfliteUri;
            SizeBytes = sizeBytes;
        }
    }
}
