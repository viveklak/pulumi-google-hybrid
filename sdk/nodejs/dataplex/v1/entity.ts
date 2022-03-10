// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Create a metadata entity.
 * Auto-naming is currently not supported for this resource.
 */
export class Entity extends pulumi.CustomResource {
    /**
     * Get an existing Entity resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Entity {
        return new Entity(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:dataplex/v1:Entity';

    /**
     * Returns true if the given object is an instance of Entity.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Entity {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Entity.__pulumiType;
    }

    /**
     * Immutable. The ID of the asset associated with the storage location containing the entity data. The entity must be with in the same zone with the asset.
     */
    public readonly asset!: pulumi.Output<string>;
    /**
     * The name of the associated Data Catalog entry.
     */
    public /*out*/ readonly catalogEntry!: pulumi.Output<string>;
    /**
     * Metadata stores that the entity is compatible with.
     */
    public /*out*/ readonly compatibility!: pulumi.Output<outputs.dataplex.v1.GoogleCloudDataplexV1EntityCompatibilityStatusResponse>;
    /**
     * The time when the entity was created.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * Immutable. The storage path of the entity data. For Cloud Storage data, this is the fully-qualified path to the entity, such as gs://bucket/path/to/data. For BigQuery data, this is the name of the table resource, such as projects/project_id/datasets/dataset_id/tables/table_id.
     */
    public readonly dataPath!: pulumi.Output<string>;
    /**
     * Optional. The set of items within the data path constituting the data in the entity, represented as a glob path. Example: gs://bucket/path/to/data/**&#47;*.csv.
     */
    public readonly dataPathPattern!: pulumi.Output<string>;
    /**
     * Optional. User friendly longer description text. Must be shorter than or equal to 1024 characters.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * Optional. Display name must be shorter than or equal to 256 characters.
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * Optional. The etag associated with the entity, which can be retrieved with a GetEntity request. Required for update and delete requests.
     */
    public readonly etag!: pulumi.Output<string>;
    /**
     * Identifies the storage format of the entity data. It does not apply to entities with data stored in BigQuery.
     */
    public readonly format!: pulumi.Output<outputs.dataplex.v1.GoogleCloudDataplexV1StorageFormatResponse>;
    /**
     * The resource name of the entity, of the form: projects/{project_number}/locations/{location_id}/lakes/{lake_id}/zones/{zone_id}/entities/{id}.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The description of the data structure and layout. The schema is not included in list responses. It is only included in SCHEMA and FULL entity views of a GetEntity response.
     */
    public readonly schema!: pulumi.Output<outputs.dataplex.v1.GoogleCloudDataplexV1SchemaResponse>;
    /**
     * Immutable. Identifies the storage system of the entity data.
     */
    public readonly system!: pulumi.Output<string>;
    /**
     * Immutable. The type of entity.
     */
    public readonly type!: pulumi.Output<string>;
    /**
     * The time when the entity was last updated.
     */
    public /*out*/ readonly updateTime!: pulumi.Output<string>;

    /**
     * Create a Entity resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: EntityArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.asset === undefined) && !opts.urn) {
                throw new Error("Missing required property 'asset'");
            }
            if ((!args || args.dataPath === undefined) && !opts.urn) {
                throw new Error("Missing required property 'dataPath'");
            }
            if ((!args || args.format === undefined) && !opts.urn) {
                throw new Error("Missing required property 'format'");
            }
            if ((!args || args.id === undefined) && !opts.urn) {
                throw new Error("Missing required property 'id'");
            }
            if ((!args || args.lakeId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'lakeId'");
            }
            if ((!args || args.schema === undefined) && !opts.urn) {
                throw new Error("Missing required property 'schema'");
            }
            if ((!args || args.system === undefined) && !opts.urn) {
                throw new Error("Missing required property 'system'");
            }
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            resourceInputs["asset"] = args ? args.asset : undefined;
            resourceInputs["dataPath"] = args ? args.dataPath : undefined;
            resourceInputs["dataPathPattern"] = args ? args.dataPathPattern : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["displayName"] = args ? args.displayName : undefined;
            resourceInputs["etag"] = args ? args.etag : undefined;
            resourceInputs["format"] = args ? args.format : undefined;
            resourceInputs["id"] = args ? args.id : undefined;
            resourceInputs["lakeId"] = args ? args.lakeId : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["schema"] = args ? args.schema : undefined;
            resourceInputs["system"] = args ? args.system : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["validateOnly"] = args ? args.validateOnly : undefined;
            resourceInputs["zone"] = args ? args.zone : undefined;
            resourceInputs["catalogEntry"] = undefined /*out*/;
            resourceInputs["compatibility"] = undefined /*out*/;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        } else {
            resourceInputs["asset"] = undefined /*out*/;
            resourceInputs["catalogEntry"] = undefined /*out*/;
            resourceInputs["compatibility"] = undefined /*out*/;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["dataPath"] = undefined /*out*/;
            resourceInputs["dataPathPattern"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["displayName"] = undefined /*out*/;
            resourceInputs["etag"] = undefined /*out*/;
            resourceInputs["format"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["schema"] = undefined /*out*/;
            resourceInputs["system"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Entity.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Entity resource.
 */
export interface EntityArgs {
    /**
     * Immutable. The ID of the asset associated with the storage location containing the entity data. The entity must be with in the same zone with the asset.
     */
    asset: pulumi.Input<string>;
    /**
     * Immutable. The storage path of the entity data. For Cloud Storage data, this is the fully-qualified path to the entity, such as gs://bucket/path/to/data. For BigQuery data, this is the name of the table resource, such as projects/project_id/datasets/dataset_id/tables/table_id.
     */
    dataPath: pulumi.Input<string>;
    /**
     * Optional. The set of items within the data path constituting the data in the entity, represented as a glob path. Example: gs://bucket/path/to/data/**&#47;*.csv.
     */
    dataPathPattern?: pulumi.Input<string>;
    /**
     * Optional. User friendly longer description text. Must be shorter than or equal to 1024 characters.
     */
    description?: pulumi.Input<string>;
    /**
     * Optional. Display name must be shorter than or equal to 256 characters.
     */
    displayName?: pulumi.Input<string>;
    /**
     * Optional. The etag associated with the entity, which can be retrieved with a GetEntity request. Required for update and delete requests.
     */
    etag?: pulumi.Input<string>;
    /**
     * Identifies the storage format of the entity data. It does not apply to entities with data stored in BigQuery.
     */
    format: pulumi.Input<inputs.dataplex.v1.GoogleCloudDataplexV1StorageFormatArgs>;
    /**
     * A user-provided entity ID. It is mutable, and will be used as the published table name. Specifying a new ID in an update entity request will override the existing value. The ID must contain only letters (a-z, A-Z), numbers (0-9), and underscores. Must begin with a letter and consist of 256 or fewer characters.
     */
    id: pulumi.Input<string>;
    lakeId: pulumi.Input<string>;
    location?: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    /**
     * The description of the data structure and layout. The schema is not included in list responses. It is only included in SCHEMA and FULL entity views of a GetEntity response.
     */
    schema: pulumi.Input<inputs.dataplex.v1.GoogleCloudDataplexV1SchemaArgs>;
    /**
     * Immutable. Identifies the storage system of the entity data.
     */
    system: pulumi.Input<enums.dataplex.v1.EntitySystem>;
    /**
     * Immutable. The type of entity.
     */
    type: pulumi.Input<enums.dataplex.v1.EntityType>;
    validateOnly?: pulumi.Input<string>;
    zone?: pulumi.Input<string>;
}