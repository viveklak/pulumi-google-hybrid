// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v3beta1

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieves the specified Environment.
func LookupEnvironment(ctx *pulumi.Context, args *LookupEnvironmentArgs, opts ...pulumi.InvokeOption) (*LookupEnvironmentResult, error) {
	var rv LookupEnvironmentResult
	err := ctx.Invoke("google-native:dialogflow/v3beta1:getEnvironment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupEnvironmentArgs struct {
	AgentId       string `pulumi:"agentId"`
	EnvironmentId string `pulumi:"environmentId"`
	Location      string `pulumi:"location"`
	Project       string `pulumi:"project"`
}

type LookupEnvironmentResult struct {
	// The human-readable description of the environment. The maximum length is 500 characters. If exceeded, the request is rejected.
	Description string `pulumi:"description"`
	// The human-readable name of the environment (unique in an agent). Limit of 64 characters.
	DisplayName string `pulumi:"displayName"`
	// The name of the environment. Format: `projects//locations//agents//environments/`.
	Name string `pulumi:"name"`
	// Update time of this environment.
	UpdateTime string `pulumi:"updateTime"`
	// A list of configurations for flow versions. You should include version configs for all flows that are reachable from `Start Flow` in the agent. Otherwise, an error will be returned.
	VersionConfigs []GoogleCloudDialogflowCxV3beta1EnvironmentVersionConfigResponse `pulumi:"versionConfigs"`
}
