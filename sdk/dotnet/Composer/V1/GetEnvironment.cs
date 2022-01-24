// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi.Utilities;

namespace Pulumi.GoogleNative.Composer.V1
{
    public static class GetEnvironment
    {
        /// <summary>
        /// Get an existing environment.
        /// </summary>
        public static Task<GetEnvironmentResult> InvokeAsync(GetEnvironmentArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetEnvironmentResult>("google-native:composer/v1:getEnvironment", args ?? new GetEnvironmentArgs(), options.WithDefaults());

        /// <summary>
        /// Get an existing environment.
        /// </summary>
        public static Output<GetEnvironmentResult> Invoke(GetEnvironmentInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetEnvironmentResult>("google-native:composer/v1:getEnvironment", args ?? new GetEnvironmentInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetEnvironmentArgs : Pulumi.InvokeArgs
    {
        [Input("environmentId", required: true)]
        public string EnvironmentId { get; set; } = null!;

        [Input("location", required: true)]
        public string Location { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        public GetEnvironmentArgs()
        {
        }
    }

    public sealed class GetEnvironmentInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("environmentId", required: true)]
        public Input<string> EnvironmentId { get; set; } = null!;

        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        public GetEnvironmentInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetEnvironmentResult
    {
        /// <summary>
        /// Configuration parameters for this environment.
        /// </summary>
        public readonly Outputs.EnvironmentConfigResponse Config;
        /// <summary>
        /// The time at which this environment was created.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// Optional. User-defined labels for this environment. The labels map can contain no more than 64 entries. Entries of the labels map are UTF8 strings that comply with the following restrictions: * Keys must conform to regexp: \p{Ll}\p{Lo}{0,62} * Values must conform to regexp: [\p{Ll}\p{Lo}\p{N}_-]{0,63} * Both keys and values are additionally constrained to be &lt;= 128 bytes in size.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Labels;
        /// <summary>
        /// The resource name of the environment, in the form: "projects/{projectId}/locations/{locationId}/environments/{environmentId}" EnvironmentId must start with a lowercase letter followed by up to 63 lowercase letters, numbers, or hyphens, and cannot end with a hyphen.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The current state of the environment.
        /// </summary>
        public readonly string State;
        /// <summary>
        /// The time at which this environment was last modified.
        /// </summary>
        public readonly string UpdateTime;
        /// <summary>
        /// The UUID (Universally Unique IDentifier) associated with this environment. This value is generated when the environment is created.
        /// </summary>
        public readonly string Uuid;

        [OutputConstructor]
        private GetEnvironmentResult(
            Outputs.EnvironmentConfigResponse config,

            string createTime,

            ImmutableDictionary<string, string> labels,

            string name,

            string state,

            string updateTime,

            string uuid)
        {
            Config = config;
            CreateTime = createTime;
            Labels = labels;
            Name = name;
            State = state;
            UpdateTime = updateTime;
            Uuid = uuid;
        }
    }
}
