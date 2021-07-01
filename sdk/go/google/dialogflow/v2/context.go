// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a context. If the specified context already exists, overrides the context.
type Context struct {
	pulumi.CustomResourceState

	// Optional. The number of conversational query requests after which the context expires. The default is `0`. If set to `0`, the context expires immediately. Contexts expire automatically after 20 minutes if there are no matching queries.
	LifespanCount pulumi.IntOutput `pulumi:"lifespanCount"`
	// Required. The unique identifier of the context. Format: `projects//agent/sessions//contexts/`, or `projects//agent/environments//users//sessions//contexts/`. The `Context ID` is always converted to lowercase, may only contain characters in a-zA-Z0-9_-% and may be at most 250 bytes long. If `Environment ID` is not specified, we assume default 'draft' environment. If `User ID` is not specified, we assume default '-' user. The following context names are reserved for internal use by Dialogflow. You should not use these contexts or create contexts with these names: * `__system_counters__` * `*_id_dialog_context` * `*_dialog_params_size`
	Name pulumi.StringOutput `pulumi:"name"`
	// Optional. The collection of parameters associated with this context. Depending on your protocol or client library language, this is a map, associative array, symbol table, dictionary, or JSON object composed of a collection of (MapKey, MapValue) pairs: - MapKey type: string - MapKey value: parameter name - MapValue type: - If parameter's entity type is a composite entity: map - Else: depending on parameter value type, could be one of string, number, boolean, null, list or map - MapValue value: - If parameter's entity type is a composite entity: map from composite entity property names to property values - Else: parameter value
	Parameters pulumi.StringMapOutput `pulumi:"parameters"`
}

// NewContext registers a new resource with the given unique name, arguments, and options.
func NewContext(ctx *pulumi.Context,
	name string, args *ContextArgs, opts ...pulumi.ResourceOption) (*Context, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EnvironmentId == nil {
		return nil, errors.New("invalid value for required argument 'EnvironmentId'")
	}
	if args.Location == nil {
		return nil, errors.New("invalid value for required argument 'Location'")
	}
	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	if args.SessionId == nil {
		return nil, errors.New("invalid value for required argument 'SessionId'")
	}
	if args.UserId == nil {
		return nil, errors.New("invalid value for required argument 'UserId'")
	}
	var resource Context
	err := ctx.RegisterResource("google-native:dialogflow/v2:Context", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetContext gets an existing Context resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetContext(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ContextState, opts ...pulumi.ResourceOption) (*Context, error) {
	var resource Context
	err := ctx.ReadResource("google-native:dialogflow/v2:Context", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Context resources.
type contextState struct {
}

type ContextState struct {
}

func (ContextState) ElementType() reflect.Type {
	return reflect.TypeOf((*contextState)(nil)).Elem()
}

type contextArgs struct {
	EnvironmentId string `pulumi:"environmentId"`
	// Optional. The number of conversational query requests after which the context expires. The default is `0`. If set to `0`, the context expires immediately. Contexts expire automatically after 20 minutes if there are no matching queries.
	LifespanCount *int   `pulumi:"lifespanCount"`
	Location      string `pulumi:"location"`
	// Required. The unique identifier of the context. Format: `projects//agent/sessions//contexts/`, or `projects//agent/environments//users//sessions//contexts/`. The `Context ID` is always converted to lowercase, may only contain characters in a-zA-Z0-9_-% and may be at most 250 bytes long. If `Environment ID` is not specified, we assume default 'draft' environment. If `User ID` is not specified, we assume default '-' user. The following context names are reserved for internal use by Dialogflow. You should not use these contexts or create contexts with these names: * `__system_counters__` * `*_id_dialog_context` * `*_dialog_params_size`
	Name *string `pulumi:"name"`
	// Optional. The collection of parameters associated with this context. Depending on your protocol or client library language, this is a map, associative array, symbol table, dictionary, or JSON object composed of a collection of (MapKey, MapValue) pairs: - MapKey type: string - MapKey value: parameter name - MapValue type: - If parameter's entity type is a composite entity: map - Else: depending on parameter value type, could be one of string, number, boolean, null, list or map - MapValue value: - If parameter's entity type is a composite entity: map from composite entity property names to property values - Else: parameter value
	Parameters map[string]string `pulumi:"parameters"`
	Project    string            `pulumi:"project"`
	SessionId  string            `pulumi:"sessionId"`
	UserId     string            `pulumi:"userId"`
}

// The set of arguments for constructing a Context resource.
type ContextArgs struct {
	EnvironmentId pulumi.StringInput
	// Optional. The number of conversational query requests after which the context expires. The default is `0`. If set to `0`, the context expires immediately. Contexts expire automatically after 20 minutes if there are no matching queries.
	LifespanCount pulumi.IntPtrInput
	Location      pulumi.StringInput
	// Required. The unique identifier of the context. Format: `projects//agent/sessions//contexts/`, or `projects//agent/environments//users//sessions//contexts/`. The `Context ID` is always converted to lowercase, may only contain characters in a-zA-Z0-9_-% and may be at most 250 bytes long. If `Environment ID` is not specified, we assume default 'draft' environment. If `User ID` is not specified, we assume default '-' user. The following context names are reserved for internal use by Dialogflow. You should not use these contexts or create contexts with these names: * `__system_counters__` * `*_id_dialog_context` * `*_dialog_params_size`
	Name pulumi.StringPtrInput
	// Optional. The collection of parameters associated with this context. Depending on your protocol or client library language, this is a map, associative array, symbol table, dictionary, or JSON object composed of a collection of (MapKey, MapValue) pairs: - MapKey type: string - MapKey value: parameter name - MapValue type: - If parameter's entity type is a composite entity: map - Else: depending on parameter value type, could be one of string, number, boolean, null, list or map - MapValue value: - If parameter's entity type is a composite entity: map from composite entity property names to property values - Else: parameter value
	Parameters pulumi.StringMapInput
	Project    pulumi.StringInput
	SessionId  pulumi.StringInput
	UserId     pulumi.StringInput
}

func (ContextArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*contextArgs)(nil)).Elem()
}

type ContextInput interface {
	pulumi.Input

	ToContextOutput() ContextOutput
	ToContextOutputWithContext(ctx context.Context) ContextOutput
}

func (*Context) ElementType() reflect.Type {
	return reflect.TypeOf((*Context)(nil))
}

func (i *Context) ToContextOutput() ContextOutput {
	return i.ToContextOutputWithContext(context.Background())
}

func (i *Context) ToContextOutputWithContext(ctx context.Context) ContextOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContextOutput)
}

type ContextOutput struct {
	*pulumi.OutputState
}

func (ContextOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Context)(nil))
}

func (o ContextOutput) ToContextOutput() ContextOutput {
	return o
}

func (o ContextOutput) ToContextOutputWithContext(ctx context.Context) ContextOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ContextOutput{})
}
