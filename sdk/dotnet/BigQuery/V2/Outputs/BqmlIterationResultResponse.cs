// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.BigQuery.V2.Outputs
{

    [OutputType]
    public sealed class BqmlIterationResultResponse
    {
        /// <summary>
        /// [Output-only, Beta] Time taken to run the training iteration in milliseconds.
        /// </summary>
        public readonly string DurationMs;
        /// <summary>
        /// [Output-only, Beta] Eval loss computed on the eval data at the end of the iteration. The eval loss is used for early stopping to avoid overfitting. No eval loss if eval_split_method option is specified as no_split or auto_split with input data size less than 500 rows.
        /// </summary>
        public readonly double EvalLoss;
        /// <summary>
        /// [Output-only, Beta] Index of the ML training iteration, starting from zero for each training run.
        /// </summary>
        public readonly int Index;
        /// <summary>
        /// [Output-only, Beta] Learning rate used for this iteration, it varies for different training iterations if learn_rate_strategy option is not constant.
        /// </summary>
        public readonly double LearnRate;
        /// <summary>
        /// [Output-only, Beta] Training loss computed on the training data at the end of the iteration. The training loss function is defined by model type.
        /// </summary>
        public readonly double TrainingLoss;

        [OutputConstructor]
        private BqmlIterationResultResponse(
            string durationMs,

            double evalLoss,

            int index,

            double learnRate,

            double trainingLoss)
        {
            DurationMs = durationMs;
            EvalLoss = evalLoss;
            Index = index;
            LearnRate = learnRate;
            TrainingLoss = trainingLoss;
        }
    }
}
