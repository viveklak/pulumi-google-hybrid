// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Create a new CryptoKey within a KeyRing. CryptoKey.purpose and CryptoKey.version_template.algorithm are required.
 * Auto-naming is currently not supported for this resource.
 * Note - this resource's API doesn't support deletion. When deleted, the resource will persist
 * on Google Cloud even though it will be deleted from Pulumi state.
 */
export class CryptoKey extends pulumi.CustomResource {
    /**
     * Get an existing CryptoKey resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): CryptoKey {
        return new CryptoKey(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:cloudkms/v1:CryptoKey';

    /**
     * Returns true if the given object is an instance of CryptoKey.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is CryptoKey {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === CryptoKey.__pulumiType;
    }

    /**
     * The time at which this CryptoKey was created.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * Immutable. The period of time that versions of this key spend in the DESTROY_SCHEDULED state before transitioning to DESTROYED. If not specified at creation time, the default duration is 24 hours.
     */
    public readonly destroyScheduledDuration!: pulumi.Output<string>;
    /**
     * Immutable. Whether this key may contain imported versions only.
     */
    public readonly importOnly!: pulumi.Output<boolean>;
    /**
     * Labels with user-defined metadata. For more information, see [Labeling Keys](https://cloud.google.com/kms/docs/labeling-keys).
     */
    public readonly labels!: pulumi.Output<{[key: string]: string}>;
    /**
     * The resource name for this CryptoKey in the format `projects/*&#47;locations/*&#47;keyRings/*&#47;cryptoKeys/*`.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * At next_rotation_time, the Key Management Service will automatically: 1. Create a new version of this CryptoKey. 2. Mark the new version as primary. Key rotations performed manually via CreateCryptoKeyVersion and UpdateCryptoKeyPrimaryVersion do not affect next_rotation_time. Keys with purpose ENCRYPT_DECRYPT support automatic rotation. For other keys, this field must be omitted.
     */
    public readonly nextRotationTime!: pulumi.Output<string>;
    /**
     * A copy of the "primary" CryptoKeyVersion that will be used by Encrypt when this CryptoKey is given in EncryptRequest.name. The CryptoKey's primary version can be updated via UpdateCryptoKeyPrimaryVersion. Keys with purpose ENCRYPT_DECRYPT may have a primary. For other keys, this field will be omitted.
     */
    public /*out*/ readonly primary!: pulumi.Output<outputs.cloudkms.v1.CryptoKeyVersionResponse>;
    /**
     * Immutable. The immutable purpose of this CryptoKey.
     */
    public readonly purpose!: pulumi.Output<string>;
    /**
     * next_rotation_time will be advanced by this period when the service automatically rotates a key. Must be at least 24 hours and at most 876,000 hours. If rotation_period is set, next_rotation_time must also be set. Keys with purpose ENCRYPT_DECRYPT support automatic rotation. For other keys, this field must be omitted.
     */
    public readonly rotationPeriod!: pulumi.Output<string>;
    /**
     * A template describing settings for new CryptoKeyVersion instances. The properties of new CryptoKeyVersion instances created by either CreateCryptoKeyVersion or auto-rotation are controlled by this template.
     */
    public readonly versionTemplate!: pulumi.Output<outputs.cloudkms.v1.CryptoKeyVersionTemplateResponse>;

    /**
     * Create a CryptoKey resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CryptoKeyArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.cryptoKeyId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'cryptoKeyId'");
            }
            if ((!args || args.keyRingId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'keyRingId'");
            }
            resourceInputs["cryptoKeyId"] = args ? args.cryptoKeyId : undefined;
            resourceInputs["destroyScheduledDuration"] = args ? args.destroyScheduledDuration : undefined;
            resourceInputs["importOnly"] = args ? args.importOnly : undefined;
            resourceInputs["keyRingId"] = args ? args.keyRingId : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["nextRotationTime"] = args ? args.nextRotationTime : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["purpose"] = args ? args.purpose : undefined;
            resourceInputs["rotationPeriod"] = args ? args.rotationPeriod : undefined;
            resourceInputs["skipInitialVersionCreation"] = args ? args.skipInitialVersionCreation : undefined;
            resourceInputs["versionTemplate"] = args ? args.versionTemplate : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["primary"] = undefined /*out*/;
        } else {
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["destroyScheduledDuration"] = undefined /*out*/;
            resourceInputs["importOnly"] = undefined /*out*/;
            resourceInputs["labels"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["nextRotationTime"] = undefined /*out*/;
            resourceInputs["primary"] = undefined /*out*/;
            resourceInputs["purpose"] = undefined /*out*/;
            resourceInputs["rotationPeriod"] = undefined /*out*/;
            resourceInputs["versionTemplate"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(CryptoKey.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a CryptoKey resource.
 */
export interface CryptoKeyArgs {
    cryptoKeyId: pulumi.Input<string>;
    /**
     * Immutable. The period of time that versions of this key spend in the DESTROY_SCHEDULED state before transitioning to DESTROYED. If not specified at creation time, the default duration is 24 hours.
     */
    destroyScheduledDuration?: pulumi.Input<string>;
    /**
     * Immutable. Whether this key may contain imported versions only.
     */
    importOnly?: pulumi.Input<boolean>;
    keyRingId: pulumi.Input<string>;
    /**
     * Labels with user-defined metadata. For more information, see [Labeling Keys](https://cloud.google.com/kms/docs/labeling-keys).
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    location?: pulumi.Input<string>;
    /**
     * At next_rotation_time, the Key Management Service will automatically: 1. Create a new version of this CryptoKey. 2. Mark the new version as primary. Key rotations performed manually via CreateCryptoKeyVersion and UpdateCryptoKeyPrimaryVersion do not affect next_rotation_time. Keys with purpose ENCRYPT_DECRYPT support automatic rotation. For other keys, this field must be omitted.
     */
    nextRotationTime?: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    /**
     * Immutable. The immutable purpose of this CryptoKey.
     */
    purpose?: pulumi.Input<enums.cloudkms.v1.CryptoKeyPurpose>;
    /**
     * next_rotation_time will be advanced by this period when the service automatically rotates a key. Must be at least 24 hours and at most 876,000 hours. If rotation_period is set, next_rotation_time must also be set. Keys with purpose ENCRYPT_DECRYPT support automatic rotation. For other keys, this field must be omitted.
     */
    rotationPeriod?: pulumi.Input<string>;
    skipInitialVersionCreation?: pulumi.Input<string>;
    /**
     * A template describing settings for new CryptoKeyVersion instances. The properties of new CryptoKeyVersion instances created by either CreateCryptoKeyVersion or auto-rotation are controlled by this template.
     */
    versionTemplate?: pulumi.Input<inputs.cloudkms.v1.CryptoKeyVersionTemplateArgs>;
}
