// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates an API product in an organization. You create API products after you have proxied backend services using API proxies. An API product is a collection of API resources combined with quota settings and metadata that you can use to deliver customized and productized API bundles to your developer community. This metadata can include: - Scope - Environments - API proxies - Extensible profile API products enable you repackage APIs on-the-fly, without having to do any additional coding or configuration. Apigee recommends that you start with a simple API product including only required elements. You then provision credentials to apps to enable them to start testing your APIs. After you have authentication and authorization working against a simple API product, you can iterate to create finer grained API products, defining different sets of API resources for each API product. **WARNING:** - If you don't specify an API proxy in the request body, *any* app associated with the product can make calls to *any* API in your entire organization. - If you don't specify an environment in the request body, the product allows access to all environments. For more information, see What is an API product?
 */
export class OrganizationApiproduct extends pulumi.CustomResource {
    /**
     * Get an existing OrganizationApiproduct resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): OrganizationApiproduct {
        return new OrganizationApiproduct(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:apigee/v1:OrganizationApiproduct';

    /**
     * Returns true if the given object is an instance of OrganizationApiproduct.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is OrganizationApiproduct {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === OrganizationApiproduct.__pulumiType;
    }

    public readonly apiResources!: pulumi.Output<string[]>;
    /**
     * Flag that specifies how API keys are approved to access the APIs defined by the API product. If set to `manual`, the consumer key is generated and returned in "pending" state. In this case, the API keys won't work until they have been explicitly approved. If set to `auto`, the consumer key is generated and returned in "approved" state and can be used immediately. **Note:** Typically, `auto` is used to provide access to free or trial API products that provide limited quota or capabilities.
     */
    public readonly approvalType!: pulumi.Output<string>;
    /**
     * Array of attributes that may be used to extend the default API product profile with customer-specific metadata. You can specify a maximum of 18 attributes. Use this property to specify the access level of the API product as either `public`, `private`, or `internal`. Only products marked `public` are available to developers in the Apigee developer portal. For example, you can set a product to `internal` while it is in development and then change access to `public` when it is ready to release on the portal. API products marked as `private` do not appear on the portal, but can be accessed by external developers.
     */
    public readonly attributes!: pulumi.Output<outputs.apigee.v1.GoogleCloudApigeeV1AttributeResponse[]>;
    /**
     * Response only. Creation time of this environment as milliseconds since epoch.
     */
    public readonly createdAt!: pulumi.Output<string>;
    /**
     * Description of the API product. Include key information about the API product that is not captured by other fields. Comma-separated list of API resources to be bundled in the API product. By default, the resource paths are mapped from the `proxy.pathsuffix` variable. The proxy path suffix is defined as the URI fragment following the ProxyEndpoint base path. For example, if the `apiResources` element is defined to be `/forecastrss` and the base path defined for the API proxy is `/weather`, then only requests to `/weather/forecastrss` are permitted by the API product. You can select a specific path, or you can select all subpaths with the following wildcard: - `/**`: Indicates that all sub-URIs are included. - `/*` : Indicates that only URIs one level down are included. By default, / supports the same resources as /** as well as the base path defined by the API proxy. For example, if the base path of the API proxy is `/v1/weatherapikey`, then the API product supports requests to `/v1/weatherapikey` and to any sub-URIs, such as `/v1/weatherapikey/forecastrss`, `/v1/weatherapikey/region/CA`, and so on. For more information, see Managing API products.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * Name displayed in the UI or developer portal to developers registering for API access.
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * Comma-separated list of environment names to which the API product is bound. Requests to environments that are not listed are rejected. By specifying one or more environments, you can bind the resources listed in the API product to a specific environment, preventing developers from accessing those resources through API proxies deployed in another environment. This setting is used, for example, to prevent resources associated with API proxies in `prod` from being accessed by API proxies deployed in `test`.
     */
    public readonly environments!: pulumi.Output<string[]>;
    /**
     * Configuration used to group Apigee proxies or remote services with graphQL operation name, graphQL operation type and quotas. This grouping allows us to precisely set quota for a particular combination of graphQL name and operation type for a particular proxy request. If graphQL name is not set, this would imply quota will be applied on all graphQL requests matching the operation type.
     */
    public readonly graphqlOperationGroup!: pulumi.Output<outputs.apigee.v1.GoogleCloudApigeeV1GraphQLOperationGroupResponse>;
    /**
     * Response only. Modified time of this environment as milliseconds since epoch.
     */
    public readonly lastModifiedAt!: pulumi.Output<string>;
    /**
     * Internal name of the API product. Characters you can use in the name are restricted to: `A-Z0-9._\-$ %`. **Note:** The internal name cannot be edited when updating the API product.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Configuration used to group Apigee proxies or remote services with resources, method types, and quotas. The resource refers to the resource URI (excluding the base path). With this grouping, the API product creator is able to fine-tune and give precise control over which REST methods have access to specific resources and how many calls can be made (using the `quota` setting). **Note:** The `api_resources` setting cannot be specified for both the API product and operation group; otherwise the call will fail.
     */
    public readonly operationGroup!: pulumi.Output<outputs.apigee.v1.GoogleCloudApigeeV1OperationGroupResponse>;
    /**
     * Comma-separated list of API proxy names to which this API product is bound. By specifying API proxies, you can associate resources in the API product with specific API proxies, preventing developers from accessing those resources through other API proxies. Apigee rejects requests to API proxies that are not listed. **Note:** The API proxy names must already exist in the specified environment as they will be validated upon creation.
     */
    public readonly proxies!: pulumi.Output<string[]>;
    /**
     * Number of request messages permitted per app by this API product for the specified `quotaInterval` and `quotaTimeUnit`. For example, a `quota` of 50, for a `quotaInterval` of 12 and a `quotaTimeUnit` of hours means 50 requests are allowed every 12 hours.
     */
    public readonly quota!: pulumi.Output<string>;
    /**
     * Time interval over which the number of request messages is calculated.
     */
    public readonly quotaInterval!: pulumi.Output<string>;
    /**
     * Time unit defined for the `quotaInterval`. Valid values include `minute`, `hour`, `day`, or `month`.
     */
    public readonly quotaTimeUnit!: pulumi.Output<string>;
    /**
     * Comma-separated list of OAuth scopes that are validated at runtime. Apigee validates that the scopes in any access token presented match the scopes defined in the OAuth policy associated with the API product.
     */
    public readonly scopes!: pulumi.Output<string[]>;

    /**
     * Create a OrganizationApiproduct resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: OrganizationApiproductArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.organizationId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'organizationId'");
            }
            inputs["apiResources"] = args ? args.apiResources : undefined;
            inputs["approvalType"] = args ? args.approvalType : undefined;
            inputs["attributes"] = args ? args.attributes : undefined;
            inputs["createdAt"] = args ? args.createdAt : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["displayName"] = args ? args.displayName : undefined;
            inputs["environments"] = args ? args.environments : undefined;
            inputs["graphqlOperationGroup"] = args ? args.graphqlOperationGroup : undefined;
            inputs["lastModifiedAt"] = args ? args.lastModifiedAt : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["operationGroup"] = args ? args.operationGroup : undefined;
            inputs["organizationId"] = args ? args.organizationId : undefined;
            inputs["proxies"] = args ? args.proxies : undefined;
            inputs["quota"] = args ? args.quota : undefined;
            inputs["quotaInterval"] = args ? args.quotaInterval : undefined;
            inputs["quotaTimeUnit"] = args ? args.quotaTimeUnit : undefined;
            inputs["scopes"] = args ? args.scopes : undefined;
        } else {
            inputs["apiResources"] = undefined /*out*/;
            inputs["approvalType"] = undefined /*out*/;
            inputs["attributes"] = undefined /*out*/;
            inputs["createdAt"] = undefined /*out*/;
            inputs["description"] = undefined /*out*/;
            inputs["displayName"] = undefined /*out*/;
            inputs["environments"] = undefined /*out*/;
            inputs["graphqlOperationGroup"] = undefined /*out*/;
            inputs["lastModifiedAt"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["operationGroup"] = undefined /*out*/;
            inputs["proxies"] = undefined /*out*/;
            inputs["quota"] = undefined /*out*/;
            inputs["quotaInterval"] = undefined /*out*/;
            inputs["quotaTimeUnit"] = undefined /*out*/;
            inputs["scopes"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(OrganizationApiproduct.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a OrganizationApiproduct resource.
 */
export interface OrganizationApiproductArgs {
    readonly apiResources?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Flag that specifies how API keys are approved to access the APIs defined by the API product. If set to `manual`, the consumer key is generated and returned in "pending" state. In this case, the API keys won't work until they have been explicitly approved. If set to `auto`, the consumer key is generated and returned in "approved" state and can be used immediately. **Note:** Typically, `auto` is used to provide access to free or trial API products that provide limited quota or capabilities.
     */
    readonly approvalType?: pulumi.Input<string>;
    /**
     * Array of attributes that may be used to extend the default API product profile with customer-specific metadata. You can specify a maximum of 18 attributes. Use this property to specify the access level of the API product as either `public`, `private`, or `internal`. Only products marked `public` are available to developers in the Apigee developer portal. For example, you can set a product to `internal` while it is in development and then change access to `public` when it is ready to release on the portal. API products marked as `private` do not appear on the portal, but can be accessed by external developers.
     */
    readonly attributes?: pulumi.Input<pulumi.Input<inputs.apigee.v1.GoogleCloudApigeeV1AttributeArgs>[]>;
    /**
     * Response only. Creation time of this environment as milliseconds since epoch.
     */
    readonly createdAt?: pulumi.Input<string>;
    /**
     * Description of the API product. Include key information about the API product that is not captured by other fields. Comma-separated list of API resources to be bundled in the API product. By default, the resource paths are mapped from the `proxy.pathsuffix` variable. The proxy path suffix is defined as the URI fragment following the ProxyEndpoint base path. For example, if the `apiResources` element is defined to be `/forecastrss` and the base path defined for the API proxy is `/weather`, then only requests to `/weather/forecastrss` are permitted by the API product. You can select a specific path, or you can select all subpaths with the following wildcard: - `/**`: Indicates that all sub-URIs are included. - `/*` : Indicates that only URIs one level down are included. By default, / supports the same resources as /** as well as the base path defined by the API proxy. For example, if the base path of the API proxy is `/v1/weatherapikey`, then the API product supports requests to `/v1/weatherapikey` and to any sub-URIs, such as `/v1/weatherapikey/forecastrss`, `/v1/weatherapikey/region/CA`, and so on. For more information, see Managing API products.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Name displayed in the UI or developer portal to developers registering for API access.
     */
    readonly displayName?: pulumi.Input<string>;
    /**
     * Comma-separated list of environment names to which the API product is bound. Requests to environments that are not listed are rejected. By specifying one or more environments, you can bind the resources listed in the API product to a specific environment, preventing developers from accessing those resources through API proxies deployed in another environment. This setting is used, for example, to prevent resources associated with API proxies in `prod` from being accessed by API proxies deployed in `test`.
     */
    readonly environments?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Configuration used to group Apigee proxies or remote services with graphQL operation name, graphQL operation type and quotas. This grouping allows us to precisely set quota for a particular combination of graphQL name and operation type for a particular proxy request. If graphQL name is not set, this would imply quota will be applied on all graphQL requests matching the operation type.
     */
    readonly graphqlOperationGroup?: pulumi.Input<inputs.apigee.v1.GoogleCloudApigeeV1GraphQLOperationGroupArgs>;
    /**
     * Response only. Modified time of this environment as milliseconds since epoch.
     */
    readonly lastModifiedAt?: pulumi.Input<string>;
    /**
     * Internal name of the API product. Characters you can use in the name are restricted to: `A-Z0-9._\-$ %`. **Note:** The internal name cannot be edited when updating the API product.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Configuration used to group Apigee proxies or remote services with resources, method types, and quotas. The resource refers to the resource URI (excluding the base path). With this grouping, the API product creator is able to fine-tune and give precise control over which REST methods have access to specific resources and how many calls can be made (using the `quota` setting). **Note:** The `api_resources` setting cannot be specified for both the API product and operation group; otherwise the call will fail.
     */
    readonly operationGroup?: pulumi.Input<inputs.apigee.v1.GoogleCloudApigeeV1OperationGroupArgs>;
    readonly organizationId: pulumi.Input<string>;
    /**
     * Comma-separated list of API proxy names to which this API product is bound. By specifying API proxies, you can associate resources in the API product with specific API proxies, preventing developers from accessing those resources through other API proxies. Apigee rejects requests to API proxies that are not listed. **Note:** The API proxy names must already exist in the specified environment as they will be validated upon creation.
     */
    readonly proxies?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Number of request messages permitted per app by this API product for the specified `quotaInterval` and `quotaTimeUnit`. For example, a `quota` of 50, for a `quotaInterval` of 12 and a `quotaTimeUnit` of hours means 50 requests are allowed every 12 hours.
     */
    readonly quota?: pulumi.Input<string>;
    /**
     * Time interval over which the number of request messages is calculated.
     */
    readonly quotaInterval?: pulumi.Input<string>;
    /**
     * Time unit defined for the `quotaInterval`. Valid values include `minute`, `hour`, `day`, or `month`.
     */
    readonly quotaTimeUnit?: pulumi.Input<string>;
    /**
     * Comma-separated list of OAuth scopes that are validated at runtime. Apigee validates that the scopes in any access token presented match the scopes defined in the OAuth policy associated with the API product.
     */
    readonly scopes?: pulumi.Input<pulumi.Input<string>[]>;
}
