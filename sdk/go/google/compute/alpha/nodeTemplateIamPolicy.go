// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package alpha

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Sets the access control policy on the specified resource. Replaces any existing policy.
type NodeTemplateIamPolicy struct {
	pulumi.CustomResourceState
}

// NewNodeTemplateIamPolicy registers a new resource with the given unique name, arguments, and options.
func NewNodeTemplateIamPolicy(ctx *pulumi.Context,
	name string, args *NodeTemplateIamPolicyArgs, opts ...pulumi.ResourceOption) (*NodeTemplateIamPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	if args.Region == nil {
		return nil, errors.New("invalid value for required argument 'Region'")
	}
	if args.Resource == nil {
		return nil, errors.New("invalid value for required argument 'Resource'")
	}
	var resource NodeTemplateIamPolicy
	err := ctx.RegisterResource("google-cloud:compute/alpha:NodeTemplateIamPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNodeTemplateIamPolicy gets an existing NodeTemplateIamPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNodeTemplateIamPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NodeTemplateIamPolicyState, opts ...pulumi.ResourceOption) (*NodeTemplateIamPolicy, error) {
	var resource NodeTemplateIamPolicy
	err := ctx.ReadResource("google-cloud:compute/alpha:NodeTemplateIamPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NodeTemplateIamPolicy resources.
type nodeTemplateIamPolicyState struct {
}

type NodeTemplateIamPolicyState struct {
}

func (NodeTemplateIamPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*nodeTemplateIamPolicyState)(nil)).Elem()
}

type nodeTemplateIamPolicyArgs struct {
	// Flatten Policy to create a backwacd compatible wire-format. Deprecated. Use 'policy' to specify bindings.
	Bindings []Binding `pulumi:"bindings"`
	// Flatten Policy to create a backward compatible wire-format. Deprecated. Use 'policy' to specify the etag.
	Etag *string `pulumi:"etag"`
	// REQUIRED: The complete policy to be applied to the 'resource'. The size of the policy is limited to a few 10s of KB. An empty policy is in general a valid policy but certain services (like Projects) might reject them.
	Policy   *Policy `pulumi:"policy"`
	Project  string  `pulumi:"project"`
	Region   string  `pulumi:"region"`
	Resource string  `pulumi:"resource"`
}

// The set of arguments for constructing a NodeTemplateIamPolicy resource.
type NodeTemplateIamPolicyArgs struct {
	// Flatten Policy to create a backwacd compatible wire-format. Deprecated. Use 'policy' to specify bindings.
	Bindings BindingArrayInput
	// Flatten Policy to create a backward compatible wire-format. Deprecated. Use 'policy' to specify the etag.
	Etag pulumi.StringPtrInput
	// REQUIRED: The complete policy to be applied to the 'resource'. The size of the policy is limited to a few 10s of KB. An empty policy is in general a valid policy but certain services (like Projects) might reject them.
	Policy   PolicyPtrInput
	Project  pulumi.StringInput
	Region   pulumi.StringInput
	Resource pulumi.StringInput
}

func (NodeTemplateIamPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*nodeTemplateIamPolicyArgs)(nil)).Elem()
}

type NodeTemplateIamPolicyInput interface {
	pulumi.Input

	ToNodeTemplateIamPolicyOutput() NodeTemplateIamPolicyOutput
	ToNodeTemplateIamPolicyOutputWithContext(ctx context.Context) NodeTemplateIamPolicyOutput
}

func (*NodeTemplateIamPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((*NodeTemplateIamPolicy)(nil))
}

func (i *NodeTemplateIamPolicy) ToNodeTemplateIamPolicyOutput() NodeTemplateIamPolicyOutput {
	return i.ToNodeTemplateIamPolicyOutputWithContext(context.Background())
}

func (i *NodeTemplateIamPolicy) ToNodeTemplateIamPolicyOutputWithContext(ctx context.Context) NodeTemplateIamPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NodeTemplateIamPolicyOutput)
}

type NodeTemplateIamPolicyOutput struct {
	*pulumi.OutputState
}

func (NodeTemplateIamPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NodeTemplateIamPolicy)(nil))
}

func (o NodeTemplateIamPolicyOutput) ToNodeTemplateIamPolicyOutput() NodeTemplateIamPolicyOutput {
	return o
}

func (o NodeTemplateIamPolicyOutput) ToNodeTemplateIamPolicyOutputWithContext(ctx context.Context) NodeTemplateIamPolicyOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(NodeTemplateIamPolicyOutput{})
}
