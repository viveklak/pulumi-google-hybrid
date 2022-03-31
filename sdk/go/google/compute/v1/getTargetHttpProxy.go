// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Returns the specified TargetHttpProxy resource. Gets a list of available target HTTP proxies by making a list() request.
func LookupTargetHttpProxy(ctx *pulumi.Context, args *LookupTargetHttpProxyArgs, opts ...pulumi.InvokeOption) (*LookupTargetHttpProxyResult, error) {
	var rv LookupTargetHttpProxyResult
	err := ctx.Invoke("google-native:compute/v1:getTargetHttpProxy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupTargetHttpProxyArgs struct {
	Project         *string `pulumi:"project"`
	TargetHttpProxy string  `pulumi:"targetHttpProxy"`
}

type LookupTargetHttpProxyResult struct {
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp string `pulumi:"creationTimestamp"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description string `pulumi:"description"`
	// Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking. This field will be ignored when inserting a TargetHttpProxy. An up-to-date fingerprint must be provided in order to patch/update the TargetHttpProxy; otherwise, the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make a get() request to retrieve the TargetHttpProxy.
	Fingerprint string `pulumi:"fingerprint"`
	// Type of resource. Always compute#targetHttpProxy for target HTTP proxies.
	Kind string `pulumi:"kind"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name string `pulumi:"name"`
	// This field only applies when the forwarding rule that references this target proxy has a loadBalancingScheme set to INTERNAL_SELF_MANAGED. When this field is set to true, Envoy proxies set up inbound traffic interception and bind to the IP address and port specified in the forwarding rule. This is generally useful when using Traffic Director to configure Envoy as a gateway or middle proxy (in other words, not a sidecar proxy). The Envoy proxy listens for inbound requests and handles requests when it receives them. The default is false.
	ProxyBind bool `pulumi:"proxyBind"`
	// URL of the region where the regional Target HTTP Proxy resides. This field is not applicable to global Target HTTP Proxies.
	Region string `pulumi:"region"`
	// Server-defined URL for the resource.
	SelfLink string `pulumi:"selfLink"`
	// URL to the UrlMap resource that defines the mapping from URL to the BackendService.
	UrlMap string `pulumi:"urlMap"`
}

func LookupTargetHttpProxyOutput(ctx *pulumi.Context, args LookupTargetHttpProxyOutputArgs, opts ...pulumi.InvokeOption) LookupTargetHttpProxyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupTargetHttpProxyResult, error) {
			args := v.(LookupTargetHttpProxyArgs)
			r, err := LookupTargetHttpProxy(ctx, &args, opts...)
			return *r, err
		}).(LookupTargetHttpProxyResultOutput)
}

type LookupTargetHttpProxyOutputArgs struct {
	Project         pulumi.StringPtrInput `pulumi:"project"`
	TargetHttpProxy pulumi.StringInput    `pulumi:"targetHttpProxy"`
}

func (LookupTargetHttpProxyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTargetHttpProxyArgs)(nil)).Elem()
}

type LookupTargetHttpProxyResultOutput struct{ *pulumi.OutputState }

func (LookupTargetHttpProxyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTargetHttpProxyResult)(nil)).Elem()
}

func (o LookupTargetHttpProxyResultOutput) ToLookupTargetHttpProxyResultOutput() LookupTargetHttpProxyResultOutput {
	return o
}

func (o LookupTargetHttpProxyResultOutput) ToLookupTargetHttpProxyResultOutputWithContext(ctx context.Context) LookupTargetHttpProxyResultOutput {
	return o
}

// Creation timestamp in RFC3339 text format.
func (o LookupTargetHttpProxyResultOutput) CreationTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTargetHttpProxyResult) string { return v.CreationTimestamp }).(pulumi.StringOutput)
}

// An optional description of this resource. Provide this property when you create the resource.
func (o LookupTargetHttpProxyResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTargetHttpProxyResult) string { return v.Description }).(pulumi.StringOutput)
}

// Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking. This field will be ignored when inserting a TargetHttpProxy. An up-to-date fingerprint must be provided in order to patch/update the TargetHttpProxy; otherwise, the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make a get() request to retrieve the TargetHttpProxy.
func (o LookupTargetHttpProxyResultOutput) Fingerprint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTargetHttpProxyResult) string { return v.Fingerprint }).(pulumi.StringOutput)
}

// Type of resource. Always compute#targetHttpProxy for target HTTP proxies.
func (o LookupTargetHttpProxyResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTargetHttpProxyResult) string { return v.Kind }).(pulumi.StringOutput)
}

// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
func (o LookupTargetHttpProxyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTargetHttpProxyResult) string { return v.Name }).(pulumi.StringOutput)
}

// This field only applies when the forwarding rule that references this target proxy has a loadBalancingScheme set to INTERNAL_SELF_MANAGED. When this field is set to true, Envoy proxies set up inbound traffic interception and bind to the IP address and port specified in the forwarding rule. This is generally useful when using Traffic Director to configure Envoy as a gateway or middle proxy (in other words, not a sidecar proxy). The Envoy proxy listens for inbound requests and handles requests when it receives them. The default is false.
func (o LookupTargetHttpProxyResultOutput) ProxyBind() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupTargetHttpProxyResult) bool { return v.ProxyBind }).(pulumi.BoolOutput)
}

// URL of the region where the regional Target HTTP Proxy resides. This field is not applicable to global Target HTTP Proxies.
func (o LookupTargetHttpProxyResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTargetHttpProxyResult) string { return v.Region }).(pulumi.StringOutput)
}

// Server-defined URL for the resource.
func (o LookupTargetHttpProxyResultOutput) SelfLink() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTargetHttpProxyResult) string { return v.SelfLink }).(pulumi.StringOutput)
}

// URL to the UrlMap resource that defines the mapping from URL to the BackendService.
func (o LookupTargetHttpProxyResultOutput) UrlMap() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTargetHttpProxyResult) string { return v.UrlMap }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupTargetHttpProxyResultOutput{})
}
