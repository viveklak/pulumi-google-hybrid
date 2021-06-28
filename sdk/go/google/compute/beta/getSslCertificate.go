// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package beta

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Returns the specified SslCertificate resource. Gets a list of available SSL certificates by making a list() request.
func LookupSslCertificate(ctx *pulumi.Context, args *LookupSslCertificateArgs, opts ...pulumi.InvokeOption) (*LookupSslCertificateResult, error) {
	var rv LookupSslCertificateResult
	err := ctx.Invoke("google-native:compute/beta:getSslCertificate", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSslCertificateArgs struct {
	Project        string `pulumi:"project"`
	SslCertificate string `pulumi:"sslCertificate"`
}

type LookupSslCertificateResult struct {
	// A value read into memory from a certificate file. The certificate file must be in PEM format. The certificate chain must be no greater than 5 certs long. The chain must include at least one intermediate cert.
	Certificate string `pulumi:"certificate"`
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp string `pulumi:"creationTimestamp"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description string `pulumi:"description"`
	// Expire time of the certificate. RFC3339
	ExpireTime string `pulumi:"expireTime"`
	// Type of the resource. Always compute#sslCertificate for SSL certificates.
	Kind string `pulumi:"kind"`
	// Configuration and status of a managed SSL certificate.
	Managed SslCertificateManagedSslCertificateResponse `pulumi:"managed"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name string `pulumi:"name"`
	// A value read into memory from a write-only private key file. The private key file must be in PEM format. For security, only insert requests include this field.
	PrivateKey string `pulumi:"privateKey"`
	// URL of the region where the regional SSL Certificate resides. This field is not applicable to global SSL Certificate.
	Region string `pulumi:"region"`
	// [Output only] Server-defined URL for the resource.
	SelfLink string `pulumi:"selfLink"`
	// Configuration and status of a self-managed SSL certificate.
	SelfManaged SslCertificateSelfManagedSslCertificateResponse `pulumi:"selfManaged"`
	// Domains associated with the certificate via Subject Alternative Name.
	SubjectAlternativeNames []string `pulumi:"subjectAlternativeNames"`
	// (Optional) Specifies the type of SSL certificate, either "SELF_MANAGED" or "MANAGED". If not specified, the certificate is self-managed and the fields certificate and private_key are used.
	Type string `pulumi:"type"`
}
