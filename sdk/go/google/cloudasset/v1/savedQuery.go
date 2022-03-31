// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a saved query in a parent project/folder/organization.
// Auto-naming is currently not supported for this resource.
type SavedQuery struct {
	pulumi.CustomResourceState

	// The query content.
	Content QueryContentResponseOutput `pulumi:"content"`
	// The create time of this saved query.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// The account's email address who has created this saved query.
	Creator pulumi.StringOutput `pulumi:"creator"`
	// The description of this saved query. This value should be fewer than 255 characters.
	Description pulumi.StringOutput `pulumi:"description"`
	// Labels applied on the resource. This value should not contain more than 10 entries. The key and value of each entry must be non-empty and fewer than 64 characters.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// The last update time of this saved query.
	LastUpdateTime pulumi.StringOutput `pulumi:"lastUpdateTime"`
	// The account's email address who has updated this saved query most recently.
	LastUpdater pulumi.StringOutput `pulumi:"lastUpdater"`
	// The resource name of the saved query. The format must be: * projects/project_number/savedQueries/saved_query_id * folders/folder_number/savedQueries/saved_query_id * organizations/organization_number/savedQueries/saved_query_id
	Name pulumi.StringOutput `pulumi:"name"`
}

// NewSavedQuery registers a new resource with the given unique name, arguments, and options.
func NewSavedQuery(ctx *pulumi.Context,
	name string, args *SavedQueryArgs, opts ...pulumi.ResourceOption) (*SavedQuery, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.SavedQueryId == nil {
		return nil, errors.New("invalid value for required argument 'SavedQueryId'")
	}
	if args.V1Id == nil {
		return nil, errors.New("invalid value for required argument 'V1Id'")
	}
	if args.V1Id1 == nil {
		return nil, errors.New("invalid value for required argument 'V1Id1'")
	}
	var resource SavedQuery
	err := ctx.RegisterResource("google-native:cloudasset/v1:SavedQuery", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSavedQuery gets an existing SavedQuery resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSavedQuery(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SavedQueryState, opts ...pulumi.ResourceOption) (*SavedQuery, error) {
	var resource SavedQuery
	err := ctx.ReadResource("google-native:cloudasset/v1:SavedQuery", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SavedQuery resources.
type savedQueryState struct {
}

type SavedQueryState struct {
}

func (SavedQueryState) ElementType() reflect.Type {
	return reflect.TypeOf((*savedQueryState)(nil)).Elem()
}

type savedQueryArgs struct {
	// The query content.
	Content *QueryContent `pulumi:"content"`
	// The description of this saved query. This value should be fewer than 255 characters.
	Description *string `pulumi:"description"`
	// Labels applied on the resource. This value should not contain more than 10 entries. The key and value of each entry must be non-empty and fewer than 64 characters.
	Labels map[string]string `pulumi:"labels"`
	// The resource name of the saved query. The format must be: * projects/project_number/savedQueries/saved_query_id * folders/folder_number/savedQueries/saved_query_id * organizations/organization_number/savedQueries/saved_query_id
	Name *string `pulumi:"name"`
	// Required. The ID to use for the saved query, which must be unique in the specified parent. It will become the final component of the saved query's resource name. This value should be 4-63 characters, and valid characters are /a-z-/. Notice that this field is required in the saved query creation, and the `name` field of the `saved_query` will be ignored.
	SavedQueryId string `pulumi:"savedQueryId"`
	V1Id         string `pulumi:"v1Id"`
	V1Id1        string `pulumi:"v1Id1"`
}

// The set of arguments for constructing a SavedQuery resource.
type SavedQueryArgs struct {
	// The query content.
	Content QueryContentPtrInput
	// The description of this saved query. This value should be fewer than 255 characters.
	Description pulumi.StringPtrInput
	// Labels applied on the resource. This value should not contain more than 10 entries. The key and value of each entry must be non-empty and fewer than 64 characters.
	Labels pulumi.StringMapInput
	// The resource name of the saved query. The format must be: * projects/project_number/savedQueries/saved_query_id * folders/folder_number/savedQueries/saved_query_id * organizations/organization_number/savedQueries/saved_query_id
	Name pulumi.StringPtrInput
	// Required. The ID to use for the saved query, which must be unique in the specified parent. It will become the final component of the saved query's resource name. This value should be 4-63 characters, and valid characters are /a-z-/. Notice that this field is required in the saved query creation, and the `name` field of the `saved_query` will be ignored.
	SavedQueryId pulumi.StringInput
	V1Id         pulumi.StringInput
	V1Id1        pulumi.StringInput
}

func (SavedQueryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*savedQueryArgs)(nil)).Elem()
}

type SavedQueryInput interface {
	pulumi.Input

	ToSavedQueryOutput() SavedQueryOutput
	ToSavedQueryOutputWithContext(ctx context.Context) SavedQueryOutput
}

func (*SavedQuery) ElementType() reflect.Type {
	return reflect.TypeOf((**SavedQuery)(nil)).Elem()
}

func (i *SavedQuery) ToSavedQueryOutput() SavedQueryOutput {
	return i.ToSavedQueryOutputWithContext(context.Background())
}

func (i *SavedQuery) ToSavedQueryOutputWithContext(ctx context.Context) SavedQueryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SavedQueryOutput)
}

type SavedQueryOutput struct{ *pulumi.OutputState }

func (SavedQueryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SavedQuery)(nil)).Elem()
}

func (o SavedQueryOutput) ToSavedQueryOutput() SavedQueryOutput {
	return o
}

func (o SavedQueryOutput) ToSavedQueryOutputWithContext(ctx context.Context) SavedQueryOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SavedQueryInput)(nil)).Elem(), &SavedQuery{})
	pulumi.RegisterOutputType(SavedQueryOutput{})
}
