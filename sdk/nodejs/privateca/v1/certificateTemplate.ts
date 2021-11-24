// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Create a new CertificateTemplate in a given Project and Location.
 * Auto-naming is currently not supported for this resource.
 */
export class CertificateTemplate extends pulumi.CustomResource {
    /**
     * Get an existing CertificateTemplate resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): CertificateTemplate {
        return new CertificateTemplate(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:privateca/v1:CertificateTemplate';

    /**
     * Returns true if the given object is an instance of CertificateTemplate.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is CertificateTemplate {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === CertificateTemplate.__pulumiType;
    }

    /**
     * The time at which this CertificateTemplate was created.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * Optional. A human-readable description of scenarios this template is intended for.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * Optional. Describes constraints on identities that may be appear in Certificates issued using this template. If this is omitted, then this template will not add restrictions on a certificate's identity.
     */
    public readonly identityConstraints!: pulumi.Output<outputs.privateca.v1.CertificateIdentityConstraintsResponse>;
    /**
     * Optional. Labels with user-defined metadata.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string}>;
    /**
     * The resource name for this CertificateTemplate in the format `projects/*&#47;locations/*&#47;certificateTemplates/*`.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Optional. Describes the set of X.509 extensions that may appear in a Certificate issued using this CertificateTemplate. If a certificate request sets extensions that don't appear in the passthrough_extensions, those extensions will be dropped. If the issuing CaPool's IssuancePolicy defines baseline_values that don't appear here, the certificate issuance request will fail. If this is omitted, then this template will not add restrictions on a certificate's X.509 extensions. These constraints do not apply to X.509 extensions set in this CertificateTemplate's predefined_values.
     */
    public readonly passthroughExtensions!: pulumi.Output<outputs.privateca.v1.CertificateExtensionConstraintsResponse>;
    /**
     * Optional. A set of X.509 values that will be applied to all issued certificates that use this template. If the certificate request includes conflicting values for the same properties, they will be overwritten by the values defined here. If the issuing CaPool's IssuancePolicy defines conflicting baseline_values for the same properties, the certificate issuance request will fail.
     */
    public readonly predefinedValues!: pulumi.Output<outputs.privateca.v1.X509ParametersResponse>;
    /**
     * The time at which this CertificateTemplate was updated.
     */
    public /*out*/ readonly updateTime!: pulumi.Output<string>;

    /**
     * Create a CertificateTemplate resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CertificateTemplateArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.certificateTemplateId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'certificateTemplateId'");
            }
            resourceInputs["certificateTemplateId"] = args ? args.certificateTemplateId : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["identityConstraints"] = args ? args.identityConstraints : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["passthroughExtensions"] = args ? args.passthroughExtensions : undefined;
            resourceInputs["predefinedValues"] = args ? args.predefinedValues : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["requestId"] = args ? args.requestId : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        } else {
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["identityConstraints"] = undefined /*out*/;
            resourceInputs["labels"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["passthroughExtensions"] = undefined /*out*/;
            resourceInputs["predefinedValues"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(CertificateTemplate.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a CertificateTemplate resource.
 */
export interface CertificateTemplateArgs {
    certificateTemplateId: pulumi.Input<string>;
    /**
     * Optional. A human-readable description of scenarios this template is intended for.
     */
    description?: pulumi.Input<string>;
    /**
     * Optional. Describes constraints on identities that may be appear in Certificates issued using this template. If this is omitted, then this template will not add restrictions on a certificate's identity.
     */
    identityConstraints?: pulumi.Input<inputs.privateca.v1.CertificateIdentityConstraintsArgs>;
    /**
     * Optional. Labels with user-defined metadata.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    location?: pulumi.Input<string>;
    /**
     * Optional. Describes the set of X.509 extensions that may appear in a Certificate issued using this CertificateTemplate. If a certificate request sets extensions that don't appear in the passthrough_extensions, those extensions will be dropped. If the issuing CaPool's IssuancePolicy defines baseline_values that don't appear here, the certificate issuance request will fail. If this is omitted, then this template will not add restrictions on a certificate's X.509 extensions. These constraints do not apply to X.509 extensions set in this CertificateTemplate's predefined_values.
     */
    passthroughExtensions?: pulumi.Input<inputs.privateca.v1.CertificateExtensionConstraintsArgs>;
    /**
     * Optional. A set of X.509 values that will be applied to all issued certificates that use this template. If the certificate request includes conflicting values for the same properties, they will be overwritten by the values defined here. If the issuing CaPool's IssuancePolicy defines conflicting baseline_values for the same properties, the certificate issuance request will fail.
     */
    predefinedValues?: pulumi.Input<inputs.privateca.v1.X509ParametersArgs>;
    project?: pulumi.Input<string>;
    requestId?: pulumi.Input<string>;
}
