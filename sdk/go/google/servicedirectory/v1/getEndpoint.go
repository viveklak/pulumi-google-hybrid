// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets an endpoint.
func LookupEndpoint(ctx *pulumi.Context, args *LookupEndpointArgs, opts ...pulumi.InvokeOption) (*LookupEndpointResult, error) {
	var rv LookupEndpointResult
	err := ctx.Invoke("google-native:servicedirectory/v1:getEndpoint", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupEndpointArgs struct {
	EndpointId  string `pulumi:"endpointId"`
	Location    string `pulumi:"location"`
	NamespaceId string `pulumi:"namespaceId"`
	Project     string `pulumi:"project"`
	ServiceId   string `pulumi:"serviceId"`
}

type LookupEndpointResult struct {
	// Optional. An IPv4 or IPv6 address. Service Directory rejects bad addresses like: * `8.8.8` * `8.8.8.8:53` * `test:bad:address` * `[::1]` * `[::1]:8080` Limited to 45 characters.
	Address string `pulumi:"address"`
	// Optional. Annotations for the endpoint. This data can be consumed by service clients. Restrictions: * The entire annotations dictionary may contain up to 512 characters, spread accoss all key-value pairs. Annotations that go beyond this limit are rejected * Valid annotation keys have two segments: an optional prefix and name, separated by a slash (/). The name segment is required and must be 63 characters or less, beginning and ending with an alphanumeric character ([a-z0-9A-Z]) with dashes (-), underscores (_), dots (.), and alphanumerics between. The prefix is optional. If specified, the prefix must be a DNS subdomain: a series of DNS labels separated by dots (.), not longer than 253 characters in total, followed by a slash (/) Annotations that fails to meet these requirements are rejected. * The `(*.)google.com/` and `(*.)googleapis.com/` prefixes are reserved for system annotations managed by Service Directory. If the user tries to write to these keyspaces, those entries are silently ignored by the system Note: This field is equivalent to the `metadata` field in the v1beta1 API. They have the same syntax and read/write to the same location in Service Directory.
	Annotations map[string]string `pulumi:"annotations"`
	// Immutable. The resource name for the endpoint in the format `projects/*/locations/*/namespaces/*/services/*/endpoints/*`.
	Name string `pulumi:"name"`
	// Optional. Service Directory rejects values outside of `[0, 65535]`.
	Port int `pulumi:"port"`
}