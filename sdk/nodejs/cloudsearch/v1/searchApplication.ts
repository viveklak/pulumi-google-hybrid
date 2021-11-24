// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates a search application. **Note:** This API requires an admin account to execute.
 */
export class SearchApplication extends pulumi.CustomResource {
    /**
     * Get an existing SearchApplication resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): SearchApplication {
        return new SearchApplication(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:cloudsearch/v1:SearchApplication';

    /**
     * Returns true if the given object is an instance of SearchApplication.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SearchApplication {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SearchApplication.__pulumiType;
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
     * Indicates whether audit logging is on/off for requests made for the search application in query APIs.
     */
    public readonly enableAuditLog!: pulumi.Output<boolean>;
    /**
     * Name of the Search Application. Format: searchapplications/{application_id}.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * IDs of the Long Running Operations (LROs) currently running for this schema. Output only field.
     */
    public /*out*/ readonly operationIds!: pulumi.Output<string[]>;
    /**
     * The default options for query interpretation
     */
    public readonly queryInterpretationConfig!: pulumi.Output<outputs.cloudsearch.v1.QueryInterpretationConfigResponse>;
    /**
     * Configuration for ranking results.
     */
    public readonly scoringConfig!: pulumi.Output<outputs.cloudsearch.v1.ScoringConfigResponse>;
    /**
     * Configuration for a sources specified in data_source_restrictions.
     */
    public readonly sourceConfig!: pulumi.Output<outputs.cloudsearch.v1.SourceConfigResponse[]>;

    /**
     * Create a SearchApplication resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: SearchApplicationArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            resourceInputs["dataSourceRestrictions"] = args ? args.dataSourceRestrictions : undefined;
            resourceInputs["defaultFacetOptions"] = args ? args.defaultFacetOptions : undefined;
            resourceInputs["defaultSortOptions"] = args ? args.defaultSortOptions : undefined;
            resourceInputs["displayName"] = args ? args.displayName : undefined;
            resourceInputs["enableAuditLog"] = args ? args.enableAuditLog : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["queryInterpretationConfig"] = args ? args.queryInterpretationConfig : undefined;
            resourceInputs["scoringConfig"] = args ? args.scoringConfig : undefined;
            resourceInputs["sourceConfig"] = args ? args.sourceConfig : undefined;
            resourceInputs["operationIds"] = undefined /*out*/;
        } else {
            resourceInputs["dataSourceRestrictions"] = undefined /*out*/;
            resourceInputs["defaultFacetOptions"] = undefined /*out*/;
            resourceInputs["defaultSortOptions"] = undefined /*out*/;
            resourceInputs["displayName"] = undefined /*out*/;
            resourceInputs["enableAuditLog"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["operationIds"] = undefined /*out*/;
            resourceInputs["queryInterpretationConfig"] = undefined /*out*/;
            resourceInputs["scoringConfig"] = undefined /*out*/;
            resourceInputs["sourceConfig"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(SearchApplication.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a SearchApplication resource.
 */
export interface SearchApplicationArgs {
    /**
     * Retrictions applied to the configurations. The maximum number of elements is 10.
     */
    dataSourceRestrictions?: pulumi.Input<pulumi.Input<inputs.cloudsearch.v1.DataSourceRestrictionArgs>[]>;
    /**
     * The default fields for returning facet results. The sources specified here also have been included in data_source_restrictions above.
     */
    defaultFacetOptions?: pulumi.Input<pulumi.Input<inputs.cloudsearch.v1.FacetOptionsArgs>[]>;
    /**
     * The default options for sorting the search results
     */
    defaultSortOptions?: pulumi.Input<inputs.cloudsearch.v1.SortOptionsArgs>;
    /**
     * Display name of the Search Application. The maximum length is 300 characters.
     */
    displayName?: pulumi.Input<string>;
    /**
     * Indicates whether audit logging is on/off for requests made for the search application in query APIs.
     */
    enableAuditLog?: pulumi.Input<boolean>;
    /**
     * Name of the Search Application. Format: searchapplications/{application_id}.
     */
    name?: pulumi.Input<string>;
    /**
     * The default options for query interpretation
     */
    queryInterpretationConfig?: pulumi.Input<inputs.cloudsearch.v1.QueryInterpretationConfigArgs>;
    /**
     * Configuration for ranking results.
     */
    scoringConfig?: pulumi.Input<inputs.cloudsearch.v1.ScoringConfigArgs>;
    /**
     * Configuration for a sources specified in data_source_restrictions.
     */
    sourceConfig?: pulumi.Input<pulumi.Input<inputs.cloudsearch.v1.SourceConfigArgs>[]>;
}
