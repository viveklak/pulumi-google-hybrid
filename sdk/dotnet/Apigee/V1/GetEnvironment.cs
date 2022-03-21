// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Apigee.V1
{
    public static class GetEnvironment
    {
        /// <summary>
        /// Gets environment details.
        /// </summary>
        public static Task<GetEnvironmentResult> InvokeAsync(GetEnvironmentArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetEnvironmentResult>("google-native:apigee/v1:getEnvironment", args ?? new GetEnvironmentArgs(), options.WithDefaults());

        /// <summary>
        /// Gets environment details.
        /// </summary>
        public static Output<GetEnvironmentResult> Invoke(GetEnvironmentInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetEnvironmentResult>("google-native:apigee/v1:getEnvironment", args ?? new GetEnvironmentInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetEnvironmentArgs : Pulumi.InvokeArgs
    {
        [Input("environmentId", required: true)]
        public string EnvironmentId { get; set; } = null!;

        [Input("organizationId", required: true)]
        public string OrganizationId { get; set; } = null!;

        public GetEnvironmentArgs()
        {
        }
    }

    public sealed class GetEnvironmentInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("environmentId", required: true)]
        public Input<string> EnvironmentId { get; set; } = null!;

        [Input("organizationId", required: true)]
        public Input<string> OrganizationId { get; set; } = null!;

        public GetEnvironmentInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetEnvironmentResult
    {
        /// <summary>
        /// Optional. API Proxy type supported by the environment. The type can be set when creating the Environment and cannot be changed.
        /// </summary>
        public readonly string ApiProxyType;
        /// <summary>
        /// Creation time of this environment as milliseconds since epoch.
        /// </summary>
        public readonly string CreatedAt;
        /// <summary>
        /// Optional. Deployment type supported by the environment. The deployment type can be set when creating the environment and cannot be changed. When you enable archive deployment, you will be **prevented from performing** a [subset of actions](/apigee/docs/api-platform/local-development/overview#prevented-actions) within the environment, including: * Managing the deployment of API proxy or shared flow revisions * Creating, updating, or deleting resource files * Creating, updating, or deleting target servers
        /// </summary>
        public readonly string DeploymentType;
        /// <summary>
        /// Optional. Description of the environment.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// Optional. Display name for this environment.
        /// </summary>
        public readonly string DisplayName;
        /// <summary>
        /// Last modification time of this environment as milliseconds since epoch.
        /// </summary>
        public readonly string LastModifiedAt;
        /// <summary>
        /// Name of the environment. Values must match the regular expression `^[.\\p{Alnum}-_]{1,255}$`
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Optional. Key-value pairs that may be used for customizing the environment.
        /// </summary>
        public readonly Outputs.GoogleCloudApigeeV1PropertiesResponse Properties;
        /// <summary>
        /// State of the environment. Values other than ACTIVE means the resource is not ready to use.
        /// </summary>
        public readonly string State;

        [OutputConstructor]
        private GetEnvironmentResult(
            string apiProxyType,

            string createdAt,

            string deploymentType,

            string description,

            string displayName,

            string lastModifiedAt,

            string name,

            Outputs.GoogleCloudApigeeV1PropertiesResponse properties,

            string state)
        {
            ApiProxyType = apiProxyType;
            CreatedAt = createdAt;
            DeploymentType = deploymentType;
            Description = description;
            DisplayName = displayName;
            LastModifiedAt = lastModifiedAt;
            Name = name;
            Properties = properties;
            State = state;
        }
    }
}
