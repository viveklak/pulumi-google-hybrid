// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dialogflow.V3.Outputs
{

    /// <summary>
    /// Settings related to NLU.
    /// </summary>
    [OutputType]
    public sealed class GoogleCloudDialogflowCxV3NluSettingsResponse
    {
        /// <summary>
        /// To filter out false positive results and still get variety in matched natural language inputs for your agent, you can tune the machine learning classification threshold. If the returned score value is less than the threshold value, then a no-match event will be triggered. The score values range from 0.0 (completely uncertain) to 1.0 (completely certain). If set to 0.0, the default of 0.3 is used.
        /// </summary>
        public readonly double ClassificationThreshold;
        /// <summary>
        /// Indicates NLU model training mode.
        /// </summary>
        public readonly string ModelTrainingMode;
        /// <summary>
        /// Indicates the type of NLU model.
        /// </summary>
        public readonly string ModelType;

        [OutputConstructor]
        private GoogleCloudDialogflowCxV3NluSettingsResponse(
            double classificationThreshold,

            string modelTrainingMode,

            string modelType)
        {
            ClassificationThreshold = classificationThreshold;
            ModelTrainingMode = modelTrainingMode;
            ModelType = modelType;
        }
    }
}
