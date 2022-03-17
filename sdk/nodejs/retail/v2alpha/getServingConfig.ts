// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Gets a ServingConfig. Returns a NotFound error if the ServingConfig does not exist.
 */
export function getServingConfig(args: GetServingConfigArgs, opts?: pulumi.InvokeOptions): Promise<GetServingConfigResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("google-native:retail/v2alpha:getServingConfig", {
        "catalogId": args.catalogId,
        "location": args.location,
        "project": args.project,
        "servingConfigId": args.servingConfigId,
    }, opts);
}

export interface GetServingConfigArgs {
    catalogId: string;
    location: string;
    project?: string;
    servingConfigId: string;
}

export interface GetServingConfigResult {
    /**
     * Condition boost specifications. If a product matches multiple conditions in the specifications, boost scores from these specifications are all applied and combined in a non-linear way. Maximum number of specifications is 100. Notice that if both ServingConfig.boost_control_ids and [SearchRequest.boost_spec] are set, the boost conditions from both places are evaluated. If a search request matches multiple boost conditions, the final boost score is equal to the sum of the boost scores from all matched boost conditions. Can only be set if solution_types is SOLUTION_TYPE_SEARCH.
     */
    readonly boostControlIds: string[];
    /**
     * The human readable serving config display name. Used in Retail UI. This field must be a UTF-8 encoded string with a length limit of 128 characters. Otherwise, an INVALID_ARGUMENT error is returned.
     */
    readonly displayName: string;
    /**
     * How much diversity to use in recommendation model results e.g. 'medium-diversity' or 'high-diversity'. Currently supported values: * 'no-diversity' * 'low-diversity' * 'medium-diversity' * 'high-diversity' * 'auto-diversity' If not specified, we choose default based on recommendation model type. Default value: 'no-diversity'. Can only be set if solution_types is SOLUTION_TYPE_RECOMMENDATION.
     */
    readonly diversityLevel: string;
    /**
     * Condition do not associate specifications. If multiple do not associate conditions match, all matching do not associate controls in the list will execute. - Order does not matter. - Maximum number of specifications is 100. Can only be set if solution_types is SOLUTION_TYPE_SEARCH.
     */
    readonly doNotAssociateControlIds: string[];
    /**
     * The specification for dynamically generated facets. Notice that only textual facets can be dynamically generated. Can only be set if solution_types is SOLUTION_TYPE_SEARCH.
     */
    readonly dynamicFacetSpec: outputs.retail.v2alpha.GoogleCloudRetailV2alphaSearchRequestDynamicFacetSpecResponse;
    /**
     * Whether to add additional category filters on the 'similar-items' model. If not specified, we enable it by default. Allowed values are: * 'no-category-match': No additional filtering of original results from the model and the customer's filters. * 'relaxed-category-match': Only keep results with categories that match at least one item categories in the PredictRequests's context item. * If customer also sends filters in the PredictRequest, then the results will satisfy both conditions (user given and category match). Can only be set if solution_types is SOLUTION_TYPE_RECOMMENDATION.
     */
    readonly enableCategoryFilterLevel: string;
    /**
     * Facet specifications for faceted search. If empty, no facets are returned. The ids refer to the ids of Control resources with only the Facet control set. These controls are assumed to be in the same Catalog as the ServingConfig. A maximum of 100 values are allowed. Otherwise, an INVALID_ARGUMENT error is returned. Can only be set if solution_types is SOLUTION_TYPE_SEARCH.
     */
    readonly facetControlIds: string[];
    /**
     * Condition filter specifications. If a product matches multiple conditions in the specifications, filters from these specifications are all applied and combined via the AND operator. Maximum number of specifications is 100. Can only be set if solution_types is SOLUTION_TYPE_SEARCH.
     */
    readonly filterControlIds: string[];
    /**
     * Condition ignore specifications. If multiple ignore conditions match, all matching ignore controls in the list will execute. - Order does not matter. - Maximum number of specifications is 100. Can only be set if solution_types is SOLUTION_TYPE_SEARCH.
     */
    readonly ignoreControlIds: string[];
    /**
     * The id of the model to use at serving time. Currently only RecommendationModels are supported: https://cloud.google.com/retail/recommendations-ai/docs/create-models Can be changed but only to a compatible model (e.g. others-you-may-like CTR to others-you-may-like CVR). Required when solution_types is SOLUTION_TYPE_RECOMMENDATION.
     */
    readonly modelId: string;
    /**
     * Immutable. Fully qualified name projects/*&#47;locations/global/catalogs/*&#47;servingConfig/*
     */
    readonly name: string;
    /**
     * Condition oneway synonyms specifications. If multiple oneway synonyms conditions match, all matching oneway synonyms controls in the list will execute. Order of controls in the list will not matter. Maximum number of specifications is 100. Can only be set if solution_types is SOLUTION_TYPE_SEARCH.
     */
    readonly onewaySynonymsControlIds: string[];
    /**
     * How much price ranking we want in serving results. Price reranking causes product items with a similar recommendation probability to be ordered by price, with the highest-priced items first. This setting could result in a decrease in click-through and conversion rates. Allowed values are: * 'no-price-reranking' * 'low-price-raranking' * 'medium-price-reranking' * 'high-price-reranking' If not specified, we choose default based on model type. Default value: 'no-price-reranking'. Can only be set if solution_types is SOLUTION_TYPE_RECOMMENDATION.
     */
    readonly priceRerankingLevel: string;
    /**
     * Condition redirect specifications. Only the first triggered redirect action is applied, even if multiple apply. Maximum number of specifications is 1000. Can only be set if solution_types is SOLUTION_TYPE_SEARCH.
     */
    readonly redirectControlIds: string[];
    /**
     * Condition replacement specifications. - Applied according to the order in the list. - A previously replaced term can not be re-replaced. - Maximum number of specifications is 100. Can only be set if solution_types is SOLUTION_TYPE_SEARCH.
     */
    readonly replacementControlIds: string[];
    /**
     * Immutable. Specifies the solution types that a serving config can be associated with. Currently we support setting only one type of solution.
     */
    readonly solutionTypes: string[];
    /**
     * Condition synonyms specifications. If multiple syonyms conditions match, all matching synonyms control in the list will execute. Order of controls in the list will not matter. Maximum number of specifications is 100. Can only be set if solution_types is SOLUTION_TYPE_SEARCH.
     */
    readonly twowaySynonymsControlIds: string[];
}

export function getServingConfigOutput(args: GetServingConfigOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetServingConfigResult> {
    return pulumi.output(args).apply(a => getServingConfig(a, opts))
}

export interface GetServingConfigOutputArgs {
    catalogId: pulumi.Input<string>;
    location: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    servingConfigId: pulumi.Input<string>;
}
