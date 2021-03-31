// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Create an OS Config patch deployment.
type PatchDeployment struct {
	pulumi.CustomResourceState
}

// NewPatchDeployment registers a new resource with the given unique name, arguments, and options.
func NewPatchDeployment(ctx *pulumi.Context,
	name string, args *PatchDeploymentArgs, opts ...pulumi.ResourceOption) (*PatchDeployment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PatchDeploymentsId == nil {
		return nil, errors.New("invalid value for required argument 'PatchDeploymentsId'")
	}
	if args.ProjectsId == nil {
		return nil, errors.New("invalid value for required argument 'ProjectsId'")
	}
	var resource PatchDeployment
	err := ctx.RegisterResource("google-cloud:osconfig/v1:PatchDeployment", name, args, &resource, opts...)
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
	err := ctx.ReadResource("google-cloud:osconfig/v1:PatchDeployment", name, id, state, &resource, opts...)
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
	// Output only. Time the patch deployment was created. Timestamp is in [RFC3339](https://www.ietf.org/rfc/rfc3339.txt) text format.
	CreateTime *string `pulumi:"createTime"`
	// Optional. Description of the patch deployment. Length of the description is limited to 1024 characters.
	Description *string `pulumi:"description"`
	// Optional. Duration of the patch. After the duration ends, the patch times out.
	Duration *string `pulumi:"duration"`
	// Required. VM instances to patch.
	InstanceFilter *PatchInstanceFilter `pulumi:"instanceFilter"`
	// Output only. The last time a patch job was started by this deployment. Timestamp is in [RFC3339](https://www.ietf.org/rfc/rfc3339.txt) text format.
	LastExecuteTime *string `pulumi:"lastExecuteTime"`
	// Unique name for the patch deployment resource in a project. The patch deployment name is in the form: `projects/{project_id}/patchDeployments/{patch_deployment_id}`. This field is ignored when you create a new patch deployment.
	Name *string `pulumi:"name"`
	// Required. Schedule a one-time execution.
	OneTimeSchedule *OneTimeSchedule `pulumi:"oneTimeSchedule"`
	// Optional. Patch configuration that is applied.
	PatchConfig        *PatchConfig `pulumi:"patchConfig"`
	PatchDeploymentsId string       `pulumi:"patchDeploymentsId"`
	ProjectsId         string       `pulumi:"projectsId"`
	// Required. Schedule recurring executions.
	RecurringSchedule *RecurringSchedule `pulumi:"recurringSchedule"`
	// Optional. Rollout strategy of the patch job.
	Rollout *PatchRollout `pulumi:"rollout"`
	// Output only. Time the patch deployment was last updated. Timestamp is in [RFC3339](https://www.ietf.org/rfc/rfc3339.txt) text format.
	UpdateTime *string `pulumi:"updateTime"`
}

// The set of arguments for constructing a PatchDeployment resource.
type PatchDeploymentArgs struct {
	// Output only. Time the patch deployment was created. Timestamp is in [RFC3339](https://www.ietf.org/rfc/rfc3339.txt) text format.
	CreateTime pulumi.StringPtrInput
	// Optional. Description of the patch deployment. Length of the description is limited to 1024 characters.
	Description pulumi.StringPtrInput
	// Optional. Duration of the patch. After the duration ends, the patch times out.
	Duration pulumi.StringPtrInput
	// Required. VM instances to patch.
	InstanceFilter PatchInstanceFilterPtrInput
	// Output only. The last time a patch job was started by this deployment. Timestamp is in [RFC3339](https://www.ietf.org/rfc/rfc3339.txt) text format.
	LastExecuteTime pulumi.StringPtrInput
	// Unique name for the patch deployment resource in a project. The patch deployment name is in the form: `projects/{project_id}/patchDeployments/{patch_deployment_id}`. This field is ignored when you create a new patch deployment.
	Name pulumi.StringPtrInput
	// Required. Schedule a one-time execution.
	OneTimeSchedule OneTimeSchedulePtrInput
	// Optional. Patch configuration that is applied.
	PatchConfig        PatchConfigPtrInput
	PatchDeploymentsId pulumi.StringInput
	ProjectsId         pulumi.StringInput
	// Required. Schedule recurring executions.
	RecurringSchedule RecurringSchedulePtrInput
	// Optional. Rollout strategy of the patch job.
	Rollout PatchRolloutPtrInput
	// Output only. Time the patch deployment was last updated. Timestamp is in [RFC3339](https://www.ietf.org/rfc/rfc3339.txt) text format.
	UpdateTime pulumi.StringPtrInput
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
	return reflect.TypeOf((*PatchDeployment)(nil))
}

func (i *PatchDeployment) ToPatchDeploymentOutput() PatchDeploymentOutput {
	return i.ToPatchDeploymentOutputWithContext(context.Background())
}

func (i *PatchDeployment) ToPatchDeploymentOutputWithContext(ctx context.Context) PatchDeploymentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PatchDeploymentOutput)
}

type PatchDeploymentOutput struct {
	*pulumi.OutputState
}

func (PatchDeploymentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PatchDeployment)(nil))
}

func (o PatchDeploymentOutput) ToPatchDeploymentOutput() PatchDeploymentOutput {
	return o
}

func (o PatchDeploymentOutput) ToPatchDeploymentOutputWithContext(ctx context.Context) PatchDeploymentOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(PatchDeploymentOutput{})
}
