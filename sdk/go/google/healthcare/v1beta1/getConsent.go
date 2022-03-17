// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the specified revision of a Consent, or the latest revision if `revision_id` is not specified in the resource name.
func LookupConsent(ctx *pulumi.Context, args *LookupConsentArgs, opts ...pulumi.InvokeOption) (*LookupConsentResult, error) {
	var rv LookupConsentResult
	err := ctx.Invoke("google-native:healthcare/v1beta1:getConsent", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupConsentArgs struct {
	ConsentId      string  `pulumi:"consentId"`
	ConsentStoreId string  `pulumi:"consentStoreId"`
	DatasetId      string  `pulumi:"datasetId"`
	Location       string  `pulumi:"location"`
	Project        *string `pulumi:"project"`
}

type LookupConsentResult struct {
	// The resource name of the Consent artifact that contains proof of the end user's consent, of the form `projects/{project_id}/locations/{location_id}/datasets/{dataset_id}/consentStores/{consent_store_id}/consentArtifacts/{consent_artifact_id}`.
	ConsentArtifact string `pulumi:"consentArtifact"`
	// Timestamp in UTC of when this Consent is considered expired.
	ExpireTime string `pulumi:"expireTime"`
	// Optional. User-supplied key-value pairs used to organize Consent resources. Metadata keys must: - be between 1 and 63 characters long - have a UTF-8 encoding of maximum 128 bytes - begin with a letter - consist of up to 63 characters including lowercase letters, numeric characters, underscores, and dashes Metadata values must be: - be between 1 and 63 characters long - have a UTF-8 encoding of maximum 128 bytes - consist of up to 63 characters including lowercase letters, numeric characters, underscores, and dashes No more than 64 metadata entries can be associated with a given consent.
	Metadata map[string]string `pulumi:"metadata"`
	// Resource name of the Consent, of the form `projects/{project_id}/locations/{location_id}/datasets/{dataset_id}/consentStores/{consent_store_id}/consents/{consent_id}`. Cannot be changed after creation.
	Name string `pulumi:"name"`
	// Optional. Represents a user's consent in terms of the resources that can be accessed and under what conditions.
	Policies []GoogleCloudHealthcareV1beta1ConsentPolicyResponse `pulumi:"policies"`
	// The timestamp that the revision was created.
	RevisionCreateTime string `pulumi:"revisionCreateTime"`
	// The revision ID of the Consent. The format is an 8-character hexadecimal string. Refer to a specific revision of a Consent by appending `@{revision_id}` to the Consent's resource name.
	RevisionId string `pulumi:"revisionId"`
	// Indicates the current state of this Consent.
	State string `pulumi:"state"`
	// Input only. The time to live for this Consent from when it is created.
	Ttl string `pulumi:"ttl"`
	// User's UUID provided by the client.
	UserId string `pulumi:"userId"`
}

func LookupConsentOutput(ctx *pulumi.Context, args LookupConsentOutputArgs, opts ...pulumi.InvokeOption) LookupConsentResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupConsentResult, error) {
			args := v.(LookupConsentArgs)
			r, err := LookupConsent(ctx, &args, opts...)
			return *r, err
		}).(LookupConsentResultOutput)
}

type LookupConsentOutputArgs struct {
	ConsentId      pulumi.StringInput    `pulumi:"consentId"`
	ConsentStoreId pulumi.StringInput    `pulumi:"consentStoreId"`
	DatasetId      pulumi.StringInput    `pulumi:"datasetId"`
	Location       pulumi.StringInput    `pulumi:"location"`
	Project        pulumi.StringPtrInput `pulumi:"project"`
}

func (LookupConsentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConsentArgs)(nil)).Elem()
}

type LookupConsentResultOutput struct{ *pulumi.OutputState }

func (LookupConsentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConsentResult)(nil)).Elem()
}

func (o LookupConsentResultOutput) ToLookupConsentResultOutput() LookupConsentResultOutput {
	return o
}

func (o LookupConsentResultOutput) ToLookupConsentResultOutputWithContext(ctx context.Context) LookupConsentResultOutput {
	return o
}

// The resource name of the Consent artifact that contains proof of the end user's consent, of the form `projects/{project_id}/locations/{location_id}/datasets/{dataset_id}/consentStores/{consent_store_id}/consentArtifacts/{consent_artifact_id}`.
func (o LookupConsentResultOutput) ConsentArtifact() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConsentResult) string { return v.ConsentArtifact }).(pulumi.StringOutput)
}

// Timestamp in UTC of when this Consent is considered expired.
func (o LookupConsentResultOutput) ExpireTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConsentResult) string { return v.ExpireTime }).(pulumi.StringOutput)
}

// Optional. User-supplied key-value pairs used to organize Consent resources. Metadata keys must: - be between 1 and 63 characters long - have a UTF-8 encoding of maximum 128 bytes - begin with a letter - consist of up to 63 characters including lowercase letters, numeric characters, underscores, and dashes Metadata values must be: - be between 1 and 63 characters long - have a UTF-8 encoding of maximum 128 bytes - consist of up to 63 characters including lowercase letters, numeric characters, underscores, and dashes No more than 64 metadata entries can be associated with a given consent.
func (o LookupConsentResultOutput) Metadata() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupConsentResult) map[string]string { return v.Metadata }).(pulumi.StringMapOutput)
}

// Resource name of the Consent, of the form `projects/{project_id}/locations/{location_id}/datasets/{dataset_id}/consentStores/{consent_store_id}/consents/{consent_id}`. Cannot be changed after creation.
func (o LookupConsentResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConsentResult) string { return v.Name }).(pulumi.StringOutput)
}

// Optional. Represents a user's consent in terms of the resources that can be accessed and under what conditions.
func (o LookupConsentResultOutput) Policies() GoogleCloudHealthcareV1beta1ConsentPolicyResponseArrayOutput {
	return o.ApplyT(func(v LookupConsentResult) []GoogleCloudHealthcareV1beta1ConsentPolicyResponse { return v.Policies }).(GoogleCloudHealthcareV1beta1ConsentPolicyResponseArrayOutput)
}

// The timestamp that the revision was created.
func (o LookupConsentResultOutput) RevisionCreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConsentResult) string { return v.RevisionCreateTime }).(pulumi.StringOutput)
}

// The revision ID of the Consent. The format is an 8-character hexadecimal string. Refer to a specific revision of a Consent by appending `@{revision_id}` to the Consent's resource name.
func (o LookupConsentResultOutput) RevisionId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConsentResult) string { return v.RevisionId }).(pulumi.StringOutput)
}

// Indicates the current state of this Consent.
func (o LookupConsentResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConsentResult) string { return v.State }).(pulumi.StringOutput)
}

// Input only. The time to live for this Consent from when it is created.
func (o LookupConsentResultOutput) Ttl() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConsentResult) string { return v.Ttl }).(pulumi.StringOutput)
}

// User's UUID provided by the client.
func (o LookupConsentResultOutput) UserId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConsentResult) string { return v.UserId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupConsentResultOutput{})
}
