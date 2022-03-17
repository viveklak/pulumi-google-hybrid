// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets details of a single connection profile.
func LookupConnectionProfile(ctx *pulumi.Context, args *LookupConnectionProfileArgs, opts ...pulumi.InvokeOption) (*LookupConnectionProfileResult, error) {
	var rv LookupConnectionProfileResult
	err := ctx.Invoke("google-native:datamigration/v1beta1:getConnectionProfile", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupConnectionProfileArgs struct {
	ConnectionProfileId string  `pulumi:"connectionProfileId"`
	Location            string  `pulumi:"location"`
	Project             *string `pulumi:"project"`
}

type LookupConnectionProfileResult struct {
	// A CloudSQL database connection profile.
	Cloudsql CloudSqlConnectionProfileResponse `pulumi:"cloudsql"`
	// The timestamp when the resource was created. A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".
	CreateTime string `pulumi:"createTime"`
	// The connection profile display name.
	DisplayName string `pulumi:"displayName"`
	// The error details in case of state FAILED.
	Error StatusResponse `pulumi:"error"`
	// The resource labels for connection profile to use to annotate any related underlying resources such as Compute Engine VMs. An object containing a list of "key": "value" pairs. Example: `{ "name": "wrench", "mass": "1.3kg", "count": "3" }`.
	Labels map[string]string `pulumi:"labels"`
	// A MySQL database connection profile.
	Mysql MySqlConnectionProfileResponse `pulumi:"mysql"`
	// The name of this connection profile resource in the form of projects/{project}/locations/{location}/connectionProfiles/{connectionProfile}.
	Name string `pulumi:"name"`
	// The database provider.
	Provider string `pulumi:"provider"`
	// The current connection profile state (e.g. DRAFT, READY, or FAILED).
	State string `pulumi:"state"`
	// The timestamp when the resource was last updated. A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".
	UpdateTime string `pulumi:"updateTime"`
}

func LookupConnectionProfileOutput(ctx *pulumi.Context, args LookupConnectionProfileOutputArgs, opts ...pulumi.InvokeOption) LookupConnectionProfileResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupConnectionProfileResult, error) {
			args := v.(LookupConnectionProfileArgs)
			r, err := LookupConnectionProfile(ctx, &args, opts...)
			return *r, err
		}).(LookupConnectionProfileResultOutput)
}

type LookupConnectionProfileOutputArgs struct {
	ConnectionProfileId pulumi.StringInput    `pulumi:"connectionProfileId"`
	Location            pulumi.StringInput    `pulumi:"location"`
	Project             pulumi.StringPtrInput `pulumi:"project"`
}

func (LookupConnectionProfileOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConnectionProfileArgs)(nil)).Elem()
}

type LookupConnectionProfileResultOutput struct{ *pulumi.OutputState }

func (LookupConnectionProfileResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConnectionProfileResult)(nil)).Elem()
}

func (o LookupConnectionProfileResultOutput) ToLookupConnectionProfileResultOutput() LookupConnectionProfileResultOutput {
	return o
}

func (o LookupConnectionProfileResultOutput) ToLookupConnectionProfileResultOutputWithContext(ctx context.Context) LookupConnectionProfileResultOutput {
	return o
}

// A CloudSQL database connection profile.
func (o LookupConnectionProfileResultOutput) Cloudsql() CloudSqlConnectionProfileResponseOutput {
	return o.ApplyT(func(v LookupConnectionProfileResult) CloudSqlConnectionProfileResponse { return v.Cloudsql }).(CloudSqlConnectionProfileResponseOutput)
}

// The timestamp when the resource was created. A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".
func (o LookupConnectionProfileResultOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectionProfileResult) string { return v.CreateTime }).(pulumi.StringOutput)
}

// The connection profile display name.
func (o LookupConnectionProfileResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectionProfileResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

// The error details in case of state FAILED.
func (o LookupConnectionProfileResultOutput) Error() StatusResponseOutput {
	return o.ApplyT(func(v LookupConnectionProfileResult) StatusResponse { return v.Error }).(StatusResponseOutput)
}

// The resource labels for connection profile to use to annotate any related underlying resources such as Compute Engine VMs. An object containing a list of "key": "value" pairs. Example: `{ "name": "wrench", "mass": "1.3kg", "count": "3" }`.
func (o LookupConnectionProfileResultOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupConnectionProfileResult) map[string]string { return v.Labels }).(pulumi.StringMapOutput)
}

// A MySQL database connection profile.
func (o LookupConnectionProfileResultOutput) Mysql() MySqlConnectionProfileResponseOutput {
	return o.ApplyT(func(v LookupConnectionProfileResult) MySqlConnectionProfileResponse { return v.Mysql }).(MySqlConnectionProfileResponseOutput)
}

// The name of this connection profile resource in the form of projects/{project}/locations/{location}/connectionProfiles/{connectionProfile}.
func (o LookupConnectionProfileResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectionProfileResult) string { return v.Name }).(pulumi.StringOutput)
}

// The database provider.
func (o LookupConnectionProfileResultOutput) Provider() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectionProfileResult) string { return v.Provider }).(pulumi.StringOutput)
}

// The current connection profile state (e.g. DRAFT, READY, or FAILED).
func (o LookupConnectionProfileResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectionProfileResult) string { return v.State }).(pulumi.StringOutput)
}

// The timestamp when the resource was last updated. A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".
func (o LookupConnectionProfileResultOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectionProfileResult) string { return v.UpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupConnectionProfileResultOutput{})
}
