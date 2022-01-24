// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Create a CaPool.
 * Auto-naming is currently not supported for this resource.
 */
export class CaPool extends pulumi.CustomResource {
    /**
     * Get an existing CaPool resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): CaPool {
        return new CaPool(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:privateca/v1:CaPool';

    /**
     * Returns true if the given object is an instance of CaPool.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is CaPool {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === CaPool.__pulumiType;
    }

    /**
     * Optional. The IssuancePolicy to control how Certificates will be issued from this CaPool.
     */
    public readonly issuancePolicy!: pulumi.Output<outputs.privateca.v1.IssuancePolicyResponse>;
    /**
     * Optional. Labels with user-defined metadata.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string}>;
    /**
     * The resource name for this CaPool in the format `projects/*&#47;locations/*&#47;caPools/*`.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Optional. The PublishingOptions to follow when issuing Certificates from any CertificateAuthority in this CaPool.
     */
    public readonly publishingOptions!: pulumi.Output<outputs.privateca.v1.PublishingOptionsResponse>;
    /**
     * Immutable. The Tier of this CaPool.
     */
    public readonly tier!: pulumi.Output<string>;

    /**
     * Create a CaPool resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CaPoolArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.caPoolId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'caPoolId'");
            }
            if ((!args || args.tier === undefined) && !opts.urn) {
                throw new Error("Missing required property 'tier'");
            }
            resourceInputs["caPoolId"] = args ? args.caPoolId : undefined;
            resourceInputs["issuancePolicy"] = args ? args.issuancePolicy : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["publishingOptions"] = args ? args.publishingOptions : undefined;
            resourceInputs["requestId"] = args ? args.requestId : undefined;
            resourceInputs["tier"] = args ? args.tier : undefined;
            resourceInputs["name"] = undefined /*out*/;
        } else {
            resourceInputs["issuancePolicy"] = undefined /*out*/;
            resourceInputs["labels"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["publishingOptions"] = undefined /*out*/;
            resourceInputs["tier"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(CaPool.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a CaPool resource.
 */
export interface CaPoolArgs {
    caPoolId: pulumi.Input<string>;
    /**
     * Optional. The IssuancePolicy to control how Certificates will be issued from this CaPool.
     */
    issuancePolicy?: pulumi.Input<inputs.privateca.v1.IssuancePolicyArgs>;
    /**
     * Optional. Labels with user-defined metadata.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    location?: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    /**
     * Optional. The PublishingOptions to follow when issuing Certificates from any CertificateAuthority in this CaPool.
     */
    publishingOptions?: pulumi.Input<inputs.privateca.v1.PublishingOptionsArgs>;
    requestId?: pulumi.Input<string>;
    /**
     * Immutable. The Tier of this CaPool.
     */
    tier: pulumi.Input<enums.privateca.v1.CaPoolTier>;
}
