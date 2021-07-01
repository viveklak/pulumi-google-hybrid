// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates new autoscaling policy.
type AutoscalingPolicy struct {
	pulumi.CustomResourceState

	BasicAlgorithm BasicAutoscalingAlgorithmResponseOutput `pulumi:"basicAlgorithm"`
	// The "resource name" of the autoscaling policy, as described in https://cloud.google.com/apis/design/resource_names. For projects.regions.autoscalingPolicies, the resource name of the policy has the following format: projects/{project_id}/regions/{region}/autoscalingPolicies/{policy_id} For projects.locations.autoscalingPolicies, the resource name of the policy has the following format: projects/{project_id}/locations/{location}/autoscalingPolicies/{policy_id}
	Name pulumi.StringOutput `pulumi:"name"`
	// Optional. Describes how the autoscaler will operate for secondary workers.
	SecondaryWorkerConfig InstanceGroupAutoscalingPolicyConfigResponseOutput `pulumi:"secondaryWorkerConfig"`
	// Required. Describes how the autoscaler will operate for primary workers.
	WorkerConfig InstanceGroupAutoscalingPolicyConfigResponseOutput `pulumi:"workerConfig"`
}

// NewAutoscalingPolicy registers a new resource with the given unique name, arguments, and options.
func NewAutoscalingPolicy(ctx *pulumi.Context,
	name string, args *AutoscalingPolicyArgs, opts ...pulumi.ResourceOption) (*AutoscalingPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Location == nil {
		return nil, errors.New("invalid value for required argument 'Location'")
	}
	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	var resource AutoscalingPolicy
	err := ctx.RegisterResource("google-native:dataproc/v1:AutoscalingPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAutoscalingPolicy gets an existing AutoscalingPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAutoscalingPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AutoscalingPolicyState, opts ...pulumi.ResourceOption) (*AutoscalingPolicy, error) {
	var resource AutoscalingPolicy
	err := ctx.ReadResource("google-native:dataproc/v1:AutoscalingPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AutoscalingPolicy resources.
type autoscalingPolicyState struct {
}

type AutoscalingPolicyState struct {
}

func (AutoscalingPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*autoscalingPolicyState)(nil)).Elem()
}

type autoscalingPolicyArgs struct {
	BasicAlgorithm *BasicAutoscalingAlgorithm `pulumi:"basicAlgorithm"`
	// Required. The policy id.The id must contain only letters (a-z, A-Z), numbers (0-9), underscores (_), and hyphens (-). Cannot begin or end with underscore or hyphen. Must consist of between 3 and 50 characters.
	Id       *string `pulumi:"id"`
	Location string  `pulumi:"location"`
	Project  string  `pulumi:"project"`
	// Optional. Describes how the autoscaler will operate for secondary workers.
	SecondaryWorkerConfig *InstanceGroupAutoscalingPolicyConfig `pulumi:"secondaryWorkerConfig"`
	// Required. Describes how the autoscaler will operate for primary workers.
	WorkerConfig *InstanceGroupAutoscalingPolicyConfig `pulumi:"workerConfig"`
}

// The set of arguments for constructing a AutoscalingPolicy resource.
type AutoscalingPolicyArgs struct {
	BasicAlgorithm BasicAutoscalingAlgorithmPtrInput
	// Required. The policy id.The id must contain only letters (a-z, A-Z), numbers (0-9), underscores (_), and hyphens (-). Cannot begin or end with underscore or hyphen. Must consist of between 3 and 50 characters.
	Id       pulumi.StringPtrInput
	Location pulumi.StringInput
	Project  pulumi.StringInput
	// Optional. Describes how the autoscaler will operate for secondary workers.
	SecondaryWorkerConfig InstanceGroupAutoscalingPolicyConfigPtrInput
	// Required. Describes how the autoscaler will operate for primary workers.
	WorkerConfig InstanceGroupAutoscalingPolicyConfigPtrInput
}

func (AutoscalingPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*autoscalingPolicyArgs)(nil)).Elem()
}

type AutoscalingPolicyInput interface {
	pulumi.Input

	ToAutoscalingPolicyOutput() AutoscalingPolicyOutput
	ToAutoscalingPolicyOutputWithContext(ctx context.Context) AutoscalingPolicyOutput
}

func (*AutoscalingPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoscalingPolicy)(nil))
}

func (i *AutoscalingPolicy) ToAutoscalingPolicyOutput() AutoscalingPolicyOutput {
	return i.ToAutoscalingPolicyOutputWithContext(context.Background())
}

func (i *AutoscalingPolicy) ToAutoscalingPolicyOutputWithContext(ctx context.Context) AutoscalingPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoscalingPolicyOutput)
}

type AutoscalingPolicyOutput struct {
	*pulumi.OutputState
}

func (AutoscalingPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoscalingPolicy)(nil))
}

func (o AutoscalingPolicyOutput) ToAutoscalingPolicyOutput() AutoscalingPolicyOutput {
	return o
}

func (o AutoscalingPolicyOutput) ToAutoscalingPolicyOutputWithContext(ctx context.Context) AutoscalingPolicyOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(AutoscalingPolicyOutput{})
}
