// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the GcpUserAccessBinding with the given name.
func LookupGcpUserAccessBinding(ctx *pulumi.Context, args *LookupGcpUserAccessBindingArgs, opts ...pulumi.InvokeOption) (*LookupGcpUserAccessBindingResult, error) {
	var rv LookupGcpUserAccessBindingResult
	err := ctx.Invoke("google-native:accesscontextmanager/v1:getGcpUserAccessBinding", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupGcpUserAccessBindingArgs struct {
	GcpUserAccessBindingId string `pulumi:"gcpUserAccessBindingId"`
	OrganizationId         string `pulumi:"organizationId"`
}

type LookupGcpUserAccessBindingResult struct {
	// Access level that a user must have to be granted access. Only one access level is supported, not multiple. This repeated field must have exactly one element. Example: "accessPolicies/9522/accessLevels/device_trusted"
	AccessLevels []string `pulumi:"accessLevels"`
	// Immutable. Google Group id whose members are subject to this binding's restrictions. See "id" in the [G Suite Directory API's Groups resource] (https://developers.google.com/admin-sdk/directory/v1/reference/groups#resource). If a group's email address/alias is changed, this resource will continue to point at the changed group. This field does not accept group email addresses or aliases. Example: "01d520gv4vjcrht"
	GroupKey string `pulumi:"groupKey"`
	// Immutable. Assigned by the server during creation. The last segment has an arbitrary length and has only URI unreserved characters (as defined by [RFC 3986 Section 2.3](https://tools.ietf.org/html/rfc3986#section-2.3)). Should not be specified by the client during creation. Example: "organizations/256/gcpUserAccessBindings/b3-BhcX_Ud5N"
	Name string `pulumi:"name"`
}

func LookupGcpUserAccessBindingOutput(ctx *pulumi.Context, args LookupGcpUserAccessBindingOutputArgs, opts ...pulumi.InvokeOption) LookupGcpUserAccessBindingResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupGcpUserAccessBindingResult, error) {
			args := v.(LookupGcpUserAccessBindingArgs)
			r, err := LookupGcpUserAccessBinding(ctx, &args, opts...)
			return *r, err
		}).(LookupGcpUserAccessBindingResultOutput)
}

type LookupGcpUserAccessBindingOutputArgs struct {
	GcpUserAccessBindingId pulumi.StringInput `pulumi:"gcpUserAccessBindingId"`
	OrganizationId         pulumi.StringInput `pulumi:"organizationId"`
}

func (LookupGcpUserAccessBindingOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGcpUserAccessBindingArgs)(nil)).Elem()
}

type LookupGcpUserAccessBindingResultOutput struct{ *pulumi.OutputState }

func (LookupGcpUserAccessBindingResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGcpUserAccessBindingResult)(nil)).Elem()
}

func (o LookupGcpUserAccessBindingResultOutput) ToLookupGcpUserAccessBindingResultOutput() LookupGcpUserAccessBindingResultOutput {
	return o
}

func (o LookupGcpUserAccessBindingResultOutput) ToLookupGcpUserAccessBindingResultOutputWithContext(ctx context.Context) LookupGcpUserAccessBindingResultOutput {
	return o
}

// Access level that a user must have to be granted access. Only one access level is supported, not multiple. This repeated field must have exactly one element. Example: "accessPolicies/9522/accessLevels/device_trusted"
func (o LookupGcpUserAccessBindingResultOutput) AccessLevels() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupGcpUserAccessBindingResult) []string { return v.AccessLevels }).(pulumi.StringArrayOutput)
}

// Immutable. Google Group id whose members are subject to this binding's restrictions. See "id" in the [G Suite Directory API's Groups resource] (https://developers.google.com/admin-sdk/directory/v1/reference/groups#resource). If a group's email address/alias is changed, this resource will continue to point at the changed group. This field does not accept group email addresses or aliases. Example: "01d520gv4vjcrht"
func (o LookupGcpUserAccessBindingResultOutput) GroupKey() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGcpUserAccessBindingResult) string { return v.GroupKey }).(pulumi.StringOutput)
}

// Immutable. Assigned by the server during creation. The last segment has an arbitrary length and has only URI unreserved characters (as defined by [RFC 3986 Section 2.3](https://tools.ietf.org/html/rfc3986#section-2.3)). Should not be specified by the client during creation. Example: "organizations/256/gcpUserAccessBindings/b3-BhcX_Ud5N"
func (o LookupGcpUserAccessBindingResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGcpUserAccessBindingResult) string { return v.Name }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupGcpUserAccessBindingResultOutput{})
}
