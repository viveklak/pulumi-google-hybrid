// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package alpha

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a SslCertificate resource in the specified project and region using the data included in the request
type RegionSslCertificate struct {
	pulumi.CustomResourceState

	// A value read into memory from a certificate file. The certificate file must be in PEM format. The certificate chain must be no greater than 5 certs long. The chain must include at least one intermediate cert.
	Certificate pulumi.StringOutput `pulumi:"certificate"`
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringOutput `pulumi:"creationTimestamp"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description pulumi.StringOutput `pulumi:"description"`
	// Expire time of the certificate. RFC3339
	ExpireTime pulumi.StringOutput `pulumi:"expireTime"`
	// Type of the resource. Always compute#sslCertificate for SSL certificates.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Configuration and status of a managed SSL certificate.
	Managed SslCertificateManagedSslCertificateResponseOutput `pulumi:"managed"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name pulumi.StringOutput `pulumi:"name"`
	// A value read into memory from a write-only private key file. The private key file must be in PEM format. For security, only insert requests include this field.
	PrivateKey pulumi.StringOutput `pulumi:"privateKey"`
	// URL of the region where the regional SSL Certificate resides. This field is not applicable to global SSL Certificate.
	Region pulumi.StringOutput `pulumi:"region"`
	// [Output only] Server-defined URL for the resource.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
	// Server-defined URL for this resource with the resource id.
	SelfLinkWithId pulumi.StringOutput `pulumi:"selfLinkWithId"`
	// Configuration and status of a self-managed SSL certificate.
	SelfManaged SslCertificateSelfManagedSslCertificateResponseOutput `pulumi:"selfManaged"`
	// Domains associated with the certificate via Subject Alternative Name.
	SubjectAlternativeNames pulumi.StringArrayOutput `pulumi:"subjectAlternativeNames"`
	// (Optional) Specifies the type of SSL certificate, either "SELF_MANAGED" or "MANAGED". If not specified, the certificate is self-managed and the fields certificate and private_key are used.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewRegionSslCertificate registers a new resource with the given unique name, arguments, and options.
func NewRegionSslCertificate(ctx *pulumi.Context,
	name string, args *RegionSslCertificateArgs, opts ...pulumi.ResourceOption) (*RegionSslCertificate, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	if args.Region == nil {
		return nil, errors.New("invalid value for required argument 'Region'")
	}
	var resource RegionSslCertificate
	err := ctx.RegisterResource("google-native:compute/alpha:RegionSslCertificate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRegionSslCertificate gets an existing RegionSslCertificate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRegionSslCertificate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RegionSslCertificateState, opts ...pulumi.ResourceOption) (*RegionSslCertificate, error) {
	var resource RegionSslCertificate
	err := ctx.ReadResource("google-native:compute/alpha:RegionSslCertificate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RegionSslCertificate resources.
type regionSslCertificateState struct {
	// A value read into memory from a certificate file. The certificate file must be in PEM format. The certificate chain must be no greater than 5 certs long. The chain must include at least one intermediate cert.
	Certificate *string `pulumi:"certificate"`
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp *string `pulumi:"creationTimestamp"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description *string `pulumi:"description"`
	// Expire time of the certificate. RFC3339
	ExpireTime *string `pulumi:"expireTime"`
	// Type of the resource. Always compute#sslCertificate for SSL certificates.
	Kind *string `pulumi:"kind"`
	// Configuration and status of a managed SSL certificate.
	Managed *SslCertificateManagedSslCertificateResponse `pulumi:"managed"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name *string `pulumi:"name"`
	// A value read into memory from a write-only private key file. The private key file must be in PEM format. For security, only insert requests include this field.
	PrivateKey *string `pulumi:"privateKey"`
	// URL of the region where the regional SSL Certificate resides. This field is not applicable to global SSL Certificate.
	Region *string `pulumi:"region"`
	// [Output only] Server-defined URL for the resource.
	SelfLink *string `pulumi:"selfLink"`
	// Server-defined URL for this resource with the resource id.
	SelfLinkWithId *string `pulumi:"selfLinkWithId"`
	// Configuration and status of a self-managed SSL certificate.
	SelfManaged *SslCertificateSelfManagedSslCertificateResponse `pulumi:"selfManaged"`
	// Domains associated with the certificate via Subject Alternative Name.
	SubjectAlternativeNames []string `pulumi:"subjectAlternativeNames"`
	// (Optional) Specifies the type of SSL certificate, either "SELF_MANAGED" or "MANAGED". If not specified, the certificate is self-managed and the fields certificate and private_key are used.
	Type *string `pulumi:"type"`
}

type RegionSslCertificateState struct {
	// A value read into memory from a certificate file. The certificate file must be in PEM format. The certificate chain must be no greater than 5 certs long. The chain must include at least one intermediate cert.
	Certificate pulumi.StringPtrInput
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringPtrInput
	// An optional description of this resource. Provide this property when you create the resource.
	Description pulumi.StringPtrInput
	// Expire time of the certificate. RFC3339
	ExpireTime pulumi.StringPtrInput
	// Type of the resource. Always compute#sslCertificate for SSL certificates.
	Kind pulumi.StringPtrInput
	// Configuration and status of a managed SSL certificate.
	Managed SslCertificateManagedSslCertificateResponsePtrInput
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name pulumi.StringPtrInput
	// A value read into memory from a write-only private key file. The private key file must be in PEM format. For security, only insert requests include this field.
	PrivateKey pulumi.StringPtrInput
	// URL of the region where the regional SSL Certificate resides. This field is not applicable to global SSL Certificate.
	Region pulumi.StringPtrInput
	// [Output only] Server-defined URL for the resource.
	SelfLink pulumi.StringPtrInput
	// Server-defined URL for this resource with the resource id.
	SelfLinkWithId pulumi.StringPtrInput
	// Configuration and status of a self-managed SSL certificate.
	SelfManaged SslCertificateSelfManagedSslCertificateResponsePtrInput
	// Domains associated with the certificate via Subject Alternative Name.
	SubjectAlternativeNames pulumi.StringArrayInput
	// (Optional) Specifies the type of SSL certificate, either "SELF_MANAGED" or "MANAGED". If not specified, the certificate is self-managed and the fields certificate and private_key are used.
	Type pulumi.StringPtrInput
}

func (RegionSslCertificateState) ElementType() reflect.Type {
	return reflect.TypeOf((*regionSslCertificateState)(nil)).Elem()
}

type regionSslCertificateArgs struct {
	// A value read into memory from a certificate file. The certificate file must be in PEM format. The certificate chain must be no greater than 5 certs long. The chain must include at least one intermediate cert.
	Certificate *string `pulumi:"certificate"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description *string `pulumi:"description"`
	// Configuration and status of a managed SSL certificate.
	Managed *SslCertificateManagedSslCertificate `pulumi:"managed"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name *string `pulumi:"name"`
	// A value read into memory from a write-only private key file. The private key file must be in PEM format. For security, only insert requests include this field.
	PrivateKey *string `pulumi:"privateKey"`
	Project    string  `pulumi:"project"`
	Region     string  `pulumi:"region"`
	RequestId  *string `pulumi:"requestId"`
	// Configuration and status of a self-managed SSL certificate.
	SelfManaged *SslCertificateSelfManagedSslCertificate `pulumi:"selfManaged"`
	// (Optional) Specifies the type of SSL certificate, either "SELF_MANAGED" or "MANAGED". If not specified, the certificate is self-managed and the fields certificate and private_key are used.
	Type *string `pulumi:"type"`
}

// The set of arguments for constructing a RegionSslCertificate resource.
type RegionSslCertificateArgs struct {
	// A value read into memory from a certificate file. The certificate file must be in PEM format. The certificate chain must be no greater than 5 certs long. The chain must include at least one intermediate cert.
	Certificate pulumi.StringPtrInput
	// An optional description of this resource. Provide this property when you create the resource.
	Description pulumi.StringPtrInput
	// Configuration and status of a managed SSL certificate.
	Managed SslCertificateManagedSslCertificatePtrInput
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name pulumi.StringPtrInput
	// A value read into memory from a write-only private key file. The private key file must be in PEM format. For security, only insert requests include this field.
	PrivateKey pulumi.StringPtrInput
	Project    pulumi.StringInput
	Region     pulumi.StringInput
	RequestId  pulumi.StringPtrInput
	// Configuration and status of a self-managed SSL certificate.
	SelfManaged SslCertificateSelfManagedSslCertificatePtrInput
	// (Optional) Specifies the type of SSL certificate, either "SELF_MANAGED" or "MANAGED". If not specified, the certificate is self-managed and the fields certificate and private_key are used.
	Type *RegionSslCertificateType
}

func (RegionSslCertificateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*regionSslCertificateArgs)(nil)).Elem()
}

type RegionSslCertificateInput interface {
	pulumi.Input

	ToRegionSslCertificateOutput() RegionSslCertificateOutput
	ToRegionSslCertificateOutputWithContext(ctx context.Context) RegionSslCertificateOutput
}

func (*RegionSslCertificate) ElementType() reflect.Type {
	return reflect.TypeOf((*RegionSslCertificate)(nil))
}

func (i *RegionSslCertificate) ToRegionSslCertificateOutput() RegionSslCertificateOutput {
	return i.ToRegionSslCertificateOutputWithContext(context.Background())
}

func (i *RegionSslCertificate) ToRegionSslCertificateOutputWithContext(ctx context.Context) RegionSslCertificateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegionSslCertificateOutput)
}

type RegionSslCertificateOutput struct {
	*pulumi.OutputState
}

func (RegionSslCertificateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RegionSslCertificate)(nil))
}

func (o RegionSslCertificateOutput) ToRegionSslCertificateOutput() RegionSslCertificateOutput {
	return o
}

func (o RegionSslCertificateOutput) ToRegionSslCertificateOutputWithContext(ctx context.Context) RegionSslCertificateOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(RegionSslCertificateOutput{})
}
