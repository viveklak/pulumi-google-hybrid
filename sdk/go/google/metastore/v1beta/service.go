// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a metastore service in a project and location.
type Service struct {
	pulumi.CustomResourceState

	// A Cloud Storage URI (starting with gs://) that specifies where artifacts related to the metastore service are stored.
	ArtifactGcsUri pulumi.StringOutput `pulumi:"artifactGcsUri"`
	// The time when the metastore service was created.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Immutable. The database type that the Metastore service stores its data.
	DatabaseType pulumi.StringOutput `pulumi:"databaseType"`
	// Immutable. Information used to configure the Dataproc Metastore service to encrypt customer data at rest. Cannot be updated.
	EncryptionConfig EncryptionConfigResponseOutput `pulumi:"encryptionConfig"`
	// The URI of the endpoint used to access the metastore service.
	EndpointUri pulumi.StringOutput `pulumi:"endpointUri"`
	// Configuration information specific to running Hive metastore software as the metastore service.
	HiveMetastoreConfig HiveMetastoreConfigResponseOutput `pulumi:"hiveMetastoreConfig"`
	// User-defined labels for the metastore service.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// The one hour maintenance window of the metastore service. This specifies when the service can be restarted for maintenance purposes in UTC time. Maintenance window is not needed for services with the SPANNER database type.
	MaintenanceWindow MaintenanceWindowResponseOutput `pulumi:"maintenanceWindow"`
	// The setting that defines how metastore metadata should be integrated with external services and systems.
	MetadataIntegration MetadataIntegrationResponseOutput `pulumi:"metadataIntegration"`
	// The metadata management activities of the metastore service.
	MetadataManagementActivity MetadataManagementActivityResponseOutput `pulumi:"metadataManagementActivity"`
	// Immutable. The relative resource name of the metastore service, of the form:projects/{project_number}/locations/{location_id}/services/{service_id}.
	Name pulumi.StringOutput `pulumi:"name"`
	// Immutable. The relative resource name of the VPC network on which the instance can be accessed. It is specified in the following form:projects/{project_number}/global/networks/{network_id}.
	Network pulumi.StringOutput `pulumi:"network"`
	// Immutable. The configuration specifying the network settings for the Dataproc Metastore service.
	NetworkConfig NetworkConfigResponseOutput `pulumi:"networkConfig"`
	// The TCP port at which the metastore service is reached. Default: 9083.
	Port pulumi.IntOutput `pulumi:"port"`
	// Immutable. The release channel of the service. If unspecified, defaults to STABLE.
	ReleaseChannel pulumi.StringOutput `pulumi:"releaseChannel"`
	// The current state of the metastore service.
	State pulumi.StringOutput `pulumi:"state"`
	// Additional information about the current state of the metastore service, if available.
	StateMessage pulumi.StringOutput `pulumi:"stateMessage"`
	// The tier of the service.
	Tier pulumi.StringOutput `pulumi:"tier"`
	// The globally unique resource identifier of the metastore service.
	Uid pulumi.StringOutput `pulumi:"uid"`
	// The time when the metastore service was last updated.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewService registers a new resource with the given unique name, arguments, and options.
func NewService(ctx *pulumi.Context,
	name string, args *ServiceArgs, opts ...pulumi.ResourceOption) (*Service, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ServiceId == nil {
		return nil, errors.New("invalid value for required argument 'ServiceId'")
	}
	var resource Service
	err := ctx.RegisterResource("google-native:metastore/v1beta:Service", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetService gets an existing Service resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetService(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServiceState, opts ...pulumi.ResourceOption) (*Service, error) {
	var resource Service
	err := ctx.ReadResource("google-native:metastore/v1beta:Service", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Service resources.
type serviceState struct {
}

type ServiceState struct {
}

func (ServiceState) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceState)(nil)).Elem()
}

type serviceArgs struct {
	// Immutable. The database type that the Metastore service stores its data.
	DatabaseType *ServiceDatabaseType `pulumi:"databaseType"`
	// Immutable. Information used to configure the Dataproc Metastore service to encrypt customer data at rest. Cannot be updated.
	EncryptionConfig *EncryptionConfig `pulumi:"encryptionConfig"`
	// Configuration information specific to running Hive metastore software as the metastore service.
	HiveMetastoreConfig *HiveMetastoreConfig `pulumi:"hiveMetastoreConfig"`
	// User-defined labels for the metastore service.
	Labels   map[string]string `pulumi:"labels"`
	Location *string           `pulumi:"location"`
	// The one hour maintenance window of the metastore service. This specifies when the service can be restarted for maintenance purposes in UTC time. Maintenance window is not needed for services with the SPANNER database type.
	MaintenanceWindow *MaintenanceWindow `pulumi:"maintenanceWindow"`
	// The setting that defines how metastore metadata should be integrated with external services and systems.
	MetadataIntegration *MetadataIntegration `pulumi:"metadataIntegration"`
	// Immutable. The relative resource name of the metastore service, of the form:projects/{project_number}/locations/{location_id}/services/{service_id}.
	Name *string `pulumi:"name"`
	// Immutable. The relative resource name of the VPC network on which the instance can be accessed. It is specified in the following form:projects/{project_number}/global/networks/{network_id}.
	Network *string `pulumi:"network"`
	// Immutable. The configuration specifying the network settings for the Dataproc Metastore service.
	NetworkConfig *NetworkConfig `pulumi:"networkConfig"`
	// The TCP port at which the metastore service is reached. Default: 9083.
	Port    *int    `pulumi:"port"`
	Project *string `pulumi:"project"`
	// Immutable. The release channel of the service. If unspecified, defaults to STABLE.
	ReleaseChannel *ServiceReleaseChannel `pulumi:"releaseChannel"`
	// Optional. A request ID. Specify a unique request ID to allow the server to ignore the request if it has completed. The server will ignore subsequent requests that provide a duplicate request ID for at least 60 minutes after the first request.For example, if an initial request times out, followed by another request with the same request ID, the server ignores the second request to prevent the creation of duplicate commitments.The request ID must be a valid UUID (https://en.wikipedia.org/wiki/Universally_unique_identifier#Format) A zero UUID (00000000-0000-0000-0000-000000000000) is not supported.
	RequestId *string `pulumi:"requestId"`
	// Required. The ID of the metastore service, which is used as the final component of the metastore service's name.This value must be between 2 and 63 characters long inclusive, begin with a letter, end with a letter or number, and consist of alpha-numeric ASCII characters or hyphens.
	ServiceId string `pulumi:"serviceId"`
	// The tier of the service.
	Tier *ServiceTier `pulumi:"tier"`
}

// The set of arguments for constructing a Service resource.
type ServiceArgs struct {
	// Immutable. The database type that the Metastore service stores its data.
	DatabaseType ServiceDatabaseTypePtrInput
	// Immutable. Information used to configure the Dataproc Metastore service to encrypt customer data at rest. Cannot be updated.
	EncryptionConfig EncryptionConfigPtrInput
	// Configuration information specific to running Hive metastore software as the metastore service.
	HiveMetastoreConfig HiveMetastoreConfigPtrInput
	// User-defined labels for the metastore service.
	Labels   pulumi.StringMapInput
	Location pulumi.StringPtrInput
	// The one hour maintenance window of the metastore service. This specifies when the service can be restarted for maintenance purposes in UTC time. Maintenance window is not needed for services with the SPANNER database type.
	MaintenanceWindow MaintenanceWindowPtrInput
	// The setting that defines how metastore metadata should be integrated with external services and systems.
	MetadataIntegration MetadataIntegrationPtrInput
	// Immutable. The relative resource name of the metastore service, of the form:projects/{project_number}/locations/{location_id}/services/{service_id}.
	Name pulumi.StringPtrInput
	// Immutable. The relative resource name of the VPC network on which the instance can be accessed. It is specified in the following form:projects/{project_number}/global/networks/{network_id}.
	Network pulumi.StringPtrInput
	// Immutable. The configuration specifying the network settings for the Dataproc Metastore service.
	NetworkConfig NetworkConfigPtrInput
	// The TCP port at which the metastore service is reached. Default: 9083.
	Port    pulumi.IntPtrInput
	Project pulumi.StringPtrInput
	// Immutable. The release channel of the service. If unspecified, defaults to STABLE.
	ReleaseChannel ServiceReleaseChannelPtrInput
	// Optional. A request ID. Specify a unique request ID to allow the server to ignore the request if it has completed. The server will ignore subsequent requests that provide a duplicate request ID for at least 60 minutes after the first request.For example, if an initial request times out, followed by another request with the same request ID, the server ignores the second request to prevent the creation of duplicate commitments.The request ID must be a valid UUID (https://en.wikipedia.org/wiki/Universally_unique_identifier#Format) A zero UUID (00000000-0000-0000-0000-000000000000) is not supported.
	RequestId pulumi.StringPtrInput
	// Required. The ID of the metastore service, which is used as the final component of the metastore service's name.This value must be between 2 and 63 characters long inclusive, begin with a letter, end with a letter or number, and consist of alpha-numeric ASCII characters or hyphens.
	ServiceId pulumi.StringInput
	// The tier of the service.
	Tier ServiceTierPtrInput
}

func (ServiceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceArgs)(nil)).Elem()
}

type ServiceInput interface {
	pulumi.Input

	ToServiceOutput() ServiceOutput
	ToServiceOutputWithContext(ctx context.Context) ServiceOutput
}

func (*Service) ElementType() reflect.Type {
	return reflect.TypeOf((**Service)(nil)).Elem()
}

func (i *Service) ToServiceOutput() ServiceOutput {
	return i.ToServiceOutputWithContext(context.Background())
}

func (i *Service) ToServiceOutputWithContext(ctx context.Context) ServiceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceOutput)
}

type ServiceOutput struct{ *pulumi.OutputState }

func (ServiceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Service)(nil)).Elem()
}

func (o ServiceOutput) ToServiceOutput() ServiceOutput {
	return o
}

func (o ServiceOutput) ToServiceOutputWithContext(ctx context.Context) ServiceOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceInput)(nil)).Elem(), &Service{})
	pulumi.RegisterOutputType(ServiceOutput{})
}
