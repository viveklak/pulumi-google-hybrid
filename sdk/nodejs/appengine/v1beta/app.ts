// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates an App Engine application for a Google Cloud Platform project. Required fields: id - The ID of the target Cloud Platform project. location - The region (https://cloud.google.com/appengine/docs/locations) where you want the App Engine application located.For more information about App Engine applications, see Managing Projects, Applications, and Billing (https://cloud.google.com/appengine/docs/standard/python/console/).
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
    public static readonly __pulumiType = 'google-native:appengine/v1beta:App';

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
     * Google Cloud Storage bucket that can be used for storing files associated with this application. This bucket is associated with the application and can be used by the gcloud deployment commands.@OutputOnly
     */
    public readonly codeBucket!: pulumi.Output<string>;
    /**
     * The type of the Cloud Firestore or Cloud Datastore database associated with this application.
     */
    public readonly databaseType!: pulumi.Output<string>;
    /**
     * Google Cloud Storage bucket that can be used by this application to store content.@OutputOnly
     */
    public readonly defaultBucket!: pulumi.Output<string>;
    /**
     * Cookie expiration policy for this application.
     */
    public readonly defaultCookieExpiration!: pulumi.Output<string>;
    /**
     * Hostname used to reach this application, as resolved by App Engine.@OutputOnly
     */
    public readonly defaultHostname!: pulumi.Output<string>;
    /**
     * HTTP path dispatch rules for requests to the application that do not explicitly target a service or version. Rules are order-dependent. Up to 20 dispatch rules can be supported.
     */
    public readonly dispatchRules!: pulumi.Output<outputs.appengine.v1beta.UrlDispatchRuleResponse[]>;
    /**
     * The feature specific settings to be used in the application.
     */
    public readonly featureSettings!: pulumi.Output<outputs.appengine.v1beta.FeatureSettingsResponse>;
    /**
     * The Google Container Registry domain used for storing managed build docker images for this application.
     */
    public readonly gcrDomain!: pulumi.Output<string>;
    public readonly iap!: pulumi.Output<outputs.appengine.v1beta.IdentityAwareProxyResponse>;
    /**
     * Location from which this application runs. Application instances run out of the data centers in the specified location, which is also where all of the application's end user content is stored.Defaults to us-central.View the list of supported locations (https://cloud.google.com/appengine/docs/locations).
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * Full path to the Application resource in the API. Example: apps/myapp.@OutputOnly
     */
    public readonly name!: pulumi.Output<string>;
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
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            inputs["authDomain"] = args ? args.authDomain : undefined;
            inputs["codeBucket"] = args ? args.codeBucket : undefined;
            inputs["databaseType"] = args ? args.databaseType : undefined;
            inputs["defaultBucket"] = args ? args.defaultBucket : undefined;
            inputs["defaultCookieExpiration"] = args ? args.defaultCookieExpiration : undefined;
            inputs["defaultHostname"] = args ? args.defaultHostname : undefined;
            inputs["dispatchRules"] = args ? args.dispatchRules : undefined;
            inputs["featureSettings"] = args ? args.featureSettings : undefined;
            inputs["gcrDomain"] = args ? args.gcrDomain : undefined;
            inputs["iap"] = args ? args.iap : undefined;
            inputs["id"] = args ? args.id : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["servingStatus"] = args ? args.servingStatus : undefined;
        } else {
            inputs["authDomain"] = undefined /*out*/;
            inputs["codeBucket"] = undefined /*out*/;
            inputs["databaseType"] = undefined /*out*/;
            inputs["defaultBucket"] = undefined /*out*/;
            inputs["defaultCookieExpiration"] = undefined /*out*/;
            inputs["defaultHostname"] = undefined /*out*/;
            inputs["dispatchRules"] = undefined /*out*/;
            inputs["featureSettings"] = undefined /*out*/;
            inputs["gcrDomain"] = undefined /*out*/;
            inputs["iap"] = undefined /*out*/;
            inputs["location"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["servingStatus"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(App.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a App resource.
 */
export interface AppArgs {
    /**
     * Google Apps authentication domain that controls which users can access this application.Defaults to open access for any Google Account.
     */
    readonly authDomain?: pulumi.Input<string>;
    /**
     * Google Cloud Storage bucket that can be used for storing files associated with this application. This bucket is associated with the application and can be used by the gcloud deployment commands.@OutputOnly
     */
    readonly codeBucket?: pulumi.Input<string>;
    /**
     * The type of the Cloud Firestore or Cloud Datastore database associated with this application.
     */
    readonly databaseType?: pulumi.Input<string>;
    /**
     * Google Cloud Storage bucket that can be used by this application to store content.@OutputOnly
     */
    readonly defaultBucket?: pulumi.Input<string>;
    /**
     * Cookie expiration policy for this application.
     */
    readonly defaultCookieExpiration?: pulumi.Input<string>;
    /**
     * Hostname used to reach this application, as resolved by App Engine.@OutputOnly
     */
    readonly defaultHostname?: pulumi.Input<string>;
    /**
     * HTTP path dispatch rules for requests to the application that do not explicitly target a service or version. Rules are order-dependent. Up to 20 dispatch rules can be supported.
     */
    readonly dispatchRules?: pulumi.Input<pulumi.Input<inputs.appengine.v1beta.UrlDispatchRuleArgs>[]>;
    /**
     * The feature specific settings to be used in the application.
     */
    readonly featureSettings?: pulumi.Input<inputs.appengine.v1beta.FeatureSettingsArgs>;
    /**
     * The Google Container Registry domain used for storing managed build docker images for this application.
     */
    readonly gcrDomain?: pulumi.Input<string>;
    readonly iap?: pulumi.Input<inputs.appengine.v1beta.IdentityAwareProxyArgs>;
    /**
     * Identifier of the Application resource. This identifier is equivalent to the project ID of the Google Cloud Platform project where you want to deploy your application. Example: myapp.
     */
    readonly id?: pulumi.Input<string>;
    /**
     * Location from which this application runs. Application instances run out of the data centers in the specified location, which is also where all of the application's end user content is stored.Defaults to us-central.View the list of supported locations (https://cloud.google.com/appengine/docs/locations).
     */
    readonly location?: pulumi.Input<string>;
    /**
     * Full path to the Application resource in the API. Example: apps/myapp.@OutputOnly
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Serving status of this application.
     */
    readonly servingStatus?: pulumi.Input<string>;
}
