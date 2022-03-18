// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Atomically updates the ResourceRecordSet collection.
 * Auto-naming is currently not supported for this resource.
 * Note - this resource's API doesn't support deletion. When deleted, the resource will persist
 * on Google Cloud even though it will be deleted from Pulumi state.
 */
export class Change extends pulumi.CustomResource {
    /**
     * Get an existing Change resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Change {
        return new Change(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:dns/v1:Change';

    /**
     * Returns true if the given object is an instance of Change.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Change {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Change.__pulumiType;
    }

    /**
     * Which ResourceRecordSets to add?
     */
    public readonly additions!: pulumi.Output<outputs.dns.v1.ResourceRecordSetResponse[]>;
    /**
     * Which ResourceRecordSets to remove? Must match existing data exactly.
     */
    public readonly deletions!: pulumi.Output<outputs.dns.v1.ResourceRecordSetResponse[]>;
    /**
     * If the DNS queries for the zone will be served.
     */
    public readonly isServing!: pulumi.Output<boolean>;
    public readonly kind!: pulumi.Output<string>;
    /**
     * The time that this operation was started by the server (output only). This is in RFC3339 text format.
     */
    public readonly startTime!: pulumi.Output<string>;
    /**
     * Status of the operation (output only). A status of "done" means that the request to update the authoritative servers has been sent, but the servers might not be updated yet.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;

    /**
     * Create a Change resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ChangeArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.managedZone === undefined) && !opts.urn) {
                throw new Error("Missing required property 'managedZone'");
            }
            resourceInputs["additions"] = args ? args.additions : undefined;
            resourceInputs["clientOperationId"] = args ? args.clientOperationId : undefined;
            resourceInputs["deletions"] = args ? args.deletions : undefined;
            resourceInputs["id"] = args ? args.id : undefined;
            resourceInputs["isServing"] = args ? args.isServing : undefined;
            resourceInputs["kind"] = args ? args.kind : undefined;
            resourceInputs["managedZone"] = args ? args.managedZone : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["startTime"] = args ? args.startTime : undefined;
            resourceInputs["status"] = undefined /*out*/;
        } else {
            resourceInputs["additions"] = undefined /*out*/;
            resourceInputs["deletions"] = undefined /*out*/;
            resourceInputs["isServing"] = undefined /*out*/;
            resourceInputs["kind"] = undefined /*out*/;
            resourceInputs["startTime"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Change.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Change resource.
 */
export interface ChangeArgs {
    /**
     * Which ResourceRecordSets to add?
     */
    additions?: pulumi.Input<pulumi.Input<inputs.dns.v1.ResourceRecordSetArgs>[]>;
    /**
     * For mutating operation requests only. An optional identifier specified by the client. Must be unique for operation resources in the Operations collection.
     */
    clientOperationId?: pulumi.Input<string>;
    /**
     * Which ResourceRecordSets to remove? Must match existing data exactly.
     */
    deletions?: pulumi.Input<pulumi.Input<inputs.dns.v1.ResourceRecordSetArgs>[]>;
    /**
     * Unique identifier for the resource; defined by the server (output only).
     */
    id?: pulumi.Input<string>;
    /**
     * If the DNS queries for the zone will be served.
     */
    isServing?: pulumi.Input<boolean>;
    kind?: pulumi.Input<string>;
    managedZone: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    /**
     * The time that this operation was started by the server (output only). This is in RFC3339 text format.
     */
    startTime?: pulumi.Input<string>;
}
