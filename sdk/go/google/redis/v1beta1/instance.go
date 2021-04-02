// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Creates a Redis instance based on the specified tier and memory size. By default, the instance is accessible from the project's [default network](https://cloud.google.com/vpc/docs/vpc). The creation is executed asynchronously and callers may check the returned operation to track its progress. Once the operation is completed the Redis instance will be fully functional. Completed longrunning.Operation will contain the new instance object in the response field. The returned operation is automatically deleted after a few hours, so there is no need to call DeleteOperation.
type Instance struct {
	pulumi.CustomResourceState

	// Optional. Only applicable to STANDARD_HA tier which protects the instance against zonal failures by provisioning it across two zones. If provided, it must be a different zone from the one provided in location_id.
	AlternativeLocationId pulumi.StringOutput `pulumi:"alternativeLocationId"`
	// Optional. Indicates whether OSS Redis AUTH is enabled for the instance. If set to "true" AUTH is enabled on the instance. Default value is "false" meaning AUTH is disabled.
	AuthEnabled pulumi.BoolOutput `pulumi:"authEnabled"`
	// Optional. The full name of the Google Compute Engine [network](https://cloud.google.com/vpc/docs/vpc) to which the instance is connected. If left unspecified, the `default` network will be used.
	AuthorizedNetwork pulumi.StringOutput `pulumi:"authorizedNetwork"`
	// Optional. The network connect mode of the Redis instance. If not provided, the connect mode defaults to DIRECT_PEERING.
	ConnectMode pulumi.StringOutput `pulumi:"connectMode"`
	// The time the instance was created.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// The current zone where the Redis endpoint is placed. For Basic Tier instances, this will always be the same as the location_id provided by the user at creation time. For Standard Tier instances, this can be either location_id or alternative_location_id and can change after a failover event.
	CurrentLocationId pulumi.StringOutput `pulumi:"currentLocationId"`
	// An arbitrary and optional user-provided name for the instance.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// Hostname or IP address of the exposed Redis endpoint used by clients to connect to the service.
	Host pulumi.StringOutput `pulumi:"host"`
	// Resource labels to represent user provided metadata
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// Optional. The zone where the instance will be provisioned. If not provided, the service will choose a zone for the instance. For STANDARD_HA tier, instances will be created across two zones for protection against zonal failures. If alternative_location_id is also provided, it must be different from location_id.
	LocationId pulumi.StringOutput `pulumi:"locationId"`
	// Required. Redis memory size in GiB.
	MemorySizeGb pulumi.IntOutput `pulumi:"memorySizeGb"`
	// Required. Unique name of the resource in this scope including project and location using the form: `projects/{project_id}/locations/{location_id}/instances/{instance_id}` Note: Redis instances are managed and addressed at regional level so location_id here refers to a GCP region; however, users may choose which specific zone (or collection of zones for cross-zone instances) an instance should be provisioned in. Refer to location_id and alternative_location_id fields for more details.
	Name pulumi.StringOutput `pulumi:"name"`
	// Cloud IAM identity used by import / export operations to transfer data to/from Cloud Storage. Format is "serviceAccount:". The value may change over time for a given instance so should be checked before each import/export operation.
	PersistenceIamIdentity pulumi.StringOutput `pulumi:"persistenceIamIdentity"`
	// The port number of the exposed Redis endpoint.
	Port pulumi.IntOutput `pulumi:"port"`
	// Optional. Redis configuration parameters, according to http://redis.io/topics/config. Currently, the only supported parameters are: Redis version 3.2 and newer: * maxmemory-policy * notify-keyspace-events Redis version 4.0 and newer: * activedefrag * lfu-decay-time * lfu-log-factor * maxmemory-gb Redis version 5.0 and newer: * stream-node-max-bytes * stream-node-max-entries
	RedisConfigs pulumi.StringMapOutput `pulumi:"redisConfigs"`
	// Optional. The version of Redis software. If not provided, latest supported version will be used. Currently, the supported values are: * `REDIS_3_2` for Redis 3.2 compatibility * `REDIS_4_0` for Redis 4.0 compatibility (default) * `REDIS_5_0` for Redis 5.0 compatibility
	RedisVersion pulumi.StringOutput `pulumi:"redisVersion"`
	// Optional. For DIRECT_PEERING mode, the CIDR range of internal addresses that are reserved for this instance. Range must be unique and non-overlapping with existing subnets in an authorized network. For PRIVATE_SERVICE_ACCESS mode, the name of one allocated IP address ranges associated with this private service access connection. If not provided, the service will choose an unused /29 block, for example, 10.0.0.0/29 or 192.168.0.0/29.
	ReservedIpRange pulumi.StringOutput `pulumi:"reservedIpRange"`
	// List of server CA certificates for the instance.
	ServerCaCerts TlsCertificateResponseArrayOutput `pulumi:"serverCaCerts"`
	// The current state of this instance.
	State pulumi.StringOutput `pulumi:"state"`
	// Additional information about the current status of this instance, if available.
	StatusMessage pulumi.StringOutput `pulumi:"statusMessage"`
	// Required. The service tier of the instance.
	Tier pulumi.StringOutput `pulumi:"tier"`
	// Optional. The TLS mode of the Redis instance. If not provided, TLS is disabled for the instance.
	TransitEncryptionMode pulumi.StringOutput `pulumi:"transitEncryptionMode"`
}

// NewInstance registers a new resource with the given unique name, arguments, and options.
func NewInstance(ctx *pulumi.Context,
	name string, args *InstanceArgs, opts ...pulumi.ResourceOption) (*Instance, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstancesId == nil {
		return nil, errors.New("invalid value for required argument 'InstancesId'")
	}
	if args.LocationsId == nil {
		return nil, errors.New("invalid value for required argument 'LocationsId'")
	}
	if args.ProjectsId == nil {
		return nil, errors.New("invalid value for required argument 'ProjectsId'")
	}
	var resource Instance
	err := ctx.RegisterResource("google-cloud:redis/v1beta1:Instance", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInstance gets an existing Instance resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInstance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InstanceState, opts ...pulumi.ResourceOption) (*Instance, error) {
	var resource Instance
	err := ctx.ReadResource("google-cloud:redis/v1beta1:Instance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Instance resources.
type instanceState struct {
	// Optional. Only applicable to STANDARD_HA tier which protects the instance against zonal failures by provisioning it across two zones. If provided, it must be a different zone from the one provided in location_id.
	AlternativeLocationId *string `pulumi:"alternativeLocationId"`
	// Optional. Indicates whether OSS Redis AUTH is enabled for the instance. If set to "true" AUTH is enabled on the instance. Default value is "false" meaning AUTH is disabled.
	AuthEnabled *bool `pulumi:"authEnabled"`
	// Optional. The full name of the Google Compute Engine [network](https://cloud.google.com/vpc/docs/vpc) to which the instance is connected. If left unspecified, the `default` network will be used.
	AuthorizedNetwork *string `pulumi:"authorizedNetwork"`
	// Optional. The network connect mode of the Redis instance. If not provided, the connect mode defaults to DIRECT_PEERING.
	ConnectMode *string `pulumi:"connectMode"`
	// The time the instance was created.
	CreateTime *string `pulumi:"createTime"`
	// The current zone where the Redis endpoint is placed. For Basic Tier instances, this will always be the same as the location_id provided by the user at creation time. For Standard Tier instances, this can be either location_id or alternative_location_id and can change after a failover event.
	CurrentLocationId *string `pulumi:"currentLocationId"`
	// An arbitrary and optional user-provided name for the instance.
	DisplayName *string `pulumi:"displayName"`
	// Hostname or IP address of the exposed Redis endpoint used by clients to connect to the service.
	Host *string `pulumi:"host"`
	// Resource labels to represent user provided metadata
	Labels map[string]string `pulumi:"labels"`
	// Optional. The zone where the instance will be provisioned. If not provided, the service will choose a zone for the instance. For STANDARD_HA tier, instances will be created across two zones for protection against zonal failures. If alternative_location_id is also provided, it must be different from location_id.
	LocationId *string `pulumi:"locationId"`
	// Required. Redis memory size in GiB.
	MemorySizeGb *int `pulumi:"memorySizeGb"`
	// Required. Unique name of the resource in this scope including project and location using the form: `projects/{project_id}/locations/{location_id}/instances/{instance_id}` Note: Redis instances are managed and addressed at regional level so location_id here refers to a GCP region; however, users may choose which specific zone (or collection of zones for cross-zone instances) an instance should be provisioned in. Refer to location_id and alternative_location_id fields for more details.
	Name *string `pulumi:"name"`
	// Cloud IAM identity used by import / export operations to transfer data to/from Cloud Storage. Format is "serviceAccount:". The value may change over time for a given instance so should be checked before each import/export operation.
	PersistenceIamIdentity *string `pulumi:"persistenceIamIdentity"`
	// The port number of the exposed Redis endpoint.
	Port *int `pulumi:"port"`
	// Optional. Redis configuration parameters, according to http://redis.io/topics/config. Currently, the only supported parameters are: Redis version 3.2 and newer: * maxmemory-policy * notify-keyspace-events Redis version 4.0 and newer: * activedefrag * lfu-decay-time * lfu-log-factor * maxmemory-gb Redis version 5.0 and newer: * stream-node-max-bytes * stream-node-max-entries
	RedisConfigs map[string]string `pulumi:"redisConfigs"`
	// Optional. The version of Redis software. If not provided, latest supported version will be used. Currently, the supported values are: * `REDIS_3_2` for Redis 3.2 compatibility * `REDIS_4_0` for Redis 4.0 compatibility (default) * `REDIS_5_0` for Redis 5.0 compatibility
	RedisVersion *string `pulumi:"redisVersion"`
	// Optional. For DIRECT_PEERING mode, the CIDR range of internal addresses that are reserved for this instance. Range must be unique and non-overlapping with existing subnets in an authorized network. For PRIVATE_SERVICE_ACCESS mode, the name of one allocated IP address ranges associated with this private service access connection. If not provided, the service will choose an unused /29 block, for example, 10.0.0.0/29 or 192.168.0.0/29.
	ReservedIpRange *string `pulumi:"reservedIpRange"`
	// List of server CA certificates for the instance.
	ServerCaCerts []TlsCertificateResponse `pulumi:"serverCaCerts"`
	// The current state of this instance.
	State *string `pulumi:"state"`
	// Additional information about the current status of this instance, if available.
	StatusMessage *string `pulumi:"statusMessage"`
	// Required. The service tier of the instance.
	Tier *string `pulumi:"tier"`
	// Optional. The TLS mode of the Redis instance. If not provided, TLS is disabled for the instance.
	TransitEncryptionMode *string `pulumi:"transitEncryptionMode"`
}

type InstanceState struct {
	// Optional. Only applicable to STANDARD_HA tier which protects the instance against zonal failures by provisioning it across two zones. If provided, it must be a different zone from the one provided in location_id.
	AlternativeLocationId pulumi.StringPtrInput
	// Optional. Indicates whether OSS Redis AUTH is enabled for the instance. If set to "true" AUTH is enabled on the instance. Default value is "false" meaning AUTH is disabled.
	AuthEnabled pulumi.BoolPtrInput
	// Optional. The full name of the Google Compute Engine [network](https://cloud.google.com/vpc/docs/vpc) to which the instance is connected. If left unspecified, the `default` network will be used.
	AuthorizedNetwork pulumi.StringPtrInput
	// Optional. The network connect mode of the Redis instance. If not provided, the connect mode defaults to DIRECT_PEERING.
	ConnectMode pulumi.StringPtrInput
	// The time the instance was created.
	CreateTime pulumi.StringPtrInput
	// The current zone where the Redis endpoint is placed. For Basic Tier instances, this will always be the same as the location_id provided by the user at creation time. For Standard Tier instances, this can be either location_id or alternative_location_id and can change after a failover event.
	CurrentLocationId pulumi.StringPtrInput
	// An arbitrary and optional user-provided name for the instance.
	DisplayName pulumi.StringPtrInput
	// Hostname or IP address of the exposed Redis endpoint used by clients to connect to the service.
	Host pulumi.StringPtrInput
	// Resource labels to represent user provided metadata
	Labels pulumi.StringMapInput
	// Optional. The zone where the instance will be provisioned. If not provided, the service will choose a zone for the instance. For STANDARD_HA tier, instances will be created across two zones for protection against zonal failures. If alternative_location_id is also provided, it must be different from location_id.
	LocationId pulumi.StringPtrInput
	// Required. Redis memory size in GiB.
	MemorySizeGb pulumi.IntPtrInput
	// Required. Unique name of the resource in this scope including project and location using the form: `projects/{project_id}/locations/{location_id}/instances/{instance_id}` Note: Redis instances are managed and addressed at regional level so location_id here refers to a GCP region; however, users may choose which specific zone (or collection of zones for cross-zone instances) an instance should be provisioned in. Refer to location_id and alternative_location_id fields for more details.
	Name pulumi.StringPtrInput
	// Cloud IAM identity used by import / export operations to transfer data to/from Cloud Storage. Format is "serviceAccount:". The value may change over time for a given instance so should be checked before each import/export operation.
	PersistenceIamIdentity pulumi.StringPtrInput
	// The port number of the exposed Redis endpoint.
	Port pulumi.IntPtrInput
	// Optional. Redis configuration parameters, according to http://redis.io/topics/config. Currently, the only supported parameters are: Redis version 3.2 and newer: * maxmemory-policy * notify-keyspace-events Redis version 4.0 and newer: * activedefrag * lfu-decay-time * lfu-log-factor * maxmemory-gb Redis version 5.0 and newer: * stream-node-max-bytes * stream-node-max-entries
	RedisConfigs pulumi.StringMapInput
	// Optional. The version of Redis software. If not provided, latest supported version will be used. Currently, the supported values are: * `REDIS_3_2` for Redis 3.2 compatibility * `REDIS_4_0` for Redis 4.0 compatibility (default) * `REDIS_5_0` for Redis 5.0 compatibility
	RedisVersion pulumi.StringPtrInput
	// Optional. For DIRECT_PEERING mode, the CIDR range of internal addresses that are reserved for this instance. Range must be unique and non-overlapping with existing subnets in an authorized network. For PRIVATE_SERVICE_ACCESS mode, the name of one allocated IP address ranges associated with this private service access connection. If not provided, the service will choose an unused /29 block, for example, 10.0.0.0/29 or 192.168.0.0/29.
	ReservedIpRange pulumi.StringPtrInput
	// List of server CA certificates for the instance.
	ServerCaCerts TlsCertificateResponseArrayInput
	// The current state of this instance.
	State pulumi.StringPtrInput
	// Additional information about the current status of this instance, if available.
	StatusMessage pulumi.StringPtrInput
	// Required. The service tier of the instance.
	Tier pulumi.StringPtrInput
	// Optional. The TLS mode of the Redis instance. If not provided, TLS is disabled for the instance.
	TransitEncryptionMode pulumi.StringPtrInput
}

func (InstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceState)(nil)).Elem()
}

type instanceArgs struct {
	// Optional. Only applicable to STANDARD_HA tier which protects the instance against zonal failures by provisioning it across two zones. If provided, it must be a different zone from the one provided in location_id.
	AlternativeLocationId *string `pulumi:"alternativeLocationId"`
	// Optional. Indicates whether OSS Redis AUTH is enabled for the instance. If set to "true" AUTH is enabled on the instance. Default value is "false" meaning AUTH is disabled.
	AuthEnabled *bool `pulumi:"authEnabled"`
	// Optional. The full name of the Google Compute Engine [network](https://cloud.google.com/vpc/docs/vpc) to which the instance is connected. If left unspecified, the `default` network will be used.
	AuthorizedNetwork *string `pulumi:"authorizedNetwork"`
	// Optional. The network connect mode of the Redis instance. If not provided, the connect mode defaults to DIRECT_PEERING.
	ConnectMode *string `pulumi:"connectMode"`
	// An arbitrary and optional user-provided name for the instance.
	DisplayName *string `pulumi:"displayName"`
	InstancesId string  `pulumi:"instancesId"`
	// Resource labels to represent user provided metadata
	Labels map[string]string `pulumi:"labels"`
	// Optional. The zone where the instance will be provisioned. If not provided, the service will choose a zone for the instance. For STANDARD_HA tier, instances will be created across two zones for protection against zonal failures. If alternative_location_id is also provided, it must be different from location_id.
	LocationId  *string `pulumi:"locationId"`
	LocationsId string  `pulumi:"locationsId"`
	// Required. Redis memory size in GiB.
	MemorySizeGb *int `pulumi:"memorySizeGb"`
	// Required. Unique name of the resource in this scope including project and location using the form: `projects/{project_id}/locations/{location_id}/instances/{instance_id}` Note: Redis instances are managed and addressed at regional level so location_id here refers to a GCP region; however, users may choose which specific zone (or collection of zones for cross-zone instances) an instance should be provisioned in. Refer to location_id and alternative_location_id fields for more details.
	Name       *string `pulumi:"name"`
	ProjectsId string  `pulumi:"projectsId"`
	// Optional. Redis configuration parameters, according to http://redis.io/topics/config. Currently, the only supported parameters are: Redis version 3.2 and newer: * maxmemory-policy * notify-keyspace-events Redis version 4.0 and newer: * activedefrag * lfu-decay-time * lfu-log-factor * maxmemory-gb Redis version 5.0 and newer: * stream-node-max-bytes * stream-node-max-entries
	RedisConfigs map[string]string `pulumi:"redisConfigs"`
	// Optional. The version of Redis software. If not provided, latest supported version will be used. Currently, the supported values are: * `REDIS_3_2` for Redis 3.2 compatibility * `REDIS_4_0` for Redis 4.0 compatibility (default) * `REDIS_5_0` for Redis 5.0 compatibility
	RedisVersion *string `pulumi:"redisVersion"`
	// Optional. For DIRECT_PEERING mode, the CIDR range of internal addresses that are reserved for this instance. Range must be unique and non-overlapping with existing subnets in an authorized network. For PRIVATE_SERVICE_ACCESS mode, the name of one allocated IP address ranges associated with this private service access connection. If not provided, the service will choose an unused /29 block, for example, 10.0.0.0/29 or 192.168.0.0/29.
	ReservedIpRange *string `pulumi:"reservedIpRange"`
	// Required. The service tier of the instance.
	Tier *string `pulumi:"tier"`
	// Optional. The TLS mode of the Redis instance. If not provided, TLS is disabled for the instance.
	TransitEncryptionMode *string `pulumi:"transitEncryptionMode"`
}

// The set of arguments for constructing a Instance resource.
type InstanceArgs struct {
	// Optional. Only applicable to STANDARD_HA tier which protects the instance against zonal failures by provisioning it across two zones. If provided, it must be a different zone from the one provided in location_id.
	AlternativeLocationId pulumi.StringPtrInput
	// Optional. Indicates whether OSS Redis AUTH is enabled for the instance. If set to "true" AUTH is enabled on the instance. Default value is "false" meaning AUTH is disabled.
	AuthEnabled pulumi.BoolPtrInput
	// Optional. The full name of the Google Compute Engine [network](https://cloud.google.com/vpc/docs/vpc) to which the instance is connected. If left unspecified, the `default` network will be used.
	AuthorizedNetwork pulumi.StringPtrInput
	// Optional. The network connect mode of the Redis instance. If not provided, the connect mode defaults to DIRECT_PEERING.
	ConnectMode pulumi.StringPtrInput
	// An arbitrary and optional user-provided name for the instance.
	DisplayName pulumi.StringPtrInput
	InstancesId pulumi.StringInput
	// Resource labels to represent user provided metadata
	Labels pulumi.StringMapInput
	// Optional. The zone where the instance will be provisioned. If not provided, the service will choose a zone for the instance. For STANDARD_HA tier, instances will be created across two zones for protection against zonal failures. If alternative_location_id is also provided, it must be different from location_id.
	LocationId  pulumi.StringPtrInput
	LocationsId pulumi.StringInput
	// Required. Redis memory size in GiB.
	MemorySizeGb pulumi.IntPtrInput
	// Required. Unique name of the resource in this scope including project and location using the form: `projects/{project_id}/locations/{location_id}/instances/{instance_id}` Note: Redis instances are managed and addressed at regional level so location_id here refers to a GCP region; however, users may choose which specific zone (or collection of zones for cross-zone instances) an instance should be provisioned in. Refer to location_id and alternative_location_id fields for more details.
	Name       pulumi.StringPtrInput
	ProjectsId pulumi.StringInput
	// Optional. Redis configuration parameters, according to http://redis.io/topics/config. Currently, the only supported parameters are: Redis version 3.2 and newer: * maxmemory-policy * notify-keyspace-events Redis version 4.0 and newer: * activedefrag * lfu-decay-time * lfu-log-factor * maxmemory-gb Redis version 5.0 and newer: * stream-node-max-bytes * stream-node-max-entries
	RedisConfigs pulumi.StringMapInput
	// Optional. The version of Redis software. If not provided, latest supported version will be used. Currently, the supported values are: * `REDIS_3_2` for Redis 3.2 compatibility * `REDIS_4_0` for Redis 4.0 compatibility (default) * `REDIS_5_0` for Redis 5.0 compatibility
	RedisVersion pulumi.StringPtrInput
	// Optional. For DIRECT_PEERING mode, the CIDR range of internal addresses that are reserved for this instance. Range must be unique and non-overlapping with existing subnets in an authorized network. For PRIVATE_SERVICE_ACCESS mode, the name of one allocated IP address ranges associated with this private service access connection. If not provided, the service will choose an unused /29 block, for example, 10.0.0.0/29 or 192.168.0.0/29.
	ReservedIpRange pulumi.StringPtrInput
	// Required. The service tier of the instance.
	Tier pulumi.StringPtrInput
	// Optional. The TLS mode of the Redis instance. If not provided, TLS is disabled for the instance.
	TransitEncryptionMode pulumi.StringPtrInput
}

func (InstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceArgs)(nil)).Elem()
}

type InstanceInput interface {
	pulumi.Input

	ToInstanceOutput() InstanceOutput
	ToInstanceOutputWithContext(ctx context.Context) InstanceOutput
}

func (*Instance) ElementType() reflect.Type {
	return reflect.TypeOf((*Instance)(nil))
}

func (i *Instance) ToInstanceOutput() InstanceOutput {
	return i.ToInstanceOutputWithContext(context.Background())
}

func (i *Instance) ToInstanceOutputWithContext(ctx context.Context) InstanceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceOutput)
}

type InstanceOutput struct {
	*pulumi.OutputState
}

func (InstanceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Instance)(nil))
}

func (o InstanceOutput) ToInstanceOutput() InstanceOutput {
	return o
}

func (o InstanceOutput) ToInstanceOutputWithContext(ctx context.Context) InstanceOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(InstanceOutput{})
}