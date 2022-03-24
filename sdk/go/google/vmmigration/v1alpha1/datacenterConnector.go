// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1alpha1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a new DatacenterConnector in a given Source.
// Auto-naming is currently not supported for this resource.
type DatacenterConnector struct {
	pulumi.CustomResourceState

	// Appliance OVA version. This is the OVA which is manually installed by the user and contains the infrastructure for the automatically updatable components on the appliance.
	ApplianceInfrastructureVersion pulumi.StringOutput `pulumi:"applianceInfrastructureVersion"`
	// Appliance last installed update bundle version. This is the version of the automatically updatable components on the appliance.
	ApplianceSoftwareVersion pulumi.StringOutput `pulumi:"applianceSoftwareVersion"`
	// The available versions for updating this appliance.
	AvailableVersions AvailableUpdatesResponseOutput `pulumi:"availableVersions"`
	// The communication channel between the datacenter connector and GCP.
	Bucket pulumi.StringOutput `pulumi:"bucket"`
	// The time the connector was created (as an API call, not when it was actually installed).
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Provides details on the state of the Datacenter Connector in case of an error.
	Error StatusResponseOutput `pulumi:"error"`
	// The connector's name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Immutable. A unique key for this connector. This key is internal to the OVA connector and is supplied with its creation during the registration process and can not be modified.
	RegistrationId pulumi.StringOutput `pulumi:"registrationId"`
	// The service account to use in the connector when communicating with the cloud.
	ServiceAccount pulumi.StringOutput `pulumi:"serviceAccount"`
	// State of the DatacenterConnector, as determined by the health checks.
	State pulumi.StringOutput `pulumi:"state"`
	// The time the state was last set.
	StateTime pulumi.StringOutput `pulumi:"stateTime"`
	// The last time the connector was updated with an API call.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
	// The status of the current / last upgradeAppliance operation.
	UpgradeStatus UpgradeStatusResponseOutput `pulumi:"upgradeStatus"`
	// The version running in the DatacenterConnector. This is supplied by the OVA connector during the registration process and can not be modified.
	Version pulumi.StringOutput `pulumi:"version"`
}

// NewDatacenterConnector registers a new resource with the given unique name, arguments, and options.
func NewDatacenterConnector(ctx *pulumi.Context,
	name string, args *DatacenterConnectorArgs, opts ...pulumi.ResourceOption) (*DatacenterConnector, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DatacenterConnectorId == nil {
		return nil, errors.New("invalid value for required argument 'DatacenterConnectorId'")
	}
	if args.SourceId == nil {
		return nil, errors.New("invalid value for required argument 'SourceId'")
	}
	var resource DatacenterConnector
	err := ctx.RegisterResource("google-native:vmmigration/v1alpha1:DatacenterConnector", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDatacenterConnector gets an existing DatacenterConnector resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDatacenterConnector(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DatacenterConnectorState, opts ...pulumi.ResourceOption) (*DatacenterConnector, error) {
	var resource DatacenterConnector
	err := ctx.ReadResource("google-native:vmmigration/v1alpha1:DatacenterConnector", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DatacenterConnector resources.
type datacenterConnectorState struct {
}

type DatacenterConnectorState struct {
}

func (DatacenterConnectorState) ElementType() reflect.Type {
	return reflect.TypeOf((*datacenterConnectorState)(nil)).Elem()
}

type datacenterConnectorArgs struct {
	// Required. The datacenterConnector identifier.
	DatacenterConnectorId string  `pulumi:"datacenterConnectorId"`
	Location              *string `pulumi:"location"`
	Project               *string `pulumi:"project"`
	// Immutable. A unique key for this connector. This key is internal to the OVA connector and is supplied with its creation during the registration process and can not be modified.
	RegistrationId *string `pulumi:"registrationId"`
	// A request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. The server will guarantee that for at least 60 minutes since the first request. For example, consider a situation where you make an initial request and t he request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported (00000000-0000-0000-0000-000000000000).
	RequestId *string `pulumi:"requestId"`
	// The service account to use in the connector when communicating with the cloud.
	ServiceAccount *string `pulumi:"serviceAccount"`
	SourceId       string  `pulumi:"sourceId"`
	// The version running in the DatacenterConnector. This is supplied by the OVA connector during the registration process and can not be modified.
	Version *string `pulumi:"version"`
}

// The set of arguments for constructing a DatacenterConnector resource.
type DatacenterConnectorArgs struct {
	// Required. The datacenterConnector identifier.
	DatacenterConnectorId pulumi.StringInput
	Location              pulumi.StringPtrInput
	Project               pulumi.StringPtrInput
	// Immutable. A unique key for this connector. This key is internal to the OVA connector and is supplied with its creation during the registration process and can not be modified.
	RegistrationId pulumi.StringPtrInput
	// A request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. The server will guarantee that for at least 60 minutes since the first request. For example, consider a situation where you make an initial request and t he request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported (00000000-0000-0000-0000-000000000000).
	RequestId pulumi.StringPtrInput
	// The service account to use in the connector when communicating with the cloud.
	ServiceAccount pulumi.StringPtrInput
	SourceId       pulumi.StringInput
	// The version running in the DatacenterConnector. This is supplied by the OVA connector during the registration process and can not be modified.
	Version pulumi.StringPtrInput
}

func (DatacenterConnectorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*datacenterConnectorArgs)(nil)).Elem()
}

type DatacenterConnectorInput interface {
	pulumi.Input

	ToDatacenterConnectorOutput() DatacenterConnectorOutput
	ToDatacenterConnectorOutputWithContext(ctx context.Context) DatacenterConnectorOutput
}

func (*DatacenterConnector) ElementType() reflect.Type {
	return reflect.TypeOf((**DatacenterConnector)(nil)).Elem()
}

func (i *DatacenterConnector) ToDatacenterConnectorOutput() DatacenterConnectorOutput {
	return i.ToDatacenterConnectorOutputWithContext(context.Background())
}

func (i *DatacenterConnector) ToDatacenterConnectorOutputWithContext(ctx context.Context) DatacenterConnectorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatacenterConnectorOutput)
}

type DatacenterConnectorOutput struct{ *pulumi.OutputState }

func (DatacenterConnectorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DatacenterConnector)(nil)).Elem()
}

func (o DatacenterConnectorOutput) ToDatacenterConnectorOutput() DatacenterConnectorOutput {
	return o
}

func (o DatacenterConnectorOutput) ToDatacenterConnectorOutputWithContext(ctx context.Context) DatacenterConnectorOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DatacenterConnectorInput)(nil)).Elem(), &DatacenterConnector{})
	pulumi.RegisterOutputType(DatacenterConnectorOutput{})
}