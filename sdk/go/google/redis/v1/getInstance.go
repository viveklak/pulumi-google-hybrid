// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the details of a specific Redis instance.
func LookupInstance(ctx *pulumi.Context, args *LookupInstanceArgs, opts ...pulumi.InvokeOption) (*LookupInstanceResult, error) {
	var rv LookupInstanceResult
	err := ctx.Invoke("google-native:redis/v1:getInstance", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupInstanceArgs struct {
	InstanceId string `pulumi:"instanceId"`
	Location   string `pulumi:"location"`
	Project    string `pulumi:"project"`
}

type LookupInstanceResult struct {
	// Optional. Only applicable to STANDARD_HA tier which protects the instance against zonal failures by provisioning it across two zones. If provided, it must be a different zone from the one provided in location_id.
	AlternativeLocationId string `pulumi:"alternativeLocationId"`
	// Optional. Indicates whether OSS Redis AUTH is enabled for the instance. If set to "true" AUTH is enabled on the instance. Default value is "false" meaning AUTH is disabled.
	AuthEnabled bool `pulumi:"authEnabled"`
	// Optional. The full name of the Google Compute Engine [network](https://cloud.google.com/vpc/docs/vpc) to which the instance is connected. If left unspecified, the `default` network will be used.
	AuthorizedNetwork string `pulumi:"authorizedNetwork"`
	// Optional. The network connect mode of the Redis instance. If not provided, the connect mode defaults to DIRECT_PEERING.
	ConnectMode string `pulumi:"connectMode"`
	// The time the instance was created.
	CreateTime string `pulumi:"createTime"`
	// The current zone where the Redis endpoint is placed. For Basic Tier instances, this will always be the same as the location_id provided by the user at creation time. For Standard Tier instances, this can be either location_id or alternative_location_id and can change after a failover event.
	CurrentLocationId string `pulumi:"currentLocationId"`
	// An arbitrary and optional user-provided name for the instance.
	DisplayName string `pulumi:"displayName"`
	// Hostname or IP address of the exposed Redis endpoint used by clients to connect to the service.
	Host string `pulumi:"host"`
	// Resource labels to represent user provided metadata
	Labels map[string]string `pulumi:"labels"`
	// Optional. The zone where the instance will be provisioned. If not provided, the service will choose a zone for the instance. For STANDARD_HA tier, instances will be created across two zones for protection against zonal failures. If alternative_location_id is also provided, it must be different from location_id.
	Location string `pulumi:"location"`
	// Redis memory size in GiB.
	MemorySizeGb int `pulumi:"memorySizeGb"`
	// Unique name of the resource in this scope including project and location using the form: `projects/{project_id}/locations/{location_id}/instances/{instance_id}` Note: Redis instances are managed and addressed at regional level so location_id here refers to a GCP region; however, users may choose which specific zone (or collection of zones for cross-zone instances) an instance should be provisioned in. Refer to location_id and alternative_location_id fields for more details.
	Name string `pulumi:"name"`
	// Cloud IAM identity used by import / export operations to transfer data to/from Cloud Storage. Format is "serviceAccount:". The value may change over time for a given instance so should be checked before each import/export operation.
	PersistenceIamIdentity string `pulumi:"persistenceIamIdentity"`
	// The port number of the exposed Redis endpoint.
	Port int `pulumi:"port"`
	// Optional. Redis configuration parameters, according to http://redis.io/topics/config. Currently, the only supported parameters are: Redis version 3.2 and newer: * maxmemory-policy * notify-keyspace-events Redis version 4.0 and newer: * activedefrag * lfu-decay-time * lfu-log-factor * maxmemory-gb Redis version 5.0 and newer: * stream-node-max-bytes * stream-node-max-entries
	RedisConfigs map[string]string `pulumi:"redisConfigs"`
	// Optional. The version of Redis software. If not provided, latest supported version will be used. Currently, the supported values are: * `REDIS_3_2` for Redis 3.2 compatibility * `REDIS_4_0` for Redis 4.0 compatibility (default) * `REDIS_5_0` for Redis 5.0 compatibility * `REDIS_6_X` for Redis 6.x compatibility
	RedisVersion string `pulumi:"redisVersion"`
	// Optional. For DIRECT_PEERING mode, the CIDR range of internal addresses that are reserved for this instance. Range must be unique and non-overlapping with existing subnets in an authorized network. For PRIVATE_SERVICE_ACCESS mode, the name of one allocated IP address ranges associated with this private service access connection. If not provided, the service will choose an unused /29 block, for example, 10.0.0.0/29 or 192.168.0.0/29.
	ReservedIpRange string `pulumi:"reservedIpRange"`
	// List of server CA certificates for the instance.
	ServerCaCerts []TlsCertificateResponse `pulumi:"serverCaCerts"`
	// The current state of this instance.
	State string `pulumi:"state"`
	// Additional information about the current status of this instance, if available.
	StatusMessage string `pulumi:"statusMessage"`
	// The service tier of the instance.
	Tier string `pulumi:"tier"`
	// Optional. The TLS mode of the Redis instance. If not provided, TLS is disabled for the instance.
	TransitEncryptionMode string `pulumi:"transitEncryptionMode"`
}
