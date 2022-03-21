// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates a zone resource within a lake.
 * Auto-naming is currently not supported for this resource.
 */
export class Zone extends pulumi.CustomResource {
    /**
     * Get an existing Zone resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Zone {
        return new Zone(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:dataplex/v1:Zone';

    /**
     * Returns true if the given object is an instance of Zone.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Zone {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Zone.__pulumiType;
    }

    /**
     * Aggregated status of the underlying assets of the zone.
     */
    public /*out*/ readonly assetStatus!: pulumi.Output<outputs.dataplex.v1.GoogleCloudDataplexV1AssetStatusResponse>;
    /**
     * The time when the zone was created.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * Optional. Description of the zone.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * Optional. Specification of the discovery feature applied to data in this zone.
     */
    public readonly discoverySpec!: pulumi.Output<outputs.dataplex.v1.GoogleCloudDataplexV1ZoneDiscoverySpecResponse>;
    /**
     * Optional. User friendly display name.
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * Optional. User defined labels for the zone.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string}>;
    /**
     * The relative resource name of the zone, of the form: projects/{project_number}/locations/{location_id}/lakes/{lake_id}/zones/{zone_id}.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Specification of the resources that are referenced by the assets within this zone.
     */
    public readonly resourceSpec!: pulumi.Output<outputs.dataplex.v1.GoogleCloudDataplexV1ZoneResourceSpecResponse>;
    /**
     * Current state of the zone.
     */
    public /*out*/ readonly state!: pulumi.Output<string>;
    /**
     * Immutable. The type of the zone.
     */
    public readonly type!: pulumi.Output<string>;
    /**
     * System generated globally unique ID for the zone. This ID will be different if the zone is deleted and re-created with the same name.
     */
    public /*out*/ readonly uid!: pulumi.Output<string>;
    /**
     * The time when the zone was last updated.
     */
    public /*out*/ readonly updateTime!: pulumi.Output<string>;

    /**
     * Create a Zone resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ZoneArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.lakeId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'lakeId'");
            }
            if ((!args || args.resourceSpec === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceSpec'");
            }
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            if ((!args || args.zoneId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'zoneId'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["discoverySpec"] = args ? args.discoverySpec : undefined;
            resourceInputs["displayName"] = args ? args.displayName : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["lakeId"] = args ? args.lakeId : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["resourceSpec"] = args ? args.resourceSpec : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["validateOnly"] = args ? args.validateOnly : undefined;
            resourceInputs["zoneId"] = args ? args.zoneId : undefined;
            resourceInputs["assetStatus"] = undefined /*out*/;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
            resourceInputs["uid"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        } else {
            resourceInputs["assetStatus"] = undefined /*out*/;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["discoverySpec"] = undefined /*out*/;
            resourceInputs["displayName"] = undefined /*out*/;
            resourceInputs["labels"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["resourceSpec"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
            resourceInputs["uid"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Zone.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Zone resource.
 */
export interface ZoneArgs {
    /**
     * Optional. Description of the zone.
     */
    description?: pulumi.Input<string>;
    /**
     * Optional. Specification of the discovery feature applied to data in this zone.
     */
    discoverySpec?: pulumi.Input<inputs.dataplex.v1.GoogleCloudDataplexV1ZoneDiscoverySpecArgs>;
    /**
     * Optional. User friendly display name.
     */
    displayName?: pulumi.Input<string>;
    /**
     * Optional. User defined labels for the zone.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    lakeId: pulumi.Input<string>;
    location?: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    /**
     * Specification of the resources that are referenced by the assets within this zone.
     */
    resourceSpec: pulumi.Input<inputs.dataplex.v1.GoogleCloudDataplexV1ZoneResourceSpecArgs>;
    /**
     * Immutable. The type of the zone.
     */
    type: pulumi.Input<enums.dataplex.v1.ZoneType>;
    /**
     * Optional. Only validate the request, but do not perform mutations. The default is false.
     */
    validateOnly?: pulumi.Input<string>;
    /**
     * Required. Zone identifier. This ID will be used to generate names such as database and dataset names when publishing metadata to Hive Metastore and BigQuery. * Must contain only lowercase letters, numbers and hyphens. * Must start with a letter. * Must end with a number or a letter. * Must be between 1-63 characters. * Must be unique across all lakes from all locations in a project. * Must not be one of the reserved IDs (i.e. "default", "global-temp")
     */
    zoneId: pulumi.Input<string>;
}
