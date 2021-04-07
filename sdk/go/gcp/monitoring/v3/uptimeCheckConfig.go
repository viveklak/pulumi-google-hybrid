// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v3

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Creates a new Uptime check configuration.
type UptimeCheckConfig struct {
	pulumi.CustomResourceState

	// The content that is expected to appear in the data returned by the target server against which the check is run. Currently, only the first entry in the content_matchers list is supported, and additional entries will be ignored. This field is optional and should only be specified if a content match is required as part of the/ Uptime check.
	ContentMatchers ContentMatcherResponseArrayOutput `pulumi:"contentMatchers"`
	// A human-friendly name for the Uptime check configuration. The display name should be unique within a Stackdriver Workspace in order to make it easier to identify; however, uniqueness is not enforced. Required.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// Contains information needed to make an HTTP or HTTPS check.
	HttpCheck HttpCheckResponseOutput `pulumi:"httpCheck"`
	// The internal checkers that this check will egress from. If is_internal is true and this list is empty, the check will egress from all the InternalCheckers configured for the project that owns this UptimeCheckConfig.
	InternalCheckers InternalCheckerResponseArrayOutput `pulumi:"internalCheckers"`
	// If this is true, then checks are made only from the 'internal_checkers'. If it is false, then checks are made only from the 'selected_regions'. It is an error to provide 'selected_regions' when is_internal is true, or to provide 'internal_checkers' when is_internal is false.
	IsInternal pulumi.BoolOutput `pulumi:"isInternal"`
	// The monitored resource (https://cloud.google.com/monitoring/api/resources) associated with the configuration. The following monitored resource types are valid for this field: uptime_url, gce_instance, gae_app, aws_ec2_instance, aws_elb_load_balancer
	MonitoredResource MonitoredResourceResponseOutput `pulumi:"monitoredResource"`
	// A unique resource name for this Uptime check configuration. The format is: projects/[PROJECT_ID_OR_NUMBER]/uptimeCheckConfigs/[UPTIME_CHECK_ID] [PROJECT_ID_OR_NUMBER] is the Workspace host project associated with the Uptime check.This field should be omitted when creating the Uptime check configuration; on create, the resource name is assigned by the server and included in the response.
	Name pulumi.StringOutput `pulumi:"name"`
	// How often, in seconds, the Uptime check is performed. Currently, the only supported values are 60s (1 minute), 300s (5 minutes), 600s (10 minutes), and 900s (15 minutes). Optional, defaults to 60s.
	Period pulumi.StringOutput `pulumi:"period"`
	// The group resource associated with the configuration.
	ResourceGroup ResourceGroupResponseOutput `pulumi:"resourceGroup"`
	// The list of regions from which the check will be run. Some regions contain one location, and others contain more than one. If this field is specified, enough regions must be provided to include a minimum of 3 locations. Not specifying this field will result in Uptime checks running from all available regions.
	SelectedRegions pulumi.StringArrayOutput `pulumi:"selectedRegions"`
	// Contains information needed to make a TCP check.
	TcpCheck TcpCheckResponseOutput `pulumi:"tcpCheck"`
	// The maximum amount of time to wait for the request to complete (must be between 1 and 60 seconds). Required.
	Timeout pulumi.StringOutput `pulumi:"timeout"`
}

// NewUptimeCheckConfig registers a new resource with the given unique name, arguments, and options.
func NewUptimeCheckConfig(ctx *pulumi.Context,
	name string, args *UptimeCheckConfigArgs, opts ...pulumi.ResourceOption) (*UptimeCheckConfig, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ProjectsId == nil {
		return nil, errors.New("invalid value for required argument 'ProjectsId'")
	}
	if args.UptimeCheckConfigsId == nil {
		return nil, errors.New("invalid value for required argument 'UptimeCheckConfigsId'")
	}
	var resource UptimeCheckConfig
	err := ctx.RegisterResource("gcp-native:monitoring/v3:UptimeCheckConfig", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUptimeCheckConfig gets an existing UptimeCheckConfig resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUptimeCheckConfig(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UptimeCheckConfigState, opts ...pulumi.ResourceOption) (*UptimeCheckConfig, error) {
	var resource UptimeCheckConfig
	err := ctx.ReadResource("gcp-native:monitoring/v3:UptimeCheckConfig", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering UptimeCheckConfig resources.
type uptimeCheckConfigState struct {
	// The content that is expected to appear in the data returned by the target server against which the check is run. Currently, only the first entry in the content_matchers list is supported, and additional entries will be ignored. This field is optional and should only be specified if a content match is required as part of the/ Uptime check.
	ContentMatchers []ContentMatcherResponse `pulumi:"contentMatchers"`
	// A human-friendly name for the Uptime check configuration. The display name should be unique within a Stackdriver Workspace in order to make it easier to identify; however, uniqueness is not enforced. Required.
	DisplayName *string `pulumi:"displayName"`
	// Contains information needed to make an HTTP or HTTPS check.
	HttpCheck *HttpCheckResponse `pulumi:"httpCheck"`
	// The internal checkers that this check will egress from. If is_internal is true and this list is empty, the check will egress from all the InternalCheckers configured for the project that owns this UptimeCheckConfig.
	InternalCheckers []InternalCheckerResponse `pulumi:"internalCheckers"`
	// If this is true, then checks are made only from the 'internal_checkers'. If it is false, then checks are made only from the 'selected_regions'. It is an error to provide 'selected_regions' when is_internal is true, or to provide 'internal_checkers' when is_internal is false.
	IsInternal *bool `pulumi:"isInternal"`
	// The monitored resource (https://cloud.google.com/monitoring/api/resources) associated with the configuration. The following monitored resource types are valid for this field: uptime_url, gce_instance, gae_app, aws_ec2_instance, aws_elb_load_balancer
	MonitoredResource *MonitoredResourceResponse `pulumi:"monitoredResource"`
	// A unique resource name for this Uptime check configuration. The format is: projects/[PROJECT_ID_OR_NUMBER]/uptimeCheckConfigs/[UPTIME_CHECK_ID] [PROJECT_ID_OR_NUMBER] is the Workspace host project associated with the Uptime check.This field should be omitted when creating the Uptime check configuration; on create, the resource name is assigned by the server and included in the response.
	Name *string `pulumi:"name"`
	// How often, in seconds, the Uptime check is performed. Currently, the only supported values are 60s (1 minute), 300s (5 minutes), 600s (10 minutes), and 900s (15 minutes). Optional, defaults to 60s.
	Period *string `pulumi:"period"`
	// The group resource associated with the configuration.
	ResourceGroup *ResourceGroupResponse `pulumi:"resourceGroup"`
	// The list of regions from which the check will be run. Some regions contain one location, and others contain more than one. If this field is specified, enough regions must be provided to include a minimum of 3 locations. Not specifying this field will result in Uptime checks running from all available regions.
	SelectedRegions []string `pulumi:"selectedRegions"`
	// Contains information needed to make a TCP check.
	TcpCheck *TcpCheckResponse `pulumi:"tcpCheck"`
	// The maximum amount of time to wait for the request to complete (must be between 1 and 60 seconds). Required.
	Timeout *string `pulumi:"timeout"`
}

type UptimeCheckConfigState struct {
	// The content that is expected to appear in the data returned by the target server against which the check is run. Currently, only the first entry in the content_matchers list is supported, and additional entries will be ignored. This field is optional and should only be specified if a content match is required as part of the/ Uptime check.
	ContentMatchers ContentMatcherResponseArrayInput
	// A human-friendly name for the Uptime check configuration. The display name should be unique within a Stackdriver Workspace in order to make it easier to identify; however, uniqueness is not enforced. Required.
	DisplayName pulumi.StringPtrInput
	// Contains information needed to make an HTTP or HTTPS check.
	HttpCheck HttpCheckResponsePtrInput
	// The internal checkers that this check will egress from. If is_internal is true and this list is empty, the check will egress from all the InternalCheckers configured for the project that owns this UptimeCheckConfig.
	InternalCheckers InternalCheckerResponseArrayInput
	// If this is true, then checks are made only from the 'internal_checkers'. If it is false, then checks are made only from the 'selected_regions'. It is an error to provide 'selected_regions' when is_internal is true, or to provide 'internal_checkers' when is_internal is false.
	IsInternal pulumi.BoolPtrInput
	// The monitored resource (https://cloud.google.com/monitoring/api/resources) associated with the configuration. The following monitored resource types are valid for this field: uptime_url, gce_instance, gae_app, aws_ec2_instance, aws_elb_load_balancer
	MonitoredResource MonitoredResourceResponsePtrInput
	// A unique resource name for this Uptime check configuration. The format is: projects/[PROJECT_ID_OR_NUMBER]/uptimeCheckConfigs/[UPTIME_CHECK_ID] [PROJECT_ID_OR_NUMBER] is the Workspace host project associated with the Uptime check.This field should be omitted when creating the Uptime check configuration; on create, the resource name is assigned by the server and included in the response.
	Name pulumi.StringPtrInput
	// How often, in seconds, the Uptime check is performed. Currently, the only supported values are 60s (1 minute), 300s (5 minutes), 600s (10 minutes), and 900s (15 minutes). Optional, defaults to 60s.
	Period pulumi.StringPtrInput
	// The group resource associated with the configuration.
	ResourceGroup ResourceGroupResponsePtrInput
	// The list of regions from which the check will be run. Some regions contain one location, and others contain more than one. If this field is specified, enough regions must be provided to include a minimum of 3 locations. Not specifying this field will result in Uptime checks running from all available regions.
	SelectedRegions pulumi.StringArrayInput
	// Contains information needed to make a TCP check.
	TcpCheck TcpCheckResponsePtrInput
	// The maximum amount of time to wait for the request to complete (must be between 1 and 60 seconds). Required.
	Timeout pulumi.StringPtrInput
}

func (UptimeCheckConfigState) ElementType() reflect.Type {
	return reflect.TypeOf((*uptimeCheckConfigState)(nil)).Elem()
}

type uptimeCheckConfigArgs struct {
	// The content that is expected to appear in the data returned by the target server against which the check is run. Currently, only the first entry in the content_matchers list is supported, and additional entries will be ignored. This field is optional and should only be specified if a content match is required as part of the/ Uptime check.
	ContentMatchers []ContentMatcher `pulumi:"contentMatchers"`
	// A human-friendly name for the Uptime check configuration. The display name should be unique within a Stackdriver Workspace in order to make it easier to identify; however, uniqueness is not enforced. Required.
	DisplayName *string `pulumi:"displayName"`
	// Contains information needed to make an HTTP or HTTPS check.
	HttpCheck *HttpCheck `pulumi:"httpCheck"`
	// The internal checkers that this check will egress from. If is_internal is true and this list is empty, the check will egress from all the InternalCheckers configured for the project that owns this UptimeCheckConfig.
	InternalCheckers []InternalChecker `pulumi:"internalCheckers"`
	// If this is true, then checks are made only from the 'internal_checkers'. If it is false, then checks are made only from the 'selected_regions'. It is an error to provide 'selected_regions' when is_internal is true, or to provide 'internal_checkers' when is_internal is false.
	IsInternal *bool `pulumi:"isInternal"`
	// The monitored resource (https://cloud.google.com/monitoring/api/resources) associated with the configuration. The following monitored resource types are valid for this field: uptime_url, gce_instance, gae_app, aws_ec2_instance, aws_elb_load_balancer
	MonitoredResource *MonitoredResource `pulumi:"monitoredResource"`
	// A unique resource name for this Uptime check configuration. The format is: projects/[PROJECT_ID_OR_NUMBER]/uptimeCheckConfigs/[UPTIME_CHECK_ID] [PROJECT_ID_OR_NUMBER] is the Workspace host project associated with the Uptime check.This field should be omitted when creating the Uptime check configuration; on create, the resource name is assigned by the server and included in the response.
	Name *string `pulumi:"name"`
	// How often, in seconds, the Uptime check is performed. Currently, the only supported values are 60s (1 minute), 300s (5 minutes), 600s (10 minutes), and 900s (15 minutes). Optional, defaults to 60s.
	Period     *string `pulumi:"period"`
	ProjectsId string  `pulumi:"projectsId"`
	// The group resource associated with the configuration.
	ResourceGroup *ResourceGroup `pulumi:"resourceGroup"`
	// The list of regions from which the check will be run. Some regions contain one location, and others contain more than one. If this field is specified, enough regions must be provided to include a minimum of 3 locations. Not specifying this field will result in Uptime checks running from all available regions.
	SelectedRegions []string `pulumi:"selectedRegions"`
	// Contains information needed to make a TCP check.
	TcpCheck *TcpCheck `pulumi:"tcpCheck"`
	// The maximum amount of time to wait for the request to complete (must be between 1 and 60 seconds). Required.
	Timeout              *string `pulumi:"timeout"`
	UptimeCheckConfigsId string  `pulumi:"uptimeCheckConfigsId"`
}

// The set of arguments for constructing a UptimeCheckConfig resource.
type UptimeCheckConfigArgs struct {
	// The content that is expected to appear in the data returned by the target server against which the check is run. Currently, only the first entry in the content_matchers list is supported, and additional entries will be ignored. This field is optional and should only be specified if a content match is required as part of the/ Uptime check.
	ContentMatchers ContentMatcherArrayInput
	// A human-friendly name for the Uptime check configuration. The display name should be unique within a Stackdriver Workspace in order to make it easier to identify; however, uniqueness is not enforced. Required.
	DisplayName pulumi.StringPtrInput
	// Contains information needed to make an HTTP or HTTPS check.
	HttpCheck HttpCheckPtrInput
	// The internal checkers that this check will egress from. If is_internal is true and this list is empty, the check will egress from all the InternalCheckers configured for the project that owns this UptimeCheckConfig.
	InternalCheckers InternalCheckerArrayInput
	// If this is true, then checks are made only from the 'internal_checkers'. If it is false, then checks are made only from the 'selected_regions'. It is an error to provide 'selected_regions' when is_internal is true, or to provide 'internal_checkers' when is_internal is false.
	IsInternal pulumi.BoolPtrInput
	// The monitored resource (https://cloud.google.com/monitoring/api/resources) associated with the configuration. The following monitored resource types are valid for this field: uptime_url, gce_instance, gae_app, aws_ec2_instance, aws_elb_load_balancer
	MonitoredResource MonitoredResourcePtrInput
	// A unique resource name for this Uptime check configuration. The format is: projects/[PROJECT_ID_OR_NUMBER]/uptimeCheckConfigs/[UPTIME_CHECK_ID] [PROJECT_ID_OR_NUMBER] is the Workspace host project associated with the Uptime check.This field should be omitted when creating the Uptime check configuration; on create, the resource name is assigned by the server and included in the response.
	Name pulumi.StringPtrInput
	// How often, in seconds, the Uptime check is performed. Currently, the only supported values are 60s (1 minute), 300s (5 minutes), 600s (10 minutes), and 900s (15 minutes). Optional, defaults to 60s.
	Period     pulumi.StringPtrInput
	ProjectsId pulumi.StringInput
	// The group resource associated with the configuration.
	ResourceGroup ResourceGroupPtrInput
	// The list of regions from which the check will be run. Some regions contain one location, and others contain more than one. If this field is specified, enough regions must be provided to include a minimum of 3 locations. Not specifying this field will result in Uptime checks running from all available regions.
	SelectedRegions pulumi.StringArrayInput
	// Contains information needed to make a TCP check.
	TcpCheck TcpCheckPtrInput
	// The maximum amount of time to wait for the request to complete (must be between 1 and 60 seconds). Required.
	Timeout              pulumi.StringPtrInput
	UptimeCheckConfigsId pulumi.StringInput
}

func (UptimeCheckConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*uptimeCheckConfigArgs)(nil)).Elem()
}

type UptimeCheckConfigInput interface {
	pulumi.Input

	ToUptimeCheckConfigOutput() UptimeCheckConfigOutput
	ToUptimeCheckConfigOutputWithContext(ctx context.Context) UptimeCheckConfigOutput
}

func (*UptimeCheckConfig) ElementType() reflect.Type {
	return reflect.TypeOf((*UptimeCheckConfig)(nil))
}

func (i *UptimeCheckConfig) ToUptimeCheckConfigOutput() UptimeCheckConfigOutput {
	return i.ToUptimeCheckConfigOutputWithContext(context.Background())
}

func (i *UptimeCheckConfig) ToUptimeCheckConfigOutputWithContext(ctx context.Context) UptimeCheckConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UptimeCheckConfigOutput)
}

type UptimeCheckConfigOutput struct {
	*pulumi.OutputState
}

func (UptimeCheckConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UptimeCheckConfig)(nil))
}

func (o UptimeCheckConfigOutput) ToUptimeCheckConfigOutput() UptimeCheckConfigOutput {
	return o
}

func (o UptimeCheckConfigOutput) ToUptimeCheckConfigOutputWithContext(ctx context.Context) UptimeCheckConfigOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(UptimeCheckConfigOutput{})
}