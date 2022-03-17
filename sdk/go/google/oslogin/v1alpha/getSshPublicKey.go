// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1alpha

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieves an SSH public key.
func LookupSshPublicKey(ctx *pulumi.Context, args *LookupSshPublicKeyArgs, opts ...pulumi.InvokeOption) (*LookupSshPublicKeyResult, error) {
	var rv LookupSshPublicKeyResult
	err := ctx.Invoke("google-native:oslogin/v1alpha:getSshPublicKey", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSshPublicKeyArgs struct {
	SshPublicKeyId string `pulumi:"sshPublicKeyId"`
	UserId         string `pulumi:"userId"`
}

type LookupSshPublicKeyResult struct {
	// An expiration time in microseconds since epoch.
	ExpirationTimeUsec string `pulumi:"expirationTimeUsec"`
	// The SHA-256 fingerprint of the SSH public key.
	Fingerprint string `pulumi:"fingerprint"`
	// Public key text in SSH format, defined by RFC4253 section 6.6.
	Key string `pulumi:"key"`
	// The canonical resource name.
	Name string `pulumi:"name"`
}

func LookupSshPublicKeyOutput(ctx *pulumi.Context, args LookupSshPublicKeyOutputArgs, opts ...pulumi.InvokeOption) LookupSshPublicKeyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSshPublicKeyResult, error) {
			args := v.(LookupSshPublicKeyArgs)
			r, err := LookupSshPublicKey(ctx, &args, opts...)
			return *r, err
		}).(LookupSshPublicKeyResultOutput)
}

type LookupSshPublicKeyOutputArgs struct {
	SshPublicKeyId pulumi.StringInput `pulumi:"sshPublicKeyId"`
	UserId         pulumi.StringInput `pulumi:"userId"`
}

func (LookupSshPublicKeyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSshPublicKeyArgs)(nil)).Elem()
}

type LookupSshPublicKeyResultOutput struct{ *pulumi.OutputState }

func (LookupSshPublicKeyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSshPublicKeyResult)(nil)).Elem()
}

func (o LookupSshPublicKeyResultOutput) ToLookupSshPublicKeyResultOutput() LookupSshPublicKeyResultOutput {
	return o
}

func (o LookupSshPublicKeyResultOutput) ToLookupSshPublicKeyResultOutputWithContext(ctx context.Context) LookupSshPublicKeyResultOutput {
	return o
}

// An expiration time in microseconds since epoch.
func (o LookupSshPublicKeyResultOutput) ExpirationTimeUsec() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSshPublicKeyResult) string { return v.ExpirationTimeUsec }).(pulumi.StringOutput)
}

// The SHA-256 fingerprint of the SSH public key.
func (o LookupSshPublicKeyResultOutput) Fingerprint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSshPublicKeyResult) string { return v.Fingerprint }).(pulumi.StringOutput)
}

// Public key text in SSH format, defined by RFC4253 section 6.6.
func (o LookupSshPublicKeyResultOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSshPublicKeyResult) string { return v.Key }).(pulumi.StringOutput)
}

// The canonical resource name.
func (o LookupSshPublicKeyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSshPublicKeyResult) string { return v.Name }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSshPublicKeyResultOutput{})
}
