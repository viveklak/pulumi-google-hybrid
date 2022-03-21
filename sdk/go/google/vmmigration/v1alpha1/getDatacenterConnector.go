// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1alpha1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets details of a single DatacenterConnector.
func LookupDatacenterConnector(ctx *pulumi.Context, args *LookupDatacenterConnectorArgs, opts ...pulumi.InvokeOption) (*LookupDatacenterConnectorResult, error) {
	var rv LookupDatacenterConnectorResult
	err := ctx.Invoke("google-native:vmmigration/v1alpha1:getDatacenterConnector", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDatacenterConnectorArgs struct {
	DatacenterConnectorId string  `pulumi:"datacenterConnectorId"`
	Location              string  `pulumi:"location"`
	Project               *string `pulumi:"project"`
	SourceId              string  `pulumi:"sourceId"`
}

type LookupDatacenterConnectorResult struct {
	// Appliance OVA version. This is the OVA which is manually installed by the user and contains the infrastructure for the automatically updatable components on the appliance.
	ApplianceInfrastructureVersion string `pulumi:"applianceInfrastructureVersion"`
	// Appliance last installed update bundle version. This is the version of the automatically updatable components on the appliance.
	ApplianceSoftwareVersion string `pulumi:"applianceSoftwareVersion"`
	// The available versions for updating this appliance.
	AvailableVersions AvailableUpdatesResponse `pulumi:"availableVersions"`
	// The communication channel between the datacenter connector and GCP.
	Bucket string `pulumi:"bucket"`
	// The time the connector was created (as an API call, not when it was actually installed).
	CreateTime string `pulumi:"createTime"`
	// Provides details on the state of the Datacenter Connector in case of an error.
	Error StatusResponse `pulumi:"error"`
	// The connector's name.
	Name string `pulumi:"name"`
	// Immutable. A unique key for this connector. This key is internal to the OVA connector and is supplied with its creation during the registration process and can not be modified.
	RegistrationId string `pulumi:"registrationId"`
	// The service account to use in the connector when communicating with the cloud.
	ServiceAccount string `pulumi:"serviceAccount"`
	// State of the DatacenterConnector, as determined by the health checks.
	State string `pulumi:"state"`
	// The time the state was last set.
	StateTime string `pulumi:"stateTime"`
	// The last time the connector was updated with an API call.
	UpdateTime string `pulumi:"updateTime"`
	// The status of the current / last upgradeAppliance operation.
	UpgradeStatus UpgradeStatusResponse `pulumi:"upgradeStatus"`
	// The version running in the DatacenterConnector. This is supplied by the OVA connector during the registration process and can not be modified.
	Version string `pulumi:"version"`
}

func LookupDatacenterConnectorOutput(ctx *pulumi.Context, args LookupDatacenterConnectorOutputArgs, opts ...pulumi.InvokeOption) LookupDatacenterConnectorResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDatacenterConnectorResult, error) {
			args := v.(LookupDatacenterConnectorArgs)
			r, err := LookupDatacenterConnector(ctx, &args, opts...)
			return *r, err
		}).(LookupDatacenterConnectorResultOutput)
}

type LookupDatacenterConnectorOutputArgs struct {
	DatacenterConnectorId pulumi.StringInput    `pulumi:"datacenterConnectorId"`
	Location              pulumi.StringInput    `pulumi:"location"`
	Project               pulumi.StringPtrInput `pulumi:"project"`
	SourceId              pulumi.StringInput    `pulumi:"sourceId"`
}

func (LookupDatacenterConnectorOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDatacenterConnectorArgs)(nil)).Elem()
}

type LookupDatacenterConnectorResultOutput struct{ *pulumi.OutputState }

func (LookupDatacenterConnectorResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDatacenterConnectorResult)(nil)).Elem()
}

func (o LookupDatacenterConnectorResultOutput) ToLookupDatacenterConnectorResultOutput() LookupDatacenterConnectorResultOutput {
	return o
}

func (o LookupDatacenterConnectorResultOutput) ToLookupDatacenterConnectorResultOutputWithContext(ctx context.Context) LookupDatacenterConnectorResultOutput {
	return o
}

// Appliance OVA version. This is the OVA which is manually installed by the user and contains the infrastructure for the automatically updatable components on the appliance.
func (o LookupDatacenterConnectorResultOutput) ApplianceInfrastructureVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatacenterConnectorResult) string { return v.ApplianceInfrastructureVersion }).(pulumi.StringOutput)
}

// Appliance last installed update bundle version. This is the version of the automatically updatable components on the appliance.
func (o LookupDatacenterConnectorResultOutput) ApplianceSoftwareVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatacenterConnectorResult) string { return v.ApplianceSoftwareVersion }).(pulumi.StringOutput)
}

// The available versions for updating this appliance.
func (o LookupDatacenterConnectorResultOutput) AvailableVersions() AvailableUpdatesResponseOutput {
	return o.ApplyT(func(v LookupDatacenterConnectorResult) AvailableUpdatesResponse { return v.AvailableVersions }).(AvailableUpdatesResponseOutput)
}

// The communication channel between the datacenter connector and GCP.
func (o LookupDatacenterConnectorResultOutput) Bucket() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatacenterConnectorResult) string { return v.Bucket }).(pulumi.StringOutput)
}

// The time the connector was created (as an API call, not when it was actually installed).
func (o LookupDatacenterConnectorResultOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatacenterConnectorResult) string { return v.CreateTime }).(pulumi.StringOutput)
}

// Provides details on the state of the Datacenter Connector in case of an error.
func (o LookupDatacenterConnectorResultOutput) Error() StatusResponseOutput {
	return o.ApplyT(func(v LookupDatacenterConnectorResult) StatusResponse { return v.Error }).(StatusResponseOutput)
}

// The connector's name.
func (o LookupDatacenterConnectorResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatacenterConnectorResult) string { return v.Name }).(pulumi.StringOutput)
}

// Immutable. A unique key for this connector. This key is internal to the OVA connector and is supplied with its creation during the registration process and can not be modified.
func (o LookupDatacenterConnectorResultOutput) RegistrationId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatacenterConnectorResult) string { return v.RegistrationId }).(pulumi.StringOutput)
}

// The service account to use in the connector when communicating with the cloud.
func (o LookupDatacenterConnectorResultOutput) ServiceAccount() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatacenterConnectorResult) string { return v.ServiceAccount }).(pulumi.StringOutput)
}

// State of the DatacenterConnector, as determined by the health checks.
func (o LookupDatacenterConnectorResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatacenterConnectorResult) string { return v.State }).(pulumi.StringOutput)
}

// The time the state was last set.
func (o LookupDatacenterConnectorResultOutput) StateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatacenterConnectorResult) string { return v.StateTime }).(pulumi.StringOutput)
}

// The last time the connector was updated with an API call.
func (o LookupDatacenterConnectorResultOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatacenterConnectorResult) string { return v.UpdateTime }).(pulumi.StringOutput)
}

// The status of the current / last upgradeAppliance operation.
func (o LookupDatacenterConnectorResultOutput) UpgradeStatus() UpgradeStatusResponseOutput {
	return o.ApplyT(func(v LookupDatacenterConnectorResult) UpgradeStatusResponse { return v.UpgradeStatus }).(UpgradeStatusResponseOutput)
}

// The version running in the DatacenterConnector. This is supplied by the OVA connector during the registration process and can not be modified.
func (o LookupDatacenterConnectorResultOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatacenterConnectorResult) string { return v.Version }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDatacenterConnectorResultOutput{})
}
