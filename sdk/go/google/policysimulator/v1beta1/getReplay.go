// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the specified Replay. Each `Replay` is available for at least 7 days.
func LookupReplay(ctx *pulumi.Context, args *LookupReplayArgs, opts ...pulumi.InvokeOption) (*LookupReplayResult, error) {
	var rv LookupReplayResult
	err := ctx.Invoke("google-native:policysimulator/v1beta1:getReplay", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupReplayArgs struct {
	Location string  `pulumi:"location"`
	Project  *string `pulumi:"project"`
	ReplayId string  `pulumi:"replayId"`
}

type LookupReplayResult struct {
	// The configuration used for the `Replay`.
	Config GoogleCloudPolicysimulatorV1beta1ReplayConfigResponse `pulumi:"config"`
	// The resource name of the `Replay`, which has the following format: `{projects|folders|organizations}/{resource-id}/locations/global/replays/{replay-id}`, where `{resource-id}` is the ID of the project, folder, or organization that owns the Replay. Example: `projects/my-example-project/locations/global/replays/506a5f7f-38ce-4d7d-8e03-479ce1833c36`
	Name string `pulumi:"name"`
	// Summary statistics about the replayed log entries.
	ResultsSummary GoogleCloudPolicysimulatorV1beta1ReplayResultsSummaryResponse `pulumi:"resultsSummary"`
	// The current state of the `Replay`.
	State string `pulumi:"state"`
}

func LookupReplayOutput(ctx *pulumi.Context, args LookupReplayOutputArgs, opts ...pulumi.InvokeOption) LookupReplayResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupReplayResult, error) {
			args := v.(LookupReplayArgs)
			r, err := LookupReplay(ctx, &args, opts...)
			return *r, err
		}).(LookupReplayResultOutput)
}

type LookupReplayOutputArgs struct {
	Location pulumi.StringInput    `pulumi:"location"`
	Project  pulumi.StringPtrInput `pulumi:"project"`
	ReplayId pulumi.StringInput    `pulumi:"replayId"`
}

func (LookupReplayOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupReplayArgs)(nil)).Elem()
}

type LookupReplayResultOutput struct{ *pulumi.OutputState }

func (LookupReplayResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupReplayResult)(nil)).Elem()
}

func (o LookupReplayResultOutput) ToLookupReplayResultOutput() LookupReplayResultOutput {
	return o
}

func (o LookupReplayResultOutput) ToLookupReplayResultOutputWithContext(ctx context.Context) LookupReplayResultOutput {
	return o
}

// The configuration used for the `Replay`.
func (o LookupReplayResultOutput) Config() GoogleCloudPolicysimulatorV1beta1ReplayConfigResponseOutput {
	return o.ApplyT(func(v LookupReplayResult) GoogleCloudPolicysimulatorV1beta1ReplayConfigResponse { return v.Config }).(GoogleCloudPolicysimulatorV1beta1ReplayConfigResponseOutput)
}

// The resource name of the `Replay`, which has the following format: `{projects|folders|organizations}/{resource-id}/locations/global/replays/{replay-id}`, where `{resource-id}` is the ID of the project, folder, or organization that owns the Replay. Example: `projects/my-example-project/locations/global/replays/506a5f7f-38ce-4d7d-8e03-479ce1833c36`
func (o LookupReplayResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReplayResult) string { return v.Name }).(pulumi.StringOutput)
}

// Summary statistics about the replayed log entries.
func (o LookupReplayResultOutput) ResultsSummary() GoogleCloudPolicysimulatorV1beta1ReplayResultsSummaryResponseOutput {
	return o.ApplyT(func(v LookupReplayResult) GoogleCloudPolicysimulatorV1beta1ReplayResultsSummaryResponse {
		return v.ResultsSummary
	}).(GoogleCloudPolicysimulatorV1beta1ReplayResultsSummaryResponseOutput)
}

// The current state of the `Replay`.
func (o LookupReplayResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReplayResult) string { return v.State }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupReplayResultOutput{})
}