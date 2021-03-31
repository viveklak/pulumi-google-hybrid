// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleCloud.Apigateway.V1beta
{
    /// <summary>
    /// Creates a new ApiConfig in a given project and location.
    /// </summary>
    [GoogleCloudResourceType("google-cloud:apigateway/v1beta:ApiConfig")]
    public partial class ApiConfig : Pulumi.CustomResource
    {
        /// <summary>
        /// Create a ApiConfig resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ApiConfig(string name, ApiConfigArgs args, CustomResourceOptions? options = null)
            : base("google-cloud:apigateway/v1beta:ApiConfig", name, args ?? new ApiConfigArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ApiConfig(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-cloud:apigateway/v1beta:ApiConfig", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing ApiConfig resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ApiConfig Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ApiConfig(name, id, options);
        }
    }

    public sealed class ApiConfigArgs : Pulumi.ResourceArgs
    {
        [Input("apisId", required: true)]
        public Input<string> ApisId { get; set; } = null!;

        [Input("configsId", required: true)]
        public Input<string> ConfigsId { get; set; } = null!;

        /// <summary>
        /// Output only. Created time.
        /// </summary>
        [Input("createTime")]
        public Input<string>? CreateTime { get; set; }

        /// <summary>
        /// Optional. Display name.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// Immutable. Gateway specific configuration.
        /// </summary>
        [Input("gatewayConfig")]
        public Input<Inputs.ApigatewayGatewayConfigArgs>? GatewayConfig { get; set; }

        /// <summary>
        /// Immutable. The Google Cloud IAM Service Account that Gateways serving this config should use to authenticate to other services. This may either be the Service Account's email (`{ACCOUNT_ID}@{PROJECT}.iam.gserviceaccount.com`) or its full resource name (`projects/{PROJECT}/accounts/{UNIQUE_ID}`). This is most often used when the service is a GCP resource such as a Cloud Run Service or an IAP-secured service.
        /// </summary>
        [Input("gatewayServiceAccount")]
        public Input<string>? GatewayServiceAccount { get; set; }

        [Input("grpcServices")]
        private InputList<Inputs.ApigatewayApiConfigGrpcServiceDefinitionArgs>? _grpcServices;

        /// <summary>
        /// Optional. gRPC service definition files. If specified, openapi_documents must not be included.
        /// </summary>
        public InputList<Inputs.ApigatewayApiConfigGrpcServiceDefinitionArgs> GrpcServices
        {
            get => _grpcServices ?? (_grpcServices = new InputList<Inputs.ApigatewayApiConfigGrpcServiceDefinitionArgs>());
            set => _grpcServices = value;
        }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// Optional. Resource labels to represent user-provided metadata. Refer to cloud documentation on labels for more details. https://cloud.google.com/compute/docs/labeling-resources
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        [Input("locationsId", required: true)]
        public Input<string> LocationsId { get; set; } = null!;

        [Input("managedServiceConfigs")]
        private InputList<Inputs.ApigatewayApiConfigFileArgs>? _managedServiceConfigs;

        /// <summary>
        /// Optional. Service Configuration files. At least one must be included when using gRPC service definitions. See https://cloud.google.com/endpoints/docs/grpc/grpc-service-config#service_configuration_overview for the expected file contents. If multiple files are specified, the files are merged with the following rules: * All singular scalar fields are merged using "last one wins" semantics in the order of the files uploaded. * Repeated fields are concatenated. * Singular embedded messages are merged using these rules for nested fields.
        /// </summary>
        public InputList<Inputs.ApigatewayApiConfigFileArgs> ManagedServiceConfigs
        {
            get => _managedServiceConfigs ?? (_managedServiceConfigs = new InputList<Inputs.ApigatewayApiConfigFileArgs>());
            set => _managedServiceConfigs = value;
        }

        /// <summary>
        /// Output only. Resource name of the API Config. Format: projects/{project}/locations/global/apis/{api}/configs/{api_config}
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("openapiDocuments")]
        private InputList<Inputs.ApigatewayApiConfigOpenApiDocumentArgs>? _openapiDocuments;

        /// <summary>
        /// Optional. OpenAPI specification documents. If specified, grpc_services and managed_service_configs must not be included.
        /// </summary>
        public InputList<Inputs.ApigatewayApiConfigOpenApiDocumentArgs> OpenapiDocuments
        {
            get => _openapiDocuments ?? (_openapiDocuments = new InputList<Inputs.ApigatewayApiConfigOpenApiDocumentArgs>());
            set => _openapiDocuments = value;
        }

        [Input("projectsId", required: true)]
        public Input<string> ProjectsId { get; set; } = null!;

        /// <summary>
        /// Output only. The ID of the associated Service Config ( https://cloud.google.com/service-infrastructure/docs/glossary#config).
        /// </summary>
        [Input("serviceConfigId")]
        public Input<string>? ServiceConfigId { get; set; }

        /// <summary>
        /// Output only. State of the API Config.
        /// </summary>
        [Input("state")]
        public Input<string>? State { get; set; }

        /// <summary>
        /// Output only. Updated time.
        /// </summary>
        [Input("updateTime")]
        public Input<string>? UpdateTime { get; set; }

        public ApiConfigArgs()
        {
        }
    }
}
