// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates a composite index. This returns a google.longrunning.Operation which may be used to track the status of the creation. The metadata for the operation will be the type IndexOperationMetadata.
 */
export class DatabaseCollectionGroupIndex extends pulumi.CustomResource {
    /**
     * Get an existing DatabaseCollectionGroupIndex resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): DatabaseCollectionGroupIndex {
        return new DatabaseCollectionGroupIndex(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:firestore/v1beta2:DatabaseCollectionGroupIndex';

    /**
     * Returns true if the given object is an instance of DatabaseCollectionGroupIndex.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DatabaseCollectionGroupIndex {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DatabaseCollectionGroupIndex.__pulumiType;
    }

    /**
     * The fields supported by this index. For composite indexes, this is always 2 or more fields. The last field entry is always for the field path `__name__`. If, on creation, `__name__` was not specified as the last field, it will be added automatically with the same direction as that of the last field defined. If the final field in a composite index is not directional, the `__name__` will be ordered ASCENDING (unless explicitly specified). For single field indexes, this will always be exactly one entry with a field path equal to the field path of the associated field.
     */
    public readonly fields!: pulumi.Output<outputs.firestore.v1beta2.GoogleFirestoreAdminV1beta2IndexFieldResponse[]>;
    /**
     * A server defined name for this index. The form of this name for composite indexes will be: `projects/{project_id}/databases/{database_id}/collectionGroups/{collection_id}/indexes/{composite_index_id}` For single field indexes, this field will be empty.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Indexes with a collection query scope specified allow queries against a collection that is the child of a specific document, specified at query time, and that has the same collection id. Indexes with a collection group query scope specified allow queries against all collections descended from a specific document, specified at query time, and that have the same collection id as this index.
     */
    public readonly queryScope!: pulumi.Output<string>;
    /**
     * The serving state of the index.
     */
    public readonly state!: pulumi.Output<string>;

    /**
     * Create a DatabaseCollectionGroupIndex resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DatabaseCollectionGroupIndexArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.collectionGroupId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'collectionGroupId'");
            }
            if ((!args || args.databaseId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'databaseId'");
            }
            if ((!args || args.project === undefined) && !opts.urn) {
                throw new Error("Missing required property 'project'");
            }
            inputs["collectionGroupId"] = args ? args.collectionGroupId : undefined;
            inputs["databaseId"] = args ? args.databaseId : undefined;
            inputs["fields"] = args ? args.fields : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["queryScope"] = args ? args.queryScope : undefined;
            inputs["state"] = args ? args.state : undefined;
        } else {
            inputs["fields"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["queryScope"] = undefined /*out*/;
            inputs["state"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(DatabaseCollectionGroupIndex.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a DatabaseCollectionGroupIndex resource.
 */
export interface DatabaseCollectionGroupIndexArgs {
    readonly collectionGroupId: pulumi.Input<string>;
    readonly databaseId: pulumi.Input<string>;
    /**
     * The fields supported by this index. For composite indexes, this is always 2 or more fields. The last field entry is always for the field path `__name__`. If, on creation, `__name__` was not specified as the last field, it will be added automatically with the same direction as that of the last field defined. If the final field in a composite index is not directional, the `__name__` will be ordered ASCENDING (unless explicitly specified). For single field indexes, this will always be exactly one entry with a field path equal to the field path of the associated field.
     */
    readonly fields?: pulumi.Input<pulumi.Input<inputs.firestore.v1beta2.GoogleFirestoreAdminV1beta2IndexFieldArgs>[]>;
    /**
     * A server defined name for this index. The form of this name for composite indexes will be: `projects/{project_id}/databases/{database_id}/collectionGroups/{collection_id}/indexes/{composite_index_id}` For single field indexes, this field will be empty.
     */
    readonly name?: pulumi.Input<string>;
    readonly project: pulumi.Input<string>;
    /**
     * Indexes with a collection query scope specified allow queries against a collection that is the child of a specific document, specified at query time, and that has the same collection id. Indexes with a collection group query scope specified allow queries against all collections descended from a specific document, specified at query time, and that have the same collection id as this index.
     */
    readonly queryScope?: pulumi.Input<string>;
    /**
     * The serving state of the index.
     */
    readonly state?: pulumi.Input<string>;
}
