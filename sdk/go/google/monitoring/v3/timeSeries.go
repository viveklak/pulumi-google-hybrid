// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v3

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Creates or adds data to one or more time series. The response is empty if all time series in the request were written. If any time series could not be written, a corresponding failure message is included in the error response.
type TimeSeries struct {
	pulumi.CustomResourceState
}

// NewTimeSeries registers a new resource with the given unique name, arguments, and options.
func NewTimeSeries(ctx *pulumi.Context,
	name string, args *TimeSeriesArgs, opts ...pulumi.ResourceOption) (*TimeSeries, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ProjectsId == nil {
		return nil, errors.New("invalid value for required argument 'ProjectsId'")
	}
	var resource TimeSeries
	err := ctx.RegisterResource("google-cloud:monitoring/v3:TimeSeries", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTimeSeries gets an existing TimeSeries resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTimeSeries(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TimeSeriesState, opts ...pulumi.ResourceOption) (*TimeSeries, error) {
	var resource TimeSeries
	err := ctx.ReadResource("google-cloud:monitoring/v3:TimeSeries", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TimeSeries resources.
type timeSeriesState struct {
}

type TimeSeriesState struct {
}

func (TimeSeriesState) ElementType() reflect.Type {
	return reflect.TypeOf((*timeSeriesState)(nil)).Elem()
}

type timeSeriesArgs struct {
	ProjectsId string `pulumi:"projectsId"`
	// Required. The new data to be added to a list of time series. Adds at most one data point to each of several time series. The new data point must be more recent than any other point in its time series. Each TimeSeries value must fully specify a unique time series by supplying all label values for the metric and the monitored resource.The maximum number of TimeSeries objects per Create request is 200.
	TimeSeries []TimeSeriesType `pulumi:"timeSeries"`
}

// The set of arguments for constructing a TimeSeries resource.
type TimeSeriesArgs struct {
	ProjectsId pulumi.StringInput
	// Required. The new data to be added to a list of time series. Adds at most one data point to each of several time series. The new data point must be more recent than any other point in its time series. Each TimeSeries value must fully specify a unique time series by supplying all label values for the metric and the monitored resource.The maximum number of TimeSeries objects per Create request is 200.
	TimeSeries TimeSeriesTypeArrayInput
}

func (TimeSeriesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*timeSeriesArgs)(nil)).Elem()
}

type TimeSeriesInput interface {
	pulumi.Input

	ToTimeSeriesOutput() TimeSeriesOutput
	ToTimeSeriesOutputWithContext(ctx context.Context) TimeSeriesOutput
}

func (*TimeSeries) ElementType() reflect.Type {
	return reflect.TypeOf((*TimeSeries)(nil))
}

func (i *TimeSeries) ToTimeSeriesOutput() TimeSeriesOutput {
	return i.ToTimeSeriesOutputWithContext(context.Background())
}

func (i *TimeSeries) ToTimeSeriesOutputWithContext(ctx context.Context) TimeSeriesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TimeSeriesOutput)
}

type TimeSeriesOutput struct {
	*pulumi.OutputState
}

func (TimeSeriesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TimeSeries)(nil))
}

func (o TimeSeriesOutput) ToTimeSeriesOutput() TimeSeriesOutput {
	return o
}

func (o TimeSeriesOutput) ToTimeSeriesOutputWithContext(ctx context.Context) TimeSeriesOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(TimeSeriesOutput{})
}
