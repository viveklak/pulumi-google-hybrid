// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the latest state of a long-running DlpJob. See https://cloud.google.com/dlp/docs/inspecting-storage and https://cloud.google.com/dlp/docs/compute-risk-analysis to learn more.
func LookupDlpJob(ctx *pulumi.Context, args *LookupDlpJobArgs, opts ...pulumi.InvokeOption) (*LookupDlpJobResult, error) {
	var rv LookupDlpJobResult
	err := ctx.Invoke("google-native:dlp/v2:getDlpJob", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDlpJobArgs struct {
	DlpJobId string  `pulumi:"dlpJobId"`
	Location string  `pulumi:"location"`
	Project  *string `pulumi:"project"`
}

type LookupDlpJobResult struct {
	// Time when the job was created.
	CreateTime string `pulumi:"createTime"`
	// Time when the job finished.
	EndTime string `pulumi:"endTime"`
	// A stream of errors encountered running the job.
	Errors []GooglePrivacyDlpV2ErrorResponse `pulumi:"errors"`
	// Results from inspecting a data source.
	InspectDetails GooglePrivacyDlpV2InspectDataSourceDetailsResponse `pulumi:"inspectDetails"`
	// If created by a job trigger, the resource name of the trigger that instantiated the job.
	JobTriggerName string `pulumi:"jobTriggerName"`
	// The server-assigned name.
	Name string `pulumi:"name"`
	// Results from analyzing risk of a data source.
	RiskDetails GooglePrivacyDlpV2AnalyzeDataSourceRiskDetailsResponse `pulumi:"riskDetails"`
	// Time when the job started.
	StartTime string `pulumi:"startTime"`
	// State of a job.
	State string `pulumi:"state"`
	// The type of job.
	Type string `pulumi:"type"`
}

func LookupDlpJobOutput(ctx *pulumi.Context, args LookupDlpJobOutputArgs, opts ...pulumi.InvokeOption) LookupDlpJobResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDlpJobResult, error) {
			args := v.(LookupDlpJobArgs)
			r, err := LookupDlpJob(ctx, &args, opts...)
			return *r, err
		}).(LookupDlpJobResultOutput)
}

type LookupDlpJobOutputArgs struct {
	DlpJobId pulumi.StringInput    `pulumi:"dlpJobId"`
	Location pulumi.StringInput    `pulumi:"location"`
	Project  pulumi.StringPtrInput `pulumi:"project"`
}

func (LookupDlpJobOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDlpJobArgs)(nil)).Elem()
}

type LookupDlpJobResultOutput struct{ *pulumi.OutputState }

func (LookupDlpJobResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDlpJobResult)(nil)).Elem()
}

func (o LookupDlpJobResultOutput) ToLookupDlpJobResultOutput() LookupDlpJobResultOutput {
	return o
}

func (o LookupDlpJobResultOutput) ToLookupDlpJobResultOutputWithContext(ctx context.Context) LookupDlpJobResultOutput {
	return o
}

// Time when the job was created.
func (o LookupDlpJobResultOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDlpJobResult) string { return v.CreateTime }).(pulumi.StringOutput)
}

// Time when the job finished.
func (o LookupDlpJobResultOutput) EndTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDlpJobResult) string { return v.EndTime }).(pulumi.StringOutput)
}

// A stream of errors encountered running the job.
func (o LookupDlpJobResultOutput) Errors() GooglePrivacyDlpV2ErrorResponseArrayOutput {
	return o.ApplyT(func(v LookupDlpJobResult) []GooglePrivacyDlpV2ErrorResponse { return v.Errors }).(GooglePrivacyDlpV2ErrorResponseArrayOutput)
}

// Results from inspecting a data source.
func (o LookupDlpJobResultOutput) InspectDetails() GooglePrivacyDlpV2InspectDataSourceDetailsResponseOutput {
	return o.ApplyT(func(v LookupDlpJobResult) GooglePrivacyDlpV2InspectDataSourceDetailsResponse { return v.InspectDetails }).(GooglePrivacyDlpV2InspectDataSourceDetailsResponseOutput)
}

// If created by a job trigger, the resource name of the trigger that instantiated the job.
func (o LookupDlpJobResultOutput) JobTriggerName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDlpJobResult) string { return v.JobTriggerName }).(pulumi.StringOutput)
}

// The server-assigned name.
func (o LookupDlpJobResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDlpJobResult) string { return v.Name }).(pulumi.StringOutput)
}

// Results from analyzing risk of a data source.
func (o LookupDlpJobResultOutput) RiskDetails() GooglePrivacyDlpV2AnalyzeDataSourceRiskDetailsResponseOutput {
	return o.ApplyT(func(v LookupDlpJobResult) GooglePrivacyDlpV2AnalyzeDataSourceRiskDetailsResponse {
		return v.RiskDetails
	}).(GooglePrivacyDlpV2AnalyzeDataSourceRiskDetailsResponseOutput)
}

// Time when the job started.
func (o LookupDlpJobResultOutput) StartTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDlpJobResult) string { return v.StartTime }).(pulumi.StringOutput)
}

// State of a job.
func (o LookupDlpJobResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDlpJobResult) string { return v.State }).(pulumi.StringOutput)
}

// The type of job.
func (o LookupDlpJobResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDlpJobResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDlpJobResultOutput{})
}
