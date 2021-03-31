// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleCloud.Serviceusage.V1beta1
{
    /// <summary>
    /// Creates a consumer override. A consumer override is applied to the consumer on its own authority to limit its own quota usage. Consumer overrides cannot be used to grant more quota than would be allowed by admin overrides, producer overrides, or the default limit of the service.
    /// </summary>
    [GoogleCloudResourceType("google-cloud:serviceusage/v1beta1:ServiceConsumerQuotaMetricLimitConsumerOverride")]
    public partial class ServiceConsumerQuotaMetricLimitConsumerOverride : Pulumi.CustomResource
    {
        /// <summary>
        /// Create a ServiceConsumerQuotaMetricLimitConsumerOverride resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ServiceConsumerQuotaMetricLimitConsumerOverride(string name, ServiceConsumerQuotaMetricLimitConsumerOverrideArgs args, CustomResourceOptions? options = null)
            : base("google-cloud:serviceusage/v1beta1:ServiceConsumerQuotaMetricLimitConsumerOverride", name, args ?? new ServiceConsumerQuotaMetricLimitConsumerOverrideArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ServiceConsumerQuotaMetricLimitConsumerOverride(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-cloud:serviceusage/v1beta1:ServiceConsumerQuotaMetricLimitConsumerOverride", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing ServiceConsumerQuotaMetricLimitConsumerOverride resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ServiceConsumerQuotaMetricLimitConsumerOverride Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ServiceConsumerQuotaMetricLimitConsumerOverride(name, id, options);
        }
    }

    public sealed class ServiceConsumerQuotaMetricLimitConsumerOverrideArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The resource name of the ancestor that requested the override. For example: "organizations/12345" or "folders/67890". Used by admin overrides only.
        /// </summary>
        [Input("adminOverrideAncestor")]
        public Input<string>? AdminOverrideAncestor { get; set; }

        [Input("consumerOverridesId", required: true)]
        public Input<string> ConsumerOverridesId { get; set; } = null!;

        [Input("consumerQuotaMetricsId", required: true)]
        public Input<string> ConsumerQuotaMetricsId { get; set; } = null!;

        [Input("dimensions")]
        private InputMap<string>? _dimensions;

        /// <summary>
        /// If this map is nonempty, then this override applies only to specific values for dimensions defined in the limit unit. For example, an override on a limit with the unit 1/{project}/{region} could contain an entry with the key "region" and the value "us-east-1"; the override is only applied to quota consumed in that region. This map has the following restrictions: * Keys that are not defined in the limit's unit are not valid keys. Any string appearing in {brackets} in the unit (besides {project} or {user}) is a defined key. * "project" is not a valid key; the project is already specified in the parent resource name. * "user" is not a valid key; the API does not support quota overrides that apply only to a specific user. * If "region" appears as a key, its value must be a valid Cloud region. * If "zone" appears as a key, its value must be a valid Cloud zone. * If any valid key other than "region" or "zone" appears in the map, then all valid keys other than "region" or "zone" must also appear in the map.
        /// </summary>
        public InputMap<string> Dimensions
        {
            get => _dimensions ?? (_dimensions = new InputMap<string>());
            set => _dimensions = value;
        }

        [Input("limitsId", required: true)]
        public Input<string> LimitsId { get; set; } = null!;

        /// <summary>
        /// The name of the metric to which this override applies. An example name would be: `compute.googleapis.com/cpus`
        /// </summary>
        [Input("metric")]
        public Input<string>? Metric { get; set; }

        /// <summary>
        /// The resource name of the override. This name is generated by the server when the override is created. Example names would be: `projects/123/services/compute.googleapis.com/consumerQuotaMetrics/compute.googleapis.com%2Fcpus/limits/%2Fproject%2Fregion/adminOverrides/4a3f2c1d` `projects/123/services/compute.googleapis.com/consumerQuotaMetrics/compute.googleapis.com%2Fcpus/limits/%2Fproject%2Fregion/consumerOverrides/4a3f2c1d` The resource name is intended to be opaque and should not be parsed for its component strings, since its representation could change in the future.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The overriding quota limit value. Can be any nonnegative integer, or -1 (unlimited quota).
        /// </summary>
        [Input("overrideValue")]
        public Input<string>? OverrideValue { get; set; }

        [Input("servicesId", required: true)]
        public Input<string> ServicesId { get; set; } = null!;

        /// <summary>
        /// The limit unit of the limit to which this override applies. An example unit would be: `1/{project}/{region}` Note that `{project}` and `{region}` are not placeholders in this example; the literal characters `{` and `}` occur in the string.
        /// </summary>
        [Input("unit")]
        public Input<string>? Unit { get; set; }

        [Input("v1beta1Id", required: true)]
        public Input<string> V1beta1Id { get; set; } = null!;

        [Input("v1beta1Id1", required: true)]
        public Input<string> V1beta1Id1 { get; set; } = null!;

        public ServiceConsumerQuotaMetricLimitConsumerOverrideArgs()
        {
        }
    }
}
