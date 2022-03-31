// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package alpha

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets information about a specific type provider.
func LookupTypeProvider(ctx *pulumi.Context, args *LookupTypeProviderArgs, opts ...pulumi.InvokeOption) (*LookupTypeProviderResult, error) {
	var rv LookupTypeProviderResult
	err := ctx.Invoke("google-native:deploymentmanager/alpha:getTypeProvider", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupTypeProviderArgs struct {
	Project      *string `pulumi:"project"`
	TypeProvider string  `pulumi:"typeProvider"`
}

type LookupTypeProviderResult struct {
	// Allows resource handling overrides for specific collections
	CollectionOverrides []CollectionOverrideResponse `pulumi:"collectionOverrides"`
	// Credential used when interacting with this type.
	Credential CredentialResponse `pulumi:"credential"`
	// List of up to 2 custom certificate authority roots to use for TLS authentication when making calls on behalf of this type provider. If set, TLS authentication will exclusively use these roots instead of relying on publicly trusted certificate authorities when validating TLS certificate authenticity. The certificates must be in base64-encoded PEM format. The maximum size of each certificate must not exceed 10KB.
	CustomCertificateAuthorityRoots []string `pulumi:"customCertificateAuthorityRoots"`
	// An optional textual description of the resource; provided by the client when the resource is created.
	Description string `pulumi:"description"`
	// Descriptor Url for the this type provider.
	DescriptorUrl string `pulumi:"descriptorUrl"`
	// Creation timestamp in RFC3339 text format.
	InsertTime string `pulumi:"insertTime"`
	// Map of One Platform labels; provided by the client when the resource is created or updated. Specifically: Label keys must be between 1 and 63 characters long and must conform to the following regular expression: `[a-z]([-a-z0-9]*[a-z0-9])?` Label values must be between 0 and 63 characters long and must conform to the regular expression `([a-z]([-a-z0-9]*[a-z0-9])?)?`
	Labels []TypeProviderLabelEntryResponse `pulumi:"labels"`
	// Name of the resource; provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name string `pulumi:"name"`
	// The Operation that most recently ran, or is currently running, on this type provider.
	Operation OperationResponse `pulumi:"operation"`
	// Options to apply when handling any resources in this service.
	Options OptionsResponse `pulumi:"options"`
	// Self link for the type provider.
	SelfLink string `pulumi:"selfLink"`
}

func LookupTypeProviderOutput(ctx *pulumi.Context, args LookupTypeProviderOutputArgs, opts ...pulumi.InvokeOption) LookupTypeProviderResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupTypeProviderResult, error) {
			args := v.(LookupTypeProviderArgs)
			r, err := LookupTypeProvider(ctx, &args, opts...)
			return *r, err
		}).(LookupTypeProviderResultOutput)
}

type LookupTypeProviderOutputArgs struct {
	Project      pulumi.StringPtrInput `pulumi:"project"`
	TypeProvider pulumi.StringInput    `pulumi:"typeProvider"`
}

func (LookupTypeProviderOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTypeProviderArgs)(nil)).Elem()
}

type LookupTypeProviderResultOutput struct{ *pulumi.OutputState }

func (LookupTypeProviderResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTypeProviderResult)(nil)).Elem()
}

func (o LookupTypeProviderResultOutput) ToLookupTypeProviderResultOutput() LookupTypeProviderResultOutput {
	return o
}

func (o LookupTypeProviderResultOutput) ToLookupTypeProviderResultOutputWithContext(ctx context.Context) LookupTypeProviderResultOutput {
	return o
}

// Allows resource handling overrides for specific collections
func (o LookupTypeProviderResultOutput) CollectionOverrides() CollectionOverrideResponseArrayOutput {
	return o.ApplyT(func(v LookupTypeProviderResult) []CollectionOverrideResponse { return v.CollectionOverrides }).(CollectionOverrideResponseArrayOutput)
}

// Credential used when interacting with this type.
func (o LookupTypeProviderResultOutput) Credential() CredentialResponseOutput {
	return o.ApplyT(func(v LookupTypeProviderResult) CredentialResponse { return v.Credential }).(CredentialResponseOutput)
}

// List of up to 2 custom certificate authority roots to use for TLS authentication when making calls on behalf of this type provider. If set, TLS authentication will exclusively use these roots instead of relying on publicly trusted certificate authorities when validating TLS certificate authenticity. The certificates must be in base64-encoded PEM format. The maximum size of each certificate must not exceed 10KB.
func (o LookupTypeProviderResultOutput) CustomCertificateAuthorityRoots() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupTypeProviderResult) []string { return v.CustomCertificateAuthorityRoots }).(pulumi.StringArrayOutput)
}

// An optional textual description of the resource; provided by the client when the resource is created.
func (o LookupTypeProviderResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTypeProviderResult) string { return v.Description }).(pulumi.StringOutput)
}

// Descriptor Url for the this type provider.
func (o LookupTypeProviderResultOutput) DescriptorUrl() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTypeProviderResult) string { return v.DescriptorUrl }).(pulumi.StringOutput)
}

// Creation timestamp in RFC3339 text format.
func (o LookupTypeProviderResultOutput) InsertTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTypeProviderResult) string { return v.InsertTime }).(pulumi.StringOutput)
}

// Map of One Platform labels; provided by the client when the resource is created or updated. Specifically: Label keys must be between 1 and 63 characters long and must conform to the following regular expression: `[a-z]([-a-z0-9]*[a-z0-9])?` Label values must be between 0 and 63 characters long and must conform to the regular expression `([a-z]([-a-z0-9]*[a-z0-9])?)?`
func (o LookupTypeProviderResultOutput) Labels() TypeProviderLabelEntryResponseArrayOutput {
	return o.ApplyT(func(v LookupTypeProviderResult) []TypeProviderLabelEntryResponse { return v.Labels }).(TypeProviderLabelEntryResponseArrayOutput)
}

// Name of the resource; provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
func (o LookupTypeProviderResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTypeProviderResult) string { return v.Name }).(pulumi.StringOutput)
}

// The Operation that most recently ran, or is currently running, on this type provider.
func (o LookupTypeProviderResultOutput) Operation() OperationResponseOutput {
	return o.ApplyT(func(v LookupTypeProviderResult) OperationResponse { return v.Operation }).(OperationResponseOutput)
}

// Options to apply when handling any resources in this service.
func (o LookupTypeProviderResultOutput) Options() OptionsResponseOutput {
	return o.ApplyT(func(v LookupTypeProviderResult) OptionsResponse { return v.Options }).(OptionsResponseOutput)
}

// Self link for the type provider.
func (o LookupTypeProviderResultOutput) SelfLink() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTypeProviderResult) string { return v.SelfLink }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupTypeProviderResultOutput{})
}
