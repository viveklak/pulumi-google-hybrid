// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package alpha

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates an instance template in the specified project using the data that is included in the request. If you are creating a new template to update an existing instance group, your new instance template must use the same network or, if applicable, the same subnetwork as the original template.
type InstanceTemplate struct {
	pulumi.CustomResourceState

	// The creation timestamp for this instance template in RFC3339 text format.
	CreationTimestamp pulumi.StringOutput `pulumi:"creationTimestamp"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description pulumi.StringOutput `pulumi:"description"`
	// The resource type, which is always compute#instanceTemplate for instance templates.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Name of the resource; provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name pulumi.StringOutput `pulumi:"name"`
	// The instance properties for this instance template.
	Properties InstancePropertiesResponseOutput `pulumi:"properties"`
	// The URL for this instance template. The server defines this URL.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
	// Server-defined URL for this resource with the resource id.
	SelfLinkWithId pulumi.StringOutput `pulumi:"selfLinkWithId"`
	// The source instance used to create the template. You can provide this as a partial or full URL to the resource. For example, the following are valid values:
	// - https://www.googleapis.com/compute/v1/projects/project/zones/zone/instances/instance
	// - projects/project/zones/zone/instances/instance
	SourceInstance pulumi.StringOutput `pulumi:"sourceInstance"`
	// The source instance params to use to create this instance template.
	SourceInstanceParams SourceInstanceParamsResponseOutput `pulumi:"sourceInstanceParams"`
}

// NewInstanceTemplate registers a new resource with the given unique name, arguments, and options.
func NewInstanceTemplate(ctx *pulumi.Context,
	name string, args *InstanceTemplateArgs, opts ...pulumi.ResourceOption) (*InstanceTemplate, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	var resource InstanceTemplate
	err := ctx.RegisterResource("google-native:compute/alpha:InstanceTemplate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInstanceTemplate gets an existing InstanceTemplate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInstanceTemplate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InstanceTemplateState, opts ...pulumi.ResourceOption) (*InstanceTemplate, error) {
	var resource InstanceTemplate
	err := ctx.ReadResource("google-native:compute/alpha:InstanceTemplate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering InstanceTemplate resources.
type instanceTemplateState struct {
}

type InstanceTemplateState struct {
}

func (InstanceTemplateState) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceTemplateState)(nil)).Elem()
}

type instanceTemplateArgs struct {
	// An optional description of this resource. Provide this property when you create the resource.
	Description *string `pulumi:"description"`
	// Name of the resource; provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name    *string `pulumi:"name"`
	Project string  `pulumi:"project"`
	// The instance properties for this instance template.
	Properties *InstanceProperties `pulumi:"properties"`
	RequestId  *string             `pulumi:"requestId"`
	// The source instance used to create the template. You can provide this as a partial or full URL to the resource. For example, the following are valid values:
	// - https://www.googleapis.com/compute/v1/projects/project/zones/zone/instances/instance
	// - projects/project/zones/zone/instances/instance
	SourceInstance *string `pulumi:"sourceInstance"`
	// The source instance params to use to create this instance template.
	SourceInstanceParams *SourceInstanceParams `pulumi:"sourceInstanceParams"`
}

// The set of arguments for constructing a InstanceTemplate resource.
type InstanceTemplateArgs struct {
	// An optional description of this resource. Provide this property when you create the resource.
	Description pulumi.StringPtrInput
	// Name of the resource; provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name    pulumi.StringPtrInput
	Project pulumi.StringInput
	// The instance properties for this instance template.
	Properties InstancePropertiesPtrInput
	RequestId  pulumi.StringPtrInput
	// The source instance used to create the template. You can provide this as a partial or full URL to the resource. For example, the following are valid values:
	// - https://www.googleapis.com/compute/v1/projects/project/zones/zone/instances/instance
	// - projects/project/zones/zone/instances/instance
	SourceInstance pulumi.StringPtrInput
	// The source instance params to use to create this instance template.
	SourceInstanceParams SourceInstanceParamsPtrInput
}

func (InstanceTemplateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceTemplateArgs)(nil)).Elem()
}

type InstanceTemplateInput interface {
	pulumi.Input

	ToInstanceTemplateOutput() InstanceTemplateOutput
	ToInstanceTemplateOutputWithContext(ctx context.Context) InstanceTemplateOutput
}

func (*InstanceTemplate) ElementType() reflect.Type {
	return reflect.TypeOf((*InstanceTemplate)(nil))
}

func (i *InstanceTemplate) ToInstanceTemplateOutput() InstanceTemplateOutput {
	return i.ToInstanceTemplateOutputWithContext(context.Background())
}

func (i *InstanceTemplate) ToInstanceTemplateOutputWithContext(ctx context.Context) InstanceTemplateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceTemplateOutput)
}

type InstanceTemplateOutput struct {
	*pulumi.OutputState
}

func (InstanceTemplateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InstanceTemplate)(nil))
}

func (o InstanceTemplateOutput) ToInstanceTemplateOutput() InstanceTemplateOutput {
	return o
}

func (o InstanceTemplateOutput) ToInstanceTemplateOutputWithContext(ctx context.Context) InstanceTemplateOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(InstanceTemplateOutput{})
}
