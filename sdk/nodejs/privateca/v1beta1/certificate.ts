// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Create a new Certificate in a given Project, Location from a particular CertificateAuthority.
 * Auto-naming is currently not supported for this resource.
 * Note - this resource's API doesn't support deletion. When deleted, the resource will persist
 * on Google Cloud even though it will be deleted from Pulumi state.
 */
export class Certificate extends pulumi.CustomResource {
    /**
     * Get an existing Certificate resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Certificate {
        return new Certificate(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:privateca/v1beta1:Certificate';

    /**
     * Returns true if the given object is an instance of Certificate.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Certificate {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Certificate.__pulumiType;
    }

    /**
     * A structured description of the issued X.509 certificate.
     */
    public /*out*/ readonly certificateDescription!: pulumi.Output<outputs.privateca.v1beta1.CertificateDescriptionResponse>;
    /**
     * Immutable. A description of the certificate and key that does not require X.509 or ASN.1.
     */
    public readonly config!: pulumi.Output<outputs.privateca.v1beta1.CertificateConfigResponse>;
    /**
     * The time at which this Certificate was created.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * Optional. Labels with user-defined metadata.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string}>;
    /**
     * Immutable. The desired lifetime of a certificate. Used to create the "not_before_time" and "not_after_time" fields inside an X.509 certificate. Note that the lifetime may be truncated if it would extend past the life of any certificate authority in the issuing chain.
     */
    public readonly lifetime!: pulumi.Output<string>;
    /**
     * The resource path for this Certificate in the format `projects/*&#47;locations/*&#47;certificateAuthorities/*&#47;certificates/*`.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The pem-encoded, signed X.509 certificate.
     */
    public /*out*/ readonly pemCertificate!: pulumi.Output<string>;
    /**
     * The chain that may be used to verify the X.509 certificate. Expected to be in issuer-to-root order according to RFC 5246.
     */
    public /*out*/ readonly pemCertificateChain!: pulumi.Output<string[]>;
    /**
     * Immutable. A pem-encoded X.509 certificate signing request (CSR).
     */
    public readonly pemCsr!: pulumi.Output<string>;
    /**
     * Details regarding the revocation of this Certificate. This Certificate is considered revoked if and only if this field is present.
     */
    public /*out*/ readonly revocationDetails!: pulumi.Output<outputs.privateca.v1beta1.RevocationDetailsResponse>;
    /**
     * The time at which this Certificate was updated.
     */
    public /*out*/ readonly updateTime!: pulumi.Output<string>;

    /**
     * Create a Certificate resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CertificateArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.certificateAuthorityId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'certificateAuthorityId'");
            }
            if ((!args || args.lifetime === undefined) && !opts.urn) {
                throw new Error("Missing required property 'lifetime'");
            }
            resourceInputs["certificateAuthorityId"] = args ? args.certificateAuthorityId : undefined;
            resourceInputs["certificateId"] = args ? args.certificateId : undefined;
            resourceInputs["config"] = args ? args.config : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["lifetime"] = args ? args.lifetime : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["pemCsr"] = args ? args.pemCsr : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["requestId"] = args ? args.requestId : undefined;
            resourceInputs["certificateDescription"] = undefined /*out*/;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["pemCertificate"] = undefined /*out*/;
            resourceInputs["pemCertificateChain"] = undefined /*out*/;
            resourceInputs["revocationDetails"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        } else {
            resourceInputs["certificateDescription"] = undefined /*out*/;
            resourceInputs["config"] = undefined /*out*/;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["labels"] = undefined /*out*/;
            resourceInputs["lifetime"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["pemCertificate"] = undefined /*out*/;
            resourceInputs["pemCertificateChain"] = undefined /*out*/;
            resourceInputs["pemCsr"] = undefined /*out*/;
            resourceInputs["revocationDetails"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Certificate.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Certificate resource.
 */
export interface CertificateArgs {
    certificateAuthorityId: pulumi.Input<string>;
    certificateId?: pulumi.Input<string>;
    /**
     * Immutable. A description of the certificate and key that does not require X.509 or ASN.1.
     */
    config?: pulumi.Input<inputs.privateca.v1beta1.CertificateConfigArgs>;
    /**
     * Optional. Labels with user-defined metadata.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Immutable. The desired lifetime of a certificate. Used to create the "not_before_time" and "not_after_time" fields inside an X.509 certificate. Note that the lifetime may be truncated if it would extend past the life of any certificate authority in the issuing chain.
     */
    lifetime: pulumi.Input<string>;
    location?: pulumi.Input<string>;
    /**
     * Immutable. A pem-encoded X.509 certificate signing request (CSR).
     */
    pemCsr?: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    requestId?: pulumi.Input<string>;
}
