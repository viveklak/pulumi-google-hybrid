// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Inserts a resource containing information about a database inside a Cloud SQL instance.
 */
export class Database extends pulumi.CustomResource {
    /**
     * Get an existing Database resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Database {
        return new Database(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:sqladmin/v1:Database';

    /**
     * Returns true if the given object is an instance of Database.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Database {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Database.__pulumiType;
    }

    /**
     * The Cloud SQL charset value.
     */
    public readonly charset!: pulumi.Output<string>;
    /**
     * The Cloud SQL collation value.
     */
    public readonly collation!: pulumi.Output<string>;
    /**
     * This field is deprecated and will be removed from a future version of the API.
     *
     * @deprecated This field is deprecated and will be removed from a future version of the API.
     */
    public readonly etag!: pulumi.Output<string>;
    /**
     * The name of the Cloud SQL instance. This does not include the project ID.
     */
    public readonly instance!: pulumi.Output<string>;
    /**
     * This is always `sql#database`.
     */
    public readonly kind!: pulumi.Output<string>;
    /**
     * The name of the database in the Cloud SQL instance. This does not include the project ID or instance name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The project ID of the project containing the Cloud SQL database. The Google apps domain is prefixed if applicable.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * The URI of this resource.
     */
    public readonly selfLink!: pulumi.Output<string>;
    public readonly sqlserverDatabaseDetails!: pulumi.Output<outputs.sqladmin.v1.SqlServerDatabaseDetailsResponse>;

    /**
     * Create a Database resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DatabaseArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.instance === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instance'");
            }
            resourceInputs["charset"] = args ? args.charset : undefined;
            resourceInputs["collation"] = args ? args.collation : undefined;
            resourceInputs["etag"] = args ? args.etag : undefined;
            resourceInputs["instance"] = args ? args.instance : undefined;
            resourceInputs["kind"] = args ? args.kind : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["selfLink"] = args ? args.selfLink : undefined;
            resourceInputs["sqlserverDatabaseDetails"] = args ? args.sqlserverDatabaseDetails : undefined;
        } else {
            resourceInputs["charset"] = undefined /*out*/;
            resourceInputs["collation"] = undefined /*out*/;
            resourceInputs["etag"] = undefined /*out*/;
            resourceInputs["instance"] = undefined /*out*/;
            resourceInputs["kind"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["project"] = undefined /*out*/;
            resourceInputs["selfLink"] = undefined /*out*/;
            resourceInputs["sqlserverDatabaseDetails"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Database.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Database resource.
 */
export interface DatabaseArgs {
    /**
     * The Cloud SQL charset value.
     */
    charset?: pulumi.Input<string>;
    /**
     * The Cloud SQL collation value.
     */
    collation?: pulumi.Input<string>;
    /**
     * This field is deprecated and will be removed from a future version of the API.
     *
     * @deprecated This field is deprecated and will be removed from a future version of the API.
     */
    etag?: pulumi.Input<string>;
    /**
     * The name of the Cloud SQL instance. This does not include the project ID.
     */
    instance: pulumi.Input<string>;
    /**
     * This is always `sql#database`.
     */
    kind?: pulumi.Input<string>;
    /**
     * The name of the database in the Cloud SQL instance. This does not include the project ID or instance name.
     */
    name?: pulumi.Input<string>;
    /**
     * The project ID of the project containing the Cloud SQL database. The Google apps domain is prefixed if applicable.
     */
    project?: pulumi.Input<string>;
    /**
     * The URI of this resource.
     */
    selfLink?: pulumi.Input<string>;
    sqlserverDatabaseDetails?: pulumi.Input<inputs.sqladmin.v1.SqlServerDatabaseDetailsArgs>;
}
