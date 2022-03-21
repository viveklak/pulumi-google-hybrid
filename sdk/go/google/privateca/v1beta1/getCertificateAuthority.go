// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Returns a CertificateAuthority.
func LookupCertificateAuthority(ctx *pulumi.Context, args *LookupCertificateAuthorityArgs, opts ...pulumi.InvokeOption) (*LookupCertificateAuthorityResult, error) {
	var rv LookupCertificateAuthorityResult
	err := ctx.Invoke("google-native:privateca/v1beta1:getCertificateAuthority", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupCertificateAuthorityArgs struct {
	CertificateAuthorityId string  `pulumi:"certificateAuthorityId"`
	Location               string  `pulumi:"location"`
	Project                *string `pulumi:"project"`
}

type LookupCertificateAuthorityResult struct {
	// URLs for accessing content published by this CA, such as the CA certificate and CRLs.
	AccessUrls AccessUrlsResponse `pulumi:"accessUrls"`
	// A structured description of this CertificateAuthority's CA certificate and its issuers. Ordered as self-to-root.
	CaCertificateDescriptions []CertificateDescriptionResponse `pulumi:"caCertificateDescriptions"`
	// Optional. The CertificateAuthorityPolicy to enforce when issuing Certificates from this CertificateAuthority.
	CertificatePolicy CertificateAuthorityPolicyResponse `pulumi:"certificatePolicy"`
	// Immutable. The config used to create a self-signed X.509 certificate or CSR.
	Config CertificateConfigResponse `pulumi:"config"`
	// The time at which this CertificateAuthority was created.
	CreateTime string `pulumi:"createTime"`
	// The time at which this CertificateAuthority will be deleted, if scheduled for deletion.
	DeleteTime string `pulumi:"deleteTime"`
	// Immutable. The name of a Cloud Storage bucket where this CertificateAuthority will publish content, such as the CA certificate and CRLs. This must be a bucket name, without any prefixes (such as `gs://`) or suffixes (such as `.googleapis.com`). For example, to use a bucket named `my-bucket`, you would simply specify `my-bucket`. If not specified, a managed bucket will be created.
	GcsBucket string `pulumi:"gcsBucket"`
	// Optional. The IssuingOptions to follow when issuing Certificates from this CertificateAuthority.
	IssuingOptions IssuingOptionsResponse `pulumi:"issuingOptions"`
	// Immutable. Used when issuing certificates for this CertificateAuthority. If this CertificateAuthority is a self-signed CertificateAuthority, this key is also used to sign the self-signed CA certificate. Otherwise, it is used to sign a CSR.
	KeySpec KeyVersionSpecResponse `pulumi:"keySpec"`
	// Optional. Labels with user-defined metadata.
	Labels map[string]string `pulumi:"labels"`
	// The desired lifetime of the CA certificate. Used to create the "not_before_time" and "not_after_time" fields inside an X.509 certificate.
	Lifetime string `pulumi:"lifetime"`
	// The resource name for this CertificateAuthority in the format `projects/*/locations/*/certificateAuthorities/*`.
	Name string `pulumi:"name"`
	// This CertificateAuthority's certificate chain, including the current CertificateAuthority's certificate. Ordered such that the root issuer is the final element (consistent with RFC 5246). For a self-signed CA, this will only list the current CertificateAuthority's certificate.
	PemCaCertificates []string `pulumi:"pemCaCertificates"`
	// The State for this CertificateAuthority.
	State string `pulumi:"state"`
	// Optional. If this is a subordinate CertificateAuthority, this field will be set with the subordinate configuration, which describes its issuers. This may be updated, but this CertificateAuthority must continue to validate.
	SubordinateConfig SubordinateConfigResponse `pulumi:"subordinateConfig"`
	// Immutable. The Tier of this CertificateAuthority.
	Tier string `pulumi:"tier"`
	// Immutable. The Type of this CertificateAuthority.
	Type string `pulumi:"type"`
	// The time at which this CertificateAuthority was updated.
	UpdateTime string `pulumi:"updateTime"`
}

func LookupCertificateAuthorityOutput(ctx *pulumi.Context, args LookupCertificateAuthorityOutputArgs, opts ...pulumi.InvokeOption) LookupCertificateAuthorityResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupCertificateAuthorityResult, error) {
			args := v.(LookupCertificateAuthorityArgs)
			r, err := LookupCertificateAuthority(ctx, &args, opts...)
			return *r, err
		}).(LookupCertificateAuthorityResultOutput)
}

type LookupCertificateAuthorityOutputArgs struct {
	CertificateAuthorityId pulumi.StringInput    `pulumi:"certificateAuthorityId"`
	Location               pulumi.StringInput    `pulumi:"location"`
	Project                pulumi.StringPtrInput `pulumi:"project"`
}

func (LookupCertificateAuthorityOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCertificateAuthorityArgs)(nil)).Elem()
}

type LookupCertificateAuthorityResultOutput struct{ *pulumi.OutputState }

func (LookupCertificateAuthorityResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCertificateAuthorityResult)(nil)).Elem()
}

func (o LookupCertificateAuthorityResultOutput) ToLookupCertificateAuthorityResultOutput() LookupCertificateAuthorityResultOutput {
	return o
}

func (o LookupCertificateAuthorityResultOutput) ToLookupCertificateAuthorityResultOutputWithContext(ctx context.Context) LookupCertificateAuthorityResultOutput {
	return o
}

// URLs for accessing content published by this CA, such as the CA certificate and CRLs.
func (o LookupCertificateAuthorityResultOutput) AccessUrls() AccessUrlsResponseOutput {
	return o.ApplyT(func(v LookupCertificateAuthorityResult) AccessUrlsResponse { return v.AccessUrls }).(AccessUrlsResponseOutput)
}

// A structured description of this CertificateAuthority's CA certificate and its issuers. Ordered as self-to-root.
func (o LookupCertificateAuthorityResultOutput) CaCertificateDescriptions() CertificateDescriptionResponseArrayOutput {
	return o.ApplyT(func(v LookupCertificateAuthorityResult) []CertificateDescriptionResponse {
		return v.CaCertificateDescriptions
	}).(CertificateDescriptionResponseArrayOutput)
}

// Optional. The CertificateAuthorityPolicy to enforce when issuing Certificates from this CertificateAuthority.
func (o LookupCertificateAuthorityResultOutput) CertificatePolicy() CertificateAuthorityPolicyResponseOutput {
	return o.ApplyT(func(v LookupCertificateAuthorityResult) CertificateAuthorityPolicyResponse {
		return v.CertificatePolicy
	}).(CertificateAuthorityPolicyResponseOutput)
}

// Immutable. The config used to create a self-signed X.509 certificate or CSR.
func (o LookupCertificateAuthorityResultOutput) Config() CertificateConfigResponseOutput {
	return o.ApplyT(func(v LookupCertificateAuthorityResult) CertificateConfigResponse { return v.Config }).(CertificateConfigResponseOutput)
}

// The time at which this CertificateAuthority was created.
func (o LookupCertificateAuthorityResultOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCertificateAuthorityResult) string { return v.CreateTime }).(pulumi.StringOutput)
}

// The time at which this CertificateAuthority will be deleted, if scheduled for deletion.
func (o LookupCertificateAuthorityResultOutput) DeleteTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCertificateAuthorityResult) string { return v.DeleteTime }).(pulumi.StringOutput)
}

// Immutable. The name of a Cloud Storage bucket where this CertificateAuthority will publish content, such as the CA certificate and CRLs. This must be a bucket name, without any prefixes (such as `gs://`) or suffixes (such as `.googleapis.com`). For example, to use a bucket named `my-bucket`, you would simply specify `my-bucket`. If not specified, a managed bucket will be created.
func (o LookupCertificateAuthorityResultOutput) GcsBucket() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCertificateAuthorityResult) string { return v.GcsBucket }).(pulumi.StringOutput)
}

// Optional. The IssuingOptions to follow when issuing Certificates from this CertificateAuthority.
func (o LookupCertificateAuthorityResultOutput) IssuingOptions() IssuingOptionsResponseOutput {
	return o.ApplyT(func(v LookupCertificateAuthorityResult) IssuingOptionsResponse { return v.IssuingOptions }).(IssuingOptionsResponseOutput)
}

// Immutable. Used when issuing certificates for this CertificateAuthority. If this CertificateAuthority is a self-signed CertificateAuthority, this key is also used to sign the self-signed CA certificate. Otherwise, it is used to sign a CSR.
func (o LookupCertificateAuthorityResultOutput) KeySpec() KeyVersionSpecResponseOutput {
	return o.ApplyT(func(v LookupCertificateAuthorityResult) KeyVersionSpecResponse { return v.KeySpec }).(KeyVersionSpecResponseOutput)
}

// Optional. Labels with user-defined metadata.
func (o LookupCertificateAuthorityResultOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupCertificateAuthorityResult) map[string]string { return v.Labels }).(pulumi.StringMapOutput)
}

// The desired lifetime of the CA certificate. Used to create the "not_before_time" and "not_after_time" fields inside an X.509 certificate.
func (o LookupCertificateAuthorityResultOutput) Lifetime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCertificateAuthorityResult) string { return v.Lifetime }).(pulumi.StringOutput)
}

// The resource name for this CertificateAuthority in the format `projects/*/locations/*/certificateAuthorities/*`.
func (o LookupCertificateAuthorityResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCertificateAuthorityResult) string { return v.Name }).(pulumi.StringOutput)
}

// This CertificateAuthority's certificate chain, including the current CertificateAuthority's certificate. Ordered such that the root issuer is the final element (consistent with RFC 5246). For a self-signed CA, this will only list the current CertificateAuthority's certificate.
func (o LookupCertificateAuthorityResultOutput) PemCaCertificates() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupCertificateAuthorityResult) []string { return v.PemCaCertificates }).(pulumi.StringArrayOutput)
}

// The State for this CertificateAuthority.
func (o LookupCertificateAuthorityResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCertificateAuthorityResult) string { return v.State }).(pulumi.StringOutput)
}

// Optional. If this is a subordinate CertificateAuthority, this field will be set with the subordinate configuration, which describes its issuers. This may be updated, but this CertificateAuthority must continue to validate.
func (o LookupCertificateAuthorityResultOutput) SubordinateConfig() SubordinateConfigResponseOutput {
	return o.ApplyT(func(v LookupCertificateAuthorityResult) SubordinateConfigResponse { return v.SubordinateConfig }).(SubordinateConfigResponseOutput)
}

// Immutable. The Tier of this CertificateAuthority.
func (o LookupCertificateAuthorityResultOutput) Tier() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCertificateAuthorityResult) string { return v.Tier }).(pulumi.StringOutput)
}

// Immutable. The Type of this CertificateAuthority.
func (o LookupCertificateAuthorityResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCertificateAuthorityResult) string { return v.Type }).(pulumi.StringOutput)
}

// The time at which this CertificateAuthority was updated.
func (o LookupCertificateAuthorityResultOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCertificateAuthorityResult) string { return v.UpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCertificateAuthorityResultOutput{})
}
