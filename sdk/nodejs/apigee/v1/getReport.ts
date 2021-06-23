// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Retrieve a custom report definition.
 */
export function getReport(args: GetReportArgs, opts?: pulumi.InvokeOptions): Promise<GetReportResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("google-native:apigee/v1:getReport", {
        "organizationId": args.organizationId,
        "reportId": args.reportId,
    }, opts);
}

export interface GetReportArgs {
    organizationId: string;
    reportId: string;
}

export interface GetReportResult {
    /**
     * This field contains the chart type for the report
     */
    readonly chartType: string;
    /**
     * Legacy field: not used. This field contains a list of comments associated with custom report
     */
    readonly comments: string[];
    /**
     * Unix time when the app was created json key: createdAt
     */
    readonly createdAt: string;
    /**
     * This contains the list of dimensions for the report
     */
    readonly dimensions: string[];
    /**
     * This is the display name for the report
     */
    readonly displayName: string;
    /**
     * Environment name
     */
    readonly environment: string;
    /**
     * This field contains the filter expression
     */
    readonly filter: string;
    /**
     * Legacy field: not used. Contains the from time for the report
     */
    readonly fromTime: string;
    /**
     * Modified time of this entity as milliseconds since epoch. json key: lastModifiedAt
     */
    readonly lastModifiedAt: string;
    /**
     * Last viewed time of this entity as milliseconds since epoch
     */
    readonly lastViewedAt: string;
    /**
     * Legacy field: not used This field contains the limit for the result retrieved
     */
    readonly limit: string;
    /**
     * Required. This contains the list of metrics
     */
    readonly metrics: outputs.apigee.v1.GoogleCloudApigeeV1CustomReportMetricResponse[];
    /**
     * Required. Unique identifier for the report T his is a legacy field used to encode custom report unique id
     */
    readonly name: string;
    /**
     * Legacy field: not used. This field contains the offset for the data
     */
    readonly offset: string;
    /**
     * Organization name
     */
    readonly organization: string;
    /**
     * This field contains report properties such as ui metadata etc.
     */
    readonly properties: outputs.apigee.v1.GoogleCloudApigeeV1ReportPropertyResponse[];
    /**
     * Legacy field: not used much. Contains the list of sort by columns
     */
    readonly sortByCols: string[];
    /**
     * Legacy field: not used much. Contains the sort order for the sort columns
     */
    readonly sortOrder: string;
    /**
     * Legacy field: not used. This field contains a list of tags associated with custom report
     */
    readonly tags: string[];
    /**
     * This field contains the time unit of aggregation for the report
     */
    readonly timeUnit: string;
    /**
     * Legacy field: not used. Contains the end time for the report
     */
    readonly toTime: string;
    /**
     * Legacy field: not used. This field contains the top k parameter value for restricting the result
     */
    readonly topk: string;
}