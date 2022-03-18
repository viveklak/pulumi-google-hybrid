// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Run.V1
{
    /// <summary>
    /// Create a service.
    /// Auto-naming is currently not supported for this resource.
    /// </summary>
    [GoogleNativeResourceType("google-native:run/v1:Service")]
    public partial class Service : Pulumi.CustomResource
    {
        /// <summary>
        /// The API version for this call such as "serving.knative.dev/v1".
        /// </summary>
        [Output("apiVersion")]
        public Output<string> ApiVersion { get; private set; } = null!;

        /// <summary>
        /// The kind of resource, in this case "Service".
        /// </summary>
        [Output("kind")]
        public Output<string> Kind { get; private set; } = null!;

        /// <summary>
        /// Metadata associated with this Service, including name, namespace, labels, and annotations. Cloud Run (fully managed) uses the following annotation keys to configure features on a Service: * `run.googleapis.com/ingress` sets the ingress settings for the Service. See [the ingress settings documentation](/run/docs/securing/ingress) for details on configuring ingress settings. * `run.googleapis.com/ingress-status` is output-only and contains the currently active ingress settings for the Service. `run.googleapis.com/ingress-status` may differ from `run.googleapis.com/ingress` while the system is processing a change to `run.googleapis.com/ingress` or if the system failed to process a change to `run.googleapis.com/ingress`. When the system has processed all changes successfully `run.googleapis.com/ingress-status` and `run.googleapis.com/ingress` are equal.
        /// </summary>
        [Output("metadata")]
        public Output<Outputs.ObjectMetaResponse> Metadata { get; private set; } = null!;

        /// <summary>
        /// Spec holds the desired state of the Service (from the client).
        /// </summary>
        [Output("spec")]
        public Output<Outputs.ServiceSpecResponse> Spec { get; private set; } = null!;

        /// <summary>
        /// Status communicates the observed state of the Service (from the controller).
        /// </summary>
        [Output("status")]
        public Output<Outputs.ServiceStatusResponse> Status { get; private set; } = null!;


        /// <summary>
        /// Create a Service resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Service(string name, ServiceArgs? args = null, CustomResourceOptions? options = null)
            : base("google-native:run/v1:Service", name, args ?? new ServiceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Service(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:run/v1:Service", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing Service resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Service Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Service(name, id, options);
        }
    }

    public sealed class ServiceArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The API version for this call such as "serving.knative.dev/v1".
        /// </summary>
        [Input("apiVersion")]
        public Input<string>? ApiVersion { get; set; }

        /// <summary>
        /// Indicates that the server should validate the request and populate default values without persisting the request. Supported values: `all`
        /// </summary>
        [Input("dryRun")]
        public Input<string>? DryRun { get; set; }

        /// <summary>
        /// The kind of resource, in this case "Service".
        /// </summary>
        [Input("kind")]
        public Input<string>? Kind { get; set; }

        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Metadata associated with this Service, including name, namespace, labels, and annotations. Cloud Run (fully managed) uses the following annotation keys to configure features on a Service: * `run.googleapis.com/ingress` sets the ingress settings for the Service. See [the ingress settings documentation](/run/docs/securing/ingress) for details on configuring ingress settings. * `run.googleapis.com/ingress-status` is output-only and contains the currently active ingress settings for the Service. `run.googleapis.com/ingress-status` may differ from `run.googleapis.com/ingress` while the system is processing a change to `run.googleapis.com/ingress` or if the system failed to process a change to `run.googleapis.com/ingress`. When the system has processed all changes successfully `run.googleapis.com/ingress-status` and `run.googleapis.com/ingress` are equal.
        /// </summary>
        [Input("metadata")]
        public Input<Inputs.ObjectMetaArgs>? Metadata { get; set; }

        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// Spec holds the desired state of the Service (from the client).
        /// </summary>
        [Input("spec")]
        public Input<Inputs.ServiceSpecArgs>? Spec { get; set; }

        public ServiceArgs()
        {
        }
    }
}
