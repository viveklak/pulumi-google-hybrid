# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities
from . import outputs

__all__ = [
    'GetRegionBackendServiceResult',
    'AwaitableGetRegionBackendServiceResult',
    'get_region_backend_service',
    'get_region_backend_service_output',
]

@pulumi.output_type
class GetRegionBackendServiceResult:
    def __init__(__self__, affinity_cookie_ttl_sec=None, backends=None, cdn_policy=None, circuit_breakers=None, compression_mode=None, connection_draining=None, connection_tracking_policy=None, consistent_hash=None, creation_timestamp=None, custom_request_headers=None, custom_response_headers=None, description=None, edge_security_policy=None, enable_cdn=None, failover_policy=None, fingerprint=None, health_checks=None, iap=None, kind=None, load_balancing_scheme=None, locality_lb_policy=None, log_config=None, max_stream_duration=None, name=None, network=None, outlier_detection=None, port=None, port_name=None, protocol=None, region=None, security_policy=None, security_settings=None, self_link=None, service_bindings=None, session_affinity=None, subsetting=None, timeout_sec=None):
        if affinity_cookie_ttl_sec and not isinstance(affinity_cookie_ttl_sec, int):
            raise TypeError("Expected argument 'affinity_cookie_ttl_sec' to be a int")
        pulumi.set(__self__, "affinity_cookie_ttl_sec", affinity_cookie_ttl_sec)
        if backends and not isinstance(backends, list):
            raise TypeError("Expected argument 'backends' to be a list")
        pulumi.set(__self__, "backends", backends)
        if cdn_policy and not isinstance(cdn_policy, dict):
            raise TypeError("Expected argument 'cdn_policy' to be a dict")
        pulumi.set(__self__, "cdn_policy", cdn_policy)
        if circuit_breakers and not isinstance(circuit_breakers, dict):
            raise TypeError("Expected argument 'circuit_breakers' to be a dict")
        pulumi.set(__self__, "circuit_breakers", circuit_breakers)
        if compression_mode and not isinstance(compression_mode, str):
            raise TypeError("Expected argument 'compression_mode' to be a str")
        pulumi.set(__self__, "compression_mode", compression_mode)
        if connection_draining and not isinstance(connection_draining, dict):
            raise TypeError("Expected argument 'connection_draining' to be a dict")
        pulumi.set(__self__, "connection_draining", connection_draining)
        if connection_tracking_policy and not isinstance(connection_tracking_policy, dict):
            raise TypeError("Expected argument 'connection_tracking_policy' to be a dict")
        pulumi.set(__self__, "connection_tracking_policy", connection_tracking_policy)
        if consistent_hash and not isinstance(consistent_hash, dict):
            raise TypeError("Expected argument 'consistent_hash' to be a dict")
        pulumi.set(__self__, "consistent_hash", consistent_hash)
        if creation_timestamp and not isinstance(creation_timestamp, str):
            raise TypeError("Expected argument 'creation_timestamp' to be a str")
        pulumi.set(__self__, "creation_timestamp", creation_timestamp)
        if custom_request_headers and not isinstance(custom_request_headers, list):
            raise TypeError("Expected argument 'custom_request_headers' to be a list")
        pulumi.set(__self__, "custom_request_headers", custom_request_headers)
        if custom_response_headers and not isinstance(custom_response_headers, list):
            raise TypeError("Expected argument 'custom_response_headers' to be a list")
        pulumi.set(__self__, "custom_response_headers", custom_response_headers)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if edge_security_policy and not isinstance(edge_security_policy, str):
            raise TypeError("Expected argument 'edge_security_policy' to be a str")
        pulumi.set(__self__, "edge_security_policy", edge_security_policy)
        if enable_cdn and not isinstance(enable_cdn, bool):
            raise TypeError("Expected argument 'enable_cdn' to be a bool")
        pulumi.set(__self__, "enable_cdn", enable_cdn)
        if failover_policy and not isinstance(failover_policy, dict):
            raise TypeError("Expected argument 'failover_policy' to be a dict")
        pulumi.set(__self__, "failover_policy", failover_policy)
        if fingerprint and not isinstance(fingerprint, str):
            raise TypeError("Expected argument 'fingerprint' to be a str")
        pulumi.set(__self__, "fingerprint", fingerprint)
        if health_checks and not isinstance(health_checks, list):
            raise TypeError("Expected argument 'health_checks' to be a list")
        pulumi.set(__self__, "health_checks", health_checks)
        if iap and not isinstance(iap, dict):
            raise TypeError("Expected argument 'iap' to be a dict")
        pulumi.set(__self__, "iap", iap)
        if kind and not isinstance(kind, str):
            raise TypeError("Expected argument 'kind' to be a str")
        pulumi.set(__self__, "kind", kind)
        if load_balancing_scheme and not isinstance(load_balancing_scheme, str):
            raise TypeError("Expected argument 'load_balancing_scheme' to be a str")
        pulumi.set(__self__, "load_balancing_scheme", load_balancing_scheme)
        if locality_lb_policy and not isinstance(locality_lb_policy, str):
            raise TypeError("Expected argument 'locality_lb_policy' to be a str")
        pulumi.set(__self__, "locality_lb_policy", locality_lb_policy)
        if log_config and not isinstance(log_config, dict):
            raise TypeError("Expected argument 'log_config' to be a dict")
        pulumi.set(__self__, "log_config", log_config)
        if max_stream_duration and not isinstance(max_stream_duration, dict):
            raise TypeError("Expected argument 'max_stream_duration' to be a dict")
        pulumi.set(__self__, "max_stream_duration", max_stream_duration)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if network and not isinstance(network, str):
            raise TypeError("Expected argument 'network' to be a str")
        pulumi.set(__self__, "network", network)
        if outlier_detection and not isinstance(outlier_detection, dict):
            raise TypeError("Expected argument 'outlier_detection' to be a dict")
        pulumi.set(__self__, "outlier_detection", outlier_detection)
        if port and not isinstance(port, int):
            raise TypeError("Expected argument 'port' to be a int")
        if port is not None:
            warnings.warn("""Deprecated in favor of portName. The TCP port to connect on the backend. The default value is 80. For Internal TCP/UDP Load Balancing and Network Load Balancing, omit port.""", DeprecationWarning)
            pulumi.log.warn("""port is deprecated: Deprecated in favor of portName. The TCP port to connect on the backend. The default value is 80. For Internal TCP/UDP Load Balancing and Network Load Balancing, omit port.""")

        pulumi.set(__self__, "port", port)
        if port_name and not isinstance(port_name, str):
            raise TypeError("Expected argument 'port_name' to be a str")
        pulumi.set(__self__, "port_name", port_name)
        if protocol and not isinstance(protocol, str):
            raise TypeError("Expected argument 'protocol' to be a str")
        pulumi.set(__self__, "protocol", protocol)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)
        if security_policy and not isinstance(security_policy, str):
            raise TypeError("Expected argument 'security_policy' to be a str")
        pulumi.set(__self__, "security_policy", security_policy)
        if security_settings and not isinstance(security_settings, dict):
            raise TypeError("Expected argument 'security_settings' to be a dict")
        pulumi.set(__self__, "security_settings", security_settings)
        if self_link and not isinstance(self_link, str):
            raise TypeError("Expected argument 'self_link' to be a str")
        pulumi.set(__self__, "self_link", self_link)
        if service_bindings and not isinstance(service_bindings, list):
            raise TypeError("Expected argument 'service_bindings' to be a list")
        pulumi.set(__self__, "service_bindings", service_bindings)
        if session_affinity and not isinstance(session_affinity, str):
            raise TypeError("Expected argument 'session_affinity' to be a str")
        pulumi.set(__self__, "session_affinity", session_affinity)
        if subsetting and not isinstance(subsetting, dict):
            raise TypeError("Expected argument 'subsetting' to be a dict")
        pulumi.set(__self__, "subsetting", subsetting)
        if timeout_sec and not isinstance(timeout_sec, int):
            raise TypeError("Expected argument 'timeout_sec' to be a int")
        pulumi.set(__self__, "timeout_sec", timeout_sec)

    @property
    @pulumi.getter(name="affinityCookieTtlSec")
    def affinity_cookie_ttl_sec(self) -> int:
        """
        Lifetime of cookies in seconds. This setting is applicable to external and internal HTTP(S) load balancers and Traffic Director and requires GENERATED_COOKIE or HTTP_COOKIE session affinity. If set to 0, the cookie is non-persistent and lasts only until the end of the browser session (or equivalent). The maximum allowed value is one day (86,400). Not supported when the backend service is referenced by a URL map that is bound to target gRPC proxy that has validateForProxyless field set to true.
        """
        return pulumi.get(self, "affinity_cookie_ttl_sec")

    @property
    @pulumi.getter
    def backends(self) -> Sequence['outputs.BackendResponse']:
        """
        The list of backends that serve this BackendService.
        """
        return pulumi.get(self, "backends")

    @property
    @pulumi.getter(name="cdnPolicy")
    def cdn_policy(self) -> 'outputs.BackendServiceCdnPolicyResponse':
        """
        Cloud CDN configuration for this BackendService. Only available for specified load balancer types.
        """
        return pulumi.get(self, "cdn_policy")

    @property
    @pulumi.getter(name="circuitBreakers")
    def circuit_breakers(self) -> 'outputs.CircuitBreakersResponse':
        return pulumi.get(self, "circuit_breakers")

    @property
    @pulumi.getter(name="compressionMode")
    def compression_mode(self) -> str:
        """
        Compress text responses using Brotli or gzip compression, based on the client's Accept-Encoding header.
        """
        return pulumi.get(self, "compression_mode")

    @property
    @pulumi.getter(name="connectionDraining")
    def connection_draining(self) -> 'outputs.ConnectionDrainingResponse':
        return pulumi.get(self, "connection_draining")

    @property
    @pulumi.getter(name="connectionTrackingPolicy")
    def connection_tracking_policy(self) -> 'outputs.BackendServiceConnectionTrackingPolicyResponse':
        """
        Connection Tracking configuration for this BackendService. Connection tracking policy settings are only available for Network Load Balancing and Internal TCP/UDP Load Balancing.
        """
        return pulumi.get(self, "connection_tracking_policy")

    @property
    @pulumi.getter(name="consistentHash")
    def consistent_hash(self) -> 'outputs.ConsistentHashLoadBalancerSettingsResponse':
        """
        Consistent Hash-based load balancing can be used to provide soft session affinity based on HTTP headers, cookies or other properties. This load balancing policy is applicable only for HTTP connections. The affinity to a particular destination host will be lost when one or more hosts are added/removed from the destination service. This field specifies parameters that control consistent hashing. This field is only applicable when localityLbPolicy is set to MAGLEV or RING_HASH. This field is applicable to either: - A regional backend service with the service_protocol set to HTTP, HTTPS, or HTTP2, and load_balancing_scheme set to INTERNAL_MANAGED. - A global backend service with the load_balancing_scheme set to INTERNAL_SELF_MANAGED. 
        """
        return pulumi.get(self, "consistent_hash")

    @property
    @pulumi.getter(name="creationTimestamp")
    def creation_timestamp(self) -> str:
        """
        Creation timestamp in RFC3339 text format.
        """
        return pulumi.get(self, "creation_timestamp")

    @property
    @pulumi.getter(name="customRequestHeaders")
    def custom_request_headers(self) -> Sequence[str]:
        """
        Headers that the load balancer adds to proxied requests. See [Creating custom headers](https://cloud.google.com/load-balancing/docs/custom-headers).
        """
        return pulumi.get(self, "custom_request_headers")

    @property
    @pulumi.getter(name="customResponseHeaders")
    def custom_response_headers(self) -> Sequence[str]:
        """
        Headers that the load balancer adds to proxied responses. See [Creating custom headers](https://cloud.google.com/load-balancing/docs/custom-headers).
        """
        return pulumi.get(self, "custom_response_headers")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        An optional description of this resource. Provide this property when you create the resource.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="edgeSecurityPolicy")
    def edge_security_policy(self) -> str:
        """
        The resource URL for the edge security policy associated with this backend service.
        """
        return pulumi.get(self, "edge_security_policy")

    @property
    @pulumi.getter(name="enableCDN")
    def enable_cdn(self) -> bool:
        """
        If true, enables Cloud CDN for the backend service of an external HTTP(S) load balancer.
        """
        return pulumi.get(self, "enable_cdn")

    @property
    @pulumi.getter(name="failoverPolicy")
    def failover_policy(self) -> 'outputs.BackendServiceFailoverPolicyResponse':
        """
        Requires at least one backend instance group to be defined as a backup (failover) backend. For load balancers that have configurable failover: [Internal TCP/UDP Load Balancing](https://cloud.google.com/load-balancing/docs/internal/failover-overview) and [external TCP/UDP Load Balancing](https://cloud.google.com/load-balancing/docs/network/networklb-failover-overview).
        """
        return pulumi.get(self, "failover_policy")

    @property
    @pulumi.getter
    def fingerprint(self) -> str:
        """
        Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking. This field will be ignored when inserting a BackendService. An up-to-date fingerprint must be provided in order to update the BackendService, otherwise the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make a get() request to retrieve a BackendService.
        """
        return pulumi.get(self, "fingerprint")

    @property
    @pulumi.getter(name="healthChecks")
    def health_checks(self) -> Sequence[str]:
        """
        The list of URLs to the healthChecks, httpHealthChecks (legacy), or httpsHealthChecks (legacy) resource for health checking this backend service. Not all backend services support legacy health checks. See Load balancer guide. Currently, at most one health check can be specified for each backend service. Backend services with instance group or zonal NEG backends must have a health check. Backend services with internet or serverless NEG backends must not have a health check.
        """
        return pulumi.get(self, "health_checks")

    @property
    @pulumi.getter
    def iap(self) -> 'outputs.BackendServiceIAPResponse':
        """
        The configurations for Identity-Aware Proxy on this resource. Not available for Internal TCP/UDP Load Balancing and Network Load Balancing.
        """
        return pulumi.get(self, "iap")

    @property
    @pulumi.getter
    def kind(self) -> str:
        """
        Type of resource. Always compute#backendService for backend services.
        """
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter(name="loadBalancingScheme")
    def load_balancing_scheme(self) -> str:
        """
        Specifies the load balancer type. A backend service created for one type of load balancer cannot be used with another. For more information, refer to Choosing a load balancer.
        """
        return pulumi.get(self, "load_balancing_scheme")

    @property
    @pulumi.getter(name="localityLbPolicy")
    def locality_lb_policy(self) -> str:
        """
        The load balancing algorithm used within the scope of the locality. The possible values are: - ROUND_ROBIN: This is a simple policy in which each healthy backend is selected in round robin order. This is the default. - LEAST_REQUEST: An O(1) algorithm which selects two random healthy hosts and picks the host which has fewer active requests. - RING_HASH: The ring/modulo hash load balancer implements consistent hashing to backends. The algorithm has the property that the addition/removal of a host from a set of N hosts only affects 1/N of the requests. - RANDOM: The load balancer selects a random healthy host. - ORIGINAL_DESTINATION: Backend host is selected based on the client connection metadata, i.e., connections are opened to the same address as the destination address of the incoming connection before the connection was redirected to the load balancer. - MAGLEV: used as a drop in replacement for the ring hash load balancer. Maglev is not as stable as ring hash but has faster table lookup build times and host selection times. For more information about Maglev, see https://ai.google/research/pubs/pub44824 This field is applicable to either: - A regional backend service with the service_protocol set to HTTP, HTTPS, or HTTP2, and load_balancing_scheme set to INTERNAL_MANAGED. - A global backend service with the load_balancing_scheme set to INTERNAL_SELF_MANAGED. If sessionAffinity is not NONE, and this field is not set to MAGLEV or RING_HASH, session affinity settings will not take effect. Only ROUND_ROBIN and RING_HASH are supported when the backend service is referenced by a URL map that is bound to target gRPC proxy that has validateForProxyless field set to true.
        """
        return pulumi.get(self, "locality_lb_policy")

    @property
    @pulumi.getter(name="logConfig")
    def log_config(self) -> 'outputs.BackendServiceLogConfigResponse':
        """
        This field denotes the logging options for the load balancer traffic served by this backend service. If logging is enabled, logs will be exported to Stackdriver.
        """
        return pulumi.get(self, "log_config")

    @property
    @pulumi.getter(name="maxStreamDuration")
    def max_stream_duration(self) -> 'outputs.DurationResponse':
        """
        Specifies the default maximum duration (timeout) for streams to this service. Duration is computed from the beginning of the stream until the response has been completely processed, including all retries. A stream that does not complete in this duration is closed. If not specified, there will be no timeout limit, i.e. the maximum duration is infinite. This value can be overridden in the PathMatcher configuration of the UrlMap that references this backend service. This field is only allowed when the loadBalancingScheme of the backend service is INTERNAL_SELF_MANAGED.
        """
        return pulumi.get(self, "max_stream_duration")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def network(self) -> str:
        """
        The URL of the network to which this backend service belongs. This field can only be specified when the load balancing scheme is set to INTERNAL.
        """
        return pulumi.get(self, "network")

    @property
    @pulumi.getter(name="outlierDetection")
    def outlier_detection(self) -> 'outputs.OutlierDetectionResponse':
        """
        Settings controlling the eviction of unhealthy hosts from the load balancing pool for the backend service. If not set, this feature is considered disabled. This field is applicable to either: - A regional backend service with the service_protocol set to HTTP, HTTPS, or HTTP2, and load_balancing_scheme set to INTERNAL_MANAGED. - A global backend service with the load_balancing_scheme set to INTERNAL_SELF_MANAGED. Not supported when the backend service is referenced by a URL map that is bound to target gRPC proxy that has validateForProxyless field set to true.
        """
        return pulumi.get(self, "outlier_detection")

    @property
    @pulumi.getter
    def port(self) -> int:
        """
        Deprecated in favor of portName. The TCP port to connect on the backend. The default value is 80. For Internal TCP/UDP Load Balancing and Network Load Balancing, omit port.
        """
        return pulumi.get(self, "port")

    @property
    @pulumi.getter(name="portName")
    def port_name(self) -> str:
        """
        A named port on a backend instance group representing the port for communication to the backend VMs in that group. The named port must be [defined on each backend instance group](https://cloud.google.com/load-balancing/docs/backend-service#named_ports). This parameter has no meaning if the backends are NEGs. For Internal TCP/UDP Load Balancing and Network Load Balancing, omit port_name.
        """
        return pulumi.get(self, "port_name")

    @property
    @pulumi.getter
    def protocol(self) -> str:
        """
        The protocol this BackendService uses to communicate with backends. Possible values are HTTP, HTTPS, HTTP2, TCP, SSL, UDP or GRPC. depending on the chosen load balancer or Traffic Director configuration. Refer to the documentation for the load balancers or for Traffic Director for more information. Must be set to GRPC when the backend service is referenced by a URL map that is bound to target gRPC proxy.
        """
        return pulumi.get(self, "protocol")

    @property
    @pulumi.getter
    def region(self) -> str:
        """
        URL of the region where the regional backend service resides. This field is not applicable to global backend services. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="securityPolicy")
    def security_policy(self) -> str:
        """
        The resource URL for the security policy associated with this backend service.
        """
        return pulumi.get(self, "security_policy")

    @property
    @pulumi.getter(name="securitySettings")
    def security_settings(self) -> 'outputs.SecuritySettingsResponse':
        """
        This field specifies the security settings that apply to this backend service. This field is applicable to a global backend service with the load_balancing_scheme set to INTERNAL_SELF_MANAGED.
        """
        return pulumi.get(self, "security_settings")

    @property
    @pulumi.getter(name="selfLink")
    def self_link(self) -> str:
        """
        Server-defined URL for the resource.
        """
        return pulumi.get(self, "self_link")

    @property
    @pulumi.getter(name="serviceBindings")
    def service_bindings(self) -> Sequence[str]:
        """
        URLs of networkservices.ServiceBinding resources. Can only be set if load balancing scheme is INTERNAL_SELF_MANAGED. If set, lists of backends and health checks must be both empty.
        """
        return pulumi.get(self, "service_bindings")

    @property
    @pulumi.getter(name="sessionAffinity")
    def session_affinity(self) -> str:
        """
        Type of session affinity to use. The default is NONE. Only NONE and HEADER_FIELD are supported when the backend service is referenced by a URL map that is bound to target gRPC proxy that has validateForProxyless field set to true. For more details, see: [Session Affinity](https://cloud.google.com/load-balancing/docs/backend-service#session_affinity).
        """
        return pulumi.get(self, "session_affinity")

    @property
    @pulumi.getter
    def subsetting(self) -> 'outputs.SubsettingResponse':
        return pulumi.get(self, "subsetting")

    @property
    @pulumi.getter(name="timeoutSec")
    def timeout_sec(self) -> int:
        """
        The backend service timeout has a different meaning depending on the type of load balancer. For more information see, Backend service settings The default is 30 seconds. The full range of timeout values allowed is 1 - 2,147,483,647 seconds. This value can be overridden in the PathMatcher configuration of the UrlMap that references this backend service. Not supported when the backend service is referenced by a URL map that is bound to target gRPC proxy that has validateForProxyless field set to true. Instead, use maxStreamDuration.
        """
        return pulumi.get(self, "timeout_sec")


class AwaitableGetRegionBackendServiceResult(GetRegionBackendServiceResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetRegionBackendServiceResult(
            affinity_cookie_ttl_sec=self.affinity_cookie_ttl_sec,
            backends=self.backends,
            cdn_policy=self.cdn_policy,
            circuit_breakers=self.circuit_breakers,
            compression_mode=self.compression_mode,
            connection_draining=self.connection_draining,
            connection_tracking_policy=self.connection_tracking_policy,
            consistent_hash=self.consistent_hash,
            creation_timestamp=self.creation_timestamp,
            custom_request_headers=self.custom_request_headers,
            custom_response_headers=self.custom_response_headers,
            description=self.description,
            edge_security_policy=self.edge_security_policy,
            enable_cdn=self.enable_cdn,
            failover_policy=self.failover_policy,
            fingerprint=self.fingerprint,
            health_checks=self.health_checks,
            iap=self.iap,
            kind=self.kind,
            load_balancing_scheme=self.load_balancing_scheme,
            locality_lb_policy=self.locality_lb_policy,
            log_config=self.log_config,
            max_stream_duration=self.max_stream_duration,
            name=self.name,
            network=self.network,
            outlier_detection=self.outlier_detection,
            port=self.port,
            port_name=self.port_name,
            protocol=self.protocol,
            region=self.region,
            security_policy=self.security_policy,
            security_settings=self.security_settings,
            self_link=self.self_link,
            service_bindings=self.service_bindings,
            session_affinity=self.session_affinity,
            subsetting=self.subsetting,
            timeout_sec=self.timeout_sec)


def get_region_backend_service(backend_service: Optional[str] = None,
                               project: Optional[str] = None,
                               region: Optional[str] = None,
                               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetRegionBackendServiceResult:
    """
    Returns the specified regional BackendService resource.
    """
    __args__ = dict()
    __args__['backendService'] = backend_service
    __args__['project'] = project
    __args__['region'] = region
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('google-native:compute/beta:getRegionBackendService', __args__, opts=opts, typ=GetRegionBackendServiceResult).value

    return AwaitableGetRegionBackendServiceResult(
        affinity_cookie_ttl_sec=__ret__.affinity_cookie_ttl_sec,
        backends=__ret__.backends,
        cdn_policy=__ret__.cdn_policy,
        circuit_breakers=__ret__.circuit_breakers,
        compression_mode=__ret__.compression_mode,
        connection_draining=__ret__.connection_draining,
        connection_tracking_policy=__ret__.connection_tracking_policy,
        consistent_hash=__ret__.consistent_hash,
        creation_timestamp=__ret__.creation_timestamp,
        custom_request_headers=__ret__.custom_request_headers,
        custom_response_headers=__ret__.custom_response_headers,
        description=__ret__.description,
        edge_security_policy=__ret__.edge_security_policy,
        enable_cdn=__ret__.enable_cdn,
        failover_policy=__ret__.failover_policy,
        fingerprint=__ret__.fingerprint,
        health_checks=__ret__.health_checks,
        iap=__ret__.iap,
        kind=__ret__.kind,
        load_balancing_scheme=__ret__.load_balancing_scheme,
        locality_lb_policy=__ret__.locality_lb_policy,
        log_config=__ret__.log_config,
        max_stream_duration=__ret__.max_stream_duration,
        name=__ret__.name,
        network=__ret__.network,
        outlier_detection=__ret__.outlier_detection,
        port=__ret__.port,
        port_name=__ret__.port_name,
        protocol=__ret__.protocol,
        region=__ret__.region,
        security_policy=__ret__.security_policy,
        security_settings=__ret__.security_settings,
        self_link=__ret__.self_link,
        service_bindings=__ret__.service_bindings,
        session_affinity=__ret__.session_affinity,
        subsetting=__ret__.subsetting,
        timeout_sec=__ret__.timeout_sec)


@_utilities.lift_output_func(get_region_backend_service)
def get_region_backend_service_output(backend_service: Optional[pulumi.Input[str]] = None,
                                      project: Optional[pulumi.Input[Optional[str]]] = None,
                                      region: Optional[pulumi.Input[str]] = None,
                                      opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetRegionBackendServiceResult]:
    """
    Returns the specified regional BackendService resource.
    """
    ...
