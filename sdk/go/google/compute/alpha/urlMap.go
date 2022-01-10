// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package alpha

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a UrlMap resource in the specified project using the data included in the request.
type UrlMap struct {
	pulumi.CustomResourceState

	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringOutput `pulumi:"creationTimestamp"`
	// defaultRouteAction takes effect when none of the hostRules match. The load balancer performs advanced routing actions, such as URL rewrites and header transformations, before forwarding the request to the selected backend. If defaultRouteAction specifies any weightedBackendServices, defaultService must not be set. Conversely if defaultService is set, defaultRouteAction cannot contain any weightedBackendServices. Only one of defaultRouteAction or defaultUrlRedirect must be set. UrlMaps for external HTTP(S) load balancers support only the urlRewrite action within defaultRouteAction. defaultRouteAction has no effect when the URL map is bound to a target gRPC proxy that has the validateForProxyless field set to true.
	DefaultRouteAction HttpRouteActionResponseOutput `pulumi:"defaultRouteAction"`
	// The full or partial URL of the defaultService resource to which traffic is directed if none of the hostRules match. If defaultRouteAction is also specified, advanced routing actions, such as URL rewrites, take effect before sending the request to the backend. However, if defaultService is specified, defaultRouteAction cannot contain any weightedBackendServices. Conversely, if routeAction specifies any weightedBackendServices, service must not be specified. Only one of defaultService, defaultUrlRedirect , or defaultRouteAction.weightedBackendService must be set. defaultService has no effect when the URL map is bound to a target gRPC proxy that has the validateForProxyless field set to true.
	DefaultService pulumi.StringOutput `pulumi:"defaultService"`
	// When none of the specified hostRules match, the request is redirected to a URL specified by defaultUrlRedirect. If defaultUrlRedirect is specified, defaultService or defaultRouteAction must not be set. Not supported when the URL map is bound to a target gRPC proxy.
	DefaultUrlRedirect HttpRedirectActionResponseOutput `pulumi:"defaultUrlRedirect"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description pulumi.StringOutput `pulumi:"description"`
	// Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking. This field is ignored when inserting a UrlMap. An up-to-date fingerprint must be provided in order to update the UrlMap, otherwise the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make a get() request to retrieve a UrlMap.
	Fingerprint pulumi.StringOutput `pulumi:"fingerprint"`
	// Specifies changes to request and response headers that need to take effect for the selected backendService. The headerAction specified here take effect after headerAction specified under pathMatcher. headerAction is not supported for load balancers that have their loadBalancingScheme set to EXTERNAL. Not supported when the URL map is bound to a target gRPC proxy that has validateForProxyless field set to true.
	HeaderAction HttpHeaderActionResponseOutput `pulumi:"headerAction"`
	// The list of host rules to use against the URL.
	HostRules HostRuleResponseArrayOutput `pulumi:"hostRules"`
	// Type of the resource. Always compute#urlMaps for url maps.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name pulumi.StringOutput `pulumi:"name"`
	// The list of named PathMatchers to use against the URL.
	PathMatchers PathMatcherResponseArrayOutput `pulumi:"pathMatchers"`
	// URL of the region where the regional URL map resides. This field is not applicable to global URL maps. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
	Region pulumi.StringOutput `pulumi:"region"`
	// Server-defined URL for the resource.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
	// The list of expected URL mapping tests. Request to update the UrlMap succeeds only if all test cases pass. You can specify a maximum of 100 tests per UrlMap. Not supported when the URL map is bound to a target gRPC proxy that has validateForProxyless field set to true.
	Tests UrlMapTestResponseArrayOutput `pulumi:"tests"`
}

// NewUrlMap registers a new resource with the given unique name, arguments, and options.
func NewUrlMap(ctx *pulumi.Context,
	name string, args *UrlMapArgs, opts ...pulumi.ResourceOption) (*UrlMap, error) {
	if args == nil {
		args = &UrlMapArgs{}
	}

	var resource UrlMap
	err := ctx.RegisterResource("google-native:compute/alpha:UrlMap", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUrlMap gets an existing UrlMap resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUrlMap(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UrlMapState, opts ...pulumi.ResourceOption) (*UrlMap, error) {
	var resource UrlMap
	err := ctx.ReadResource("google-native:compute/alpha:UrlMap", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering UrlMap resources.
type urlMapState struct {
}

type UrlMapState struct {
}

func (UrlMapState) ElementType() reflect.Type {
	return reflect.TypeOf((*urlMapState)(nil)).Elem()
}

type urlMapArgs struct {
	// defaultRouteAction takes effect when none of the hostRules match. The load balancer performs advanced routing actions, such as URL rewrites and header transformations, before forwarding the request to the selected backend. If defaultRouteAction specifies any weightedBackendServices, defaultService must not be set. Conversely if defaultService is set, defaultRouteAction cannot contain any weightedBackendServices. Only one of defaultRouteAction or defaultUrlRedirect must be set. UrlMaps for external HTTP(S) load balancers support only the urlRewrite action within defaultRouteAction. defaultRouteAction has no effect when the URL map is bound to a target gRPC proxy that has the validateForProxyless field set to true.
	DefaultRouteAction *HttpRouteAction `pulumi:"defaultRouteAction"`
	// The full or partial URL of the defaultService resource to which traffic is directed if none of the hostRules match. If defaultRouteAction is also specified, advanced routing actions, such as URL rewrites, take effect before sending the request to the backend. However, if defaultService is specified, defaultRouteAction cannot contain any weightedBackendServices. Conversely, if routeAction specifies any weightedBackendServices, service must not be specified. Only one of defaultService, defaultUrlRedirect , or defaultRouteAction.weightedBackendService must be set. defaultService has no effect when the URL map is bound to a target gRPC proxy that has the validateForProxyless field set to true.
	DefaultService *string `pulumi:"defaultService"`
	// When none of the specified hostRules match, the request is redirected to a URL specified by defaultUrlRedirect. If defaultUrlRedirect is specified, defaultService or defaultRouteAction must not be set. Not supported when the URL map is bound to a target gRPC proxy.
	DefaultUrlRedirect *HttpRedirectAction `pulumi:"defaultUrlRedirect"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description *string `pulumi:"description"`
	// Specifies changes to request and response headers that need to take effect for the selected backendService. The headerAction specified here take effect after headerAction specified under pathMatcher. headerAction is not supported for load balancers that have their loadBalancingScheme set to EXTERNAL. Not supported when the URL map is bound to a target gRPC proxy that has validateForProxyless field set to true.
	HeaderAction *HttpHeaderAction `pulumi:"headerAction"`
	// The list of host rules to use against the URL.
	HostRules []HostRule `pulumi:"hostRules"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name *string `pulumi:"name"`
	// The list of named PathMatchers to use against the URL.
	PathMatchers []PathMatcher `pulumi:"pathMatchers"`
	Project      *string       `pulumi:"project"`
	RequestId    *string       `pulumi:"requestId"`
	// The list of expected URL mapping tests. Request to update the UrlMap succeeds only if all test cases pass. You can specify a maximum of 100 tests per UrlMap. Not supported when the URL map is bound to a target gRPC proxy that has validateForProxyless field set to true.
	Tests []UrlMapTest `pulumi:"tests"`
}

// The set of arguments for constructing a UrlMap resource.
type UrlMapArgs struct {
	// defaultRouteAction takes effect when none of the hostRules match. The load balancer performs advanced routing actions, such as URL rewrites and header transformations, before forwarding the request to the selected backend. If defaultRouteAction specifies any weightedBackendServices, defaultService must not be set. Conversely if defaultService is set, defaultRouteAction cannot contain any weightedBackendServices. Only one of defaultRouteAction or defaultUrlRedirect must be set. UrlMaps for external HTTP(S) load balancers support only the urlRewrite action within defaultRouteAction. defaultRouteAction has no effect when the URL map is bound to a target gRPC proxy that has the validateForProxyless field set to true.
	DefaultRouteAction HttpRouteActionPtrInput
	// The full or partial URL of the defaultService resource to which traffic is directed if none of the hostRules match. If defaultRouteAction is also specified, advanced routing actions, such as URL rewrites, take effect before sending the request to the backend. However, if defaultService is specified, defaultRouteAction cannot contain any weightedBackendServices. Conversely, if routeAction specifies any weightedBackendServices, service must not be specified. Only one of defaultService, defaultUrlRedirect , or defaultRouteAction.weightedBackendService must be set. defaultService has no effect when the URL map is bound to a target gRPC proxy that has the validateForProxyless field set to true.
	DefaultService pulumi.StringPtrInput
	// When none of the specified hostRules match, the request is redirected to a URL specified by defaultUrlRedirect. If defaultUrlRedirect is specified, defaultService or defaultRouteAction must not be set. Not supported when the URL map is bound to a target gRPC proxy.
	DefaultUrlRedirect HttpRedirectActionPtrInput
	// An optional description of this resource. Provide this property when you create the resource.
	Description pulumi.StringPtrInput
	// Specifies changes to request and response headers that need to take effect for the selected backendService. The headerAction specified here take effect after headerAction specified under pathMatcher. headerAction is not supported for load balancers that have their loadBalancingScheme set to EXTERNAL. Not supported when the URL map is bound to a target gRPC proxy that has validateForProxyless field set to true.
	HeaderAction HttpHeaderActionPtrInput
	// The list of host rules to use against the URL.
	HostRules HostRuleArrayInput
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name pulumi.StringPtrInput
	// The list of named PathMatchers to use against the URL.
	PathMatchers PathMatcherArrayInput
	Project      pulumi.StringPtrInput
	RequestId    pulumi.StringPtrInput
	// The list of expected URL mapping tests. Request to update the UrlMap succeeds only if all test cases pass. You can specify a maximum of 100 tests per UrlMap. Not supported when the URL map is bound to a target gRPC proxy that has validateForProxyless field set to true.
	Tests UrlMapTestArrayInput
}

func (UrlMapArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*urlMapArgs)(nil)).Elem()
}

type UrlMapInput interface {
	pulumi.Input

	ToUrlMapOutput() UrlMapOutput
	ToUrlMapOutputWithContext(ctx context.Context) UrlMapOutput
}

func (*UrlMap) ElementType() reflect.Type {
	return reflect.TypeOf((**UrlMap)(nil)).Elem()
}

func (i *UrlMap) ToUrlMapOutput() UrlMapOutput {
	return i.ToUrlMapOutputWithContext(context.Background())
}

func (i *UrlMap) ToUrlMapOutputWithContext(ctx context.Context) UrlMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UrlMapOutput)
}

type UrlMapOutput struct{ *pulumi.OutputState }

func (UrlMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UrlMap)(nil)).Elem()
}

func (o UrlMapOutput) ToUrlMapOutput() UrlMapOutput {
	return o
}

func (o UrlMapOutput) ToUrlMapOutputWithContext(ctx context.Context) UrlMapOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*UrlMapInput)(nil)).Elem(), &UrlMap{})
	pulumi.RegisterOutputType(UrlMapOutput{})
}
