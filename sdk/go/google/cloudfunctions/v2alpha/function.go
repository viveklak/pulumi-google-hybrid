// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v2alpha

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a new function. If a function with the given name already exists in the specified project, the long running operation will return `ALREADY_EXISTS` error.
type Function struct {
	pulumi.CustomResourceState

	// Describes the Build step of the function that builds a container from the given source.
	BuildConfig BuildConfigResponseOutput `pulumi:"buildConfig"`
	// User-provided description of a function.
	Description pulumi.StringOutput `pulumi:"description"`
	// Describe whether the function is gen1 or gen2.
	Environment pulumi.StringOutput `pulumi:"environment"`
	// An Eventarc trigger managed by Google Cloud Functions that fires events in response to a condition in another service.
	EventTrigger EventTriggerResponseOutput `pulumi:"eventTrigger"`
	// Labels associated with this Cloud Function.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// A user-defined name of the function. Function names must be unique globally and match pattern `projects/*/locations/*/functions/*`
	Name pulumi.StringOutput `pulumi:"name"`
	// Describes the Service being deployed. Currently deploys services to Cloud Run (fully managed).
	ServiceConfig ServiceConfigResponseOutput `pulumi:"serviceConfig"`
	// State of the function.
	State pulumi.StringOutput `pulumi:"state"`
	// State Messages for this Cloud Function.
	StateMessages GoogleCloudFunctionsV2alphaStateMessageResponseArrayOutput `pulumi:"stateMessages"`
	// The last update timestamp of a Cloud Function.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewFunction registers a new resource with the given unique name, arguments, and options.
func NewFunction(ctx *pulumi.Context,
	name string, args *FunctionArgs, opts ...pulumi.ResourceOption) (*Function, error) {
	if args == nil {
		args = &FunctionArgs{}
	}

	var resource Function
	err := ctx.RegisterResource("google-native:cloudfunctions/v2alpha:Function", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFunction gets an existing Function resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFunction(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FunctionState, opts ...pulumi.ResourceOption) (*Function, error) {
	var resource Function
	err := ctx.ReadResource("google-native:cloudfunctions/v2alpha:Function", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Function resources.
type functionState struct {
}

type FunctionState struct {
}

func (FunctionState) ElementType() reflect.Type {
	return reflect.TypeOf((*functionState)(nil)).Elem()
}

type functionArgs struct {
	// Describes the Build step of the function that builds a container from the given source.
	BuildConfig *BuildConfig `pulumi:"buildConfig"`
	// User-provided description of a function.
	Description *string `pulumi:"description"`
	// Describe whether the function is gen1 or gen2.
	Environment *FunctionEnvironment `pulumi:"environment"`
	// An Eventarc trigger managed by Google Cloud Functions that fires events in response to a condition in another service.
	EventTrigger *EventTrigger `pulumi:"eventTrigger"`
	// The ID to use for the function, which will become the final component of the function's resource name. This value should be 4-63 characters, and valid characters are /a-z-/.
	FunctionId *string `pulumi:"functionId"`
	// Labels associated with this Cloud Function.
	Labels   map[string]string `pulumi:"labels"`
	Location *string           `pulumi:"location"`
	// A user-defined name of the function. Function names must be unique globally and match pattern `projects/*/locations/*/functions/*`
	Name    *string `pulumi:"name"`
	Project *string `pulumi:"project"`
	// Describes the Service being deployed. Currently deploys services to Cloud Run (fully managed).
	ServiceConfig *ServiceConfig `pulumi:"serviceConfig"`
}

// The set of arguments for constructing a Function resource.
type FunctionArgs struct {
	// Describes the Build step of the function that builds a container from the given source.
	BuildConfig BuildConfigPtrInput
	// User-provided description of a function.
	Description pulumi.StringPtrInput
	// Describe whether the function is gen1 or gen2.
	Environment FunctionEnvironmentPtrInput
	// An Eventarc trigger managed by Google Cloud Functions that fires events in response to a condition in another service.
	EventTrigger EventTriggerPtrInput
	// The ID to use for the function, which will become the final component of the function's resource name. This value should be 4-63 characters, and valid characters are /a-z-/.
	FunctionId pulumi.StringPtrInput
	// Labels associated with this Cloud Function.
	Labels   pulumi.StringMapInput
	Location pulumi.StringPtrInput
	// A user-defined name of the function. Function names must be unique globally and match pattern `projects/*/locations/*/functions/*`
	Name    pulumi.StringPtrInput
	Project pulumi.StringPtrInput
	// Describes the Service being deployed. Currently deploys services to Cloud Run (fully managed).
	ServiceConfig ServiceConfigPtrInput
}

func (FunctionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*functionArgs)(nil)).Elem()
}

type FunctionInput interface {
	pulumi.Input

	ToFunctionOutput() FunctionOutput
	ToFunctionOutputWithContext(ctx context.Context) FunctionOutput
}

func (*Function) ElementType() reflect.Type {
	return reflect.TypeOf((**Function)(nil)).Elem()
}

func (i *Function) ToFunctionOutput() FunctionOutput {
	return i.ToFunctionOutputWithContext(context.Background())
}

func (i *Function) ToFunctionOutputWithContext(ctx context.Context) FunctionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FunctionOutput)
}

type FunctionOutput struct{ *pulumi.OutputState }

func (FunctionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Function)(nil)).Elem()
}

func (o FunctionOutput) ToFunctionOutput() FunctionOutput {
	return o
}

func (o FunctionOutput) ToFunctionOutputWithContext(ctx context.Context) FunctionOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FunctionInput)(nil)).Elem(), &Function{})
	pulumi.RegisterOutputType(FunctionOutput{})
}
