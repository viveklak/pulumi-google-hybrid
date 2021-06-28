// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package beta

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Returns the specified TargetTcpProxy resource. Gets a list of available target TCP proxies by making a list() request.
func LookupTargetTcpProxy(ctx *pulumi.Context, args *LookupTargetTcpProxyArgs, opts ...pulumi.InvokeOption) (*LookupTargetTcpProxyResult, error) {
	var rv LookupTargetTcpProxyResult
	err := ctx.Invoke("google-native:compute/beta:getTargetTcpProxy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupTargetTcpProxyArgs struct {
	Project        string `pulumi:"project"`
	TargetTcpProxy string `pulumi:"targetTcpProxy"`
}

type LookupTargetTcpProxyResult struct {
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp string `pulumi:"creationTimestamp"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description string `pulumi:"description"`
	// Type of the resource. Always compute#targetTcpProxy for target TCP proxies.
	Kind string `pulumi:"kind"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name string `pulumi:"name"`
	// This field only applies when the forwarding rule that references this target proxy has a loadBalancingScheme set to INTERNAL_SELF_MANAGED.
	//
	// When this field is set to true, Envoy proxies set up inbound traffic interception and bind to the IP address and port specified in the forwarding rule. This is generally useful when using Traffic Director to configure Envoy as a gateway or middle proxy (in other words, not a sidecar proxy). The Envoy proxy listens for inbound requests and handles requests when it receives them.
	//
	// The default is false.
	ProxyBind bool `pulumi:"proxyBind"`
	// Specifies the type of proxy header to append before sending data to the backend, either NONE or PROXY_V1. The default is NONE.
	ProxyHeader string `pulumi:"proxyHeader"`
	// Server-defined URL for the resource.
	SelfLink string `pulumi:"selfLink"`
	// URL to the BackendService resource.
	Service string `pulumi:"service"`
}
