// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Uploads a ZIP-formatted shared flow configuration bundle to an organization. If the shared flow already exists, this creates a new revision of it. If the shared flow does not exist, this creates it. Once imported, the shared flow revision must be deployed before it can be accessed at runtime. The size limit of a shared flow bundle is 15 MB.
type OrganizationSharedflow struct {
	pulumi.CustomResourceState

	// The id of the most recently created revision for this shared flow.
	LatestRevisionId pulumi.StringOutput `pulumi:"latestRevisionId"`
	// Metadata describing the shared flow.
	MetaData GoogleCloudApigeeV1EntityMetadataResponseOutput `pulumi:"metaData"`
	// The ID of the shared flow.
	Name pulumi.StringOutput `pulumi:"name"`
	// A list of revisions of this shared flow.
	Revision pulumi.StringArrayOutput `pulumi:"revision"`
}

// NewOrganizationSharedflow registers a new resource with the given unique name, arguments, and options.
func NewOrganizationSharedflow(ctx *pulumi.Context,
	name string, args *OrganizationSharedflowArgs, opts ...pulumi.ResourceOption) (*OrganizationSharedflow, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Action == nil {
		return nil, errors.New("invalid value for required argument 'Action'")
	}
	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.OrganizationId == nil {
		return nil, errors.New("invalid value for required argument 'OrganizationId'")
	}
	var resource OrganizationSharedflow
	err := ctx.RegisterResource("google-native:apigee/v1:OrganizationSharedflow", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOrganizationSharedflow gets an existing OrganizationSharedflow resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOrganizationSharedflow(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OrganizationSharedflowState, opts ...pulumi.ResourceOption) (*OrganizationSharedflow, error) {
	var resource OrganizationSharedflow
	err := ctx.ReadResource("google-native:apigee/v1:OrganizationSharedflow", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering OrganizationSharedflow resources.
type organizationSharedflowState struct {
	// The id of the most recently created revision for this shared flow.
	LatestRevisionId *string `pulumi:"latestRevisionId"`
	// Metadata describing the shared flow.
	MetaData *GoogleCloudApigeeV1EntityMetadataResponse `pulumi:"metaData"`
	// The ID of the shared flow.
	Name *string `pulumi:"name"`
	// A list of revisions of this shared flow.
	Revision []string `pulumi:"revision"`
}

type OrganizationSharedflowState struct {
	// The id of the most recently created revision for this shared flow.
	LatestRevisionId pulumi.StringPtrInput
	// Metadata describing the shared flow.
	MetaData GoogleCloudApigeeV1EntityMetadataResponsePtrInput
	// The ID of the shared flow.
	Name pulumi.StringPtrInput
	// A list of revisions of this shared flow.
	Revision pulumi.StringArrayInput
}

func (OrganizationSharedflowState) ElementType() reflect.Type {
	return reflect.TypeOf((*organizationSharedflowState)(nil)).Elem()
}

type organizationSharedflowArgs struct {
	Action string `pulumi:"action"`
	// The HTTP Content-Type header value specifying the content type of the body.
	ContentType *string `pulumi:"contentType"`
	// The HTTP request/response body as raw binary.
	Data *string `pulumi:"data"`
	// Application specific response metadata. Must be set in the first response for streaming APIs.
	Extensions     []map[string]string `pulumi:"extensions"`
	Name           string              `pulumi:"name"`
	OrganizationId string              `pulumi:"organizationId"`
}

// The set of arguments for constructing a OrganizationSharedflow resource.
type OrganizationSharedflowArgs struct {
	Action pulumi.StringInput
	// The HTTP Content-Type header value specifying the content type of the body.
	ContentType pulumi.StringPtrInput
	// The HTTP request/response body as raw binary.
	Data pulumi.StringPtrInput
	// Application specific response metadata. Must be set in the first response for streaming APIs.
	Extensions     pulumi.StringMapArrayInput
	Name           pulumi.StringInput
	OrganizationId pulumi.StringInput
}

func (OrganizationSharedflowArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*organizationSharedflowArgs)(nil)).Elem()
}

type OrganizationSharedflowInput interface {
	pulumi.Input

	ToOrganizationSharedflowOutput() OrganizationSharedflowOutput
	ToOrganizationSharedflowOutputWithContext(ctx context.Context) OrganizationSharedflowOutput
}

func (*OrganizationSharedflow) ElementType() reflect.Type {
	return reflect.TypeOf((*OrganizationSharedflow)(nil))
}

func (i *OrganizationSharedflow) ToOrganizationSharedflowOutput() OrganizationSharedflowOutput {
	return i.ToOrganizationSharedflowOutputWithContext(context.Background())
}

func (i *OrganizationSharedflow) ToOrganizationSharedflowOutputWithContext(ctx context.Context) OrganizationSharedflowOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrganizationSharedflowOutput)
}

type OrganizationSharedflowOutput struct {
	*pulumi.OutputState
}

func (OrganizationSharedflowOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OrganizationSharedflow)(nil))
}

func (o OrganizationSharedflowOutput) ToOrganizationSharedflowOutput() OrganizationSharedflowOutput {
	return o
}

func (o OrganizationSharedflowOutput) ToOrganizationSharedflowOutputWithContext(ctx context.Context) OrganizationSharedflowOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(OrganizationSharedflowOutput{})
}
