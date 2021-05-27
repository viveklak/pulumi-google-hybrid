// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Requests the creation of a new IosApp in the specified FirebaseProject. The result of this call is an `Operation` which can be used to track the provisioning process. The `Operation` is automatically deleted after completion, so there is no need to call `DeleteOperation`.
 */
export class IosApp extends pulumi.CustomResource {
    /**
     * Get an existing IosApp resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): IosApp {
        return new IosApp(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:firebase/v1beta1:IosApp';

    /**
     * Returns true if the given object is an instance of IosApp.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is IosApp {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === IosApp.__pulumiType;
    }

    /**
     * Immutable. The globally unique, Firebase-assigned identifier for the `IosApp`. This identifier should be treated as an opaque token, as the data format is not specified.
     */
    public readonly appId!: pulumi.Output<string>;
    /**
     * The automatically generated Apple ID assigned to the iOS app by Apple in the iOS App Store.
     */
    public readonly appStoreId!: pulumi.Output<string>;
    /**
     * Immutable. The canonical bundle ID of the iOS app as it would appear in the iOS AppStore.
     */
    public readonly bundleId!: pulumi.Output<string>;
    /**
     * The user-assigned display name for the `IosApp`.
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * The resource name of the IosApp, in the format: projects/PROJECT_IDENTIFIER /iosApps/APP_ID * PROJECT_IDENTIFIER: the parent Project's [`ProjectNumber`](../projects#FirebaseProject.FIELDS.project_number) ***(recommended)*** or its [`ProjectId`](../projects#FirebaseProject.FIELDS.project_id). Learn more about using project identifiers in Google's [AIP 2510 standard](https://google.aip.dev/cloud/2510). Note that the value for PROJECT_IDENTIFIER in any response body will be the `ProjectId`. * APP_ID: the globally unique, Firebase-assigned identifier for the App (see [`appId`](../projects.iosApps#IosApp.FIELDS.app_id)).
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Immutable. A user-assigned unique identifier of the parent FirebaseProject for the `IosApp`.
     */
    public readonly project!: pulumi.Output<string>;

    /**
     * Create a IosApp resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: IosAppArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.project === undefined) && !opts.urn) {
                throw new Error("Missing required property 'project'");
            }
            inputs["appId"] = args ? args.appId : undefined;
            inputs["appStoreId"] = args ? args.appStoreId : undefined;
            inputs["bundleId"] = args ? args.bundleId : undefined;
            inputs["displayName"] = args ? args.displayName : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["project"] = args ? args.project : undefined;
        } else {
            inputs["appId"] = undefined /*out*/;
            inputs["appStoreId"] = undefined /*out*/;
            inputs["bundleId"] = undefined /*out*/;
            inputs["displayName"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["project"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(IosApp.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a IosApp resource.
 */
export interface IosAppArgs {
    /**
     * Immutable. The globally unique, Firebase-assigned identifier for the `IosApp`. This identifier should be treated as an opaque token, as the data format is not specified.
     */
    readonly appId?: pulumi.Input<string>;
    /**
     * The automatically generated Apple ID assigned to the iOS app by Apple in the iOS App Store.
     */
    readonly appStoreId?: pulumi.Input<string>;
    /**
     * Immutable. The canonical bundle ID of the iOS app as it would appear in the iOS AppStore.
     */
    readonly bundleId?: pulumi.Input<string>;
    /**
     * The user-assigned display name for the `IosApp`.
     */
    readonly displayName?: pulumi.Input<string>;
    /**
     * The resource name of the IosApp, in the format: projects/PROJECT_IDENTIFIER /iosApps/APP_ID * PROJECT_IDENTIFIER: the parent Project's [`ProjectNumber`](../projects#FirebaseProject.FIELDS.project_number) ***(recommended)*** or its [`ProjectId`](../projects#FirebaseProject.FIELDS.project_id). Learn more about using project identifiers in Google's [AIP 2510 standard](https://google.aip.dev/cloud/2510). Note that the value for PROJECT_IDENTIFIER in any response body will be the `ProjectId`. * APP_ID: the globally unique, Firebase-assigned identifier for the App (see [`appId`](../projects.iosApps#IosApp.FIELDS.app_id)).
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Immutable. A user-assigned unique identifier of the parent FirebaseProject for the `IosApp`.
     */
    readonly project: pulumi.Input<string>;
}
