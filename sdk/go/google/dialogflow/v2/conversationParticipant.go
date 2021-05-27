// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a new participant in a conversation.
type ConversationParticipant struct {
	pulumi.CustomResourceState

	// Optional. The unique identifier of this participant. Format: `projects//locations//conversations//participants/`.
	Name pulumi.StringOutput `pulumi:"name"`
	// Immutable. The role this participant plays in the conversation. This field must be set during participant creation and is then immutable.
	Role pulumi.StringOutput `pulumi:"role"`
	// Optional. Label applied to streams representing this participant in SIPREC XML metadata and SDP. This is used to assign transcriptions from that media stream to this participant. This field can be updated.
	SipRecordingMediaLabel pulumi.StringOutput `pulumi:"sipRecordingMediaLabel"`
}

// NewConversationParticipant registers a new resource with the given unique name, arguments, and options.
func NewConversationParticipant(ctx *pulumi.Context,
	name string, args *ConversationParticipantArgs, opts ...pulumi.ResourceOption) (*ConversationParticipant, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ConversationId == nil {
		return nil, errors.New("invalid value for required argument 'ConversationId'")
	}
	if args.Location == nil {
		return nil, errors.New("invalid value for required argument 'Location'")
	}
	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	var resource ConversationParticipant
	err := ctx.RegisterResource("google-native:dialogflow/v2:ConversationParticipant", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetConversationParticipant gets an existing ConversationParticipant resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetConversationParticipant(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConversationParticipantState, opts ...pulumi.ResourceOption) (*ConversationParticipant, error) {
	var resource ConversationParticipant
	err := ctx.ReadResource("google-native:dialogflow/v2:ConversationParticipant", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ConversationParticipant resources.
type conversationParticipantState struct {
	// Optional. The unique identifier of this participant. Format: `projects//locations//conversations//participants/`.
	Name *string `pulumi:"name"`
	// Immutable. The role this participant plays in the conversation. This field must be set during participant creation and is then immutable.
	Role *string `pulumi:"role"`
	// Optional. Label applied to streams representing this participant in SIPREC XML metadata and SDP. This is used to assign transcriptions from that media stream to this participant. This field can be updated.
	SipRecordingMediaLabel *string `pulumi:"sipRecordingMediaLabel"`
}

type ConversationParticipantState struct {
	// Optional. The unique identifier of this participant. Format: `projects//locations//conversations//participants/`.
	Name pulumi.StringPtrInput
	// Immutable. The role this participant plays in the conversation. This field must be set during participant creation and is then immutable.
	Role pulumi.StringPtrInput
	// Optional. Label applied to streams representing this participant in SIPREC XML metadata and SDP. This is used to assign transcriptions from that media stream to this participant. This field can be updated.
	SipRecordingMediaLabel pulumi.StringPtrInput
}

func (ConversationParticipantState) ElementType() reflect.Type {
	return reflect.TypeOf((*conversationParticipantState)(nil)).Elem()
}

type conversationParticipantArgs struct {
	ConversationId string `pulumi:"conversationId"`
	Location       string `pulumi:"location"`
	// Optional. The unique identifier of this participant. Format: `projects//locations//conversations//participants/`.
	Name    *string `pulumi:"name"`
	Project string  `pulumi:"project"`
	// Immutable. The role this participant plays in the conversation. This field must be set during participant creation and is then immutable.
	Role *string `pulumi:"role"`
	// Optional. Label applied to streams representing this participant in SIPREC XML metadata and SDP. This is used to assign transcriptions from that media stream to this participant. This field can be updated.
	SipRecordingMediaLabel *string `pulumi:"sipRecordingMediaLabel"`
}

// The set of arguments for constructing a ConversationParticipant resource.
type ConversationParticipantArgs struct {
	ConversationId pulumi.StringInput
	Location       pulumi.StringInput
	// Optional. The unique identifier of this participant. Format: `projects//locations//conversations//participants/`.
	Name    pulumi.StringPtrInput
	Project pulumi.StringInput
	// Immutable. The role this participant plays in the conversation. This field must be set during participant creation and is then immutable.
	Role pulumi.StringPtrInput
	// Optional. Label applied to streams representing this participant in SIPREC XML metadata and SDP. This is used to assign transcriptions from that media stream to this participant. This field can be updated.
	SipRecordingMediaLabel pulumi.StringPtrInput
}

func (ConversationParticipantArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*conversationParticipantArgs)(nil)).Elem()
}

type ConversationParticipantInput interface {
	pulumi.Input

	ToConversationParticipantOutput() ConversationParticipantOutput
	ToConversationParticipantOutputWithContext(ctx context.Context) ConversationParticipantOutput
}

func (*ConversationParticipant) ElementType() reflect.Type {
	return reflect.TypeOf((*ConversationParticipant)(nil))
}

func (i *ConversationParticipant) ToConversationParticipantOutput() ConversationParticipantOutput {
	return i.ToConversationParticipantOutputWithContext(context.Background())
}

func (i *ConversationParticipant) ToConversationParticipantOutputWithContext(ctx context.Context) ConversationParticipantOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConversationParticipantOutput)
}

type ConversationParticipantOutput struct {
	*pulumi.OutputState
}

func (ConversationParticipantOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConversationParticipant)(nil))
}

func (o ConversationParticipantOutput) ToConversationParticipantOutput() ConversationParticipantOutput {
	return o
}

func (o ConversationParticipantOutput) ToConversationParticipantOutputWithContext(ctx context.Context) ConversationParticipantOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ConversationParticipantOutput{})
}
