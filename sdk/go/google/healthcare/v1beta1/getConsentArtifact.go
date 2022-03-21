// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the specified Consent artifact.
func LookupConsentArtifact(ctx *pulumi.Context, args *LookupConsentArtifactArgs, opts ...pulumi.InvokeOption) (*LookupConsentArtifactResult, error) {
	var rv LookupConsentArtifactResult
	err := ctx.Invoke("google-native:healthcare/v1beta1:getConsentArtifact", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupConsentArtifactArgs struct {
	ConsentArtifactId string  `pulumi:"consentArtifactId"`
	ConsentStoreId    string  `pulumi:"consentStoreId"`
	DatasetId         string  `pulumi:"datasetId"`
	Location          string  `pulumi:"location"`
	Project           *string `pulumi:"project"`
}

type LookupConsentArtifactResult struct {
	// Optional. Screenshots, PDFs, or other binary information documenting the user's consent.
	ConsentContentScreenshots []ImageResponse `pulumi:"consentContentScreenshots"`
	// Optional. An string indicating the version of the consent information shown to the user.
	ConsentContentVersion string `pulumi:"consentContentVersion"`
	// Optional. A signature from a guardian.
	GuardianSignature SignatureResponse `pulumi:"guardianSignature"`
	// Optional. Metadata associated with the Consent artifact. For example, the consent locale or user agent version.
	Metadata map[string]string `pulumi:"metadata"`
	// Resource name of the Consent artifact, of the form `projects/{project_id}/locations/{location_id}/datasets/{dataset_id}/consentStores/{consent_store_id}/consentArtifacts/{consent_artifact_id}`. Cannot be changed after creation.
	Name string `pulumi:"name"`
	// User's UUID provided by the client.
	UserId string `pulumi:"userId"`
	// Optional. User's signature.
	UserSignature SignatureResponse `pulumi:"userSignature"`
	// Optional. A signature from a witness.
	WitnessSignature SignatureResponse `pulumi:"witnessSignature"`
}

func LookupConsentArtifactOutput(ctx *pulumi.Context, args LookupConsentArtifactOutputArgs, opts ...pulumi.InvokeOption) LookupConsentArtifactResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupConsentArtifactResult, error) {
			args := v.(LookupConsentArtifactArgs)
			r, err := LookupConsentArtifact(ctx, &args, opts...)
			return *r, err
		}).(LookupConsentArtifactResultOutput)
}

type LookupConsentArtifactOutputArgs struct {
	ConsentArtifactId pulumi.StringInput    `pulumi:"consentArtifactId"`
	ConsentStoreId    pulumi.StringInput    `pulumi:"consentStoreId"`
	DatasetId         pulumi.StringInput    `pulumi:"datasetId"`
	Location          pulumi.StringInput    `pulumi:"location"`
	Project           pulumi.StringPtrInput `pulumi:"project"`
}

func (LookupConsentArtifactOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConsentArtifactArgs)(nil)).Elem()
}

type LookupConsentArtifactResultOutput struct{ *pulumi.OutputState }

func (LookupConsentArtifactResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConsentArtifactResult)(nil)).Elem()
}

func (o LookupConsentArtifactResultOutput) ToLookupConsentArtifactResultOutput() LookupConsentArtifactResultOutput {
	return o
}

func (o LookupConsentArtifactResultOutput) ToLookupConsentArtifactResultOutputWithContext(ctx context.Context) LookupConsentArtifactResultOutput {
	return o
}

// Optional. Screenshots, PDFs, or other binary information documenting the user's consent.
func (o LookupConsentArtifactResultOutput) ConsentContentScreenshots() ImageResponseArrayOutput {
	return o.ApplyT(func(v LookupConsentArtifactResult) []ImageResponse { return v.ConsentContentScreenshots }).(ImageResponseArrayOutput)
}

// Optional. An string indicating the version of the consent information shown to the user.
func (o LookupConsentArtifactResultOutput) ConsentContentVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConsentArtifactResult) string { return v.ConsentContentVersion }).(pulumi.StringOutput)
}

// Optional. A signature from a guardian.
func (o LookupConsentArtifactResultOutput) GuardianSignature() SignatureResponseOutput {
	return o.ApplyT(func(v LookupConsentArtifactResult) SignatureResponse { return v.GuardianSignature }).(SignatureResponseOutput)
}

// Optional. Metadata associated with the Consent artifact. For example, the consent locale or user agent version.
func (o LookupConsentArtifactResultOutput) Metadata() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupConsentArtifactResult) map[string]string { return v.Metadata }).(pulumi.StringMapOutput)
}

// Resource name of the Consent artifact, of the form `projects/{project_id}/locations/{location_id}/datasets/{dataset_id}/consentStores/{consent_store_id}/consentArtifacts/{consent_artifact_id}`. Cannot be changed after creation.
func (o LookupConsentArtifactResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConsentArtifactResult) string { return v.Name }).(pulumi.StringOutput)
}

// User's UUID provided by the client.
func (o LookupConsentArtifactResultOutput) UserId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConsentArtifactResult) string { return v.UserId }).(pulumi.StringOutput)
}

// Optional. User's signature.
func (o LookupConsentArtifactResultOutput) UserSignature() SignatureResponseOutput {
	return o.ApplyT(func(v LookupConsentArtifactResult) SignatureResponse { return v.UserSignature }).(SignatureResponseOutput)
}

// Optional. A signature from a witness.
func (o LookupConsentArtifactResultOutput) WitnessSignature() SignatureResponseOutput {
	return o.ApplyT(func(v LookupConsentArtifactResult) SignatureResponse { return v.WitnessSignature }).(SignatureResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupConsentArtifactResultOutput{})
}
