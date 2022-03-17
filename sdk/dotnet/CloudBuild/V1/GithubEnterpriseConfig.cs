// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.CloudBuild.V1
{
    /// <summary>
    /// Create an association between a GCP project and a GitHub Enterprise server.
    /// </summary>
    [GoogleNativeResourceType("google-native:cloudbuild/v1:GithubEnterpriseConfig")]
    public partial class GithubEnterpriseConfig : Pulumi.CustomResource
    {
        /// <summary>
        /// The GitHub app id of the Cloud Build app on the GitHub Enterprise server.
        /// </summary>
        [Output("appId")]
        public Output<string> AppId { get; private set; } = null!;

        /// <summary>
        /// Time when the installation was associated with the project.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// Name to display for this config.
        /// </summary>
        [Output("displayName")]
        public Output<string> DisplayName { get; private set; } = null!;

        /// <summary>
        /// The URL of the github enterprise host the configuration is for.
        /// </summary>
        [Output("hostUrl")]
        public Output<string> HostUrl { get; private set; } = null!;

        /// <summary>
        /// Optional. The full resource name for the GitHubEnterpriseConfig For example: "projects/{$project_id}/githubEnterpriseConfigs/{$config_id}"
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Optional. The network to be used when reaching out to the GitHub Enterprise server. The VPC network must be enabled for private service connection. This should be set if the GitHub Enterprise server is hosted on-premises and not reachable by public internet. If this field is left empty, no network peering will occur and calls to the GitHub Enterprise server will be made over the public internet. Must be in the format `projects/{project}/global/networks/{network}`, where {project} is a project number or id and {network} is the name of a VPC network in the project.
        /// </summary>
        [Output("peeredNetwork")]
        public Output<string> PeeredNetwork { get; private set; } = null!;

        /// <summary>
        /// Names of secrets in Secret Manager.
        /// </summary>
        [Output("secrets")]
        public Output<Outputs.GitHubEnterpriseSecretsResponse> Secrets { get; private set; } = null!;

        /// <summary>
        /// Optional. SSL certificate to use for requests to GitHub Enterprise.
        /// </summary>
        [Output("sslCa")]
        public Output<string> SslCa { get; private set; } = null!;

        /// <summary>
        /// The key that should be attached to webhook calls to the ReceiveWebhook endpoint.
        /// </summary>
        [Output("webhookKey")]
        public Output<string> WebhookKey { get; private set; } = null!;


        /// <summary>
        /// Create a GithubEnterpriseConfig resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public GithubEnterpriseConfig(string name, GithubEnterpriseConfigArgs args, CustomResourceOptions? options = null)
            : base("google-native:cloudbuild/v1:GithubEnterpriseConfig", name, args ?? new GithubEnterpriseConfigArgs(), MakeResourceOptions(options, ""))
        {
        }

        private GithubEnterpriseConfig(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:cloudbuild/v1:GithubEnterpriseConfig", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing GithubEnterpriseConfig resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static GithubEnterpriseConfig Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new GithubEnterpriseConfig(name, id, options);
        }
    }

    public sealed class GithubEnterpriseConfigArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The GitHub app id of the Cloud Build app on the GitHub Enterprise server.
        /// </summary>
        [Input("appId", required: true)]
        public Input<string> AppId { get; set; } = null!;

        /// <summary>
        /// Name to display for this config.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// Optional. The ID to use for the GithubEnterpriseConfig, which will become the final component of the GithubEnterpriseConfig’s resource name. ghe_config_id must meet the following requirements: + They must contain only alphanumeric characters and dashes. + They can be 1-64 characters long. + They must begin and end with an alphanumeric character
        /// </summary>
        [Input("gheConfigId")]
        public Input<string>? GheConfigId { get; set; }

        /// <summary>
        /// The URL of the github enterprise host the configuration is for.
        /// </summary>
        [Input("hostUrl")]
        public Input<string>? HostUrl { get; set; }

        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Optional. The full resource name for the GitHubEnterpriseConfig For example: "projects/{$project_id}/githubEnterpriseConfigs/{$config_id}"
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Optional. The network to be used when reaching out to the GitHub Enterprise server. The VPC network must be enabled for private service connection. This should be set if the GitHub Enterprise server is hosted on-premises and not reachable by public internet. If this field is left empty, no network peering will occur and calls to the GitHub Enterprise server will be made over the public internet. Must be in the format `projects/{project}/global/networks/{network}`, where {project} is a project number or id and {network} is the name of a VPC network in the project.
        /// </summary>
        [Input("peeredNetwork")]
        public Input<string>? PeeredNetwork { get; set; }

        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// ID of the project.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// Names of secrets in Secret Manager.
        /// </summary>
        [Input("secrets")]
        public Input<Inputs.GitHubEnterpriseSecretsArgs>? Secrets { get; set; }

        /// <summary>
        /// Optional. SSL certificate to use for requests to GitHub Enterprise.
        /// </summary>
        [Input("sslCa")]
        public Input<string>? SslCa { get; set; }

        /// <summary>
        /// The key that should be attached to webhook calls to the ReceiveWebhook endpoint.
        /// </summary>
        [Input("webhookKey")]
        public Input<string>? WebhookKey { get; set; }

        public GithubEnterpriseConfigArgs()
        {
        }
    }
}
