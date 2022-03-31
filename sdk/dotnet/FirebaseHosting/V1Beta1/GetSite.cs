// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.FirebaseHosting.V1Beta1
{
    public static class GetSite
    {
        /// <summary>
        /// Gets the specified Hosting Site.
        /// </summary>
        public static Task<GetSiteResult> InvokeAsync(GetSiteArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetSiteResult>("google-native:firebasehosting/v1beta1:getSite", args ?? new GetSiteArgs(), options.WithDefaults());

        /// <summary>
        /// Gets the specified Hosting Site.
        /// </summary>
        public static Output<GetSiteResult> Invoke(GetSiteInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetSiteResult>("google-native:firebasehosting/v1beta1:getSite", args ?? new GetSiteInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSiteArgs : Pulumi.InvokeArgs
    {
        [Input("project")]
        public string? Project { get; set; }

        [Input("siteId", required: true)]
        public string SiteId { get; set; } = null!;

        public GetSiteArgs()
        {
        }
    }

    public sealed class GetSiteInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("siteId", required: true)]
        public Input<string> SiteId { get; set; } = null!;

        public GetSiteInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetSiteResult
    {
        /// <summary>
        /// Optional. The [ID of a Web App](https://firebase.google.com/docs/projects/api/reference/rest/v1beta1/projects.webApps#WebApp.FIELDS.app_id) associated with the Hosting site.
        /// </summary>
        public readonly string AppId;
        /// <summary>
        /// The default URL for the Hosting site.
        /// </summary>
        public readonly string DefaultUrl;
        /// <summary>
        /// Optional. User-specified labels for the Hosting site.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Labels;
        /// <summary>
        /// The fully-qualified resource name of the Hosting site, in the format: projects/PROJECT_IDENTIFIER/sites/SITE_ID PROJECT_IDENTIFIER: the Firebase project's [`ProjectNumber`](https://firebase.google.com/docs/projects/api/reference/rest/v1beta1/projects#FirebaseProject.FIELDS.project_number) ***(recommended)*** or its [`ProjectId`](https://firebase.google.com/docs/projects/api/reference/rest/v1beta1/projects#FirebaseProject.FIELDS.project_id). Learn more about using project identifiers in Google's [AIP 2510 standard](https://google.aip.dev/cloud/2510).
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The type of Hosting site. Every Firebase project has a `DEFAULT_SITE`, which is created when Hosting is provisioned for the project. All additional sites are `USER_SITE`.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetSiteResult(
            string appId,

            string defaultUrl,

            ImmutableDictionary<string, string> labels,

            string name,

            string type)
        {
            AppId = appId;
            DefaultUrl = defaultUrl;
            Labels = labels;
            Name = name;
            Type = type;
        }
    }
}
