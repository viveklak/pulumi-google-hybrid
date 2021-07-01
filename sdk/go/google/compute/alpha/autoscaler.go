// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package alpha

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates an autoscaler in the specified project using the data included in the request.
type Autoscaler struct {
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
	// Server-defined URL for this resource with the resource id.
	SelfLinkWithId pulumi.StringOutput `pulumi:"selfLinkWithId"`
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

// NewAutoscaler registers a new resource with the given unique name, arguments, and options.
func NewAutoscaler(ctx *pulumi.Context,
	name string, args *AutoscalerArgs, opts ...pulumi.ResourceOption) (*Autoscaler, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	if args.Zone == nil {
		return nil, errors.New("invalid value for required argument 'Zone'")
	}
	var resource Autoscaler
	err := ctx.RegisterResource("google-native:compute/alpha:Autoscaler", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAutoscaler gets an existing Autoscaler resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAutoscaler(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AutoscalerState, opts ...pulumi.ResourceOption) (*Autoscaler, error) {
	var resource Autoscaler
	err := ctx.ReadResource("google-native:compute/alpha:Autoscaler", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Autoscaler resources.
type autoscalerState struct {
}

type AutoscalerState struct {
}

func (AutoscalerState) ElementType() reflect.Type {
	return reflect.TypeOf((*autoscalerState)(nil)).Elem()
}

type autoscalerArgs struct {
	// The configuration parameters for the autoscaling algorithm. You can define one or more of the policies for an autoscaler: cpuUtilization, customMetricUtilizations, and loadBalancingUtilization.
	//
	// If none of these are specified, the default will be to autoscale based on cpuUtilization to 0.6 or 60%.
	AutoscalingPolicy *AutoscalingPolicy `pulumi:"autoscalingPolicy"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description *string `pulumi:"description"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name      *string `pulumi:"name"`
	Project   string  `pulumi:"project"`
	RequestId *string `pulumi:"requestId"`
	// URL of the managed instance group that this autoscaler will scale. This field is required when creating an autoscaler.
	Target *string `pulumi:"target"`
	Zone   string  `pulumi:"zone"`
}

// The set of arguments for constructing a Autoscaler resource.
type AutoscalerArgs struct {
	// The configuration parameters for the autoscaling algorithm. You can define one or more of the policies for an autoscaler: cpuUtilization, customMetricUtilizations, and loadBalancingUtilization.
	//
	// If none of these are specified, the default will be to autoscale based on cpuUtilization to 0.6 or 60%.
	AutoscalingPolicy AutoscalingPolicyPtrInput
	// An optional description of this resource. Provide this property when you create the resource.
	Description pulumi.StringPtrInput
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name      pulumi.StringPtrInput
	Project   pulumi.StringInput
	RequestId pulumi.StringPtrInput
	// URL of the managed instance group that this autoscaler will scale. This field is required when creating an autoscaler.
	Target pulumi.StringPtrInput
	Zone   pulumi.StringInput
}

func (AutoscalerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*autoscalerArgs)(nil)).Elem()
}

type AutoscalerInput interface {
	pulumi.Input

	ToAutoscalerOutput() AutoscalerOutput
	ToAutoscalerOutputWithContext(ctx context.Context) AutoscalerOutput
}

func (*Autoscaler) ElementType() reflect.Type {
	return reflect.TypeOf((*Autoscaler)(nil))
}

func (i *Autoscaler) ToAutoscalerOutput() AutoscalerOutput {
	return i.ToAutoscalerOutputWithContext(context.Background())
}

func (i *Autoscaler) ToAutoscalerOutputWithContext(ctx context.Context) AutoscalerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoscalerOutput)
}

type AutoscalerOutput struct {
	*pulumi.OutputState
}

func (AutoscalerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Autoscaler)(nil))
}

func (o AutoscalerOutput) ToAutoscalerOutput() AutoscalerOutput {
	return o
}

func (o AutoscalerOutput) ToAutoscalerOutputWithContext(ctx context.Context) AutoscalerOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(AutoscalerOutput{})
}
