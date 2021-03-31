// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Creates a new consent store in the parent dataset. Attempting to create a consent store with the same ID as an existing store fails with an ALREADY_EXISTS error.
 */
export class DatasetConsentStore extends pulumi.CustomResource {
    /**
     * Get an existing DatasetConsentStore resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): DatasetConsentStore {
        return new DatasetConsentStore(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-cloud:healthcare/v1:DatasetConsentStore';

    /**
     * Returns true if the given object is an instance of DatasetConsentStore.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DatasetConsentStore {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DatasetConsentStore.__pulumiType;
    }


    /**
     * Create a DatasetConsentStore resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DatasetConsentStoreArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.consentStoresId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'consentStoresId'");
            }
            if ((!args || args.datasetsId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'datasetsId'");
            }
            if ((!args || args.locationsId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'locationsId'");
            }
            if ((!args || args.projectsId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'projectsId'");
            }
            inputs["consentStoresId"] = args ? args.consentStoresId : undefined;
            inputs["datasetsId"] = args ? args.datasetsId : undefined;
            inputs["defaultConsentTtl"] = args ? args.defaultConsentTtl : undefined;
            inputs["enableConsentCreateOnUpdate"] = args ? args.enableConsentCreateOnUpdate : undefined;
            inputs["labels"] = args ? args.labels : undefined;
            inputs["locationsId"] = args ? args.locationsId : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["projectsId"] = args ? args.projectsId : undefined;
        } else {
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(DatasetConsentStore.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a DatasetConsentStore resource.
 */
export interface DatasetConsentStoreArgs {
    readonly consentStoresId: pulumi.Input<string>;
    readonly datasetsId: pulumi.Input<string>;
    /**
     * Optional. Default time to live for Consents created in this store. Must be at least 24 hours. Updating this field will not affect the expiration time of existing consents.
     */
    readonly defaultConsentTtl?: pulumi.Input<string>;
    /**
     * Optional. If `true`, UpdateConsent creates the Consent if it does not already exist. If unspecified, defaults to `false`.
     */
    readonly enableConsentCreateOnUpdate?: pulumi.Input<boolean>;
    /**
     * Optional. User-supplied key-value pairs used to organize consent stores. Label keys must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must conform to the following PCRE regular expression: \p{Ll}\p{Lo}{0,62}. Label values must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must conform to the following PCRE regular expression: [\p{Ll}\p{Lo}\p{N}_-]{0,63}. No more than 64 labels can be associated with a given store. For more information: https://cloud.google.com/healthcare/docs/how-tos/labeling-resources
     */
    readonly labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    readonly locationsId: pulumi.Input<string>;
    /**
     * Resource name of the consent store, of the form `projects/{project_id}/locations/{location_id}/datasets/{dataset_id}/consentStores/{consent_store_id}`. Cannot be changed after creation.
     */
    readonly name?: pulumi.Input<string>;
    readonly projectsId: pulumi.Input<string>;
}
