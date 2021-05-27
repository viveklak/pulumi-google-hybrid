// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates an API product in an organization. You create API products after you have proxied backend services using API proxies. An API product is a collection of API resources combined with quota settings and metadata that you can use to deliver customized and productized API bundles to your developer community. This metadata can include: - Scope - Environments - API proxies - Extensible profile API products enable you repackage APIs on-the-fly, without having to do any additional coding or configuration. Apigee recommends that you start with a simple API product including only required elements. You then provision credentials to apps to enable them to start testing your APIs. After you have authentication and authorization working against a simple API product, you can iterate to create finer grained API products, defining different sets of API resources for each API product. **WARNING:** - If you don't specify an API proxy in the request body, *any* app associated with the product can make calls to *any* API in your entire organization. - If you don't specify an environment in the request body, the product allows access to all environments. For more information, see What is an API product?
type OrganizationApiproduct struct {
	pulumi.CustomResourceState

	ApiResources pulumi.StringArrayOutput `pulumi:"apiResources"`
	// Flag that specifies how API keys are approved to access the APIs defined by the API product. If set to `manual`, the consumer key is generated and returned in "pending" state. In this case, the API keys won't work until they have been explicitly approved. If set to `auto`, the consumer key is generated and returned in "approved" state and can be used immediately. **Note:** Typically, `auto` is used to provide access to free or trial API products that provide limited quota or capabilities.
	ApprovalType pulumi.StringOutput `pulumi:"approvalType"`
	// Array of attributes that may be used to extend the default API product profile with customer-specific metadata. You can specify a maximum of 18 attributes. Use this property to specify the access level of the API product as either `public`, `private`, or `internal`. Only products marked `public` are available to developers in the Apigee developer portal. For example, you can set a product to `internal` while it is in development and then change access to `public` when it is ready to release on the portal. API products marked as `private` do not appear on the portal, but can be accessed by external developers.
	Attributes GoogleCloudApigeeV1AttributeResponseArrayOutput `pulumi:"attributes"`
	// Response only. Creation time of this environment as milliseconds since epoch.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// Description of the API product. Include key information about the API product that is not captured by other fields. Comma-separated list of API resources to be bundled in the API product. By default, the resource paths are mapped from the `proxy.pathsuffix` variable. The proxy path suffix is defined as the URI fragment following the ProxyEndpoint base path. For example, if the `apiResources` element is defined to be `/forecastrss` and the base path defined for the API proxy is `/weather`, then only requests to `/weather/forecastrss` are permitted by the API product. You can select a specific path, or you can select all subpaths with the following wildcard: - `/**`: Indicates that all sub-URIs are included. - `/*` : Indicates that only URIs one level down are included. By default, / supports the same resources as /** as well as the base path defined by the API proxy. For example, if the base path of the API proxy is `/v1/weatherapikey`, then the API product supports requests to `/v1/weatherapikey` and to any sub-URIs, such as `/v1/weatherapikey/forecastrss`, `/v1/weatherapikey/region/CA`, and so on. For more information, see Managing API products.
	Description pulumi.StringOutput `pulumi:"description"`
	// Name displayed in the UI or developer portal to developers registering for API access.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// Comma-separated list of environment names to which the API product is bound. Requests to environments that are not listed are rejected. By specifying one or more environments, you can bind the resources listed in the API product to a specific environment, preventing developers from accessing those resources through API proxies deployed in another environment. This setting is used, for example, to prevent resources associated with API proxies in `prod` from being accessed by API proxies deployed in `test`.
	Environments pulumi.StringArrayOutput `pulumi:"environments"`
	// Configuration used to group Apigee proxies or remote services with graphQL operation name, graphQL operation type and quotas. This grouping allows us to precisely set quota for a particular combination of graphQL name and operation type for a particular proxy request. If graphQL name is not set, this would imply quota will be applied on all graphQL requests matching the operation type.
	GraphqlOperationGroup GoogleCloudApigeeV1GraphQLOperationGroupResponseOutput `pulumi:"graphqlOperationGroup"`
	// Response only. Modified time of this environment as milliseconds since epoch.
	LastModifiedAt pulumi.StringOutput `pulumi:"lastModifiedAt"`
	// Internal name of the API product. Characters you can use in the name are restricted to: `A-Z0-9._\-$ %`. **Note:** The internal name cannot be edited when updating the API product.
	Name pulumi.StringOutput `pulumi:"name"`
	// Configuration used to group Apigee proxies or remote services with resources, method types, and quotas. The resource refers to the resource URI (excluding the base path). With this grouping, the API product creator is able to fine-tune and give precise control over which REST methods have access to specific resources and how many calls can be made (using the `quota` setting). **Note:** The `api_resources` setting cannot be specified for both the API product and operation group; otherwise the call will fail.
	OperationGroup GoogleCloudApigeeV1OperationGroupResponseOutput `pulumi:"operationGroup"`
	// Comma-separated list of API proxy names to which this API product is bound. By specifying API proxies, you can associate resources in the API product with specific API proxies, preventing developers from accessing those resources through other API proxies. Apigee rejects requests to API proxies that are not listed. **Note:** The API proxy names must already exist in the specified environment as they will be validated upon creation.
	Proxies pulumi.StringArrayOutput `pulumi:"proxies"`
	// Number of request messages permitted per app by this API product for the specified `quotaInterval` and `quotaTimeUnit`. For example, a `quota` of 50, for a `quotaInterval` of 12 and a `quotaTimeUnit` of hours means 50 requests are allowed every 12 hours.
	Quota pulumi.StringOutput `pulumi:"quota"`
	// Time interval over which the number of request messages is calculated.
	QuotaInterval pulumi.StringOutput `pulumi:"quotaInterval"`
	// Time unit defined for the `quotaInterval`. Valid values include `minute`, `hour`, `day`, or `month`.
	QuotaTimeUnit pulumi.StringOutput `pulumi:"quotaTimeUnit"`
	// Comma-separated list of OAuth scopes that are validated at runtime. Apigee validates that the scopes in any access token presented match the scopes defined in the OAuth policy associated with the API product.
	Scopes pulumi.StringArrayOutput `pulumi:"scopes"`
}

// NewOrganizationApiproduct registers a new resource with the given unique name, arguments, and options.
func NewOrganizationApiproduct(ctx *pulumi.Context,
	name string, args *OrganizationApiproductArgs, opts ...pulumi.ResourceOption) (*OrganizationApiproduct, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.OrganizationId == nil {
		return nil, errors.New("invalid value for required argument 'OrganizationId'")
	}
	var resource OrganizationApiproduct
	err := ctx.RegisterResource("google-native:apigee/v1:OrganizationApiproduct", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOrganizationApiproduct gets an existing OrganizationApiproduct resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOrganizationApiproduct(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OrganizationApiproductState, opts ...pulumi.ResourceOption) (*OrganizationApiproduct, error) {
	var resource OrganizationApiproduct
	err := ctx.ReadResource("google-native:apigee/v1:OrganizationApiproduct", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering OrganizationApiproduct resources.
type organizationApiproductState struct {
	ApiResources []string `pulumi:"apiResources"`
	// Flag that specifies how API keys are approved to access the APIs defined by the API product. If set to `manual`, the consumer key is generated and returned in "pending" state. In this case, the API keys won't work until they have been explicitly approved. If set to `auto`, the consumer key is generated and returned in "approved" state and can be used immediately. **Note:** Typically, `auto` is used to provide access to free or trial API products that provide limited quota or capabilities.
	ApprovalType *string `pulumi:"approvalType"`
	// Array of attributes that may be used to extend the default API product profile with customer-specific metadata. You can specify a maximum of 18 attributes. Use this property to specify the access level of the API product as either `public`, `private`, or `internal`. Only products marked `public` are available to developers in the Apigee developer portal. For example, you can set a product to `internal` while it is in development and then change access to `public` when it is ready to release on the portal. API products marked as `private` do not appear on the portal, but can be accessed by external developers.
	Attributes []GoogleCloudApigeeV1AttributeResponse `pulumi:"attributes"`
	// Response only. Creation time of this environment as milliseconds since epoch.
	CreatedAt *string `pulumi:"createdAt"`
	// Description of the API product. Include key information about the API product that is not captured by other fields. Comma-separated list of API resources to be bundled in the API product. By default, the resource paths are mapped from the `proxy.pathsuffix` variable. The proxy path suffix is defined as the URI fragment following the ProxyEndpoint base path. For example, if the `apiResources` element is defined to be `/forecastrss` and the base path defined for the API proxy is `/weather`, then only requests to `/weather/forecastrss` are permitted by the API product. You can select a specific path, or you can select all subpaths with the following wildcard: - `/**`: Indicates that all sub-URIs are included. - `/*` : Indicates that only URIs one level down are included. By default, / supports the same resources as /** as well as the base path defined by the API proxy. For example, if the base path of the API proxy is `/v1/weatherapikey`, then the API product supports requests to `/v1/weatherapikey` and to any sub-URIs, such as `/v1/weatherapikey/forecastrss`, `/v1/weatherapikey/region/CA`, and so on. For more information, see Managing API products.
	Description *string `pulumi:"description"`
	// Name displayed in the UI or developer portal to developers registering for API access.
	DisplayName *string `pulumi:"displayName"`
	// Comma-separated list of environment names to which the API product is bound. Requests to environments that are not listed are rejected. By specifying one or more environments, you can bind the resources listed in the API product to a specific environment, preventing developers from accessing those resources through API proxies deployed in another environment. This setting is used, for example, to prevent resources associated with API proxies in `prod` from being accessed by API proxies deployed in `test`.
	Environments []string `pulumi:"environments"`
	// Configuration used to group Apigee proxies or remote services with graphQL operation name, graphQL operation type and quotas. This grouping allows us to precisely set quota for a particular combination of graphQL name and operation type for a particular proxy request. If graphQL name is not set, this would imply quota will be applied on all graphQL requests matching the operation type.
	GraphqlOperationGroup *GoogleCloudApigeeV1GraphQLOperationGroupResponse `pulumi:"graphqlOperationGroup"`
	// Response only. Modified time of this environment as milliseconds since epoch.
	LastModifiedAt *string `pulumi:"lastModifiedAt"`
	// Internal name of the API product. Characters you can use in the name are restricted to: `A-Z0-9._\-$ %`. **Note:** The internal name cannot be edited when updating the API product.
	Name *string `pulumi:"name"`
	// Configuration used to group Apigee proxies or remote services with resources, method types, and quotas. The resource refers to the resource URI (excluding the base path). With this grouping, the API product creator is able to fine-tune and give precise control over which REST methods have access to specific resources and how many calls can be made (using the `quota` setting). **Note:** The `api_resources` setting cannot be specified for both the API product and operation group; otherwise the call will fail.
	OperationGroup *GoogleCloudApigeeV1OperationGroupResponse `pulumi:"operationGroup"`
	// Comma-separated list of API proxy names to which this API product is bound. By specifying API proxies, you can associate resources in the API product with specific API proxies, preventing developers from accessing those resources through other API proxies. Apigee rejects requests to API proxies that are not listed. **Note:** The API proxy names must already exist in the specified environment as they will be validated upon creation.
	Proxies []string `pulumi:"proxies"`
	// Number of request messages permitted per app by this API product for the specified `quotaInterval` and `quotaTimeUnit`. For example, a `quota` of 50, for a `quotaInterval` of 12 and a `quotaTimeUnit` of hours means 50 requests are allowed every 12 hours.
	Quota *string `pulumi:"quota"`
	// Time interval over which the number of request messages is calculated.
	QuotaInterval *string `pulumi:"quotaInterval"`
	// Time unit defined for the `quotaInterval`. Valid values include `minute`, `hour`, `day`, or `month`.
	QuotaTimeUnit *string `pulumi:"quotaTimeUnit"`
	// Comma-separated list of OAuth scopes that are validated at runtime. Apigee validates that the scopes in any access token presented match the scopes defined in the OAuth policy associated with the API product.
	Scopes []string `pulumi:"scopes"`
}

type OrganizationApiproductState struct {
	ApiResources pulumi.StringArrayInput
	// Flag that specifies how API keys are approved to access the APIs defined by the API product. If set to `manual`, the consumer key is generated and returned in "pending" state. In this case, the API keys won't work until they have been explicitly approved. If set to `auto`, the consumer key is generated and returned in "approved" state and can be used immediately. **Note:** Typically, `auto` is used to provide access to free or trial API products that provide limited quota or capabilities.
	ApprovalType pulumi.StringPtrInput
	// Array of attributes that may be used to extend the default API product profile with customer-specific metadata. You can specify a maximum of 18 attributes. Use this property to specify the access level of the API product as either `public`, `private`, or `internal`. Only products marked `public` are available to developers in the Apigee developer portal. For example, you can set a product to `internal` while it is in development and then change access to `public` when it is ready to release on the portal. API products marked as `private` do not appear on the portal, but can be accessed by external developers.
	Attributes GoogleCloudApigeeV1AttributeResponseArrayInput
	// Response only. Creation time of this environment as milliseconds since epoch.
	CreatedAt pulumi.StringPtrInput
	// Description of the API product. Include key information about the API product that is not captured by other fields. Comma-separated list of API resources to be bundled in the API product. By default, the resource paths are mapped from the `proxy.pathsuffix` variable. The proxy path suffix is defined as the URI fragment following the ProxyEndpoint base path. For example, if the `apiResources` element is defined to be `/forecastrss` and the base path defined for the API proxy is `/weather`, then only requests to `/weather/forecastrss` are permitted by the API product. You can select a specific path, or you can select all subpaths with the following wildcard: - `/**`: Indicates that all sub-URIs are included. - `/*` : Indicates that only URIs one level down are included. By default, / supports the same resources as /** as well as the base path defined by the API proxy. For example, if the base path of the API proxy is `/v1/weatherapikey`, then the API product supports requests to `/v1/weatherapikey` and to any sub-URIs, such as `/v1/weatherapikey/forecastrss`, `/v1/weatherapikey/region/CA`, and so on. For more information, see Managing API products.
	Description pulumi.StringPtrInput
	// Name displayed in the UI or developer portal to developers registering for API access.
	DisplayName pulumi.StringPtrInput
	// Comma-separated list of environment names to which the API product is bound. Requests to environments that are not listed are rejected. By specifying one or more environments, you can bind the resources listed in the API product to a specific environment, preventing developers from accessing those resources through API proxies deployed in another environment. This setting is used, for example, to prevent resources associated with API proxies in `prod` from being accessed by API proxies deployed in `test`.
	Environments pulumi.StringArrayInput
	// Configuration used to group Apigee proxies or remote services with graphQL operation name, graphQL operation type and quotas. This grouping allows us to precisely set quota for a particular combination of graphQL name and operation type for a particular proxy request. If graphQL name is not set, this would imply quota will be applied on all graphQL requests matching the operation type.
	GraphqlOperationGroup GoogleCloudApigeeV1GraphQLOperationGroupResponsePtrInput
	// Response only. Modified time of this environment as milliseconds since epoch.
	LastModifiedAt pulumi.StringPtrInput
	// Internal name of the API product. Characters you can use in the name are restricted to: `A-Z0-9._\-$ %`. **Note:** The internal name cannot be edited when updating the API product.
	Name pulumi.StringPtrInput
	// Configuration used to group Apigee proxies or remote services with resources, method types, and quotas. The resource refers to the resource URI (excluding the base path). With this grouping, the API product creator is able to fine-tune and give precise control over which REST methods have access to specific resources and how many calls can be made (using the `quota` setting). **Note:** The `api_resources` setting cannot be specified for both the API product and operation group; otherwise the call will fail.
	OperationGroup GoogleCloudApigeeV1OperationGroupResponsePtrInput
	// Comma-separated list of API proxy names to which this API product is bound. By specifying API proxies, you can associate resources in the API product with specific API proxies, preventing developers from accessing those resources through other API proxies. Apigee rejects requests to API proxies that are not listed. **Note:** The API proxy names must already exist in the specified environment as they will be validated upon creation.
	Proxies pulumi.StringArrayInput
	// Number of request messages permitted per app by this API product for the specified `quotaInterval` and `quotaTimeUnit`. For example, a `quota` of 50, for a `quotaInterval` of 12 and a `quotaTimeUnit` of hours means 50 requests are allowed every 12 hours.
	Quota pulumi.StringPtrInput
	// Time interval over which the number of request messages is calculated.
	QuotaInterval pulumi.StringPtrInput
	// Time unit defined for the `quotaInterval`. Valid values include `minute`, `hour`, `day`, or `month`.
	QuotaTimeUnit pulumi.StringPtrInput
	// Comma-separated list of OAuth scopes that are validated at runtime. Apigee validates that the scopes in any access token presented match the scopes defined in the OAuth policy associated with the API product.
	Scopes pulumi.StringArrayInput
}

func (OrganizationApiproductState) ElementType() reflect.Type {
	return reflect.TypeOf((*organizationApiproductState)(nil)).Elem()
}

type organizationApiproductArgs struct {
	ApiResources []string `pulumi:"apiResources"`
	// Flag that specifies how API keys are approved to access the APIs defined by the API product. If set to `manual`, the consumer key is generated and returned in "pending" state. In this case, the API keys won't work until they have been explicitly approved. If set to `auto`, the consumer key is generated and returned in "approved" state and can be used immediately. **Note:** Typically, `auto` is used to provide access to free or trial API products that provide limited quota or capabilities.
	ApprovalType *string `pulumi:"approvalType"`
	// Array of attributes that may be used to extend the default API product profile with customer-specific metadata. You can specify a maximum of 18 attributes. Use this property to specify the access level of the API product as either `public`, `private`, or `internal`. Only products marked `public` are available to developers in the Apigee developer portal. For example, you can set a product to `internal` while it is in development and then change access to `public` when it is ready to release on the portal. API products marked as `private` do not appear on the portal, but can be accessed by external developers.
	Attributes []GoogleCloudApigeeV1Attribute `pulumi:"attributes"`
	// Response only. Creation time of this environment as milliseconds since epoch.
	CreatedAt *string `pulumi:"createdAt"`
	// Description of the API product. Include key information about the API product that is not captured by other fields. Comma-separated list of API resources to be bundled in the API product. By default, the resource paths are mapped from the `proxy.pathsuffix` variable. The proxy path suffix is defined as the URI fragment following the ProxyEndpoint base path. For example, if the `apiResources` element is defined to be `/forecastrss` and the base path defined for the API proxy is `/weather`, then only requests to `/weather/forecastrss` are permitted by the API product. You can select a specific path, or you can select all subpaths with the following wildcard: - `/**`: Indicates that all sub-URIs are included. - `/*` : Indicates that only URIs one level down are included. By default, / supports the same resources as /** as well as the base path defined by the API proxy. For example, if the base path of the API proxy is `/v1/weatherapikey`, then the API product supports requests to `/v1/weatherapikey` and to any sub-URIs, such as `/v1/weatherapikey/forecastrss`, `/v1/weatherapikey/region/CA`, and so on. For more information, see Managing API products.
	Description *string `pulumi:"description"`
	// Name displayed in the UI or developer portal to developers registering for API access.
	DisplayName *string `pulumi:"displayName"`
	// Comma-separated list of environment names to which the API product is bound. Requests to environments that are not listed are rejected. By specifying one or more environments, you can bind the resources listed in the API product to a specific environment, preventing developers from accessing those resources through API proxies deployed in another environment. This setting is used, for example, to prevent resources associated with API proxies in `prod` from being accessed by API proxies deployed in `test`.
	Environments []string `pulumi:"environments"`
	// Configuration used to group Apigee proxies or remote services with graphQL operation name, graphQL operation type and quotas. This grouping allows us to precisely set quota for a particular combination of graphQL name and operation type for a particular proxy request. If graphQL name is not set, this would imply quota will be applied on all graphQL requests matching the operation type.
	GraphqlOperationGroup *GoogleCloudApigeeV1GraphQLOperationGroup `pulumi:"graphqlOperationGroup"`
	// Response only. Modified time of this environment as milliseconds since epoch.
	LastModifiedAt *string `pulumi:"lastModifiedAt"`
	// Internal name of the API product. Characters you can use in the name are restricted to: `A-Z0-9._\-$ %`. **Note:** The internal name cannot be edited when updating the API product.
	Name *string `pulumi:"name"`
	// Configuration used to group Apigee proxies or remote services with resources, method types, and quotas. The resource refers to the resource URI (excluding the base path). With this grouping, the API product creator is able to fine-tune and give precise control over which REST methods have access to specific resources and how many calls can be made (using the `quota` setting). **Note:** The `api_resources` setting cannot be specified for both the API product and operation group; otherwise the call will fail.
	OperationGroup *GoogleCloudApigeeV1OperationGroup `pulumi:"operationGroup"`
	OrganizationId string                             `pulumi:"organizationId"`
	// Comma-separated list of API proxy names to which this API product is bound. By specifying API proxies, you can associate resources in the API product with specific API proxies, preventing developers from accessing those resources through other API proxies. Apigee rejects requests to API proxies that are not listed. **Note:** The API proxy names must already exist in the specified environment as they will be validated upon creation.
	Proxies []string `pulumi:"proxies"`
	// Number of request messages permitted per app by this API product for the specified `quotaInterval` and `quotaTimeUnit`. For example, a `quota` of 50, for a `quotaInterval` of 12 and a `quotaTimeUnit` of hours means 50 requests are allowed every 12 hours.
	Quota *string `pulumi:"quota"`
	// Time interval over which the number of request messages is calculated.
	QuotaInterval *string `pulumi:"quotaInterval"`
	// Time unit defined for the `quotaInterval`. Valid values include `minute`, `hour`, `day`, or `month`.
	QuotaTimeUnit *string `pulumi:"quotaTimeUnit"`
	// Comma-separated list of OAuth scopes that are validated at runtime. Apigee validates that the scopes in any access token presented match the scopes defined in the OAuth policy associated with the API product.
	Scopes []string `pulumi:"scopes"`
}

// The set of arguments for constructing a OrganizationApiproduct resource.
type OrganizationApiproductArgs struct {
	ApiResources pulumi.StringArrayInput
	// Flag that specifies how API keys are approved to access the APIs defined by the API product. If set to `manual`, the consumer key is generated and returned in "pending" state. In this case, the API keys won't work until they have been explicitly approved. If set to `auto`, the consumer key is generated and returned in "approved" state and can be used immediately. **Note:** Typically, `auto` is used to provide access to free or trial API products that provide limited quota or capabilities.
	ApprovalType pulumi.StringPtrInput
	// Array of attributes that may be used to extend the default API product profile with customer-specific metadata. You can specify a maximum of 18 attributes. Use this property to specify the access level of the API product as either `public`, `private`, or `internal`. Only products marked `public` are available to developers in the Apigee developer portal. For example, you can set a product to `internal` while it is in development and then change access to `public` when it is ready to release on the portal. API products marked as `private` do not appear on the portal, but can be accessed by external developers.
	Attributes GoogleCloudApigeeV1AttributeArrayInput
	// Response only. Creation time of this environment as milliseconds since epoch.
	CreatedAt pulumi.StringPtrInput
	// Description of the API product. Include key information about the API product that is not captured by other fields. Comma-separated list of API resources to be bundled in the API product. By default, the resource paths are mapped from the `proxy.pathsuffix` variable. The proxy path suffix is defined as the URI fragment following the ProxyEndpoint base path. For example, if the `apiResources` element is defined to be `/forecastrss` and the base path defined for the API proxy is `/weather`, then only requests to `/weather/forecastrss` are permitted by the API product. You can select a specific path, or you can select all subpaths with the following wildcard: - `/**`: Indicates that all sub-URIs are included. - `/*` : Indicates that only URIs one level down are included. By default, / supports the same resources as /** as well as the base path defined by the API proxy. For example, if the base path of the API proxy is `/v1/weatherapikey`, then the API product supports requests to `/v1/weatherapikey` and to any sub-URIs, such as `/v1/weatherapikey/forecastrss`, `/v1/weatherapikey/region/CA`, and so on. For more information, see Managing API products.
	Description pulumi.StringPtrInput
	// Name displayed in the UI or developer portal to developers registering for API access.
	DisplayName pulumi.StringPtrInput
	// Comma-separated list of environment names to which the API product is bound. Requests to environments that are not listed are rejected. By specifying one or more environments, you can bind the resources listed in the API product to a specific environment, preventing developers from accessing those resources through API proxies deployed in another environment. This setting is used, for example, to prevent resources associated with API proxies in `prod` from being accessed by API proxies deployed in `test`.
	Environments pulumi.StringArrayInput
	// Configuration used to group Apigee proxies or remote services with graphQL operation name, graphQL operation type and quotas. This grouping allows us to precisely set quota for a particular combination of graphQL name and operation type for a particular proxy request. If graphQL name is not set, this would imply quota will be applied on all graphQL requests matching the operation type.
	GraphqlOperationGroup GoogleCloudApigeeV1GraphQLOperationGroupPtrInput
	// Response only. Modified time of this environment as milliseconds since epoch.
	LastModifiedAt pulumi.StringPtrInput
	// Internal name of the API product. Characters you can use in the name are restricted to: `A-Z0-9._\-$ %`. **Note:** The internal name cannot be edited when updating the API product.
	Name pulumi.StringPtrInput
	// Configuration used to group Apigee proxies or remote services with resources, method types, and quotas. The resource refers to the resource URI (excluding the base path). With this grouping, the API product creator is able to fine-tune and give precise control over which REST methods have access to specific resources and how many calls can be made (using the `quota` setting). **Note:** The `api_resources` setting cannot be specified for both the API product and operation group; otherwise the call will fail.
	OperationGroup GoogleCloudApigeeV1OperationGroupPtrInput
	OrganizationId pulumi.StringInput
	// Comma-separated list of API proxy names to which this API product is bound. By specifying API proxies, you can associate resources in the API product with specific API proxies, preventing developers from accessing those resources through other API proxies. Apigee rejects requests to API proxies that are not listed. **Note:** The API proxy names must already exist in the specified environment as they will be validated upon creation.
	Proxies pulumi.StringArrayInput
	// Number of request messages permitted per app by this API product for the specified `quotaInterval` and `quotaTimeUnit`. For example, a `quota` of 50, for a `quotaInterval` of 12 and a `quotaTimeUnit` of hours means 50 requests are allowed every 12 hours.
	Quota pulumi.StringPtrInput
	// Time interval over which the number of request messages is calculated.
	QuotaInterval pulumi.StringPtrInput
	// Time unit defined for the `quotaInterval`. Valid values include `minute`, `hour`, `day`, or `month`.
	QuotaTimeUnit pulumi.StringPtrInput
	// Comma-separated list of OAuth scopes that are validated at runtime. Apigee validates that the scopes in any access token presented match the scopes defined in the OAuth policy associated with the API product.
	Scopes pulumi.StringArrayInput
}

func (OrganizationApiproductArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*organizationApiproductArgs)(nil)).Elem()
}

type OrganizationApiproductInput interface {
	pulumi.Input

	ToOrganizationApiproductOutput() OrganizationApiproductOutput
	ToOrganizationApiproductOutputWithContext(ctx context.Context) OrganizationApiproductOutput
}

func (*OrganizationApiproduct) ElementType() reflect.Type {
	return reflect.TypeOf((*OrganizationApiproduct)(nil))
}

func (i *OrganizationApiproduct) ToOrganizationApiproductOutput() OrganizationApiproductOutput {
	return i.ToOrganizationApiproductOutputWithContext(context.Background())
}

func (i *OrganizationApiproduct) ToOrganizationApiproductOutputWithContext(ctx context.Context) OrganizationApiproductOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrganizationApiproductOutput)
}

type OrganizationApiproductOutput struct {
	*pulumi.OutputState
}

func (OrganizationApiproductOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OrganizationApiproduct)(nil))
}

func (o OrganizationApiproductOutput) ToOrganizationApiproductOutput() OrganizationApiproductOutput {
	return o
}

func (o OrganizationApiproductOutput) ToOrganizationApiproductOutputWithContext(ctx context.Context) OrganizationApiproductOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(OrganizationApiproductOutput{})
}
