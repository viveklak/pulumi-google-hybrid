// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1alpha1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this method to create a route for a private connectivity in a project and location.
// Auto-naming is currently not supported for this resource.
type Route struct {
	pulumi.CustomResourceState

	// The create time of the resource.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Destination address for connection
	DestinationAddress pulumi.StringOutput `pulumi:"destinationAddress"`
	// Destination port for connection
	DestinationPort pulumi.IntOutput `pulumi:"destinationPort"`
	// Display name.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// Labels.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// The resource's name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The update time of the resource.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewRoute registers a new resource with the given unique name, arguments, and options.
func NewRoute(ctx *pulumi.Context,
	name string, args *RouteArgs, opts ...pulumi.ResourceOption) (*Route, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DestinationAddress == nil {
		return nil, errors.New("invalid value for required argument 'DestinationAddress'")
	}
	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	if args.PrivateConnectionId == nil {
		return nil, errors.New("invalid value for required argument 'PrivateConnectionId'")
	}
	if args.RouteId == nil {
		return nil, errors.New("invalid value for required argument 'RouteId'")
	}
	var resource Route
	err := ctx.RegisterResource("google-native:datastream/v1alpha1:Route", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRoute gets an existing Route resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRoute(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RouteState, opts ...pulumi.ResourceOption) (*Route, error) {
	var resource Route
	err := ctx.ReadResource("google-native:datastream/v1alpha1:Route", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Route resources.
type routeState struct {
}

type RouteState struct {
}

func (RouteState) ElementType() reflect.Type {
	return reflect.TypeOf((*routeState)(nil)).Elem()
}

type routeArgs struct {
	// Destination address for connection
	DestinationAddress string `pulumi:"destinationAddress"`
	// Destination port for connection
	DestinationPort *int `pulumi:"destinationPort"`
	// Display name.
	DisplayName string `pulumi:"displayName"`
	// Labels.
	Labels              map[string]string `pulumi:"labels"`
	Location            *string           `pulumi:"location"`
	PrivateConnectionId string            `pulumi:"privateConnectionId"`
	Project             *string           `pulumi:"project"`
	RequestId           *string           `pulumi:"requestId"`
	RouteId             string            `pulumi:"routeId"`
}

// The set of arguments for constructing a Route resource.
type RouteArgs struct {
	// Destination address for connection
	DestinationAddress pulumi.StringInput
	// Destination port for connection
	DestinationPort pulumi.IntPtrInput
	// Display name.
	DisplayName pulumi.StringInput
	// Labels.
	Labels              pulumi.StringMapInput
	Location            pulumi.StringPtrInput
	PrivateConnectionId pulumi.StringInput
	Project             pulumi.StringPtrInput
	RequestId           pulumi.StringPtrInput
	RouteId             pulumi.StringInput
}

func (RouteArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*routeArgs)(nil)).Elem()
}

type RouteInput interface {
	pulumi.Input

	ToRouteOutput() RouteOutput
	ToRouteOutputWithContext(ctx context.Context) RouteOutput
}

func (*Route) ElementType() reflect.Type {
	return reflect.TypeOf((**Route)(nil)).Elem()
}

func (i *Route) ToRouteOutput() RouteOutput {
	return i.ToRouteOutputWithContext(context.Background())
}

func (i *Route) ToRouteOutputWithContext(ctx context.Context) RouteOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouteOutput)
}

type RouteOutput struct{ *pulumi.OutputState }

func (RouteOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Route)(nil)).Elem()
}

func (o RouteOutput) ToRouteOutput() RouteOutput {
	return o
}

func (o RouteOutput) ToRouteOutputWithContext(ctx context.Context) RouteOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RouteInput)(nil)).Elem(), &Route{})
	pulumi.RegisterOutputType(RouteOutput{})
}
