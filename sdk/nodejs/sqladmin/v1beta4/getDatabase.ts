// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Retrieves a resource containing information about a database inside a Cloud SQL instance.
 */
export function getDatabase(args: GetDatabaseArgs, opts?: pulumi.InvokeOptions): Promise<GetDatabaseResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("google-native:sqladmin/v1beta4:getDatabase", {
        "database": args.database,
        "instance": args.instance,
        "project": args.project,
    }, opts);
}

export interface GetDatabaseArgs {
    database: string;
    instance: string;
    project?: string;
}

export interface GetDatabaseResult {
    /**
     * The Cloud SQL charset value.
     */
    readonly charset: string;
    /**
     * The Cloud SQL collation value.
     */
    readonly collation: string;
    /**
     * The name of the Cloud SQL instance. This does not include the project ID.
     */
    readonly instance: string;
    /**
     * This is always **sql#database**.
     */
    readonly kind: string;
    /**
     * The name of the database in the Cloud SQL instance. This does not include the project ID or instance name.
     */
    readonly name: string;
    /**
     * The project ID of the project containing the Cloud SQL database. The Google apps domain is prefixed if applicable.
     */
    readonly project: string;
    /**
     * The URI of this resource.
     */
    readonly selfLink: string;
    readonly sqlserverDatabaseDetails: outputs.sqladmin.v1beta4.SqlServerDatabaseDetailsResponse;
}

export function getDatabaseOutput(args: GetDatabaseOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDatabaseResult> {
    return pulumi.output(args).apply(a => getDatabase(a, opts))
}

export interface GetDatabaseOutputArgs {
    database: pulumi.Input<string>;
    instance: pulumi.Input<string>;
    project?: pulumi.Input<string>;
}
