// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Ml.V1.Outputs
{

    /// <summary>
    /// Represents a set of hyperparameters to optimize.
    /// </summary>
    [OutputType]
    public sealed class GoogleCloudMlV1__HyperparameterSpecResponse
    {
        /// <summary>
        /// Optional. The search algorithm specified for the hyperparameter tuning job. Uses the default AI Platform hyperparameter tuning algorithm if unspecified.
        /// </summary>
        public readonly string Algorithm;
        /// <summary>
        /// Optional. Indicates if the hyperparameter tuning job enables auto trial early stopping.
        /// </summary>
        public readonly bool EnableTrialEarlyStopping;
        /// <summary>
        /// The type of goal to use for tuning. Available types are `MAXIMIZE` and `MINIMIZE`. Defaults to `MAXIMIZE`.
        /// </summary>
        public readonly string Goal;
        /// <summary>
        /// Optional. The TensorFlow summary tag name to use for optimizing trials. For current versions of TensorFlow, this tag name should exactly match what is shown in TensorBoard, including all scopes. For versions of TensorFlow prior to 0.12, this should be only the tag passed to tf.Summary. By default, "training/hptuning/metric" will be used.
        /// </summary>
        public readonly string HyperparameterMetricTag;
        /// <summary>
        /// Optional. The number of failed trials that need to be seen before failing the hyperparameter tuning job. You can specify this field to override the default failing criteria for AI Platform hyperparameter tuning jobs. Defaults to zero, which means the service decides when a hyperparameter job should fail.
        /// </summary>
        public readonly int MaxFailedTrials;
        /// <summary>
        /// Optional. The number of training trials to run concurrently. You can reduce the time it takes to perform hyperparameter tuning by adding trials in parallel. However, each trail only benefits from the information gained in completed trials. That means that a trial does not get access to the results of trials running at the same time, which could reduce the quality of the overall optimization. Each trial will use the same scale tier and machine types. Defaults to one.
        /// </summary>
        public readonly int MaxParallelTrials;
        /// <summary>
        /// Optional. How many training trials should be attempted to optimize the specified hyperparameters. Defaults to one.
        /// </summary>
        public readonly int MaxTrials;
        /// <summary>
        /// The set of parameters to tune.
        /// </summary>
        public readonly ImmutableArray<Outputs.GoogleCloudMlV1__ParameterSpecResponse> Params;
        /// <summary>
        /// Optional. The prior hyperparameter tuning job id that users hope to continue with. The job id will be used to find the corresponding vizier study guid and resume the study.
        /// </summary>
        public readonly string ResumePreviousJobId;

        [OutputConstructor]
        private GoogleCloudMlV1__HyperparameterSpecResponse(
            string algorithm,

            bool enableTrialEarlyStopping,

            string goal,

            string hyperparameterMetricTag,

            int maxFailedTrials,

            int maxParallelTrials,

            int maxTrials,

            ImmutableArray<Outputs.GoogleCloudMlV1__ParameterSpecResponse> @params,

            string resumePreviousJobId)
        {
            Algorithm = algorithm;
            EnableTrialEarlyStopping = enableTrialEarlyStopping;
            Goal = goal;
            HyperparameterMetricTag = hyperparameterMetricTag;
            MaxFailedTrials = maxFailedTrials;
            MaxParallelTrials = maxParallelTrials;
            MaxTrials = maxTrials;
            Params = @params;
            ResumePreviousJobId = resumePreviousJobId;
        }
    }
}
