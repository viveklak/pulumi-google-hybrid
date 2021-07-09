// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates a new Consent in the parent consent store.
 */
export class Consent extends pulumi.CustomResource {
    /**
     * Get an existing Consent resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Consent {
        return new Consent(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:healthcare/v1:Consent';

    /**
     * Returns true if the given object is an instance of Consent.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Consent {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Consent.__pulumiType;
    }

    /**
     * The resource name of the Consent artifact that contains proof of the end user's consent, of the form `projects/{project_id}/locations/{location_id}/datasets/{dataset_id}/consentStores/{consent_store_id}/consentArtifacts/{consent_artifact_id}`.
     */
    public readonly consentArtifact!: pulumi.Output<string>;
    /**
     * Timestamp in UTC of when this Consent is considered expired.
     */
    public readonly expireTime!: pulumi.Output<string>;
    /**
     * Optional. User-supplied key-value pairs used to organize Consent resources. Metadata keys must: - be between 1 and 63 characters long - have a UTF-8 encoding of maximum 128 bytes - begin with a letter - consist of up to 63 characters including lowercase letters, numeric characters, underscores, and dashes Metadata values must be: - be between 1 and 63 characters long - have a UTF-8 encoding of maximum 128 bytes - consist of up to 63 characters including lowercase letters, numeric characters, underscores, and dashes No more than 64 metadata entries can be associated with a given consent.
     */
    public readonly metadata!: pulumi.Output<{[key: string]: string}>;
    /**
     * Resource name of the Consent, of the form `projects/{project_id}/locations/{location_id}/datasets/{dataset_id}/consentStores/{consent_store_id}/consents/{consent_id}`. Cannot be changed after creation.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Optional. Represents a user's consent in terms of the resources that can be accessed and under what conditions.
     */
    public readonly policies!: pulumi.Output<outputs.healthcare.v1.GoogleCloudHealthcareV1ConsentPolicyResponse[]>;
    /**
     * The timestamp that the revision was created.
     */
    public /*out*/ readonly revisionCreateTime!: pulumi.Output<string>;
    /**
     * The revision ID of the Consent. The format is an 8-character hexadecimal string. Refer to a specific revision of a Consent by appending `@{revision_id}` to the Consent's resource name.
     */
    public /*out*/ readonly revisionId!: pulumi.Output<string>;
    /**
     * Indicates the current state of this Consent.
     */
    public readonly state!: pulumi.Output<string>;
    /**
     * Input only. The time to live for this Consent from when it is created.
     */
    public readonly ttl!: pulumi.Output<string>;
    /**
     * User's UUID provided by the client.
     */
    public readonly userId!: pulumi.Output<string>;

    /**
     * Create a Consent resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ConsentArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.consentArtifact === undefined) && !opts.urn) {
                throw new Error("Missing required property 'consentArtifact'");
            }
            if ((!args || args.consentStoreId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'consentStoreId'");
            }
            if ((!args || args.datasetId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'datasetId'");
            }
            if ((!args || args.location === undefined) && !opts.urn) {
                throw new Error("Missing required property 'location'");
            }
            if ((!args || args.project === undefined) && !opts.urn) {
                throw new Error("Missing required property 'project'");
            }
            if ((!args || args.state === undefined) && !opts.urn) {
                throw new Error("Missing required property 'state'");
            }
            if ((!args || args.userId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'userId'");
            }
            inputs["consentArtifact"] = args ? args.consentArtifact : undefined;
            inputs["consentStoreId"] = args ? args.consentStoreId : undefined;
            inputs["datasetId"] = args ? args.datasetId : undefined;
            inputs["expireTime"] = args ? args.expireTime : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["metadata"] = args ? args.metadata : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["policies"] = args ? args.policies : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["state"] = args ? args.state : undefined;
            inputs["ttl"] = args ? args.ttl : undefined;
            inputs["userId"] = args ? args.userId : undefined;
            inputs["revisionCreateTime"] = undefined /*out*/;
            inputs["revisionId"] = undefined /*out*/;
        } else {
            inputs["consentArtifact"] = undefined /*out*/;
            inputs["expireTime"] = undefined /*out*/;
            inputs["metadata"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["policies"] = undefined /*out*/;
            inputs["revisionCreateTime"] = undefined /*out*/;
            inputs["revisionId"] = undefined /*out*/;
            inputs["state"] = undefined /*out*/;
            inputs["ttl"] = undefined /*out*/;
            inputs["userId"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Consent.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Consent resource.
 */
export interface ConsentArgs {
    /**
     * The resource name of the Consent artifact that contains proof of the end user's consent, of the form `projects/{project_id}/locations/{location_id}/datasets/{dataset_id}/consentStores/{consent_store_id}/consentArtifacts/{consent_artifact_id}`.
     */
    consentArtifact: pulumi.Input<string>;
    consentStoreId: pulumi.Input<string>;
    datasetId: pulumi.Input<string>;
    /**
     * Timestamp in UTC of when this Consent is considered expired.
     */
    expireTime?: pulumi.Input<string>;
    location: pulumi.Input<string>;
    /**
     * Optional. User-supplied key-value pairs used to organize Consent resources. Metadata keys must: - be between 1 and 63 characters long - have a UTF-8 encoding of maximum 128 bytes - begin with a letter - consist of up to 63 characters including lowercase letters, numeric characters, underscores, and dashes Metadata values must be: - be between 1 and 63 characters long - have a UTF-8 encoding of maximum 128 bytes - consist of up to 63 characters including lowercase letters, numeric characters, underscores, and dashes No more than 64 metadata entries can be associated with a given consent.
     */
    metadata?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Resource name of the Consent, of the form `projects/{project_id}/locations/{location_id}/datasets/{dataset_id}/consentStores/{consent_store_id}/consents/{consent_id}`. Cannot be changed after creation.
     */
    name?: pulumi.Input<string>;
    /**
     * Optional. Represents a user's consent in terms of the resources that can be accessed and under what conditions.
     */
    policies?: pulumi.Input<pulumi.Input<inputs.healthcare.v1.GoogleCloudHealthcareV1ConsentPolicyArgs>[]>;
    project: pulumi.Input<string>;
    /**
     * Indicates the current state of this Consent.
     */
    state: pulumi.Input<enums.healthcare.v1.ConsentState>;
    /**
     * Input only. The time to live for this Consent from when it is created.
     */
    ttl?: pulumi.Input<string>;
    /**
     * User's UUID provided by the client.
     */
    userId: pulumi.Input<string>;
}
