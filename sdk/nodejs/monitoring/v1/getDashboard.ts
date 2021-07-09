// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Fetches a specific dashboard.This method requires the monitoring.dashboards.get permission on the specified dashboard. For more information, see Cloud Identity and Access Management (https://cloud.google.com/iam).
 */
export function getDashboard(args: GetDashboardArgs, opts?: pulumi.InvokeOptions): Promise<GetDashboardResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("google-native:monitoring/v1:getDashboard", {
        "dashboardId": args.dashboardId,
        "project": args.project,
    }, opts);
}

export interface GetDashboardArgs {
    dashboardId: string;
    project: string;
}

export interface GetDashboardResult {
    /**
     * The content is divided into equally spaced columns and the widgets are arranged vertically.
     */
    readonly columnLayout: outputs.monitoring.v1.ColumnLayoutResponse;
    /**
     * The mutable, human-readable name.
     */
    readonly displayName: string;
    /**
     * etag is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other. An etag is returned in the response to GetDashboard, and users are expected to put that etag in the request to UpdateDashboard to ensure that their change will be applied to the same version of the Dashboard configuration. The field should not be passed during dashboard creation.
     */
    readonly etag: string;
    /**
     * Content is arranged with a basic layout that re-flows a simple list of informational elements like widgets or tiles.
     */
    readonly gridLayout: outputs.monitoring.v1.GridLayoutResponse;
    /**
     * The content is arranged as a grid of tiles, with each content widget occupying one or more grid blocks.
     */
    readonly mosaicLayout: outputs.monitoring.v1.MosaicLayoutResponse;
    /**
     * Immutable. The resource name of the dashboard.
     */
    readonly name: string;
    /**
     * The content is divided into equally spaced rows and the widgets are arranged horizontally.
     */
    readonly rowLayout: outputs.monitoring.v1.RowLayoutResponse;
}
