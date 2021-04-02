// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Represents preferences for sending email notifications for transfer run events.
type EmailPreferences struct {
	// If true, email notifications will be sent on transfer run failures.
	EnableFailureEmail *bool `pulumi:"enableFailureEmail"`
}

// EmailPreferencesInput is an input type that accepts EmailPreferencesArgs and EmailPreferencesOutput values.
// You can construct a concrete instance of `EmailPreferencesInput` via:
//
//          EmailPreferencesArgs{...}
type EmailPreferencesInput interface {
	pulumi.Input

	ToEmailPreferencesOutput() EmailPreferencesOutput
	ToEmailPreferencesOutputWithContext(context.Context) EmailPreferencesOutput
}

// Represents preferences for sending email notifications for transfer run events.
type EmailPreferencesArgs struct {
	// If true, email notifications will be sent on transfer run failures.
	EnableFailureEmail pulumi.BoolPtrInput `pulumi:"enableFailureEmail"`
}

func (EmailPreferencesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EmailPreferences)(nil)).Elem()
}

func (i EmailPreferencesArgs) ToEmailPreferencesOutput() EmailPreferencesOutput {
	return i.ToEmailPreferencesOutputWithContext(context.Background())
}

func (i EmailPreferencesArgs) ToEmailPreferencesOutputWithContext(ctx context.Context) EmailPreferencesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EmailPreferencesOutput)
}

func (i EmailPreferencesArgs) ToEmailPreferencesPtrOutput() EmailPreferencesPtrOutput {
	return i.ToEmailPreferencesPtrOutputWithContext(context.Background())
}

func (i EmailPreferencesArgs) ToEmailPreferencesPtrOutputWithContext(ctx context.Context) EmailPreferencesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EmailPreferencesOutput).ToEmailPreferencesPtrOutputWithContext(ctx)
}

// EmailPreferencesPtrInput is an input type that accepts EmailPreferencesArgs, EmailPreferencesPtr and EmailPreferencesPtrOutput values.
// You can construct a concrete instance of `EmailPreferencesPtrInput` via:
//
//          EmailPreferencesArgs{...}
//
//  or:
//
//          nil
type EmailPreferencesPtrInput interface {
	pulumi.Input

	ToEmailPreferencesPtrOutput() EmailPreferencesPtrOutput
	ToEmailPreferencesPtrOutputWithContext(context.Context) EmailPreferencesPtrOutput
}

type emailPreferencesPtrType EmailPreferencesArgs

func EmailPreferencesPtr(v *EmailPreferencesArgs) EmailPreferencesPtrInput {
	return (*emailPreferencesPtrType)(v)
}

func (*emailPreferencesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EmailPreferences)(nil)).Elem()
}

func (i *emailPreferencesPtrType) ToEmailPreferencesPtrOutput() EmailPreferencesPtrOutput {
	return i.ToEmailPreferencesPtrOutputWithContext(context.Background())
}

func (i *emailPreferencesPtrType) ToEmailPreferencesPtrOutputWithContext(ctx context.Context) EmailPreferencesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EmailPreferencesPtrOutput)
}

// Represents preferences for sending email notifications for transfer run events.
type EmailPreferencesOutput struct{ *pulumi.OutputState }

func (EmailPreferencesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EmailPreferences)(nil)).Elem()
}

func (o EmailPreferencesOutput) ToEmailPreferencesOutput() EmailPreferencesOutput {
	return o
}

func (o EmailPreferencesOutput) ToEmailPreferencesOutputWithContext(ctx context.Context) EmailPreferencesOutput {
	return o
}

func (o EmailPreferencesOutput) ToEmailPreferencesPtrOutput() EmailPreferencesPtrOutput {
	return o.ToEmailPreferencesPtrOutputWithContext(context.Background())
}

func (o EmailPreferencesOutput) ToEmailPreferencesPtrOutputWithContext(ctx context.Context) EmailPreferencesPtrOutput {
	return o.ApplyT(func(v EmailPreferences) *EmailPreferences {
		return &v
	}).(EmailPreferencesPtrOutput)
}

// If true, email notifications will be sent on transfer run failures.
func (o EmailPreferencesOutput) EnableFailureEmail() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v EmailPreferences) *bool { return v.EnableFailureEmail }).(pulumi.BoolPtrOutput)
}

type EmailPreferencesPtrOutput struct{ *pulumi.OutputState }

func (EmailPreferencesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EmailPreferences)(nil)).Elem()
}

func (o EmailPreferencesPtrOutput) ToEmailPreferencesPtrOutput() EmailPreferencesPtrOutput {
	return o
}

func (o EmailPreferencesPtrOutput) ToEmailPreferencesPtrOutputWithContext(ctx context.Context) EmailPreferencesPtrOutput {
	return o
}

func (o EmailPreferencesPtrOutput) Elem() EmailPreferencesOutput {
	return o.ApplyT(func(v *EmailPreferences) EmailPreferences { return *v }).(EmailPreferencesOutput)
}

// If true, email notifications will be sent on transfer run failures.
func (o EmailPreferencesPtrOutput) EnableFailureEmail() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EmailPreferences) *bool {
		if v == nil {
			return nil
		}
		return v.EnableFailureEmail
	}).(pulumi.BoolPtrOutput)
}

// Represents preferences for sending email notifications for transfer run events.
type EmailPreferencesResponse struct {
	// If true, email notifications will be sent on transfer run failures.
	EnableFailureEmail bool `pulumi:"enableFailureEmail"`
}

// EmailPreferencesResponseInput is an input type that accepts EmailPreferencesResponseArgs and EmailPreferencesResponseOutput values.
// You can construct a concrete instance of `EmailPreferencesResponseInput` via:
//
//          EmailPreferencesResponseArgs{...}
type EmailPreferencesResponseInput interface {
	pulumi.Input

	ToEmailPreferencesResponseOutput() EmailPreferencesResponseOutput
	ToEmailPreferencesResponseOutputWithContext(context.Context) EmailPreferencesResponseOutput
}

// Represents preferences for sending email notifications for transfer run events.
type EmailPreferencesResponseArgs struct {
	// If true, email notifications will be sent on transfer run failures.
	EnableFailureEmail pulumi.BoolInput `pulumi:"enableFailureEmail"`
}

func (EmailPreferencesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EmailPreferencesResponse)(nil)).Elem()
}

func (i EmailPreferencesResponseArgs) ToEmailPreferencesResponseOutput() EmailPreferencesResponseOutput {
	return i.ToEmailPreferencesResponseOutputWithContext(context.Background())
}

func (i EmailPreferencesResponseArgs) ToEmailPreferencesResponseOutputWithContext(ctx context.Context) EmailPreferencesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EmailPreferencesResponseOutput)
}

func (i EmailPreferencesResponseArgs) ToEmailPreferencesResponsePtrOutput() EmailPreferencesResponsePtrOutput {
	return i.ToEmailPreferencesResponsePtrOutputWithContext(context.Background())
}

func (i EmailPreferencesResponseArgs) ToEmailPreferencesResponsePtrOutputWithContext(ctx context.Context) EmailPreferencesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EmailPreferencesResponseOutput).ToEmailPreferencesResponsePtrOutputWithContext(ctx)
}

// EmailPreferencesResponsePtrInput is an input type that accepts EmailPreferencesResponseArgs, EmailPreferencesResponsePtr and EmailPreferencesResponsePtrOutput values.
// You can construct a concrete instance of `EmailPreferencesResponsePtrInput` via:
//
//          EmailPreferencesResponseArgs{...}
//
//  or:
//
//          nil
type EmailPreferencesResponsePtrInput interface {
	pulumi.Input

	ToEmailPreferencesResponsePtrOutput() EmailPreferencesResponsePtrOutput
	ToEmailPreferencesResponsePtrOutputWithContext(context.Context) EmailPreferencesResponsePtrOutput
}

type emailPreferencesResponsePtrType EmailPreferencesResponseArgs

func EmailPreferencesResponsePtr(v *EmailPreferencesResponseArgs) EmailPreferencesResponsePtrInput {
	return (*emailPreferencesResponsePtrType)(v)
}

func (*emailPreferencesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EmailPreferencesResponse)(nil)).Elem()
}

func (i *emailPreferencesResponsePtrType) ToEmailPreferencesResponsePtrOutput() EmailPreferencesResponsePtrOutput {
	return i.ToEmailPreferencesResponsePtrOutputWithContext(context.Background())
}

func (i *emailPreferencesResponsePtrType) ToEmailPreferencesResponsePtrOutputWithContext(ctx context.Context) EmailPreferencesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EmailPreferencesResponsePtrOutput)
}

// Represents preferences for sending email notifications for transfer run events.
type EmailPreferencesResponseOutput struct{ *pulumi.OutputState }

func (EmailPreferencesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EmailPreferencesResponse)(nil)).Elem()
}

func (o EmailPreferencesResponseOutput) ToEmailPreferencesResponseOutput() EmailPreferencesResponseOutput {
	return o
}

func (o EmailPreferencesResponseOutput) ToEmailPreferencesResponseOutputWithContext(ctx context.Context) EmailPreferencesResponseOutput {
	return o
}

func (o EmailPreferencesResponseOutput) ToEmailPreferencesResponsePtrOutput() EmailPreferencesResponsePtrOutput {
	return o.ToEmailPreferencesResponsePtrOutputWithContext(context.Background())
}

func (o EmailPreferencesResponseOutput) ToEmailPreferencesResponsePtrOutputWithContext(ctx context.Context) EmailPreferencesResponsePtrOutput {
	return o.ApplyT(func(v EmailPreferencesResponse) *EmailPreferencesResponse {
		return &v
	}).(EmailPreferencesResponsePtrOutput)
}

// If true, email notifications will be sent on transfer run failures.
func (o EmailPreferencesResponseOutput) EnableFailureEmail() pulumi.BoolOutput {
	return o.ApplyT(func(v EmailPreferencesResponse) bool { return v.EnableFailureEmail }).(pulumi.BoolOutput)
}

type EmailPreferencesResponsePtrOutput struct{ *pulumi.OutputState }

func (EmailPreferencesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EmailPreferencesResponse)(nil)).Elem()
}

func (o EmailPreferencesResponsePtrOutput) ToEmailPreferencesResponsePtrOutput() EmailPreferencesResponsePtrOutput {
	return o
}

func (o EmailPreferencesResponsePtrOutput) ToEmailPreferencesResponsePtrOutputWithContext(ctx context.Context) EmailPreferencesResponsePtrOutput {
	return o
}

func (o EmailPreferencesResponsePtrOutput) Elem() EmailPreferencesResponseOutput {
	return o.ApplyT(func(v *EmailPreferencesResponse) EmailPreferencesResponse { return *v }).(EmailPreferencesResponseOutput)
}

// If true, email notifications will be sent on transfer run failures.
func (o EmailPreferencesResponsePtrOutput) EnableFailureEmail() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EmailPreferencesResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.EnableFailureEmail
	}).(pulumi.BoolPtrOutput)
}

// Options customizing the data transfer schedule.
type ScheduleOptions struct {
	// If true, automatic scheduling of data transfer runs for this configuration will be disabled. The runs can be started on ad-hoc basis using StartManualTransferRuns API. When automatic scheduling is disabled, the TransferConfig.schedule field will be ignored.
	DisableAutoScheduling *bool `pulumi:"disableAutoScheduling"`
	// Defines time to stop scheduling transfer runs. A transfer run cannot be scheduled at or after the end time. The end time can be changed at any moment. The time when a data transfer can be trigerred manually is not limited by this option.
	EndTime *string `pulumi:"endTime"`
	// Specifies time to start scheduling transfer runs. The first run will be scheduled at or after the start time according to a recurrence pattern defined in the schedule string. The start time can be changed at any moment. The time when a data transfer can be trigerred manually is not limited by this option.
	StartTime *string `pulumi:"startTime"`
}

// ScheduleOptionsInput is an input type that accepts ScheduleOptionsArgs and ScheduleOptionsOutput values.
// You can construct a concrete instance of `ScheduleOptionsInput` via:
//
//          ScheduleOptionsArgs{...}
type ScheduleOptionsInput interface {
	pulumi.Input

	ToScheduleOptionsOutput() ScheduleOptionsOutput
	ToScheduleOptionsOutputWithContext(context.Context) ScheduleOptionsOutput
}

// Options customizing the data transfer schedule.
type ScheduleOptionsArgs struct {
	// If true, automatic scheduling of data transfer runs for this configuration will be disabled. The runs can be started on ad-hoc basis using StartManualTransferRuns API. When automatic scheduling is disabled, the TransferConfig.schedule field will be ignored.
	DisableAutoScheduling pulumi.BoolPtrInput `pulumi:"disableAutoScheduling"`
	// Defines time to stop scheduling transfer runs. A transfer run cannot be scheduled at or after the end time. The end time can be changed at any moment. The time when a data transfer can be trigerred manually is not limited by this option.
	EndTime pulumi.StringPtrInput `pulumi:"endTime"`
	// Specifies time to start scheduling transfer runs. The first run will be scheduled at or after the start time according to a recurrence pattern defined in the schedule string. The start time can be changed at any moment. The time when a data transfer can be trigerred manually is not limited by this option.
	StartTime pulumi.StringPtrInput `pulumi:"startTime"`
}

func (ScheduleOptionsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ScheduleOptions)(nil)).Elem()
}

func (i ScheduleOptionsArgs) ToScheduleOptionsOutput() ScheduleOptionsOutput {
	return i.ToScheduleOptionsOutputWithContext(context.Background())
}

func (i ScheduleOptionsArgs) ToScheduleOptionsOutputWithContext(ctx context.Context) ScheduleOptionsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScheduleOptionsOutput)
}

func (i ScheduleOptionsArgs) ToScheduleOptionsPtrOutput() ScheduleOptionsPtrOutput {
	return i.ToScheduleOptionsPtrOutputWithContext(context.Background())
}

func (i ScheduleOptionsArgs) ToScheduleOptionsPtrOutputWithContext(ctx context.Context) ScheduleOptionsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScheduleOptionsOutput).ToScheduleOptionsPtrOutputWithContext(ctx)
}

// ScheduleOptionsPtrInput is an input type that accepts ScheduleOptionsArgs, ScheduleOptionsPtr and ScheduleOptionsPtrOutput values.
// You can construct a concrete instance of `ScheduleOptionsPtrInput` via:
//
//          ScheduleOptionsArgs{...}
//
//  or:
//
//          nil
type ScheduleOptionsPtrInput interface {
	pulumi.Input

	ToScheduleOptionsPtrOutput() ScheduleOptionsPtrOutput
	ToScheduleOptionsPtrOutputWithContext(context.Context) ScheduleOptionsPtrOutput
}

type scheduleOptionsPtrType ScheduleOptionsArgs

func ScheduleOptionsPtr(v *ScheduleOptionsArgs) ScheduleOptionsPtrInput {
	return (*scheduleOptionsPtrType)(v)
}

func (*scheduleOptionsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ScheduleOptions)(nil)).Elem()
}

func (i *scheduleOptionsPtrType) ToScheduleOptionsPtrOutput() ScheduleOptionsPtrOutput {
	return i.ToScheduleOptionsPtrOutputWithContext(context.Background())
}

func (i *scheduleOptionsPtrType) ToScheduleOptionsPtrOutputWithContext(ctx context.Context) ScheduleOptionsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScheduleOptionsPtrOutput)
}

// Options customizing the data transfer schedule.
type ScheduleOptionsOutput struct{ *pulumi.OutputState }

func (ScheduleOptionsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScheduleOptions)(nil)).Elem()
}

func (o ScheduleOptionsOutput) ToScheduleOptionsOutput() ScheduleOptionsOutput {
	return o
}

func (o ScheduleOptionsOutput) ToScheduleOptionsOutputWithContext(ctx context.Context) ScheduleOptionsOutput {
	return o
}

func (o ScheduleOptionsOutput) ToScheduleOptionsPtrOutput() ScheduleOptionsPtrOutput {
	return o.ToScheduleOptionsPtrOutputWithContext(context.Background())
}

func (o ScheduleOptionsOutput) ToScheduleOptionsPtrOutputWithContext(ctx context.Context) ScheduleOptionsPtrOutput {
	return o.ApplyT(func(v ScheduleOptions) *ScheduleOptions {
		return &v
	}).(ScheduleOptionsPtrOutput)
}

// If true, automatic scheduling of data transfer runs for this configuration will be disabled. The runs can be started on ad-hoc basis using StartManualTransferRuns API. When automatic scheduling is disabled, the TransferConfig.schedule field will be ignored.
func (o ScheduleOptionsOutput) DisableAutoScheduling() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ScheduleOptions) *bool { return v.DisableAutoScheduling }).(pulumi.BoolPtrOutput)
}

// Defines time to stop scheduling transfer runs. A transfer run cannot be scheduled at or after the end time. The end time can be changed at any moment. The time when a data transfer can be trigerred manually is not limited by this option.
func (o ScheduleOptionsOutput) EndTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScheduleOptions) *string { return v.EndTime }).(pulumi.StringPtrOutput)
}

// Specifies time to start scheduling transfer runs. The first run will be scheduled at or after the start time according to a recurrence pattern defined in the schedule string. The start time can be changed at any moment. The time when a data transfer can be trigerred manually is not limited by this option.
func (o ScheduleOptionsOutput) StartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScheduleOptions) *string { return v.StartTime }).(pulumi.StringPtrOutput)
}

type ScheduleOptionsPtrOutput struct{ *pulumi.OutputState }

func (ScheduleOptionsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ScheduleOptions)(nil)).Elem()
}

func (o ScheduleOptionsPtrOutput) ToScheduleOptionsPtrOutput() ScheduleOptionsPtrOutput {
	return o
}

func (o ScheduleOptionsPtrOutput) ToScheduleOptionsPtrOutputWithContext(ctx context.Context) ScheduleOptionsPtrOutput {
	return o
}

func (o ScheduleOptionsPtrOutput) Elem() ScheduleOptionsOutput {
	return o.ApplyT(func(v *ScheduleOptions) ScheduleOptions { return *v }).(ScheduleOptionsOutput)
}

// If true, automatic scheduling of data transfer runs for this configuration will be disabled. The runs can be started on ad-hoc basis using StartManualTransferRuns API. When automatic scheduling is disabled, the TransferConfig.schedule field will be ignored.
func (o ScheduleOptionsPtrOutput) DisableAutoScheduling() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ScheduleOptions) *bool {
		if v == nil {
			return nil
		}
		return v.DisableAutoScheduling
	}).(pulumi.BoolPtrOutput)
}

// Defines time to stop scheduling transfer runs. A transfer run cannot be scheduled at or after the end time. The end time can be changed at any moment. The time when a data transfer can be trigerred manually is not limited by this option.
func (o ScheduleOptionsPtrOutput) EndTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScheduleOptions) *string {
		if v == nil {
			return nil
		}
		return v.EndTime
	}).(pulumi.StringPtrOutput)
}

// Specifies time to start scheduling transfer runs. The first run will be scheduled at or after the start time according to a recurrence pattern defined in the schedule string. The start time can be changed at any moment. The time when a data transfer can be trigerred manually is not limited by this option.
func (o ScheduleOptionsPtrOutput) StartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScheduleOptions) *string {
		if v == nil {
			return nil
		}
		return v.StartTime
	}).(pulumi.StringPtrOutput)
}

// Options customizing the data transfer schedule.
type ScheduleOptionsResponse struct {
	// If true, automatic scheduling of data transfer runs for this configuration will be disabled. The runs can be started on ad-hoc basis using StartManualTransferRuns API. When automatic scheduling is disabled, the TransferConfig.schedule field will be ignored.
	DisableAutoScheduling bool `pulumi:"disableAutoScheduling"`
	// Defines time to stop scheduling transfer runs. A transfer run cannot be scheduled at or after the end time. The end time can be changed at any moment. The time when a data transfer can be trigerred manually is not limited by this option.
	EndTime string `pulumi:"endTime"`
	// Specifies time to start scheduling transfer runs. The first run will be scheduled at or after the start time according to a recurrence pattern defined in the schedule string. The start time can be changed at any moment. The time when a data transfer can be trigerred manually is not limited by this option.
	StartTime string `pulumi:"startTime"`
}

// ScheduleOptionsResponseInput is an input type that accepts ScheduleOptionsResponseArgs and ScheduleOptionsResponseOutput values.
// You can construct a concrete instance of `ScheduleOptionsResponseInput` via:
//
//          ScheduleOptionsResponseArgs{...}
type ScheduleOptionsResponseInput interface {
	pulumi.Input

	ToScheduleOptionsResponseOutput() ScheduleOptionsResponseOutput
	ToScheduleOptionsResponseOutputWithContext(context.Context) ScheduleOptionsResponseOutput
}

// Options customizing the data transfer schedule.
type ScheduleOptionsResponseArgs struct {
	// If true, automatic scheduling of data transfer runs for this configuration will be disabled. The runs can be started on ad-hoc basis using StartManualTransferRuns API. When automatic scheduling is disabled, the TransferConfig.schedule field will be ignored.
	DisableAutoScheduling pulumi.BoolInput `pulumi:"disableAutoScheduling"`
	// Defines time to stop scheduling transfer runs. A transfer run cannot be scheduled at or after the end time. The end time can be changed at any moment. The time when a data transfer can be trigerred manually is not limited by this option.
	EndTime pulumi.StringInput `pulumi:"endTime"`
	// Specifies time to start scheduling transfer runs. The first run will be scheduled at or after the start time according to a recurrence pattern defined in the schedule string. The start time can be changed at any moment. The time when a data transfer can be trigerred manually is not limited by this option.
	StartTime pulumi.StringInput `pulumi:"startTime"`
}

func (ScheduleOptionsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ScheduleOptionsResponse)(nil)).Elem()
}

func (i ScheduleOptionsResponseArgs) ToScheduleOptionsResponseOutput() ScheduleOptionsResponseOutput {
	return i.ToScheduleOptionsResponseOutputWithContext(context.Background())
}

func (i ScheduleOptionsResponseArgs) ToScheduleOptionsResponseOutputWithContext(ctx context.Context) ScheduleOptionsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScheduleOptionsResponseOutput)
}

func (i ScheduleOptionsResponseArgs) ToScheduleOptionsResponsePtrOutput() ScheduleOptionsResponsePtrOutput {
	return i.ToScheduleOptionsResponsePtrOutputWithContext(context.Background())
}

func (i ScheduleOptionsResponseArgs) ToScheduleOptionsResponsePtrOutputWithContext(ctx context.Context) ScheduleOptionsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScheduleOptionsResponseOutput).ToScheduleOptionsResponsePtrOutputWithContext(ctx)
}

// ScheduleOptionsResponsePtrInput is an input type that accepts ScheduleOptionsResponseArgs, ScheduleOptionsResponsePtr and ScheduleOptionsResponsePtrOutput values.
// You can construct a concrete instance of `ScheduleOptionsResponsePtrInput` via:
//
//          ScheduleOptionsResponseArgs{...}
//
//  or:
//
//          nil
type ScheduleOptionsResponsePtrInput interface {
	pulumi.Input

	ToScheduleOptionsResponsePtrOutput() ScheduleOptionsResponsePtrOutput
	ToScheduleOptionsResponsePtrOutputWithContext(context.Context) ScheduleOptionsResponsePtrOutput
}

type scheduleOptionsResponsePtrType ScheduleOptionsResponseArgs

func ScheduleOptionsResponsePtr(v *ScheduleOptionsResponseArgs) ScheduleOptionsResponsePtrInput {
	return (*scheduleOptionsResponsePtrType)(v)
}

func (*scheduleOptionsResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ScheduleOptionsResponse)(nil)).Elem()
}

func (i *scheduleOptionsResponsePtrType) ToScheduleOptionsResponsePtrOutput() ScheduleOptionsResponsePtrOutput {
	return i.ToScheduleOptionsResponsePtrOutputWithContext(context.Background())
}

func (i *scheduleOptionsResponsePtrType) ToScheduleOptionsResponsePtrOutputWithContext(ctx context.Context) ScheduleOptionsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScheduleOptionsResponsePtrOutput)
}

// Options customizing the data transfer schedule.
type ScheduleOptionsResponseOutput struct{ *pulumi.OutputState }

func (ScheduleOptionsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScheduleOptionsResponse)(nil)).Elem()
}

func (o ScheduleOptionsResponseOutput) ToScheduleOptionsResponseOutput() ScheduleOptionsResponseOutput {
	return o
}

func (o ScheduleOptionsResponseOutput) ToScheduleOptionsResponseOutputWithContext(ctx context.Context) ScheduleOptionsResponseOutput {
	return o
}

func (o ScheduleOptionsResponseOutput) ToScheduleOptionsResponsePtrOutput() ScheduleOptionsResponsePtrOutput {
	return o.ToScheduleOptionsResponsePtrOutputWithContext(context.Background())
}

func (o ScheduleOptionsResponseOutput) ToScheduleOptionsResponsePtrOutputWithContext(ctx context.Context) ScheduleOptionsResponsePtrOutput {
	return o.ApplyT(func(v ScheduleOptionsResponse) *ScheduleOptionsResponse {
		return &v
	}).(ScheduleOptionsResponsePtrOutput)
}

// If true, automatic scheduling of data transfer runs for this configuration will be disabled. The runs can be started on ad-hoc basis using StartManualTransferRuns API. When automatic scheduling is disabled, the TransferConfig.schedule field will be ignored.
func (o ScheduleOptionsResponseOutput) DisableAutoScheduling() pulumi.BoolOutput {
	return o.ApplyT(func(v ScheduleOptionsResponse) bool { return v.DisableAutoScheduling }).(pulumi.BoolOutput)
}

// Defines time to stop scheduling transfer runs. A transfer run cannot be scheduled at or after the end time. The end time can be changed at any moment. The time when a data transfer can be trigerred manually is not limited by this option.
func (o ScheduleOptionsResponseOutput) EndTime() pulumi.StringOutput {
	return o.ApplyT(func(v ScheduleOptionsResponse) string { return v.EndTime }).(pulumi.StringOutput)
}

// Specifies time to start scheduling transfer runs. The first run will be scheduled at or after the start time according to a recurrence pattern defined in the schedule string. The start time can be changed at any moment. The time when a data transfer can be trigerred manually is not limited by this option.
func (o ScheduleOptionsResponseOutput) StartTime() pulumi.StringOutput {
	return o.ApplyT(func(v ScheduleOptionsResponse) string { return v.StartTime }).(pulumi.StringOutput)
}

type ScheduleOptionsResponsePtrOutput struct{ *pulumi.OutputState }

func (ScheduleOptionsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ScheduleOptionsResponse)(nil)).Elem()
}

func (o ScheduleOptionsResponsePtrOutput) ToScheduleOptionsResponsePtrOutput() ScheduleOptionsResponsePtrOutput {
	return o
}

func (o ScheduleOptionsResponsePtrOutput) ToScheduleOptionsResponsePtrOutputWithContext(ctx context.Context) ScheduleOptionsResponsePtrOutput {
	return o
}

func (o ScheduleOptionsResponsePtrOutput) Elem() ScheduleOptionsResponseOutput {
	return o.ApplyT(func(v *ScheduleOptionsResponse) ScheduleOptionsResponse { return *v }).(ScheduleOptionsResponseOutput)
}

// If true, automatic scheduling of data transfer runs for this configuration will be disabled. The runs can be started on ad-hoc basis using StartManualTransferRuns API. When automatic scheduling is disabled, the TransferConfig.schedule field will be ignored.
func (o ScheduleOptionsResponsePtrOutput) DisableAutoScheduling() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ScheduleOptionsResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.DisableAutoScheduling
	}).(pulumi.BoolPtrOutput)
}

// Defines time to stop scheduling transfer runs. A transfer run cannot be scheduled at or after the end time. The end time can be changed at any moment. The time when a data transfer can be trigerred manually is not limited by this option.
func (o ScheduleOptionsResponsePtrOutput) EndTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScheduleOptionsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.EndTime
	}).(pulumi.StringPtrOutput)
}

// Specifies time to start scheduling transfer runs. The first run will be scheduled at or after the start time according to a recurrence pattern defined in the schedule string. The start time can be changed at any moment. The time when a data transfer can be trigerred manually is not limited by this option.
func (o ScheduleOptionsResponsePtrOutput) StartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScheduleOptionsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.StartTime
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(EmailPreferencesOutput{})
	pulumi.RegisterOutputType(EmailPreferencesPtrOutput{})
	pulumi.RegisterOutputType(EmailPreferencesResponseOutput{})
	pulumi.RegisterOutputType(EmailPreferencesResponsePtrOutput{})
	pulumi.RegisterOutputType(ScheduleOptionsOutput{})
	pulumi.RegisterOutputType(ScheduleOptionsPtrOutput{})
	pulumi.RegisterOutputType(ScheduleOptionsResponseOutput{})
	pulumi.RegisterOutputType(ScheduleOptionsResponsePtrOutput{})
}