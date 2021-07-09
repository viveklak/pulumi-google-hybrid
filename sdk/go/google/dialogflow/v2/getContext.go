// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieves the specified context.
func LookupContext(ctx *pulumi.Context, args *LookupContextArgs, opts ...pulumi.InvokeOption) (*LookupContextResult, error) {
	var rv LookupContextResult
	err := ctx.Invoke("google-native:dialogflow/v2:getContext", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupContextArgs struct {
	ContextId     string `pulumi:"contextId"`
	EnvironmentId string `pulumi:"environmentId"`
	Location      string `pulumi:"location"`
	Project       string `pulumi:"project"`
	SessionId     string `pulumi:"sessionId"`
	UserId        string `pulumi:"userId"`
}

type LookupContextResult struct {
	// Optional. The number of conversational query requests after which the context expires. The default is `0`. If set to `0`, the context expires immediately. Contexts expire automatically after 20 minutes if there are no matching queries.
	LifespanCount int `pulumi:"lifespanCount"`
	// The unique identifier of the context. Format: `projects//agent/sessions//contexts/`, or `projects//agent/environments//users//sessions//contexts/`. The `Context ID` is always converted to lowercase, may only contain characters in a-zA-Z0-9_-% and may be at most 250 bytes long. If `Environment ID` is not specified, we assume default 'draft' environment. If `User ID` is not specified, we assume default '-' user. The following context names are reserved for internal use by Dialogflow. You should not use these contexts or create contexts with these names: * `__system_counters__` * `*_id_dialog_context` * `*_dialog_params_size`
	Name string `pulumi:"name"`
	// Optional. The collection of parameters associated with this context. Depending on your protocol or client library language, this is a map, associative array, symbol table, dictionary, or JSON object composed of a collection of (MapKey, MapValue) pairs: - MapKey type: string - MapKey value: parameter name - MapValue type: - If parameter's entity type is a composite entity: map - Else: depending on parameter value type, could be one of string, number, boolean, null, list or map - MapValue value: - If parameter's entity type is a composite entity: map from composite entity property names to property values - Else: parameter value
	Parameters map[string]string `pulumi:"parameters"`
}
