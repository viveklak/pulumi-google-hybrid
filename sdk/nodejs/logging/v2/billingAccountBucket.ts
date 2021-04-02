// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Creates a bucket that can be used to store log entries. Once a bucket has been created, the region cannot be changed.
 */
export class BillingAccountBucket extends pulumi.CustomResource {
    /**
     * Get an existing BillingAccountBucket resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): BillingAccountBucket {
        return new BillingAccountBucket(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-cloud:logging/v2:BillingAccountBucket';

    /**
     * Returns true if the given object is an instance of BillingAccountBucket.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is BillingAccountBucket {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === BillingAccountBucket.__pulumiType;
    }


    /**
     * Create a BillingAccountBucket resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: BillingAccountBucketArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.billingAccountsId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'billingAccountsId'");
            }
            if ((!args || args.bucketsId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'bucketsId'");
            }
            if ((!args || args.locationsId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'locationsId'");
            }
            inputs["billingAccountsId"] = args ? args.billingAccountsId : undefined;
            inputs["bucketsId"] = args ? args.bucketsId : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["locationsId"] = args ? args.locationsId : undefined;
            inputs["locked"] = args ? args.locked : undefined;
            inputs["retentionDays"] = args ? args.retentionDays : undefined;
        } else {
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(BillingAccountBucket.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a BillingAccountBucket resource.
 */
export interface BillingAccountBucketArgs {
    readonly billingAccountsId: pulumi.Input<string>;
    readonly bucketsId: pulumi.Input<string>;
    /**
     * Describes this bucket.
     */
    readonly description?: pulumi.Input<string>;
    readonly locationsId: pulumi.Input<string>;
    /**
     * Whether the bucket has been locked. The retention period on a locked bucket may not be changed. Locked buckets may only be deleted if they are empty.
     */
    readonly locked?: pulumi.Input<boolean>;
    /**
     * Logs will be retained by default for this amount of time, after which they will automatically be deleted. The minimum retention period is 1 day. If this value is set to zero at bucket creation time, the default time of 30 days will be used.
     */
    readonly retentionDays?: pulumi.Input<number>;
}