// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates an App Engine application for a Google Cloud Platform project. Required fields: id - The ID of the target Cloud Platform project. location - The region (https://cloud.google.com/appengine/docs/locations) where you want the App Engine application located.For more information about App Engine applications, see Managing Projects, Applications, and Billing (https://cloud.google.com/appengine/docs/standard/python/console/).
 * Auto-naming is currently not supported for this resource.
 * Note - this resource's API doesn't support deletion. When deleted, the resource will persist
 * on Google Cloud even though it will be deleted from Pulumi state.
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
    public static readonly __pulumiType = 'google-native:appengine/v1:App';

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
     * Google Apps authentication domain that controls which users can access this application.Defaults to open access for any Google Account.
     */
    public readonly authDomain!: pulumi.Output<string>;
    /**
     * Google Cloud Storage bucket that can be used for storing files associated with this application. This bucket is associated with the application and can be used by the gcloud deployment commands.
     */
    public /*out*/ readonly codeBucket!: pulumi.Output<string>;
    /**
     * The type of the Cloud Firestore or Cloud Datastore database associated with this application.
     */
    public readonly databaseType!: pulumi.Output<string>;
    /**
     * Google Cloud Storage bucket that can be used by this application to store content.
     */
    public /*out*/ readonly defaultBucket!: pulumi.Output<string>;
    /**
     * Cookie expiration policy for this application.
     */
    public readonly defaultCookieExpiration!: pulumi.Output<string>;
    /**
     * Hostname used to reach this application, as resolved by App Engine.
     */
    public /*out*/ readonly defaultHostname!: pulumi.Output<string>;
    /**
     * HTTP path dispatch rules for requests to the application that do not explicitly target a service or version. Rules are order-dependent. Up to 20 dispatch rules can be supported.
     */
    public readonly dispatchRules!: pulumi.Output<outputs.appengine.v1.UrlDispatchRuleResponse[]>;
    /**
     * The feature specific settings to be used in the application.
     */
    public readonly featureSettings!: pulumi.Output<outputs.appengine.v1.FeatureSettingsResponse>;
    /**
     * The Google Container Registry domain used for storing managed build docker images for this application.
     */
    public readonly gcrDomain!: pulumi.Output<string>;
    public readonly iap!: pulumi.Output<outputs.appengine.v1.IdentityAwareProxyResponse>;
    /**
     * Location from which this application runs. Application instances run out of the data centers in the specified location, which is also where all of the application's end user content is stored.Defaults to us-central.View the list of supported locations (https://cloud.google.com/appengine/docs/locations).
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * Full path to the Application resource in the API. Example: apps/myapp.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The service account associated with the application. This is the app-level default identity. If no identity provided during create version, Admin API will fallback to this one.
     */
    public readonly serviceAccount!: pulumi.Output<string>;
    /**
     * Serving status of this application.
     */
    public readonly servingStatus!: pulumi.Output<string>;

    /**
     * Create a App resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: AppArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            resourceInputs["authDomain"] = args ? args.authDomain : undefined;
            resourceInputs["databaseType"] = args ? args.databaseType : undefined;
            resourceInputs["defaultCookieExpiration"] = args ? args.defaultCookieExpiration : undefined;
            resourceInputs["dispatchRules"] = args ? args.dispatchRules : undefined;
            resourceInputs["featureSettings"] = args ? args.featureSettings : undefined;
            resourceInputs["gcrDomain"] = args ? args.gcrDomain : undefined;
            resourceInputs["iap"] = args ? args.iap : undefined;
            resourceInputs["id"] = args ? args.id : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["serviceAccount"] = args ? args.serviceAccount : undefined;
            resourceInputs["servingStatus"] = args ? args.servingStatus : undefined;
            resourceInputs["codeBucket"] = undefined /*out*/;
            resourceInputs["defaultBucket"] = undefined /*out*/;
            resourceInputs["defaultHostname"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
        } else {
            resourceInputs["authDomain"] = undefined /*out*/;
            resourceInputs["codeBucket"] = undefined /*out*/;
            resourceInputs["databaseType"] = undefined /*out*/;
            resourceInputs["defaultBucket"] = undefined /*out*/;
            resourceInputs["defaultCookieExpiration"] = undefined /*out*/;
            resourceInputs["defaultHostname"] = undefined /*out*/;
            resourceInputs["dispatchRules"] = undefined /*out*/;
            resourceInputs["featureSettings"] = undefined /*out*/;
            resourceInputs["gcrDomain"] = undefined /*out*/;
            resourceInputs["iap"] = undefined /*out*/;
            resourceInputs["location"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["serviceAccount"] = undefined /*out*/;
            resourceInputs["servingStatus"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(App.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a App resource.
 */
export interface AppArgs {
    /**
     * Google Apps authentication domain that controls which users can access this application.Defaults to open access for any Google Account.
     */
    authDomain?: pulumi.Input<string>;
    /**
     * The type of the Cloud Firestore or Cloud Datastore database associated with this application.
     */
    databaseType?: pulumi.Input<enums.appengine.v1.AppDatabaseType>;
    /**
     * Cookie expiration policy for this application.
     */
    defaultCookieExpiration?: pulumi.Input<string>;
    /**
     * HTTP path dispatch rules for requests to the application that do not explicitly target a service or version. Rules are order-dependent. Up to 20 dispatch rules can be supported.
     */
    dispatchRules?: pulumi.Input<pulumi.Input<inputs.appengine.v1.UrlDispatchRuleArgs>[]>;
    /**
     * The feature specific settings to be used in the application.
     */
    featureSettings?: pulumi.Input<inputs.appengine.v1.FeatureSettingsArgs>;
    /**
     * The Google Container Registry domain used for storing managed build docker images for this application.
     */
    gcrDomain?: pulumi.Input<string>;
    iap?: pulumi.Input<inputs.appengine.v1.IdentityAwareProxyArgs>;
    /**
     * Identifier of the Application resource. This identifier is equivalent to the project ID of the Google Cloud Platform project where you want to deploy your application. Example: myapp.
     */
    id?: pulumi.Input<string>;
    /**
     * Location from which this application runs. Application instances run out of the data centers in the specified location, which is also where all of the application's end user content is stored.Defaults to us-central.View the list of supported locations (https://cloud.google.com/appengine/docs/locations).
     */
    location?: pulumi.Input<string>;
    /**
     * The service account associated with the application. This is the app-level default identity. If no identity provided during create version, Admin API will fallback to this one.
     */
    serviceAccount?: pulumi.Input<string>;
    /**
     * Serving status of this application.
     */
    servingStatus?: pulumi.Input<enums.appengine.v1.AppServingStatus>;
}
