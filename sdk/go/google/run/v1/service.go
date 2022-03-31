// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Create a service.
// Auto-naming is currently not supported for this resource.
type Service struct {
	pulumi.CustomResourceState

	// The API version for this call such as "serving.knative.dev/v1".
	ApiVersion pulumi.StringOutput `pulumi:"apiVersion"`
	// The kind of resource, in this case "Service".
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Metadata associated with this Service, including name, namespace, labels, and annotations. Cloud Run (fully managed) uses the following annotation keys to configure features on a Service: * `run.googleapis.com/ingress` sets the ingress settings for the Service. See [the ingress settings documentation](/run/docs/securing/ingress) for details on configuring ingress settings. * `run.googleapis.com/ingress-status` is output-only and contains the currently active ingress settings for the Service. `run.googleapis.com/ingress-status` may differ from `run.googleapis.com/ingress` while the system is processing a change to `run.googleapis.com/ingress` or if the system failed to process a change to `run.googleapis.com/ingress`. When the system has processed all changes successfully `run.googleapis.com/ingress-status` and `run.googleapis.com/ingress` are equal.
	Metadata ObjectMetaResponseOutput `pulumi:"metadata"`
	// Spec holds the desired state of the Service (from the client).
	Spec ServiceSpecResponseOutput `pulumi:"spec"`
	// Status communicates the observed state of the Service (from the controller).
	Status ServiceStatusResponseOutput `pulumi:"status"`
}

// NewService registers a new resource with the given unique name, arguments, and options.
func NewService(ctx *pulumi.Context,
	name string, args *ServiceArgs, opts ...pulumi.ResourceOption) (*Service, error) {
	if args == nil {
		args = &ServiceArgs{}
	}

	var resource Service
	err := ctx.RegisterResource("google-native:run/v1:Service", name, args, &resource, opts...)
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
	err := ctx.ReadResource("google-native:run/v1:Service", name, id, state, &resource, opts...)
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
	// The API version for this call such as "serving.knative.dev/v1".
	ApiVersion *string `pulumi:"apiVersion"`
	// Indicates that the server should validate the request and populate default values without persisting the request. Supported values: `all`
	DryRun *string `pulumi:"dryRun"`
	// The kind of resource, in this case "Service".
	Kind     *string `pulumi:"kind"`
	Location *string `pulumi:"location"`
	// Metadata associated with this Service, including name, namespace, labels, and annotations. Cloud Run (fully managed) uses the following annotation keys to configure features on a Service: * `run.googleapis.com/ingress` sets the ingress settings for the Service. See [the ingress settings documentation](/run/docs/securing/ingress) for details on configuring ingress settings. * `run.googleapis.com/ingress-status` is output-only and contains the currently active ingress settings for the Service. `run.googleapis.com/ingress-status` may differ from `run.googleapis.com/ingress` while the system is processing a change to `run.googleapis.com/ingress` or if the system failed to process a change to `run.googleapis.com/ingress`. When the system has processed all changes successfully `run.googleapis.com/ingress-status` and `run.googleapis.com/ingress` are equal.
	Metadata *ObjectMeta `pulumi:"metadata"`
	Project  *string     `pulumi:"project"`
	// Spec holds the desired state of the Service (from the client).
	Spec *ServiceSpec `pulumi:"spec"`
	// Status communicates the observed state of the Service (from the controller).
	Status *ServiceStatus `pulumi:"status"`
}

// The set of arguments for constructing a Service resource.
type ServiceArgs struct {
	// The API version for this call such as "serving.knative.dev/v1".
	ApiVersion pulumi.StringPtrInput
	// Indicates that the server should validate the request and populate default values without persisting the request. Supported values: `all`
	DryRun pulumi.StringPtrInput
	// The kind of resource, in this case "Service".
	Kind     pulumi.StringPtrInput
	Location pulumi.StringPtrInput
	// Metadata associated with this Service, including name, namespace, labels, and annotations. Cloud Run (fully managed) uses the following annotation keys to configure features on a Service: * `run.googleapis.com/ingress` sets the ingress settings for the Service. See [the ingress settings documentation](/run/docs/securing/ingress) for details on configuring ingress settings. * `run.googleapis.com/ingress-status` is output-only and contains the currently active ingress settings for the Service. `run.googleapis.com/ingress-status` may differ from `run.googleapis.com/ingress` while the system is processing a change to `run.googleapis.com/ingress` or if the system failed to process a change to `run.googleapis.com/ingress`. When the system has processed all changes successfully `run.googleapis.com/ingress-status` and `run.googleapis.com/ingress` are equal.
	Metadata ObjectMetaPtrInput
	Project  pulumi.StringPtrInput
	// Spec holds the desired state of the Service (from the client).
	Spec ServiceSpecPtrInput
	// Status communicates the observed state of the Service (from the controller).
	Status ServiceStatusPtrInput
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
