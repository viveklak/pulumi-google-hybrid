// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package beta

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a UrlMap resource in the specified project using the data included in the request.
type RegionUrlMap struct {
	pulumi.CustomResourceState

	// [Output Only] Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringOutput `pulumi:"creationTimestamp"`
	// defaultRouteAction takes effect when none of the  hostRules match. The load balancer performs advanced routing actions like URL rewrites, header transformations, etc. prior to forwarding the request to the selected backend. If defaultRouteAction specifies any weightedBackendServices, defaultService must not be set. Conversely if defaultService is set, defaultRouteAction cannot contain any  weightedBackendServices.
	// Only one of defaultRouteAction or defaultUrlRedirect must be set.
	// UrlMaps for external HTTP(S) load balancers support only the urlRewrite action within defaultRouteAction.
	// defaultRouteAction has no effect when the URL map is bound to target gRPC proxy that has validateForProxyless field set to true.
	DefaultRouteAction HttpRouteActionResponseOutput `pulumi:"defaultRouteAction"`
	// The full or partial URL of the defaultService resource to which traffic is directed if none of the hostRules match. If defaultRouteAction is additionally specified, advanced routing actions like URL Rewrites, etc. take effect prior to sending the request to the backend. However, if defaultService is specified, defaultRouteAction cannot contain any weightedBackendServices. Conversely, if routeAction specifies any weightedBackendServices, service must not be specified.
	// Only one of defaultService, defaultUrlRedirect  or defaultRouteAction.weightedBackendService must be set.
	// defaultService has no effect when the URL map is bound to target gRPC proxy that has validateForProxyless field set to true.
	DefaultService pulumi.StringOutput `pulumi:"defaultService"`
	// When none of the specified hostRules match, the request is redirected to a URL specified by defaultUrlRedirect.
	// If defaultUrlRedirect is specified, defaultService or defaultRouteAction must not be set.
	// Not supported when the URL map is bound to target gRPC proxy.
	DefaultUrlRedirect HttpRedirectActionResponseOutput `pulumi:"defaultUrlRedirect"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description pulumi.StringOutput `pulumi:"description"`
	// Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking. This field will be ignored when inserting a UrlMap. An up-to-date fingerprint must be provided in order to update the UrlMap, otherwise the request will fail with error 412 conditionNotMet.
	//
	// To see the latest fingerprint, make a get() request to retrieve a UrlMap.
	Fingerprint pulumi.StringOutput `pulumi:"fingerprint"`
	// Specifies changes to request and response headers that need to take effect for the selected backendService.
	// The headerAction specified here take effect after headerAction specified under pathMatcher.
	// Note that headerAction is not supported for Loadbalancers that have their loadBalancingScheme set to EXTERNAL.
	// Not supported when the URL map is bound to target gRPC proxy that has validateForProxyless field set to true.
	HeaderAction HttpHeaderActionResponseOutput `pulumi:"headerAction"`
	// The list of HostRules to use against the URL.
	HostRules HostRuleResponseArrayOutput `pulumi:"hostRules"`
	// [Output Only] Type of the resource. Always compute#urlMaps for url maps.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name pulumi.StringOutput `pulumi:"name"`
	// The list of named PathMatchers to use against the URL.
	PathMatchers PathMatcherResponseArrayOutput `pulumi:"pathMatchers"`
	// [Output Only] URL of the region where the regional URL map resides. This field is not applicable to global URL maps. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
	Region pulumi.StringOutput `pulumi:"region"`
	// [Output Only] Server-defined URL for the resource.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
	// The list of expected URL mapping tests. Request to update this UrlMap will succeed only if all of the test cases pass. You can specify a maximum of 100 tests per UrlMap.
	// Not supported when the URL map is bound to target gRPC proxy that has validateForProxyless field set to true.
	Tests UrlMapTestResponseArrayOutput `pulumi:"tests"`
}

// NewRegionUrlMap registers a new resource with the given unique name, arguments, and options.
func NewRegionUrlMap(ctx *pulumi.Context,
	name string, args *RegionUrlMapArgs, opts ...pulumi.ResourceOption) (*RegionUrlMap, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	if args.Region == nil {
		return nil, errors.New("invalid value for required argument 'Region'")
	}
	var resource RegionUrlMap
	err := ctx.RegisterResource("google-native:compute/beta:RegionUrlMap", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRegionUrlMap gets an existing RegionUrlMap resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRegionUrlMap(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RegionUrlMapState, opts ...pulumi.ResourceOption) (*RegionUrlMap, error) {
	var resource RegionUrlMap
	err := ctx.ReadResource("google-native:compute/beta:RegionUrlMap", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RegionUrlMap resources.
type regionUrlMapState struct {
	// [Output Only] Creation timestamp in RFC3339 text format.
	CreationTimestamp *string `pulumi:"creationTimestamp"`
	// defaultRouteAction takes effect when none of the  hostRules match. The load balancer performs advanced routing actions like URL rewrites, header transformations, etc. prior to forwarding the request to the selected backend. If defaultRouteAction specifies any weightedBackendServices, defaultService must not be set. Conversely if defaultService is set, defaultRouteAction cannot contain any  weightedBackendServices.
	// Only one of defaultRouteAction or defaultUrlRedirect must be set.
	// UrlMaps for external HTTP(S) load balancers support only the urlRewrite action within defaultRouteAction.
	// defaultRouteAction has no effect when the URL map is bound to target gRPC proxy that has validateForProxyless field set to true.
	DefaultRouteAction *HttpRouteActionResponse `pulumi:"defaultRouteAction"`
	// The full or partial URL of the defaultService resource to which traffic is directed if none of the hostRules match. If defaultRouteAction is additionally specified, advanced routing actions like URL Rewrites, etc. take effect prior to sending the request to the backend. However, if defaultService is specified, defaultRouteAction cannot contain any weightedBackendServices. Conversely, if routeAction specifies any weightedBackendServices, service must not be specified.
	// Only one of defaultService, defaultUrlRedirect  or defaultRouteAction.weightedBackendService must be set.
	// defaultService has no effect when the URL map is bound to target gRPC proxy that has validateForProxyless field set to true.
	DefaultService *string `pulumi:"defaultService"`
	// When none of the specified hostRules match, the request is redirected to a URL specified by defaultUrlRedirect.
	// If defaultUrlRedirect is specified, defaultService or defaultRouteAction must not be set.
	// Not supported when the URL map is bound to target gRPC proxy.
	DefaultUrlRedirect *HttpRedirectActionResponse `pulumi:"defaultUrlRedirect"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description *string `pulumi:"description"`
	// Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking. This field will be ignored when inserting a UrlMap. An up-to-date fingerprint must be provided in order to update the UrlMap, otherwise the request will fail with error 412 conditionNotMet.
	//
	// To see the latest fingerprint, make a get() request to retrieve a UrlMap.
	Fingerprint *string `pulumi:"fingerprint"`
	// Specifies changes to request and response headers that need to take effect for the selected backendService.
	// The headerAction specified here take effect after headerAction specified under pathMatcher.
	// Note that headerAction is not supported for Loadbalancers that have their loadBalancingScheme set to EXTERNAL.
	// Not supported when the URL map is bound to target gRPC proxy that has validateForProxyless field set to true.
	HeaderAction *HttpHeaderActionResponse `pulumi:"headerAction"`
	// The list of HostRules to use against the URL.
	HostRules []HostRuleResponse `pulumi:"hostRules"`
	// [Output Only] Type of the resource. Always compute#urlMaps for url maps.
	Kind *string `pulumi:"kind"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name *string `pulumi:"name"`
	// The list of named PathMatchers to use against the URL.
	PathMatchers []PathMatcherResponse `pulumi:"pathMatchers"`
	// [Output Only] URL of the region where the regional URL map resides. This field is not applicable to global URL maps. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
	Region *string `pulumi:"region"`
	// [Output Only] Server-defined URL for the resource.
	SelfLink *string `pulumi:"selfLink"`
	// The list of expected URL mapping tests. Request to update this UrlMap will succeed only if all of the test cases pass. You can specify a maximum of 100 tests per UrlMap.
	// Not supported when the URL map is bound to target gRPC proxy that has validateForProxyless field set to true.
	Tests []UrlMapTestResponse `pulumi:"tests"`
}

type RegionUrlMapState struct {
	// [Output Only] Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringPtrInput
	// defaultRouteAction takes effect when none of the  hostRules match. The load balancer performs advanced routing actions like URL rewrites, header transformations, etc. prior to forwarding the request to the selected backend. If defaultRouteAction specifies any weightedBackendServices, defaultService must not be set. Conversely if defaultService is set, defaultRouteAction cannot contain any  weightedBackendServices.
	// Only one of defaultRouteAction or defaultUrlRedirect must be set.
	// UrlMaps for external HTTP(S) load balancers support only the urlRewrite action within defaultRouteAction.
	// defaultRouteAction has no effect when the URL map is bound to target gRPC proxy that has validateForProxyless field set to true.
	DefaultRouteAction HttpRouteActionResponsePtrInput
	// The full or partial URL of the defaultService resource to which traffic is directed if none of the hostRules match. If defaultRouteAction is additionally specified, advanced routing actions like URL Rewrites, etc. take effect prior to sending the request to the backend. However, if defaultService is specified, defaultRouteAction cannot contain any weightedBackendServices. Conversely, if routeAction specifies any weightedBackendServices, service must not be specified.
	// Only one of defaultService, defaultUrlRedirect  or defaultRouteAction.weightedBackendService must be set.
	// defaultService has no effect when the URL map is bound to target gRPC proxy that has validateForProxyless field set to true.
	DefaultService pulumi.StringPtrInput
	// When none of the specified hostRules match, the request is redirected to a URL specified by defaultUrlRedirect.
	// If defaultUrlRedirect is specified, defaultService or defaultRouteAction must not be set.
	// Not supported when the URL map is bound to target gRPC proxy.
	DefaultUrlRedirect HttpRedirectActionResponsePtrInput
	// An optional description of this resource. Provide this property when you create the resource.
	Description pulumi.StringPtrInput
	// Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking. This field will be ignored when inserting a UrlMap. An up-to-date fingerprint must be provided in order to update the UrlMap, otherwise the request will fail with error 412 conditionNotMet.
	//
	// To see the latest fingerprint, make a get() request to retrieve a UrlMap.
	Fingerprint pulumi.StringPtrInput
	// Specifies changes to request and response headers that need to take effect for the selected backendService.
	// The headerAction specified here take effect after headerAction specified under pathMatcher.
	// Note that headerAction is not supported for Loadbalancers that have their loadBalancingScheme set to EXTERNAL.
	// Not supported when the URL map is bound to target gRPC proxy that has validateForProxyless field set to true.
	HeaderAction HttpHeaderActionResponsePtrInput
	// The list of HostRules to use against the URL.
	HostRules HostRuleResponseArrayInput
	// [Output Only] Type of the resource. Always compute#urlMaps for url maps.
	Kind pulumi.StringPtrInput
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name pulumi.StringPtrInput
	// The list of named PathMatchers to use against the URL.
	PathMatchers PathMatcherResponseArrayInput
	// [Output Only] URL of the region where the regional URL map resides. This field is not applicable to global URL maps. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
	Region pulumi.StringPtrInput
	// [Output Only] Server-defined URL for the resource.
	SelfLink pulumi.StringPtrInput
	// The list of expected URL mapping tests. Request to update this UrlMap will succeed only if all of the test cases pass. You can specify a maximum of 100 tests per UrlMap.
	// Not supported when the URL map is bound to target gRPC proxy that has validateForProxyless field set to true.
	Tests UrlMapTestResponseArrayInput
}

func (RegionUrlMapState) ElementType() reflect.Type {
	return reflect.TypeOf((*regionUrlMapState)(nil)).Elem()
}

type regionUrlMapArgs struct {
	// [Output Only] Creation timestamp in RFC3339 text format.
	CreationTimestamp *string `pulumi:"creationTimestamp"`
	// defaultRouteAction takes effect when none of the  hostRules match. The load balancer performs advanced routing actions like URL rewrites, header transformations, etc. prior to forwarding the request to the selected backend. If defaultRouteAction specifies any weightedBackendServices, defaultService must not be set. Conversely if defaultService is set, defaultRouteAction cannot contain any  weightedBackendServices.
	// Only one of defaultRouteAction or defaultUrlRedirect must be set.
	// UrlMaps for external HTTP(S) load balancers support only the urlRewrite action within defaultRouteAction.
	// defaultRouteAction has no effect when the URL map is bound to target gRPC proxy that has validateForProxyless field set to true.
	DefaultRouteAction *HttpRouteAction `pulumi:"defaultRouteAction"`
	// The full or partial URL of the defaultService resource to which traffic is directed if none of the hostRules match. If defaultRouteAction is additionally specified, advanced routing actions like URL Rewrites, etc. take effect prior to sending the request to the backend. However, if defaultService is specified, defaultRouteAction cannot contain any weightedBackendServices. Conversely, if routeAction specifies any weightedBackendServices, service must not be specified.
	// Only one of defaultService, defaultUrlRedirect  or defaultRouteAction.weightedBackendService must be set.
	// defaultService has no effect when the URL map is bound to target gRPC proxy that has validateForProxyless field set to true.
	DefaultService *string `pulumi:"defaultService"`
	// When none of the specified hostRules match, the request is redirected to a URL specified by defaultUrlRedirect.
	// If defaultUrlRedirect is specified, defaultService or defaultRouteAction must not be set.
	// Not supported when the URL map is bound to target gRPC proxy.
	DefaultUrlRedirect *HttpRedirectAction `pulumi:"defaultUrlRedirect"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description *string `pulumi:"description"`
	// Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking. This field will be ignored when inserting a UrlMap. An up-to-date fingerprint must be provided in order to update the UrlMap, otherwise the request will fail with error 412 conditionNotMet.
	//
	// To see the latest fingerprint, make a get() request to retrieve a UrlMap.
	Fingerprint *string `pulumi:"fingerprint"`
	// Specifies changes to request and response headers that need to take effect for the selected backendService.
	// The headerAction specified here take effect after headerAction specified under pathMatcher.
	// Note that headerAction is not supported for Loadbalancers that have their loadBalancingScheme set to EXTERNAL.
	// Not supported when the URL map is bound to target gRPC proxy that has validateForProxyless field set to true.
	HeaderAction *HttpHeaderAction `pulumi:"headerAction"`
	// The list of HostRules to use against the URL.
	HostRules []HostRule `pulumi:"hostRules"`
	// [Output Only] The unique identifier for the resource. This identifier is defined by the server.
	Id *string `pulumi:"id"`
	// [Output Only] Type of the resource. Always compute#urlMaps for url maps.
	Kind *string `pulumi:"kind"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name *string `pulumi:"name"`
	// The list of named PathMatchers to use against the URL.
	PathMatchers []PathMatcher `pulumi:"pathMatchers"`
	Project      string        `pulumi:"project"`
	// [Output Only] URL of the region where the regional URL map resides. This field is not applicable to global URL maps. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
	Region    string  `pulumi:"region"`
	RequestId *string `pulumi:"requestId"`
	// [Output Only] Server-defined URL for the resource.
	SelfLink *string `pulumi:"selfLink"`
	// The list of expected URL mapping tests. Request to update this UrlMap will succeed only if all of the test cases pass. You can specify a maximum of 100 tests per UrlMap.
	// Not supported when the URL map is bound to target gRPC proxy that has validateForProxyless field set to true.
	Tests []UrlMapTest `pulumi:"tests"`
}

// The set of arguments for constructing a RegionUrlMap resource.
type RegionUrlMapArgs struct {
	// [Output Only] Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringPtrInput
	// defaultRouteAction takes effect when none of the  hostRules match. The load balancer performs advanced routing actions like URL rewrites, header transformations, etc. prior to forwarding the request to the selected backend. If defaultRouteAction specifies any weightedBackendServices, defaultService must not be set. Conversely if defaultService is set, defaultRouteAction cannot contain any  weightedBackendServices.
	// Only one of defaultRouteAction or defaultUrlRedirect must be set.
	// UrlMaps for external HTTP(S) load balancers support only the urlRewrite action within defaultRouteAction.
	// defaultRouteAction has no effect when the URL map is bound to target gRPC proxy that has validateForProxyless field set to true.
	DefaultRouteAction HttpRouteActionPtrInput
	// The full or partial URL of the defaultService resource to which traffic is directed if none of the hostRules match. If defaultRouteAction is additionally specified, advanced routing actions like URL Rewrites, etc. take effect prior to sending the request to the backend. However, if defaultService is specified, defaultRouteAction cannot contain any weightedBackendServices. Conversely, if routeAction specifies any weightedBackendServices, service must not be specified.
	// Only one of defaultService, defaultUrlRedirect  or defaultRouteAction.weightedBackendService must be set.
	// defaultService has no effect when the URL map is bound to target gRPC proxy that has validateForProxyless field set to true.
	DefaultService pulumi.StringPtrInput
	// When none of the specified hostRules match, the request is redirected to a URL specified by defaultUrlRedirect.
	// If defaultUrlRedirect is specified, defaultService or defaultRouteAction must not be set.
	// Not supported when the URL map is bound to target gRPC proxy.
	DefaultUrlRedirect HttpRedirectActionPtrInput
	// An optional description of this resource. Provide this property when you create the resource.
	Description pulumi.StringPtrInput
	// Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking. This field will be ignored when inserting a UrlMap. An up-to-date fingerprint must be provided in order to update the UrlMap, otherwise the request will fail with error 412 conditionNotMet.
	//
	// To see the latest fingerprint, make a get() request to retrieve a UrlMap.
	Fingerprint pulumi.StringPtrInput
	// Specifies changes to request and response headers that need to take effect for the selected backendService.
	// The headerAction specified here take effect after headerAction specified under pathMatcher.
	// Note that headerAction is not supported for Loadbalancers that have their loadBalancingScheme set to EXTERNAL.
	// Not supported when the URL map is bound to target gRPC proxy that has validateForProxyless field set to true.
	HeaderAction HttpHeaderActionPtrInput
	// The list of HostRules to use against the URL.
	HostRules HostRuleArrayInput
	// [Output Only] The unique identifier for the resource. This identifier is defined by the server.
	Id pulumi.StringPtrInput
	// [Output Only] Type of the resource. Always compute#urlMaps for url maps.
	Kind pulumi.StringPtrInput
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name pulumi.StringPtrInput
	// The list of named PathMatchers to use against the URL.
	PathMatchers PathMatcherArrayInput
	Project      pulumi.StringInput
	// [Output Only] URL of the region where the regional URL map resides. This field is not applicable to global URL maps. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
	Region    pulumi.StringInput
	RequestId pulumi.StringPtrInput
	// [Output Only] Server-defined URL for the resource.
	SelfLink pulumi.StringPtrInput
	// The list of expected URL mapping tests. Request to update this UrlMap will succeed only if all of the test cases pass. You can specify a maximum of 100 tests per UrlMap.
	// Not supported when the URL map is bound to target gRPC proxy that has validateForProxyless field set to true.
	Tests UrlMapTestArrayInput
}

func (RegionUrlMapArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*regionUrlMapArgs)(nil)).Elem()
}

type RegionUrlMapInput interface {
	pulumi.Input

	ToRegionUrlMapOutput() RegionUrlMapOutput
	ToRegionUrlMapOutputWithContext(ctx context.Context) RegionUrlMapOutput
}

func (*RegionUrlMap) ElementType() reflect.Type {
	return reflect.TypeOf((*RegionUrlMap)(nil))
}

func (i *RegionUrlMap) ToRegionUrlMapOutput() RegionUrlMapOutput {
	return i.ToRegionUrlMapOutputWithContext(context.Background())
}

func (i *RegionUrlMap) ToRegionUrlMapOutputWithContext(ctx context.Context) RegionUrlMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegionUrlMapOutput)
}

type RegionUrlMapOutput struct {
	*pulumi.OutputState
}

func (RegionUrlMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RegionUrlMap)(nil))
}

func (o RegionUrlMapOutput) ToRegionUrlMapOutput() RegionUrlMapOutput {
	return o
}

func (o RegionUrlMapOutput) ToRegionUrlMapOutputWithContext(ctx context.Context) RegionUrlMapOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(RegionUrlMapOutput{})
}
