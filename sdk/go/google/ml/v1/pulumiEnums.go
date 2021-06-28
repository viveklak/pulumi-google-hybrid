// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Required. The optimization goal of the metric.
type GoogleCloudMlV1_StudyConfig_MetricSpecGoal pulumi.String

const (
	// Goal Type will default to maximize.
	GoogleCloudMlV1_StudyConfig_MetricSpecGoalGoalTypeUnspecified = GoogleCloudMlV1_StudyConfig_MetricSpecGoal("GOAL_TYPE_UNSPECIFIED")
	// Maximize the goal metric.
	GoogleCloudMlV1_StudyConfig_MetricSpecGoalMaximize = GoogleCloudMlV1_StudyConfig_MetricSpecGoal("MAXIMIZE")
	// Minimize the goal metric.
	GoogleCloudMlV1_StudyConfig_MetricSpecGoalMinimize = GoogleCloudMlV1_StudyConfig_MetricSpecGoal("MINIMIZE")
)

func (GoogleCloudMlV1_StudyConfig_MetricSpecGoal) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e GoogleCloudMlV1_StudyConfig_MetricSpecGoal) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e GoogleCloudMlV1_StudyConfig_MetricSpecGoal) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e GoogleCloudMlV1_StudyConfig_MetricSpecGoal) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e GoogleCloudMlV1_StudyConfig_MetricSpecGoal) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

// How the parameter should be scaled. Leave unset for categorical parameters.
type GoogleCloudMlV1_StudyConfig_ParameterSpecScaleType pulumi.String

const (
	// By default, no scaling is applied.
	GoogleCloudMlV1_StudyConfig_ParameterSpecScaleTypeScaleTypeUnspecified = GoogleCloudMlV1_StudyConfig_ParameterSpecScaleType("SCALE_TYPE_UNSPECIFIED")
	// Scales the feasible space to (0, 1) linearly.
	GoogleCloudMlV1_StudyConfig_ParameterSpecScaleTypeUnitLinearScale = GoogleCloudMlV1_StudyConfig_ParameterSpecScaleType("UNIT_LINEAR_SCALE")
	// Scales the feasible space logarithmically to (0, 1). The entire feasible space must be strictly positive.
	GoogleCloudMlV1_StudyConfig_ParameterSpecScaleTypeUnitLogScale = GoogleCloudMlV1_StudyConfig_ParameterSpecScaleType("UNIT_LOG_SCALE")
	// Scales the feasible space "reverse" logarithmically to (0, 1). The result is that values close to the top of the feasible space are spread out more than points near the bottom. The entire feasible space must be strictly positive.
	GoogleCloudMlV1_StudyConfig_ParameterSpecScaleTypeUnitReverseLogScale = GoogleCloudMlV1_StudyConfig_ParameterSpecScaleType("UNIT_REVERSE_LOG_SCALE")
)

func (GoogleCloudMlV1_StudyConfig_ParameterSpecScaleType) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e GoogleCloudMlV1_StudyConfig_ParameterSpecScaleType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e GoogleCloudMlV1_StudyConfig_ParameterSpecScaleType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e GoogleCloudMlV1_StudyConfig_ParameterSpecScaleType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e GoogleCloudMlV1_StudyConfig_ParameterSpecScaleType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

// Required. The type of the parameter.
type GoogleCloudMlV1_StudyConfig_ParameterSpecType pulumi.String

const (
	// You must specify a valid type. Using this unspecified type will result in an error.
	GoogleCloudMlV1_StudyConfig_ParameterSpecTypeParameterTypeUnspecified = GoogleCloudMlV1_StudyConfig_ParameterSpecType("PARAMETER_TYPE_UNSPECIFIED")
	// Type for real-valued parameters.
	GoogleCloudMlV1_StudyConfig_ParameterSpecTypeDouble = GoogleCloudMlV1_StudyConfig_ParameterSpecType("DOUBLE")
	// Type for integral parameters.
	GoogleCloudMlV1_StudyConfig_ParameterSpecTypeInteger = GoogleCloudMlV1_StudyConfig_ParameterSpecType("INTEGER")
	// The parameter is categorical, with a value chosen from the categories field.
	GoogleCloudMlV1_StudyConfig_ParameterSpecTypeCategorical = GoogleCloudMlV1_StudyConfig_ParameterSpecType("CATEGORICAL")
	// The parameter is real valued, with a fixed set of feasible points. If `type==DISCRETE`, feasible_points must be provided, and {`min_value`, `max_value`} will be ignored.
	GoogleCloudMlV1_StudyConfig_ParameterSpecTypeDiscrete = GoogleCloudMlV1_StudyConfig_ParameterSpecType("DISCRETE")
)

func (GoogleCloudMlV1_StudyConfig_ParameterSpecType) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e GoogleCloudMlV1_StudyConfig_ParameterSpecType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e GoogleCloudMlV1_StudyConfig_ParameterSpecType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e GoogleCloudMlV1_StudyConfig_ParameterSpecType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e GoogleCloudMlV1_StudyConfig_ParameterSpecType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

// The type of accelerator to use.
type GoogleCloudMlV1__AcceleratorConfigType pulumi.String

const (
	// Unspecified accelerator type. Default to no GPU.
	GoogleCloudMlV1__AcceleratorConfigTypeAcceleratorTypeUnspecified = GoogleCloudMlV1__AcceleratorConfigType("ACCELERATOR_TYPE_UNSPECIFIED")
	// Nvidia Tesla K80 GPU.
	GoogleCloudMlV1__AcceleratorConfigTypeNvidiaTeslaK80 = GoogleCloudMlV1__AcceleratorConfigType("NVIDIA_TESLA_K80")
	// Nvidia Tesla P100 GPU.
	GoogleCloudMlV1__AcceleratorConfigTypeNvidiaTeslaP100 = GoogleCloudMlV1__AcceleratorConfigType("NVIDIA_TESLA_P100")
	// Nvidia V100 GPU.
	GoogleCloudMlV1__AcceleratorConfigTypeNvidiaTeslaV100 = GoogleCloudMlV1__AcceleratorConfigType("NVIDIA_TESLA_V100")
	// Nvidia Tesla P4 GPU.
	GoogleCloudMlV1__AcceleratorConfigTypeNvidiaTeslaP4 = GoogleCloudMlV1__AcceleratorConfigType("NVIDIA_TESLA_P4")
	// Nvidia T4 GPU.
	GoogleCloudMlV1__AcceleratorConfigTypeNvidiaTeslaT4 = GoogleCloudMlV1__AcceleratorConfigType("NVIDIA_TESLA_T4")
	// Nvidia A100 GPU.
	GoogleCloudMlV1__AcceleratorConfigTypeNvidiaTeslaA100 = GoogleCloudMlV1__AcceleratorConfigType("NVIDIA_TESLA_A100")
	// TPU v2.
	GoogleCloudMlV1__AcceleratorConfigTypeTpuV2 = GoogleCloudMlV1__AcceleratorConfigType("TPU_V2")
	// TPU v3.
	GoogleCloudMlV1__AcceleratorConfigTypeTpuV3 = GoogleCloudMlV1__AcceleratorConfigType("TPU_V3")
)

func (GoogleCloudMlV1__AcceleratorConfigType) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e GoogleCloudMlV1__AcceleratorConfigType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e GoogleCloudMlV1__AcceleratorConfigType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e GoogleCloudMlV1__AcceleratorConfigType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e GoogleCloudMlV1__AcceleratorConfigType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

// Optional. The search algorithm specified for the hyperparameter tuning job. Uses the default AI Platform hyperparameter tuning algorithm if unspecified.
type GoogleCloudMlV1__HyperparameterSpecAlgorithm pulumi.String

const (
	// The default algorithm used by the hyperparameter tuning service. This is a Bayesian optimization algorithm.
	GoogleCloudMlV1__HyperparameterSpecAlgorithmAlgorithmUnspecified = GoogleCloudMlV1__HyperparameterSpecAlgorithm("ALGORITHM_UNSPECIFIED")
	// Simple grid search within the feasible space. To use grid search, all parameters must be `INTEGER`, `CATEGORICAL`, or `DISCRETE`.
	GoogleCloudMlV1__HyperparameterSpecAlgorithmGridSearch = GoogleCloudMlV1__HyperparameterSpecAlgorithm("GRID_SEARCH")
	// Simple random search within the feasible space.
	GoogleCloudMlV1__HyperparameterSpecAlgorithmRandomSearch = GoogleCloudMlV1__HyperparameterSpecAlgorithm("RANDOM_SEARCH")
)

func (GoogleCloudMlV1__HyperparameterSpecAlgorithm) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e GoogleCloudMlV1__HyperparameterSpecAlgorithm) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e GoogleCloudMlV1__HyperparameterSpecAlgorithm) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e GoogleCloudMlV1__HyperparameterSpecAlgorithm) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e GoogleCloudMlV1__HyperparameterSpecAlgorithm) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

// Required. The type of goal to use for tuning. Available types are `MAXIMIZE` and `MINIMIZE`. Defaults to `MAXIMIZE`.
type GoogleCloudMlV1__HyperparameterSpecGoal pulumi.String

const (
	// Goal Type will default to maximize.
	GoogleCloudMlV1__HyperparameterSpecGoalGoalTypeUnspecified = GoogleCloudMlV1__HyperparameterSpecGoal("GOAL_TYPE_UNSPECIFIED")
	// Maximize the goal metric.
	GoogleCloudMlV1__HyperparameterSpecGoalMaximize = GoogleCloudMlV1__HyperparameterSpecGoal("MAXIMIZE")
	// Minimize the goal metric.
	GoogleCloudMlV1__HyperparameterSpecGoalMinimize = GoogleCloudMlV1__HyperparameterSpecGoal("MINIMIZE")
)

func (GoogleCloudMlV1__HyperparameterSpecGoal) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e GoogleCloudMlV1__HyperparameterSpecGoal) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e GoogleCloudMlV1__HyperparameterSpecGoal) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e GoogleCloudMlV1__HyperparameterSpecGoal) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e GoogleCloudMlV1__HyperparameterSpecGoal) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

// metric name.
type GoogleCloudMlV1__MetricSpecName pulumi.String

const (
	// Unspecified MetricName.
	GoogleCloudMlV1__MetricSpecNameMetricNameUnspecified = GoogleCloudMlV1__MetricSpecName("METRIC_NAME_UNSPECIFIED")
	// CPU usage.
	GoogleCloudMlV1__MetricSpecNameCpuUsage = GoogleCloudMlV1__MetricSpecName("CPU_USAGE")
	// GPU duty cycle.
	GoogleCloudMlV1__MetricSpecNameGpuDutyCycle = GoogleCloudMlV1__MetricSpecName("GPU_DUTY_CYCLE")
)

func (GoogleCloudMlV1__MetricSpecName) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e GoogleCloudMlV1__MetricSpecName) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e GoogleCloudMlV1__MetricSpecName) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e GoogleCloudMlV1__MetricSpecName) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e GoogleCloudMlV1__MetricSpecName) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

// Optional. How the parameter should be scaled to the hypercube. Leave unset for categorical parameters. Some kind of scaling is strongly recommended for real or integral parameters (e.g., `UNIT_LINEAR_SCALE`).
type GoogleCloudMlV1__ParameterSpecScaleType pulumi.String

const (
	// By default, no scaling is applied.
	GoogleCloudMlV1__ParameterSpecScaleTypeNone = GoogleCloudMlV1__ParameterSpecScaleType("NONE")
	// Scales the feasible space to (0, 1) linearly.
	GoogleCloudMlV1__ParameterSpecScaleTypeUnitLinearScale = GoogleCloudMlV1__ParameterSpecScaleType("UNIT_LINEAR_SCALE")
	// Scales the feasible space logarithmically to (0, 1). The entire feasible space must be strictly positive.
	GoogleCloudMlV1__ParameterSpecScaleTypeUnitLogScale = GoogleCloudMlV1__ParameterSpecScaleType("UNIT_LOG_SCALE")
	// Scales the feasible space "reverse" logarithmically to (0, 1). The result is that values close to the top of the feasible space are spread out more than points near the bottom. The entire feasible space must be strictly positive.
	GoogleCloudMlV1__ParameterSpecScaleTypeUnitReverseLogScale = GoogleCloudMlV1__ParameterSpecScaleType("UNIT_REVERSE_LOG_SCALE")
)

func (GoogleCloudMlV1__ParameterSpecScaleType) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e GoogleCloudMlV1__ParameterSpecScaleType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e GoogleCloudMlV1__ParameterSpecScaleType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e GoogleCloudMlV1__ParameterSpecScaleType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e GoogleCloudMlV1__ParameterSpecScaleType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

// Required. The type of the parameter.
type GoogleCloudMlV1__ParameterSpecType pulumi.String

const (
	// You must specify a valid type. Using this unspecified type will result in an error.
	GoogleCloudMlV1__ParameterSpecTypeParameterTypeUnspecified = GoogleCloudMlV1__ParameterSpecType("PARAMETER_TYPE_UNSPECIFIED")
	// Type for real-valued parameters.
	GoogleCloudMlV1__ParameterSpecTypeDouble = GoogleCloudMlV1__ParameterSpecType("DOUBLE")
	// Type for integral parameters.
	GoogleCloudMlV1__ParameterSpecTypeInteger = GoogleCloudMlV1__ParameterSpecType("INTEGER")
	// The parameter is categorical, with a value chosen from the categories field.
	GoogleCloudMlV1__ParameterSpecTypeCategorical = GoogleCloudMlV1__ParameterSpecType("CATEGORICAL")
	// The parameter is real valued, with a fixed set of feasible points. If `type==DISCRETE`, feasible_points must be provided, and {`min_value`, `max_value`} will be ignored.
	GoogleCloudMlV1__ParameterSpecTypeDiscrete = GoogleCloudMlV1__ParameterSpecType("DISCRETE")
)

func (GoogleCloudMlV1__ParameterSpecType) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e GoogleCloudMlV1__ParameterSpecType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e GoogleCloudMlV1__ParameterSpecType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e GoogleCloudMlV1__ParameterSpecType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e GoogleCloudMlV1__ParameterSpecType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

// Required. The format of the input data files.
type GoogleCloudMlV1__PredictionInputDataFormat pulumi.String

const (
	// Unspecified format.
	GoogleCloudMlV1__PredictionInputDataFormatDataFormatUnspecified = GoogleCloudMlV1__PredictionInputDataFormat("DATA_FORMAT_UNSPECIFIED")
	// Each line of the file is a JSON dictionary representing one record.
	GoogleCloudMlV1__PredictionInputDataFormatJson = GoogleCloudMlV1__PredictionInputDataFormat("JSON")
	// Deprecated. Use JSON instead.
	GoogleCloudMlV1__PredictionInputDataFormatText = GoogleCloudMlV1__PredictionInputDataFormat("TEXT")
	// The source file is a TFRecord file. Currently available only for input data.
	GoogleCloudMlV1__PredictionInputDataFormatTfRecord = GoogleCloudMlV1__PredictionInputDataFormat("TF_RECORD")
	// The source file is a GZIP-compressed TFRecord file. Currently available only for input data.
	GoogleCloudMlV1__PredictionInputDataFormatTfRecordGzip = GoogleCloudMlV1__PredictionInputDataFormat("TF_RECORD_GZIP")
	// Values are comma-separated rows, with keys in a separate file. Currently available only for output data.
	GoogleCloudMlV1__PredictionInputDataFormatCsv = GoogleCloudMlV1__PredictionInputDataFormat("CSV")
)

func (GoogleCloudMlV1__PredictionInputDataFormat) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e GoogleCloudMlV1__PredictionInputDataFormat) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e GoogleCloudMlV1__PredictionInputDataFormat) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e GoogleCloudMlV1__PredictionInputDataFormat) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e GoogleCloudMlV1__PredictionInputDataFormat) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

// Optional. Format of the output data files, defaults to JSON.
type GoogleCloudMlV1__PredictionInputOutputDataFormat pulumi.String

const (
	// Unspecified format.
	GoogleCloudMlV1__PredictionInputOutputDataFormatDataFormatUnspecified = GoogleCloudMlV1__PredictionInputOutputDataFormat("DATA_FORMAT_UNSPECIFIED")
	// Each line of the file is a JSON dictionary representing one record.
	GoogleCloudMlV1__PredictionInputOutputDataFormatJson = GoogleCloudMlV1__PredictionInputOutputDataFormat("JSON")
	// Deprecated. Use JSON instead.
	GoogleCloudMlV1__PredictionInputOutputDataFormatText = GoogleCloudMlV1__PredictionInputOutputDataFormat("TEXT")
	// The source file is a TFRecord file. Currently available only for input data.
	GoogleCloudMlV1__PredictionInputOutputDataFormatTfRecord = GoogleCloudMlV1__PredictionInputOutputDataFormat("TF_RECORD")
	// The source file is a GZIP-compressed TFRecord file. Currently available only for input data.
	GoogleCloudMlV1__PredictionInputOutputDataFormatTfRecordGzip = GoogleCloudMlV1__PredictionInputOutputDataFormat("TF_RECORD_GZIP")
	// Values are comma-separated rows, with keys in a separate file. Currently available only for output data.
	GoogleCloudMlV1__PredictionInputOutputDataFormatCsv = GoogleCloudMlV1__PredictionInputOutputDataFormat("CSV")
)

func (GoogleCloudMlV1__PredictionInputOutputDataFormat) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e GoogleCloudMlV1__PredictionInputOutputDataFormat) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e GoogleCloudMlV1__PredictionInputOutputDataFormat) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e GoogleCloudMlV1__PredictionInputOutputDataFormat) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e GoogleCloudMlV1__PredictionInputOutputDataFormat) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

// The search algorithm specified for the study.
type GoogleCloudMlV1__StudyConfigAlgorithm pulumi.String

const (
	// The default algorithm used by the Cloud AI Platform Vizier service.
	GoogleCloudMlV1__StudyConfigAlgorithmAlgorithmUnspecified = GoogleCloudMlV1__StudyConfigAlgorithm("ALGORITHM_UNSPECIFIED")
	// Gaussian Process Bandit.
	GoogleCloudMlV1__StudyConfigAlgorithmGaussianProcessBandit = GoogleCloudMlV1__StudyConfigAlgorithm("GAUSSIAN_PROCESS_BANDIT")
	// Simple grid search within the feasible space. To use grid search, all parameters must be `INTEGER`, `CATEGORICAL`, or `DISCRETE`.
	GoogleCloudMlV1__StudyConfigAlgorithmGridSearch = GoogleCloudMlV1__StudyConfigAlgorithm("GRID_SEARCH")
	// Simple random search within the feasible space.
	GoogleCloudMlV1__StudyConfigAlgorithmRandomSearch = GoogleCloudMlV1__StudyConfigAlgorithm("RANDOM_SEARCH")
)

func (GoogleCloudMlV1__StudyConfigAlgorithm) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e GoogleCloudMlV1__StudyConfigAlgorithm) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e GoogleCloudMlV1__StudyConfigAlgorithm) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e GoogleCloudMlV1__StudyConfigAlgorithm) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e GoogleCloudMlV1__StudyConfigAlgorithm) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

// Required. Specifies the machine types, the number of replicas for workers and parameter servers.
type GoogleCloudMlV1__TrainingInputScaleTier pulumi.String

const (
	// A single worker instance. This tier is suitable for learning how to use Cloud ML, and for experimenting with new models using small datasets.
	GoogleCloudMlV1__TrainingInputScaleTierBasic = GoogleCloudMlV1__TrainingInputScaleTier("BASIC")
	// Many workers and a few parameter servers.
	GoogleCloudMlV1__TrainingInputScaleTierStandard1 = GoogleCloudMlV1__TrainingInputScaleTier("STANDARD_1")
	// A large number of workers with many parameter servers.
	GoogleCloudMlV1__TrainingInputScaleTierPremium1 = GoogleCloudMlV1__TrainingInputScaleTier("PREMIUM_1")
	// A single worker instance [with a GPU](/ai-platform/training/docs/using-gpus).
	GoogleCloudMlV1__TrainingInputScaleTierBasicGpu = GoogleCloudMlV1__TrainingInputScaleTier("BASIC_GPU")
	// A single worker instance with a [Cloud TPU](/ml-engine/docs/tensorflow/using-tpus).
	GoogleCloudMlV1__TrainingInputScaleTierBasicTpu = GoogleCloudMlV1__TrainingInputScaleTier("BASIC_TPU")
	// The CUSTOM tier is not a set tier, but rather enables you to use your own cluster specification. When you use this tier, set values to configure your processing cluster according to these guidelines: * You _must_ set `TrainingInput.masterType` to specify the type of machine to use for your master node. This is the only required setting. * You _may_ set `TrainingInput.workerCount` to specify the number of workers to use. If you specify one or more workers, you _must_ also set `TrainingInput.workerType` to specify the type of machine to use for your worker nodes. * You _may_ set `TrainingInput.parameterServerCount` to specify the number of parameter servers to use. If you specify one or more parameter servers, you _must_ also set `TrainingInput.parameterServerType` to specify the type of machine to use for your parameter servers. Note that all of your workers must use the same machine type, which can be different from your parameter server type and master type. Your parameter servers must likewise use the same machine type, which can be different from your worker type and master type.
	GoogleCloudMlV1__TrainingInputScaleTierCustom = GoogleCloudMlV1__TrainingInputScaleTier("CUSTOM")
)

func (GoogleCloudMlV1__TrainingInputScaleTier) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e GoogleCloudMlV1__TrainingInputScaleTier) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e GoogleCloudMlV1__TrainingInputScaleTier) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e GoogleCloudMlV1__TrainingInputScaleTier) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e GoogleCloudMlV1__TrainingInputScaleTier) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

// The log type that this config enables.
type GoogleIamV1__AuditLogConfigLogType pulumi.String

const (
	// Default case. Should never be this.
	GoogleIamV1__AuditLogConfigLogTypeLogTypeUnspecified = GoogleIamV1__AuditLogConfigLogType("LOG_TYPE_UNSPECIFIED")
	// Admin reads. Example: CloudIAM getIamPolicy
	GoogleIamV1__AuditLogConfigLogTypeAdminRead = GoogleIamV1__AuditLogConfigLogType("ADMIN_READ")
	// Data writes. Example: CloudSQL Users create
	GoogleIamV1__AuditLogConfigLogTypeDataWrite = GoogleIamV1__AuditLogConfigLogType("DATA_WRITE")
	// Data reads. Example: CloudSQL Users list
	GoogleIamV1__AuditLogConfigLogTypeDataRead = GoogleIamV1__AuditLogConfigLogType("DATA_READ")
)

func (GoogleIamV1__AuditLogConfigLogType) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e GoogleIamV1__AuditLogConfigLogType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e GoogleIamV1__AuditLogConfigLogType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e GoogleIamV1__AuditLogConfigLogType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e GoogleIamV1__AuditLogConfigLogType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

// The detailed state of a trial.
type TrialStateEnum pulumi.String

const (
	// The trial state is unspecified.
	TrialStateEnumStateUnspecified = TrialStateEnum("STATE_UNSPECIFIED")
	// Indicates that a specific trial has been requested, but it has not yet been suggested by the service.
	TrialStateEnumRequested = TrialStateEnum("REQUESTED")
	// Indicates that the trial has been suggested.
	TrialStateEnumActive = TrialStateEnum("ACTIVE")
	// Indicates that the trial is done, and either has a final_measurement set, or is marked as trial_infeasible.
	TrialStateEnumCompleted = TrialStateEnum("COMPLETED")
	// Indicates that the trial should stop according to the service.
	TrialStateEnumStopping = TrialStateEnum("STOPPING")
)

func (TrialStateEnum) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e TrialStateEnum) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e TrialStateEnum) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e TrialStateEnum) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e TrialStateEnum) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

// Optional. The machine learning framework AI Platform uses to train this version of the model. Valid values are `TENSORFLOW`, `SCIKIT_LEARN`, `XGBOOST`. If you do not specify a framework, AI Platform will analyze files in the deployment_uri to determine a framework. If you choose `SCIKIT_LEARN` or `XGBOOST`, you must also set the runtime version of the model to 1.4 or greater. Do **not** specify a framework if you're deploying a [custom prediction routine](/ai-platform/prediction/docs/custom-prediction-routines) or if you're using a [custom container](/ai-platform/prediction/docs/use-custom-container).
type VersionFramework pulumi.String

const (
	// Unspecified framework. Assigns a value based on the file suffix.
	VersionFrameworkFrameworkUnspecified = VersionFramework("FRAMEWORK_UNSPECIFIED")
	// Tensorflow framework.
	VersionFrameworkTensorflow = VersionFramework("TENSORFLOW")
	// Scikit-learn framework.
	VersionFrameworkScikitLearn = VersionFramework("SCIKIT_LEARN")
	// XGBoost framework.
	VersionFrameworkXgboost = VersionFramework("XGBOOST")
)

func (VersionFramework) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e VersionFramework) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e VersionFramework) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e VersionFramework) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e VersionFramework) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}
