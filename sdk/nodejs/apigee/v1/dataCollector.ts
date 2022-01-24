// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates a new data collector.
 */
export class DataCollector extends pulumi.CustomResource {
    /**
     * Get an existing DataCollector resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): DataCollector {
        return new DataCollector(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:apigee/v1:DataCollector';

    /**
     * Returns true if the given object is an instance of DataCollector.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DataCollector {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DataCollector.__pulumiType;
    }

    /**
     * The time at which the data collector was created in milliseconds since the epoch.
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * A description of the data collector.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * The time at which the Data Collector was last updated in milliseconds since the epoch.
     */
    public /*out*/ readonly lastModifiedAt!: pulumi.Output<string>;
    /**
     * ID of the data collector. Must begin with `dc_`.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Immutable. The type of data this data collector will collect.
     */
    public readonly type!: pulumi.Output<string>;

    /**
     * Create a DataCollector resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DataCollectorArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.organizationId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'organizationId'");
            }
            resourceInputs["dataCollectorId"] = args ? args.dataCollectorId : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["organizationId"] = args ? args.organizationId : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["lastModifiedAt"] = undefined /*out*/;
        } else {
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["lastModifiedAt"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(DataCollector.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a DataCollector resource.
 */
export interface DataCollectorArgs {
    dataCollectorId?: pulumi.Input<string>;
    /**
     * A description of the data collector.
     */
    description?: pulumi.Input<string>;
    /**
     * ID of the data collector. Must begin with `dc_`.
     */
    name?: pulumi.Input<string>;
    organizationId: pulumi.Input<string>;
    /**
     * Immutable. The type of data this data collector will collect.
     */
    type?: pulumi.Input<enums.apigee.v1.DataCollectorType>;
}
