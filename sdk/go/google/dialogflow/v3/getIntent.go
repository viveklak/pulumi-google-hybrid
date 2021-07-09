// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v3

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieves the specified intent.
func LookupIntent(ctx *pulumi.Context, args *LookupIntentArgs, opts ...pulumi.InvokeOption) (*LookupIntentResult, error) {
	var rv LookupIntentResult
	err := ctx.Invoke("google-native:dialogflow/v3:getIntent", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupIntentArgs struct {
	AgentId      string  `pulumi:"agentId"`
	IntentId     string  `pulumi:"intentId"`
	LanguageCode *string `pulumi:"languageCode"`
	Location     string  `pulumi:"location"`
	Project      string  `pulumi:"project"`
}

type LookupIntentResult struct {
	// Human readable description for better understanding an intent like its scope, content, result etc. Maximum character limit: 140 characters.
	Description string `pulumi:"description"`
	// The human-readable name of the intent, unique within the agent.
	DisplayName string `pulumi:"displayName"`
	// Indicates whether this is a fallback intent. Currently only default fallback intent is allowed in the agent, which is added upon agent creation. Adding training phrases to fallback intent is useful in the case of requests that are mistakenly matched, since training phrases assigned to fallback intents act as negative examples that triggers no-match event.
	IsFallback bool `pulumi:"isFallback"`
	// The key/value metadata to label an intent. Labels can contain lowercase letters, digits and the symbols '-' and '_'. International characters are allowed, including letters from unicase alphabets. Keys must start with a letter. Keys and values can be no longer than 63 characters and no more than 128 bytes. Prefix "sys." is reserved for Dialogflow defined labels. Currently allowed Dialogflow defined labels include: * sys.head * sys.contextual The above labels do not require value. "sys.head" means the intent is a head intent. "sys.contextual" means the intent is a contextual intent.
	Labels map[string]string `pulumi:"labels"`
	// The unique identifier of the intent. Required for the Intents.UpdateIntent method. Intents.CreateIntent populates the name automatically. Format: `projects//locations//agents//intents/`.
	Name string `pulumi:"name"`
	// The collection of parameters associated with the intent.
	Parameters []GoogleCloudDialogflowCxV3IntentParameterResponse `pulumi:"parameters"`
	// The priority of this intent. Higher numbers represent higher priorities. - If the supplied value is unspecified or 0, the service translates the value to 500,000, which corresponds to the `Normal` priority in the console. - If the supplied value is negative, the intent is ignored in runtime detect intent requests.
	Priority int `pulumi:"priority"`
	// The collection of training phrases the agent is trained on to identify the intent.
	TrainingPhrases []GoogleCloudDialogflowCxV3IntentTrainingPhraseResponse `pulumi:"trainingPhrases"`
}
