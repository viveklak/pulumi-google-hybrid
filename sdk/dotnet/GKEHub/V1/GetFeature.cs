// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.GKEHub.V1
{
    public static class GetFeature
    {
        /// <summary>
        /// Gets details of a single Feature.
        /// </summary>
        public static Task<GetFeatureResult> InvokeAsync(GetFeatureArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetFeatureResult>("google-native:gkehub/v1:getFeature", args ?? new GetFeatureArgs(), options.WithDefaults());

        /// <summary>
        /// Gets details of a single Feature.
        /// </summary>
        public static Output<GetFeatureResult> Invoke(GetFeatureInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetFeatureResult>("google-native:gkehub/v1:getFeature", args ?? new GetFeatureInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetFeatureArgs : Pulumi.InvokeArgs
    {
        [Input("featureId", required: true)]
        public string FeatureId { get; set; } = null!;

        [Input("location", required: true)]
        public string Location { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        public GetFeatureArgs()
        {
        }
    }

    public sealed class GetFeatureInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("featureId", required: true)]
        public Input<string> FeatureId { get; set; } = null!;

        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        public GetFeatureInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetFeatureResult
    {
        /// <summary>
        /// When the Feature resource was created.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// When the Feature resource was deleted.
        /// </summary>
        public readonly string DeleteTime;
        /// <summary>
        /// GCP labels for this Feature.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Labels;
        /// <summary>
        /// Optional. Membership-specific configuration for this Feature. If this Feature does not support any per-Membership configuration, this field may be unused. The keys indicate which Membership the configuration is for, in the form: `projects/{p}/locations/{l}/memberships/{m}` Where {p} is the project, {l} is a valid location and {m} is a valid Membership in this project at that location. {p} WILL match the Feature's project. {p} will always be returned as the project number, but the project ID is also accepted during input. If the same Membership is specified in the map twice (using the project ID form, and the project number form), exactly ONE of the entries will be saved, with no guarantees as to which. For this reason, it is recommended the same format be used for all entries when mutating a Feature.
        /// </summary>
        public readonly ImmutableDictionary<string, string> MembershipSpecs;
        /// <summary>
        /// Membership-specific Feature status. If this Feature does report any per-Membership status, this field may be unused. The keys indicate which Membership the state is for, in the form: `projects/{p}/locations/{l}/memberships/{m}` Where {p} is the project number, {l} is a valid location and {m} is a valid Membership in this project at that location. {p} MUST match the Feature's project number.
        /// </summary>
        public readonly ImmutableDictionary<string, string> MembershipStates;
        /// <summary>
        /// The full, unique name of this Feature resource in the format `projects/*/locations/*/features/*`.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// State of the Feature resource itself.
        /// </summary>
        public readonly Outputs.FeatureResourceStateResponse ResourceState;
        /// <summary>
        /// Optional. Hub-wide Feature configuration. If this Feature does not support any Hub-wide configuration, this field may be unused.
        /// </summary>
        public readonly Outputs.CommonFeatureSpecResponse Spec;
        /// <summary>
        /// The Hub-wide Feature state.
        /// </summary>
        public readonly Outputs.CommonFeatureStateResponse State;
        /// <summary>
        /// When the Feature resource was last updated.
        /// </summary>
        public readonly string UpdateTime;

        [OutputConstructor]
        private GetFeatureResult(
            string createTime,

            string deleteTime,

            ImmutableDictionary<string, string> labels,

            ImmutableDictionary<string, string> membershipSpecs,

            ImmutableDictionary<string, string> membershipStates,

            string name,

            Outputs.FeatureResourceStateResponse resourceState,

            Outputs.CommonFeatureSpecResponse spec,

            Outputs.CommonFeatureStateResponse state,

            string updateTime)
        {
            CreateTime = createTime;
            DeleteTime = deleteTime;
            Labels = labels;
            MembershipSpecs = membershipSpecs;
            MembershipStates = membershipStates;
            Name = name;
            ResourceState = resourceState;
            Spec = spec;
            State = state;
            UpdateTime = updateTime;
        }
    }
}
