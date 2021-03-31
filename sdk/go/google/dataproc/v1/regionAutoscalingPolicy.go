// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Creates new autoscaling policy.
type RegionAutoscalingPolicy struct {
	pulumi.CustomResourceState
}

// NewRegionAutoscalingPolicy registers a new resource with the given unique name, arguments, and options.
func NewRegionAutoscalingPolicy(ctx *pulumi.Context,
	name string, args *RegionAutoscalingPolicyArgs, opts ...pulumi.ResourceOption) (*RegionAutoscalingPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AutoscalingPoliciesId == nil {
		return nil, errors.New("invalid value for required argument 'AutoscalingPoliciesId'")
	}
	if args.ProjectsId == nil {
		return nil, errors.New("invalid value for required argument 'ProjectsId'")
	}
	if args.RegionsId == nil {
		return nil, errors.New("invalid value for required argument 'RegionsId'")
	}
	var resource RegionAutoscalingPolicy
	err := ctx.RegisterResource("google-cloud:dataproc/v1:RegionAutoscalingPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRegionAutoscalingPolicy gets an existing RegionAutoscalingPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRegionAutoscalingPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RegionAutoscalingPolicyState, opts ...pulumi.ResourceOption) (*RegionAutoscalingPolicy, error) {
	var resource RegionAutoscalingPolicy
	err := ctx.ReadResource("google-cloud:dataproc/v1:RegionAutoscalingPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RegionAutoscalingPolicy resources.
type regionAutoscalingPolicyState struct {
}

type RegionAutoscalingPolicyState struct {
}

func (RegionAutoscalingPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*regionAutoscalingPolicyState)(nil)).Elem()
}

type regionAutoscalingPolicyArgs struct {
	AutoscalingPoliciesId string                     `pulumi:"autoscalingPoliciesId"`
	BasicAlgorithm        *BasicAutoscalingAlgorithm `pulumi:"basicAlgorithm"`
	// Required. The policy id.The id must contain only letters (a-z, A-Z), numbers (0-9), underscores (_), and hyphens (-). Cannot begin or end with underscore or hyphen. Must consist of between 3 and 50 characters.
	Id *string `pulumi:"id"`
	// Output only. The "resource name" of the autoscaling policy, as described in https://cloud.google.com/apis/design/resource_names. For projects.regions.autoscalingPolicies, the resource name of the policy has the following format: projects/{project_id}/regions/{region}/autoscalingPolicies/{policy_id} For projects.locations.autoscalingPolicies, the resource name of the policy has the following format: projects/{project_id}/locations/{location}/autoscalingPolicies/{policy_id}
	Name       *string `pulumi:"name"`
	ProjectsId string  `pulumi:"projectsId"`
	RegionsId  string  `pulumi:"regionsId"`
	// Optional. Describes how the autoscaler will operate for secondary workers.
	SecondaryWorkerConfig *InstanceGroupAutoscalingPolicyConfig `pulumi:"secondaryWorkerConfig"`
	// Required. Describes how the autoscaler will operate for primary workers.
	WorkerConfig *InstanceGroupAutoscalingPolicyConfig `pulumi:"workerConfig"`
}

// The set of arguments for constructing a RegionAutoscalingPolicy resource.
type RegionAutoscalingPolicyArgs struct {
	AutoscalingPoliciesId pulumi.StringInput
	BasicAlgorithm        BasicAutoscalingAlgorithmPtrInput
	// Required. The policy id.The id must contain only letters (a-z, A-Z), numbers (0-9), underscores (_), and hyphens (-). Cannot begin or end with underscore or hyphen. Must consist of between 3 and 50 characters.
	Id pulumi.StringPtrInput
	// Output only. The "resource name" of the autoscaling policy, as described in https://cloud.google.com/apis/design/resource_names. For projects.regions.autoscalingPolicies, the resource name of the policy has the following format: projects/{project_id}/regions/{region}/autoscalingPolicies/{policy_id} For projects.locations.autoscalingPolicies, the resource name of the policy has the following format: projects/{project_id}/locations/{location}/autoscalingPolicies/{policy_id}
	Name       pulumi.StringPtrInput
	ProjectsId pulumi.StringInput
	RegionsId  pulumi.StringInput
	// Optional. Describes how the autoscaler will operate for secondary workers.
	SecondaryWorkerConfig InstanceGroupAutoscalingPolicyConfigPtrInput
	// Required. Describes how the autoscaler will operate for primary workers.
	WorkerConfig InstanceGroupAutoscalingPolicyConfigPtrInput
}

func (RegionAutoscalingPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*regionAutoscalingPolicyArgs)(nil)).Elem()
}

type RegionAutoscalingPolicyInput interface {
	pulumi.Input

	ToRegionAutoscalingPolicyOutput() RegionAutoscalingPolicyOutput
	ToRegionAutoscalingPolicyOutputWithContext(ctx context.Context) RegionAutoscalingPolicyOutput
}

func (*RegionAutoscalingPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((*RegionAutoscalingPolicy)(nil))
}

func (i *RegionAutoscalingPolicy) ToRegionAutoscalingPolicyOutput() RegionAutoscalingPolicyOutput {
	return i.ToRegionAutoscalingPolicyOutputWithContext(context.Background())
}

func (i *RegionAutoscalingPolicy) ToRegionAutoscalingPolicyOutputWithContext(ctx context.Context) RegionAutoscalingPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegionAutoscalingPolicyOutput)
}

type RegionAutoscalingPolicyOutput struct {
	*pulumi.OutputState
}

func (RegionAutoscalingPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RegionAutoscalingPolicy)(nil))
}

func (o RegionAutoscalingPolicyOutput) ToRegionAutoscalingPolicyOutput() RegionAutoscalingPolicyOutput {
	return o
}

func (o RegionAutoscalingPolicyOutput) ToRegionAutoscalingPolicyOutputWithContext(ctx context.Context) RegionAutoscalingPolicyOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(RegionAutoscalingPolicyOutput{})
}
