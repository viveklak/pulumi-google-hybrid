// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a new EndpointPolicy in a given project and location.
type EndpointPolicy struct {
	pulumi.CustomResourceState

	// Optional. This field specifies the URL of AuthorizationPolicy resource that applies authorization policies to the inbound traffic at the matched endpoints. Refer to Authorization. If this field is not specified, authorization is disabled(no authz checks) for this endpoint.
	AuthorizationPolicy pulumi.StringOutput `pulumi:"authorizationPolicy"`
	// Optional. A URL referring to a ClientTlsPolicy resource. ClientTlsPolicy can be set to specify the authentication for traffic from the proxy to the actual endpoints. More specifically, it is applied to the outgoing traffic from the proxy to the endpoint. This is typically used for sidecar model where the proxy identifies itself as endpoint to the control plane, with the connection between sidecar and endpoint requiring authentication. If this field is not set, authentication is disabled(open). Applicable only when EndpointPolicyType is SIDECAR_PROXY.
	ClientTlsPolicy pulumi.StringOutput `pulumi:"clientTlsPolicy"`
	// The timestamp when the resource was created.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Optional. A free-text description of the resource. Max length 1024 characters.
	Description pulumi.StringOutput `pulumi:"description"`
	// A matcher that selects endpoints to which the policies should be applied.
	EndpointMatcher EndpointMatcherResponseOutput `pulumi:"endpointMatcher"`
	// Optional. Set of label tags associated with the EndpointPolicy resource.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// Name of the EndpointPolicy resource. It matches pattern `projects/{project}/locations/global/endpointPolicies/{endpoint_policy}`.
	Name pulumi.StringOutput `pulumi:"name"`
	// Optional. A URL referring to ServerTlsPolicy resource. ServerTlsPolicy is used to determine the authentication policy to be applied to terminate the inbound traffic at the identified backends. If this field is not set, authentication is disabled(open) for this endpoint.
	ServerTlsPolicy pulumi.StringOutput `pulumi:"serverTlsPolicy"`
	// Optional. Port selector for the (matched) endpoints. If no port selector is provided, the matched config is applied to all ports.
	TrafficPortSelector TrafficPortSelectorResponseOutput `pulumi:"trafficPortSelector"`
	// The type of endpoint policy. This is primarily used to validate the configuration.
	Type pulumi.StringOutput `pulumi:"type"`
	// The timestamp when the resource was updated.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewEndpointPolicy registers a new resource with the given unique name, arguments, and options.
func NewEndpointPolicy(ctx *pulumi.Context,
	name string, args *EndpointPolicyArgs, opts ...pulumi.ResourceOption) (*EndpointPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EndpointMatcher == nil {
		return nil, errors.New("invalid value for required argument 'EndpointMatcher'")
	}
	if args.EndpointPolicyId == nil {
		return nil, errors.New("invalid value for required argument 'EndpointPolicyId'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	var resource EndpointPolicy
	err := ctx.RegisterResource("google-native:networkservices/v1beta1:EndpointPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEndpointPolicy gets an existing EndpointPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEndpointPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EndpointPolicyState, opts ...pulumi.ResourceOption) (*EndpointPolicy, error) {
	var resource EndpointPolicy
	err := ctx.ReadResource("google-native:networkservices/v1beta1:EndpointPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EndpointPolicy resources.
type endpointPolicyState struct {
}

type EndpointPolicyState struct {
}

func (EndpointPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*endpointPolicyState)(nil)).Elem()
}

type endpointPolicyArgs struct {
	// Optional. This field specifies the URL of AuthorizationPolicy resource that applies authorization policies to the inbound traffic at the matched endpoints. Refer to Authorization. If this field is not specified, authorization is disabled(no authz checks) for this endpoint.
	AuthorizationPolicy *string `pulumi:"authorizationPolicy"`
	// Optional. A URL referring to a ClientTlsPolicy resource. ClientTlsPolicy can be set to specify the authentication for traffic from the proxy to the actual endpoints. More specifically, it is applied to the outgoing traffic from the proxy to the endpoint. This is typically used for sidecar model where the proxy identifies itself as endpoint to the control plane, with the connection between sidecar and endpoint requiring authentication. If this field is not set, authentication is disabled(open). Applicable only when EndpointPolicyType is SIDECAR_PROXY.
	ClientTlsPolicy *string `pulumi:"clientTlsPolicy"`
	// Optional. A free-text description of the resource. Max length 1024 characters.
	Description *string `pulumi:"description"`
	// A matcher that selects endpoints to which the policies should be applied.
	EndpointMatcher EndpointMatcher `pulumi:"endpointMatcher"`
	// Required. Short name of the EndpointPolicy resource to be created. E.g. "CustomECS".
	EndpointPolicyId string `pulumi:"endpointPolicyId"`
	// Optional. Set of label tags associated with the EndpointPolicy resource.
	Labels   map[string]string `pulumi:"labels"`
	Location *string           `pulumi:"location"`
	// Name of the EndpointPolicy resource. It matches pattern `projects/{project}/locations/global/endpointPolicies/{endpoint_policy}`.
	Name    *string `pulumi:"name"`
	Project *string `pulumi:"project"`
	// Optional. A URL referring to ServerTlsPolicy resource. ServerTlsPolicy is used to determine the authentication policy to be applied to terminate the inbound traffic at the identified backends. If this field is not set, authentication is disabled(open) for this endpoint.
	ServerTlsPolicy *string `pulumi:"serverTlsPolicy"`
	// Optional. Port selector for the (matched) endpoints. If no port selector is provided, the matched config is applied to all ports.
	TrafficPortSelector *TrafficPortSelector `pulumi:"trafficPortSelector"`
	// The type of endpoint policy. This is primarily used to validate the configuration.
	Type EndpointPolicyType `pulumi:"type"`
}

// The set of arguments for constructing a EndpointPolicy resource.
type EndpointPolicyArgs struct {
	// Optional. This field specifies the URL of AuthorizationPolicy resource that applies authorization policies to the inbound traffic at the matched endpoints. Refer to Authorization. If this field is not specified, authorization is disabled(no authz checks) for this endpoint.
	AuthorizationPolicy pulumi.StringPtrInput
	// Optional. A URL referring to a ClientTlsPolicy resource. ClientTlsPolicy can be set to specify the authentication for traffic from the proxy to the actual endpoints. More specifically, it is applied to the outgoing traffic from the proxy to the endpoint. This is typically used for sidecar model where the proxy identifies itself as endpoint to the control plane, with the connection between sidecar and endpoint requiring authentication. If this field is not set, authentication is disabled(open). Applicable only when EndpointPolicyType is SIDECAR_PROXY.
	ClientTlsPolicy pulumi.StringPtrInput
	// Optional. A free-text description of the resource. Max length 1024 characters.
	Description pulumi.StringPtrInput
	// A matcher that selects endpoints to which the policies should be applied.
	EndpointMatcher EndpointMatcherInput
	// Required. Short name of the EndpointPolicy resource to be created. E.g. "CustomECS".
	EndpointPolicyId pulumi.StringInput
	// Optional. Set of label tags associated with the EndpointPolicy resource.
	Labels   pulumi.StringMapInput
	Location pulumi.StringPtrInput
	// Name of the EndpointPolicy resource. It matches pattern `projects/{project}/locations/global/endpointPolicies/{endpoint_policy}`.
	Name    pulumi.StringPtrInput
	Project pulumi.StringPtrInput
	// Optional. A URL referring to ServerTlsPolicy resource. ServerTlsPolicy is used to determine the authentication policy to be applied to terminate the inbound traffic at the identified backends. If this field is not set, authentication is disabled(open) for this endpoint.
	ServerTlsPolicy pulumi.StringPtrInput
	// Optional. Port selector for the (matched) endpoints. If no port selector is provided, the matched config is applied to all ports.
	TrafficPortSelector TrafficPortSelectorPtrInput
	// The type of endpoint policy. This is primarily used to validate the configuration.
	Type EndpointPolicyTypeInput
}

func (EndpointPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*endpointPolicyArgs)(nil)).Elem()
}

type EndpointPolicyInput interface {
	pulumi.Input

	ToEndpointPolicyOutput() EndpointPolicyOutput
	ToEndpointPolicyOutputWithContext(ctx context.Context) EndpointPolicyOutput
}

func (*EndpointPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**EndpointPolicy)(nil)).Elem()
}

func (i *EndpointPolicy) ToEndpointPolicyOutput() EndpointPolicyOutput {
	return i.ToEndpointPolicyOutputWithContext(context.Background())
}

func (i *EndpointPolicy) ToEndpointPolicyOutputWithContext(ctx context.Context) EndpointPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointPolicyOutput)
}

type EndpointPolicyOutput struct{ *pulumi.OutputState }

func (EndpointPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EndpointPolicy)(nil)).Elem()
}

func (o EndpointPolicyOutput) ToEndpointPolicyOutput() EndpointPolicyOutput {
	return o
}

func (o EndpointPolicyOutput) ToEndpointPolicyOutputWithContext(ctx context.Context) EndpointPolicyOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EndpointPolicyInput)(nil)).Elem(), &EndpointPolicy{})
	pulumi.RegisterOutputType(EndpointPolicyOutput{})
}
