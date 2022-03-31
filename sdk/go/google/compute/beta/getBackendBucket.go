// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package beta

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Returns the specified BackendBucket resource. Gets a list of available backend buckets by making a list() request.
func LookupBackendBucket(ctx *pulumi.Context, args *LookupBackendBucketArgs, opts ...pulumi.InvokeOption) (*LookupBackendBucketResult, error) {
	var rv LookupBackendBucketResult
	err := ctx.Invoke("google-native:compute/beta:getBackendBucket", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupBackendBucketArgs struct {
	BackendBucket string  `pulumi:"backendBucket"`
	Project       *string `pulumi:"project"`
}

type LookupBackendBucketResult struct {
	// Cloud Storage bucket name.
	BucketName string `pulumi:"bucketName"`
	// Cloud CDN configuration for this BackendBucket.
	CdnPolicy BackendBucketCdnPolicyResponse `pulumi:"cdnPolicy"`
	// Compress text responses using Brotli or gzip compression, based on the client's Accept-Encoding header.
	CompressionMode string `pulumi:"compressionMode"`
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp string `pulumi:"creationTimestamp"`
	// Headers that the HTTP/S load balancer should add to proxied responses.
	CustomResponseHeaders []string `pulumi:"customResponseHeaders"`
	// An optional textual description of the resource; provided by the client when the resource is created.
	Description string `pulumi:"description"`
	// The resource URL for the edge security policy associated with this backend bucket.
	EdgeSecurityPolicy string `pulumi:"edgeSecurityPolicy"`
	// If true, enable Cloud CDN for this BackendBucket.
	EnableCdn bool `pulumi:"enableCdn"`
	// Type of the resource.
	Kind string `pulumi:"kind"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name string `pulumi:"name"`
	// Server-defined URL for the resource.
	SelfLink string `pulumi:"selfLink"`
}

func LookupBackendBucketOutput(ctx *pulumi.Context, args LookupBackendBucketOutputArgs, opts ...pulumi.InvokeOption) LookupBackendBucketResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupBackendBucketResult, error) {
			args := v.(LookupBackendBucketArgs)
			r, err := LookupBackendBucket(ctx, &args, opts...)
			return *r, err
		}).(LookupBackendBucketResultOutput)
}

type LookupBackendBucketOutputArgs struct {
	BackendBucket pulumi.StringInput    `pulumi:"backendBucket"`
	Project       pulumi.StringPtrInput `pulumi:"project"`
}

func (LookupBackendBucketOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBackendBucketArgs)(nil)).Elem()
}

type LookupBackendBucketResultOutput struct{ *pulumi.OutputState }

func (LookupBackendBucketResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBackendBucketResult)(nil)).Elem()
}

func (o LookupBackendBucketResultOutput) ToLookupBackendBucketResultOutput() LookupBackendBucketResultOutput {
	return o
}

func (o LookupBackendBucketResultOutput) ToLookupBackendBucketResultOutputWithContext(ctx context.Context) LookupBackendBucketResultOutput {
	return o
}

// Cloud Storage bucket name.
func (o LookupBackendBucketResultOutput) BucketName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBackendBucketResult) string { return v.BucketName }).(pulumi.StringOutput)
}

// Cloud CDN configuration for this BackendBucket.
func (o LookupBackendBucketResultOutput) CdnPolicy() BackendBucketCdnPolicyResponseOutput {
	return o.ApplyT(func(v LookupBackendBucketResult) BackendBucketCdnPolicyResponse { return v.CdnPolicy }).(BackendBucketCdnPolicyResponseOutput)
}

// Compress text responses using Brotli or gzip compression, based on the client's Accept-Encoding header.
func (o LookupBackendBucketResultOutput) CompressionMode() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBackendBucketResult) string { return v.CompressionMode }).(pulumi.StringOutput)
}

// Creation timestamp in RFC3339 text format.
func (o LookupBackendBucketResultOutput) CreationTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBackendBucketResult) string { return v.CreationTimestamp }).(pulumi.StringOutput)
}

// Headers that the HTTP/S load balancer should add to proxied responses.
func (o LookupBackendBucketResultOutput) CustomResponseHeaders() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupBackendBucketResult) []string { return v.CustomResponseHeaders }).(pulumi.StringArrayOutput)
}

// An optional textual description of the resource; provided by the client when the resource is created.
func (o LookupBackendBucketResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBackendBucketResult) string { return v.Description }).(pulumi.StringOutput)
}

// The resource URL for the edge security policy associated with this backend bucket.
func (o LookupBackendBucketResultOutput) EdgeSecurityPolicy() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBackendBucketResult) string { return v.EdgeSecurityPolicy }).(pulumi.StringOutput)
}

// If true, enable Cloud CDN for this BackendBucket.
func (o LookupBackendBucketResultOutput) EnableCdn() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupBackendBucketResult) bool { return v.EnableCdn }).(pulumi.BoolOutput)
}

// Type of the resource.
func (o LookupBackendBucketResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBackendBucketResult) string { return v.Kind }).(pulumi.StringOutput)
}

// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
func (o LookupBackendBucketResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBackendBucketResult) string { return v.Name }).(pulumi.StringOutput)
}

// Server-defined URL for the resource.
func (o LookupBackendBucketResultOutput) SelfLink() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBackendBucketResult) string { return v.SelfLink }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupBackendBucketResultOutput{})
}
