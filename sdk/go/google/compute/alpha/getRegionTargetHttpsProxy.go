// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package alpha

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Returns the specified TargetHttpsProxy resource in the specified region. Gets a list of available target HTTP proxies by making a list() request.
func LookupRegionTargetHttpsProxy(ctx *pulumi.Context, args *LookupRegionTargetHttpsProxyArgs, opts ...pulumi.InvokeOption) (*LookupRegionTargetHttpsProxyResult, error) {
	var rv LookupRegionTargetHttpsProxyResult
	err := ctx.Invoke("google-native:compute/alpha:getRegionTargetHttpsProxy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupRegionTargetHttpsProxyArgs struct {
	Project          string `pulumi:"project"`
	Region           string `pulumi:"region"`
	TargetHttpsProxy string `pulumi:"targetHttpsProxy"`
}

type LookupRegionTargetHttpsProxyResult struct {
	// Optional. A URL referring to a networksecurity.AuthorizationPolicy resource that describes how the proxy should authorize inbound traffic. If left blank, access will not be restricted by an authorization policy.
	// Refer to the AuthorizationPolicy resource for additional details.
	// authorizationPolicy only applies to a global TargetHttpsProxy attached to globalForwardingRules with the loadBalancingScheme set to INTERNAL_SELF_MANAGED.
	// Note: This field currently has no impact.
	AuthorizationPolicy string `pulumi:"authorizationPolicy"`
	// URL of a certificate map that identifies a certificate map associated with the given target proxy. This field can only be set for global target proxies. If set, sslCertificates will be ignored.
	CertificateMap string `pulumi:"certificateMap"`
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp string `pulumi:"creationTimestamp"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description string `pulumi:"description"`
	// Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking. This field will be ignored when inserting a TargetHttpsProxy. An up-to-date fingerprint must be provided in order to patch the TargetHttpsProxy; otherwise, the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make a get() request to retrieve the TargetHttpsProxy.
	Fingerprint string `pulumi:"fingerprint"`
	// URLs to networkservices.HttpFilter resources enabled for xDS clients using this configuration. For example, https://networkservices.googleapis.com/beta/projects/project/locations/locationhttpFilters/httpFilter Only filters that handle outbound connection and stream events may be specified. These filters work in conjunction with a default set of HTTP filters that may already be configured by Traffic Director. Traffic Director will determine the final location of these filters within xDS configuration based on the name of the HTTP filter. If Traffic Director positions multiple filters at the same location, those filters will be in the same order as specified in this list.
	// httpFilters only applies for loadbalancers with loadBalancingScheme set to INTERNAL_SELF_MANAGED. See ForwardingRule for more details.
	HttpFilters []string `pulumi:"httpFilters"`
	// Type of resource. Always compute#targetHttpsProxy for target HTTPS proxies.
	Kind string `pulumi:"kind"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name string `pulumi:"name"`
	// This field only applies when the forwarding rule that references this target proxy has a loadBalancingScheme set to INTERNAL_SELF_MANAGED.
	//
	// When this field is set to true, Envoy proxies set up inbound traffic interception and bind to the IP address and port specified in the forwarding rule. This is generally useful when using Traffic Director to configure Envoy as a gateway or middle proxy (in other words, not a sidecar proxy). The Envoy proxy listens for inbound requests and handles requests when it receives them.
	//
	// The default is false.
	ProxyBind bool `pulumi:"proxyBind"`
	// Specifies the QUIC override policy for this TargetHttpsProxy resource. This setting determines whether the load balancer attempts to negotiate QUIC with clients. You can specify NONE, ENABLE, or DISABLE.
	// - When quic-override is set to NONE, Google manages whether QUIC is used.
	// - When quic-override is set to ENABLE, the load balancer uses QUIC when possible.
	// - When quic-override is set to DISABLE, the load balancer doesn't use QUIC.
	// - If the quic-override flag is not specified, NONE is implied.
	QuicOverride string `pulumi:"quicOverride"`
	// URL of the region where the regional TargetHttpsProxy resides. This field is not applicable to global TargetHttpsProxies.
	Region string `pulumi:"region"`
	// Server-defined URL for the resource.
	SelfLink string `pulumi:"selfLink"`
	// Server-defined URL for this resource with the resource id.
	SelfLinkWithId string `pulumi:"selfLinkWithId"`
	// Optional. A URL referring to a networksecurity.ServerTlsPolicy resource that describes how the proxy should authenticate inbound traffic.
	// serverTlsPolicy only applies to a global TargetHttpsProxy attached to globalForwardingRules with the loadBalancingScheme set to INTERNAL_SELF_MANAGED.
	// If left blank, communications are not encrypted.
	// Note: This field currently has no impact.
	ServerTlsPolicy string `pulumi:"serverTlsPolicy"`
	// URLs to SslCertificate resources that are used to authenticate connections between users and the load balancer. At least one SSL certificate must be specified. Currently, you may specify up to 15 SSL certificates.
	SslCertificates []string `pulumi:"sslCertificates"`
	// URL of SslPolicy resource that will be associated with the TargetHttpsProxy resource. If not set, the TargetHttpsProxy resource has no SSL policy configured.
	SslPolicy string `pulumi:"sslPolicy"`
	// A fully-qualified or valid partial URL to the UrlMap resource that defines the mapping from URL to the BackendService. For example, the following are all valid URLs for specifying a URL map:
	// - https://www.googleapis.compute/v1/projects/project/global/urlMaps/url-map
	// - projects/project/global/urlMaps/url-map
	// - global/urlMaps/url-map
	UrlMap string `pulumi:"urlMap"`
}
