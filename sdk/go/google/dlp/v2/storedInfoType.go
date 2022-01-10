// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a pre-built stored infoType to be used for inspection. See https://cloud.google.com/dlp/docs/creating-stored-infotypes to learn more.
// Auto-naming is currently not supported for this resource.
type StoredInfoType struct {
	pulumi.CustomResourceState

	// Current version of the stored info type.
	CurrentVersion GooglePrivacyDlpV2StoredInfoTypeVersionResponseOutput `pulumi:"currentVersion"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Pending versions of the stored info type. Empty if no versions are pending.
	PendingVersions GooglePrivacyDlpV2StoredInfoTypeVersionResponseArrayOutput `pulumi:"pendingVersions"`
}

// NewStoredInfoType registers a new resource with the given unique name, arguments, and options.
func NewStoredInfoType(ctx *pulumi.Context,
	name string, args *StoredInfoTypeArgs, opts ...pulumi.ResourceOption) (*StoredInfoType, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Config == nil {
		return nil, errors.New("invalid value for required argument 'Config'")
	}
	var resource StoredInfoType
	err := ctx.RegisterResource("google-native:dlp/v2:StoredInfoType", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetStoredInfoType gets an existing StoredInfoType resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetStoredInfoType(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StoredInfoTypeState, opts ...pulumi.ResourceOption) (*StoredInfoType, error) {
	var resource StoredInfoType
	err := ctx.ReadResource("google-native:dlp/v2:StoredInfoType", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering StoredInfoType resources.
type storedInfoTypeState struct {
}

type StoredInfoTypeState struct {
}

func (StoredInfoTypeState) ElementType() reflect.Type {
	return reflect.TypeOf((*storedInfoTypeState)(nil)).Elem()
}

type storedInfoTypeArgs struct {
	// Configuration of the storedInfoType to create.
	Config   GooglePrivacyDlpV2StoredInfoTypeConfig `pulumi:"config"`
	Location *string                                `pulumi:"location"`
	Project  *string                                `pulumi:"project"`
	// The storedInfoType ID can contain uppercase and lowercase letters, numbers, and hyphens; that is, it must match the regular expression: `[a-zA-Z\d-_]+`. The maximum length is 100 characters. Can be empty to allow the system to generate one.
	StoredInfoTypeId *string `pulumi:"storedInfoTypeId"`
}

// The set of arguments for constructing a StoredInfoType resource.
type StoredInfoTypeArgs struct {
	// Configuration of the storedInfoType to create.
	Config   GooglePrivacyDlpV2StoredInfoTypeConfigInput
	Location pulumi.StringPtrInput
	Project  pulumi.StringPtrInput
	// The storedInfoType ID can contain uppercase and lowercase letters, numbers, and hyphens; that is, it must match the regular expression: `[a-zA-Z\d-_]+`. The maximum length is 100 characters. Can be empty to allow the system to generate one.
	StoredInfoTypeId pulumi.StringPtrInput
}

func (StoredInfoTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*storedInfoTypeArgs)(nil)).Elem()
}

type StoredInfoTypeInput interface {
	pulumi.Input

	ToStoredInfoTypeOutput() StoredInfoTypeOutput
	ToStoredInfoTypeOutputWithContext(ctx context.Context) StoredInfoTypeOutput
}

func (*StoredInfoType) ElementType() reflect.Type {
	return reflect.TypeOf((**StoredInfoType)(nil)).Elem()
}

func (i *StoredInfoType) ToStoredInfoTypeOutput() StoredInfoTypeOutput {
	return i.ToStoredInfoTypeOutputWithContext(context.Background())
}

func (i *StoredInfoType) ToStoredInfoTypeOutputWithContext(ctx context.Context) StoredInfoTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StoredInfoTypeOutput)
}

type StoredInfoTypeOutput struct{ *pulumi.OutputState }

func (StoredInfoTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StoredInfoType)(nil)).Elem()
}

func (o StoredInfoTypeOutput) ToStoredInfoTypeOutput() StoredInfoTypeOutput {
	return o
}

func (o StoredInfoTypeOutput) ToStoredInfoTypeOutputWithContext(ctx context.Context) StoredInfoTypeOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*StoredInfoTypeInput)(nil)).Elem(), &StoredInfoType{})
	pulumi.RegisterOutputType(StoredInfoTypeOutput{})
}
