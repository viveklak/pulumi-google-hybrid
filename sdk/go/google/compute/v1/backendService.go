// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a BackendService resource in the specified project using the data included in the request. For more information, see Backend services overview .
type BackendService struct {
	pulumi.CustomResourceState

	// Lifetime of cookies in seconds. This setting is applicable to external and internal HTTP(S) load balancers and Traffic Director and requires GENERATED_COOKIE or HTTP_COOKIE session affinity. If set to 0, the cookie is non-persistent and lasts only until the end of the browser session (or equivalent). The maximum allowed value is one day (86,400). Not supported when the backend service is referenced by a URL map that is bound to target gRPC proxy that has validateForProxyless field set to true.
	AffinityCookieTtlSec pulumi.IntOutput `pulumi:"affinityCookieTtlSec"`
	// The list of backends that serve this BackendService.
	Backends BackendResponseArrayOutput `pulumi:"backends"`
	// Cloud CDN configuration for this BackendService. Only available for specified load balancer types.
	CdnPolicy          BackendServiceCdnPolicyResponseOutput `pulumi:"cdnPolicy"`
	CircuitBreakers    CircuitBreakersResponseOutput         `pulumi:"circuitBreakers"`
	ConnectionDraining ConnectionDrainingResponseOutput      `pulumi:"connectionDraining"`
	// Connection Tracking configuration for this BackendService. Connection tracking policy settings are only available for Network Load Balancing and Internal TCP/UDP Load Balancing.
	ConnectionTrackingPolicy BackendServiceConnectionTrackingPolicyResponseOutput `pulumi:"connectionTrackingPolicy"`
	// Consistent Hash-based load balancing can be used to provide soft session affinity based on HTTP headers, cookies or other properties. This load balancing policy is applicable only for HTTP connections. The affinity to a particular destination host will be lost when one or more hosts are added/removed from the destination service. This field specifies parameters that control consistent hashing. This field is only applicable when localityLbPolicy is set to MAGLEV or RING_HASH. This field is applicable to either: - A regional backend service with the service_protocol set to HTTP, HTTPS, or HTTP2, and load_balancing_scheme set to INTERNAL_MANAGED. - A global backend service with the load_balancing_scheme set to INTERNAL_SELF_MANAGED.
	ConsistentHash ConsistentHashLoadBalancerSettingsResponseOutput `pulumi:"consistentHash"`
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringOutput `pulumi:"creationTimestamp"`
	// Headers that the load balancer adds to proxied requests. See [Creating custom headers](https://cloud.google.com/load-balancing/docs/custom-headers).
	CustomRequestHeaders pulumi.StringArrayOutput `pulumi:"customRequestHeaders"`
	// Headers that the load balancer adds to proxied responses. See [Creating custom headers](https://cloud.google.com/load-balancing/docs/custom-headers).
	CustomResponseHeaders pulumi.StringArrayOutput `pulumi:"customResponseHeaders"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description pulumi.StringOutput `pulumi:"description"`
	// The resource URL for the edge security policy associated with this backend service.
	EdgeSecurityPolicy pulumi.StringOutput `pulumi:"edgeSecurityPolicy"`
	// If true, enables Cloud CDN for the backend service of an external HTTP(S) load balancer.
	EnableCDN pulumi.BoolOutput `pulumi:"enableCDN"`
	// Requires at least one backend instance group to be defined as a backup (failover) backend. For load balancers that have configurable failover: [Internal TCP/UDP Load Balancing](https://cloud.google.com/load-balancing/docs/internal/failover-overview) and [external TCP/UDP Load Balancing](https://cloud.google.com/load-balancing/docs/network/networklb-failover-overview).
	FailoverPolicy BackendServiceFailoverPolicyResponseOutput `pulumi:"failoverPolicy"`
	// Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking. This field will be ignored when inserting a BackendService. An up-to-date fingerprint must be provided in order to update the BackendService, otherwise the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make a get() request to retrieve a BackendService.
	Fingerprint pulumi.StringOutput `pulumi:"fingerprint"`
	// The list of URLs to the healthChecks, httpHealthChecks (legacy), or httpsHealthChecks (legacy) resource for health checking this backend service. Not all backend services support legacy health checks. See Load balancer guide. Currently, at most one health check can be specified for each backend service. Backend services with instance group or zonal NEG backends must have a health check. Backend services with internet or serverless NEG backends must not have a health check.
	HealthChecks pulumi.StringArrayOutput `pulumi:"healthChecks"`
	// The configurations for Identity-Aware Proxy on this resource. Not available for Internal TCP/UDP Load Balancing and Network Load Balancing.
	Iap BackendServiceIAPResponseOutput `pulumi:"iap"`
	// Type of resource. Always compute#backendService for backend services.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Specifies the load balancer type. A backend service created for one type of load balancer cannot be used with another. For more information, refer to Choosing a load balancer.
	LoadBalancingScheme pulumi.StringOutput `pulumi:"loadBalancingScheme"`
	// The load balancing algorithm used within the scope of the locality. The possible values are: - ROUND_ROBIN: This is a simple policy in which each healthy backend is selected in round robin order. This is the default. - LEAST_REQUEST: An O(1) algorithm which selects two random healthy hosts and picks the host which has fewer active requests. - RING_HASH: The ring/modulo hash load balancer implements consistent hashing to backends. The algorithm has the property that the addition/removal of a host from a set of N hosts only affects 1/N of the requests. - RANDOM: The load balancer selects a random healthy host. - ORIGINAL_DESTINATION: Backend host is selected based on the client connection metadata, i.e., connections are opened to the same address as the destination address of the incoming connection before the connection was redirected to the load balancer. - MAGLEV: used as a drop in replacement for the ring hash load balancer. Maglev is not as stable as ring hash but has faster table lookup build times and host selection times. For more information about Maglev, see https://ai.google/research/pubs/pub44824 This field is applicable to either: - A regional backend service with the service_protocol set to HTTP, HTTPS, or HTTP2, and load_balancing_scheme set to INTERNAL_MANAGED. - A global backend service with the load_balancing_scheme set to INTERNAL_SELF_MANAGED. If sessionAffinity is not NONE, and this field is not set to MAGLEV or RING_HASH, session affinity settings will not take effect. Only ROUND_ROBIN and RING_HASH are supported when the backend service is referenced by a URL map that is bound to target gRPC proxy that has validateForProxyless field set to true.
	LocalityLbPolicy pulumi.StringOutput `pulumi:"localityLbPolicy"`
	// This field denotes the logging options for the load balancer traffic served by this backend service. If logging is enabled, logs will be exported to Stackdriver.
	LogConfig BackendServiceLogConfigResponseOutput `pulumi:"logConfig"`
	// Specifies the default maximum duration (timeout) for streams to this service. Duration is computed from the beginning of the stream until the response has been completely processed, including all retries. A stream that does not complete in this duration is closed. If not specified, there will be no timeout limit, i.e. the maximum duration is infinite. This value can be overridden in the PathMatcher configuration of the UrlMap that references this backend service. This field is only allowed when the loadBalancingScheme of the backend service is INTERNAL_SELF_MANAGED.
	MaxStreamDuration DurationResponseOutput `pulumi:"maxStreamDuration"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name pulumi.StringOutput `pulumi:"name"`
	// The URL of the network to which this backend service belongs. This field can only be specified when the load balancing scheme is set to INTERNAL.
	Network pulumi.StringOutput `pulumi:"network"`
	// Settings controlling the eviction of unhealthy hosts from the load balancing pool for the backend service. If not set, this feature is considered disabled. This field is applicable to either: - A regional backend service with the service_protocol set to HTTP, HTTPS, or HTTP2, and load_balancing_scheme set to INTERNAL_MANAGED. - A global backend service with the load_balancing_scheme set to INTERNAL_SELF_MANAGED. Not supported when the backend service is referenced by a URL map that is bound to target gRPC proxy that has validateForProxyless field set to true.
	OutlierDetection OutlierDetectionResponseOutput `pulumi:"outlierDetection"`
	// Deprecated in favor of portName. The TCP port to connect on the backend. The default value is 80. For Internal TCP/UDP Load Balancing and Network Load Balancing, omit port.
	//
	// Deprecated: Deprecated in favor of portName. The TCP port to connect on the backend. The default value is 80. For Internal TCP/UDP Load Balancing and Network Load Balancing, omit port.
	Port pulumi.IntOutput `pulumi:"port"`
	// A named port on a backend instance group representing the port for communication to the backend VMs in that group. The named port must be [defined on each backend instance group](https://cloud.google.com/load-balancing/docs/backend-service#named_ports). This parameter has no meaning if the backends are NEGs. For Internal TCP/UDP Load Balancing and Network Load Balancing, omit port_name.
	PortName pulumi.StringOutput `pulumi:"portName"`
	// The protocol this BackendService uses to communicate with backends. Possible values are HTTP, HTTPS, HTTP2, TCP, SSL, UDP or GRPC. depending on the chosen load balancer or Traffic Director configuration. Refer to the documentation for the load balancers or for Traffic Director for more information. Must be set to GRPC when the backend service is referenced by a URL map that is bound to target gRPC proxy.
	Protocol pulumi.StringOutput `pulumi:"protocol"`
	// URL of the region where the regional backend service resides. This field is not applicable to global backend services. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
	Region pulumi.StringOutput `pulumi:"region"`
	// The resource URL for the security policy associated with this backend service.
	SecurityPolicy pulumi.StringOutput `pulumi:"securityPolicy"`
	// This field specifies the security settings that apply to this backend service. This field is applicable to a global backend service with the load_balancing_scheme set to INTERNAL_SELF_MANAGED.
	SecuritySettings SecuritySettingsResponseOutput `pulumi:"securitySettings"`
	// Server-defined URL for the resource.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
	// Type of session affinity to use. The default is NONE. Only NONE and HEADER_FIELD are supported when the backend service is referenced by a URL map that is bound to target gRPC proxy that has validateForProxyless field set to true. For more details, see: [Session Affinity](https://cloud.google.com/load-balancing/docs/backend-service#session_affinity).
	SessionAffinity pulumi.StringOutput      `pulumi:"sessionAffinity"`
	Subsetting      SubsettingResponseOutput `pulumi:"subsetting"`
	// The backend service timeout has a different meaning depending on the type of load balancer. For more information see, Backend service settings. The default is 30 seconds. The full range of timeout values allowed goes from 1 through 2,147,483,647 seconds. This value can be overridden in the PathMatcher configuration of the UrlMap that references this backend service. Not supported when the backend service is referenced by a URL map that is bound to target gRPC proxy that has validateForProxyless field set to true. Instead, use maxStreamDuration.
	TimeoutSec pulumi.IntOutput `pulumi:"timeoutSec"`
}

// NewBackendService registers a new resource with the given unique name, arguments, and options.
func NewBackendService(ctx *pulumi.Context,
	name string, args *BackendServiceArgs, opts ...pulumi.ResourceOption) (*BackendService, error) {
	if args == nil {
		args = &BackendServiceArgs{}
	}

	var resource BackendService
	err := ctx.RegisterResource("google-native:compute/v1:BackendService", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBackendService gets an existing BackendService resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBackendService(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BackendServiceState, opts ...pulumi.ResourceOption) (*BackendService, error) {
	var resource BackendService
	err := ctx.ReadResource("google-native:compute/v1:BackendService", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BackendService resources.
type backendServiceState struct {
}

type BackendServiceState struct {
}

func (BackendServiceState) ElementType() reflect.Type {
	return reflect.TypeOf((*backendServiceState)(nil)).Elem()
}

type backendServiceArgs struct {
	// Lifetime of cookies in seconds. This setting is applicable to external and internal HTTP(S) load balancers and Traffic Director and requires GENERATED_COOKIE or HTTP_COOKIE session affinity. If set to 0, the cookie is non-persistent and lasts only until the end of the browser session (or equivalent). The maximum allowed value is one day (86,400). Not supported when the backend service is referenced by a URL map that is bound to target gRPC proxy that has validateForProxyless field set to true.
	AffinityCookieTtlSec *int `pulumi:"affinityCookieTtlSec"`
	// The list of backends that serve this BackendService.
	Backends []Backend `pulumi:"backends"`
	// Cloud CDN configuration for this BackendService. Only available for specified load balancer types.
	CdnPolicy          *BackendServiceCdnPolicy `pulumi:"cdnPolicy"`
	CircuitBreakers    *CircuitBreakers         `pulumi:"circuitBreakers"`
	ConnectionDraining *ConnectionDraining      `pulumi:"connectionDraining"`
	// Connection Tracking configuration for this BackendService. Connection tracking policy settings are only available for Network Load Balancing and Internal TCP/UDP Load Balancing.
	ConnectionTrackingPolicy *BackendServiceConnectionTrackingPolicy `pulumi:"connectionTrackingPolicy"`
	// Consistent Hash-based load balancing can be used to provide soft session affinity based on HTTP headers, cookies or other properties. This load balancing policy is applicable only for HTTP connections. The affinity to a particular destination host will be lost when one or more hosts are added/removed from the destination service. This field specifies parameters that control consistent hashing. This field is only applicable when localityLbPolicy is set to MAGLEV or RING_HASH. This field is applicable to either: - A regional backend service with the service_protocol set to HTTP, HTTPS, or HTTP2, and load_balancing_scheme set to INTERNAL_MANAGED. - A global backend service with the load_balancing_scheme set to INTERNAL_SELF_MANAGED.
	ConsistentHash *ConsistentHashLoadBalancerSettings `pulumi:"consistentHash"`
	// Headers that the load balancer adds to proxied requests. See [Creating custom headers](https://cloud.google.com/load-balancing/docs/custom-headers).
	CustomRequestHeaders []string `pulumi:"customRequestHeaders"`
	// Headers that the load balancer adds to proxied responses. See [Creating custom headers](https://cloud.google.com/load-balancing/docs/custom-headers).
	CustomResponseHeaders []string `pulumi:"customResponseHeaders"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description *string `pulumi:"description"`
	// If true, enables Cloud CDN for the backend service of an external HTTP(S) load balancer.
	EnableCDN *bool `pulumi:"enableCDN"`
	// Requires at least one backend instance group to be defined as a backup (failover) backend. For load balancers that have configurable failover: [Internal TCP/UDP Load Balancing](https://cloud.google.com/load-balancing/docs/internal/failover-overview) and [external TCP/UDP Load Balancing](https://cloud.google.com/load-balancing/docs/network/networklb-failover-overview).
	FailoverPolicy *BackendServiceFailoverPolicy `pulumi:"failoverPolicy"`
	// The list of URLs to the healthChecks, httpHealthChecks (legacy), or httpsHealthChecks (legacy) resource for health checking this backend service. Not all backend services support legacy health checks. See Load balancer guide. Currently, at most one health check can be specified for each backend service. Backend services with instance group or zonal NEG backends must have a health check. Backend services with internet or serverless NEG backends must not have a health check.
	HealthChecks []string `pulumi:"healthChecks"`
	// The configurations for Identity-Aware Proxy on this resource. Not available for Internal TCP/UDP Load Balancing and Network Load Balancing.
	Iap *BackendServiceIAP `pulumi:"iap"`
	// Specifies the load balancer type. A backend service created for one type of load balancer cannot be used with another. For more information, refer to Choosing a load balancer.
	LoadBalancingScheme *BackendServiceLoadBalancingScheme `pulumi:"loadBalancingScheme"`
	// The load balancing algorithm used within the scope of the locality. The possible values are: - ROUND_ROBIN: This is a simple policy in which each healthy backend is selected in round robin order. This is the default. - LEAST_REQUEST: An O(1) algorithm which selects two random healthy hosts and picks the host which has fewer active requests. - RING_HASH: The ring/modulo hash load balancer implements consistent hashing to backends. The algorithm has the property that the addition/removal of a host from a set of N hosts only affects 1/N of the requests. - RANDOM: The load balancer selects a random healthy host. - ORIGINAL_DESTINATION: Backend host is selected based on the client connection metadata, i.e., connections are opened to the same address as the destination address of the incoming connection before the connection was redirected to the load balancer. - MAGLEV: used as a drop in replacement for the ring hash load balancer. Maglev is not as stable as ring hash but has faster table lookup build times and host selection times. For more information about Maglev, see https://ai.google/research/pubs/pub44824 This field is applicable to either: - A regional backend service with the service_protocol set to HTTP, HTTPS, or HTTP2, and load_balancing_scheme set to INTERNAL_MANAGED. - A global backend service with the load_balancing_scheme set to INTERNAL_SELF_MANAGED. If sessionAffinity is not NONE, and this field is not set to MAGLEV or RING_HASH, session affinity settings will not take effect. Only ROUND_ROBIN and RING_HASH are supported when the backend service is referenced by a URL map that is bound to target gRPC proxy that has validateForProxyless field set to true.
	LocalityLbPolicy *BackendServiceLocalityLbPolicy `pulumi:"localityLbPolicy"`
	// This field denotes the logging options for the load balancer traffic served by this backend service. If logging is enabled, logs will be exported to Stackdriver.
	LogConfig *BackendServiceLogConfig `pulumi:"logConfig"`
	// Specifies the default maximum duration (timeout) for streams to this service. Duration is computed from the beginning of the stream until the response has been completely processed, including all retries. A stream that does not complete in this duration is closed. If not specified, there will be no timeout limit, i.e. the maximum duration is infinite. This value can be overridden in the PathMatcher configuration of the UrlMap that references this backend service. This field is only allowed when the loadBalancingScheme of the backend service is INTERNAL_SELF_MANAGED.
	MaxStreamDuration *Duration `pulumi:"maxStreamDuration"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name *string `pulumi:"name"`
	// The URL of the network to which this backend service belongs. This field can only be specified when the load balancing scheme is set to INTERNAL.
	Network *string `pulumi:"network"`
	// Settings controlling the eviction of unhealthy hosts from the load balancing pool for the backend service. If not set, this feature is considered disabled. This field is applicable to either: - A regional backend service with the service_protocol set to HTTP, HTTPS, or HTTP2, and load_balancing_scheme set to INTERNAL_MANAGED. - A global backend service with the load_balancing_scheme set to INTERNAL_SELF_MANAGED. Not supported when the backend service is referenced by a URL map that is bound to target gRPC proxy that has validateForProxyless field set to true.
	OutlierDetection *OutlierDetection `pulumi:"outlierDetection"`
	// Deprecated in favor of portName. The TCP port to connect on the backend. The default value is 80. For Internal TCP/UDP Load Balancing and Network Load Balancing, omit port.
	//
	// Deprecated: Deprecated in favor of portName. The TCP port to connect on the backend. The default value is 80. For Internal TCP/UDP Load Balancing and Network Load Balancing, omit port.
	Port *int `pulumi:"port"`
	// A named port on a backend instance group representing the port for communication to the backend VMs in that group. The named port must be [defined on each backend instance group](https://cloud.google.com/load-balancing/docs/backend-service#named_ports). This parameter has no meaning if the backends are NEGs. For Internal TCP/UDP Load Balancing and Network Load Balancing, omit port_name.
	PortName *string `pulumi:"portName"`
	Project  *string `pulumi:"project"`
	// The protocol this BackendService uses to communicate with backends. Possible values are HTTP, HTTPS, HTTP2, TCP, SSL, UDP or GRPC. depending on the chosen load balancer or Traffic Director configuration. Refer to the documentation for the load balancers or for Traffic Director for more information. Must be set to GRPC when the backend service is referenced by a URL map that is bound to target gRPC proxy.
	Protocol *BackendServiceProtocol `pulumi:"protocol"`
	// An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
	RequestId *string `pulumi:"requestId"`
	// This field specifies the security settings that apply to this backend service. This field is applicable to a global backend service with the load_balancing_scheme set to INTERNAL_SELF_MANAGED.
	SecuritySettings *SecuritySettings `pulumi:"securitySettings"`
	// Type of session affinity to use. The default is NONE. Only NONE and HEADER_FIELD are supported when the backend service is referenced by a URL map that is bound to target gRPC proxy that has validateForProxyless field set to true. For more details, see: [Session Affinity](https://cloud.google.com/load-balancing/docs/backend-service#session_affinity).
	SessionAffinity *BackendServiceSessionAffinity `pulumi:"sessionAffinity"`
	Subsetting      *Subsetting                    `pulumi:"subsetting"`
	// The backend service timeout has a different meaning depending on the type of load balancer. For more information see, Backend service settings. The default is 30 seconds. The full range of timeout values allowed goes from 1 through 2,147,483,647 seconds. This value can be overridden in the PathMatcher configuration of the UrlMap that references this backend service. Not supported when the backend service is referenced by a URL map that is bound to target gRPC proxy that has validateForProxyless field set to true. Instead, use maxStreamDuration.
	TimeoutSec *int `pulumi:"timeoutSec"`
}

// The set of arguments for constructing a BackendService resource.
type BackendServiceArgs struct {
	// Lifetime of cookies in seconds. This setting is applicable to external and internal HTTP(S) load balancers and Traffic Director and requires GENERATED_COOKIE or HTTP_COOKIE session affinity. If set to 0, the cookie is non-persistent and lasts only until the end of the browser session (or equivalent). The maximum allowed value is one day (86,400). Not supported when the backend service is referenced by a URL map that is bound to target gRPC proxy that has validateForProxyless field set to true.
	AffinityCookieTtlSec pulumi.IntPtrInput
	// The list of backends that serve this BackendService.
	Backends BackendArrayInput
	// Cloud CDN configuration for this BackendService. Only available for specified load balancer types.
	CdnPolicy          BackendServiceCdnPolicyPtrInput
	CircuitBreakers    CircuitBreakersPtrInput
	ConnectionDraining ConnectionDrainingPtrInput
	// Connection Tracking configuration for this BackendService. Connection tracking policy settings are only available for Network Load Balancing and Internal TCP/UDP Load Balancing.
	ConnectionTrackingPolicy BackendServiceConnectionTrackingPolicyPtrInput
	// Consistent Hash-based load balancing can be used to provide soft session affinity based on HTTP headers, cookies or other properties. This load balancing policy is applicable only for HTTP connections. The affinity to a particular destination host will be lost when one or more hosts are added/removed from the destination service. This field specifies parameters that control consistent hashing. This field is only applicable when localityLbPolicy is set to MAGLEV or RING_HASH. This field is applicable to either: - A regional backend service with the service_protocol set to HTTP, HTTPS, or HTTP2, and load_balancing_scheme set to INTERNAL_MANAGED. - A global backend service with the load_balancing_scheme set to INTERNAL_SELF_MANAGED.
	ConsistentHash ConsistentHashLoadBalancerSettingsPtrInput
	// Headers that the load balancer adds to proxied requests. See [Creating custom headers](https://cloud.google.com/load-balancing/docs/custom-headers).
	CustomRequestHeaders pulumi.StringArrayInput
	// Headers that the load balancer adds to proxied responses. See [Creating custom headers](https://cloud.google.com/load-balancing/docs/custom-headers).
	CustomResponseHeaders pulumi.StringArrayInput
	// An optional description of this resource. Provide this property when you create the resource.
	Description pulumi.StringPtrInput
	// If true, enables Cloud CDN for the backend service of an external HTTP(S) load balancer.
	EnableCDN pulumi.BoolPtrInput
	// Requires at least one backend instance group to be defined as a backup (failover) backend. For load balancers that have configurable failover: [Internal TCP/UDP Load Balancing](https://cloud.google.com/load-balancing/docs/internal/failover-overview) and [external TCP/UDP Load Balancing](https://cloud.google.com/load-balancing/docs/network/networklb-failover-overview).
	FailoverPolicy BackendServiceFailoverPolicyPtrInput
	// The list of URLs to the healthChecks, httpHealthChecks (legacy), or httpsHealthChecks (legacy) resource for health checking this backend service. Not all backend services support legacy health checks. See Load balancer guide. Currently, at most one health check can be specified for each backend service. Backend services with instance group or zonal NEG backends must have a health check. Backend services with internet or serverless NEG backends must not have a health check.
	HealthChecks pulumi.StringArrayInput
	// The configurations for Identity-Aware Proxy on this resource. Not available for Internal TCP/UDP Load Balancing and Network Load Balancing.
	Iap BackendServiceIAPPtrInput
	// Specifies the load balancer type. A backend service created for one type of load balancer cannot be used with another. For more information, refer to Choosing a load balancer.
	LoadBalancingScheme BackendServiceLoadBalancingSchemePtrInput
	// The load balancing algorithm used within the scope of the locality. The possible values are: - ROUND_ROBIN: This is a simple policy in which each healthy backend is selected in round robin order. This is the default. - LEAST_REQUEST: An O(1) algorithm which selects two random healthy hosts and picks the host which has fewer active requests. - RING_HASH: The ring/modulo hash load balancer implements consistent hashing to backends. The algorithm has the property that the addition/removal of a host from a set of N hosts only affects 1/N of the requests. - RANDOM: The load balancer selects a random healthy host. - ORIGINAL_DESTINATION: Backend host is selected based on the client connection metadata, i.e., connections are opened to the same address as the destination address of the incoming connection before the connection was redirected to the load balancer. - MAGLEV: used as a drop in replacement for the ring hash load balancer. Maglev is not as stable as ring hash but has faster table lookup build times and host selection times. For more information about Maglev, see https://ai.google/research/pubs/pub44824 This field is applicable to either: - A regional backend service with the service_protocol set to HTTP, HTTPS, or HTTP2, and load_balancing_scheme set to INTERNAL_MANAGED. - A global backend service with the load_balancing_scheme set to INTERNAL_SELF_MANAGED. If sessionAffinity is not NONE, and this field is not set to MAGLEV or RING_HASH, session affinity settings will not take effect. Only ROUND_ROBIN and RING_HASH are supported when the backend service is referenced by a URL map that is bound to target gRPC proxy that has validateForProxyless field set to true.
	LocalityLbPolicy BackendServiceLocalityLbPolicyPtrInput
	// This field denotes the logging options for the load balancer traffic served by this backend service. If logging is enabled, logs will be exported to Stackdriver.
	LogConfig BackendServiceLogConfigPtrInput
	// Specifies the default maximum duration (timeout) for streams to this service. Duration is computed from the beginning of the stream until the response has been completely processed, including all retries. A stream that does not complete in this duration is closed. If not specified, there will be no timeout limit, i.e. the maximum duration is infinite. This value can be overridden in the PathMatcher configuration of the UrlMap that references this backend service. This field is only allowed when the loadBalancingScheme of the backend service is INTERNAL_SELF_MANAGED.
	MaxStreamDuration DurationPtrInput
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name pulumi.StringPtrInput
	// The URL of the network to which this backend service belongs. This field can only be specified when the load balancing scheme is set to INTERNAL.
	Network pulumi.StringPtrInput
	// Settings controlling the eviction of unhealthy hosts from the load balancing pool for the backend service. If not set, this feature is considered disabled. This field is applicable to either: - A regional backend service with the service_protocol set to HTTP, HTTPS, or HTTP2, and load_balancing_scheme set to INTERNAL_MANAGED. - A global backend service with the load_balancing_scheme set to INTERNAL_SELF_MANAGED. Not supported when the backend service is referenced by a URL map that is bound to target gRPC proxy that has validateForProxyless field set to true.
	OutlierDetection OutlierDetectionPtrInput
	// Deprecated in favor of portName. The TCP port to connect on the backend. The default value is 80. For Internal TCP/UDP Load Balancing and Network Load Balancing, omit port.
	//
	// Deprecated: Deprecated in favor of portName. The TCP port to connect on the backend. The default value is 80. For Internal TCP/UDP Load Balancing and Network Load Balancing, omit port.
	Port pulumi.IntPtrInput
	// A named port on a backend instance group representing the port for communication to the backend VMs in that group. The named port must be [defined on each backend instance group](https://cloud.google.com/load-balancing/docs/backend-service#named_ports). This parameter has no meaning if the backends are NEGs. For Internal TCP/UDP Load Balancing and Network Load Balancing, omit port_name.
	PortName pulumi.StringPtrInput
	Project  pulumi.StringPtrInput
	// The protocol this BackendService uses to communicate with backends. Possible values are HTTP, HTTPS, HTTP2, TCP, SSL, UDP or GRPC. depending on the chosen load balancer or Traffic Director configuration. Refer to the documentation for the load balancers or for Traffic Director for more information. Must be set to GRPC when the backend service is referenced by a URL map that is bound to target gRPC proxy.
	Protocol BackendServiceProtocolPtrInput
	// An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
	RequestId pulumi.StringPtrInput
	// This field specifies the security settings that apply to this backend service. This field is applicable to a global backend service with the load_balancing_scheme set to INTERNAL_SELF_MANAGED.
	SecuritySettings SecuritySettingsPtrInput
	// Type of session affinity to use. The default is NONE. Only NONE and HEADER_FIELD are supported when the backend service is referenced by a URL map that is bound to target gRPC proxy that has validateForProxyless field set to true. For more details, see: [Session Affinity](https://cloud.google.com/load-balancing/docs/backend-service#session_affinity).
	SessionAffinity BackendServiceSessionAffinityPtrInput
	Subsetting      SubsettingPtrInput
	// The backend service timeout has a different meaning depending on the type of load balancer. For more information see, Backend service settings. The default is 30 seconds. The full range of timeout values allowed goes from 1 through 2,147,483,647 seconds. This value can be overridden in the PathMatcher configuration of the UrlMap that references this backend service. Not supported when the backend service is referenced by a URL map that is bound to target gRPC proxy that has validateForProxyless field set to true. Instead, use maxStreamDuration.
	TimeoutSec pulumi.IntPtrInput
}

func (BackendServiceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*backendServiceArgs)(nil)).Elem()
}

type BackendServiceInput interface {
	pulumi.Input

	ToBackendServiceOutput() BackendServiceOutput
	ToBackendServiceOutputWithContext(ctx context.Context) BackendServiceOutput
}

func (*BackendService) ElementType() reflect.Type {
	return reflect.TypeOf((**BackendService)(nil)).Elem()
}

func (i *BackendService) ToBackendServiceOutput() BackendServiceOutput {
	return i.ToBackendServiceOutputWithContext(context.Background())
}

func (i *BackendService) ToBackendServiceOutputWithContext(ctx context.Context) BackendServiceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackendServiceOutput)
}

type BackendServiceOutput struct{ *pulumi.OutputState }

func (BackendServiceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BackendService)(nil)).Elem()
}

func (o BackendServiceOutput) ToBackendServiceOutput() BackendServiceOutput {
	return o
}

func (o BackendServiceOutput) ToBackendServiceOutputWithContext(ctx context.Context) BackendServiceOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BackendServiceInput)(nil)).Elem(), &BackendService{})
	pulumi.RegisterOutputType(BackendServiceOutput{})
}
