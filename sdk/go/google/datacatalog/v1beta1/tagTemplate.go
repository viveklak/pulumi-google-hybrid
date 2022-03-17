// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a tag template. The user should enable the Data Catalog API in the project identified by the `parent` parameter (see [Data Catalog Resource Project](https://cloud.google.com/data-catalog/docs/concepts/resource-project) for more information).
type TagTemplate struct {
	pulumi.CustomResourceState

	// The display name for this template. Defaults to an empty string.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// Map of tag template field IDs to the settings for the field. This map is an exhaustive list of the allowed fields. This map must contain at least one field and at most 500 fields. The keys to this map are tag template field IDs. Field IDs can contain letters (both uppercase and lowercase), numbers (0-9) and underscores (_). Field IDs must be at least 1 character long and at most 64 characters long. Field IDs must start with a letter or underscore.
	Fields pulumi.StringMapOutput `pulumi:"fields"`
	// The resource name of the tag template in URL format. Example: * projects/{project_id}/locations/{location}/tagTemplates/{tag_template_id} Note that this TagTemplate and its child resources may not actually be stored in the location in this name.
	Name pulumi.StringOutput `pulumi:"name"`
}

// NewTagTemplate registers a new resource with the given unique name, arguments, and options.
func NewTagTemplate(ctx *pulumi.Context,
	name string, args *TagTemplateArgs, opts ...pulumi.ResourceOption) (*TagTemplate, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Fields == nil {
		return nil, errors.New("invalid value for required argument 'Fields'")
	}
	if args.TagTemplateId == nil {
		return nil, errors.New("invalid value for required argument 'TagTemplateId'")
	}
	var resource TagTemplate
	err := ctx.RegisterResource("google-native:datacatalog/v1beta1:TagTemplate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTagTemplate gets an existing TagTemplate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTagTemplate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TagTemplateState, opts ...pulumi.ResourceOption) (*TagTemplate, error) {
	var resource TagTemplate
	err := ctx.ReadResource("google-native:datacatalog/v1beta1:TagTemplate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TagTemplate resources.
type tagTemplateState struct {
}

type TagTemplateState struct {
}

func (TagTemplateState) ElementType() reflect.Type {
	return reflect.TypeOf((*tagTemplateState)(nil)).Elem()
}

type tagTemplateArgs struct {
	// The display name for this template. Defaults to an empty string.
	DisplayName *string `pulumi:"displayName"`
	// Map of tag template field IDs to the settings for the field. This map is an exhaustive list of the allowed fields. This map must contain at least one field and at most 500 fields. The keys to this map are tag template field IDs. Field IDs can contain letters (both uppercase and lowercase), numbers (0-9) and underscores (_). Field IDs must be at least 1 character long and at most 64 characters long. Field IDs must start with a letter or underscore.
	Fields   map[string]string `pulumi:"fields"`
	Location *string           `pulumi:"location"`
	// The resource name of the tag template in URL format. Example: * projects/{project_id}/locations/{location}/tagTemplates/{tag_template_id} Note that this TagTemplate and its child resources may not actually be stored in the location in this name.
	Name    *string `pulumi:"name"`
	Project *string `pulumi:"project"`
	// Required. The id of the tag template to create.
	TagTemplateId string `pulumi:"tagTemplateId"`
}

// The set of arguments for constructing a TagTemplate resource.
type TagTemplateArgs struct {
	// The display name for this template. Defaults to an empty string.
	DisplayName pulumi.StringPtrInput
	// Map of tag template field IDs to the settings for the field. This map is an exhaustive list of the allowed fields. This map must contain at least one field and at most 500 fields. The keys to this map are tag template field IDs. Field IDs can contain letters (both uppercase and lowercase), numbers (0-9) and underscores (_). Field IDs must be at least 1 character long and at most 64 characters long. Field IDs must start with a letter or underscore.
	Fields   pulumi.StringMapInput
	Location pulumi.StringPtrInput
	// The resource name of the tag template in URL format. Example: * projects/{project_id}/locations/{location}/tagTemplates/{tag_template_id} Note that this TagTemplate and its child resources may not actually be stored in the location in this name.
	Name    pulumi.StringPtrInput
	Project pulumi.StringPtrInput
	// Required. The id of the tag template to create.
	TagTemplateId pulumi.StringInput
}

func (TagTemplateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*tagTemplateArgs)(nil)).Elem()
}

type TagTemplateInput interface {
	pulumi.Input

	ToTagTemplateOutput() TagTemplateOutput
	ToTagTemplateOutputWithContext(ctx context.Context) TagTemplateOutput
}

func (*TagTemplate) ElementType() reflect.Type {
	return reflect.TypeOf((**TagTemplate)(nil)).Elem()
}

func (i *TagTemplate) ToTagTemplateOutput() TagTemplateOutput {
	return i.ToTagTemplateOutputWithContext(context.Background())
}

func (i *TagTemplate) ToTagTemplateOutputWithContext(ctx context.Context) TagTemplateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TagTemplateOutput)
}

type TagTemplateOutput struct{ *pulumi.OutputState }

func (TagTemplateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TagTemplate)(nil)).Elem()
}

func (o TagTemplateOutput) ToTagTemplateOutput() TagTemplateOutput {
	return o
}

func (o TagTemplateOutput) ToTagTemplateOutputWithContext(ctx context.Context) TagTemplateOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TagTemplateInput)(nil)).Elem(), &TagTemplate{})
	pulumi.RegisterOutputType(TagTemplateOutput{})
}
