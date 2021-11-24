// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates an app associated with a developer. This API associates the developer app with the specified API product and auto-generates an API key for the app to use in calls to API proxies inside that API product. The `name` is the unique ID of the app that you can use in API calls. The `DisplayName` (set as an attribute) appears in the UI. If you don't set the `DisplayName` attribute, the `name` appears in the UI.
 */
export class App extends pulumi.CustomResource {
    /**
     * Get an existing App resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): App {
        return new App(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:apigee/v1:App';

    /**
     * Returns true if the given object is an instance of App.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is App {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === App.__pulumiType;
    }

    /**
     * List of API products associated with the developer app.
     */
    public readonly apiProducts!: pulumi.Output<string[]>;
    /**
     * Developer app family.
     */
    public readonly appFamily!: pulumi.Output<string>;
    /**
     * ID of the developer app.
     */
    public readonly appId!: pulumi.Output<string>;
    /**
     * List of attributes for the developer app.
     */
    public readonly attributes!: pulumi.Output<outputs.apigee.v1.GoogleCloudApigeeV1AttributeResponse[]>;
    /**
     * Callback URL used by OAuth 2.0 authorization servers to communicate authorization codes back to developer apps.
     */
    public readonly callbackUrl!: pulumi.Output<string>;
    /**
     * Time the developer app was created in milliseconds since epoch.
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * Set of credentials for the developer app consisting of the consumer key/secret pairs associated with the API products.
     */
    public /*out*/ readonly credentials!: pulumi.Output<outputs.apigee.v1.GoogleCloudApigeeV1CredentialResponse[]>;
    /**
     * ID of the developer.
     */
    public readonly developerId!: pulumi.Output<string>;
    /**
     * Expiration time, in milliseconds, for the consumer key that is generated for the developer app. If not set or left to the default value of `-1`, the API key never expires. The expiration time can't be updated after it is set.
     */
    public readonly keyExpiresIn!: pulumi.Output<string>;
    /**
     * Time the developer app was modified in milliseconds since epoch.
     */
    public /*out*/ readonly lastModifiedAt!: pulumi.Output<string>;
    /**
     * Name of the developer app.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Scopes to apply to the developer app. The specified scopes must already exist for the API product that you associate with the developer app.
     */
    public readonly scopes!: pulumi.Output<string[]>;
    /**
     * Status of the credential. Valid values include `approved` or `revoked`.
     */
    public readonly status!: pulumi.Output<string>;

    /**
     * Create a App resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AppArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.developerId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'developerId'");
            }
            if ((!args || args.organizationId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'organizationId'");
            }
            resourceInputs["apiProducts"] = args ? args.apiProducts : undefined;
            resourceInputs["appFamily"] = args ? args.appFamily : undefined;
            resourceInputs["appId"] = args ? args.appId : undefined;
            resourceInputs["attributes"] = args ? args.attributes : undefined;
            resourceInputs["callbackUrl"] = args ? args.callbackUrl : undefined;
            resourceInputs["developerId"] = args ? args.developerId : undefined;
            resourceInputs["keyExpiresIn"] = args ? args.keyExpiresIn : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["organizationId"] = args ? args.organizationId : undefined;
            resourceInputs["scopes"] = args ? args.scopes : undefined;
            resourceInputs["status"] = args ? args.status : undefined;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["credentials"] = undefined /*out*/;
            resourceInputs["lastModifiedAt"] = undefined /*out*/;
        } else {
            resourceInputs["apiProducts"] = undefined /*out*/;
            resourceInputs["appFamily"] = undefined /*out*/;
            resourceInputs["appId"] = undefined /*out*/;
            resourceInputs["attributes"] = undefined /*out*/;
            resourceInputs["callbackUrl"] = undefined /*out*/;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["credentials"] = undefined /*out*/;
            resourceInputs["developerId"] = undefined /*out*/;
            resourceInputs["keyExpiresIn"] = undefined /*out*/;
            resourceInputs["lastModifiedAt"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["scopes"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(App.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a App resource.
 */
export interface AppArgs {
    /**
     * List of API products associated with the developer app.
     */
    apiProducts?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Developer app family.
     */
    appFamily?: pulumi.Input<string>;
    /**
     * ID of the developer app.
     */
    appId?: pulumi.Input<string>;
    /**
     * List of attributes for the developer app.
     */
    attributes?: pulumi.Input<pulumi.Input<inputs.apigee.v1.GoogleCloudApigeeV1AttributeArgs>[]>;
    /**
     * Callback URL used by OAuth 2.0 authorization servers to communicate authorization codes back to developer apps.
     */
    callbackUrl?: pulumi.Input<string>;
    /**
     * ID of the developer.
     */
    developerId: pulumi.Input<string>;
    /**
     * Expiration time, in milliseconds, for the consumer key that is generated for the developer app. If not set or left to the default value of `-1`, the API key never expires. The expiration time can't be updated after it is set.
     */
    keyExpiresIn?: pulumi.Input<string>;
    /**
     * Name of the developer app.
     */
    name?: pulumi.Input<string>;
    organizationId: pulumi.Input<string>;
    /**
     * Scopes to apply to the developer app. The specified scopes must already exist for the API product that you associate with the developer app.
     */
    scopes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Status of the credential. Valid values include `approved` or `revoked`.
     */
    status?: pulumi.Input<string>;
}
