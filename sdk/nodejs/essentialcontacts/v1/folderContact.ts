// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Adds a new contact for a resource.
 * Auto-naming is currently not supported for this resource.
 */
export class FolderContact extends pulumi.CustomResource {
    /**
     * Get an existing FolderContact resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): FolderContact {
        return new FolderContact(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:essentialcontacts/v1:FolderContact';

    /**
     * Returns true if the given object is an instance of FolderContact.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FolderContact {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FolderContact.__pulumiType;
    }

    /**
     * The email address to send notifications to. This does not need to be a Google account.
     */
    public readonly email!: pulumi.Output<string>;
    /**
     * The preferred language for notifications, as a ISO 639-1 language code. See [Supported languages](https://cloud.google.com/resource-manager/docs/managing-notification-contacts#supported-languages) for a list of supported languages.
     */
    public readonly languageTag!: pulumi.Output<string>;
    /**
     * The identifier for the contact. Format: {resource_type}/{resource_id}/contacts/{contact_id}
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The categories of notifications that the contact will receive communications for.
     */
    public readonly notificationCategorySubscriptions!: pulumi.Output<string[]>;
    /**
     * The last time the validation_state was updated, either manually or automatically. A contact is considered stale if its validation state was updated more than 1 year ago.
     */
    public readonly validateTime!: pulumi.Output<string>;
    /**
     * The validity of the contact. A contact is considered valid if it is the correct recipient for notifications for a particular resource.
     */
    public readonly validationState!: pulumi.Output<string>;

    /**
     * Create a FolderContact resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FolderContactArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.email === undefined) && !opts.urn) {
                throw new Error("Missing required property 'email'");
            }
            if ((!args || args.folderId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'folderId'");
            }
            resourceInputs["email"] = args ? args.email : undefined;
            resourceInputs["folderId"] = args ? args.folderId : undefined;
            resourceInputs["languageTag"] = args ? args.languageTag : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["notificationCategorySubscriptions"] = args ? args.notificationCategorySubscriptions : undefined;
            resourceInputs["validateTime"] = args ? args.validateTime : undefined;
            resourceInputs["validationState"] = args ? args.validationState : undefined;
        } else {
            resourceInputs["email"] = undefined /*out*/;
            resourceInputs["languageTag"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["notificationCategorySubscriptions"] = undefined /*out*/;
            resourceInputs["validateTime"] = undefined /*out*/;
            resourceInputs["validationState"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(FolderContact.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a FolderContact resource.
 */
export interface FolderContactArgs {
    /**
     * The email address to send notifications to. This does not need to be a Google account.
     */
    email: pulumi.Input<string>;
    folderId: pulumi.Input<string>;
    /**
     * The preferred language for notifications, as a ISO 639-1 language code. See [Supported languages](https://cloud.google.com/resource-manager/docs/managing-notification-contacts#supported-languages) for a list of supported languages.
     */
    languageTag?: pulumi.Input<string>;
    /**
     * The identifier for the contact. Format: {resource_type}/{resource_id}/contacts/{contact_id}
     */
    name?: pulumi.Input<string>;
    /**
     * The categories of notifications that the contact will receive communications for.
     */
    notificationCategorySubscriptions?: pulumi.Input<pulumi.Input<enums.essentialcontacts.v1.FolderContactNotificationCategorySubscriptionsItem>[]>;
    /**
     * The last time the validation_state was updated, either manually or automatically. A contact is considered stale if its validation state was updated more than 1 year ago.
     */
    validateTime?: pulumi.Input<string>;
    /**
     * The validity of the contact. A contact is considered valid if it is the correct recipient for notifications for a particular resource.
     */
    validationState?: pulumi.Input<enums.essentialcontacts.v1.FolderContactValidationState>;
}
