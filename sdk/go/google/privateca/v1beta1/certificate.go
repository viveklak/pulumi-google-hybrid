// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Create a new Certificate in a given Project, Location from a particular CertificateAuthority.
type Certificate struct {
	pulumi.CustomResourceState

	// A structured description of the issued X.509 certificate.
	CertificateDescription CertificateDescriptionResponseOutput `pulumi:"certificateDescription"`
	// Immutable. A description of the certificate and key that does not require X.509 or ASN.1.
	Config CertificateConfigResponseOutput `pulumi:"config"`
	// The time at which this Certificate was created.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Optional. Labels with user-defined metadata.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// Immutable. The desired lifetime of a certificate. Used to create the "not_before_time" and "not_after_time" fields inside an X.509 certificate. Note that the lifetime may be truncated if it would extend past the life of any certificate authority in the issuing chain.
	Lifetime pulumi.StringOutput `pulumi:"lifetime"`
	// The resource path for this Certificate in the format `projects/*/locations/*/certificateAuthorities/*/certificates/*`.
	Name pulumi.StringOutput `pulumi:"name"`
	// The pem-encoded, signed X.509 certificate.
	PemCertificate pulumi.StringOutput `pulumi:"pemCertificate"`
	// The chain that may be used to verify the X.509 certificate. Expected to be in issuer-to-root order according to RFC 5246.
	PemCertificateChain pulumi.StringArrayOutput `pulumi:"pemCertificateChain"`
	// Immutable. A pem-encoded X.509 certificate signing request (CSR).
	PemCsr pulumi.StringOutput `pulumi:"pemCsr"`
	// Details regarding the revocation of this Certificate. This Certificate is considered revoked if and only if this field is present.
	RevocationDetails RevocationDetailsResponseOutput `pulumi:"revocationDetails"`
	// The time at which this Certificate was updated.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewCertificate registers a new resource with the given unique name, arguments, and options.
func NewCertificate(ctx *pulumi.Context,
	name string, args *CertificateArgs, opts ...pulumi.ResourceOption) (*Certificate, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CertificateAuthorityId == nil {
		return nil, errors.New("invalid value for required argument 'CertificateAuthorityId'")
	}
	if args.Lifetime == nil {
		return nil, errors.New("invalid value for required argument 'Lifetime'")
	}
	if args.Location == nil {
		return nil, errors.New("invalid value for required argument 'Location'")
	}
	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	var resource Certificate
	err := ctx.RegisterResource("google-native:privateca/v1beta1:Certificate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCertificate gets an existing Certificate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCertificate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CertificateState, opts ...pulumi.ResourceOption) (*Certificate, error) {
	var resource Certificate
	err := ctx.ReadResource("google-native:privateca/v1beta1:Certificate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Certificate resources.
type certificateState struct {
	// A structured description of the issued X.509 certificate.
	CertificateDescription *CertificateDescriptionResponse `pulumi:"certificateDescription"`
	// Immutable. A description of the certificate and key that does not require X.509 or ASN.1.
	Config *CertificateConfigResponse `pulumi:"config"`
	// The time at which this Certificate was created.
	CreateTime *string `pulumi:"createTime"`
	// Optional. Labels with user-defined metadata.
	Labels map[string]string `pulumi:"labels"`
	// Immutable. The desired lifetime of a certificate. Used to create the "not_before_time" and "not_after_time" fields inside an X.509 certificate. Note that the lifetime may be truncated if it would extend past the life of any certificate authority in the issuing chain.
	Lifetime *string `pulumi:"lifetime"`
	// The resource path for this Certificate in the format `projects/*/locations/*/certificateAuthorities/*/certificates/*`.
	Name *string `pulumi:"name"`
	// The pem-encoded, signed X.509 certificate.
	PemCertificate *string `pulumi:"pemCertificate"`
	// The chain that may be used to verify the X.509 certificate. Expected to be in issuer-to-root order according to RFC 5246.
	PemCertificateChain []string `pulumi:"pemCertificateChain"`
	// Immutable. A pem-encoded X.509 certificate signing request (CSR).
	PemCsr *string `pulumi:"pemCsr"`
	// Details regarding the revocation of this Certificate. This Certificate is considered revoked if and only if this field is present.
	RevocationDetails *RevocationDetailsResponse `pulumi:"revocationDetails"`
	// The time at which this Certificate was updated.
	UpdateTime *string `pulumi:"updateTime"`
}

type CertificateState struct {
	// A structured description of the issued X.509 certificate.
	CertificateDescription CertificateDescriptionResponsePtrInput
	// Immutable. A description of the certificate and key that does not require X.509 or ASN.1.
	Config CertificateConfigResponsePtrInput
	// The time at which this Certificate was created.
	CreateTime pulumi.StringPtrInput
	// Optional. Labels with user-defined metadata.
	Labels pulumi.StringMapInput
	// Immutable. The desired lifetime of a certificate. Used to create the "not_before_time" and "not_after_time" fields inside an X.509 certificate. Note that the lifetime may be truncated if it would extend past the life of any certificate authority in the issuing chain.
	Lifetime pulumi.StringPtrInput
	// The resource path for this Certificate in the format `projects/*/locations/*/certificateAuthorities/*/certificates/*`.
	Name pulumi.StringPtrInput
	// The pem-encoded, signed X.509 certificate.
	PemCertificate pulumi.StringPtrInput
	// The chain that may be used to verify the X.509 certificate. Expected to be in issuer-to-root order according to RFC 5246.
	PemCertificateChain pulumi.StringArrayInput
	// Immutable. A pem-encoded X.509 certificate signing request (CSR).
	PemCsr pulumi.StringPtrInput
	// Details regarding the revocation of this Certificate. This Certificate is considered revoked if and only if this field is present.
	RevocationDetails RevocationDetailsResponsePtrInput
	// The time at which this Certificate was updated.
	UpdateTime pulumi.StringPtrInput
}

func (CertificateState) ElementType() reflect.Type {
	return reflect.TypeOf((*certificateState)(nil)).Elem()
}

type certificateArgs struct {
	CertificateAuthorityId string  `pulumi:"certificateAuthorityId"`
	CertificateId          *string `pulumi:"certificateId"`
	// Immutable. A description of the certificate and key that does not require X.509 or ASN.1.
	Config *CertificateConfig `pulumi:"config"`
	// Optional. Labels with user-defined metadata.
	Labels map[string]string `pulumi:"labels"`
	// Immutable. The desired lifetime of a certificate. Used to create the "not_before_time" and "not_after_time" fields inside an X.509 certificate. Note that the lifetime may be truncated if it would extend past the life of any certificate authority in the issuing chain.
	Lifetime string `pulumi:"lifetime"`
	Location string `pulumi:"location"`
	// Immutable. A pem-encoded X.509 certificate signing request (CSR).
	PemCsr    *string `pulumi:"pemCsr"`
	Project   string  `pulumi:"project"`
	RequestId *string `pulumi:"requestId"`
}

// The set of arguments for constructing a Certificate resource.
type CertificateArgs struct {
	CertificateAuthorityId pulumi.StringInput
	CertificateId          pulumi.StringPtrInput
	// Immutable. A description of the certificate and key that does not require X.509 or ASN.1.
	Config CertificateConfigPtrInput
	// Optional. Labels with user-defined metadata.
	Labels pulumi.StringMapInput
	// Immutable. The desired lifetime of a certificate. Used to create the "not_before_time" and "not_after_time" fields inside an X.509 certificate. Note that the lifetime may be truncated if it would extend past the life of any certificate authority in the issuing chain.
	Lifetime pulumi.StringInput
	Location pulumi.StringInput
	// Immutable. A pem-encoded X.509 certificate signing request (CSR).
	PemCsr    pulumi.StringPtrInput
	Project   pulumi.StringInput
	RequestId pulumi.StringPtrInput
}

func (CertificateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*certificateArgs)(nil)).Elem()
}

type CertificateInput interface {
	pulumi.Input

	ToCertificateOutput() CertificateOutput
	ToCertificateOutputWithContext(ctx context.Context) CertificateOutput
}

func (*Certificate) ElementType() reflect.Type {
	return reflect.TypeOf((*Certificate)(nil))
}

func (i *Certificate) ToCertificateOutput() CertificateOutput {
	return i.ToCertificateOutputWithContext(context.Background())
}

func (i *Certificate) ToCertificateOutputWithContext(ctx context.Context) CertificateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificateOutput)
}

type CertificateOutput struct {
	*pulumi.OutputState
}

func (CertificateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Certificate)(nil))
}

func (o CertificateOutput) ToCertificateOutput() CertificateOutput {
	return o
}

func (o CertificateOutput) ToCertificateOutputWithContext(ctx context.Context) CertificateOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(CertificateOutput{})
}
