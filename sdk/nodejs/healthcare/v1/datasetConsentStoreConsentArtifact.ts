// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates a new Consent artifact in the parent consent store.
 */
export class DatasetConsentStoreConsentArtifact extends pulumi.CustomResource {
    /**
     * Get an existing DatasetConsentStoreConsentArtifact resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): DatasetConsentStoreConsentArtifact {
        return new DatasetConsentStoreConsentArtifact(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-cloud:healthcare/v1:DatasetConsentStoreConsentArtifact';

    /**
     * Returns true if the given object is an instance of DatasetConsentStoreConsentArtifact.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DatasetConsentStoreConsentArtifact {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DatasetConsentStoreConsentArtifact.__pulumiType;
    }


    /**
     * Create a DatasetConsentStoreConsentArtifact resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DatasetConsentStoreConsentArtifactArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.consentArtifactsId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'consentArtifactsId'");
            }
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
            inputs["consentArtifactsId"] = args ? args.consentArtifactsId : undefined;
            inputs["consentContentScreenshots"] = args ? args.consentContentScreenshots : undefined;
            inputs["consentContentVersion"] = args ? args.consentContentVersion : undefined;
            inputs["consentStoresId"] = args ? args.consentStoresId : undefined;
            inputs["datasetsId"] = args ? args.datasetsId : undefined;
            inputs["guardianSignature"] = args ? args.guardianSignature : undefined;
            inputs["locationsId"] = args ? args.locationsId : undefined;
            inputs["metadata"] = args ? args.metadata : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["projectsId"] = args ? args.projectsId : undefined;
            inputs["userId"] = args ? args.userId : undefined;
            inputs["userSignature"] = args ? args.userSignature : undefined;
            inputs["witnessSignature"] = args ? args.witnessSignature : undefined;
        } else {
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(DatasetConsentStoreConsentArtifact.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a DatasetConsentStoreConsentArtifact resource.
 */
export interface DatasetConsentStoreConsentArtifactArgs {
    readonly consentArtifactsId: pulumi.Input<string>;
    /**
     * Optional. Screenshots, PDFs, or other binary information documenting the user's consent.
     */
    readonly consentContentScreenshots?: pulumi.Input<pulumi.Input<inputs.healthcare.v1.Image>[]>;
    /**
     * Optional. An string indicating the version of the consent information shown to the user.
     */
    readonly consentContentVersion?: pulumi.Input<string>;
    readonly consentStoresId: pulumi.Input<string>;
    readonly datasetsId: pulumi.Input<string>;
    /**
     * Optional. A signature from a guardian.
     */
    readonly guardianSignature?: pulumi.Input<inputs.healthcare.v1.Signature>;
    readonly locationsId: pulumi.Input<string>;
    /**
     * Optional. Metadata associated with the Consent artifact. For example, the consent locale or user agent version.
     */
    readonly metadata?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Resource name of the Consent artifact, of the form `projects/{project_id}/locations/{location_id}/datasets/{dataset_id}/consentStores/{consent_store_id}/consentArtifacts/{consent_artifact_id}`. Cannot be changed after creation.
     */
    readonly name?: pulumi.Input<string>;
    readonly projectsId: pulumi.Input<string>;
    /**
     * Required. User's UUID provided by the client.
     */
    readonly userId?: pulumi.Input<string>;
    /**
     * Optional. User's signature.
     */
    readonly userSignature?: pulumi.Input<inputs.healthcare.v1.Signature>;
    /**
     * Optional. A signature from a witness.
     */
    readonly witnessSignature?: pulumi.Input<inputs.healthcare.v1.Signature>;
}
