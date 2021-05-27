// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a debug session for a deployed API Proxy revision.
type OrganizationEnvironmentApiRevisionDebugsession struct {
	pulumi.CustomResourceState

	// Optional. The number of request to be traced. Min = 1, Max = 15, Default = 10.
	Count pulumi.IntOutput `pulumi:"count"`
	// Optional. A conditional statement which is evaluated against the request message to determine if it should be traced. Syntax matches that of on API Proxy bundle flow Condition.
	Filter pulumi.StringOutput `pulumi:"filter"`
	// A unique ID for this DebugSession.
	Name pulumi.StringOutput `pulumi:"name"`
	// Optional. The time in seconds after which this DebugSession should end. This value will override the value in query param, if both are provided.
	Timeout pulumi.StringOutput `pulumi:"timeout"`
	// Optional. The maximum number of bytes captured from the response payload. Min = 0, Max = 5120, Default = 5120.
	Tracesize pulumi.IntOutput `pulumi:"tracesize"`
	// Optional. The length of time, in seconds, that this debug session is valid, starting from when it's received in the control plane. Min = 1, Max = 15, Default = 10.
	Validity pulumi.IntOutput `pulumi:"validity"`
}

// NewOrganizationEnvironmentApiRevisionDebugsession registers a new resource with the given unique name, arguments, and options.
func NewOrganizationEnvironmentApiRevisionDebugsession(ctx *pulumi.Context,
	name string, args *OrganizationEnvironmentApiRevisionDebugsessionArgs, opts ...pulumi.ResourceOption) (*OrganizationEnvironmentApiRevisionDebugsession, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApiId == nil {
		return nil, errors.New("invalid value for required argument 'ApiId'")
	}
	if args.EnvironmentId == nil {
		return nil, errors.New("invalid value for required argument 'EnvironmentId'")
	}
	if args.OrganizationId == nil {
		return nil, errors.New("invalid value for required argument 'OrganizationId'")
	}
	if args.RevisionId == nil {
		return nil, errors.New("invalid value for required argument 'RevisionId'")
	}
	var resource OrganizationEnvironmentApiRevisionDebugsession
	err := ctx.RegisterResource("google-native:apigee/v1:OrganizationEnvironmentApiRevisionDebugsession", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOrganizationEnvironmentApiRevisionDebugsession gets an existing OrganizationEnvironmentApiRevisionDebugsession resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOrganizationEnvironmentApiRevisionDebugsession(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OrganizationEnvironmentApiRevisionDebugsessionState, opts ...pulumi.ResourceOption) (*OrganizationEnvironmentApiRevisionDebugsession, error) {
	var resource OrganizationEnvironmentApiRevisionDebugsession
	err := ctx.ReadResource("google-native:apigee/v1:OrganizationEnvironmentApiRevisionDebugsession", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering OrganizationEnvironmentApiRevisionDebugsession resources.
type organizationEnvironmentApiRevisionDebugsessionState struct {
	// Optional. The number of request to be traced. Min = 1, Max = 15, Default = 10.
	Count *int `pulumi:"count"`
	// Optional. A conditional statement which is evaluated against the request message to determine if it should be traced. Syntax matches that of on API Proxy bundle flow Condition.
	Filter *string `pulumi:"filter"`
	// A unique ID for this DebugSession.
	Name *string `pulumi:"name"`
	// Optional. The time in seconds after which this DebugSession should end. This value will override the value in query param, if both are provided.
	Timeout *string `pulumi:"timeout"`
	// Optional. The maximum number of bytes captured from the response payload. Min = 0, Max = 5120, Default = 5120.
	Tracesize *int `pulumi:"tracesize"`
	// Optional. The length of time, in seconds, that this debug session is valid, starting from when it's received in the control plane. Min = 1, Max = 15, Default = 10.
	Validity *int `pulumi:"validity"`
}

type OrganizationEnvironmentApiRevisionDebugsessionState struct {
	// Optional. The number of request to be traced. Min = 1, Max = 15, Default = 10.
	Count pulumi.IntPtrInput
	// Optional. A conditional statement which is evaluated against the request message to determine if it should be traced. Syntax matches that of on API Proxy bundle flow Condition.
	Filter pulumi.StringPtrInput
	// A unique ID for this DebugSession.
	Name pulumi.StringPtrInput
	// Optional. The time in seconds after which this DebugSession should end. This value will override the value in query param, if both are provided.
	Timeout pulumi.StringPtrInput
	// Optional. The maximum number of bytes captured from the response payload. Min = 0, Max = 5120, Default = 5120.
	Tracesize pulumi.IntPtrInput
	// Optional. The length of time, in seconds, that this debug session is valid, starting from when it's received in the control plane. Min = 1, Max = 15, Default = 10.
	Validity pulumi.IntPtrInput
}

func (OrganizationEnvironmentApiRevisionDebugsessionState) ElementType() reflect.Type {
	return reflect.TypeOf((*organizationEnvironmentApiRevisionDebugsessionState)(nil)).Elem()
}

type organizationEnvironmentApiRevisionDebugsessionArgs struct {
	ApiId string `pulumi:"apiId"`
	// Optional. The number of request to be traced. Min = 1, Max = 15, Default = 10.
	Count         *int   `pulumi:"count"`
	EnvironmentId string `pulumi:"environmentId"`
	// Optional. A conditional statement which is evaluated against the request message to determine if it should be traced. Syntax matches that of on API Proxy bundle flow Condition.
	Filter *string `pulumi:"filter"`
	// A unique ID for this DebugSession.
	Name           *string `pulumi:"name"`
	OrganizationId string  `pulumi:"organizationId"`
	RevisionId     string  `pulumi:"revisionId"`
	// Optional. The time in seconds after which this DebugSession should end. This value will override the value in query param, if both are provided.
	Timeout *string `pulumi:"timeout"`
	// Optional. The maximum number of bytes captured from the response payload. Min = 0, Max = 5120, Default = 5120.
	Tracesize *int `pulumi:"tracesize"`
	// Optional. The length of time, in seconds, that this debug session is valid, starting from when it's received in the control plane. Min = 1, Max = 15, Default = 10.
	Validity *int `pulumi:"validity"`
}

// The set of arguments for constructing a OrganizationEnvironmentApiRevisionDebugsession resource.
type OrganizationEnvironmentApiRevisionDebugsessionArgs struct {
	ApiId pulumi.StringInput
	// Optional. The number of request to be traced. Min = 1, Max = 15, Default = 10.
	Count         pulumi.IntPtrInput
	EnvironmentId pulumi.StringInput
	// Optional. A conditional statement which is evaluated against the request message to determine if it should be traced. Syntax matches that of on API Proxy bundle flow Condition.
	Filter pulumi.StringPtrInput
	// A unique ID for this DebugSession.
	Name           pulumi.StringPtrInput
	OrganizationId pulumi.StringInput
	RevisionId     pulumi.StringInput
	// Optional. The time in seconds after which this DebugSession should end. This value will override the value in query param, if both are provided.
	Timeout pulumi.StringPtrInput
	// Optional. The maximum number of bytes captured from the response payload. Min = 0, Max = 5120, Default = 5120.
	Tracesize pulumi.IntPtrInput
	// Optional. The length of time, in seconds, that this debug session is valid, starting from when it's received in the control plane. Min = 1, Max = 15, Default = 10.
	Validity pulumi.IntPtrInput
}

func (OrganizationEnvironmentApiRevisionDebugsessionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*organizationEnvironmentApiRevisionDebugsessionArgs)(nil)).Elem()
}

type OrganizationEnvironmentApiRevisionDebugsessionInput interface {
	pulumi.Input

	ToOrganizationEnvironmentApiRevisionDebugsessionOutput() OrganizationEnvironmentApiRevisionDebugsessionOutput
	ToOrganizationEnvironmentApiRevisionDebugsessionOutputWithContext(ctx context.Context) OrganizationEnvironmentApiRevisionDebugsessionOutput
}

func (*OrganizationEnvironmentApiRevisionDebugsession) ElementType() reflect.Type {
	return reflect.TypeOf((*OrganizationEnvironmentApiRevisionDebugsession)(nil))
}

func (i *OrganizationEnvironmentApiRevisionDebugsession) ToOrganizationEnvironmentApiRevisionDebugsessionOutput() OrganizationEnvironmentApiRevisionDebugsessionOutput {
	return i.ToOrganizationEnvironmentApiRevisionDebugsessionOutputWithContext(context.Background())
}

func (i *OrganizationEnvironmentApiRevisionDebugsession) ToOrganizationEnvironmentApiRevisionDebugsessionOutputWithContext(ctx context.Context) OrganizationEnvironmentApiRevisionDebugsessionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrganizationEnvironmentApiRevisionDebugsessionOutput)
}

type OrganizationEnvironmentApiRevisionDebugsessionOutput struct {
	*pulumi.OutputState
}

func (OrganizationEnvironmentApiRevisionDebugsessionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OrganizationEnvironmentApiRevisionDebugsession)(nil))
}

func (o OrganizationEnvironmentApiRevisionDebugsessionOutput) ToOrganizationEnvironmentApiRevisionDebugsessionOutput() OrganizationEnvironmentApiRevisionDebugsessionOutput {
	return o
}

func (o OrganizationEnvironmentApiRevisionDebugsessionOutput) ToOrganizationEnvironmentApiRevisionDebugsessionOutputWithContext(ctx context.Context) OrganizationEnvironmentApiRevisionDebugsessionOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(OrganizationEnvironmentApiRevisionDebugsessionOutput{})
}
