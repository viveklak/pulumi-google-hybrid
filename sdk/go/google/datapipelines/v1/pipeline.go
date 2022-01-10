// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a pipeline. For a batch pipeline, you can pass scheduler information. Data Pipelines uses the scheduler information to create an internal scheduler that runs jobs periodically. If the internal scheduler is not configured, you can use RunPipeline to run jobs.
type Pipeline struct {
	pulumi.CustomResourceState

	// Immutable. The timestamp when the pipeline was initially created. Set by the Data Pipelines service.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// The display name of the pipeline. It can contain only letters ([A-Za-z]), numbers ([0-9]), hyphens (-), and underscores (_).
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// Number of jobs.
	JobCount pulumi.IntOutput `pulumi:"jobCount"`
	// Immutable. The timestamp when the pipeline was last modified. Set by the Data Pipelines service.
	LastUpdateTime pulumi.StringOutput `pulumi:"lastUpdateTime"`
	// The pipeline name. For example: `projects/PROJECT_ID/locations/LOCATION_ID/pipelines/PIPELINE_ID`. * `PROJECT_ID` can contain letters ([A-Za-z]), numbers ([0-9]), hyphens (-), colons (:), and periods (.). For more information, see [Identifying projects](https://cloud.google.com/resource-manager/docs/creating-managing-projects#identifying_projects). * `LOCATION_ID` is the canonical ID for the pipeline's location. The list of available locations can be obtained by calling `google.cloud.location.Locations.ListLocations`. Note that the Data Pipelines service is not available in all regions. It depends on Cloud Scheduler, an App Engine application, so it's only available in [App Engine regions](https://cloud.google.com/about/locations#region). * `PIPELINE_ID` is the ID of the pipeline. Must be unique for the selected project and location.
	Name pulumi.StringOutput `pulumi:"name"`
	// Immutable. The sources of the pipeline (for example, Dataplex). The keys and values are set by the corresponding sources during pipeline creation.
	PipelineSources pulumi.StringMapOutput `pulumi:"pipelineSources"`
	// Internal scheduling information for a pipeline. If this information is provided, periodic jobs will be created per the schedule. If not, users are responsible for creating jobs externally.
	ScheduleInfo GoogleCloudDatapipelinesV1ScheduleSpecResponseOutput `pulumi:"scheduleInfo"`
	// Optional. A service account email to be used with the Cloud Scheduler job. If not specified, the default compute engine service account will be used.
	SchedulerServiceAccountEmail pulumi.StringOutput `pulumi:"schedulerServiceAccountEmail"`
	// The state of the pipeline. When the pipeline is created, the state is set to 'PIPELINE_STATE_ACTIVE' by default. State changes can be requested by setting the state to stopping, paused, or resuming. State cannot be changed through UpdatePipeline requests.
	State pulumi.StringOutput `pulumi:"state"`
	// The type of the pipeline. This field affects the scheduling of the pipeline and the type of metrics to show for the pipeline.
	Type pulumi.StringOutput `pulumi:"type"`
	// Workload information for creating new jobs.
	Workload GoogleCloudDatapipelinesV1WorkloadResponseOutput `pulumi:"workload"`
}

// NewPipeline registers a new resource with the given unique name, arguments, and options.
func NewPipeline(ctx *pulumi.Context,
	name string, args *PipelineArgs, opts ...pulumi.ResourceOption) (*Pipeline, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	if args.State == nil {
		return nil, errors.New("invalid value for required argument 'State'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	var resource Pipeline
	err := ctx.RegisterResource("google-native:datapipelines/v1:Pipeline", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPipeline gets an existing Pipeline resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPipeline(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PipelineState, opts ...pulumi.ResourceOption) (*Pipeline, error) {
	var resource Pipeline
	err := ctx.ReadResource("google-native:datapipelines/v1:Pipeline", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Pipeline resources.
type pipelineState struct {
}

type PipelineState struct {
}

func (PipelineState) ElementType() reflect.Type {
	return reflect.TypeOf((*pipelineState)(nil)).Elem()
}

type pipelineArgs struct {
	// The display name of the pipeline. It can contain only letters ([A-Za-z]), numbers ([0-9]), hyphens (-), and underscores (_).
	DisplayName string  `pulumi:"displayName"`
	Location    *string `pulumi:"location"`
	// The pipeline name. For example: `projects/PROJECT_ID/locations/LOCATION_ID/pipelines/PIPELINE_ID`. * `PROJECT_ID` can contain letters ([A-Za-z]), numbers ([0-9]), hyphens (-), colons (:), and periods (.). For more information, see [Identifying projects](https://cloud.google.com/resource-manager/docs/creating-managing-projects#identifying_projects). * `LOCATION_ID` is the canonical ID for the pipeline's location. The list of available locations can be obtained by calling `google.cloud.location.Locations.ListLocations`. Note that the Data Pipelines service is not available in all regions. It depends on Cloud Scheduler, an App Engine application, so it's only available in [App Engine regions](https://cloud.google.com/about/locations#region). * `PIPELINE_ID` is the ID of the pipeline. Must be unique for the selected project and location.
	Name *string `pulumi:"name"`
	// Immutable. The sources of the pipeline (for example, Dataplex). The keys and values are set by the corresponding sources during pipeline creation.
	PipelineSources map[string]string `pulumi:"pipelineSources"`
	Project         *string           `pulumi:"project"`
	// Internal scheduling information for a pipeline. If this information is provided, periodic jobs will be created per the schedule. If not, users are responsible for creating jobs externally.
	ScheduleInfo *GoogleCloudDatapipelinesV1ScheduleSpec `pulumi:"scheduleInfo"`
	// Optional. A service account email to be used with the Cloud Scheduler job. If not specified, the default compute engine service account will be used.
	SchedulerServiceAccountEmail *string `pulumi:"schedulerServiceAccountEmail"`
	// The state of the pipeline. When the pipeline is created, the state is set to 'PIPELINE_STATE_ACTIVE' by default. State changes can be requested by setting the state to stopping, paused, or resuming. State cannot be changed through UpdatePipeline requests.
	State PipelineStateEnum `pulumi:"state"`
	// The type of the pipeline. This field affects the scheduling of the pipeline and the type of metrics to show for the pipeline.
	Type PipelineType `pulumi:"type"`
	// Workload information for creating new jobs.
	Workload *GoogleCloudDatapipelinesV1Workload `pulumi:"workload"`
}

// The set of arguments for constructing a Pipeline resource.
type PipelineArgs struct {
	// The display name of the pipeline. It can contain only letters ([A-Za-z]), numbers ([0-9]), hyphens (-), and underscores (_).
	DisplayName pulumi.StringInput
	Location    pulumi.StringPtrInput
	// The pipeline name. For example: `projects/PROJECT_ID/locations/LOCATION_ID/pipelines/PIPELINE_ID`. * `PROJECT_ID` can contain letters ([A-Za-z]), numbers ([0-9]), hyphens (-), colons (:), and periods (.). For more information, see [Identifying projects](https://cloud.google.com/resource-manager/docs/creating-managing-projects#identifying_projects). * `LOCATION_ID` is the canonical ID for the pipeline's location. The list of available locations can be obtained by calling `google.cloud.location.Locations.ListLocations`. Note that the Data Pipelines service is not available in all regions. It depends on Cloud Scheduler, an App Engine application, so it's only available in [App Engine regions](https://cloud.google.com/about/locations#region). * `PIPELINE_ID` is the ID of the pipeline. Must be unique for the selected project and location.
	Name pulumi.StringPtrInput
	// Immutable. The sources of the pipeline (for example, Dataplex). The keys and values are set by the corresponding sources during pipeline creation.
	PipelineSources pulumi.StringMapInput
	Project         pulumi.StringPtrInput
	// Internal scheduling information for a pipeline. If this information is provided, periodic jobs will be created per the schedule. If not, users are responsible for creating jobs externally.
	ScheduleInfo GoogleCloudDatapipelinesV1ScheduleSpecPtrInput
	// Optional. A service account email to be used with the Cloud Scheduler job. If not specified, the default compute engine service account will be used.
	SchedulerServiceAccountEmail pulumi.StringPtrInput
	// The state of the pipeline. When the pipeline is created, the state is set to 'PIPELINE_STATE_ACTIVE' by default. State changes can be requested by setting the state to stopping, paused, or resuming. State cannot be changed through UpdatePipeline requests.
	State PipelineStateEnumInput
	// The type of the pipeline. This field affects the scheduling of the pipeline and the type of metrics to show for the pipeline.
	Type PipelineTypeInput
	// Workload information for creating new jobs.
	Workload GoogleCloudDatapipelinesV1WorkloadPtrInput
}

func (PipelineArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*pipelineArgs)(nil)).Elem()
}

type PipelineInput interface {
	pulumi.Input

	ToPipelineOutput() PipelineOutput
	ToPipelineOutputWithContext(ctx context.Context) PipelineOutput
}

func (*Pipeline) ElementType() reflect.Type {
	return reflect.TypeOf((**Pipeline)(nil)).Elem()
}

func (i *Pipeline) ToPipelineOutput() PipelineOutput {
	return i.ToPipelineOutputWithContext(context.Background())
}

func (i *Pipeline) ToPipelineOutputWithContext(ctx context.Context) PipelineOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PipelineOutput)
}

type PipelineOutput struct{ *pulumi.OutputState }

func (PipelineOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Pipeline)(nil)).Elem()
}

func (o PipelineOutput) ToPipelineOutput() PipelineOutput {
	return o
}

func (o PipelineOutput) ToPipelineOutputWithContext(ctx context.Context) PipelineOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PipelineInput)(nil)).Elem(), &Pipeline{})
	pulumi.RegisterOutputType(PipelineOutput{})
}
