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
    /// Options for automatically scaling a model.
    /// </summary>
    [OutputType]
    public sealed class GoogleCloudMlV1__AutoScalingResponse
    {
        /// <summary>
        /// The maximum number of nodes to scale this model under load. The actual value will depend on resource quota and availability.
        /// </summary>
        public readonly int MaxNodes;
        /// <summary>
        /// MetricSpec contains the specifications to use to calculate the desired nodes count.
        /// </summary>
        public readonly ImmutableArray<Outputs.GoogleCloudMlV1__MetricSpecResponse> Metrics;
        /// <summary>
        /// Optional. The minimum number of nodes to allocate for this model. These nodes are always up, starting from the time the model is deployed. Therefore, the cost of operating this model will be at least `rate` * `min_nodes` * number of hours since last billing cycle, where `rate` is the cost per node-hour as documented in the [pricing guide](/ml-engine/docs/pricing), even if no predictions are performed. There is additional cost for each prediction performed. Unlike manual scaling, if the load gets too heavy for the nodes that are up, the service will automatically add nodes to handle the increased load as well as scale back as traffic drops, always maintaining at least `min_nodes`. You will be charged for the time in which additional nodes are used. If `min_nodes` is not specified and AutoScaling is used with a [legacy (MLS1) machine type](/ml-engine/docs/machine-types-online-prediction), `min_nodes` defaults to 0, in which case, when traffic to a model stops (and after a cool-down period), nodes will be shut down and no charges will be incurred until traffic to the model resumes. If `min_nodes` is not specified and AutoScaling is used with a [Compute Engine (N1) machine type](/ml-engine/docs/machine-types-online-prediction), `min_nodes` defaults to 1. `min_nodes` must be at least 1 for use with a Compute Engine machine type. You can set `min_nodes` when creating the model version, and you can also update `min_nodes` for an existing version: update_body.json: { 'autoScaling': { 'minNodes': 5 } } HTTP request: PATCH https://ml.googleapis.com/v1/{name=projects/*/models/*/versions/*}?update_mask=autoScaling.minNodes -d @./update_body.json 
        /// </summary>
        public readonly int MinNodes;

        [OutputConstructor]
        private GoogleCloudMlV1__AutoScalingResponse(
            int maxNodes,

            ImmutableArray<Outputs.GoogleCloudMlV1__MetricSpecResponse> metrics,

            int minNodes)
        {
            MaxNodes = maxNodes;
            Metrics = metrics;
            MinNodes = minNodes;
        }
    }
}
