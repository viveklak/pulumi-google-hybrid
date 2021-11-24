// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates a pre-built stored infoType to be used for inspection. See https://cloud.google.com/dlp/docs/creating-stored-infotypes to learn more.
 * Auto-naming is currently not supported for this resource.
 */
export class StoredInfoType extends pulumi.CustomResource {
    /**
     * Get an existing StoredInfoType resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): StoredInfoType {
        return new StoredInfoType(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:dlp/v2:StoredInfoType';

    /**
     * Returns true if the given object is an instance of StoredInfoType.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is StoredInfoType {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === StoredInfoType.__pulumiType;
    }

    /**
     * Current version of the stored info type.
     */
    public /*out*/ readonly currentVersion!: pulumi.Output<outputs.dlp.v2.GooglePrivacyDlpV2StoredInfoTypeVersionResponse>;
    /**
     * Resource name.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Pending versions of the stored info type. Empty if no versions are pending.
     */
    public /*out*/ readonly pendingVersions!: pulumi.Output<outputs.dlp.v2.GooglePrivacyDlpV2StoredInfoTypeVersionResponse[]>;

    /**
     * Create a StoredInfoType resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: StoredInfoTypeArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.config === undefined) && !opts.urn) {
                throw new Error("Missing required property 'config'");
            }
            resourceInputs["config"] = args ? args.config : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["storedInfoTypeId"] = args ? args.storedInfoTypeId : undefined;
            resourceInputs["currentVersion"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["pendingVersions"] = undefined /*out*/;
        } else {
            resourceInputs["currentVersion"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["pendingVersions"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(StoredInfoType.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a StoredInfoType resource.
 */
export interface StoredInfoTypeArgs {
    /**
     * Configuration of the storedInfoType to create.
     */
    config: pulumi.Input<inputs.dlp.v2.GooglePrivacyDlpV2StoredInfoTypeConfigArgs>;
    location?: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    /**
     * The storedInfoType ID can contain uppercase and lowercase letters, numbers, and hyphens; that is, it must match the regular expression: `[a-zA-Z\d-_]+`. The maximum length is 100 characters. Can be empty to allow the system to generate one.
     */
    storedInfoTypeId?: pulumi.Input<string>;
}
