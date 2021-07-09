// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets information about a cluster.
func LookupCluster(ctx *pulumi.Context, args *LookupClusterArgs, opts ...pulumi.InvokeOption) (*LookupClusterResult, error) {
	var rv LookupClusterResult
	err := ctx.Invoke("google-native:bigtableadmin/v2:getCluster", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupClusterArgs struct {
	ClusterId  string `pulumi:"clusterId"`
	InstanceId string `pulumi:"instanceId"`
	Project    string `pulumi:"project"`
}

type LookupClusterResult struct {
	// Immutable. The type of storage used by this cluster to serve its parent instance's tables, unless explicitly overridden.
	DefaultStorageType string `pulumi:"defaultStorageType"`
	// Immutable. The encryption configuration for CMEK-protected clusters.
	EncryptionConfig EncryptionConfigResponse `pulumi:"encryptionConfig"`
	// Immutable. The location where this cluster's nodes and storage reside. For best performance, clients should be located as close as possible to this cluster. Currently only zones are supported, so values should be of the form `projects/{project}/locations/{zone}`.
	Location string `pulumi:"location"`
	// The unique name of the cluster. Values are of the form `projects/{project}/instances/{instance}/clusters/a-z*`.
	Name string `pulumi:"name"`
	// The number of nodes allocated to this cluster. More nodes enable higher throughput and more consistent performance.
	ServeNodes int `pulumi:"serveNodes"`
	// The current state of the cluster.
	State string `pulumi:"state"`
}
