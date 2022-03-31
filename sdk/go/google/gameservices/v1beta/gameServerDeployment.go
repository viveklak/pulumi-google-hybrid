// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a new game server deployment in a given project and location.
type GameServerDeployment struct {
	pulumi.CustomResourceState

	// The creation time.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Human readable description of the game server deployment.
	Description pulumi.StringOutput `pulumi:"description"`
	// Used to perform consistent read-modify-write updates. If not set, a blind "overwrite" update happens.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The labels associated with this game server deployment. Each label is a key-value pair.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// The resource name of the game server deployment, in the following form: `projects/{project}/locations/{locationId}/gameServerDeployments/{deploymentId}`. For example, `projects/my-project/locations/global/gameServerDeployments/my-deployment`.
	Name pulumi.StringOutput `pulumi:"name"`
	// The last-modified time.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewGameServerDeployment registers a new resource with the given unique name, arguments, and options.
func NewGameServerDeployment(ctx *pulumi.Context,
	name string, args *GameServerDeploymentArgs, opts ...pulumi.ResourceOption) (*GameServerDeployment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DeploymentId == nil {
		return nil, errors.New("invalid value for required argument 'DeploymentId'")
	}
	var resource GameServerDeployment
	err := ctx.RegisterResource("google-native:gameservices/v1beta:GameServerDeployment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGameServerDeployment gets an existing GameServerDeployment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGameServerDeployment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GameServerDeploymentState, opts ...pulumi.ResourceOption) (*GameServerDeployment, error) {
	var resource GameServerDeployment
	err := ctx.ReadResource("google-native:gameservices/v1beta:GameServerDeployment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GameServerDeployment resources.
type gameServerDeploymentState struct {
}

type GameServerDeploymentState struct {
}

func (GameServerDeploymentState) ElementType() reflect.Type {
	return reflect.TypeOf((*gameServerDeploymentState)(nil)).Elem()
}

type gameServerDeploymentArgs struct {
	// Required. The ID of the game server deployment resource to create.
	DeploymentId string `pulumi:"deploymentId"`
	// Human readable description of the game server deployment.
	Description *string `pulumi:"description"`
	// Used to perform consistent read-modify-write updates. If not set, a blind "overwrite" update happens.
	Etag *string `pulumi:"etag"`
	// The labels associated with this game server deployment. Each label is a key-value pair.
	Labels   map[string]string `pulumi:"labels"`
	Location *string           `pulumi:"location"`
	// The resource name of the game server deployment, in the following form: `projects/{project}/locations/{locationId}/gameServerDeployments/{deploymentId}`. For example, `projects/my-project/locations/global/gameServerDeployments/my-deployment`.
	Name    *string `pulumi:"name"`
	Project *string `pulumi:"project"`
}

// The set of arguments for constructing a GameServerDeployment resource.
type GameServerDeploymentArgs struct {
	// Required. The ID of the game server deployment resource to create.
	DeploymentId pulumi.StringInput
	// Human readable description of the game server deployment.
	Description pulumi.StringPtrInput
	// Used to perform consistent read-modify-write updates. If not set, a blind "overwrite" update happens.
	Etag pulumi.StringPtrInput
	// The labels associated with this game server deployment. Each label is a key-value pair.
	Labels   pulumi.StringMapInput
	Location pulumi.StringPtrInput
	// The resource name of the game server deployment, in the following form: `projects/{project}/locations/{locationId}/gameServerDeployments/{deploymentId}`. For example, `projects/my-project/locations/global/gameServerDeployments/my-deployment`.
	Name    pulumi.StringPtrInput
	Project pulumi.StringPtrInput
}

func (GameServerDeploymentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*gameServerDeploymentArgs)(nil)).Elem()
}

type GameServerDeploymentInput interface {
	pulumi.Input

	ToGameServerDeploymentOutput() GameServerDeploymentOutput
	ToGameServerDeploymentOutputWithContext(ctx context.Context) GameServerDeploymentOutput
}

func (*GameServerDeployment) ElementType() reflect.Type {
	return reflect.TypeOf((**GameServerDeployment)(nil)).Elem()
}

func (i *GameServerDeployment) ToGameServerDeploymentOutput() GameServerDeploymentOutput {
	return i.ToGameServerDeploymentOutputWithContext(context.Background())
}

func (i *GameServerDeployment) ToGameServerDeploymentOutputWithContext(ctx context.Context) GameServerDeploymentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GameServerDeploymentOutput)
}

type GameServerDeploymentOutput struct{ *pulumi.OutputState }

func (GameServerDeploymentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GameServerDeployment)(nil)).Elem()
}

func (o GameServerDeploymentOutput) ToGameServerDeploymentOutput() GameServerDeploymentOutput {
	return o
}

func (o GameServerDeploymentOutput) ToGameServerDeploymentOutputWithContext(ctx context.Context) GameServerDeploymentOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GameServerDeploymentInput)(nil)).Elem(), &GameServerDeployment{})
	pulumi.RegisterOutputType(GameServerDeploymentOutput{})
}
