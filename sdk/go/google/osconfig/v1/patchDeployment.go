// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Create an OS Config patch deployment.
type PatchDeployment struct {
	pulumi.CustomResourceState

	// Time the patch deployment was created. Timestamp is in [RFC3339](https://www.ietf.org/rfc/rfc3339.txt) text format.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Optional. Description of the patch deployment. Length of the description is limited to 1024 characters.
	Description pulumi.StringOutput `pulumi:"description"`
	// Optional. Duration of the patch. After the duration ends, the patch times out.
	Duration pulumi.StringOutput `pulumi:"duration"`
	// VM instances to patch.
	InstanceFilter PatchInstanceFilterResponseOutput `pulumi:"instanceFilter"`
	// The last time a patch job was started by this deployment. Timestamp is in [RFC3339](https://www.ietf.org/rfc/rfc3339.txt) text format.
	LastExecuteTime pulumi.StringOutput `pulumi:"lastExecuteTime"`
	// Unique name for the patch deployment resource in a project. The patch deployment name is in the form: `projects/{project_id}/patchDeployments/{patch_deployment_id}`. This field is ignored when you create a new patch deployment.
	Name pulumi.StringOutput `pulumi:"name"`
	// Schedule a one-time execution.
	OneTimeSchedule OneTimeScheduleResponseOutput `pulumi:"oneTimeSchedule"`
	// Optional. Patch configuration that is applied.
	PatchConfig PatchConfigResponseOutput `pulumi:"patchConfig"`
	// Schedule recurring executions.
	RecurringSchedule RecurringScheduleResponseOutput `pulumi:"recurringSchedule"`
	// Optional. Rollout strategy of the patch job.
	Rollout PatchRolloutResponseOutput `pulumi:"rollout"`
	// Current state of the patch deployment.
	State pulumi.StringOutput `pulumi:"state"`
	// Time the patch deployment was last updated. Timestamp is in [RFC3339](https://www.ietf.org/rfc/rfc3339.txt) text format.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewPatchDeployment registers a new resource with the given unique name, arguments, and options.
func NewPatchDeployment(ctx *pulumi.Context,
	name string, args *PatchDeploymentArgs, opts ...pulumi.ResourceOption) (*PatchDeployment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceFilter == nil {
		return nil, errors.New("invalid value for required argument 'InstanceFilter'")
	}
	if args.OneTimeSchedule == nil {
		return nil, errors.New("invalid value for required argument 'OneTimeSchedule'")
	}
	if args.PatchDeploymentId == nil {
		return nil, errors.New("invalid value for required argument 'PatchDeploymentId'")
	}
	if args.RecurringSchedule == nil {
		return nil, errors.New("invalid value for required argument 'RecurringSchedule'")
	}
	var resource PatchDeployment
	err := ctx.RegisterResource("google-native:osconfig/v1:PatchDeployment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPatchDeployment gets an existing PatchDeployment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPatchDeployment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PatchDeploymentState, opts ...pulumi.ResourceOption) (*PatchDeployment, error) {
	var resource PatchDeployment
	err := ctx.ReadResource("google-native:osconfig/v1:PatchDeployment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PatchDeployment resources.
type patchDeploymentState struct {
}

type PatchDeploymentState struct {
}

func (PatchDeploymentState) ElementType() reflect.Type {
	return reflect.TypeOf((*patchDeploymentState)(nil)).Elem()
}

type patchDeploymentArgs struct {
	// Optional. Description of the patch deployment. Length of the description is limited to 1024 characters.
	Description *string `pulumi:"description"`
	// Optional. Duration of the patch. After the duration ends, the patch times out.
	Duration *string `pulumi:"duration"`
	// VM instances to patch.
	InstanceFilter PatchInstanceFilter `pulumi:"instanceFilter"`
	// Unique name for the patch deployment resource in a project. The patch deployment name is in the form: `projects/{project_id}/patchDeployments/{patch_deployment_id}`. This field is ignored when you create a new patch deployment.
	Name *string `pulumi:"name"`
	// Schedule a one-time execution.
	OneTimeSchedule OneTimeSchedule `pulumi:"oneTimeSchedule"`
	// Optional. Patch configuration that is applied.
	PatchConfig *PatchConfig `pulumi:"patchConfig"`
	// Required. A name for the patch deployment in the project. When creating a name the following rules apply: * Must contain only lowercase letters, numbers, and hyphens. * Must start with a letter. * Must be between 1-63 characters. * Must end with a number or a letter. * Must be unique within the project.
	PatchDeploymentId string  `pulumi:"patchDeploymentId"`
	Project           *string `pulumi:"project"`
	// Schedule recurring executions.
	RecurringSchedule RecurringSchedule `pulumi:"recurringSchedule"`
	// Optional. Rollout strategy of the patch job.
	Rollout *PatchRollout `pulumi:"rollout"`
}

// The set of arguments for constructing a PatchDeployment resource.
type PatchDeploymentArgs struct {
	// Optional. Description of the patch deployment. Length of the description is limited to 1024 characters.
	Description pulumi.StringPtrInput
	// Optional. Duration of the patch. After the duration ends, the patch times out.
	Duration pulumi.StringPtrInput
	// VM instances to patch.
	InstanceFilter PatchInstanceFilterInput
	// Unique name for the patch deployment resource in a project. The patch deployment name is in the form: `projects/{project_id}/patchDeployments/{patch_deployment_id}`. This field is ignored when you create a new patch deployment.
	Name pulumi.StringPtrInput
	// Schedule a one-time execution.
	OneTimeSchedule OneTimeScheduleInput
	// Optional. Patch configuration that is applied.
	PatchConfig PatchConfigPtrInput
	// Required. A name for the patch deployment in the project. When creating a name the following rules apply: * Must contain only lowercase letters, numbers, and hyphens. * Must start with a letter. * Must be between 1-63 characters. * Must end with a number or a letter. * Must be unique within the project.
	PatchDeploymentId pulumi.StringInput
	Project           pulumi.StringPtrInput
	// Schedule recurring executions.
	RecurringSchedule RecurringScheduleInput
	// Optional. Rollout strategy of the patch job.
	Rollout PatchRolloutPtrInput
}

func (PatchDeploymentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*patchDeploymentArgs)(nil)).Elem()
}

type PatchDeploymentInput interface {
	pulumi.Input

	ToPatchDeploymentOutput() PatchDeploymentOutput
	ToPatchDeploymentOutputWithContext(ctx context.Context) PatchDeploymentOutput
}

func (*PatchDeployment) ElementType() reflect.Type {
	return reflect.TypeOf((**PatchDeployment)(nil)).Elem()
}

func (i *PatchDeployment) ToPatchDeploymentOutput() PatchDeploymentOutput {
	return i.ToPatchDeploymentOutputWithContext(context.Background())
}

func (i *PatchDeployment) ToPatchDeploymentOutputWithContext(ctx context.Context) PatchDeploymentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PatchDeploymentOutput)
}

type PatchDeploymentOutput struct{ *pulumi.OutputState }

func (PatchDeploymentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PatchDeployment)(nil)).Elem()
}

func (o PatchDeploymentOutput) ToPatchDeploymentOutput() PatchDeploymentOutput {
	return o
}

func (o PatchDeploymentOutput) ToPatchDeploymentOutputWithContext(ctx context.Context) PatchDeploymentOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PatchDeploymentInput)(nil)).Elem(), &PatchDeployment{})
	pulumi.RegisterOutputType(PatchDeploymentOutput{})
}
