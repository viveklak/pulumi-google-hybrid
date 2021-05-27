// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

/**
 * Submit a query to be processed in the background. If the submission of the query succeeds, the API returns a 201 status and an ID that refer to the query. In addition to the HTTP status 201, the `state` of "enqueued" means that the request succeeded.
 */
export class OrganizationEnvironmentQuery extends pulumi.CustomResource {
    /**
     * Get an existing OrganizationEnvironmentQuery resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): OrganizationEnvironmentQuery {
        return new OrganizationEnvironmentQuery(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:apigee/v1:OrganizationEnvironmentQuery';

    /**
     * Returns true if the given object is an instance of OrganizationEnvironmentQuery.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is OrganizationEnvironmentQuery {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === OrganizationEnvironmentQuery.__pulumiType;
    }

    /**
     * Creation time of the query.
     */
    public /*out*/ readonly created!: pulumi.Output<string>;
    /**
     * Hostname is available only when query is executed at host level.
     */
    public readonly envgroupHostname!: pulumi.Output<string>;
    /**
     * Error is set when query fails.
     */
    public /*out*/ readonly error!: pulumi.Output<string>;
    /**
     * ExecutionTime is available only after the query is completed.
     */
    public /*out*/ readonly executionTime!: pulumi.Output<string>;
    /**
     * Asynchronous Query Name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Contains information like metrics, dimenstions etc of the AsyncQuery.
     */
    public /*out*/ readonly queryParams!: pulumi.Output<outputs.apigee.v1.GoogleCloudApigeeV1QueryMetadataResponse>;
    /**
     * Asynchronous Report ID.
     */
    public readonly reportDefinitionId!: pulumi.Output<string>;
    /**
     * Result is available only after the query is completed.
     */
    public /*out*/ readonly result!: pulumi.Output<outputs.apigee.v1.GoogleCloudApigeeV1AsyncQueryResultResponse>;
    /**
     * ResultFileSize is available only after the query is completed.
     */
    public /*out*/ readonly resultFileSize!: pulumi.Output<string>;
    /**
     * ResultRows is available only after the query is completed.
     */
    public /*out*/ readonly resultRows!: pulumi.Output<string>;
    /**
     * Self link of the query. Example: `/organizations/myorg/environments/myenv/queries/9cfc0d85-0f30-46d6-ae6f-318d0cb961bd` or following format if query is running at host level: `/organizations/myorg/hostQueries/9cfc0d85-0f30-46d6-ae6f-318d0cb961bd`
     */
    public /*out*/ readonly self!: pulumi.Output<string>;
    /**
     * Query state could be "enqueued", "running", "completed", "failed".
     */
    public /*out*/ readonly state!: pulumi.Output<string>;
    /**
     * Last updated timestamp for the query.
     */
    public /*out*/ readonly updated!: pulumi.Output<string>;

    /**
     * Create a OrganizationEnvironmentQuery resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: OrganizationEnvironmentQueryArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.environmentId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'environmentId'");
            }
            if ((!args || args.organizationId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'organizationId'");
            }
            inputs["csvDelimiter"] = args ? args.csvDelimiter : undefined;
            inputs["dimensions"] = args ? args.dimensions : undefined;
            inputs["envgroupHostname"] = args ? args.envgroupHostname : undefined;
            inputs["environmentId"] = args ? args.environmentId : undefined;
            inputs["filter"] = args ? args.filter : undefined;
            inputs["groupByTimeUnit"] = args ? args.groupByTimeUnit : undefined;
            inputs["limit"] = args ? args.limit : undefined;
            inputs["metrics"] = args ? args.metrics : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["organizationId"] = args ? args.organizationId : undefined;
            inputs["outputFormat"] = args ? args.outputFormat : undefined;
            inputs["reportDefinitionId"] = args ? args.reportDefinitionId : undefined;
            inputs["timeRange"] = args ? args.timeRange : undefined;
            inputs["created"] = undefined /*out*/;
            inputs["error"] = undefined /*out*/;
            inputs["executionTime"] = undefined /*out*/;
            inputs["queryParams"] = undefined /*out*/;
            inputs["result"] = undefined /*out*/;
            inputs["resultFileSize"] = undefined /*out*/;
            inputs["resultRows"] = undefined /*out*/;
            inputs["self"] = undefined /*out*/;
            inputs["state"] = undefined /*out*/;
            inputs["updated"] = undefined /*out*/;
        } else {
            inputs["created"] = undefined /*out*/;
            inputs["envgroupHostname"] = undefined /*out*/;
            inputs["error"] = undefined /*out*/;
            inputs["executionTime"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["queryParams"] = undefined /*out*/;
            inputs["reportDefinitionId"] = undefined /*out*/;
            inputs["result"] = undefined /*out*/;
            inputs["resultFileSize"] = undefined /*out*/;
            inputs["resultRows"] = undefined /*out*/;
            inputs["self"] = undefined /*out*/;
            inputs["state"] = undefined /*out*/;
            inputs["updated"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(OrganizationEnvironmentQuery.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a OrganizationEnvironmentQuery resource.
 */
export interface OrganizationEnvironmentQueryArgs {
    /**
     * Delimiter used in the CSV file, if `outputFormat` is set to `csv`. Defaults to the `,` (comma) character. Supported delimiter characters include comma (`,`), pipe (`|`), and tab (`\t`).
     */
    readonly csvDelimiter?: pulumi.Input<string>;
    /**
     * A list of dimensions. https://docs.apigee.com/api-platform/analytics/analytics-reference#dimensions
     */
    readonly dimensions?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Hostname needs to be specified if query intends to run at host level. This field is only allowed when query is submitted by CreateHostAsyncQuery where analytics data will be grouped by organization and hostname.
     */
    readonly envgroupHostname?: pulumi.Input<string>;
    readonly environmentId: pulumi.Input<string>;
    /**
     * Boolean expression that can be used to filter data. Filter expressions can be combined using AND/OR terms and should be fully parenthesized to avoid ambiguity. See Analytics metrics, dimensions, and filters reference https://docs.apigee.com/api-platform/analytics/analytics-reference for more information on the fields available to filter on. For more information on the tokens that you use to build filter expressions, see Filter expression syntax. https://docs.apigee.com/api-platform/analytics/asynch-reports-api#filter-expression-syntax
     */
    readonly filter?: pulumi.Input<string>;
    /**
     * Time unit used to group the result set. Valid values include: second, minute, hour, day, week, or month. If a query includes groupByTimeUnit, then the result is an aggregation based on the specified time unit and the resultant timestamp does not include milliseconds precision. If a query omits groupByTimeUnit, then the resultant timestamp includes milliseconds precision.
     */
    readonly groupByTimeUnit?: pulumi.Input<string>;
    /**
     * Maximum number of rows that can be returned in the result.
     */
    readonly limit?: pulumi.Input<number>;
    /**
     * A list of Metrics.
     */
    readonly metrics?: pulumi.Input<pulumi.Input<inputs.apigee.v1.GoogleCloudApigeeV1QueryMetricArgs>[]>;
    /**
     * Asynchronous Query Name.
     */
    readonly name?: pulumi.Input<string>;
    readonly organizationId: pulumi.Input<string>;
    /**
     * Valid values include: `csv` or `json`. Defaults to `json`. Note: Configure the delimiter for CSV output using the csvDelimiter property.
     */
    readonly outputFormat?: pulumi.Input<string>;
    /**
     * Asynchronous Report ID.
     */
    readonly reportDefinitionId?: pulumi.Input<string>;
    /**
     * Required. Time range for the query. Can use the following predefined strings to specify the time range: `last60minutes` `last24hours` `last7days` Or, specify the timeRange as a structure describing start and end timestamps in the ISO format: yyyy-mm-ddThh:mm:ssZ. Example: "timeRange": { "start": "2018-07-29T00:13:00Z", "end": "2018-08-01T00:18:00Z" }
     */
    readonly timeRange?: any;
}
