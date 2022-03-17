// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package beta

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Returns the specified TargetSslProxy resource. Gets a list of available target SSL proxies by making a list() request.
func LookupTargetSslProxy(ctx *pulumi.Context, args *LookupTargetSslProxyArgs, opts ...pulumi.InvokeOption) (*LookupTargetSslProxyResult, error) {
	var rv LookupTargetSslProxyResult
	err := ctx.Invoke("google-native:compute/beta:getTargetSslProxy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupTargetSslProxyArgs struct {
	Project        *string `pulumi:"project"`
	TargetSslProxy string  `pulumi:"targetSslProxy"`
}

type LookupTargetSslProxyResult struct {
	// URL of a certificate map that identifies a certificate map associated with the given target proxy. This field can only be set for global target proxies. If set, sslCertificates will be ignored.
	CertificateMap string `pulumi:"certificateMap"`
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp string `pulumi:"creationTimestamp"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description string `pulumi:"description"`
	// Type of the resource. Always compute#targetSslProxy for target SSL proxies.
	Kind string `pulumi:"kind"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name string `pulumi:"name"`
	// Specifies the type of proxy header to append before sending data to the backend, either NONE or PROXY_V1. The default is NONE.
	ProxyHeader string `pulumi:"proxyHeader"`
	// Server-defined URL for the resource.
	SelfLink string `pulumi:"selfLink"`
	// URL to the BackendService resource.
	Service string `pulumi:"service"`
	// URLs to SslCertificate resources that are used to authenticate connections to Backends. At least one SSL certificate must be specified. Currently, you may specify up to 15 SSL certificates. sslCertificates do not apply when the load balancing scheme is set to INTERNAL_SELF_MANAGED.
	SslCertificates []string `pulumi:"sslCertificates"`
	// URL of SslPolicy resource that will be associated with the TargetSslProxy resource. If not set, the TargetSslProxy resource will not have any SSL policy configured.
	SslPolicy string `pulumi:"sslPolicy"`
}

func LookupTargetSslProxyOutput(ctx *pulumi.Context, args LookupTargetSslProxyOutputArgs, opts ...pulumi.InvokeOption) LookupTargetSslProxyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupTargetSslProxyResult, error) {
			args := v.(LookupTargetSslProxyArgs)
			r, err := LookupTargetSslProxy(ctx, &args, opts...)
			return *r, err
		}).(LookupTargetSslProxyResultOutput)
}

type LookupTargetSslProxyOutputArgs struct {
	Project        pulumi.StringPtrInput `pulumi:"project"`
	TargetSslProxy pulumi.StringInput    `pulumi:"targetSslProxy"`
}

func (LookupTargetSslProxyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTargetSslProxyArgs)(nil)).Elem()
}

type LookupTargetSslProxyResultOutput struct{ *pulumi.OutputState }

func (LookupTargetSslProxyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTargetSslProxyResult)(nil)).Elem()
}

func (o LookupTargetSslProxyResultOutput) ToLookupTargetSslProxyResultOutput() LookupTargetSslProxyResultOutput {
	return o
}

func (o LookupTargetSslProxyResultOutput) ToLookupTargetSslProxyResultOutputWithContext(ctx context.Context) LookupTargetSslProxyResultOutput {
	return o
}

// URL of a certificate map that identifies a certificate map associated with the given target proxy. This field can only be set for global target proxies. If set, sslCertificates will be ignored.
func (o LookupTargetSslProxyResultOutput) CertificateMap() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTargetSslProxyResult) string { return v.CertificateMap }).(pulumi.StringOutput)
}

// Creation timestamp in RFC3339 text format.
func (o LookupTargetSslProxyResultOutput) CreationTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTargetSslProxyResult) string { return v.CreationTimestamp }).(pulumi.StringOutput)
}

// An optional description of this resource. Provide this property when you create the resource.
func (o LookupTargetSslProxyResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTargetSslProxyResult) string { return v.Description }).(pulumi.StringOutput)
}

// Type of the resource. Always compute#targetSslProxy for target SSL proxies.
func (o LookupTargetSslProxyResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTargetSslProxyResult) string { return v.Kind }).(pulumi.StringOutput)
}

// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
func (o LookupTargetSslProxyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTargetSslProxyResult) string { return v.Name }).(pulumi.StringOutput)
}

// Specifies the type of proxy header to append before sending data to the backend, either NONE or PROXY_V1. The default is NONE.
func (o LookupTargetSslProxyResultOutput) ProxyHeader() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTargetSslProxyResult) string { return v.ProxyHeader }).(pulumi.StringOutput)
}

// Server-defined URL for the resource.
func (o LookupTargetSslProxyResultOutput) SelfLink() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTargetSslProxyResult) string { return v.SelfLink }).(pulumi.StringOutput)
}

// URL to the BackendService resource.
func (o LookupTargetSslProxyResultOutput) Service() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTargetSslProxyResult) string { return v.Service }).(pulumi.StringOutput)
}

// URLs to SslCertificate resources that are used to authenticate connections to Backends. At least one SSL certificate must be specified. Currently, you may specify up to 15 SSL certificates. sslCertificates do not apply when the load balancing scheme is set to INTERNAL_SELF_MANAGED.
func (o LookupTargetSslProxyResultOutput) SslCertificates() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupTargetSslProxyResult) []string { return v.SslCertificates }).(pulumi.StringArrayOutput)
}

// URL of SslPolicy resource that will be associated with the TargetSslProxy resource. If not set, the TargetSslProxy resource will not have any SSL policy configured.
func (o LookupTargetSslProxyResultOutput) SslPolicy() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTargetSslProxyResult) string { return v.SslPolicy }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupTargetSslProxyResultOutput{})
}
