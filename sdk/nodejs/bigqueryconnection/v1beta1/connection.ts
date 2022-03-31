// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates a new connection.
 */
export class Connection extends pulumi.CustomResource {
    /**
     * Get an existing Connection resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Connection {
        return new Connection(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:bigqueryconnection/v1beta1:Connection';

    /**
     * Returns true if the given object is an instance of Connection.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Connection {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Connection.__pulumiType;
    }

    /**
     * Cloud SQL properties.
     */
    public readonly cloudSql!: pulumi.Output<outputs.bigqueryconnection.v1beta1.CloudSqlPropertiesResponse>;
    /**
     * The creation timestamp of the connection.
     */
    public /*out*/ readonly creationTime!: pulumi.Output<string>;
    /**
     * User provided description.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * User provided display name for the connection.
     */
    public readonly friendlyName!: pulumi.Output<string>;
    /**
     * True, if credential is configured for this connection.
     */
    public /*out*/ readonly hasCredential!: pulumi.Output<boolean>;
    /**
     * The last update timestamp of the connection.
     */
    public /*out*/ readonly lastModifiedTime!: pulumi.Output<string>;
    /**
     * The resource name of the connection in the form of: `projects/{project_id}/locations/{location_id}/connections/{connection_id}`
     */
    public readonly name!: pulumi.Output<string>;

    /**
     * Create a Connection resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ConnectionArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            resourceInputs["cloudSql"] = args ? args.cloudSql : undefined;
            resourceInputs["connectionId"] = args ? args.connectionId : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["friendlyName"] = args ? args.friendlyName : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["creationTime"] = undefined /*out*/;
            resourceInputs["hasCredential"] = undefined /*out*/;
            resourceInputs["lastModifiedTime"] = undefined /*out*/;
        } else {
            resourceInputs["cloudSql"] = undefined /*out*/;
            resourceInputs["creationTime"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["friendlyName"] = undefined /*out*/;
            resourceInputs["hasCredential"] = undefined /*out*/;
            resourceInputs["lastModifiedTime"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Connection.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Connection resource.
 */
export interface ConnectionArgs {
    /**
     * Cloud SQL properties.
     */
    cloudSql?: pulumi.Input<inputs.bigqueryconnection.v1beta1.CloudSqlPropertiesArgs>;
    /**
     * Optional. Connection id that should be assigned to the created connection.
     */
    connectionId?: pulumi.Input<string>;
    /**
     * User provided description.
     */
    description?: pulumi.Input<string>;
    /**
     * User provided display name for the connection.
     */
    friendlyName?: pulumi.Input<string>;
    location?: pulumi.Input<string>;
    /**
     * The resource name of the connection in the form of: `projects/{project_id}/locations/{location_id}/connections/{connection_id}`
     */
    name?: pulumi.Input<string>;
    project?: pulumi.Input<string>;
}
