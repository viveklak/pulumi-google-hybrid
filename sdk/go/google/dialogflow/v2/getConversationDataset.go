// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieves the specified conversation dataset.
func LookupConversationDataset(ctx *pulumi.Context, args *LookupConversationDatasetArgs, opts ...pulumi.InvokeOption) (*LookupConversationDatasetResult, error) {
	var rv LookupConversationDatasetResult
	err := ctx.Invoke("google-native:dialogflow/v2:getConversationDataset", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupConversationDatasetArgs struct {
	ConversationDatasetId string  `pulumi:"conversationDatasetId"`
	Location              string  `pulumi:"location"`
	Project               *string `pulumi:"project"`
}

type LookupConversationDatasetResult struct {
	// The number of conversations this conversation dataset contains.
	ConversationCount string `pulumi:"conversationCount"`
	// Metadata set during conversation data import.
	ConversationInfo GoogleCloudDialogflowV2ConversationInfoResponse `pulumi:"conversationInfo"`
	// Creation time of this dataset.
	CreateTime string `pulumi:"createTime"`
	// Optional. The description of the dataset. Maximum of 10000 bytes.
	Description string `pulumi:"description"`
	// The display name of the dataset. Maximum of 64 bytes.
	DisplayName string `pulumi:"displayName"`
	// Input configurations set during conversation data import.
	InputConfig GoogleCloudDialogflowV2InputConfigResponse `pulumi:"inputConfig"`
	// ConversationDataset resource name. Format: `projects//locations//conversationDatasets/`
	Name string `pulumi:"name"`
}

func LookupConversationDatasetOutput(ctx *pulumi.Context, args LookupConversationDatasetOutputArgs, opts ...pulumi.InvokeOption) LookupConversationDatasetResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupConversationDatasetResult, error) {
			args := v.(LookupConversationDatasetArgs)
			r, err := LookupConversationDataset(ctx, &args, opts...)
			return *r, err
		}).(LookupConversationDatasetResultOutput)
}

type LookupConversationDatasetOutputArgs struct {
	ConversationDatasetId pulumi.StringInput    `pulumi:"conversationDatasetId"`
	Location              pulumi.StringInput    `pulumi:"location"`
	Project               pulumi.StringPtrInput `pulumi:"project"`
}

func (LookupConversationDatasetOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConversationDatasetArgs)(nil)).Elem()
}

type LookupConversationDatasetResultOutput struct{ *pulumi.OutputState }

func (LookupConversationDatasetResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConversationDatasetResult)(nil)).Elem()
}

func (o LookupConversationDatasetResultOutput) ToLookupConversationDatasetResultOutput() LookupConversationDatasetResultOutput {
	return o
}

func (o LookupConversationDatasetResultOutput) ToLookupConversationDatasetResultOutputWithContext(ctx context.Context) LookupConversationDatasetResultOutput {
	return o
}

// The number of conversations this conversation dataset contains.
func (o LookupConversationDatasetResultOutput) ConversationCount() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConversationDatasetResult) string { return v.ConversationCount }).(pulumi.StringOutput)
}

// Metadata set during conversation data import.
func (o LookupConversationDatasetResultOutput) ConversationInfo() GoogleCloudDialogflowV2ConversationInfoResponseOutput {
	return o.ApplyT(func(v LookupConversationDatasetResult) GoogleCloudDialogflowV2ConversationInfoResponse {
		return v.ConversationInfo
	}).(GoogleCloudDialogflowV2ConversationInfoResponseOutput)
}

// Creation time of this dataset.
func (o LookupConversationDatasetResultOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConversationDatasetResult) string { return v.CreateTime }).(pulumi.StringOutput)
}

// Optional. The description of the dataset. Maximum of 10000 bytes.
func (o LookupConversationDatasetResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConversationDatasetResult) string { return v.Description }).(pulumi.StringOutput)
}

// The display name of the dataset. Maximum of 64 bytes.
func (o LookupConversationDatasetResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConversationDatasetResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

// Input configurations set during conversation data import.
func (o LookupConversationDatasetResultOutput) InputConfig() GoogleCloudDialogflowV2InputConfigResponseOutput {
	return o.ApplyT(func(v LookupConversationDatasetResult) GoogleCloudDialogflowV2InputConfigResponse {
		return v.InputConfig
	}).(GoogleCloudDialogflowV2InputConfigResponseOutput)
}

// ConversationDataset resource name. Format: `projects//locations//conversationDatasets/`
func (o LookupConversationDatasetResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConversationDatasetResult) string { return v.Name }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupConversationDatasetResultOutput{})
}
