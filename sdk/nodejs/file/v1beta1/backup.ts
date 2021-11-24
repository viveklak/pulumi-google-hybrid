// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Creates a backup.
 * Auto-naming is currently not supported for this resource.
 */
export class Backup extends pulumi.CustomResource {
    /**
     * Get an existing Backup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Backup {
        return new Backup(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:file/v1beta1:Backup';

    /**
     * Returns true if the given object is an instance of Backup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Backup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Backup.__pulumiType;
    }

    /**
     * Capacity of the source file share when the backup was created.
     */
    public /*out*/ readonly capacityGb!: pulumi.Output<string>;
    /**
     * The time when the backup was created.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * A description of the backup with 2048 characters or less. Requests with longer descriptions will be rejected.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * Amount of bytes that will be downloaded if the backup is restored
     */
    public /*out*/ readonly downloadBytes!: pulumi.Output<string>;
    /**
     * Resource labels to represent user provided metadata.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string}>;
    /**
     * The resource name of the backup, in the format `projects/{project_id}/locations/{location_id}/backups/{backup_id}`.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Reserved for future use.
     */
    public /*out*/ readonly satisfiesPzs!: pulumi.Output<boolean>;
    /**
     * Name of the file share in the source Cloud Filestore instance that the backup is created from.
     */
    public readonly sourceFileShare!: pulumi.Output<string>;
    /**
     * The resource name of the source Cloud Filestore instance, in the format `projects/{project_id}/locations/{location_id}/instances/{instance_id}`, used to create this backup.
     */
    public readonly sourceInstance!: pulumi.Output<string>;
    /**
     * The service tier of the source Cloud Filestore instance that this backup is created from.
     */
    public /*out*/ readonly sourceInstanceTier!: pulumi.Output<string>;
    /**
     * The backup state.
     */
    public /*out*/ readonly state!: pulumi.Output<string>;
    /**
     * The size of the storage used by the backup. As backups share storage, this number is expected to change with backup creation/deletion.
     */
    public /*out*/ readonly storageBytes!: pulumi.Output<string>;

    /**
     * Create a Backup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: BackupArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.backupId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'backupId'");
            }
            resourceInputs["backupId"] = args ? args.backupId : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["sourceFileShare"] = args ? args.sourceFileShare : undefined;
            resourceInputs["sourceInstance"] = args ? args.sourceInstance : undefined;
            resourceInputs["capacityGb"] = undefined /*out*/;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["downloadBytes"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["satisfiesPzs"] = undefined /*out*/;
            resourceInputs["sourceInstanceTier"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
            resourceInputs["storageBytes"] = undefined /*out*/;
        } else {
            resourceInputs["capacityGb"] = undefined /*out*/;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["downloadBytes"] = undefined /*out*/;
            resourceInputs["labels"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["satisfiesPzs"] = undefined /*out*/;
            resourceInputs["sourceFileShare"] = undefined /*out*/;
            resourceInputs["sourceInstance"] = undefined /*out*/;
            resourceInputs["sourceInstanceTier"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
            resourceInputs["storageBytes"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Backup.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Backup resource.
 */
export interface BackupArgs {
    backupId: pulumi.Input<string>;
    /**
     * A description of the backup with 2048 characters or less. Requests with longer descriptions will be rejected.
     */
    description?: pulumi.Input<string>;
    /**
     * Resource labels to represent user provided metadata.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    location?: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    /**
     * Name of the file share in the source Cloud Filestore instance that the backup is created from.
     */
    sourceFileShare?: pulumi.Input<string>;
    /**
     * The resource name of the source Cloud Filestore instance, in the format `projects/{project_id}/locations/{location_id}/instances/{instance_id}`, used to create this backup.
     */
    sourceInstance?: pulumi.Input<string>;
}
