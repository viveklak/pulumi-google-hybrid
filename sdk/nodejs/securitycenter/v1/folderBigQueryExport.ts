// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Creates a big query export.
 * Auto-naming is currently not supported for this resource.
 */
export class FolderBigQueryExport extends pulumi.CustomResource {
    /**
     * Get an existing FolderBigQueryExport resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): FolderBigQueryExport {
        return new FolderBigQueryExport(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:securitycenter/v1:FolderBigQueryExport';

    /**
     * Returns true if the given object is an instance of FolderBigQueryExport.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FolderBigQueryExport {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FolderBigQueryExport.__pulumiType;
    }

    /**
     * The time at which the big query export was created. This field is set by the server and will be ignored if provided on export on creation.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * The dataset to write findings' updates to. Its format is "projects/[project_id]/datasets/[bigquery_dataset_id]". BigQuery Dataset unique ID must contain only letters (a-z, A-Z), numbers (0-9), or underscores (_).
     */
    public readonly dataset!: pulumi.Output<string>;
    /**
     * The description of the export (max of 1024 characters).
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * Expression that defines the filter to apply across create/update events of findings. The expression is a list of zero or more restrictions combined via logical operators `AND` and `OR`. Parentheses are supported, and `OR` has higher precedence than `AND`. Restrictions have the form ` ` and may have a `-` character in front of them to indicate negation. The fields map to those defined in the corresponding resource. The supported operators are: * `=` for all value types. * `>`, `<`, `>=`, `<=` for integer values. * `:`, meaning substring matching, for strings. The supported value types are: * string literals in quotes. * integer literals without quotes. * boolean literals `true` and `false` without quotes. Please see the proto documentation in the finding (https://source.corp.google.com/piper///depot/google3/google/cloud/securitycenter/v1/finding.proto) and in the ListFindingsRequest for valid filter syntax. (https://source.corp.google.com/piper///depot/google3/google/cloud/securitycenter/v1/securitycenter_service.proto).
     */
    public readonly filter!: pulumi.Output<string>;
    /**
     * Email address of the user who last edited the big query export. This field is set by the server and will be ignored if provided on export creation or update.
     */
    public /*out*/ readonly mostRecentEditor!: pulumi.Output<string>;
    /**
     * The relative resource name of this export. See: https://cloud.google.com/apis/design/resource_names#relative_resource_name. Example format: "organizations/{organization_id}/bigQueryExports/{export_id}" Example format: "folders/{folder_id}/bigQueryExports/{export_id}" Example format: "projects/{project_id}/bigQueryExports/{export_id}" This field is provided in responses, and is ignored when provided in create requests.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The service account that needs permission to create table, upload data to the big query dataset.
     */
    public /*out*/ readonly principal!: pulumi.Output<string>;
    /**
     * The most recent time at which the big export was updated. This field is set by the server and will be ignored if provided on export creation or update.
     */
    public /*out*/ readonly updateTime!: pulumi.Output<string>;

    /**
     * Create a FolderBigQueryExport resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FolderBigQueryExportArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.bigQueryExportId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'bigQueryExportId'");
            }
            if ((!args || args.folderId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'folderId'");
            }
            resourceInputs["bigQueryExportId"] = args ? args.bigQueryExportId : undefined;
            resourceInputs["dataset"] = args ? args.dataset : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["filter"] = args ? args.filter : undefined;
            resourceInputs["folderId"] = args ? args.folderId : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["mostRecentEditor"] = undefined /*out*/;
            resourceInputs["principal"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        } else {
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["dataset"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["filter"] = undefined /*out*/;
            resourceInputs["mostRecentEditor"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["principal"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(FolderBigQueryExport.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a FolderBigQueryExport resource.
 */
export interface FolderBigQueryExportArgs {
    bigQueryExportId: pulumi.Input<string>;
    /**
     * The dataset to write findings' updates to. Its format is "projects/[project_id]/datasets/[bigquery_dataset_id]". BigQuery Dataset unique ID must contain only letters (a-z, A-Z), numbers (0-9), or underscores (_).
     */
    dataset?: pulumi.Input<string>;
    /**
     * The description of the export (max of 1024 characters).
     */
    description?: pulumi.Input<string>;
    /**
     * Expression that defines the filter to apply across create/update events of findings. The expression is a list of zero or more restrictions combined via logical operators `AND` and `OR`. Parentheses are supported, and `OR` has higher precedence than `AND`. Restrictions have the form ` ` and may have a `-` character in front of them to indicate negation. The fields map to those defined in the corresponding resource. The supported operators are: * `=` for all value types. * `>`, `<`, `>=`, `<=` for integer values. * `:`, meaning substring matching, for strings. The supported value types are: * string literals in quotes. * integer literals without quotes. * boolean literals `true` and `false` without quotes. Please see the proto documentation in the finding (https://source.corp.google.com/piper///depot/google3/google/cloud/securitycenter/v1/finding.proto) and in the ListFindingsRequest for valid filter syntax. (https://source.corp.google.com/piper///depot/google3/google/cloud/securitycenter/v1/securitycenter_service.proto).
     */
    filter?: pulumi.Input<string>;
    folderId: pulumi.Input<string>;
    /**
     * The relative resource name of this export. See: https://cloud.google.com/apis/design/resource_names#relative_resource_name. Example format: "organizations/{organization_id}/bigQueryExports/{export_id}" Example format: "folders/{folder_id}/bigQueryExports/{export_id}" Example format: "projects/{project_id}/bigQueryExports/{export_id}" This field is provided in responses, and is ignored when provided in create requests.
     */
    name?: pulumi.Input<string>;
}