// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a new game server cluster in a given project and location.
type GameServerCluster struct {
	pulumi.CustomResourceState

	// The game server cluster connection information. This information is used to manage game server clusters.
	ConnectionInfo GameServerClusterConnectionInfoResponseOutput `pulumi:"connectionInfo"`
	// The creation time.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Human readable description of the cluster.
	Description pulumi.StringOutput `pulumi:"description"`
	// ETag of the resource.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The labels associated with this game server cluster. Each label is a key-value pair.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// Required. The resource name of the game server cluster, in the following form: `projects/{project}/locations/{location}/realms/{realm}/gameServerClusters/{cluster}`. For example, `projects/my-project/locations/{location}/realms/zanzibar/gameServerClusters/my-onprem-cluster`.
	Name pulumi.StringOutput `pulumi:"name"`
	// The last-modified time.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewGameServerCluster registers a new resource with the given unique name, arguments, and options.
func NewGameServerCluster(ctx *pulumi.Context,
	name string, args *GameServerClusterArgs, opts ...pulumi.ResourceOption) (*GameServerCluster, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.GameServerClusterId == nil {
		return nil, errors.New("invalid value for required argument 'GameServerClusterId'")
	}
	if args.Location == nil {
		return nil, errors.New("invalid value for required argument 'Location'")
	}
	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	if args.RealmId == nil {
		return nil, errors.New("invalid value for required argument 'RealmId'")
	}
	var resource GameServerCluster
	err := ctx.RegisterResource("google-native:gameservices/v1:GameServerCluster", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGameServerCluster gets an existing GameServerCluster resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGameServerCluster(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GameServerClusterState, opts ...pulumi.ResourceOption) (*GameServerCluster, error) {
	var resource GameServerCluster
	err := ctx.ReadResource("google-native:gameservices/v1:GameServerCluster", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GameServerCluster resources.
type gameServerClusterState struct {
}

type GameServerClusterState struct {
}

func (GameServerClusterState) ElementType() reflect.Type {
	return reflect.TypeOf((*gameServerClusterState)(nil)).Elem()
}

type gameServerClusterArgs struct {
	// The game server cluster connection information. This information is used to manage game server clusters.
	ConnectionInfo *GameServerClusterConnectionInfo `pulumi:"connectionInfo"`
	// Human readable description of the cluster.
	Description *string `pulumi:"description"`
	// ETag of the resource.
	Etag                *string `pulumi:"etag"`
	GameServerClusterId string  `pulumi:"gameServerClusterId"`
	// The labels associated with this game server cluster. Each label is a key-value pair.
	Labels   map[string]string `pulumi:"labels"`
	Location string            `pulumi:"location"`
	// Required. The resource name of the game server cluster, in the following form: `projects/{project}/locations/{location}/realms/{realm}/gameServerClusters/{cluster}`. For example, `projects/my-project/locations/{location}/realms/zanzibar/gameServerClusters/my-onprem-cluster`.
	Name    *string `pulumi:"name"`
	Project string  `pulumi:"project"`
	RealmId string  `pulumi:"realmId"`
}

// The set of arguments for constructing a GameServerCluster resource.
type GameServerClusterArgs struct {
	// The game server cluster connection information. This information is used to manage game server clusters.
	ConnectionInfo GameServerClusterConnectionInfoPtrInput
	// Human readable description of the cluster.
	Description pulumi.StringPtrInput
	// ETag of the resource.
	Etag                pulumi.StringPtrInput
	GameServerClusterId pulumi.StringInput
	// The labels associated with this game server cluster. Each label is a key-value pair.
	Labels   pulumi.StringMapInput
	Location pulumi.StringInput
	// Required. The resource name of the game server cluster, in the following form: `projects/{project}/locations/{location}/realms/{realm}/gameServerClusters/{cluster}`. For example, `projects/my-project/locations/{location}/realms/zanzibar/gameServerClusters/my-onprem-cluster`.
	Name    pulumi.StringPtrInput
	Project pulumi.StringInput
	RealmId pulumi.StringInput
}

func (GameServerClusterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*gameServerClusterArgs)(nil)).Elem()
}

type GameServerClusterInput interface {
	pulumi.Input

	ToGameServerClusterOutput() GameServerClusterOutput
	ToGameServerClusterOutputWithContext(ctx context.Context) GameServerClusterOutput
}

func (*GameServerCluster) ElementType() reflect.Type {
	return reflect.TypeOf((*GameServerCluster)(nil))
}

func (i *GameServerCluster) ToGameServerClusterOutput() GameServerClusterOutput {
	return i.ToGameServerClusterOutputWithContext(context.Background())
}

func (i *GameServerCluster) ToGameServerClusterOutputWithContext(ctx context.Context) GameServerClusterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GameServerClusterOutput)
}

type GameServerClusterOutput struct {
	*pulumi.OutputState
}

func (GameServerClusterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GameServerCluster)(nil))
}

func (o GameServerClusterOutput) ToGameServerClusterOutput() GameServerClusterOutput {
	return o
}

func (o GameServerClusterOutput) ToGameServerClusterOutputWithContext(ctx context.Context) GameServerClusterOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(GameServerClusterOutput{})
}
