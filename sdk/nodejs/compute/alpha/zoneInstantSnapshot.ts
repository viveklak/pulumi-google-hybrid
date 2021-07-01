// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Creates an instant snapshot in the specified zone.
 */
export class ZoneInstantSnapshot extends pulumi.CustomResource {
    /**
     * Get an existing ZoneInstantSnapshot resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ZoneInstantSnapshot {
        return new ZoneInstantSnapshot(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:compute/alpha:ZoneInstantSnapshot';

    /**
     * Returns true if the given object is an instance of ZoneInstantSnapshot.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ZoneInstantSnapshot {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ZoneInstantSnapshot.__pulumiType;
    }

    /**
     * Creation timestamp in RFC3339 text format.
     */
    public /*out*/ readonly creationTimestamp!: pulumi.Output<string>;
    /**
     * An optional description of this resource. Provide this property when you create the resource.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * Size of the source disk, specified in GB.
     */
    public /*out*/ readonly diskSizeGb!: pulumi.Output<string>;
    /**
     * Whether to attempt an application consistent instant snapshot by informing the OS to prepare for the snapshot process. Currently only supported on Windows instances using the Volume Shadow Copy Service (VSS).
     */
    public readonly guestFlush!: pulumi.Output<boolean>;
    /**
     * Type of the resource. Always compute#instantSnapshot for InstantSnapshot resources.
     */
    public /*out*/ readonly kind!: pulumi.Output<string>;
    /**
     * A fingerprint for the labels being applied to this InstantSnapshot, which is essentially a hash of the labels set used for optimistic locking. The fingerprint is initially generated by Compute Engine and changes after every request to modify or update labels. You must always provide an up-to-date fingerprint hash in order to update or change labels, otherwise the request will fail with error 412 conditionNotMet.
     *
     * To see the latest fingerprint, make a get() request to retrieve a InstantSnapshot.
     */
    public /*out*/ readonly labelFingerprint!: pulumi.Output<string>;
    /**
     * Labels to apply to this InstantSnapshot. These can be later modified by the setLabels method. Label values may be empty.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string}>;
    /**
     * Name of the resource; provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * URL of the region where the instant snapshot resides. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
     */
    public /*out*/ readonly region!: pulumi.Output<string>;
    /**
     * Reserved for future use.
     */
    public /*out*/ readonly satisfiesPzs!: pulumi.Output<boolean>;
    /**
     * Server-defined URL for the resource.
     */
    public /*out*/ readonly selfLink!: pulumi.Output<string>;
    /**
     * Server-defined URL for this resource's resource id.
     */
    public /*out*/ readonly selfLinkWithId!: pulumi.Output<string>;
    /**
     * URL of the source disk used to create this instant snapshot. Note that the source disk must be in the same zone/region as the instant snapshot to be created. This can be a full or valid partial URL. For example, the following are valid values:  
     * - https://www.googleapis.com/compute/v1/projects/project/zones/zone/disks/disk  
     * - https://www.googleapis.com/compute/v1/projects/project/regions/region/disks/disk  
     * - projects/project/zones/zone/disks/disk  
     * - projects/project/regions/region/disks/disk  
     * - zones/zone/disks/disk  
     * - regions/region/disks/disk
     */
    public readonly sourceDisk!: pulumi.Output<string>;
    /**
     * The ID value of the disk used to create this InstantSnapshot. This value may be used to determine whether the InstantSnapshot was taken from the current or a previous instance of a given disk name.
     */
    public /*out*/ readonly sourceDiskId!: pulumi.Output<string>;
    /**
     * The status of the instantSnapshot. This can be CREATING, DELETING, FAILED, or READY.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * URL of the zone where the instant snapshot resides. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
     */
    public readonly zone!: pulumi.Output<string>;

    /**
     * Create a ZoneInstantSnapshot resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ZoneInstantSnapshotArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.project === undefined) && !opts.urn) {
                throw new Error("Missing required property 'project'");
            }
            if ((!args || args.zone === undefined) && !opts.urn) {
                throw new Error("Missing required property 'zone'");
            }
            inputs["description"] = args ? args.description : undefined;
            inputs["guestFlush"] = args ? args.guestFlush : undefined;
            inputs["labels"] = args ? args.labels : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["requestId"] = args ? args.requestId : undefined;
            inputs["sourceDisk"] = args ? args.sourceDisk : undefined;
            inputs["zone"] = args ? args.zone : undefined;
            inputs["creationTimestamp"] = undefined /*out*/;
            inputs["diskSizeGb"] = undefined /*out*/;
            inputs["kind"] = undefined /*out*/;
            inputs["labelFingerprint"] = undefined /*out*/;
            inputs["region"] = undefined /*out*/;
            inputs["satisfiesPzs"] = undefined /*out*/;
            inputs["selfLink"] = undefined /*out*/;
            inputs["selfLinkWithId"] = undefined /*out*/;
            inputs["sourceDiskId"] = undefined /*out*/;
            inputs["status"] = undefined /*out*/;
        } else {
            inputs["creationTimestamp"] = undefined /*out*/;
            inputs["description"] = undefined /*out*/;
            inputs["diskSizeGb"] = undefined /*out*/;
            inputs["guestFlush"] = undefined /*out*/;
            inputs["kind"] = undefined /*out*/;
            inputs["labelFingerprint"] = undefined /*out*/;
            inputs["labels"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["region"] = undefined /*out*/;
            inputs["satisfiesPzs"] = undefined /*out*/;
            inputs["selfLink"] = undefined /*out*/;
            inputs["selfLinkWithId"] = undefined /*out*/;
            inputs["sourceDisk"] = undefined /*out*/;
            inputs["sourceDiskId"] = undefined /*out*/;
            inputs["status"] = undefined /*out*/;
            inputs["zone"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(ZoneInstantSnapshot.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a ZoneInstantSnapshot resource.
 */
export interface ZoneInstantSnapshotArgs {
    /**
     * An optional description of this resource. Provide this property when you create the resource.
     */
    description?: pulumi.Input<string>;
    /**
     * Whether to attempt an application consistent instant snapshot by informing the OS to prepare for the snapshot process. Currently only supported on Windows instances using the Volume Shadow Copy Service (VSS).
     */
    guestFlush?: pulumi.Input<boolean>;
    /**
     * Labels to apply to this InstantSnapshot. These can be later modified by the setLabels method. Label values may be empty.
     */
    labels?: pulumi.Input<{[key: string]: string}>;
    /**
     * Name of the resource; provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
     */
    name?: pulumi.Input<string>;
    project: pulumi.Input<string>;
    requestId?: pulumi.Input<string>;
    /**
     * URL of the source disk used to create this instant snapshot. Note that the source disk must be in the same zone/region as the instant snapshot to be created. This can be a full or valid partial URL. For example, the following are valid values:  
     * - https://www.googleapis.com/compute/v1/projects/project/zones/zone/disks/disk  
     * - https://www.googleapis.com/compute/v1/projects/project/regions/region/disks/disk  
     * - projects/project/zones/zone/disks/disk  
     * - projects/project/regions/region/disks/disk  
     * - zones/zone/disks/disk  
     * - regions/region/disks/disk
     */
    sourceDisk?: pulumi.Input<string>;
    zone: pulumi.Input<string>;
}
