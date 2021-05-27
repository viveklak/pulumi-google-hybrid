// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates a regional BackendService resource in the specified project using the data included in the request. For more information, see  Backend services overview.
 */
export class RegionBackendService extends pulumi.CustomResource {
    /**
     * Get an existing RegionBackendService resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): RegionBackendService {
        return new RegionBackendService(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:compute/beta:RegionBackendService';

    /**
     * Returns true if the given object is an instance of RegionBackendService.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RegionBackendService {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RegionBackendService.__pulumiType;
    }

    /**
     * Lifetime of cookies in seconds. Only applicable if the loadBalancingScheme is EXTERNAL, INTERNAL_SELF_MANAGED, or INTERNAL_MANAGED, the protocol is HTTP or HTTPS, and the sessionAffinity is GENERATED_COOKIE, or HTTP_COOKIE.
     *
     * If set to 0, the cookie is non-persistent and lasts only until the end of the browser session (or equivalent). The maximum allowed value is one day (86,400).
     *
     * Not supported when the backend service is referenced by a URL map that is bound to target gRPC proxy that has validateForProxyless field set to true.
     */
    public readonly affinityCookieTtlSec!: pulumi.Output<number>;
    /**
     * The list of backends that serve this BackendService.
     */
    public readonly backends!: pulumi.Output<outputs.compute.beta.BackendResponse[]>;
    /**
     * Cloud CDN configuration for this BackendService. Only available for  external HTTP(S) Load Balancing.
     */
    public readonly cdnPolicy!: pulumi.Output<outputs.compute.beta.BackendServiceCdnPolicyResponse>;
    /**
     * Settings controlling the volume of connections to a backend service. If not set, this feature is considered disabled.
     *
     * This field is applicable to either:  
     * - A regional backend service with the service_protocol set to HTTP, HTTPS, or HTTP2, and load_balancing_scheme set to INTERNAL_MANAGED. 
     * - A global backend service with the load_balancing_scheme set to INTERNAL_SELF_MANAGED.  
     *
     * Not supported when the backend service is referenced by a URL map that is bound to target gRPC proxy that has validateForProxyless field set to true.
     */
    public readonly circuitBreakers!: pulumi.Output<outputs.compute.beta.CircuitBreakersResponse>;
    public readonly connectionDraining!: pulumi.Output<outputs.compute.beta.ConnectionDrainingResponse>;
    public readonly connectionTrackingPolicy!: pulumi.Output<outputs.compute.beta.BackendServiceConnectionTrackingPolicyResponse>;
    /**
     * Consistent Hash-based load balancing can be used to provide soft session affinity based on HTTP headers, cookies or other properties. This load balancing policy is applicable only for HTTP connections. The affinity to a particular destination host will be lost when one or more hosts are added/removed from the destination service. This field specifies parameters that control consistent hashing. This field is only applicable when localityLbPolicy is set to MAGLEV or RING_HASH.
     *
     * This field is applicable to either:  
     * - A regional backend service with the service_protocol set to HTTP, HTTPS, or HTTP2, and load_balancing_scheme set to INTERNAL_MANAGED. 
     * - A global backend service with the load_balancing_scheme set to INTERNAL_SELF_MANAGED.  
     *
     * Not supported when the backend service is referenced by a URL map that is bound to target gRPC proxy that has validateForProxyless field set to true.
     */
    public readonly consistentHash!: pulumi.Output<outputs.compute.beta.ConsistentHashLoadBalancerSettingsResponse>;
    /**
     * [Output Only] Creation timestamp in RFC3339 text format.
     */
    public readonly creationTimestamp!: pulumi.Output<string>;
    /**
     * Headers that the HTTP/S load balancer should add to proxied requests.
     */
    public readonly customRequestHeaders!: pulumi.Output<string[]>;
    /**
     * Headers that the HTTP/S load balancer should add to proxied responses.
     */
    public readonly customResponseHeaders!: pulumi.Output<string[]>;
    /**
     * An optional description of this resource. Provide this property when you create the resource.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * If true, enables Cloud CDN for the backend service. Only applicable if the loadBalancingScheme is EXTERNAL and the protocol is HTTP or HTTPS.
     */
    public readonly enableCDN!: pulumi.Output<boolean>;
    /**
     * Applicable only to Failover for Internal TCP/UDP Load Balancing and Network Load Balancing. Requires at least one backend instance group to be defined as a backup (failover) backend.
     */
    public readonly failoverPolicy!: pulumi.Output<outputs.compute.beta.BackendServiceFailoverPolicyResponse>;
    /**
     * Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking. This field will be ignored when inserting a BackendService. An up-to-date fingerprint must be provided in order to update the BackendService, otherwise the request will fail with error 412 conditionNotMet.
     *
     * To see the latest fingerprint, make a get() request to retrieve a BackendService.
     */
    public readonly fingerprint!: pulumi.Output<string>;
    /**
     * The list of URLs to the healthChecks, httpHealthChecks (legacy), or httpsHealthChecks (legacy) resource for health checking this backend service. Not all backend services support legacy health checks. See  Load balancer guide. Currently, at most one health check can be specified for each backend service. Backend services with instance group or zonal NEG backends must have a health check. Backend services with internet or serverless NEG backends must not have a health check.
     */
    public readonly healthChecks!: pulumi.Output<string[]>;
    /**
     * The configurations for Identity-Aware Proxy on this resource. Not available for Internal TCP/UDP Load Balancing and Network Load Balancing.
     */
    public readonly iap!: pulumi.Output<outputs.compute.beta.BackendServiceIAPResponse>;
    /**
     * [Output Only] Type of resource. Always compute#backendService for backend services.
     */
    public readonly kind!: pulumi.Output<string>;
    /**
     * Specifies the load balancer type. Choose EXTERNAL for external HTTP(S), SSL Proxy, TCP Proxy and Network Load Balancing. Choose  INTERNAL for Internal TCP/UDP Load Balancing. Choose  INTERNAL_MANAGED for Internal HTTP(S) Load Balancing.  INTERNAL_SELF_MANAGED for Traffic Director. A backend service created for one type of load balancer cannot be used with another. For more information, refer to Choosing a load balancer.
     */
    public readonly loadBalancingScheme!: pulumi.Output<string>;
    /**
     * The load balancing algorithm used within the scope of the locality. The possible values are:  
     * - ROUND_ROBIN: This is a simple policy in which each healthy backend is selected in round robin order. This is the default. 
     * - LEAST_REQUEST: An O(1) algorithm which selects two random healthy hosts and picks the host which has fewer active requests. 
     * - RING_HASH: The ring/modulo hash load balancer implements consistent hashing to backends. The algorithm has the property that the addition/removal of a host from a set of N hosts only affects 1/N of the requests. 
     * - RANDOM: The load balancer selects a random healthy host. 
     * - ORIGINAL_DESTINATION: Backend host is selected based on the client connection metadata, i.e., connections are opened to the same address as the destination address of the incoming connection before the connection was redirected to the load balancer. 
     * - MAGLEV: used as a drop in replacement for the ring hash load balancer. Maglev is not as stable as ring hash but has faster table lookup build times and host selection times. For more information about Maglev, see https://ai.google/research/pubs/pub44824 
     *
     * This field is applicable to either:  
     * - A regional backend service with the service_protocol set to HTTP, HTTPS, or HTTP2, and load_balancing_scheme set to INTERNAL_MANAGED. 
     * - A global backend service with the load_balancing_scheme set to INTERNAL_SELF_MANAGED.  
     *
     * If sessionAffinity is not NONE, and this field is not set to MAGLEV or RING_HASH, session affinity settings will not take effect.
     *
     * Only the default ROUND_ROBIN policy is supported when the backend service is referenced by a URL map that is bound to target gRPC proxy that has validateForProxyless field set to true.
     */
    public readonly localityLbPolicy!: pulumi.Output<string>;
    /**
     * This field denotes the logging options for the load balancer traffic served by this backend service. If logging is enabled, logs will be exported to Stackdriver.
     */
    public readonly logConfig!: pulumi.Output<outputs.compute.beta.BackendServiceLogConfigResponse>;
    /**
     * Specifies the default maximum duration (timeout) for streams to this service. Duration is computed from the beginning of the stream until the response has been completely processed, including all retries. A stream that does not complete in this duration is closed.
     * If not specified, there will be no timeout limit, i.e. the maximum duration is infinite.
     * This field is only allowed when the loadBalancingScheme of the backend service is INTERNAL_SELF_MANAGED.
     */
    public readonly maxStreamDuration!: pulumi.Output<outputs.compute.beta.DurationResponse>;
    /**
     * Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The URL of the network to which this backend service belongs. This field can only be specified when the load balancing scheme is set to INTERNAL.
     */
    public readonly network!: pulumi.Output<string>;
    /**
     * Settings controlling the eviction of unhealthy hosts from the load balancing pool for the backend service. If not set, this feature is considered disabled.
     *
     * This field is applicable to either:  
     * - A regional backend service with the service_protocol set to HTTP, HTTPS, or HTTP2, and load_balancing_scheme set to INTERNAL_MANAGED. 
     * - A global backend service with the load_balancing_scheme set to INTERNAL_SELF_MANAGED.  
     *
     * Not supported when the backend service is referenced by a URL map that is bound to target gRPC proxy that has validateForProxyless field set to true.
     */
    public readonly outlierDetection!: pulumi.Output<outputs.compute.beta.OutlierDetectionResponse>;
    /**
     * Deprecated in favor of portName. The TCP port to connect on the backend. The default value is 80.
     *
     * Backend services for Internal TCP/UDP Load Balancing and Network Load Balancing require you omit port.
     */
    public readonly port!: pulumi.Output<number>;
    /**
     * A named port on a backend instance group representing the port for communication to the backend VMs in that group. Required when the loadBalancingScheme is EXTERNAL (except Network Load Balancing), INTERNAL_MANAGED, or  INTERNAL_SELF_MANAGED and the backends are instance groups. The named port must be defined on each backend instance group. This parameter has no meaning if the backends are NEGs.
     *
     *
     *
     * Backend services for Internal TCP/UDP Load Balancing and Network Load Balancing require you omit port_name.
     */
    public readonly portName!: pulumi.Output<string>;
    /**
     * The protocol this BackendService uses to communicate with backends.
     *
     * Possible values are HTTP, HTTPS, HTTP2, TCP, SSL, UDP or GRPC. depending on the chosen load balancer or Traffic Director configuration. Refer to the documentation for the load balancer or for Traffic Director for more information.
     *
     * Must be set to GRPC when the backend service is referenced by a URL map that is bound to target gRPC proxy.
     */
    public readonly protocol!: pulumi.Output<string>;
    /**
     * [Output Only] URL of the region where the regional backend service resides. This field is not applicable to global backend services. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * [Output Only] The resource URL for the security policy associated with this backend service.
     */
    public readonly securityPolicy!: pulumi.Output<string>;
    /**
     * This field specifies the security policy that applies to this backend service. This field is applicable to either:  
     * - A regional backend service with the service_protocol set to HTTP, HTTPS, or HTTP2, and load_balancing_scheme set to INTERNAL_MANAGED. 
     * - A global backend service with the load_balancing_scheme set to INTERNAL_SELF_MANAGED.
     */
    public readonly securitySettings!: pulumi.Output<outputs.compute.beta.SecuritySettingsResponse>;
    /**
     * [Output Only] Server-defined URL for the resource.
     */
    public readonly selfLink!: pulumi.Output<string>;
    /**
     * Type of session affinity to use. The default is NONE.
     *
     * When the loadBalancingScheme is EXTERNAL: * For Network Load Balancing, the possible values are NONE, CLIENT_IP, CLIENT_IP_PROTO, or  CLIENT_IP_PORT_PROTO. * For all other load balancers that use loadBalancingScheme=EXTERNAL, the possible values are NONE, CLIENT_IP, or GENERATED_COOKIE. * You can use GENERATED_COOKIE if the protocol is HTTP, HTTP2, or HTTPS.
     *
     * When the loadBalancingScheme is INTERNAL, possible values are NONE, CLIENT_IP, CLIENT_IP_PROTO, or CLIENT_IP_PORT_PROTO.
     *
     * When the loadBalancingScheme is INTERNAL_SELF_MANAGED, or INTERNAL_MANAGED, possible values are NONE, CLIENT_IP, GENERATED_COOKIE, HEADER_FIELD, or HTTP_COOKIE.
     *
     * Not supported when the backend service is referenced by a URL map that is bound to target gRPC proxy that has validateForProxyless field set to true.
     */
    public readonly sessionAffinity!: pulumi.Output<string>;
    public readonly subsetting!: pulumi.Output<outputs.compute.beta.SubsettingResponse>;
    /**
     * The backend service timeout has a different meaning depending on the type of load balancer. For more information see,  Backend service settings The default is 30 seconds. The full range of timeout values allowed is 1 - 2,147,483,647 seconds.
     */
    public readonly timeoutSec!: pulumi.Output<number>;

    /**
     * Create a RegionBackendService resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RegionBackendServiceArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.project === undefined) && !opts.urn) {
                throw new Error("Missing required property 'project'");
            }
            if ((!args || args.region === undefined) && !opts.urn) {
                throw new Error("Missing required property 'region'");
            }
            inputs["affinityCookieTtlSec"] = args ? args.affinityCookieTtlSec : undefined;
            inputs["backends"] = args ? args.backends : undefined;
            inputs["cdnPolicy"] = args ? args.cdnPolicy : undefined;
            inputs["circuitBreakers"] = args ? args.circuitBreakers : undefined;
            inputs["connectionDraining"] = args ? args.connectionDraining : undefined;
            inputs["connectionTrackingPolicy"] = args ? args.connectionTrackingPolicy : undefined;
            inputs["consistentHash"] = args ? args.consistentHash : undefined;
            inputs["creationTimestamp"] = args ? args.creationTimestamp : undefined;
            inputs["customRequestHeaders"] = args ? args.customRequestHeaders : undefined;
            inputs["customResponseHeaders"] = args ? args.customResponseHeaders : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["enableCDN"] = args ? args.enableCDN : undefined;
            inputs["failoverPolicy"] = args ? args.failoverPolicy : undefined;
            inputs["fingerprint"] = args ? args.fingerprint : undefined;
            inputs["healthChecks"] = args ? args.healthChecks : undefined;
            inputs["iap"] = args ? args.iap : undefined;
            inputs["id"] = args ? args.id : undefined;
            inputs["kind"] = args ? args.kind : undefined;
            inputs["loadBalancingScheme"] = args ? args.loadBalancingScheme : undefined;
            inputs["localityLbPolicy"] = args ? args.localityLbPolicy : undefined;
            inputs["logConfig"] = args ? args.logConfig : undefined;
            inputs["maxStreamDuration"] = args ? args.maxStreamDuration : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["network"] = args ? args.network : undefined;
            inputs["outlierDetection"] = args ? args.outlierDetection : undefined;
            inputs["port"] = args ? args.port : undefined;
            inputs["portName"] = args ? args.portName : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["protocol"] = args ? args.protocol : undefined;
            inputs["region"] = args ? args.region : undefined;
            inputs["requestId"] = args ? args.requestId : undefined;
            inputs["securityPolicy"] = args ? args.securityPolicy : undefined;
            inputs["securitySettings"] = args ? args.securitySettings : undefined;
            inputs["selfLink"] = args ? args.selfLink : undefined;
            inputs["sessionAffinity"] = args ? args.sessionAffinity : undefined;
            inputs["subsetting"] = args ? args.subsetting : undefined;
            inputs["timeoutSec"] = args ? args.timeoutSec : undefined;
        } else {
            inputs["affinityCookieTtlSec"] = undefined /*out*/;
            inputs["backends"] = undefined /*out*/;
            inputs["cdnPolicy"] = undefined /*out*/;
            inputs["circuitBreakers"] = undefined /*out*/;
            inputs["connectionDraining"] = undefined /*out*/;
            inputs["connectionTrackingPolicy"] = undefined /*out*/;
            inputs["consistentHash"] = undefined /*out*/;
            inputs["creationTimestamp"] = undefined /*out*/;
            inputs["customRequestHeaders"] = undefined /*out*/;
            inputs["customResponseHeaders"] = undefined /*out*/;
            inputs["description"] = undefined /*out*/;
            inputs["enableCDN"] = undefined /*out*/;
            inputs["failoverPolicy"] = undefined /*out*/;
            inputs["fingerprint"] = undefined /*out*/;
            inputs["healthChecks"] = undefined /*out*/;
            inputs["iap"] = undefined /*out*/;
            inputs["kind"] = undefined /*out*/;
            inputs["loadBalancingScheme"] = undefined /*out*/;
            inputs["localityLbPolicy"] = undefined /*out*/;
            inputs["logConfig"] = undefined /*out*/;
            inputs["maxStreamDuration"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["network"] = undefined /*out*/;
            inputs["outlierDetection"] = undefined /*out*/;
            inputs["port"] = undefined /*out*/;
            inputs["portName"] = undefined /*out*/;
            inputs["protocol"] = undefined /*out*/;
            inputs["region"] = undefined /*out*/;
            inputs["securityPolicy"] = undefined /*out*/;
            inputs["securitySettings"] = undefined /*out*/;
            inputs["selfLink"] = undefined /*out*/;
            inputs["sessionAffinity"] = undefined /*out*/;
            inputs["subsetting"] = undefined /*out*/;
            inputs["timeoutSec"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(RegionBackendService.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a RegionBackendService resource.
 */
export interface RegionBackendServiceArgs {
    /**
     * Lifetime of cookies in seconds. Only applicable if the loadBalancingScheme is EXTERNAL, INTERNAL_SELF_MANAGED, or INTERNAL_MANAGED, the protocol is HTTP or HTTPS, and the sessionAffinity is GENERATED_COOKIE, or HTTP_COOKIE.
     *
     * If set to 0, the cookie is non-persistent and lasts only until the end of the browser session (or equivalent). The maximum allowed value is one day (86,400).
     *
     * Not supported when the backend service is referenced by a URL map that is bound to target gRPC proxy that has validateForProxyless field set to true.
     */
    readonly affinityCookieTtlSec?: pulumi.Input<number>;
    /**
     * The list of backends that serve this BackendService.
     */
    readonly backends?: pulumi.Input<pulumi.Input<inputs.compute.beta.BackendArgs>[]>;
    /**
     * Cloud CDN configuration for this BackendService. Only available for  external HTTP(S) Load Balancing.
     */
    readonly cdnPolicy?: pulumi.Input<inputs.compute.beta.BackendServiceCdnPolicyArgs>;
    /**
     * Settings controlling the volume of connections to a backend service. If not set, this feature is considered disabled.
     *
     * This field is applicable to either:  
     * - A regional backend service with the service_protocol set to HTTP, HTTPS, or HTTP2, and load_balancing_scheme set to INTERNAL_MANAGED. 
     * - A global backend service with the load_balancing_scheme set to INTERNAL_SELF_MANAGED.  
     *
     * Not supported when the backend service is referenced by a URL map that is bound to target gRPC proxy that has validateForProxyless field set to true.
     */
    readonly circuitBreakers?: pulumi.Input<inputs.compute.beta.CircuitBreakersArgs>;
    readonly connectionDraining?: pulumi.Input<inputs.compute.beta.ConnectionDrainingArgs>;
    readonly connectionTrackingPolicy?: pulumi.Input<inputs.compute.beta.BackendServiceConnectionTrackingPolicyArgs>;
    /**
     * Consistent Hash-based load balancing can be used to provide soft session affinity based on HTTP headers, cookies or other properties. This load balancing policy is applicable only for HTTP connections. The affinity to a particular destination host will be lost when one or more hosts are added/removed from the destination service. This field specifies parameters that control consistent hashing. This field is only applicable when localityLbPolicy is set to MAGLEV or RING_HASH.
     *
     * This field is applicable to either:  
     * - A regional backend service with the service_protocol set to HTTP, HTTPS, or HTTP2, and load_balancing_scheme set to INTERNAL_MANAGED. 
     * - A global backend service with the load_balancing_scheme set to INTERNAL_SELF_MANAGED.  
     *
     * Not supported when the backend service is referenced by a URL map that is bound to target gRPC proxy that has validateForProxyless field set to true.
     */
    readonly consistentHash?: pulumi.Input<inputs.compute.beta.ConsistentHashLoadBalancerSettingsArgs>;
    /**
     * [Output Only] Creation timestamp in RFC3339 text format.
     */
    readonly creationTimestamp?: pulumi.Input<string>;
    /**
     * Headers that the HTTP/S load balancer should add to proxied requests.
     */
    readonly customRequestHeaders?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Headers that the HTTP/S load balancer should add to proxied responses.
     */
    readonly customResponseHeaders?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * An optional description of this resource. Provide this property when you create the resource.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * If true, enables Cloud CDN for the backend service. Only applicable if the loadBalancingScheme is EXTERNAL and the protocol is HTTP or HTTPS.
     */
    readonly enableCDN?: pulumi.Input<boolean>;
    /**
     * Applicable only to Failover for Internal TCP/UDP Load Balancing and Network Load Balancing. Requires at least one backend instance group to be defined as a backup (failover) backend.
     */
    readonly failoverPolicy?: pulumi.Input<inputs.compute.beta.BackendServiceFailoverPolicyArgs>;
    /**
     * Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking. This field will be ignored when inserting a BackendService. An up-to-date fingerprint must be provided in order to update the BackendService, otherwise the request will fail with error 412 conditionNotMet.
     *
     * To see the latest fingerprint, make a get() request to retrieve a BackendService.
     */
    readonly fingerprint?: pulumi.Input<string>;
    /**
     * The list of URLs to the healthChecks, httpHealthChecks (legacy), or httpsHealthChecks (legacy) resource for health checking this backend service. Not all backend services support legacy health checks. See  Load balancer guide. Currently, at most one health check can be specified for each backend service. Backend services with instance group or zonal NEG backends must have a health check. Backend services with internet or serverless NEG backends must not have a health check.
     */
    readonly healthChecks?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The configurations for Identity-Aware Proxy on this resource. Not available for Internal TCP/UDP Load Balancing and Network Load Balancing.
     */
    readonly iap?: pulumi.Input<inputs.compute.beta.BackendServiceIAPArgs>;
    /**
     * [Output Only] The unique identifier for the resource. This identifier is defined by the server.
     */
    readonly id?: pulumi.Input<string>;
    /**
     * [Output Only] Type of resource. Always compute#backendService for backend services.
     */
    readonly kind?: pulumi.Input<string>;
    /**
     * Specifies the load balancer type. Choose EXTERNAL for external HTTP(S), SSL Proxy, TCP Proxy and Network Load Balancing. Choose  INTERNAL for Internal TCP/UDP Load Balancing. Choose  INTERNAL_MANAGED for Internal HTTP(S) Load Balancing.  INTERNAL_SELF_MANAGED for Traffic Director. A backend service created for one type of load balancer cannot be used with another. For more information, refer to Choosing a load balancer.
     */
    readonly loadBalancingScheme?: pulumi.Input<string>;
    /**
     * The load balancing algorithm used within the scope of the locality. The possible values are:  
     * - ROUND_ROBIN: This is a simple policy in which each healthy backend is selected in round robin order. This is the default. 
     * - LEAST_REQUEST: An O(1) algorithm which selects two random healthy hosts and picks the host which has fewer active requests. 
     * - RING_HASH: The ring/modulo hash load balancer implements consistent hashing to backends. The algorithm has the property that the addition/removal of a host from a set of N hosts only affects 1/N of the requests. 
     * - RANDOM: The load balancer selects a random healthy host. 
     * - ORIGINAL_DESTINATION: Backend host is selected based on the client connection metadata, i.e., connections are opened to the same address as the destination address of the incoming connection before the connection was redirected to the load balancer. 
     * - MAGLEV: used as a drop in replacement for the ring hash load balancer. Maglev is not as stable as ring hash but has faster table lookup build times and host selection times. For more information about Maglev, see https://ai.google/research/pubs/pub44824 
     *
     * This field is applicable to either:  
     * - A regional backend service with the service_protocol set to HTTP, HTTPS, or HTTP2, and load_balancing_scheme set to INTERNAL_MANAGED. 
     * - A global backend service with the load_balancing_scheme set to INTERNAL_SELF_MANAGED.  
     *
     * If sessionAffinity is not NONE, and this field is not set to MAGLEV or RING_HASH, session affinity settings will not take effect.
     *
     * Only the default ROUND_ROBIN policy is supported when the backend service is referenced by a URL map that is bound to target gRPC proxy that has validateForProxyless field set to true.
     */
    readonly localityLbPolicy?: pulumi.Input<string>;
    /**
     * This field denotes the logging options for the load balancer traffic served by this backend service. If logging is enabled, logs will be exported to Stackdriver.
     */
    readonly logConfig?: pulumi.Input<inputs.compute.beta.BackendServiceLogConfigArgs>;
    /**
     * Specifies the default maximum duration (timeout) for streams to this service. Duration is computed from the beginning of the stream until the response has been completely processed, including all retries. A stream that does not complete in this duration is closed.
     * If not specified, there will be no timeout limit, i.e. the maximum duration is infinite.
     * This field is only allowed when the loadBalancingScheme of the backend service is INTERNAL_SELF_MANAGED.
     */
    readonly maxStreamDuration?: pulumi.Input<inputs.compute.beta.DurationArgs>;
    /**
     * Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The URL of the network to which this backend service belongs. This field can only be specified when the load balancing scheme is set to INTERNAL.
     */
    readonly network?: pulumi.Input<string>;
    /**
     * Settings controlling the eviction of unhealthy hosts from the load balancing pool for the backend service. If not set, this feature is considered disabled.
     *
     * This field is applicable to either:  
     * - A regional backend service with the service_protocol set to HTTP, HTTPS, or HTTP2, and load_balancing_scheme set to INTERNAL_MANAGED. 
     * - A global backend service with the load_balancing_scheme set to INTERNAL_SELF_MANAGED.  
     *
     * Not supported when the backend service is referenced by a URL map that is bound to target gRPC proxy that has validateForProxyless field set to true.
     */
    readonly outlierDetection?: pulumi.Input<inputs.compute.beta.OutlierDetectionArgs>;
    /**
     * Deprecated in favor of portName. The TCP port to connect on the backend. The default value is 80.
     *
     * Backend services for Internal TCP/UDP Load Balancing and Network Load Balancing require you omit port.
     */
    readonly port?: pulumi.Input<number>;
    /**
     * A named port on a backend instance group representing the port for communication to the backend VMs in that group. Required when the loadBalancingScheme is EXTERNAL (except Network Load Balancing), INTERNAL_MANAGED, or  INTERNAL_SELF_MANAGED and the backends are instance groups. The named port must be defined on each backend instance group. This parameter has no meaning if the backends are NEGs.
     *
     *
     *
     * Backend services for Internal TCP/UDP Load Balancing and Network Load Balancing require you omit port_name.
     */
    readonly portName?: pulumi.Input<string>;
    readonly project: pulumi.Input<string>;
    /**
     * The protocol this BackendService uses to communicate with backends.
     *
     * Possible values are HTTP, HTTPS, HTTP2, TCP, SSL, UDP or GRPC. depending on the chosen load balancer or Traffic Director configuration. Refer to the documentation for the load balancer or for Traffic Director for more information.
     *
     * Must be set to GRPC when the backend service is referenced by a URL map that is bound to target gRPC proxy.
     */
    readonly protocol?: pulumi.Input<string>;
    /**
     * [Output Only] URL of the region where the regional backend service resides. This field is not applicable to global backend services. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
     */
    readonly region: pulumi.Input<string>;
    readonly requestId?: pulumi.Input<string>;
    /**
     * [Output Only] The resource URL for the security policy associated with this backend service.
     */
    readonly securityPolicy?: pulumi.Input<string>;
    /**
     * This field specifies the security policy that applies to this backend service. This field is applicable to either:  
     * - A regional backend service with the service_protocol set to HTTP, HTTPS, or HTTP2, and load_balancing_scheme set to INTERNAL_MANAGED. 
     * - A global backend service with the load_balancing_scheme set to INTERNAL_SELF_MANAGED.
     */
    readonly securitySettings?: pulumi.Input<inputs.compute.beta.SecuritySettingsArgs>;
    /**
     * [Output Only] Server-defined URL for the resource.
     */
    readonly selfLink?: pulumi.Input<string>;
    /**
     * Type of session affinity to use. The default is NONE.
     *
     * When the loadBalancingScheme is EXTERNAL: * For Network Load Balancing, the possible values are NONE, CLIENT_IP, CLIENT_IP_PROTO, or  CLIENT_IP_PORT_PROTO. * For all other load balancers that use loadBalancingScheme=EXTERNAL, the possible values are NONE, CLIENT_IP, or GENERATED_COOKIE. * You can use GENERATED_COOKIE if the protocol is HTTP, HTTP2, or HTTPS.
     *
     * When the loadBalancingScheme is INTERNAL, possible values are NONE, CLIENT_IP, CLIENT_IP_PROTO, or CLIENT_IP_PORT_PROTO.
     *
     * When the loadBalancingScheme is INTERNAL_SELF_MANAGED, or INTERNAL_MANAGED, possible values are NONE, CLIENT_IP, GENERATED_COOKIE, HEADER_FIELD, or HTTP_COOKIE.
     *
     * Not supported when the backend service is referenced by a URL map that is bound to target gRPC proxy that has validateForProxyless field set to true.
     */
    readonly sessionAffinity?: pulumi.Input<string>;
    readonly subsetting?: pulumi.Input<inputs.compute.beta.SubsettingArgs>;
    /**
     * The backend service timeout has a different meaning depending on the type of load balancer. For more information see,  Backend service settings The default is 30 seconds. The full range of timeout values allowed is 1 - 2,147,483,647 seconds.
     */
    readonly timeoutSec?: pulumi.Input<number>;
}
