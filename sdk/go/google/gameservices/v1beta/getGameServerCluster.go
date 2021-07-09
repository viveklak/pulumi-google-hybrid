// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets details of a single game server cluster.
func LookupGameServerCluster(ctx *pulumi.Context, args *LookupGameServerClusterArgs, opts ...pulumi.InvokeOption) (*LookupGameServerClusterResult, error) {
	var rv LookupGameServerClusterResult
	err := ctx.Invoke("google-native:gameservices/v1beta:getGameServerCluster", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupGameServerClusterArgs struct {
	GameServerClusterId string `pulumi:"gameServerClusterId"`
	Location            string `pulumi:"location"`
	Project             string `pulumi:"project"`
	RealmId             string `pulumi:"realmId"`
}

type LookupGameServerClusterResult struct {
	// Optional. The allocation priority assigned to the game server cluster. Game server clusters receive new game server allocations based on the relative allocation priorites set for each cluster, if the realm is configured for multicluster allocation.
	AllocationPriority string `pulumi:"allocationPriority"`
	// The game server cluster connection information. This information is used to manage game server clusters.
	ConnectionInfo GameServerClusterConnectionInfoResponse `pulumi:"connectionInfo"`
	// The creation time.
	CreateTime string `pulumi:"createTime"`
	// Human readable description of the cluster.
	Description string `pulumi:"description"`
	// ETag of the resource.
	Etag string `pulumi:"etag"`
	// The labels associated with this game server cluster. Each label is a key-value pair.
	Labels map[string]string `pulumi:"labels"`
	// The resource name of the game server cluster, in the following form: `projects/{project}/locations/{location}/realms/{realm}/gameServerClusters/{cluster}`. For example, `projects/my-project/locations/{location}/realms/zanzibar/gameServerClusters/my-onprem-cluster`.
	Name string `pulumi:"name"`
	// The last-modified time.
	UpdateTime string `pulumi:"updateTime"`
}
