// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Firebase.V1Beta1
{
    /// <summary>
    /// Requests the creation of a new WebApp in the specified FirebaseProject. The result of this call is an `Operation` which can be used to track the provisioning process. The `Operation` is automatically deleted after completion, so there is no need to call `DeleteOperation`.
    /// </summary>
    [GoogleNativeResourceType("google-native:firebase/v1beta1:WebApp")]
    public partial class WebApp : Pulumi.CustomResource
    {
        /// <summary>
        /// Immutable. The globally unique, Firebase-assigned identifier for the `WebApp`. This identifier should be treated as an opaque token, as the data format is not specified.
        /// </summary>
        [Output("appId")]
        public Output<string> AppId { get; private set; } = null!;

        /// <summary>
        /// The URLs where the `WebApp` is hosted.
        /// </summary>
        [Output("appUrls")]
        public Output<ImmutableArray<string>> AppUrls { get; private set; } = null!;

        /// <summary>
        /// The user-assigned display name for the `WebApp`.
        /// </summary>
        [Output("displayName")]
        public Output<string> DisplayName { get; private set; } = null!;

        /// <summary>
        /// The resource name of the WebApp, in the format: projects/PROJECT_IDENTIFIER /webApps/APP_ID * PROJECT_IDENTIFIER: the parent Project's [`ProjectNumber`](../projects#FirebaseProject.FIELDS.project_number) ***(recommended)*** or its [`ProjectId`](../projects#FirebaseProject.FIELDS.project_id). Learn more about using project identifiers in Google's [AIP 2510 standard](https://google.aip.dev/cloud/2510). Note that the value for PROJECT_IDENTIFIER in any response body will be the `ProjectId`. * APP_ID: the globally unique, Firebase-assigned identifier for the App (see [`appId`](../projects.webApps#WebApp.FIELDS.app_id)).
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Immutable. A user-assigned unique identifier of the parent FirebaseProject for the `WebApp`.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// Immutable. A unique, Firebase-assigned identifier for the `WebApp`. This identifier is only used to populate the `namespace` value for the `WebApp`. For most use cases, use `appId` to identify or reference the App. The `webId` value is only unique within a `FirebaseProject` and its associated Apps.
        /// </summary>
        [Output("webId")]
        public Output<string> WebId { get; private set; } = null!;


        /// <summary>
        /// Create a WebApp resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public WebApp(string name, WebAppArgs args, CustomResourceOptions? options = null)
            : base("google-native:firebase/v1beta1:WebApp", name, args ?? new WebAppArgs(), MakeResourceOptions(options, ""))
        {
        }

        private WebApp(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:firebase/v1beta1:WebApp", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing WebApp resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static WebApp Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new WebApp(name, id, options);
        }
    }

    public sealed class WebAppArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Immutable. The globally unique, Firebase-assigned identifier for the `WebApp`. This identifier should be treated as an opaque token, as the data format is not specified.
        /// </summary>
        [Input("appId")]
        public Input<string>? AppId { get; set; }

        [Input("appUrls")]
        private InputList<string>? _appUrls;

        /// <summary>
        /// The URLs where the `WebApp` is hosted.
        /// </summary>
        public InputList<string> AppUrls
        {
            get => _appUrls ?? (_appUrls = new InputList<string>());
            set => _appUrls = value;
        }

        /// <summary>
        /// The user-assigned display name for the `WebApp`.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// The resource name of the WebApp, in the format: projects/PROJECT_IDENTIFIER /webApps/APP_ID * PROJECT_IDENTIFIER: the parent Project's [`ProjectNumber`](../projects#FirebaseProject.FIELDS.project_number) ***(recommended)*** or its [`ProjectId`](../projects#FirebaseProject.FIELDS.project_id). Learn more about using project identifiers in Google's [AIP 2510 standard](https://google.aip.dev/cloud/2510). Note that the value for PROJECT_IDENTIFIER in any response body will be the `ProjectId`. * APP_ID: the globally unique, Firebase-assigned identifier for the App (see [`appId`](../projects.webApps#WebApp.FIELDS.app_id)).
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Immutable. A user-assigned unique identifier of the parent FirebaseProject for the `WebApp`.
        /// </summary>
        [Input("project", required: true)]
        public Input<string> Project { get; set; } = null!;

        public WebAppArgs()
        {
        }
    }
}
