// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieves a conversation participant.
func LookupParticipant(ctx *pulumi.Context, args *LookupParticipantArgs, opts ...pulumi.InvokeOption) (*LookupParticipantResult, error) {
	var rv LookupParticipantResult
	err := ctx.Invoke("google-native:dialogflow/v2:getParticipant", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupParticipantArgs struct {
	ConversationId string  `pulumi:"conversationId"`
	Location       string  `pulumi:"location"`
	ParticipantId  string  `pulumi:"participantId"`
	Project        *string `pulumi:"project"`
}

type LookupParticipantResult struct {
	// Optional. Key-value filters on the metadata of documents returned by article suggestion. If specified, article suggestion only returns suggested documents that match all filters in their Document.metadata. Multiple values for a metadata key should be concatenated by comma. For example, filters to match all documents that have 'US' or 'CA' in their market metadata values and 'agent' in their user metadata values will be ```documents_metadata_filters { key: "market" value: "US,CA" } documents_metadata_filters { key: "user" value: "agent" }```
	DocumentsMetadataFilters map[string]string `pulumi:"documentsMetadataFilters"`
	// Optional. The unique identifier of this participant. Format: `projects//locations//conversations//participants/`.
	Name string `pulumi:"name"`
	// Immutable. The role this participant plays in the conversation. This field must be set during participant creation and is then immutable.
	Role string `pulumi:"role"`
	// Optional. Label applied to streams representing this participant in SIPREC XML metadata and SDP. This is used to assign transcriptions from that media stream to this participant. This field can be updated.
	SipRecordingMediaLabel string `pulumi:"sipRecordingMediaLabel"`
}

func LookupParticipantOutput(ctx *pulumi.Context, args LookupParticipantOutputArgs, opts ...pulumi.InvokeOption) LookupParticipantResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupParticipantResult, error) {
			args := v.(LookupParticipantArgs)
			r, err := LookupParticipant(ctx, &args, opts...)
			return *r, err
		}).(LookupParticipantResultOutput)
}

type LookupParticipantOutputArgs struct {
	ConversationId pulumi.StringInput    `pulumi:"conversationId"`
	Location       pulumi.StringInput    `pulumi:"location"`
	ParticipantId  pulumi.StringInput    `pulumi:"participantId"`
	Project        pulumi.StringPtrInput `pulumi:"project"`
}

func (LookupParticipantOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupParticipantArgs)(nil)).Elem()
}

type LookupParticipantResultOutput struct{ *pulumi.OutputState }

func (LookupParticipantResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupParticipantResult)(nil)).Elem()
}

func (o LookupParticipantResultOutput) ToLookupParticipantResultOutput() LookupParticipantResultOutput {
	return o
}

func (o LookupParticipantResultOutput) ToLookupParticipantResultOutputWithContext(ctx context.Context) LookupParticipantResultOutput {
	return o
}

// Optional. Key-value filters on the metadata of documents returned by article suggestion. If specified, article suggestion only returns suggested documents that match all filters in their Document.metadata. Multiple values for a metadata key should be concatenated by comma. For example, filters to match all documents that have 'US' or 'CA' in their market metadata values and 'agent' in their user metadata values will be ```documents_metadata_filters { key: "market" value: "US,CA" } documents_metadata_filters { key: "user" value: "agent" }```
func (o LookupParticipantResultOutput) DocumentsMetadataFilters() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupParticipantResult) map[string]string { return v.DocumentsMetadataFilters }).(pulumi.StringMapOutput)
}

// Optional. The unique identifier of this participant. Format: `projects//locations//conversations//participants/`.
func (o LookupParticipantResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupParticipantResult) string { return v.Name }).(pulumi.StringOutput)
}

// Immutable. The role this participant plays in the conversation. This field must be set during participant creation and is then immutable.
func (o LookupParticipantResultOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v LookupParticipantResult) string { return v.Role }).(pulumi.StringOutput)
}

// Optional. Label applied to streams representing this participant in SIPREC XML metadata and SDP. This is used to assign transcriptions from that media stream to this participant. This field can be updated.
func (o LookupParticipantResultOutput) SipRecordingMediaLabel() pulumi.StringOutput {
	return o.ApplyT(func(v LookupParticipantResult) string { return v.SipRecordingMediaLabel }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupParticipantResultOutput{})
}
