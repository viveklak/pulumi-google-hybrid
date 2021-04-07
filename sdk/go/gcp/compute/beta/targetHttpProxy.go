// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package beta

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Creates a TargetHttpProxy resource in the specified project using the data included in the request.
type TargetHttpProxy struct {
	pulumi.CustomResourceState

	// [Output Only] Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringOutput `pulumi:"creationTimestamp"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description pulumi.StringOutput `pulumi:"description"`
	// Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking. This field will be ignored when inserting a TargetHttpProxy. An up-to-date fingerprint must be provided in order to patch/update the TargetHttpProxy; otherwise, the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make a get() request to retrieve the TargetHttpProxy.
	Fingerprint pulumi.StringOutput `pulumi:"fingerprint"`
	// URLs to networkservices.HttpFilter resources enabled for xDS clients using this configuration. For example, https://networkservices.googleapis.com/v1alpha1/projects/project/locations/locationhttpFilters/httpFilter Only filters that handle outbound connection and stream events may be specified. These filters work in conjunction with a default set of HTTP filters that may already be configured by Traffic Director. Traffic Director will determine the final location of these filters within xDS configuration based on the name of the HTTP filter. If Traffic Director positions multiple filters at the same location, those filters will be in the same order as specified in this list.
	// httpFilters only applies for loadbalancers with loadBalancingScheme set to INTERNAL_SELF_MANAGED. See ForwardingRule for more details.
	HttpFilters pulumi.StringArrayOutput `pulumi:"httpFilters"`
	// [Output Only] Type of resource. Always compute#targetHttpProxy for target HTTP proxies.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name pulumi.StringOutput `pulumi:"name"`
	// This field only applies when the forwarding rule that references this target proxy has a loadBalancingScheme set to INTERNAL_SELF_MANAGED.
	//
	// When this field is set to true, Envoy proxies set up inbound traffic interception and bind to the IP address and port specified in the forwarding rule. This is generally useful when using Traffic Director to configure Envoy as a gateway or middle proxy (in other words, not a sidecar proxy). The Envoy proxy listens for inbound requests and handles requests when it receives them.
	//
	// The default is false.
	ProxyBind pulumi.BoolOutput `pulumi:"proxyBind"`
	// [Output Only] URL of the region where the regional Target HTTP Proxy resides. This field is not applicable to global Target HTTP Proxies.
	Region pulumi.StringOutput `pulumi:"region"`
	// [Output Only] Server-defined URL for the resource.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
	// URL to the UrlMap resource that defines the mapping from URL to the BackendService.
	UrlMap pulumi.StringOutput `pulumi:"urlMap"`
}

// NewTargetHttpProxy registers a new resource with the given unique name, arguments, and options.
func NewTargetHttpProxy(ctx *pulumi.Context,
	name string, args *TargetHttpProxyArgs, opts ...pulumi.ResourceOption) (*TargetHttpProxy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	if args.TargetHttpProxy == nil {
		return nil, errors.New("invalid value for required argument 'TargetHttpProxy'")
	}
	var resource TargetHttpProxy
	err := ctx.RegisterResource("gcp-native:compute/beta:TargetHttpProxy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTargetHttpProxy gets an existing TargetHttpProxy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTargetHttpProxy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TargetHttpProxyState, opts ...pulumi.ResourceOption) (*TargetHttpProxy, error) {
	var resource TargetHttpProxy
	err := ctx.ReadResource("gcp-native:compute/beta:TargetHttpProxy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TargetHttpProxy resources.
type targetHttpProxyState struct {
	// [Output Only] Creation timestamp in RFC3339 text format.
	CreationTimestamp *string `pulumi:"creationTimestamp"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description *string `pulumi:"description"`
	// Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking. This field will be ignored when inserting a TargetHttpProxy. An up-to-date fingerprint must be provided in order to patch/update the TargetHttpProxy; otherwise, the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make a get() request to retrieve the TargetHttpProxy.
	Fingerprint *string `pulumi:"fingerprint"`
	// URLs to networkservices.HttpFilter resources enabled for xDS clients using this configuration. For example, https://networkservices.googleapis.com/v1alpha1/projects/project/locations/locationhttpFilters/httpFilter Only filters that handle outbound connection and stream events may be specified. These filters work in conjunction with a default set of HTTP filters that may already be configured by Traffic Director. Traffic Director will determine the final location of these filters within xDS configuration based on the name of the HTTP filter. If Traffic Director positions multiple filters at the same location, those filters will be in the same order as specified in this list.
	// httpFilters only applies for loadbalancers with loadBalancingScheme set to INTERNAL_SELF_MANAGED. See ForwardingRule for more details.
	HttpFilters []string `pulumi:"httpFilters"`
	// [Output Only] Type of resource. Always compute#targetHttpProxy for target HTTP proxies.
	Kind *string `pulumi:"kind"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name *string `pulumi:"name"`
	// This field only applies when the forwarding rule that references this target proxy has a loadBalancingScheme set to INTERNAL_SELF_MANAGED.
	//
	// When this field is set to true, Envoy proxies set up inbound traffic interception and bind to the IP address and port specified in the forwarding rule. This is generally useful when using Traffic Director to configure Envoy as a gateway or middle proxy (in other words, not a sidecar proxy). The Envoy proxy listens for inbound requests and handles requests when it receives them.
	//
	// The default is false.
	ProxyBind *bool `pulumi:"proxyBind"`
	// [Output Only] URL of the region where the regional Target HTTP Proxy resides. This field is not applicable to global Target HTTP Proxies.
	Region *string `pulumi:"region"`
	// [Output Only] Server-defined URL for the resource.
	SelfLink *string `pulumi:"selfLink"`
	// URL to the UrlMap resource that defines the mapping from URL to the BackendService.
	UrlMap *string `pulumi:"urlMap"`
}

type TargetHttpProxyState struct {
	// [Output Only] Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringPtrInput
	// An optional description of this resource. Provide this property when you create the resource.
	Description pulumi.StringPtrInput
	// Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking. This field will be ignored when inserting a TargetHttpProxy. An up-to-date fingerprint must be provided in order to patch/update the TargetHttpProxy; otherwise, the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make a get() request to retrieve the TargetHttpProxy.
	Fingerprint pulumi.StringPtrInput
	// URLs to networkservices.HttpFilter resources enabled for xDS clients using this configuration. For example, https://networkservices.googleapis.com/v1alpha1/projects/project/locations/locationhttpFilters/httpFilter Only filters that handle outbound connection and stream events may be specified. These filters work in conjunction with a default set of HTTP filters that may already be configured by Traffic Director. Traffic Director will determine the final location of these filters within xDS configuration based on the name of the HTTP filter. If Traffic Director positions multiple filters at the same location, those filters will be in the same order as specified in this list.
	// httpFilters only applies for loadbalancers with loadBalancingScheme set to INTERNAL_SELF_MANAGED. See ForwardingRule for more details.
	HttpFilters pulumi.StringArrayInput
	// [Output Only] Type of resource. Always compute#targetHttpProxy for target HTTP proxies.
	Kind pulumi.StringPtrInput
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name pulumi.StringPtrInput
	// This field only applies when the forwarding rule that references this target proxy has a loadBalancingScheme set to INTERNAL_SELF_MANAGED.
	//
	// When this field is set to true, Envoy proxies set up inbound traffic interception and bind to the IP address and port specified in the forwarding rule. This is generally useful when using Traffic Director to configure Envoy as a gateway or middle proxy (in other words, not a sidecar proxy). The Envoy proxy listens for inbound requests and handles requests when it receives them.
	//
	// The default is false.
	ProxyBind pulumi.BoolPtrInput
	// [Output Only] URL of the region where the regional Target HTTP Proxy resides. This field is not applicable to global Target HTTP Proxies.
	Region pulumi.StringPtrInput
	// [Output Only] Server-defined URL for the resource.
	SelfLink pulumi.StringPtrInput
	// URL to the UrlMap resource that defines the mapping from URL to the BackendService.
	UrlMap pulumi.StringPtrInput
}

func (TargetHttpProxyState) ElementType() reflect.Type {
	return reflect.TypeOf((*targetHttpProxyState)(nil)).Elem()
}

type targetHttpProxyArgs struct {
	// [Output Only] Creation timestamp in RFC3339 text format.
	CreationTimestamp *string `pulumi:"creationTimestamp"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description *string `pulumi:"description"`
	// Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking. This field will be ignored when inserting a TargetHttpProxy. An up-to-date fingerprint must be provided in order to patch/update the TargetHttpProxy; otherwise, the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make a get() request to retrieve the TargetHttpProxy.
	Fingerprint *string `pulumi:"fingerprint"`
	// URLs to networkservices.HttpFilter resources enabled for xDS clients using this configuration. For example, https://networkservices.googleapis.com/v1alpha1/projects/project/locations/locationhttpFilters/httpFilter Only filters that handle outbound connection and stream events may be specified. These filters work in conjunction with a default set of HTTP filters that may already be configured by Traffic Director. Traffic Director will determine the final location of these filters within xDS configuration based on the name of the HTTP filter. If Traffic Director positions multiple filters at the same location, those filters will be in the same order as specified in this list.
	// httpFilters only applies for loadbalancers with loadBalancingScheme set to INTERNAL_SELF_MANAGED. See ForwardingRule for more details.
	HttpFilters []string `pulumi:"httpFilters"`
	// [Output Only] The unique identifier for the resource. This identifier is defined by the server.
	Id *string `pulumi:"id"`
	// [Output Only] Type of resource. Always compute#targetHttpProxy for target HTTP proxies.
	Kind *string `pulumi:"kind"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name    *string `pulumi:"name"`
	Project string  `pulumi:"project"`
	// This field only applies when the forwarding rule that references this target proxy has a loadBalancingScheme set to INTERNAL_SELF_MANAGED.
	//
	// When this field is set to true, Envoy proxies set up inbound traffic interception and bind to the IP address and port specified in the forwarding rule. This is generally useful when using Traffic Director to configure Envoy as a gateway or middle proxy (in other words, not a sidecar proxy). The Envoy proxy listens for inbound requests and handles requests when it receives them.
	//
	// The default is false.
	ProxyBind *bool `pulumi:"proxyBind"`
	// [Output Only] URL of the region where the regional Target HTTP Proxy resides. This field is not applicable to global Target HTTP Proxies.
	Region *string `pulumi:"region"`
	// [Output Only] Server-defined URL for the resource.
	SelfLink        *string `pulumi:"selfLink"`
	TargetHttpProxy string  `pulumi:"targetHttpProxy"`
	// URL to the UrlMap resource that defines the mapping from URL to the BackendService.
	UrlMap *string `pulumi:"urlMap"`
}

// The set of arguments for constructing a TargetHttpProxy resource.
type TargetHttpProxyArgs struct {
	// [Output Only] Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringPtrInput
	// An optional description of this resource. Provide this property when you create the resource.
	Description pulumi.StringPtrInput
	// Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking. This field will be ignored when inserting a TargetHttpProxy. An up-to-date fingerprint must be provided in order to patch/update the TargetHttpProxy; otherwise, the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make a get() request to retrieve the TargetHttpProxy.
	Fingerprint pulumi.StringPtrInput
	// URLs to networkservices.HttpFilter resources enabled for xDS clients using this configuration. For example, https://networkservices.googleapis.com/v1alpha1/projects/project/locations/locationhttpFilters/httpFilter Only filters that handle outbound connection and stream events may be specified. These filters work in conjunction with a default set of HTTP filters that may already be configured by Traffic Director. Traffic Director will determine the final location of these filters within xDS configuration based on the name of the HTTP filter. If Traffic Director positions multiple filters at the same location, those filters will be in the same order as specified in this list.
	// httpFilters only applies for loadbalancers with loadBalancingScheme set to INTERNAL_SELF_MANAGED. See ForwardingRule for more details.
	HttpFilters pulumi.StringArrayInput
	// [Output Only] The unique identifier for the resource. This identifier is defined by the server.
	Id pulumi.StringPtrInput
	// [Output Only] Type of resource. Always compute#targetHttpProxy for target HTTP proxies.
	Kind pulumi.StringPtrInput
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name    pulumi.StringPtrInput
	Project pulumi.StringInput
	// This field only applies when the forwarding rule that references this target proxy has a loadBalancingScheme set to INTERNAL_SELF_MANAGED.
	//
	// When this field is set to true, Envoy proxies set up inbound traffic interception and bind to the IP address and port specified in the forwarding rule. This is generally useful when using Traffic Director to configure Envoy as a gateway or middle proxy (in other words, not a sidecar proxy). The Envoy proxy listens for inbound requests and handles requests when it receives them.
	//
	// The default is false.
	ProxyBind pulumi.BoolPtrInput
	// [Output Only] URL of the region where the regional Target HTTP Proxy resides. This field is not applicable to global Target HTTP Proxies.
	Region pulumi.StringPtrInput
	// [Output Only] Server-defined URL for the resource.
	SelfLink        pulumi.StringPtrInput
	TargetHttpProxy pulumi.StringInput
	// URL to the UrlMap resource that defines the mapping from URL to the BackendService.
	UrlMap pulumi.StringPtrInput
}

func (TargetHttpProxyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*targetHttpProxyArgs)(nil)).Elem()
}

type TargetHttpProxyInput interface {
	pulumi.Input

	ToTargetHttpProxyOutput() TargetHttpProxyOutput
	ToTargetHttpProxyOutputWithContext(ctx context.Context) TargetHttpProxyOutput
}

func (*TargetHttpProxy) ElementType() reflect.Type {
	return reflect.TypeOf((*TargetHttpProxy)(nil))
}

func (i *TargetHttpProxy) ToTargetHttpProxyOutput() TargetHttpProxyOutput {
	return i.ToTargetHttpProxyOutputWithContext(context.Background())
}

func (i *TargetHttpProxy) ToTargetHttpProxyOutputWithContext(ctx context.Context) TargetHttpProxyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TargetHttpProxyOutput)
}

type TargetHttpProxyOutput struct {
	*pulumi.OutputState
}

func (TargetHttpProxyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TargetHttpProxy)(nil))
}

func (o TargetHttpProxyOutput) ToTargetHttpProxyOutput() TargetHttpProxyOutput {
	return o
}

func (o TargetHttpProxyOutput) ToTargetHttpProxyOutputWithContext(ctx context.Context) TargetHttpProxyOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(TargetHttpProxyOutput{})
}