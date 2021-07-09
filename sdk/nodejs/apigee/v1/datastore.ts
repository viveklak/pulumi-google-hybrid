// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Create a Datastore for an org
 */
export class Datastore extends pulumi.CustomResource {
    /**
     * Get an existing Datastore resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Datastore {
        return new Datastore(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:apigee/v1:Datastore';

    /**
     * Returns true if the given object is an instance of Datastore.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Datastore {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Datastore.__pulumiType;
    }

    /**
     * Datastore create time, in milliseconds since the epoch of 1970-01-01T00:00:00Z
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * Datastore Configurations.
     */
    public readonly datastoreConfig!: pulumi.Output<outputs.apigee.v1.GoogleCloudApigeeV1DatastoreConfigResponse>;
    /**
     * Display name in UI
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * Datastore last update time, in milliseconds since the epoch of 1970-01-01T00:00:00Z
     */
    public /*out*/ readonly lastUpdateTime!: pulumi.Output<string>;
    /**
     * Organization that the datastore belongs to
     */
    public /*out*/ readonly org!: pulumi.Output<string>;
    /**
     * Resource link of Datastore. Example: `/organizations/{org}/analytics/datastores/{uuid}`
     */
    public /*out*/ readonly self!: pulumi.Output<string>;
    /**
     * Destination storage type. Supported types `gcs` or `bigquery`.
     */
    public readonly targetType!: pulumi.Output<string>;

    /**
     * Create a Datastore resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DatastoreArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.displayName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'displayName'");
            }
            if ((!args || args.organizationId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'organizationId'");
            }
            inputs["datastoreConfig"] = args ? args.datastoreConfig : undefined;
            inputs["displayName"] = args ? args.displayName : undefined;
            inputs["organizationId"] = args ? args.organizationId : undefined;
            inputs["targetType"] = args ? args.targetType : undefined;
            inputs["createTime"] = undefined /*out*/;
            inputs["lastUpdateTime"] = undefined /*out*/;
            inputs["org"] = undefined /*out*/;
            inputs["self"] = undefined /*out*/;
        } else {
            inputs["createTime"] = undefined /*out*/;
            inputs["datastoreConfig"] = undefined /*out*/;
            inputs["displayName"] = undefined /*out*/;
            inputs["lastUpdateTime"] = undefined /*out*/;
            inputs["org"] = undefined /*out*/;
            inputs["self"] = undefined /*out*/;
            inputs["targetType"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Datastore.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Datastore resource.
 */
export interface DatastoreArgs {
    /**
     * Datastore Configurations.
     */
    datastoreConfig?: pulumi.Input<inputs.apigee.v1.GoogleCloudApigeeV1DatastoreConfigArgs>;
    /**
     * Display name in UI
     */
    displayName: pulumi.Input<string>;
    organizationId: pulumi.Input<string>;
    /**
     * Destination storage type. Supported types `gcs` or `bigquery`.
     */
    targetType?: pulumi.Input<string>;
}
