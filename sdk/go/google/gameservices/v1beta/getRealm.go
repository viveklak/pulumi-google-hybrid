// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets details of a single realm.
func LookupRealm(ctx *pulumi.Context, args *LookupRealmArgs, opts ...pulumi.InvokeOption) (*LookupRealmResult, error) {
	var rv LookupRealmResult
	err := ctx.Invoke("google-native:gameservices/v1beta:getRealm", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupRealmArgs struct {
	Location string  `pulumi:"location"`
	Project  *string `pulumi:"project"`
	RealmId  string  `pulumi:"realmId"`
}

type LookupRealmResult struct {
	// The creation time.
	CreateTime string `pulumi:"createTime"`
	// Human readable description of the realm.
	Description string `pulumi:"description"`
	// Used to perform consistent read-modify-write updates. If not set, a blind "overwrite" update happens.
	Etag string `pulumi:"etag"`
	// The labels associated with this realm. Each label is a key-value pair.
	Labels map[string]string `pulumi:"labels"`
	// The resource name of the realm, in the following form: `projects/{project}/locations/{locationId}/realms/{realmId}`. For example, `projects/my-project/locations/global/realms/my-realm`.
	Name string `pulumi:"name"`
	// Time zone where all policies targeting this realm are evaluated. The value of this field must be from the [IANA time zone database](https://www.iana.org/time-zones).
	TimeZone string `pulumi:"timeZone"`
	// The last-modified time.
	UpdateTime string `pulumi:"updateTime"`
}

func LookupRealmOutput(ctx *pulumi.Context, args LookupRealmOutputArgs, opts ...pulumi.InvokeOption) LookupRealmResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRealmResult, error) {
			args := v.(LookupRealmArgs)
			r, err := LookupRealm(ctx, &args, opts...)
			return *r, err
		}).(LookupRealmResultOutput)
}

type LookupRealmOutputArgs struct {
	Location pulumi.StringInput    `pulumi:"location"`
	Project  pulumi.StringPtrInput `pulumi:"project"`
	RealmId  pulumi.StringInput    `pulumi:"realmId"`
}

func (LookupRealmOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRealmArgs)(nil)).Elem()
}

type LookupRealmResultOutput struct{ *pulumi.OutputState }

func (LookupRealmResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRealmResult)(nil)).Elem()
}

func (o LookupRealmResultOutput) ToLookupRealmResultOutput() LookupRealmResultOutput {
	return o
}

func (o LookupRealmResultOutput) ToLookupRealmResultOutputWithContext(ctx context.Context) LookupRealmResultOutput {
	return o
}

// The creation time.
func (o LookupRealmResultOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRealmResult) string { return v.CreateTime }).(pulumi.StringOutput)
}

// Human readable description of the realm.
func (o LookupRealmResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRealmResult) string { return v.Description }).(pulumi.StringOutput)
}

// Used to perform consistent read-modify-write updates. If not set, a blind "overwrite" update happens.
func (o LookupRealmResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRealmResult) string { return v.Etag }).(pulumi.StringOutput)
}

// The labels associated with this realm. Each label is a key-value pair.
func (o LookupRealmResultOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupRealmResult) map[string]string { return v.Labels }).(pulumi.StringMapOutput)
}

// The resource name of the realm, in the following form: `projects/{project}/locations/{locationId}/realms/{realmId}`. For example, `projects/my-project/locations/global/realms/my-realm`.
func (o LookupRealmResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRealmResult) string { return v.Name }).(pulumi.StringOutput)
}

// Time zone where all policies targeting this realm are evaluated. The value of this field must be from the [IANA time zone database](https://www.iana.org/time-zones).
func (o LookupRealmResultOutput) TimeZone() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRealmResult) string { return v.TimeZone }).(pulumi.StringOutput)
}

// The last-modified time.
func (o LookupRealmResultOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRealmResult) string { return v.UpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRealmResultOutput{})
}
