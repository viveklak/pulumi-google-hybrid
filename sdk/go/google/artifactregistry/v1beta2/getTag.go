// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets a tag.
func LookupTag(ctx *pulumi.Context, args *LookupTagArgs, opts ...pulumi.InvokeOption) (*LookupTagResult, error) {
	var rv LookupTagResult
	err := ctx.Invoke("google-native:artifactregistry/v1beta2:getTag", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupTagArgs struct {
	Location     string  `pulumi:"location"`
	PackageId    string  `pulumi:"packageId"`
	Project      *string `pulumi:"project"`
	RepositoryId string  `pulumi:"repositoryId"`
	TagId        string  `pulumi:"tagId"`
}

type LookupTagResult struct {
	// The name of the tag, for example: "projects/p1/locations/us-central1/repositories/repo1/packages/pkg1/tags/tag1". If the package part contains slashes, the slashes are escaped. The tag part can only have characters in [a-zA-Z0-9\-._~:@], anything else must be URL encoded.
	Name string `pulumi:"name"`
	// The name of the version the tag refers to, for example: "projects/p1/locations/us-central1/repositories/repo1/packages/pkg1/versions/sha256:5243811" If the package or version ID parts contain slashes, the slashes are escaped.
	Version string `pulumi:"version"`
}

func LookupTagOutput(ctx *pulumi.Context, args LookupTagOutputArgs, opts ...pulumi.InvokeOption) LookupTagResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupTagResult, error) {
			args := v.(LookupTagArgs)
			r, err := LookupTag(ctx, &args, opts...)
			return *r, err
		}).(LookupTagResultOutput)
}

type LookupTagOutputArgs struct {
	Location     pulumi.StringInput    `pulumi:"location"`
	PackageId    pulumi.StringInput    `pulumi:"packageId"`
	Project      pulumi.StringPtrInput `pulumi:"project"`
	RepositoryId pulumi.StringInput    `pulumi:"repositoryId"`
	TagId        pulumi.StringInput    `pulumi:"tagId"`
}

func (LookupTagOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTagArgs)(nil)).Elem()
}

type LookupTagResultOutput struct{ *pulumi.OutputState }

func (LookupTagResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTagResult)(nil)).Elem()
}

func (o LookupTagResultOutput) ToLookupTagResultOutput() LookupTagResultOutput {
	return o
}

func (o LookupTagResultOutput) ToLookupTagResultOutputWithContext(ctx context.Context) LookupTagResultOutput {
	return o
}

// The name of the tag, for example: "projects/p1/locations/us-central1/repositories/repo1/packages/pkg1/tags/tag1". If the package part contains slashes, the slashes are escaped. The tag part can only have characters in [a-zA-Z0-9\-._~:@], anything else must be URL encoded.
func (o LookupTagResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTagResult) string { return v.Name }).(pulumi.StringOutput)
}

// The name of the version the tag refers to, for example: "projects/p1/locations/us-central1/repositories/repo1/packages/pkg1/versions/sha256:5243811" If the package or version ID parts contain slashes, the slashes are escaped.
func (o LookupTagResultOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTagResult) string { return v.Version }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupTagResultOutput{})
}
