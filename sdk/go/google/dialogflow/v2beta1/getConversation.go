// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v2beta1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieves the specific conversation.
func LookupConversation(ctx *pulumi.Context, args *LookupConversationArgs, opts ...pulumi.InvokeOption) (*LookupConversationResult, error) {
	var rv LookupConversationResult
	err := ctx.Invoke("google-native:dialogflow/v2beta1:getConversation", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupConversationArgs struct {
	ConversationId string  `pulumi:"conversationId"`
	Location       string  `pulumi:"location"`
	Project        *string `pulumi:"project"`
}

type LookupConversationResult struct {
	// The Conversation Profile to be used to configure this Conversation. This field cannot be updated. Format: `projects//locations//conversationProfiles/`.
	ConversationProfile string `pulumi:"conversationProfile"`
	// The stage of a conversation. It indicates whether the virtual agent or a human agent is handling the conversation. If the conversation is created with the conversation profile that has Dialogflow config set, defaults to ConversationStage.VIRTUAL_AGENT_STAGE; Otherwise, defaults to ConversationStage.HUMAN_ASSIST_STAGE. If the conversation is created with the conversation profile that has Dialogflow config set but explicitly sets conversation_stage to ConversationStage.HUMAN_ASSIST_STAGE, it skips ConversationStage.VIRTUAL_AGENT_STAGE stage and directly goes to ConversationStage.HUMAN_ASSIST_STAGE.
	ConversationStage string `pulumi:"conversationStage"`
	// The time the conversation was finished.
	EndTime string `pulumi:"endTime"`
	// The current state of the Conversation.
	LifecycleState string `pulumi:"lifecycleState"`
	// The unique identifier of this conversation. Format: `projects//locations//conversations/`.
	Name string `pulumi:"name"`
	// Required if the conversation is to be connected over telephony.
	PhoneNumber GoogleCloudDialogflowV2beta1ConversationPhoneNumberResponse `pulumi:"phoneNumber"`
	// The time the conversation was started.
	StartTime string `pulumi:"startTime"`
}

func LookupConversationOutput(ctx *pulumi.Context, args LookupConversationOutputArgs, opts ...pulumi.InvokeOption) LookupConversationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupConversationResult, error) {
			args := v.(LookupConversationArgs)
			r, err := LookupConversation(ctx, &args, opts...)
			return *r, err
		}).(LookupConversationResultOutput)
}

type LookupConversationOutputArgs struct {
	ConversationId pulumi.StringInput    `pulumi:"conversationId"`
	Location       pulumi.StringInput    `pulumi:"location"`
	Project        pulumi.StringPtrInput `pulumi:"project"`
}

func (LookupConversationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConversationArgs)(nil)).Elem()
}

type LookupConversationResultOutput struct{ *pulumi.OutputState }

func (LookupConversationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConversationResult)(nil)).Elem()
}

func (o LookupConversationResultOutput) ToLookupConversationResultOutput() LookupConversationResultOutput {
	return o
}

func (o LookupConversationResultOutput) ToLookupConversationResultOutputWithContext(ctx context.Context) LookupConversationResultOutput {
	return o
}

// The Conversation Profile to be used to configure this Conversation. This field cannot be updated. Format: `projects//locations//conversationProfiles/`.
func (o LookupConversationResultOutput) ConversationProfile() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConversationResult) string { return v.ConversationProfile }).(pulumi.StringOutput)
}

// The stage of a conversation. It indicates whether the virtual agent or a human agent is handling the conversation. If the conversation is created with the conversation profile that has Dialogflow config set, defaults to ConversationStage.VIRTUAL_AGENT_STAGE; Otherwise, defaults to ConversationStage.HUMAN_ASSIST_STAGE. If the conversation is created with the conversation profile that has Dialogflow config set but explicitly sets conversation_stage to ConversationStage.HUMAN_ASSIST_STAGE, it skips ConversationStage.VIRTUAL_AGENT_STAGE stage and directly goes to ConversationStage.HUMAN_ASSIST_STAGE.
func (o LookupConversationResultOutput) ConversationStage() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConversationResult) string { return v.ConversationStage }).(pulumi.StringOutput)
}

// The time the conversation was finished.
func (o LookupConversationResultOutput) EndTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConversationResult) string { return v.EndTime }).(pulumi.StringOutput)
}

// The current state of the Conversation.
func (o LookupConversationResultOutput) LifecycleState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConversationResult) string { return v.LifecycleState }).(pulumi.StringOutput)
}

// The unique identifier of this conversation. Format: `projects//locations//conversations/`.
func (o LookupConversationResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConversationResult) string { return v.Name }).(pulumi.StringOutput)
}

// Required if the conversation is to be connected over telephony.
func (o LookupConversationResultOutput) PhoneNumber() GoogleCloudDialogflowV2beta1ConversationPhoneNumberResponseOutput {
	return o.ApplyT(func(v LookupConversationResult) GoogleCloudDialogflowV2beta1ConversationPhoneNumberResponse {
		return v.PhoneNumber
	}).(GoogleCloudDialogflowV2beta1ConversationPhoneNumberResponseOutput)
}

// The time the conversation was started.
func (o LookupConversationResultOutput) StartTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConversationResult) string { return v.StartTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupConversationResultOutput{})
}
