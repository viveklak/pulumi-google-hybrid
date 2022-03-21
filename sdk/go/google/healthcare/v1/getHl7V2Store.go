// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the specified HL7v2 store.
func LookupHl7V2Store(ctx *pulumi.Context, args *LookupHl7V2StoreArgs, opts ...pulumi.InvokeOption) (*LookupHl7V2StoreResult, error) {
	var rv LookupHl7V2StoreResult
	err := ctx.Invoke("google-native:healthcare/v1:getHl7V2Store", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupHl7V2StoreArgs struct {
	DatasetId    string  `pulumi:"datasetId"`
	Hl7V2StoreId string  `pulumi:"hl7V2StoreId"`
	Location     string  `pulumi:"location"`
	Project      *string `pulumi:"project"`
}

type LookupHl7V2StoreResult struct {
	// User-supplied key-value pairs used to organize HL7v2 stores. Label keys must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must conform to the following PCRE regular expression: \p{Ll}\p{Lo}{0,62} Label values are optional, must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must conform to the following PCRE regular expression: [\p{Ll}\p{Lo}\p{N}_-]{0,63} No more than 64 labels can be associated with a given store.
	Labels map[string]string `pulumi:"labels"`
	// Resource name of the HL7v2 store, of the form `projects/{project_id}/datasets/{dataset_id}/hl7V2Stores/{hl7v2_store_id}`.
	Name string `pulumi:"name"`
	// A list of notification configs. Each configuration uses a filter to determine whether to publish a message (both Ingest & Create) on the corresponding notification destination. Only the message name is sent as part of the notification. Supplied by the client.
	NotificationConfigs []Hl7V2NotificationConfigResponse `pulumi:"notificationConfigs"`
	// The configuration for the parser. It determines how the server parses the messages.
	ParserConfig ParserConfigResponse `pulumi:"parserConfig"`
	// Determines whether to reject duplicate messages. A duplicate message is a message with the same raw bytes as a message that has already been ingested/created in this HL7v2 store. The default value is false, meaning that the store accepts the duplicate messages and it also returns the same ACK message in the IngestMessageResponse as has been returned previously. Note that only one resource is created in the store. When this field is set to true, CreateMessage/IngestMessage requests with a duplicate message will be rejected by the store, and IngestMessageErrorDetail returns a NACK message upon rejection.
	RejectDuplicateMessage bool `pulumi:"rejectDuplicateMessage"`
}

func LookupHl7V2StoreOutput(ctx *pulumi.Context, args LookupHl7V2StoreOutputArgs, opts ...pulumi.InvokeOption) LookupHl7V2StoreResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupHl7V2StoreResult, error) {
			args := v.(LookupHl7V2StoreArgs)
			r, err := LookupHl7V2Store(ctx, &args, opts...)
			return *r, err
		}).(LookupHl7V2StoreResultOutput)
}

type LookupHl7V2StoreOutputArgs struct {
	DatasetId    pulumi.StringInput    `pulumi:"datasetId"`
	Hl7V2StoreId pulumi.StringInput    `pulumi:"hl7V2StoreId"`
	Location     pulumi.StringInput    `pulumi:"location"`
	Project      pulumi.StringPtrInput `pulumi:"project"`
}

func (LookupHl7V2StoreOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupHl7V2StoreArgs)(nil)).Elem()
}

type LookupHl7V2StoreResultOutput struct{ *pulumi.OutputState }

func (LookupHl7V2StoreResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupHl7V2StoreResult)(nil)).Elem()
}

func (o LookupHl7V2StoreResultOutput) ToLookupHl7V2StoreResultOutput() LookupHl7V2StoreResultOutput {
	return o
}

func (o LookupHl7V2StoreResultOutput) ToLookupHl7V2StoreResultOutputWithContext(ctx context.Context) LookupHl7V2StoreResultOutput {
	return o
}

// User-supplied key-value pairs used to organize HL7v2 stores. Label keys must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must conform to the following PCRE regular expression: \p{Ll}\p{Lo}{0,62} Label values are optional, must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must conform to the following PCRE regular expression: [\p{Ll}\p{Lo}\p{N}_-]{0,63} No more than 64 labels can be associated with a given store.
func (o LookupHl7V2StoreResultOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupHl7V2StoreResult) map[string]string { return v.Labels }).(pulumi.StringMapOutput)
}

// Resource name of the HL7v2 store, of the form `projects/{project_id}/datasets/{dataset_id}/hl7V2Stores/{hl7v2_store_id}`.
func (o LookupHl7V2StoreResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHl7V2StoreResult) string { return v.Name }).(pulumi.StringOutput)
}

// A list of notification configs. Each configuration uses a filter to determine whether to publish a message (both Ingest & Create) on the corresponding notification destination. Only the message name is sent as part of the notification. Supplied by the client.
func (o LookupHl7V2StoreResultOutput) NotificationConfigs() Hl7V2NotificationConfigResponseArrayOutput {
	return o.ApplyT(func(v LookupHl7V2StoreResult) []Hl7V2NotificationConfigResponse { return v.NotificationConfigs }).(Hl7V2NotificationConfigResponseArrayOutput)
}

// The configuration for the parser. It determines how the server parses the messages.
func (o LookupHl7V2StoreResultOutput) ParserConfig() ParserConfigResponseOutput {
	return o.ApplyT(func(v LookupHl7V2StoreResult) ParserConfigResponse { return v.ParserConfig }).(ParserConfigResponseOutput)
}

// Determines whether to reject duplicate messages. A duplicate message is a message with the same raw bytes as a message that has already been ingested/created in this HL7v2 store. The default value is false, meaning that the store accepts the duplicate messages and it also returns the same ACK message in the IngestMessageResponse as has been returned previously. Note that only one resource is created in the store. When this field is set to true, CreateMessage/IngestMessage requests with a duplicate message will be rejected by the store, and IngestMessageErrorDetail returns a NACK message upon rejection.
func (o LookupHl7V2StoreResultOutput) RejectDuplicateMessage() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupHl7V2StoreResult) bool { return v.RejectDuplicateMessage }).(pulumi.BoolOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupHl7V2StoreResultOutput{})
}
