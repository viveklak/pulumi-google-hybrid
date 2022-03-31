// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets an attestor. Returns NOT_FOUND if the attestor does not exist.
func LookupAttestor(ctx *pulumi.Context, args *LookupAttestorArgs, opts ...pulumi.InvokeOption) (*LookupAttestorResult, error) {
	var rv LookupAttestorResult
	err := ctx.Invoke("google-native:binaryauthorization/v1beta1:getAttestor", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAttestorArgs struct {
	AttestorId string  `pulumi:"attestorId"`
	Project    *string `pulumi:"project"`
}

type LookupAttestorResult struct {
	// Optional. A descriptive comment. This field may be updated. The field may be displayed in chooser dialogs.
	Description string `pulumi:"description"`
	// Optional. Used to prevent updating the attestor when another request has updated it since it was retrieved.
	Etag string `pulumi:"etag"`
	// The resource name, in the format: `projects/*/attestors/*`. This field may not be updated.
	Name string `pulumi:"name"`
	// Time when the attestor was last updated.
	UpdateTime string `pulumi:"updateTime"`
	// A Drydock ATTESTATION_AUTHORITY Note, created by the user.
	UserOwnedDrydockNote UserOwnedDrydockNoteResponse `pulumi:"userOwnedDrydockNote"`
}

func LookupAttestorOutput(ctx *pulumi.Context, args LookupAttestorOutputArgs, opts ...pulumi.InvokeOption) LookupAttestorResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAttestorResult, error) {
			args := v.(LookupAttestorArgs)
			r, err := LookupAttestor(ctx, &args, opts...)
			return *r, err
		}).(LookupAttestorResultOutput)
}

type LookupAttestorOutputArgs struct {
	AttestorId pulumi.StringInput    `pulumi:"attestorId"`
	Project    pulumi.StringPtrInput `pulumi:"project"`
}

func (LookupAttestorOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAttestorArgs)(nil)).Elem()
}

type LookupAttestorResultOutput struct{ *pulumi.OutputState }

func (LookupAttestorResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAttestorResult)(nil)).Elem()
}

func (o LookupAttestorResultOutput) ToLookupAttestorResultOutput() LookupAttestorResultOutput {
	return o
}

func (o LookupAttestorResultOutput) ToLookupAttestorResultOutputWithContext(ctx context.Context) LookupAttestorResultOutput {
	return o
}

// Optional. A descriptive comment. This field may be updated. The field may be displayed in chooser dialogs.
func (o LookupAttestorResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAttestorResult) string { return v.Description }).(pulumi.StringOutput)
}

// Optional. Used to prevent updating the attestor when another request has updated it since it was retrieved.
func (o LookupAttestorResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAttestorResult) string { return v.Etag }).(pulumi.StringOutput)
}

// The resource name, in the format: `projects/*/attestors/*`. This field may not be updated.
func (o LookupAttestorResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAttestorResult) string { return v.Name }).(pulumi.StringOutput)
}

// Time when the attestor was last updated.
func (o LookupAttestorResultOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAttestorResult) string { return v.UpdateTime }).(pulumi.StringOutput)
}

// A Drydock ATTESTATION_AUTHORITY Note, created by the user.
func (o LookupAttestorResultOutput) UserOwnedDrydockNote() UserOwnedDrydockNoteResponseOutput {
	return o.ApplyT(func(v LookupAttestorResult) UserOwnedDrydockNoteResponse { return v.UserOwnedDrydockNote }).(UserOwnedDrydockNoteResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAttestorResultOutput{})
}
