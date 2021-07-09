// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets details of a single realm.
func LookupRealm(ctx *pulumi.Context, args *LookupRealmArgs, opts ...pulumi.InvokeOption) (*LookupRealmResult, error) {
	var rv LookupRealmResult
	err := ctx.Invoke("google-native:gameservices/v1:getRealm", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupRealmArgs struct {
	Location string `pulumi:"location"`
	Project  string `pulumi:"project"`
	RealmId  string `pulumi:"realmId"`
}

type LookupRealmResult struct {
	// The creation time.
	CreateTime string `pulumi:"createTime"`
	// Human readable description of the realm.
	Description string `pulumi:"description"`
	// ETag of the resource.
	Etag string `pulumi:"etag"`
	// The labels associated with this realm. Each label is a key-value pair.
	Labels map[string]string `pulumi:"labels"`
	// The resource name of the realm, in the following form: `projects/{project}/locations/{location}/realms/{realm}`. For example, `projects/my-project/locations/{location}/realms/my-realm`.
	Name string `pulumi:"name"`
	// Time zone where all policies targeting this realm are evaluated. The value of this field must be from the IANA time zone database: https://www.iana.org/time-zones.
	TimeZone string `pulumi:"timeZone"`
	// The last-modified time.
	UpdateTime string `pulumi:"updateTime"`
}
