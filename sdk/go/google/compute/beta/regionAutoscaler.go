// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package beta

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates an autoscaler in the specified project using the data included in the request.
type RegionAutoscaler struct {
	pulumi.CustomResourceState

	// The configuration parameters for the autoscaling algorithm. You can define one or more of the policies for an autoscaler: cpuUtilization, customMetricUtilizations, and loadBalancingUtilization.
	//
	// If none of these are specified, the default will be to autoscale based on cpuUtilization to 0.6 or 60%.
	AutoscalingPolicy AutoscalingPolicyResponseOutput `pulumi:"autoscalingPolicy"`
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringOutput `pulumi:"creationTimestamp"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description pulumi.StringOutput `pulumi:"description"`
	// Type of the resource. Always compute#autoscaler for autoscalers.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name pulumi.StringOutput `pulumi:"name"`
	// Target recommended MIG size (number of instances) computed by autoscaler. Autoscaler calculates the recommended MIG size even when the autoscaling policy mode is different from ON. This field is empty when autoscaler is not connected to an existing managed instance group or autoscaler did not generate its prediction.
	RecommendedSize pulumi.IntOutput `pulumi:"recommendedSize"`
	// URL of the region where the instance group resides (for autoscalers living in regional scope).
	Region pulumi.StringOutput `pulumi:"region"`
	// Status information of existing scaling schedules.
	ScalingScheduleStatus pulumi.StringMapOutput `pulumi:"scalingScheduleStatus"`
	// Server-defined URL for the resource.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
	// The status of the autoscaler configuration. Current set of possible values:
	// - PENDING: Autoscaler backend hasn't read new/updated configuration.
	// - DELETING: Configuration is being deleted.
	// - ACTIVE: Configuration is acknowledged to be effective. Some warnings might be present in the statusDetails field.
	// - ERROR: Configuration has errors. Actionable for users. Details are present in the statusDetails field.  New values might be added in the future.
	Status pulumi.StringOutput `pulumi:"status"`
	// Human-readable details about the current state of the autoscaler. Read the documentation for Commonly returned status messages for examples of status messages you might encounter.
	StatusDetails AutoscalerStatusDetailsResponseArrayOutput `pulumi:"statusDetails"`
	// URL of the managed instance group that this autoscaler will scale. This field is required when creating an autoscaler.
	Target pulumi.StringOutput `pulumi:"target"`
	// URL of the zone where the instance group resides (for autoscalers living in zonal scope).
	Zone pulumi.StringOutput `pulumi:"zone"`
}

// NewRegionAutoscaler registers a new resource with the given unique name, arguments, and options.
func NewRegionAutoscaler(ctx *pulumi.Context,
	name string, args *RegionAutoscalerArgs, opts ...pulumi.ResourceOption) (*RegionAutoscaler, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	if args.Region == nil {
		return nil, errors.New("invalid value for required argument 'Region'")
	}
	var resource RegionAutoscaler
	err := ctx.RegisterResource("google-native:compute/beta:RegionAutoscaler", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRegionAutoscaler gets an existing RegionAutoscaler resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRegionAutoscaler(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RegionAutoscalerState, opts ...pulumi.ResourceOption) (*RegionAutoscaler, error) {
	var resource RegionAutoscaler
	err := ctx.ReadResource("google-native:compute/beta:RegionAutoscaler", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RegionAutoscaler resources.
type regionAutoscalerState struct {
	// The configuration parameters for the autoscaling algorithm. You can define one or more of the policies for an autoscaler: cpuUtilization, customMetricUtilizations, and loadBalancingUtilization.
	//
	// If none of these are specified, the default will be to autoscale based on cpuUtilization to 0.6 or 60%.
	AutoscalingPolicy *AutoscalingPolicyResponse `pulumi:"autoscalingPolicy"`
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp *string `pulumi:"creationTimestamp"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description *string `pulumi:"description"`
	// Type of the resource. Always compute#autoscaler for autoscalers.
	Kind *string `pulumi:"kind"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name *string `pulumi:"name"`
	// Target recommended MIG size (number of instances) computed by autoscaler. Autoscaler calculates the recommended MIG size even when the autoscaling policy mode is different from ON. This field is empty when autoscaler is not connected to an existing managed instance group or autoscaler did not generate its prediction.
	RecommendedSize *int `pulumi:"recommendedSize"`
	// URL of the region where the instance group resides (for autoscalers living in regional scope).
	Region *string `pulumi:"region"`
	// Status information of existing scaling schedules.
	ScalingScheduleStatus map[string]string `pulumi:"scalingScheduleStatus"`
	// Server-defined URL for the resource.
	SelfLink *string `pulumi:"selfLink"`
	// The status of the autoscaler configuration. Current set of possible values:
	// - PENDING: Autoscaler backend hasn't read new/updated configuration.
	// - DELETING: Configuration is being deleted.
	// - ACTIVE: Configuration is acknowledged to be effective. Some warnings might be present in the statusDetails field.
	// - ERROR: Configuration has errors. Actionable for users. Details are present in the statusDetails field.  New values might be added in the future.
	Status *string `pulumi:"status"`
	// Human-readable details about the current state of the autoscaler. Read the documentation for Commonly returned status messages for examples of status messages you might encounter.
	StatusDetails []AutoscalerStatusDetailsResponse `pulumi:"statusDetails"`
	// URL of the managed instance group that this autoscaler will scale. This field is required when creating an autoscaler.
	Target *string `pulumi:"target"`
	// URL of the zone where the instance group resides (for autoscalers living in zonal scope).
	Zone *string `pulumi:"zone"`
}

type RegionAutoscalerState struct {
	// The configuration parameters for the autoscaling algorithm. You can define one or more of the policies for an autoscaler: cpuUtilization, customMetricUtilizations, and loadBalancingUtilization.
	//
	// If none of these are specified, the default will be to autoscale based on cpuUtilization to 0.6 or 60%.
	AutoscalingPolicy AutoscalingPolicyResponsePtrInput
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringPtrInput
	// An optional description of this resource. Provide this property when you create the resource.
	Description pulumi.StringPtrInput
	// Type of the resource. Always compute#autoscaler for autoscalers.
	Kind pulumi.StringPtrInput
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name pulumi.StringPtrInput
	// Target recommended MIG size (number of instances) computed by autoscaler. Autoscaler calculates the recommended MIG size even when the autoscaling policy mode is different from ON. This field is empty when autoscaler is not connected to an existing managed instance group or autoscaler did not generate its prediction.
	RecommendedSize pulumi.IntPtrInput
	// URL of the region where the instance group resides (for autoscalers living in regional scope).
	Region pulumi.StringPtrInput
	// Status information of existing scaling schedules.
	ScalingScheduleStatus pulumi.StringMapInput
	// Server-defined URL for the resource.
	SelfLink pulumi.StringPtrInput
	// The status of the autoscaler configuration. Current set of possible values:
	// - PENDING: Autoscaler backend hasn't read new/updated configuration.
	// - DELETING: Configuration is being deleted.
	// - ACTIVE: Configuration is acknowledged to be effective. Some warnings might be present in the statusDetails field.
	// - ERROR: Configuration has errors. Actionable for users. Details are present in the statusDetails field.  New values might be added in the future.
	Status pulumi.StringPtrInput
	// Human-readable details about the current state of the autoscaler. Read the documentation for Commonly returned status messages for examples of status messages you might encounter.
	StatusDetails AutoscalerStatusDetailsResponseArrayInput
	// URL of the managed instance group that this autoscaler will scale. This field is required when creating an autoscaler.
	Target pulumi.StringPtrInput
	// URL of the zone where the instance group resides (for autoscalers living in zonal scope).
	Zone pulumi.StringPtrInput
}

func (RegionAutoscalerState) ElementType() reflect.Type {
	return reflect.TypeOf((*regionAutoscalerState)(nil)).Elem()
}

type regionAutoscalerArgs struct {
	// The configuration parameters for the autoscaling algorithm. You can define one or more of the policies for an autoscaler: cpuUtilization, customMetricUtilizations, and loadBalancingUtilization.
	//
	// If none of these are specified, the default will be to autoscale based on cpuUtilization to 0.6 or 60%.
	AutoscalingPolicy *AutoscalingPolicy `pulumi:"autoscalingPolicy"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description *string `pulumi:"description"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name      *string `pulumi:"name"`
	Project   string  `pulumi:"project"`
	Region    string  `pulumi:"region"`
	RequestId *string `pulumi:"requestId"`
	// URL of the managed instance group that this autoscaler will scale. This field is required when creating an autoscaler.
	Target *string `pulumi:"target"`
}

// The set of arguments for constructing a RegionAutoscaler resource.
type RegionAutoscalerArgs struct {
	// The configuration parameters for the autoscaling algorithm. You can define one or more of the policies for an autoscaler: cpuUtilization, customMetricUtilizations, and loadBalancingUtilization.
	//
	// If none of these are specified, the default will be to autoscale based on cpuUtilization to 0.6 or 60%.
	AutoscalingPolicy AutoscalingPolicyPtrInput
	// An optional description of this resource. Provide this property when you create the resource.
	Description pulumi.StringPtrInput
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name      pulumi.StringPtrInput
	Project   pulumi.StringInput
	Region    pulumi.StringInput
	RequestId pulumi.StringPtrInput
	// URL of the managed instance group that this autoscaler will scale. This field is required when creating an autoscaler.
	Target pulumi.StringPtrInput
}

func (RegionAutoscalerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*regionAutoscalerArgs)(nil)).Elem()
}

type RegionAutoscalerInput interface {
	pulumi.Input

	ToRegionAutoscalerOutput() RegionAutoscalerOutput
	ToRegionAutoscalerOutputWithContext(ctx context.Context) RegionAutoscalerOutput
}

func (*RegionAutoscaler) ElementType() reflect.Type {
	return reflect.TypeOf((*RegionAutoscaler)(nil))
}

func (i *RegionAutoscaler) ToRegionAutoscalerOutput() RegionAutoscalerOutput {
	return i.ToRegionAutoscalerOutputWithContext(context.Background())
}

func (i *RegionAutoscaler) ToRegionAutoscalerOutputWithContext(ctx context.Context) RegionAutoscalerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegionAutoscalerOutput)
}

type RegionAutoscalerOutput struct {
	*pulumi.OutputState
}

func (RegionAutoscalerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RegionAutoscaler)(nil))
}

func (o RegionAutoscalerOutput) ToRegionAutoscalerOutput() RegionAutoscalerOutput {
	return o
}

func (o RegionAutoscalerOutput) ToRegionAutoscalerOutputWithContext(ctx context.Context) RegionAutoscalerOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(RegionAutoscalerOutput{})
}
