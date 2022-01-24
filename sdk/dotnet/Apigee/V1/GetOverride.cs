// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi.Utilities;

namespace Pulumi.GoogleNative.Apigee.V1
{
    public static class GetOverride
    {
        /// <summary>
        /// Gets a trace configuration override.
        /// </summary>
        public static Task<GetOverrideResult> InvokeAsync(GetOverrideArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetOverrideResult>("google-native:apigee/v1:getOverride", args ?? new GetOverrideArgs(), options.WithDefaults());

        /// <summary>
        /// Gets a trace configuration override.
        /// </summary>
        public static Output<GetOverrideResult> Invoke(GetOverrideInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetOverrideResult>("google-native:apigee/v1:getOverride", args ?? new GetOverrideInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetOverrideArgs : Pulumi.InvokeArgs
    {
        [Input("environmentId", required: true)]
        public string EnvironmentId { get; set; } = null!;

        [Input("organizationId", required: true)]
        public string OrganizationId { get; set; } = null!;

        [Input("overrideId", required: true)]
        public string OverrideId { get; set; } = null!;

        public GetOverrideArgs()
        {
        }
    }

    public sealed class GetOverrideInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("environmentId", required: true)]
        public Input<string> EnvironmentId { get; set; } = null!;

        [Input("organizationId", required: true)]
        public Input<string> OrganizationId { get; set; } = null!;

        [Input("overrideId", required: true)]
        public Input<string> OverrideId { get; set; } = null!;

        public GetOverrideInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetOverrideResult
    {
        /// <summary>
        /// ID of the API proxy that will have its trace configuration overridden.
        /// </summary>
        public readonly string ApiProxy;
        /// <summary>
        /// ID of the trace configuration override specified as a system-generated UUID.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Trace configuration to override.
        /// </summary>
        public readonly Outputs.GoogleCloudApigeeV1TraceSamplingConfigResponse SamplingConfig;

        [OutputConstructor]
        private GetOverrideResult(
            string apiProxy,

            string name,

            Outputs.GoogleCloudApigeeV1TraceSamplingConfigResponse samplingConfig)
        {
            ApiProxy = apiProxy;
            Name = name;
            SamplingConfig = samplingConfig;
        }
    }
}
