// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates a search application. **Note:** This API requires an admin account to execute.
 */
export class SettingSearchapplication extends pulumi.CustomResource {
    /**
     * Get an existing SettingSearchapplication resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): SettingSearchapplication {
        return new SettingSearchapplication(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-cloud:cloudsearch/v1:SettingSearchapplication';

    /**
     * Returns true if the given object is an instance of SettingSearchapplication.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SettingSearchapplication {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SettingSearchapplication.__pulumiType;
    }

    /**
     * Retrictions applied to the configurations. The maximum number of elements is 10.
     */
    public readonly dataSourceRestrictions!: pulumi.Output<outputs.cloudsearch.v1.DataSourceRestrictionResponse[]>;
    /**
     * The default fields for returning facet results. The sources specified here also have been included in data_source_restrictions above.
     */
    public readonly defaultFacetOptions!: pulumi.Output<outputs.cloudsearch.v1.FacetOptionsResponse[]>;
    /**
     * The default options for sorting the search results
     */
    public readonly defaultSortOptions!: pulumi.Output<outputs.cloudsearch.v1.SortOptionsResponse>;
    /**
     * Display name of the Search Application. The maximum length is 300 characters.
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * Name of the Search Application. Format: searchapplications/{application_id}.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * IDs of the Long Running Operations (LROs) currently running for this schema. Output only field.
     */
    public /*out*/ readonly operationIds!: pulumi.Output<string[]>;
    /**
     * Configuration for ranking results.
     */
    public readonly scoringConfig!: pulumi.Output<outputs.cloudsearch.v1.ScoringConfigResponse>;
    /**
     * Configuration for a sources specified in data_source_restrictions.
     */
    public readonly sourceConfig!: pulumi.Output<outputs.cloudsearch.v1.SourceConfigResponse[]>;

    /**
     * Create a SettingSearchapplication resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SettingSearchapplicationArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.searchapplicationsId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'searchapplicationsId'");
            }
            inputs["dataSourceRestrictions"] = args ? args.dataSourceRestrictions : undefined;
            inputs["defaultFacetOptions"] = args ? args.defaultFacetOptions : undefined;
            inputs["defaultSortOptions"] = args ? args.defaultSortOptions : undefined;
            inputs["displayName"] = args ? args.displayName : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["scoringConfig"] = args ? args.scoringConfig : undefined;
            inputs["searchapplicationsId"] = args ? args.searchapplicationsId : undefined;
            inputs["sourceConfig"] = args ? args.sourceConfig : undefined;
            inputs["operationIds"] = undefined /*out*/;
        } else {
            inputs["dataSourceRestrictions"] = undefined /*out*/;
            inputs["defaultFacetOptions"] = undefined /*out*/;
            inputs["defaultSortOptions"] = undefined /*out*/;
            inputs["displayName"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["operationIds"] = undefined /*out*/;
            inputs["scoringConfig"] = undefined /*out*/;
            inputs["sourceConfig"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(SettingSearchapplication.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a SettingSearchapplication resource.
 */
export interface SettingSearchapplicationArgs {
    /**
     * Retrictions applied to the configurations. The maximum number of elements is 10.
     */
    readonly dataSourceRestrictions?: pulumi.Input<pulumi.Input<inputs.cloudsearch.v1.DataSourceRestriction>[]>;
    /**
     * The default fields for returning facet results. The sources specified here also have been included in data_source_restrictions above.
     */
    readonly defaultFacetOptions?: pulumi.Input<pulumi.Input<inputs.cloudsearch.v1.FacetOptions>[]>;
    /**
     * The default options for sorting the search results
     */
    readonly defaultSortOptions?: pulumi.Input<inputs.cloudsearch.v1.SortOptions>;
    /**
     * Display name of the Search Application. The maximum length is 300 characters.
     */
    readonly displayName?: pulumi.Input<string>;
    /**
     * Name of the Search Application. Format: searchapplications/{application_id}.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Configuration for ranking results.
     */
    readonly scoringConfig?: pulumi.Input<inputs.cloudsearch.v1.ScoringConfig>;
    readonly searchapplicationsId: pulumi.Input<string>;
    /**
     * Configuration for a sources specified in data_source_restrictions.
     */
    readonly sourceConfig?: pulumi.Input<pulumi.Input<inputs.cloudsearch.v1.SourceConfig>[]>;
}