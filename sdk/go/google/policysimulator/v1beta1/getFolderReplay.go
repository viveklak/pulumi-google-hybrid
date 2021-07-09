// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the specified Replay. Each `Replay` is available for at least 7 days.
func LookupFolderReplay(ctx *pulumi.Context, args *LookupFolderReplayArgs, opts ...pulumi.InvokeOption) (*LookupFolderReplayResult, error) {
	var rv LookupFolderReplayResult
	err := ctx.Invoke("google-native:policysimulator/v1beta1:getFolderReplay", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupFolderReplayArgs struct {
	FolderId string `pulumi:"folderId"`
	Location string `pulumi:"location"`
	ReplayId string `pulumi:"replayId"`
}

type LookupFolderReplayResult struct {
	// The configuration used for the `Replay`.
	Config GoogleCloudPolicysimulatorV1beta1ReplayConfigResponse `pulumi:"config"`
	// The resource name of the `Replay`, which has the following format: `{projects|folders|organizations}/{resource-id}/locations/global/replays/{replay-id}`, where `{resource-id}` is the ID of the project, folder, or organization that owns the Replay. Example: `projects/my-example-project/locations/global/replays/506a5f7f-38ce-4d7d-8e03-479ce1833c36`
	Name string `pulumi:"name"`
	// Summary statistics about the replayed log entries.
	ResultsSummary GoogleCloudPolicysimulatorV1beta1ReplayResultsSummaryResponse `pulumi:"resultsSummary"`
	// The current state of the `Replay`.
	State string `pulumi:"state"`
}
