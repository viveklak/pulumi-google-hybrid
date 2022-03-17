// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Fetches a specific dashboard.This method requires the monitoring.dashboards.get permission on the specified dashboard. For more information, see Cloud Identity and Access Management (https://cloud.google.com/iam).
func LookupDashboard(ctx *pulumi.Context, args *LookupDashboardArgs, opts ...pulumi.InvokeOption) (*LookupDashboardResult, error) {
	var rv LookupDashboardResult
	err := ctx.Invoke("google-native:monitoring/v1:getDashboard", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDashboardArgs struct {
	DashboardId string  `pulumi:"dashboardId"`
	Project     *string `pulumi:"project"`
}

type LookupDashboardResult struct {
	// The content is divided into equally spaced columns and the widgets are arranged vertically.
	ColumnLayout ColumnLayoutResponse `pulumi:"columnLayout"`
	// The mutable, human-readable name.
	DisplayName string `pulumi:"displayName"`
	// etag is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other. An etag is returned in the response to GetDashboard, and users are expected to put that etag in the request to UpdateDashboard to ensure that their change will be applied to the same version of the Dashboard configuration. The field should not be passed during dashboard creation.
	Etag string `pulumi:"etag"`
	// Content is arranged with a basic layout that re-flows a simple list of informational elements like widgets or tiles.
	GridLayout GridLayoutResponse `pulumi:"gridLayout"`
	// Labels applied to the dashboard
	Labels map[string]string `pulumi:"labels"`
	// The content is arranged as a grid of tiles, with each content widget occupying one or more grid blocks.
	MosaicLayout MosaicLayoutResponse `pulumi:"mosaicLayout"`
	// Immutable. The resource name of the dashboard.
	Name string `pulumi:"name"`
	// The content is divided into equally spaced rows and the widgets are arranged horizontally.
	RowLayout RowLayoutResponse `pulumi:"rowLayout"`
}

func LookupDashboardOutput(ctx *pulumi.Context, args LookupDashboardOutputArgs, opts ...pulumi.InvokeOption) LookupDashboardResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDashboardResult, error) {
			args := v.(LookupDashboardArgs)
			r, err := LookupDashboard(ctx, &args, opts...)
			return *r, err
		}).(LookupDashboardResultOutput)
}

type LookupDashboardOutputArgs struct {
	DashboardId pulumi.StringInput    `pulumi:"dashboardId"`
	Project     pulumi.StringPtrInput `pulumi:"project"`
}

func (LookupDashboardOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDashboardArgs)(nil)).Elem()
}

type LookupDashboardResultOutput struct{ *pulumi.OutputState }

func (LookupDashboardResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDashboardResult)(nil)).Elem()
}

func (o LookupDashboardResultOutput) ToLookupDashboardResultOutput() LookupDashboardResultOutput {
	return o
}

func (o LookupDashboardResultOutput) ToLookupDashboardResultOutputWithContext(ctx context.Context) LookupDashboardResultOutput {
	return o
}

// The content is divided into equally spaced columns and the widgets are arranged vertically.
func (o LookupDashboardResultOutput) ColumnLayout() ColumnLayoutResponseOutput {
	return o.ApplyT(func(v LookupDashboardResult) ColumnLayoutResponse { return v.ColumnLayout }).(ColumnLayoutResponseOutput)
}

// The mutable, human-readable name.
func (o LookupDashboardResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDashboardResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

// etag is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other. An etag is returned in the response to GetDashboard, and users are expected to put that etag in the request to UpdateDashboard to ensure that their change will be applied to the same version of the Dashboard configuration. The field should not be passed during dashboard creation.
func (o LookupDashboardResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDashboardResult) string { return v.Etag }).(pulumi.StringOutput)
}

// Content is arranged with a basic layout that re-flows a simple list of informational elements like widgets or tiles.
func (o LookupDashboardResultOutput) GridLayout() GridLayoutResponseOutput {
	return o.ApplyT(func(v LookupDashboardResult) GridLayoutResponse { return v.GridLayout }).(GridLayoutResponseOutput)
}

// Labels applied to the dashboard
func (o LookupDashboardResultOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupDashboardResult) map[string]string { return v.Labels }).(pulumi.StringMapOutput)
}

// The content is arranged as a grid of tiles, with each content widget occupying one or more grid blocks.
func (o LookupDashboardResultOutput) MosaicLayout() MosaicLayoutResponseOutput {
	return o.ApplyT(func(v LookupDashboardResult) MosaicLayoutResponse { return v.MosaicLayout }).(MosaicLayoutResponseOutput)
}

// Immutable. The resource name of the dashboard.
func (o LookupDashboardResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDashboardResult) string { return v.Name }).(pulumi.StringOutput)
}

// The content is divided into equally spaced rows and the widgets are arranged horizontally.
func (o LookupDashboardResultOutput) RowLayout() RowLayoutResponseOutput {
	return o.ApplyT(func(v LookupDashboardResult) RowLayoutResponse { return v.RowLayout }).(RowLayoutResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDashboardResultOutput{})
}
