// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Returns a Certificate.
func LookupCertificate(ctx *pulumi.Context, args *LookupCertificateArgs, opts ...pulumi.InvokeOption) (*LookupCertificateResult, error) {
	var rv LookupCertificateResult
	err := ctx.Invoke("google-native:privateca/v1:getCertificate", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupCertificateArgs struct {
	CaPoolId      string  `pulumi:"caPoolId"`
	CertificateId string  `pulumi:"certificateId"`
	Location      string  `pulumi:"location"`
	Project       *string `pulumi:"project"`
}

type LookupCertificateResult struct {
	// A structured description of the issued X.509 certificate.
	CertificateDescription CertificateDescriptionResponse `pulumi:"certificateDescription"`
	// Immutable. The resource name for a CertificateTemplate used to issue this certificate, in the format `projects/*/locations/*/certificateTemplates/*`. If this is specified, the caller must have the necessary permission to use this template. If this is omitted, no template will be used. This template must be in the same location as the Certificate.
	CertificateTemplate string `pulumi:"certificateTemplate"`
	// Immutable. A description of the certificate and key that does not require X.509 or ASN.1.
	Config CertificateConfigResponse `pulumi:"config"`
	// The time at which this Certificate was created.
	CreateTime string `pulumi:"createTime"`
	// The resource name of the issuing CertificateAuthority in the format `projects/*/locations/*/caPools/*/certificateAuthorities/*`.
	IssuerCertificateAuthority string `pulumi:"issuerCertificateAuthority"`
	// Optional. Labels with user-defined metadata.
	Labels map[string]string `pulumi:"labels"`
	// Immutable. The desired lifetime of a certificate. Used to create the "not_before_time" and "not_after_time" fields inside an X.509 certificate. Note that the lifetime may be truncated if it would extend past the life of any certificate authority in the issuing chain.
	Lifetime string `pulumi:"lifetime"`
	// The resource name for this Certificate in the format `projects/*/locations/*/caPools/*/certificates/*`.
	Name string `pulumi:"name"`
	// The pem-encoded, signed X.509 certificate.
	PemCertificate string `pulumi:"pemCertificate"`
	// The chain that may be used to verify the X.509 certificate. Expected to be in issuer-to-root order according to RFC 5246.
	PemCertificateChain []string `pulumi:"pemCertificateChain"`
	// Immutable. A pem-encoded X.509 certificate signing request (CSR).
	PemCsr string `pulumi:"pemCsr"`
	// Details regarding the revocation of this Certificate. This Certificate is considered revoked if and only if this field is present.
	RevocationDetails RevocationDetailsResponse `pulumi:"revocationDetails"`
	// Immutable. Specifies how the Certificate's identity fields are to be decided. If this is omitted, the `DEFAULT` subject mode will be used.
	SubjectMode string `pulumi:"subjectMode"`
	// The time at which this Certificate was updated.
	UpdateTime string `pulumi:"updateTime"`
}

func LookupCertificateOutput(ctx *pulumi.Context, args LookupCertificateOutputArgs, opts ...pulumi.InvokeOption) LookupCertificateResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupCertificateResult, error) {
			args := v.(LookupCertificateArgs)
			r, err := LookupCertificate(ctx, &args, opts...)
			return *r, err
		}).(LookupCertificateResultOutput)
}

type LookupCertificateOutputArgs struct {
	CaPoolId      pulumi.StringInput    `pulumi:"caPoolId"`
	CertificateId pulumi.StringInput    `pulumi:"certificateId"`
	Location      pulumi.StringInput    `pulumi:"location"`
	Project       pulumi.StringPtrInput `pulumi:"project"`
}

func (LookupCertificateOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCertificateArgs)(nil)).Elem()
}

type LookupCertificateResultOutput struct{ *pulumi.OutputState }

func (LookupCertificateResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCertificateResult)(nil)).Elem()
}

func (o LookupCertificateResultOutput) ToLookupCertificateResultOutput() LookupCertificateResultOutput {
	return o
}

func (o LookupCertificateResultOutput) ToLookupCertificateResultOutputWithContext(ctx context.Context) LookupCertificateResultOutput {
	return o
}

// A structured description of the issued X.509 certificate.
func (o LookupCertificateResultOutput) CertificateDescription() CertificateDescriptionResponseOutput {
	return o.ApplyT(func(v LookupCertificateResult) CertificateDescriptionResponse { return v.CertificateDescription }).(CertificateDescriptionResponseOutput)
}

// Immutable. The resource name for a CertificateTemplate used to issue this certificate, in the format `projects/*/locations/*/certificateTemplates/*`. If this is specified, the caller must have the necessary permission to use this template. If this is omitted, no template will be used. This template must be in the same location as the Certificate.
func (o LookupCertificateResultOutput) CertificateTemplate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCertificateResult) string { return v.CertificateTemplate }).(pulumi.StringOutput)
}

// Immutable. A description of the certificate and key that does not require X.509 or ASN.1.
func (o LookupCertificateResultOutput) Config() CertificateConfigResponseOutput {
	return o.ApplyT(func(v LookupCertificateResult) CertificateConfigResponse { return v.Config }).(CertificateConfigResponseOutput)
}

// The time at which this Certificate was created.
func (o LookupCertificateResultOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCertificateResult) string { return v.CreateTime }).(pulumi.StringOutput)
}

// The resource name of the issuing CertificateAuthority in the format `projects/*/locations/*/caPools/*/certificateAuthorities/*`.
func (o LookupCertificateResultOutput) IssuerCertificateAuthority() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCertificateResult) string { return v.IssuerCertificateAuthority }).(pulumi.StringOutput)
}

// Optional. Labels with user-defined metadata.
func (o LookupCertificateResultOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupCertificateResult) map[string]string { return v.Labels }).(pulumi.StringMapOutput)
}

// Immutable. The desired lifetime of a certificate. Used to create the "not_before_time" and "not_after_time" fields inside an X.509 certificate. Note that the lifetime may be truncated if it would extend past the life of any certificate authority in the issuing chain.
func (o LookupCertificateResultOutput) Lifetime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCertificateResult) string { return v.Lifetime }).(pulumi.StringOutput)
}

// The resource name for this Certificate in the format `projects/*/locations/*/caPools/*/certificates/*`.
func (o LookupCertificateResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCertificateResult) string { return v.Name }).(pulumi.StringOutput)
}

// The pem-encoded, signed X.509 certificate.
func (o LookupCertificateResultOutput) PemCertificate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCertificateResult) string { return v.PemCertificate }).(pulumi.StringOutput)
}

// The chain that may be used to verify the X.509 certificate. Expected to be in issuer-to-root order according to RFC 5246.
func (o LookupCertificateResultOutput) PemCertificateChain() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupCertificateResult) []string { return v.PemCertificateChain }).(pulumi.StringArrayOutput)
}

// Immutable. A pem-encoded X.509 certificate signing request (CSR).
func (o LookupCertificateResultOutput) PemCsr() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCertificateResult) string { return v.PemCsr }).(pulumi.StringOutput)
}

// Details regarding the revocation of this Certificate. This Certificate is considered revoked if and only if this field is present.
func (o LookupCertificateResultOutput) RevocationDetails() RevocationDetailsResponseOutput {
	return o.ApplyT(func(v LookupCertificateResult) RevocationDetailsResponse { return v.RevocationDetails }).(RevocationDetailsResponseOutput)
}

// Immutable. Specifies how the Certificate's identity fields are to be decided. If this is omitted, the `DEFAULT` subject mode will be used.
func (o LookupCertificateResultOutput) SubjectMode() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCertificateResult) string { return v.SubjectMode }).(pulumi.StringOutput)
}

// The time at which this Certificate was updated.
func (o LookupCertificateResultOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCertificateResult) string { return v.UpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCertificateResultOutput{})
}
