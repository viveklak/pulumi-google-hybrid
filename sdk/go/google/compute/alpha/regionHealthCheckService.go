// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package alpha

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a regional HealthCheckService resource in the specified project and region using the data included in the request.
type RegionHealthCheckService struct {
	pulumi.CustomResourceState

	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringOutput `pulumi:"creationTimestamp"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description pulumi.StringOutput `pulumi:"description"`
	// Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking. This field will be ignored when inserting a HealthCheckService. An up-to-date fingerprint must be provided in order to patch/update the HealthCheckService; Otherwise, the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make a get() request to retrieve the HealthCheckService.
	Fingerprint pulumi.StringOutput `pulumi:"fingerprint"`
	// List of URLs to the HealthCheck resources. Must have at least one HealthCheck, and not more than 10. HealthCheck resources must have portSpecification=USE_SERVING_PORT. For regional HealthCheckService, the HealthCheck must be regional and in the same region. For global HealthCheckService, HealthCheck must be global. Mix of regional and global HealthChecks is not supported. Multiple regional HealthChecks must belong to the same region. Regional HealthChecks</code? must belong to the same region as zones of NEGs.
	HealthChecks pulumi.StringArrayOutput `pulumi:"healthChecks"`
	// Optional. Policy for how the results from multiple health checks for the same endpoint are aggregated. Defaults to NO_AGGREGATION if unspecified.
	// - NO_AGGREGATION. An EndpointHealth message is returned for each backend in the health check service.
	// - AND. If any backend's health check reports UNHEALTHY, then UNHEALTHY is the HealthState of the entire health check service. If all backend's are healthy, the HealthState of the health check service is HEALTHY. .
	HealthStatusAggregationPolicy pulumi.StringOutput `pulumi:"healthStatusAggregationPolicy"`
	// [Output only] Type of the resource. Always compute#healthCheckServicefor health check services.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Name of the resource. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name pulumi.StringOutput `pulumi:"name"`
	// List of URLs to the NetworkEndpointGroup resources. Must not have more than 100. For regional HealthCheckService, NEGs must be in zones in the region of the HealthCheckService.
	NetworkEndpointGroups pulumi.StringArrayOutput `pulumi:"networkEndpointGroups"`
	// List of URLs to the NotificationEndpoint resources. Must not have more than 10. A list of endpoints for receiving notifications of change in health status. For regional HealthCheckService, NotificationEndpoint must be regional and in the same region. For global HealthCheckService, NotificationEndpoint must be global.
	NotificationEndpoints pulumi.StringArrayOutput `pulumi:"notificationEndpoints"`
	// URL of the region where the health check service resides. This field is not applicable to global health check services. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
	Region pulumi.StringOutput `pulumi:"region"`
	// Server-defined URL for the resource.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
	// Server-defined URL with id for the resource.
	SelfLinkWithId pulumi.StringOutput `pulumi:"selfLinkWithId"`
}

// NewRegionHealthCheckService registers a new resource with the given unique name, arguments, and options.
func NewRegionHealthCheckService(ctx *pulumi.Context,
	name string, args *RegionHealthCheckServiceArgs, opts ...pulumi.ResourceOption) (*RegionHealthCheckService, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	if args.Region == nil {
		return nil, errors.New("invalid value for required argument 'Region'")
	}
	var resource RegionHealthCheckService
	err := ctx.RegisterResource("google-native:compute/alpha:RegionHealthCheckService", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRegionHealthCheckService gets an existing RegionHealthCheckService resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRegionHealthCheckService(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RegionHealthCheckServiceState, opts ...pulumi.ResourceOption) (*RegionHealthCheckService, error) {
	var resource RegionHealthCheckService
	err := ctx.ReadResource("google-native:compute/alpha:RegionHealthCheckService", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RegionHealthCheckService resources.
type regionHealthCheckServiceState struct {
}

type RegionHealthCheckServiceState struct {
}

func (RegionHealthCheckServiceState) ElementType() reflect.Type {
	return reflect.TypeOf((*regionHealthCheckServiceState)(nil)).Elem()
}

type regionHealthCheckServiceArgs struct {
	// An optional description of this resource. Provide this property when you create the resource.
	Description *string `pulumi:"description"`
	// List of URLs to the HealthCheck resources. Must have at least one HealthCheck, and not more than 10. HealthCheck resources must have portSpecification=USE_SERVING_PORT. For regional HealthCheckService, the HealthCheck must be regional and in the same region. For global HealthCheckService, HealthCheck must be global. Mix of regional and global HealthChecks is not supported. Multiple regional HealthChecks must belong to the same region. Regional HealthChecks</code? must belong to the same region as zones of NEGs.
	HealthChecks []string `pulumi:"healthChecks"`
	// Optional. Policy for how the results from multiple health checks for the same endpoint are aggregated. Defaults to NO_AGGREGATION if unspecified.
	// - NO_AGGREGATION. An EndpointHealth message is returned for each backend in the health check service.
	// - AND. If any backend's health check reports UNHEALTHY, then UNHEALTHY is the HealthState of the entire health check service. If all backend's are healthy, the HealthState of the health check service is HEALTHY. .
	HealthStatusAggregationPolicy *string `pulumi:"healthStatusAggregationPolicy"`
	// Name of the resource. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name *string `pulumi:"name"`
	// List of URLs to the NetworkEndpointGroup resources. Must not have more than 100. For regional HealthCheckService, NEGs must be in zones in the region of the HealthCheckService.
	NetworkEndpointGroups []string `pulumi:"networkEndpointGroups"`
	// List of URLs to the NotificationEndpoint resources. Must not have more than 10. A list of endpoints for receiving notifications of change in health status. For regional HealthCheckService, NotificationEndpoint must be regional and in the same region. For global HealthCheckService, NotificationEndpoint must be global.
	NotificationEndpoints []string `pulumi:"notificationEndpoints"`
	Project               string   `pulumi:"project"`
	Region                string   `pulumi:"region"`
	RequestId             *string  `pulumi:"requestId"`
}

// The set of arguments for constructing a RegionHealthCheckService resource.
type RegionHealthCheckServiceArgs struct {
	// An optional description of this resource. Provide this property when you create the resource.
	Description pulumi.StringPtrInput
	// List of URLs to the HealthCheck resources. Must have at least one HealthCheck, and not more than 10. HealthCheck resources must have portSpecification=USE_SERVING_PORT. For regional HealthCheckService, the HealthCheck must be regional and in the same region. For global HealthCheckService, HealthCheck must be global. Mix of regional and global HealthChecks is not supported. Multiple regional HealthChecks must belong to the same region. Regional HealthChecks</code? must belong to the same region as zones of NEGs.
	HealthChecks pulumi.StringArrayInput
	// Optional. Policy for how the results from multiple health checks for the same endpoint are aggregated. Defaults to NO_AGGREGATION if unspecified.
	// - NO_AGGREGATION. An EndpointHealth message is returned for each backend in the health check service.
	// - AND. If any backend's health check reports UNHEALTHY, then UNHEALTHY is the HealthState of the entire health check service. If all backend's are healthy, the HealthState of the health check service is HEALTHY. .
	HealthStatusAggregationPolicy *RegionHealthCheckServiceHealthStatusAggregationPolicy
	// Name of the resource. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name pulumi.StringPtrInput
	// List of URLs to the NetworkEndpointGroup resources. Must not have more than 100. For regional HealthCheckService, NEGs must be in zones in the region of the HealthCheckService.
	NetworkEndpointGroups pulumi.StringArrayInput
	// List of URLs to the NotificationEndpoint resources. Must not have more than 10. A list of endpoints for receiving notifications of change in health status. For regional HealthCheckService, NotificationEndpoint must be regional and in the same region. For global HealthCheckService, NotificationEndpoint must be global.
	NotificationEndpoints pulumi.StringArrayInput
	Project               pulumi.StringInput
	Region                pulumi.StringInput
	RequestId             pulumi.StringPtrInput
}

func (RegionHealthCheckServiceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*regionHealthCheckServiceArgs)(nil)).Elem()
}

type RegionHealthCheckServiceInput interface {
	pulumi.Input

	ToRegionHealthCheckServiceOutput() RegionHealthCheckServiceOutput
	ToRegionHealthCheckServiceOutputWithContext(ctx context.Context) RegionHealthCheckServiceOutput
}

func (*RegionHealthCheckService) ElementType() reflect.Type {
	return reflect.TypeOf((*RegionHealthCheckService)(nil))
}

func (i *RegionHealthCheckService) ToRegionHealthCheckServiceOutput() RegionHealthCheckServiceOutput {
	return i.ToRegionHealthCheckServiceOutputWithContext(context.Background())
}

func (i *RegionHealthCheckService) ToRegionHealthCheckServiceOutputWithContext(ctx context.Context) RegionHealthCheckServiceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegionHealthCheckServiceOutput)
}

type RegionHealthCheckServiceOutput struct {
	*pulumi.OutputState
}

func (RegionHealthCheckServiceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RegionHealthCheckService)(nil))
}

func (o RegionHealthCheckServiceOutput) ToRegionHealthCheckServiceOutput() RegionHealthCheckServiceOutput {
	return o
}

func (o RegionHealthCheckServiceOutput) ToRegionHealthCheckServiceOutputWithContext(ctx context.Context) RegionHealthCheckServiceOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(RegionHealthCheckServiceOutput{})
}
