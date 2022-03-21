// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.AppEngine.V1Beta
{
    public static class GetVersion
    {
        /// <summary>
        /// Gets the specified Version resource. By default, only a BASIC_VIEW will be returned. Specify the FULL_VIEW parameter to get the full resource.
        /// </summary>
        public static Task<GetVersionResult> InvokeAsync(GetVersionArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetVersionResult>("google-native:appengine/v1beta:getVersion", args ?? new GetVersionArgs(), options.WithDefaults());

        /// <summary>
        /// Gets the specified Version resource. By default, only a BASIC_VIEW will be returned. Specify the FULL_VIEW parameter to get the full resource.
        /// </summary>
        public static Output<GetVersionResult> Invoke(GetVersionInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetVersionResult>("google-native:appengine/v1beta:getVersion", args ?? new GetVersionInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetVersionArgs : Pulumi.InvokeArgs
    {
        [Input("appId", required: true)]
        public string AppId { get; set; } = null!;

        [Input("serviceId", required: true)]
        public string ServiceId { get; set; } = null!;

        [Input("versionId", required: true)]
        public string VersionId { get; set; } = null!;

        [Input("view")]
        public string? View { get; set; }

        public GetVersionArgs()
        {
        }
    }

    public sealed class GetVersionInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("appId", required: true)]
        public Input<string> AppId { get; set; } = null!;

        [Input("serviceId", required: true)]
        public Input<string> ServiceId { get; set; } = null!;

        [Input("versionId", required: true)]
        public Input<string> VersionId { get; set; } = null!;

        [Input("view")]
        public Input<string>? View { get; set; }

        public GetVersionInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetVersionResult
    {
        /// <summary>
        /// Serving configuration for Google Cloud Endpoints (https://cloud.google.com/appengine/docs/python/endpoints/).Only returned in GET requests if view=FULL is set.
        /// </summary>
        public readonly Outputs.ApiConfigHandlerResponse ApiConfig;
        /// <summary>
        /// app_engine_apis allows second generation runtimes to access the App Engine APIs.
        /// </summary>
        public readonly bool AppEngineApis;
        /// <summary>
        /// Automatic scaling is based on request rate, response latencies, and other application metrics. Instances are dynamically created and destroyed as needed in order to handle traffic.
        /// </summary>
        public readonly Outputs.AutomaticScalingResponse AutomaticScaling;
        /// <summary>
        /// A service with basic scaling will create an instance when the application receives a request. The instance will be turned down when the app becomes idle. Basic scaling is ideal for work that is intermittent or driven by user activity.
        /// </summary>
        public readonly Outputs.BasicScalingResponse BasicScaling;
        /// <summary>
        /// Metadata settings that are supplied to this version to enable beta runtime features.
        /// </summary>
        public readonly ImmutableDictionary<string, string> BetaSettings;
        /// <summary>
        /// Environment variables available to the build environment.Only returned in GET requests if view=FULL is set.
        /// </summary>
        public readonly ImmutableDictionary<string, string> BuildEnvVariables;
        /// <summary>
        /// Time that this version was created.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// Email address of the user who created this version.
        /// </summary>
        public readonly string CreatedBy;
        /// <summary>
        /// Duration that static files should be cached by web proxies and browsers. Only applicable if the corresponding StaticFilesHandler (https://cloud.google.com/appengine/docs/admin-api/reference/rest/v1beta/apps.services.versions#StaticFilesHandler) does not specify its own expiration time.Only returned in GET requests if view=FULL is set.
        /// </summary>
        public readonly string DefaultExpiration;
        /// <summary>
        /// Code and application artifacts that make up this version.Only returned in GET requests if view=FULL is set.
        /// </summary>
        public readonly Outputs.DeploymentResponse Deployment;
        /// <summary>
        /// Total size in bytes of all the files that are included in this version and currently hosted on the App Engine disk.
        /// </summary>
        public readonly string DiskUsageBytes;
        /// <summary>
        /// Cloud Endpoints configuration.If endpoints_api_service is set, the Cloud Endpoints Extensible Service Proxy will be provided to serve the API implemented by the app.
        /// </summary>
        public readonly Outputs.EndpointsApiServiceResponse EndpointsApiService;
        /// <summary>
        /// The entrypoint for the application.
        /// </summary>
        public readonly Outputs.EntrypointResponse Entrypoint;
        /// <summary>
        /// App Engine execution environment for this version.Defaults to standard.
        /// </summary>
        public readonly string Env;
        /// <summary>
        /// Environment variables available to the application.Only returned in GET requests if view=FULL is set.
        /// </summary>
        public readonly ImmutableDictionary<string, string> EnvVariables;
        /// <summary>
        /// Custom static error pages. Limited to 10KB per page.Only returned in GET requests if view=FULL is set.
        /// </summary>
        public readonly ImmutableArray<Outputs.ErrorHandlerResponse> ErrorHandlers;
        /// <summary>
        /// An ordered list of URL-matching patterns that should be applied to incoming requests. The first matching URL handles the request and other request handlers are not attempted.Only returned in GET requests if view=FULL is set.
        /// </summary>
        public readonly ImmutableArray<Outputs.UrlMapResponse> Handlers;
        /// <summary>
        /// Configures health checking for instances. Unhealthy instances are stopped and replaced with new instances. Only applicable in the App Engine flexible environment.Only returned in GET requests if view=FULL is set.
        /// </summary>
        public readonly Outputs.HealthCheckResponse HealthCheck;
        /// <summary>
        /// Before an application can receive email or XMPP messages, the application must be configured to enable the service.
        /// </summary>
        public readonly ImmutableArray<string> InboundServices;
        /// <summary>
        /// Instance class that is used to run this version. Valid values are: AutomaticScaling: F1, F2, F4, F4_1G ManualScaling or BasicScaling: B1, B2, B4, B8, B4_1GDefaults to F1 for AutomaticScaling and B1 for ManualScaling or BasicScaling.
        /// </summary>
        public readonly string InstanceClass;
        /// <summary>
        /// Configuration for third-party Python runtime libraries that are required by the application.Only returned in GET requests if view=FULL is set.
        /// </summary>
        public readonly ImmutableArray<Outputs.LibraryResponse> Libraries;
        /// <summary>
        /// Configures liveness health checking for instances. Unhealthy instances are stopped and replaced with new instancesOnly returned in GET requests if view=FULL is set.
        /// </summary>
        public readonly Outputs.LivenessCheckResponse LivenessCheck;
        /// <summary>
        /// A service with manual scaling runs continuously, allowing you to perform complex initialization and rely on the state of its memory over time. Manually scaled versions are sometimes referred to as "backends".
        /// </summary>
        public readonly Outputs.ManualScalingResponse ManualScaling;
        /// <summary>
        /// Full path to the Version resource in the API. Example: apps/myapp/services/default/versions/v1.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Extra network settings. Only applicable in the App Engine flexible environment.
        /// </summary>
        public readonly Outputs.NetworkResponse Network;
        /// <summary>
        /// Files that match this pattern will not be built into this version. Only applicable for Go runtimes.Only returned in GET requests if view=FULL is set.
        /// </summary>
        public readonly string NobuildFilesRegex;
        /// <summary>
        /// Configures readiness health checking for instances. Unhealthy instances are not put into the backend traffic rotation.Only returned in GET requests if view=FULL is set.
        /// </summary>
        public readonly Outputs.ReadinessCheckResponse ReadinessCheck;
        /// <summary>
        /// Machine resources for this version. Only applicable in the App Engine flexible environment.
        /// </summary>
        public readonly Outputs.ResourcesResponse Resources;
        /// <summary>
        /// Desired runtime. Example: python27.
        /// </summary>
        public readonly string Runtime;
        /// <summary>
        /// The version of the API in the given runtime environment. Please see the app.yaml reference for valid values at https://cloud.google.com/appengine/docs/standard//config/appref
        /// </summary>
        public readonly string RuntimeApiVersion;
        /// <summary>
        /// The channel of the runtime to use. Only available for some runtimes. Defaults to the default channel.
        /// </summary>
        public readonly string RuntimeChannel;
        /// <summary>
        /// The path or name of the app's main executable.
        /// </summary>
        public readonly string RuntimeMainExecutablePath;
        /// <summary>
        /// The identity that the deployed version will run as. Admin API will use the App Engine Appspot service account as default if this field is neither provided in app.yaml file nor through CLI flag.
        /// </summary>
        public readonly string ServiceAccount;
        /// <summary>
        /// Current serving status of this version. Only the versions with a SERVING status create instances and can be billed.SERVING_STATUS_UNSPECIFIED is an invalid value. Defaults to SERVING.
        /// </summary>
        public readonly string ServingStatus;
        /// <summary>
        /// Whether multiple requests can be dispatched to this version at once.
        /// </summary>
        public readonly bool Threadsafe;
        /// <summary>
        /// Serving URL for this version. Example: "https://myversion-dot-myservice-dot-myapp.appspot.com"
        /// </summary>
        public readonly string VersionUrl;
        /// <summary>
        /// Whether to deploy this version in a container on a virtual machine.
        /// </summary>
        public readonly bool Vm;
        /// <summary>
        /// Enables VPC connectivity for standard apps.
        /// </summary>
        public readonly Outputs.VpcAccessConnectorResponse VpcAccessConnector;
        /// <summary>
        /// The Google Compute Engine zones that are supported by this version in the App Engine flexible environment. Deprecated.
        /// </summary>
        public readonly ImmutableArray<string> Zones;

        [OutputConstructor]
        private GetVersionResult(
            Outputs.ApiConfigHandlerResponse apiConfig,

            bool appEngineApis,

            Outputs.AutomaticScalingResponse automaticScaling,

            Outputs.BasicScalingResponse basicScaling,

            ImmutableDictionary<string, string> betaSettings,

            ImmutableDictionary<string, string> buildEnvVariables,

            string createTime,

            string createdBy,

            string defaultExpiration,

            Outputs.DeploymentResponse deployment,

            string diskUsageBytes,

            Outputs.EndpointsApiServiceResponse endpointsApiService,

            Outputs.EntrypointResponse entrypoint,

            string env,

            ImmutableDictionary<string, string> envVariables,

            ImmutableArray<Outputs.ErrorHandlerResponse> errorHandlers,

            ImmutableArray<Outputs.UrlMapResponse> handlers,

            Outputs.HealthCheckResponse healthCheck,

            ImmutableArray<string> inboundServices,

            string instanceClass,

            ImmutableArray<Outputs.LibraryResponse> libraries,

            Outputs.LivenessCheckResponse livenessCheck,

            Outputs.ManualScalingResponse manualScaling,

            string name,

            Outputs.NetworkResponse network,

            string nobuildFilesRegex,

            Outputs.ReadinessCheckResponse readinessCheck,

            Outputs.ResourcesResponse resources,

            string runtime,

            string runtimeApiVersion,

            string runtimeChannel,

            string runtimeMainExecutablePath,

            string serviceAccount,

            string servingStatus,

            bool threadsafe,

            string versionUrl,

            bool vm,

            Outputs.VpcAccessConnectorResponse vpcAccessConnector,

            ImmutableArray<string> zones)
        {
            ApiConfig = apiConfig;
            AppEngineApis = appEngineApis;
            AutomaticScaling = automaticScaling;
            BasicScaling = basicScaling;
            BetaSettings = betaSettings;
            BuildEnvVariables = buildEnvVariables;
            CreateTime = createTime;
            CreatedBy = createdBy;
            DefaultExpiration = defaultExpiration;
            Deployment = deployment;
            DiskUsageBytes = diskUsageBytes;
            EndpointsApiService = endpointsApiService;
            Entrypoint = entrypoint;
            Env = env;
            EnvVariables = envVariables;
            ErrorHandlers = errorHandlers;
            Handlers = handlers;
            HealthCheck = healthCheck;
            InboundServices = inboundServices;
            InstanceClass = instanceClass;
            Libraries = libraries;
            LivenessCheck = livenessCheck;
            ManualScaling = manualScaling;
            Name = name;
            Network = network;
            NobuildFilesRegex = nobuildFilesRegex;
            ReadinessCheck = readinessCheck;
            Resources = resources;
            Runtime = runtime;
            RuntimeApiVersion = runtimeApiVersion;
            RuntimeChannel = runtimeChannel;
            RuntimeMainExecutablePath = runtimeMainExecutablePath;
            ServiceAccount = serviceAccount;
            ServingStatus = servingStatus;
            Threadsafe = threadsafe;
            VersionUrl = versionUrl;
            Vm = vm;
            VpcAccessConnector = vpcAccessConnector;
            Zones = zones;
        }
    }
}
