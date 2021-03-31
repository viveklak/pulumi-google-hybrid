// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Sets the access control policy on the specified resource. Replaces any existing policy.Can return NOT_FOUND, INVALID_ARGUMENT, and PERMISSION_DENIED errors.
type RegionClusterIamPolicy struct {
	pulumi.CustomResourceState
}

// NewRegionClusterIamPolicy registers a new resource with the given unique name, arguments, and options.
func NewRegionClusterIamPolicy(ctx *pulumi.Context,
	name string, args *RegionClusterIamPolicyArgs, opts ...pulumi.ResourceOption) (*RegionClusterIamPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClustersId == nil {
		return nil, errors.New("invalid value for required argument 'ClustersId'")
	}
	if args.ProjectsId == nil {
		return nil, errors.New("invalid value for required argument 'ProjectsId'")
	}
	if args.RegionsId == nil {
		return nil, errors.New("invalid value for required argument 'RegionsId'")
	}
	var resource RegionClusterIamPolicy
	err := ctx.RegisterResource("google-cloud:dataproc/v1:RegionClusterIamPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRegionClusterIamPolicy gets an existing RegionClusterIamPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRegionClusterIamPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RegionClusterIamPolicyState, opts ...pulumi.ResourceOption) (*RegionClusterIamPolicy, error) {
	var resource RegionClusterIamPolicy
	err := ctx.ReadResource("google-cloud:dataproc/v1:RegionClusterIamPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RegionClusterIamPolicy resources.
type regionClusterIamPolicyState struct {
}

type RegionClusterIamPolicyState struct {
}

func (RegionClusterIamPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*regionClusterIamPolicyState)(nil)).Elem()
}

type regionClusterIamPolicyArgs struct {
	ClustersId string `pulumi:"clustersId"`
	// REQUIRED: The complete policy to be applied to the resource. The size of the policy is limited to a few 10s of KB. An empty policy is a valid policy but certain Cloud Platform services (such as Projects) might reject them.
	Policy     *Policy `pulumi:"policy"`
	ProjectsId string  `pulumi:"projectsId"`
	RegionsId  string  `pulumi:"regionsId"`
}

// The set of arguments for constructing a RegionClusterIamPolicy resource.
type RegionClusterIamPolicyArgs struct {
	ClustersId pulumi.StringInput
	// REQUIRED: The complete policy to be applied to the resource. The size of the policy is limited to a few 10s of KB. An empty policy is a valid policy but certain Cloud Platform services (such as Projects) might reject them.
	Policy     PolicyPtrInput
	ProjectsId pulumi.StringInput
	RegionsId  pulumi.StringInput
}

func (RegionClusterIamPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*regionClusterIamPolicyArgs)(nil)).Elem()
}

type RegionClusterIamPolicyInput interface {
	pulumi.Input

	ToRegionClusterIamPolicyOutput() RegionClusterIamPolicyOutput
	ToRegionClusterIamPolicyOutputWithContext(ctx context.Context) RegionClusterIamPolicyOutput
}

func (*RegionClusterIamPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((*RegionClusterIamPolicy)(nil))
}

func (i *RegionClusterIamPolicy) ToRegionClusterIamPolicyOutput() RegionClusterIamPolicyOutput {
	return i.ToRegionClusterIamPolicyOutputWithContext(context.Background())
}

func (i *RegionClusterIamPolicy) ToRegionClusterIamPolicyOutputWithContext(ctx context.Context) RegionClusterIamPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegionClusterIamPolicyOutput)
}

type RegionClusterIamPolicyOutput struct {
	*pulumi.OutputState
}

func (RegionClusterIamPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RegionClusterIamPolicy)(nil))
}

func (o RegionClusterIamPolicyOutput) ToRegionClusterIamPolicyOutput() RegionClusterIamPolicyOutput {
	return o
}

func (o RegionClusterIamPolicyOutput) ToRegionClusterIamPolicyOutputWithContext(ctx context.Context) RegionClusterIamPolicyOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(RegionClusterIamPolicyOutput{})
}
