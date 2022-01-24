// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Gets a sink.
 */
export function getBillingAccountSink(args: GetBillingAccountSinkArgs, opts?: pulumi.InvokeOptions): Promise<GetBillingAccountSinkResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("google-native:logging/v2:getBillingAccountSink", {
        "billingAccountId": args.billingAccountId,
        "sinkId": args.sinkId,
    }, opts);
}

export interface GetBillingAccountSinkArgs {
    billingAccountId: string;
    sinkId: string;
}

export interface GetBillingAccountSinkResult {
    /**
     * Optional. Options that affect sinks exporting data to BigQuery.
     */
    readonly bigqueryOptions: outputs.logging.v2.BigQueryOptionsResponse;
    /**
     * The creation timestamp of the sink.This field may not be present for older sinks.
     */
    readonly createTime: string;
    /**
     * Optional. A description of this sink.The maximum length of the description is 8000 characters.
     */
    readonly description: string;
    /**
     * The export destination: "storage.googleapis.com/[GCS_BUCKET]" "bigquery.googleapis.com/projects/[PROJECT_ID]/datasets/[DATASET]" "pubsub.googleapis.com/projects/[PROJECT_ID]/topics/[TOPIC_ID]" The sink's writer_identity, set when the sink is created, must have permission to write to the destination or else the log entries are not exported. For more information, see Exporting Logs with Sinks (https://cloud.google.com/logging/docs/api/tasks/exporting-logs).
     */
    readonly destination: string;
    /**
     * Optional. If set to true, then this sink is disabled and it does not export any log entries.
     */
    readonly disabled: boolean;
    /**
     * Optional. Log entries that match any of these exclusion filters will not be exported.If a log entry is matched by both filter and one of exclusion_filters it will not be exported.
     */
    readonly exclusions: outputs.logging.v2.LogExclusionResponse[];
    /**
     * Optional. An advanced logs filter (https://cloud.google.com/logging/docs/view/advanced-queries). The only exported log entries are those that are in the resource owning the sink and that match the filter.For example:logName="projects/[PROJECT_ID]/logs/[LOG_ID]" AND severity>=ERROR
     */
    readonly filter: string;
    /**
     * Optional. This field applies only to sinks owned by organizations and folders. If the field is false, the default, only the logs owned by the sink's parent resource are available for export. If the field is true, then log entries from all the projects, folders, and billing accounts contained in the sink's parent resource are also available for export. Whether a particular log entry from the children is exported depends on the sink's filter expression.For example, if this field is true, then the filter resource.type=gce_instance would export all Compute Engine VM instance log entries from all projects in the sink's parent.To only export entries from certain child projects, filter on the project part of the log name:logName:("projects/test-project1/" OR "projects/test-project2/") AND resource.type=gce_instance
     */
    readonly includeChildren: boolean;
    /**
     * The client-assigned sink identifier, unique within the project.For example: "my-syslog-errors-to-pubsub". Sink identifiers are limited to 100 characters and can include only the following characters: upper and lower-case alphanumeric characters, underscores, hyphens, and periods. First character has to be alphanumeric.
     */
    readonly name: string;
    /**
     * The last update timestamp of the sink.This field may not be present for older sinks.
     */
    readonly updateTime: string;
    /**
     * An IAM identity—a service account or group—under which Cloud Logging writes the exported log entries to the sink's destination. This field is set by sinks.create and sinks.update based on the value of unique_writer_identity in those methods.Until you grant this identity write-access to the destination, log entry exports from this sink will fail. For more information, see Granting Access for a Resource (https://cloud.google.com/iam/docs/granting-roles-to-service-accounts#granting_access_to_a_service_account_for_a_resource). Consult the destination service's documentation to determine the appropriate IAM roles to assign to the identity.Sinks that have a destination that is a log bucket in the same project as the sink do not have a writer_identity and no additional permissions are required.
     */
    readonly writerIdentity: string;
}

export function getBillingAccountSinkOutput(args: GetBillingAccountSinkOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetBillingAccountSinkResult> {
    return pulumi.output(args).apply(a => getBillingAccountSink(a, opts))
}

export interface GetBillingAccountSinkOutputArgs {
    billingAccountId: pulumi.Input<string>;
    sinkId: pulumi.Input<string>;
}
