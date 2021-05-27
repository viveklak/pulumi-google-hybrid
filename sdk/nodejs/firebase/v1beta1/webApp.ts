// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Requests the creation of a new WebApp in the specified FirebaseProject. The result of this call is an `Operation` which can be used to track the provisioning process. The `Operation` is automatically deleted after completion, so there is no need to call `DeleteOperation`.
 */
export class WebApp extends pulumi.CustomResource {
    /**
     * Get an existing WebApp resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): WebApp {
        return new WebApp(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:firebase/v1beta1:WebApp';

    /**
     * Returns true if the given object is an instance of WebApp.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is WebApp {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === WebApp.__pulumiType;
    }

    /**
     * Immutable. The globally unique, Firebase-assigned identifier for the `WebApp`. This identifier should be treated as an opaque token, as the data format is not specified.
     */
    public readonly appId!: pulumi.Output<string>;
    /**
     * The URLs where the `WebApp` is hosted.
     */
    public readonly appUrls!: pulumi.Output<string[]>;
    /**
     * The user-assigned display name for the `WebApp`.
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * The resource name of the WebApp, in the format: projects/PROJECT_IDENTIFIER /webApps/APP_ID * PROJECT_IDENTIFIER: the parent Project's [`ProjectNumber`](../projects#FirebaseProject.FIELDS.project_number) ***(recommended)*** or its [`ProjectId`](../projects#FirebaseProject.FIELDS.project_id). Learn more about using project identifiers in Google's [AIP 2510 standard](https://google.aip.dev/cloud/2510). Note that the value for PROJECT_IDENTIFIER in any response body will be the `ProjectId`. * APP_ID: the globally unique, Firebase-assigned identifier for the App (see [`appId`](../projects.webApps#WebApp.FIELDS.app_id)).
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Immutable. A user-assigned unique identifier of the parent FirebaseProject for the `WebApp`.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * Immutable. A unique, Firebase-assigned identifier for the `WebApp`. This identifier is only used to populate the `namespace` value for the `WebApp`. For most use cases, use `appId` to identify or reference the App. The `webId` value is only unique within a `FirebaseProject` and its associated Apps.
     */
    public /*out*/ readonly webId!: pulumi.Output<string>;

    /**
     * Create a WebApp resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: WebAppArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.project === undefined) && !opts.urn) {
                throw new Error("Missing required property 'project'");
            }
            inputs["appId"] = args ? args.appId : undefined;
            inputs["appUrls"] = args ? args.appUrls : undefined;
            inputs["displayName"] = args ? args.displayName : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["webId"] = undefined /*out*/;
        } else {
            inputs["appId"] = undefined /*out*/;
            inputs["appUrls"] = undefined /*out*/;
            inputs["displayName"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["project"] = undefined /*out*/;
            inputs["webId"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(WebApp.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a WebApp resource.
 */
export interface WebAppArgs {
    /**
     * Immutable. The globally unique, Firebase-assigned identifier for the `WebApp`. This identifier should be treated as an opaque token, as the data format is not specified.
     */
    readonly appId?: pulumi.Input<string>;
    /**
     * The URLs where the `WebApp` is hosted.
     */
    readonly appUrls?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The user-assigned display name for the `WebApp`.
     */
    readonly displayName?: pulumi.Input<string>;
    /**
     * The resource name of the WebApp, in the format: projects/PROJECT_IDENTIFIER /webApps/APP_ID * PROJECT_IDENTIFIER: the parent Project's [`ProjectNumber`](../projects#FirebaseProject.FIELDS.project_number) ***(recommended)*** or its [`ProjectId`](../projects#FirebaseProject.FIELDS.project_id). Learn more about using project identifiers in Google's [AIP 2510 standard](https://google.aip.dev/cloud/2510). Note that the value for PROJECT_IDENTIFIER in any response body will be the `ProjectId`. * APP_ID: the globally unique, Firebase-assigned identifier for the App (see [`appId`](../projects.webApps#WebApp.FIELDS.app_id)).
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Immutable. A user-assigned unique identifier of the parent FirebaseProject for the `WebApp`.
     */
    readonly project: pulumi.Input<string>;
}
