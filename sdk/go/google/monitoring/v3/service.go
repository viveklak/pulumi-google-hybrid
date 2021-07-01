// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v3

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Create a Service.
type Service struct {
	pulumi.CustomResourceState

	// Type used for App Engine services.
	AppEngine AppEngineResponseOutput `pulumi:"appEngine"`
	// Type used for Cloud Endpoints services.
	CloudEndpoints CloudEndpointsResponseOutput `pulumi:"cloudEndpoints"`
	// Type used for Istio services that live in a Kubernetes cluster.
	ClusterIstio ClusterIstioResponseOutput `pulumi:"clusterIstio"`
	// Custom service type.
	Custom CustomResponseOutput `pulumi:"custom"`
	// Name used for UI elements listing this Service.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// Type used for canonical services scoped to an Istio mesh. Metrics for Istio are documented here (https://istio.io/latest/docs/reference/config/metrics/)
	IstioCanonicalService IstioCanonicalServiceResponseOutput `pulumi:"istioCanonicalService"`
	// Type used for Istio services scoped to an Istio mesh.
	MeshIstio MeshIstioResponseOutput `pulumi:"meshIstio"`
	// Resource name for this Service. The format is: projects/[PROJECT_ID_OR_NUMBER]/services/[SERVICE_ID]
	Name pulumi.StringOutput `pulumi:"name"`
	// Configuration for how to query telemetry on a Service.
	Telemetry TelemetryResponseOutput `pulumi:"telemetry"`
}

// NewService registers a new resource with the given unique name, arguments, and options.
func NewService(ctx *pulumi.Context,
	name string, args *ServiceArgs, opts ...pulumi.ResourceOption) (*Service, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.V3Id == nil {
		return nil, errors.New("invalid value for required argument 'V3Id'")
	}
	if args.V3Id1 == nil {
		return nil, errors.New("invalid value for required argument 'V3Id1'")
	}
	var resource Service
	err := ctx.RegisterResource("google-native:monitoring/v3:Service", name, args, &resource, opts...)
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
	err := ctx.ReadResource("google-native:monitoring/v3:Service", name, id, state, &resource, opts...)
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
	// Type used for App Engine services.
	AppEngine *AppEngine `pulumi:"appEngine"`
	// Type used for Cloud Endpoints services.
	CloudEndpoints *CloudEndpoints `pulumi:"cloudEndpoints"`
	// Type used for Istio services that live in a Kubernetes cluster.
	ClusterIstio *ClusterIstio `pulumi:"clusterIstio"`
	// Custom service type.
	Custom *Custom `pulumi:"custom"`
	// Name used for UI elements listing this Service.
	DisplayName *string `pulumi:"displayName"`
	// Type used for canonical services scoped to an Istio mesh. Metrics for Istio are documented here (https://istio.io/latest/docs/reference/config/metrics/)
	IstioCanonicalService *IstioCanonicalService `pulumi:"istioCanonicalService"`
	// Type used for Istio services scoped to an Istio mesh.
	MeshIstio *MeshIstio `pulumi:"meshIstio"`
	// Resource name for this Service. The format is: projects/[PROJECT_ID_OR_NUMBER]/services/[SERVICE_ID]
	Name      *string `pulumi:"name"`
	ServiceId *string `pulumi:"serviceId"`
	// Configuration for how to query telemetry on a Service.
	Telemetry *Telemetry `pulumi:"telemetry"`
	V3Id      string     `pulumi:"v3Id"`
	V3Id1     string     `pulumi:"v3Id1"`
}

// The set of arguments for constructing a Service resource.
type ServiceArgs struct {
	// Type used for App Engine services.
	AppEngine AppEnginePtrInput
	// Type used for Cloud Endpoints services.
	CloudEndpoints CloudEndpointsPtrInput
	// Type used for Istio services that live in a Kubernetes cluster.
	ClusterIstio ClusterIstioPtrInput
	// Custom service type.
	Custom CustomPtrInput
	// Name used for UI elements listing this Service.
	DisplayName pulumi.StringPtrInput
	// Type used for canonical services scoped to an Istio mesh. Metrics for Istio are documented here (https://istio.io/latest/docs/reference/config/metrics/)
	IstioCanonicalService IstioCanonicalServicePtrInput
	// Type used for Istio services scoped to an Istio mesh.
	MeshIstio MeshIstioPtrInput
	// Resource name for this Service. The format is: projects/[PROJECT_ID_OR_NUMBER]/services/[SERVICE_ID]
	Name      pulumi.StringPtrInput
	ServiceId pulumi.StringPtrInput
	// Configuration for how to query telemetry on a Service.
	Telemetry TelemetryPtrInput
	V3Id      pulumi.StringInput
	V3Id1     pulumi.StringInput
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
	return reflect.TypeOf((*Service)(nil))
}

func (i *Service) ToServiceOutput() ServiceOutput {
	return i.ToServiceOutputWithContext(context.Background())
}

func (i *Service) ToServiceOutputWithContext(ctx context.Context) ServiceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceOutput)
}

type ServiceOutput struct {
	*pulumi.OutputState
}

func (ServiceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Service)(nil))
}

func (o ServiceOutput) ToServiceOutput() ServiceOutput {
	return o
}

func (o ServiceOutput) ToServiceOutputWithContext(ctx context.Context) ServiceOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ServiceOutput{})
}
