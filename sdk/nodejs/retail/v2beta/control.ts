// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates a Control. If the Control to create already exists, an ALREADY_EXISTS error is returned.
 * Auto-naming is currently not supported for this resource.
 */
export class Control extends pulumi.CustomResource {
    /**
     * Get an existing Control resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Control {
        return new Control(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:retail/v2beta:Control';

    /**
     * Returns true if the given object is an instance of Control.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Control {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Control.__pulumiType;
    }

    /**
     * List of serving configuration ids that that are associated with this control. Note the association is managed via the ServingConfig, this is an output only denormalizeed view. Assumed to be in the same catalog.
     */
    public /*out*/ readonly associatedServingConfigIds!: pulumi.Output<string[]>;
    /**
     * The human readable control display name. Used in Retail UI. This field must be a UTF-8 encoded string with a length limit of 128 characters. Otherwise, an INVALID_ARGUMENT error is thrown.
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * A facet specification to perform faceted search.
     */
    public readonly facetSpec!: pulumi.Output<outputs.retail.v2beta.GoogleCloudRetailV2betaSearchRequestFacetSpecResponse>;
    /**
     * Immutable. Fully qualified name projects/*&#47;locations/global/catalogs/*&#47;controls/*
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * A rule control - a condition-action pair. Enacts a set action when the condition is triggered. For example: Boost "gShoe" when query full matches "Running Shoes".
     */
    public readonly rule!: pulumi.Output<outputs.retail.v2beta.GoogleCloudRetailV2betaRuleResponse>;
    /**
     * Specifies the use case for the control. Affects what condition fields can be set. Only settable by search controls. Will default to SEARCH_SOLUTION_USE_CASE_SEARCH if not specified. Currently only allow one search_solution_use_case per control.
     */
    public readonly searchSolutionUseCase!: pulumi.Output<string[]>;
    /**
     * Immutable. The solution types that the serving config is used for. Currently we support setting only one type of solution at creation time. Only `SOLUTION_TYPE_SEARCH` value is supported at the moment. If no solution type is provided at creation time, will default to SOLUTION_TYPE_SEARCH.
     */
    public readonly solutionTypes!: pulumi.Output<string[]>;

    /**
     * Create a Control resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ControlArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.catalogId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'catalogId'");
            }
            if ((!args || args.controlId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'controlId'");
            }
            if ((!args || args.displayName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'displayName'");
            }
            if ((!args || args.searchSolutionUseCase === undefined) && !opts.urn) {
                throw new Error("Missing required property 'searchSolutionUseCase'");
            }
            if ((!args || args.solutionTypes === undefined) && !opts.urn) {
                throw new Error("Missing required property 'solutionTypes'");
            }
            resourceInputs["catalogId"] = args ? args.catalogId : undefined;
            resourceInputs["controlId"] = args ? args.controlId : undefined;
            resourceInputs["displayName"] = args ? args.displayName : undefined;
            resourceInputs["facetSpec"] = args ? args.facetSpec : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["rule"] = args ? args.rule : undefined;
            resourceInputs["searchSolutionUseCase"] = args ? args.searchSolutionUseCase : undefined;
            resourceInputs["solutionTypes"] = args ? args.solutionTypes : undefined;
            resourceInputs["associatedServingConfigIds"] = undefined /*out*/;
        } else {
            resourceInputs["associatedServingConfigIds"] = undefined /*out*/;
            resourceInputs["displayName"] = undefined /*out*/;
            resourceInputs["facetSpec"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["rule"] = undefined /*out*/;
            resourceInputs["searchSolutionUseCase"] = undefined /*out*/;
            resourceInputs["solutionTypes"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Control.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Control resource.
 */
export interface ControlArgs {
    catalogId: pulumi.Input<string>;
    /**
     * Required. The ID to use for the Control, which will become the final component of the Control's resource name. This value should be 4-63 characters, and valid characters are /a-z-_/.
     */
    controlId: pulumi.Input<string>;
    /**
     * The human readable control display name. Used in Retail UI. This field must be a UTF-8 encoded string with a length limit of 128 characters. Otherwise, an INVALID_ARGUMENT error is thrown.
     */
    displayName: pulumi.Input<string>;
    /**
     * A facet specification to perform faceted search.
     */
    facetSpec?: pulumi.Input<inputs.retail.v2beta.GoogleCloudRetailV2betaSearchRequestFacetSpecArgs>;
    location?: pulumi.Input<string>;
    /**
     * Immutable. Fully qualified name projects/*&#47;locations/global/catalogs/*&#47;controls/*
     */
    name?: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    /**
     * A rule control - a condition-action pair. Enacts a set action when the condition is triggered. For example: Boost "gShoe" when query full matches "Running Shoes".
     */
    rule?: pulumi.Input<inputs.retail.v2beta.GoogleCloudRetailV2betaRuleArgs>;
    /**
     * Specifies the use case for the control. Affects what condition fields can be set. Only settable by search controls. Will default to SEARCH_SOLUTION_USE_CASE_SEARCH if not specified. Currently only allow one search_solution_use_case per control.
     */
    searchSolutionUseCase: pulumi.Input<pulumi.Input<enums.retail.v2beta.ControlSearchSolutionUseCaseItem>[]>;
    /**
     * Immutable. The solution types that the serving config is used for. Currently we support setting only one type of solution at creation time. Only `SOLUTION_TYPE_SEARCH` value is supported at the moment. If no solution type is provided at creation time, will default to SOLUTION_TYPE_SEARCH.
     */
    solutionTypes: pulumi.Input<pulumi.Input<enums.retail.v2beta.ControlSolutionTypesItem>[]>;
}
