// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v2beta1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates an agent version. The new version points to the agent instance in the "default" environment.
type Version struct {
	pulumi.CustomResourceState

	// The creation time of this version. This field is read-only, i.e., it cannot be set by create and update methods.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Optional. The developer-provided description of this version.
	Description pulumi.StringOutput `pulumi:"description"`
	// The unique identifier of this agent version. Supported formats: - `projects//agent/versions/` - `projects//locations//agent/versions/`
	Name pulumi.StringOutput `pulumi:"name"`
	// The status of this version. This field is read-only and cannot be set by create and update methods.
	Status pulumi.StringOutput `pulumi:"status"`
	// The sequential number of this version. This field is read-only which means it cannot be set by create and update methods.
	VersionNumber pulumi.IntOutput `pulumi:"versionNumber"`
}

// NewVersion registers a new resource with the given unique name, arguments, and options.
func NewVersion(ctx *pulumi.Context,
	name string, args *VersionArgs, opts ...pulumi.ResourceOption) (*Version, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Location == nil {
		return nil, errors.New("invalid value for required argument 'Location'")
	}
	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	var resource Version
	err := ctx.RegisterResource("google-native:dialogflow/v2beta1:Version", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVersion gets an existing Version resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVersion(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VersionState, opts ...pulumi.ResourceOption) (*Version, error) {
	var resource Version
	err := ctx.ReadResource("google-native:dialogflow/v2beta1:Version", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Version resources.
type versionState struct {
}

type VersionState struct {
}

func (VersionState) ElementType() reflect.Type {
	return reflect.TypeOf((*versionState)(nil)).Elem()
}

type versionArgs struct {
	// Optional. The developer-provided description of this version.
	Description *string `pulumi:"description"`
	Location    string  `pulumi:"location"`
	Project     string  `pulumi:"project"`
}

// The set of arguments for constructing a Version resource.
type VersionArgs struct {
	// Optional. The developer-provided description of this version.
	Description pulumi.StringPtrInput
	Location    pulumi.StringInput
	Project     pulumi.StringInput
}

func (VersionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*versionArgs)(nil)).Elem()
}

type VersionInput interface {
	pulumi.Input

	ToVersionOutput() VersionOutput
	ToVersionOutputWithContext(ctx context.Context) VersionOutput
}

func (*Version) ElementType() reflect.Type {
	return reflect.TypeOf((*Version)(nil))
}

func (i *Version) ToVersionOutput() VersionOutput {
	return i.ToVersionOutputWithContext(context.Background())
}

func (i *Version) ToVersionOutputWithContext(ctx context.Context) VersionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VersionOutput)
}

type VersionOutput struct {
	*pulumi.OutputState
}

func (VersionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Version)(nil))
}

func (o VersionOutput) ToVersionOutput() VersionOutput {
	return o
}

func (o VersionOutput) ToVersionOutputWithContext(ctx context.Context) VersionOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(VersionOutput{})
}
