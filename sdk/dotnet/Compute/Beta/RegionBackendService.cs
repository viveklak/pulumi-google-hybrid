// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Beta
{
    /// <summary>
    /// Creates a regional BackendService resource in the specified project using the data included in the request. For more information, see Backend services overview.
    /// </summary>
    [GoogleNativeResourceType("google-native:compute/beta:RegionBackendService")]
    public partial class RegionBackendService : Pulumi.CustomResource
    {
        /// <summary>
        /// Lifetime of cookies in seconds. Only applicable if the loadBalancingScheme is EXTERNAL, INTERNAL_SELF_MANAGED, or INTERNAL_MANAGED, the protocol is HTTP or HTTPS, and the sessionAffinity is GENERATED_COOKIE, or HTTP_COOKIE. If set to 0, the cookie is non-persistent and lasts only until the end of the browser session (or equivalent). The maximum allowed value is one day ( 86,400). Not supported when the backend service is referenced by a URL map that is bound to target gRPC proxy that has validateForProxyless field set to true.
        /// </summary>
        [Output("affinityCookieTtlSec")]
        public Output<int> AffinityCookieTtlSec { get; private set; } = null!;

        /// <summary>
        /// The list of backends that serve this BackendService.
        /// </summary>
        [Output("backends")]
        public Output<ImmutableArray<Outputs.BackendResponse>> Backends { get; private set; } = null!;

        /// <summary>
        /// Cloud CDN configuration for this BackendService. Only available for external HTTP(S) Load Balancing.
        /// </summary>
        [Output("cdnPolicy")]
        public Output<Outputs.BackendServiceCdnPolicyResponse> CdnPolicy { get; private set; } = null!;

        [Output("circuitBreakers")]
        public Output<Outputs.CircuitBreakersResponse> CircuitBreakers { get; private set; } = null!;

        /// <summary>
        /// Compress text responses using Brotli or gzip compression, based on the client’s Accept-Encoding header.
        /// </summary>
        [Output("compressionMode")]
        public Output<string> CompressionMode { get; private set; } = null!;

        [Output("connectionDraining")]
        public Output<Outputs.ConnectionDrainingResponse> ConnectionDraining { get; private set; } = null!;

        [Output("connectionTrackingPolicy")]
        public Output<Outputs.BackendServiceConnectionTrackingPolicyResponse> ConnectionTrackingPolicy { get; private set; } = null!;

        /// <summary>
        /// Consistent Hash-based load balancing can be used to provide soft session affinity based on HTTP headers, cookies or other properties. This load balancing policy is applicable only for HTTP connections. The affinity to a particular destination host will be lost when one or more hosts are added/removed from the destination service. This field specifies parameters that control consistent hashing. This field is only applicable when localityLbPolicy is set to MAGLEV or RING_HASH. This field is applicable to either: - A regional backend service with the service_protocol set to HTTP, HTTPS, or HTTP2, and load_balancing_scheme set to INTERNAL_MANAGED. - A global backend service with the load_balancing_scheme set to INTERNAL_SELF_MANAGED. Not supported when the backend service is referenced by a URL map that is bound to target gRPC proxy that has validateForProxyless field set to true.
        /// </summary>
        [Output("consistentHash")]
        public Output<Outputs.ConsistentHashLoadBalancerSettingsResponse> ConsistentHash { get; private set; } = null!;

        /// <summary>
        /// [Output Only] Creation timestamp in RFC3339 text format.
        /// </summary>
        [Output("creationTimestamp")]
        public Output<string> CreationTimestamp { get; private set; } = null!;

        /// <summary>
        /// Headers that the HTTP/S load balancer should add to proxied requests.
        /// </summary>
        [Output("customRequestHeaders")]
        public Output<ImmutableArray<string>> CustomRequestHeaders { get; private set; } = null!;

        /// <summary>
        /// Headers that the HTTP/S load balancer should add to proxied responses.
        /// </summary>
        [Output("customResponseHeaders")]
        public Output<ImmutableArray<string>> CustomResponseHeaders { get; private set; } = null!;

        /// <summary>
        /// An optional description of this resource. Provide this property when you create the resource.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// [Output Only] The resource URL for the edge security policy associated with this backend service.
        /// </summary>
        [Output("edgeSecurityPolicy")]
        public Output<string> EdgeSecurityPolicy { get; private set; } = null!;

        /// <summary>
        /// If true, enables Cloud CDN for the backend service. Only applicable if the loadBalancingScheme is EXTERNAL and the protocol is HTTP or HTTPS.
        /// </summary>
        [Output("enableCDN")]
        public Output<bool> EnableCDN { get; private set; } = null!;

        /// <summary>
        /// Applicable only to Failover for Internal TCP/UDP Load Balancing and Network Load Balancing. Requires at least one backend instance group to be defined as a backup (failover) backend.
        /// </summary>
        [Output("failoverPolicy")]
        public Output<Outputs.BackendServiceFailoverPolicyResponse> FailoverPolicy { get; private set; } = null!;

        /// <summary>
        /// Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking. This field will be ignored when inserting a BackendService. An up-to-date fingerprint must be provided in order to update the BackendService, otherwise the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make a get() request to retrieve a BackendService.
        /// </summary>
        [Output("fingerprint")]
        public Output<string> Fingerprint { get; private set; } = null!;

        /// <summary>
        /// The list of URLs to the healthChecks, httpHealthChecks (legacy), or httpsHealthChecks (legacy) resource for health checking this backend service. Not all backend services support legacy health checks. See Load balancer guide. Currently, at most one health check can be specified for each backend service. Backend services with instance group or zonal NEG backends must have a health check. Backend services with internet or serverless NEG backends must not have a health check.
        /// </summary>
        [Output("healthChecks")]
        public Output<ImmutableArray<string>> HealthChecks { get; private set; } = null!;

        /// <summary>
        /// The configurations for Identity-Aware Proxy on this resource. Not available for Internal TCP/UDP Load Balancing and Network Load Balancing.
        /// </summary>
        [Output("iap")]
        public Output<Outputs.BackendServiceIAPResponse> Iap { get; private set; } = null!;

        /// <summary>
        /// [Output Only] Type of resource. Always compute#backendService for backend services.
        /// </summary>
        [Output("kind")]
        public Output<string> Kind { get; private set; } = null!;

        /// <summary>
        /// Specifies the load balancer type. Choose EXTERNAL for external HTTP(S), SSL Proxy, TCP Proxy and Network Load Balancing. Choose INTERNAL for Internal TCP/UDP Load Balancing. Choose INTERNAL_MANAGED for Internal HTTP(S) Load Balancing. INTERNAL_SELF_MANAGED for Traffic Director. A backend service created for one type of load balancer cannot be used with another. For more information, refer to Choosing a load balancer.
        /// </summary>
        [Output("loadBalancingScheme")]
        public Output<string> LoadBalancingScheme { get; private set; } = null!;

        /// <summary>
        /// The load balancing algorithm used within the scope of the locality. The possible values are: - ROUND_ROBIN: This is a simple policy in which each healthy backend is selected in round robin order. This is the default. - LEAST_REQUEST: An O(1) algorithm which selects two random healthy hosts and picks the host which has fewer active requests. - RING_HASH: The ring/modulo hash load balancer implements consistent hashing to backends. The algorithm has the property that the addition/removal of a host from a set of N hosts only affects 1/N of the requests. - RANDOM: The load balancer selects a random healthy host. - ORIGINAL_DESTINATION: Backend host is selected based on the client connection metadata, i.e., connections are opened to the same address as the destination address of the incoming connection before the connection was redirected to the load balancer. - MAGLEV: used as a drop in replacement for the ring hash load balancer. Maglev is not as stable as ring hash but has faster table lookup build times and host selection times. For more information about Maglev, see https://ai.google/research/pubs/pub44824 This field is applicable to either: - A regional backend service with the service_protocol set to HTTP, HTTPS, or HTTP2, and load_balancing_scheme set to INTERNAL_MANAGED. - A global backend service with the load_balancing_scheme set to INTERNAL_SELF_MANAGED. If sessionAffinity is not NONE, and this field is not set to MAGLEV or RING_HASH, session affinity settings will not take effect. Only the default ROUND_ROBIN policy is supported when the backend service is referenced by a URL map that is bound to target gRPC proxy that has validateForProxyless field set to true.
        /// </summary>
        [Output("localityLbPolicy")]
        public Output<string> LocalityLbPolicy { get; private set; } = null!;

        /// <summary>
        /// This field denotes the logging options for the load balancer traffic served by this backend service. If logging is enabled, logs will be exported to Stackdriver.
        /// </summary>
        [Output("logConfig")]
        public Output<Outputs.BackendServiceLogConfigResponse> LogConfig { get; private set; } = null!;

        /// <summary>
        /// Specifies the default maximum duration (timeout) for streams to this service. Duration is computed from the beginning of the stream until the response has been completely processed, including all retries. A stream that does not complete in this duration is closed. If not specified, there will be no timeout limit, i.e. the maximum duration is infinite. This value can be overridden in the PathMatcher configuration of the UrlMap that references this backend service. This field is only allowed when the loadBalancingScheme of the backend service is INTERNAL_SELF_MANAGED.
        /// </summary>
        [Output("maxStreamDuration")]
        public Output<Outputs.DurationResponse> MaxStreamDuration { get; private set; } = null!;

        /// <summary>
        /// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The URL of the network to which this backend service belongs. This field can only be specified when the load balancing scheme is set to INTERNAL.
        /// </summary>
        [Output("network")]
        public Output<string> Network { get; private set; } = null!;

        /// <summary>
        /// Settings controlling the eviction of unhealthy hosts from the load balancing pool for the backend service. If not set, this feature is considered disabled. This field is applicable to either: - A regional backend service with the service_protocol set to HTTP, HTTPS, or HTTP2, and load_balancing_scheme set to INTERNAL_MANAGED. - A global backend service with the load_balancing_scheme set to INTERNAL_SELF_MANAGED. Not supported when the backend service is referenced by a URL map that is bound to target gRPC proxy that has validateForProxyless field set to true.
        /// </summary>
        [Output("outlierDetection")]
        public Output<Outputs.OutlierDetectionResponse> OutlierDetection { get; private set; } = null!;

        /// <summary>
        /// A named port on a backend instance group representing the port for communication to the backend VMs in that group. Required when the loadBalancingScheme is EXTERNAL (except Network Load Balancing), INTERNAL_MANAGED, or INTERNAL_SELF_MANAGED and the backends are instance groups. The named port must be defined on each backend instance group. This parameter has no meaning if the backends are NEGs. Backend services for Internal TCP/UDP Load Balancing and Network Load Balancing require you omit port_name.
        /// </summary>
        [Output("portName")]
        public Output<string> PortName { get; private set; } = null!;

        /// <summary>
        /// The protocol this BackendService uses to communicate with backends. Possible values are HTTP, HTTPS, HTTP2, TCP, SSL, UDP or GRPC. depending on the chosen load balancer or Traffic Director configuration. Refer to the documentation for the load balancer or for Traffic Director for more information. Must be set to GRPC when the backend service is referenced by a URL map that is bound to target gRPC proxy.
        /// </summary>
        [Output("protocol")]
        public Output<string> Protocol { get; private set; } = null!;

        /// <summary>
        /// [Output Only] URL of the region where the regional backend service resides. This field is not applicable to global backend services. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// [Output Only] The resource URL for the security policy associated with this backend service.
        /// </summary>
        [Output("securityPolicy")]
        public Output<string> SecurityPolicy { get; private set; } = null!;

        /// <summary>
        /// This field specifies the security policy that applies to this backend service. This field is applicable to either: - A regional backend service with the service_protocol set to HTTP, HTTPS, or HTTP2, and load_balancing_scheme set to INTERNAL_MANAGED. - A global backend service with the load_balancing_scheme set to INTERNAL_SELF_MANAGED. 
        /// </summary>
        [Output("securitySettings")]
        public Output<Outputs.SecuritySettingsResponse> SecuritySettings { get; private set; } = null!;

        /// <summary>
        /// [Output Only] Server-defined URL for the resource.
        /// </summary>
        [Output("selfLink")]
        public Output<string> SelfLink { get; private set; } = null!;

        /// <summary>
        /// Type of session affinity to use. The default is NONE. When the loadBalancingScheme is EXTERNAL: * For Network Load Balancing, the possible values are NONE, CLIENT_IP, CLIENT_IP_PROTO, or CLIENT_IP_PORT_PROTO. * For all other load balancers that use loadBalancingScheme=EXTERNAL, the possible values are NONE, CLIENT_IP, or GENERATED_COOKIE. * You can use GENERATED_COOKIE if the protocol is HTTP, HTTP2, or HTTPS. When the loadBalancingScheme is INTERNAL, possible values are NONE, CLIENT_IP, CLIENT_IP_PROTO, or CLIENT_IP_PORT_PROTO. When the loadBalancingScheme is INTERNAL_SELF_MANAGED, or INTERNAL_MANAGED, possible values are NONE, CLIENT_IP, GENERATED_COOKIE, HEADER_FIELD, or HTTP_COOKIE. Not supported when the backend service is referenced by a URL map that is bound to target gRPC proxy that has validateForProxyless field set to true.
        /// </summary>
        [Output("sessionAffinity")]
        public Output<string> SessionAffinity { get; private set; } = null!;

        [Output("subsetting")]
        public Output<Outputs.SubsettingResponse> Subsetting { get; private set; } = null!;

        /// <summary>
        /// Not supported when the backend service is referenced by a URL map that is bound to target gRPC proxy that has validateForProxyless field set to true. Instead, use maxStreamDuration.
        /// </summary>
        [Output("timeoutSec")]
        public Output<int> TimeoutSec { get; private set; } = null!;


        /// <summary>
        /// Create a RegionBackendService resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RegionBackendService(string name, RegionBackendServiceArgs args, CustomResourceOptions? options = null)
            : base("google-native:compute/beta:RegionBackendService", name, args ?? new RegionBackendServiceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private RegionBackendService(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:compute/beta:RegionBackendService", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing RegionBackendService resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RegionBackendService Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new RegionBackendService(name, id, options);
        }
    }

    public sealed class RegionBackendServiceArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Lifetime of cookies in seconds. Only applicable if the loadBalancingScheme is EXTERNAL, INTERNAL_SELF_MANAGED, or INTERNAL_MANAGED, the protocol is HTTP or HTTPS, and the sessionAffinity is GENERATED_COOKIE, or HTTP_COOKIE. If set to 0, the cookie is non-persistent and lasts only until the end of the browser session (or equivalent). The maximum allowed value is one day ( 86,400). Not supported when the backend service is referenced by a URL map that is bound to target gRPC proxy that has validateForProxyless field set to true.
        /// </summary>
        [Input("affinityCookieTtlSec")]
        public Input<int>? AffinityCookieTtlSec { get; set; }

        [Input("backends")]
        private InputList<Inputs.BackendArgs>? _backends;

        /// <summary>
        /// The list of backends that serve this BackendService.
        /// </summary>
        public InputList<Inputs.BackendArgs> Backends
        {
            get => _backends ?? (_backends = new InputList<Inputs.BackendArgs>());
            set => _backends = value;
        }

        /// <summary>
        /// Cloud CDN configuration for this BackendService. Only available for external HTTP(S) Load Balancing.
        /// </summary>
        [Input("cdnPolicy")]
        public Input<Inputs.BackendServiceCdnPolicyArgs>? CdnPolicy { get; set; }

        [Input("circuitBreakers")]
        public Input<Inputs.CircuitBreakersArgs>? CircuitBreakers { get; set; }

        /// <summary>
        /// Compress text responses using Brotli or gzip compression, based on the client’s Accept-Encoding header.
        /// </summary>
        [Input("compressionMode")]
        public Input<Pulumi.GoogleNative.Compute.Beta.RegionBackendServiceCompressionMode>? CompressionMode { get; set; }

        [Input("connectionDraining")]
        public Input<Inputs.ConnectionDrainingArgs>? ConnectionDraining { get; set; }

        [Input("connectionTrackingPolicy")]
        public Input<Inputs.BackendServiceConnectionTrackingPolicyArgs>? ConnectionTrackingPolicy { get; set; }

        /// <summary>
        /// Consistent Hash-based load balancing can be used to provide soft session affinity based on HTTP headers, cookies or other properties. This load balancing policy is applicable only for HTTP connections. The affinity to a particular destination host will be lost when one or more hosts are added/removed from the destination service. This field specifies parameters that control consistent hashing. This field is only applicable when localityLbPolicy is set to MAGLEV or RING_HASH. This field is applicable to either: - A regional backend service with the service_protocol set to HTTP, HTTPS, or HTTP2, and load_balancing_scheme set to INTERNAL_MANAGED. - A global backend service with the load_balancing_scheme set to INTERNAL_SELF_MANAGED. Not supported when the backend service is referenced by a URL map that is bound to target gRPC proxy that has validateForProxyless field set to true.
        /// </summary>
        [Input("consistentHash")]
        public Input<Inputs.ConsistentHashLoadBalancerSettingsArgs>? ConsistentHash { get; set; }

        /// <summary>
        /// [Output Only] Creation timestamp in RFC3339 text format.
        /// </summary>
        [Input("creationTimestamp")]
        public Input<string>? CreationTimestamp { get; set; }

        [Input("customRequestHeaders")]
        private InputList<string>? _customRequestHeaders;

        /// <summary>
        /// Headers that the HTTP/S load balancer should add to proxied requests.
        /// </summary>
        public InputList<string> CustomRequestHeaders
        {
            get => _customRequestHeaders ?? (_customRequestHeaders = new InputList<string>());
            set => _customRequestHeaders = value;
        }

        [Input("customResponseHeaders")]
        private InputList<string>? _customResponseHeaders;

        /// <summary>
        /// Headers that the HTTP/S load balancer should add to proxied responses.
        /// </summary>
        public InputList<string> CustomResponseHeaders
        {
            get => _customResponseHeaders ?? (_customResponseHeaders = new InputList<string>());
            set => _customResponseHeaders = value;
        }

        /// <summary>
        /// An optional description of this resource. Provide this property when you create the resource.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// [Output Only] The resource URL for the edge security policy associated with this backend service.
        /// </summary>
        [Input("edgeSecurityPolicy")]
        public Input<string>? EdgeSecurityPolicy { get; set; }

        /// <summary>
        /// If true, enables Cloud CDN for the backend service. Only applicable if the loadBalancingScheme is EXTERNAL and the protocol is HTTP or HTTPS.
        /// </summary>
        [Input("enableCDN")]
        public Input<bool>? EnableCDN { get; set; }

        /// <summary>
        /// Applicable only to Failover for Internal TCP/UDP Load Balancing and Network Load Balancing. Requires at least one backend instance group to be defined as a backup (failover) backend.
        /// </summary>
        [Input("failoverPolicy")]
        public Input<Inputs.BackendServiceFailoverPolicyArgs>? FailoverPolicy { get; set; }

        /// <summary>
        /// Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking. This field will be ignored when inserting a BackendService. An up-to-date fingerprint must be provided in order to update the BackendService, otherwise the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make a get() request to retrieve a BackendService.
        /// </summary>
        [Input("fingerprint")]
        public Input<string>? Fingerprint { get; set; }

        [Input("healthChecks")]
        private InputList<string>? _healthChecks;

        /// <summary>
        /// The list of URLs to the healthChecks, httpHealthChecks (legacy), or httpsHealthChecks (legacy) resource for health checking this backend service. Not all backend services support legacy health checks. See Load balancer guide. Currently, at most one health check can be specified for each backend service. Backend services with instance group or zonal NEG backends must have a health check. Backend services with internet or serverless NEG backends must not have a health check.
        /// </summary>
        public InputList<string> HealthChecks
        {
            get => _healthChecks ?? (_healthChecks = new InputList<string>());
            set => _healthChecks = value;
        }

        /// <summary>
        /// The configurations for Identity-Aware Proxy on this resource. Not available for Internal TCP/UDP Load Balancing and Network Load Balancing.
        /// </summary>
        [Input("iap")]
        public Input<Inputs.BackendServiceIAPArgs>? Iap { get; set; }

        /// <summary>
        /// [Output Only] The unique identifier for the resource. This identifier is defined by the server.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// [Output Only] Type of resource. Always compute#backendService for backend services.
        /// </summary>
        [Input("kind")]
        public Input<string>? Kind { get; set; }

        /// <summary>
        /// Specifies the load balancer type. Choose EXTERNAL for external HTTP(S), SSL Proxy, TCP Proxy and Network Load Balancing. Choose INTERNAL for Internal TCP/UDP Load Balancing. Choose INTERNAL_MANAGED for Internal HTTP(S) Load Balancing. INTERNAL_SELF_MANAGED for Traffic Director. A backend service created for one type of load balancer cannot be used with another. For more information, refer to Choosing a load balancer.
        /// </summary>
        [Input("loadBalancingScheme")]
        public Input<Pulumi.GoogleNative.Compute.Beta.RegionBackendServiceLoadBalancingScheme>? LoadBalancingScheme { get; set; }

        /// <summary>
        /// The load balancing algorithm used within the scope of the locality. The possible values are: - ROUND_ROBIN: This is a simple policy in which each healthy backend is selected in round robin order. This is the default. - LEAST_REQUEST: An O(1) algorithm which selects two random healthy hosts and picks the host which has fewer active requests. - RING_HASH: The ring/modulo hash load balancer implements consistent hashing to backends. The algorithm has the property that the addition/removal of a host from a set of N hosts only affects 1/N of the requests. - RANDOM: The load balancer selects a random healthy host. - ORIGINAL_DESTINATION: Backend host is selected based on the client connection metadata, i.e., connections are opened to the same address as the destination address of the incoming connection before the connection was redirected to the load balancer. - MAGLEV: used as a drop in replacement for the ring hash load balancer. Maglev is not as stable as ring hash but has faster table lookup build times and host selection times. For more information about Maglev, see https://ai.google/research/pubs/pub44824 This field is applicable to either: - A regional backend service with the service_protocol set to HTTP, HTTPS, or HTTP2, and load_balancing_scheme set to INTERNAL_MANAGED. - A global backend service with the load_balancing_scheme set to INTERNAL_SELF_MANAGED. If sessionAffinity is not NONE, and this field is not set to MAGLEV or RING_HASH, session affinity settings will not take effect. Only the default ROUND_ROBIN policy is supported when the backend service is referenced by a URL map that is bound to target gRPC proxy that has validateForProxyless field set to true.
        /// </summary>
        [Input("localityLbPolicy")]
        public Input<Pulumi.GoogleNative.Compute.Beta.RegionBackendServiceLocalityLbPolicy>? LocalityLbPolicy { get; set; }

        /// <summary>
        /// This field denotes the logging options for the load balancer traffic served by this backend service. If logging is enabled, logs will be exported to Stackdriver.
        /// </summary>
        [Input("logConfig")]
        public Input<Inputs.BackendServiceLogConfigArgs>? LogConfig { get; set; }

        /// <summary>
        /// Specifies the default maximum duration (timeout) for streams to this service. Duration is computed from the beginning of the stream until the response has been completely processed, including all retries. A stream that does not complete in this duration is closed. If not specified, there will be no timeout limit, i.e. the maximum duration is infinite. This value can be overridden in the PathMatcher configuration of the UrlMap that references this backend service. This field is only allowed when the loadBalancingScheme of the backend service is INTERNAL_SELF_MANAGED.
        /// </summary>
        [Input("maxStreamDuration")]
        public Input<Inputs.DurationArgs>? MaxStreamDuration { get; set; }

        /// <summary>
        /// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The URL of the network to which this backend service belongs. This field can only be specified when the load balancing scheme is set to INTERNAL.
        /// </summary>
        [Input("network")]
        public Input<string>? Network { get; set; }

        /// <summary>
        /// Settings controlling the eviction of unhealthy hosts from the load balancing pool for the backend service. If not set, this feature is considered disabled. This field is applicable to either: - A regional backend service with the service_protocol set to HTTP, HTTPS, or HTTP2, and load_balancing_scheme set to INTERNAL_MANAGED. - A global backend service with the load_balancing_scheme set to INTERNAL_SELF_MANAGED. Not supported when the backend service is referenced by a URL map that is bound to target gRPC proxy that has validateForProxyless field set to true.
        /// </summary>
        [Input("outlierDetection")]
        public Input<Inputs.OutlierDetectionArgs>? OutlierDetection { get; set; }

        /// <summary>
        /// A named port on a backend instance group representing the port for communication to the backend VMs in that group. Required when the loadBalancingScheme is EXTERNAL (except Network Load Balancing), INTERNAL_MANAGED, or INTERNAL_SELF_MANAGED and the backends are instance groups. The named port must be defined on each backend instance group. This parameter has no meaning if the backends are NEGs. Backend services for Internal TCP/UDP Load Balancing and Network Load Balancing require you omit port_name.
        /// </summary>
        [Input("portName")]
        public Input<string>? PortName { get; set; }

        [Input("project", required: true)]
        public Input<string> Project { get; set; } = null!;

        /// <summary>
        /// The protocol this BackendService uses to communicate with backends. Possible values are HTTP, HTTPS, HTTP2, TCP, SSL, UDP or GRPC. depending on the chosen load balancer or Traffic Director configuration. Refer to the documentation for the load balancer or for Traffic Director for more information. Must be set to GRPC when the backend service is referenced by a URL map that is bound to target gRPC proxy.
        /// </summary>
        [Input("protocol")]
        public Input<Pulumi.GoogleNative.Compute.Beta.RegionBackendServiceProtocol>? Protocol { get; set; }

        /// <summary>
        /// [Output Only] URL of the region where the regional backend service resides. This field is not applicable to global backend services. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
        /// </summary>
        [Input("region", required: true)]
        public Input<string> Region { get; set; } = null!;

        [Input("requestId")]
        public Input<string>? RequestId { get; set; }

        /// <summary>
        /// [Output Only] The resource URL for the security policy associated with this backend service.
        /// </summary>
        [Input("securityPolicy")]
        public Input<string>? SecurityPolicy { get; set; }

        /// <summary>
        /// This field specifies the security policy that applies to this backend service. This field is applicable to either: - A regional backend service with the service_protocol set to HTTP, HTTPS, or HTTP2, and load_balancing_scheme set to INTERNAL_MANAGED. - A global backend service with the load_balancing_scheme set to INTERNAL_SELF_MANAGED. 
        /// </summary>
        [Input("securitySettings")]
        public Input<Inputs.SecuritySettingsArgs>? SecuritySettings { get; set; }

        /// <summary>
        /// [Output Only] Server-defined URL for the resource.
        /// </summary>
        [Input("selfLink")]
        public Input<string>? SelfLink { get; set; }

        /// <summary>
        /// Type of session affinity to use. The default is NONE. When the loadBalancingScheme is EXTERNAL: * For Network Load Balancing, the possible values are NONE, CLIENT_IP, CLIENT_IP_PROTO, or CLIENT_IP_PORT_PROTO. * For all other load balancers that use loadBalancingScheme=EXTERNAL, the possible values are NONE, CLIENT_IP, or GENERATED_COOKIE. * You can use GENERATED_COOKIE if the protocol is HTTP, HTTP2, or HTTPS. When the loadBalancingScheme is INTERNAL, possible values are NONE, CLIENT_IP, CLIENT_IP_PROTO, or CLIENT_IP_PORT_PROTO. When the loadBalancingScheme is INTERNAL_SELF_MANAGED, or INTERNAL_MANAGED, possible values are NONE, CLIENT_IP, GENERATED_COOKIE, HEADER_FIELD, or HTTP_COOKIE. Not supported when the backend service is referenced by a URL map that is bound to target gRPC proxy that has validateForProxyless field set to true.
        /// </summary>
        [Input("sessionAffinity")]
        public Input<Pulumi.GoogleNative.Compute.Beta.RegionBackendServiceSessionAffinity>? SessionAffinity { get; set; }

        [Input("subsetting")]
        public Input<Inputs.SubsettingArgs>? Subsetting { get; set; }

        /// <summary>
        /// Not supported when the backend service is referenced by a URL map that is bound to target gRPC proxy that has validateForProxyless field set to true. Instead, use maxStreamDuration.
        /// </summary>
        [Input("timeoutSec")]
        public Input<int>? TimeoutSec { get; set; }

        public RegionBackendServiceArgs()
        {
        }
    }
}
