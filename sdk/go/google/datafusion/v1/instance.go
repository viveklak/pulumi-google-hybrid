// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a new Data Fusion instance in the specified project and location.
type Instance struct {
	pulumi.CustomResourceState

	// List of accelerators enabled for this CDF instance.
	Accelerators AcceleratorResponseArrayOutput `pulumi:"accelerators"`
	// Endpoint on which the REST APIs is accessible.
	ApiEndpoint pulumi.StringOutput `pulumi:"apiEndpoint"`
	// Available versions that the instance can be upgraded to using UpdateInstanceRequest.
	AvailableVersion VersionResponseArrayOutput `pulumi:"availableVersion"`
	// The time the instance was created.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// User-managed service account to set on Dataproc when Cloud Data Fusion creates Dataproc to run data processing pipelines. This allows users to have fine-grained access control on Dataproc's accesses to cloud resources.
	DataprocServiceAccount pulumi.StringOutput `pulumi:"dataprocServiceAccount"`
	// A description of this instance.
	Description pulumi.StringOutput `pulumi:"description"`
	// Display name for an instance.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// Option to enable granular role-based access control.
	EnableRbac pulumi.BoolOutput `pulumi:"enableRbac"`
	// Option to enable Stackdriver Logging.
	EnableStackdriverLogging pulumi.BoolOutput `pulumi:"enableStackdriverLogging"`
	// Option to enable Stackdriver Monitoring.
	EnableStackdriverMonitoring pulumi.BoolOutput `pulumi:"enableStackdriverMonitoring"`
	// Cloud Storage bucket generated by Data Fusion in the customer project.
	GcsBucket pulumi.StringOutput `pulumi:"gcsBucket"`
	// The resource labels for instance to use to annotate any related underlying resources such as GCE VMs. The character '=' is not allowed to be used within the labels.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// The name of this instance is in the form of projects/{project}/locations/{location}/instances/{instance}.
	Name pulumi.StringOutput `pulumi:"name"`
	// Network configuration options. These are required when a private Data Fusion instance is to be created.
	NetworkConfig NetworkConfigResponseOutput `pulumi:"networkConfig"`
	// Map of additional options used to configure the behavior of Data Fusion instance.
	Options pulumi.StringMapOutput `pulumi:"options"`
	// P4 service account for the customer project.
	P4ServiceAccount pulumi.StringOutput `pulumi:"p4ServiceAccount"`
	// Specifies whether the Data Fusion instance should be private. If set to true, all Data Fusion nodes will have private IP addresses and will not be able to access the public internet.
	PrivateInstance pulumi.BoolOutput `pulumi:"privateInstance"`
	// Endpoint on which the Data Fusion UI is accessible.
	ServiceEndpoint pulumi.StringOutput `pulumi:"serviceEndpoint"`
	// The current state of this Data Fusion instance.
	State pulumi.StringOutput `pulumi:"state"`
	// Additional information about the current state of this Data Fusion instance if available.
	StateMessage pulumi.StringOutput `pulumi:"stateMessage"`
	// The name of the tenant project.
	TenantProjectId pulumi.StringOutput `pulumi:"tenantProjectId"`
	// Instance type.
	Type pulumi.StringOutput `pulumi:"type"`
	// The time the instance was last updated.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
	// Current version of the Data Fusion. Only specifiable in Update.
	Version pulumi.StringOutput `pulumi:"version"`
	// Name of the zone in which the Data Fusion instance will be created. Only DEVELOPER instances use this field.
	Zone pulumi.StringOutput `pulumi:"zone"`
}

// NewInstance registers a new resource with the given unique name, arguments, and options.
func NewInstance(ctx *pulumi.Context,
	name string, args *InstanceArgs, opts ...pulumi.ResourceOption) (*Instance, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Location == nil {
		return nil, errors.New("invalid value for required argument 'Location'")
	}
	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	var resource Instance
	err := ctx.RegisterResource("google-native:datafusion/v1:Instance", name, args, &resource, opts...)
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
	err := ctx.ReadResource("google-native:datafusion/v1:Instance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Instance resources.
type instanceState struct {
	// List of accelerators enabled for this CDF instance.
	Accelerators []AcceleratorResponse `pulumi:"accelerators"`
	// Endpoint on which the REST APIs is accessible.
	ApiEndpoint *string `pulumi:"apiEndpoint"`
	// Available versions that the instance can be upgraded to using UpdateInstanceRequest.
	AvailableVersion []VersionResponse `pulumi:"availableVersion"`
	// The time the instance was created.
	CreateTime *string `pulumi:"createTime"`
	// User-managed service account to set on Dataproc when Cloud Data Fusion creates Dataproc to run data processing pipelines. This allows users to have fine-grained access control on Dataproc's accesses to cloud resources.
	DataprocServiceAccount *string `pulumi:"dataprocServiceAccount"`
	// A description of this instance.
	Description *string `pulumi:"description"`
	// Display name for an instance.
	DisplayName *string `pulumi:"displayName"`
	// Option to enable granular role-based access control.
	EnableRbac *bool `pulumi:"enableRbac"`
	// Option to enable Stackdriver Logging.
	EnableStackdriverLogging *bool `pulumi:"enableStackdriverLogging"`
	// Option to enable Stackdriver Monitoring.
	EnableStackdriverMonitoring *bool `pulumi:"enableStackdriverMonitoring"`
	// Cloud Storage bucket generated by Data Fusion in the customer project.
	GcsBucket *string `pulumi:"gcsBucket"`
	// The resource labels for instance to use to annotate any related underlying resources such as GCE VMs. The character '=' is not allowed to be used within the labels.
	Labels map[string]string `pulumi:"labels"`
	// The name of this instance is in the form of projects/{project}/locations/{location}/instances/{instance}.
	Name *string `pulumi:"name"`
	// Network configuration options. These are required when a private Data Fusion instance is to be created.
	NetworkConfig *NetworkConfigResponse `pulumi:"networkConfig"`
	// Map of additional options used to configure the behavior of Data Fusion instance.
	Options map[string]string `pulumi:"options"`
	// P4 service account for the customer project.
	P4ServiceAccount *string `pulumi:"p4ServiceAccount"`
	// Specifies whether the Data Fusion instance should be private. If set to true, all Data Fusion nodes will have private IP addresses and will not be able to access the public internet.
	PrivateInstance *bool `pulumi:"privateInstance"`
	// Endpoint on which the Data Fusion UI is accessible.
	ServiceEndpoint *string `pulumi:"serviceEndpoint"`
	// The current state of this Data Fusion instance.
	State *string `pulumi:"state"`
	// Additional information about the current state of this Data Fusion instance if available.
	StateMessage *string `pulumi:"stateMessage"`
	// The name of the tenant project.
	TenantProjectId *string `pulumi:"tenantProjectId"`
	// Instance type.
	Type *string `pulumi:"type"`
	// The time the instance was last updated.
	UpdateTime *string `pulumi:"updateTime"`
	// Current version of the Data Fusion. Only specifiable in Update.
	Version *string `pulumi:"version"`
	// Name of the zone in which the Data Fusion instance will be created. Only DEVELOPER instances use this field.
	Zone *string `pulumi:"zone"`
}

type InstanceState struct {
	// List of accelerators enabled for this CDF instance.
	Accelerators AcceleratorResponseArrayInput
	// Endpoint on which the REST APIs is accessible.
	ApiEndpoint pulumi.StringPtrInput
	// Available versions that the instance can be upgraded to using UpdateInstanceRequest.
	AvailableVersion VersionResponseArrayInput
	// The time the instance was created.
	CreateTime pulumi.StringPtrInput
	// User-managed service account to set on Dataproc when Cloud Data Fusion creates Dataproc to run data processing pipelines. This allows users to have fine-grained access control on Dataproc's accesses to cloud resources.
	DataprocServiceAccount pulumi.StringPtrInput
	// A description of this instance.
	Description pulumi.StringPtrInput
	// Display name for an instance.
	DisplayName pulumi.StringPtrInput
	// Option to enable granular role-based access control.
	EnableRbac pulumi.BoolPtrInput
	// Option to enable Stackdriver Logging.
	EnableStackdriverLogging pulumi.BoolPtrInput
	// Option to enable Stackdriver Monitoring.
	EnableStackdriverMonitoring pulumi.BoolPtrInput
	// Cloud Storage bucket generated by Data Fusion in the customer project.
	GcsBucket pulumi.StringPtrInput
	// The resource labels for instance to use to annotate any related underlying resources such as GCE VMs. The character '=' is not allowed to be used within the labels.
	Labels pulumi.StringMapInput
	// The name of this instance is in the form of projects/{project}/locations/{location}/instances/{instance}.
	Name pulumi.StringPtrInput
	// Network configuration options. These are required when a private Data Fusion instance is to be created.
	NetworkConfig NetworkConfigResponsePtrInput
	// Map of additional options used to configure the behavior of Data Fusion instance.
	Options pulumi.StringMapInput
	// P4 service account for the customer project.
	P4ServiceAccount pulumi.StringPtrInput
	// Specifies whether the Data Fusion instance should be private. If set to true, all Data Fusion nodes will have private IP addresses and will not be able to access the public internet.
	PrivateInstance pulumi.BoolPtrInput
	// Endpoint on which the Data Fusion UI is accessible.
	ServiceEndpoint pulumi.StringPtrInput
	// The current state of this Data Fusion instance.
	State pulumi.StringPtrInput
	// Additional information about the current state of this Data Fusion instance if available.
	StateMessage pulumi.StringPtrInput
	// The name of the tenant project.
	TenantProjectId pulumi.StringPtrInput
	// Instance type.
	Type pulumi.StringPtrInput
	// The time the instance was last updated.
	UpdateTime pulumi.StringPtrInput
	// Current version of the Data Fusion. Only specifiable in Update.
	Version pulumi.StringPtrInput
	// Name of the zone in which the Data Fusion instance will be created. Only DEVELOPER instances use this field.
	Zone pulumi.StringPtrInput
}

func (InstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceState)(nil)).Elem()
}

type instanceArgs struct {
	// List of accelerators enabled for this CDF instance.
	Accelerators []Accelerator `pulumi:"accelerators"`
	// Available versions that the instance can be upgraded to using UpdateInstanceRequest.
	AvailableVersion []Version `pulumi:"availableVersion"`
	// User-managed service account to set on Dataproc when Cloud Data Fusion creates Dataproc to run data processing pipelines. This allows users to have fine-grained access control on Dataproc's accesses to cloud resources.
	DataprocServiceAccount *string `pulumi:"dataprocServiceAccount"`
	// A description of this instance.
	Description *string `pulumi:"description"`
	// Display name for an instance.
	DisplayName *string `pulumi:"displayName"`
	// Option to enable granular role-based access control.
	EnableRbac *bool `pulumi:"enableRbac"`
	// Option to enable Stackdriver Logging.
	EnableStackdriverLogging *bool `pulumi:"enableStackdriverLogging"`
	// Option to enable Stackdriver Monitoring.
	EnableStackdriverMonitoring *bool   `pulumi:"enableStackdriverMonitoring"`
	InstanceId                  *string `pulumi:"instanceId"`
	// The resource labels for instance to use to annotate any related underlying resources such as GCE VMs. The character '=' is not allowed to be used within the labels.
	Labels   map[string]string `pulumi:"labels"`
	Location string            `pulumi:"location"`
	// Network configuration options. These are required when a private Data Fusion instance is to be created.
	NetworkConfig *NetworkConfig `pulumi:"networkConfig"`
	// Map of additional options used to configure the behavior of Data Fusion instance.
	Options map[string]string `pulumi:"options"`
	// Specifies whether the Data Fusion instance should be private. If set to true, all Data Fusion nodes will have private IP addresses and will not be able to access the public internet.
	PrivateInstance *bool  `pulumi:"privateInstance"`
	Project         string `pulumi:"project"`
	// Instance type.
	Type string `pulumi:"type"`
	// Current version of the Data Fusion. Only specifiable in Update.
	Version *string `pulumi:"version"`
	// Name of the zone in which the Data Fusion instance will be created. Only DEVELOPER instances use this field.
	Zone *string `pulumi:"zone"`
}

// The set of arguments for constructing a Instance resource.
type InstanceArgs struct {
	// List of accelerators enabled for this CDF instance.
	Accelerators AcceleratorArrayInput
	// Available versions that the instance can be upgraded to using UpdateInstanceRequest.
	AvailableVersion VersionArrayInput
	// User-managed service account to set on Dataproc when Cloud Data Fusion creates Dataproc to run data processing pipelines. This allows users to have fine-grained access control on Dataproc's accesses to cloud resources.
	DataprocServiceAccount pulumi.StringPtrInput
	// A description of this instance.
	Description pulumi.StringPtrInput
	// Display name for an instance.
	DisplayName pulumi.StringPtrInput
	// Option to enable granular role-based access control.
	EnableRbac pulumi.BoolPtrInput
	// Option to enable Stackdriver Logging.
	EnableStackdriverLogging pulumi.BoolPtrInput
	// Option to enable Stackdriver Monitoring.
	EnableStackdriverMonitoring pulumi.BoolPtrInput
	InstanceId                  pulumi.StringPtrInput
	// The resource labels for instance to use to annotate any related underlying resources such as GCE VMs. The character '=' is not allowed to be used within the labels.
	Labels   pulumi.StringMapInput
	Location pulumi.StringInput
	// Network configuration options. These are required when a private Data Fusion instance is to be created.
	NetworkConfig NetworkConfigPtrInput
	// Map of additional options used to configure the behavior of Data Fusion instance.
	Options pulumi.StringMapInput
	// Specifies whether the Data Fusion instance should be private. If set to true, all Data Fusion nodes will have private IP addresses and will not be able to access the public internet.
	PrivateInstance pulumi.BoolPtrInput
	Project         pulumi.StringInput
	// Instance type.
	Type InstanceType
	// Current version of the Data Fusion. Only specifiable in Update.
	Version pulumi.StringPtrInput
	// Name of the zone in which the Data Fusion instance will be created. Only DEVELOPER instances use this field.
	Zone pulumi.StringPtrInput
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
