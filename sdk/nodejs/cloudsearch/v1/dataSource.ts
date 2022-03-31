// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates a datasource. **Note:** This API requires an admin account to execute.
 */
export class DataSource extends pulumi.CustomResource {
    /**
     * Get an existing DataSource resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): DataSource {
        return new DataSource(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:cloudsearch/v1:DataSource';

    /**
     * Returns true if the given object is an instance of DataSource.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DataSource {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DataSource.__pulumiType;
    }

    /**
     * If true, sets the datasource to read-only mode. In read-only mode, the Indexing API rejects any requests to index or delete items in this source. Enabling read-only mode does not stop the processing of previously accepted data.
     */
    public readonly disableModifications!: pulumi.Output<boolean>;
    /**
     * Disable serving any search or assist results.
     */
    public readonly disableServing!: pulumi.Output<boolean>;
    /**
     * Display name of the datasource The maximum length is 300 characters.
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * List of service accounts that have indexing access.
     */
    public readonly indexingServiceAccounts!: pulumi.Output<string[]>;
    /**
     * This field restricts visibility to items at the datasource level. Items within the datasource are restricted to the union of users and groups included in this field. Note that, this does not ensure access to a specific item, as users need to have ACL permissions on the contained items. This ensures a high level access on the entire datasource, and that the individual items are not shared outside this visibility.
     */
    public readonly itemsVisibility!: pulumi.Output<outputs.cloudsearch.v1.GSuitePrincipalResponse[]>;
    /**
     * Name of the datasource resource. Format: datasources/{source_id}. The name is ignored when creating a datasource.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * IDs of the Long Running Operations (LROs) currently running for this schema.
     */
    public readonly operationIds!: pulumi.Output<string[]>;
    /**
     * Can a user request to get thumbnail URI for Items indexed in this data source.
     */
    public readonly returnThumbnailUrls!: pulumi.Output<boolean>;
    /**
     * A short name or alias for the source. This value will be used to match the 'source' operator. For example, if the short name is *<value>* then queries like *source:<value>* will only return results for this source. The value must be unique across all datasources. The value must only contain alphanumeric characters (a-zA-Z0-9). The value cannot start with 'google' and cannot be one of the following: mail, gmail, docs, drive, groups, sites, calendar, hangouts, gplus, keep, people, teams. Its maximum length is 32 characters.
     */
    public readonly shortName!: pulumi.Output<string>;

    /**
     * Create a DataSource resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DataSourceArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.displayName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'displayName'");
            }
            resourceInputs["disableModifications"] = args ? args.disableModifications : undefined;
            resourceInputs["disableServing"] = args ? args.disableServing : undefined;
            resourceInputs["displayName"] = args ? args.displayName : undefined;
            resourceInputs["indexingServiceAccounts"] = args ? args.indexingServiceAccounts : undefined;
            resourceInputs["itemsVisibility"] = args ? args.itemsVisibility : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["operationIds"] = args ? args.operationIds : undefined;
            resourceInputs["returnThumbnailUrls"] = args ? args.returnThumbnailUrls : undefined;
            resourceInputs["shortName"] = args ? args.shortName : undefined;
        } else {
            resourceInputs["disableModifications"] = undefined /*out*/;
            resourceInputs["disableServing"] = undefined /*out*/;
            resourceInputs["displayName"] = undefined /*out*/;
            resourceInputs["indexingServiceAccounts"] = undefined /*out*/;
            resourceInputs["itemsVisibility"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["operationIds"] = undefined /*out*/;
            resourceInputs["returnThumbnailUrls"] = undefined /*out*/;
            resourceInputs["shortName"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(DataSource.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a DataSource resource.
 */
export interface DataSourceArgs {
    /**
     * If true, sets the datasource to read-only mode. In read-only mode, the Indexing API rejects any requests to index or delete items in this source. Enabling read-only mode does not stop the processing of previously accepted data.
     */
    disableModifications?: pulumi.Input<boolean>;
    /**
     * Disable serving any search or assist results.
     */
    disableServing?: pulumi.Input<boolean>;
    /**
     * Display name of the datasource The maximum length is 300 characters.
     */
    displayName: pulumi.Input<string>;
    /**
     * List of service accounts that have indexing access.
     */
    indexingServiceAccounts?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * This field restricts visibility to items at the datasource level. Items within the datasource are restricted to the union of users and groups included in this field. Note that, this does not ensure access to a specific item, as users need to have ACL permissions on the contained items. This ensures a high level access on the entire datasource, and that the individual items are not shared outside this visibility.
     */
    itemsVisibility?: pulumi.Input<pulumi.Input<inputs.cloudsearch.v1.GSuitePrincipalArgs>[]>;
    /**
     * Name of the datasource resource. Format: datasources/{source_id}. The name is ignored when creating a datasource.
     */
    name?: pulumi.Input<string>;
    /**
     * IDs of the Long Running Operations (LROs) currently running for this schema.
     */
    operationIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Can a user request to get thumbnail URI for Items indexed in this data source.
     */
    returnThumbnailUrls?: pulumi.Input<boolean>;
    /**
     * A short name or alias for the source. This value will be used to match the 'source' operator. For example, if the short name is *<value>* then queries like *source:<value>* will only return results for this source. The value must be unique across all datasources. The value must only contain alphanumeric characters (a-zA-Z0-9). The value cannot start with 'google' and cannot be one of the following: mail, gmail, docs, drive, groups, sites, calendar, hangouts, gplus, keep, people, teams. Its maximum length is 32 characters.
     */
    shortName?: pulumi.Input<string>;
}
