// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates an alias from a key/certificate pair. The structure of the request is controlled by the `format` query parameter: - `keycertfile` - Separate PEM-encoded key and certificate files are uploaded. Set `Content-Type: multipart/form-data` and include the `keyFile`, `certFile`, and `password` (if keys are encrypted) fields in the request body. If uploading to a truststore, omit `keyFile`. - `pkcs12` - A PKCS12 file is uploaded. Set `Content-Type: multipart/form-data`, provide the file in the `file` field, and include the `password` field if the file is encrypted in the request body. - `selfsignedcert` - A new private key and certificate are generated. Set `Content-Type: application/json` and include CertificateGenerationSpec in the request body.
 * Auto-naming is currently not supported for this resource.
 */
export class Alias extends pulumi.CustomResource {
    /**
     * Get an existing Alias resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Alias {
        return new Alias(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:apigee/v1:Alias';

    /**
     * Returns true if the given object is an instance of Alias.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Alias {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Alias.__pulumiType;
    }

    /**
     * Resource ID for this alias. Values must match the regular expression `[^/]{1,255}`.
     */
    public readonly alias!: pulumi.Output<string>;
    /**
     * Chain of certificates under this alias.
     */
    public /*out*/ readonly certsInfo!: pulumi.Output<outputs.apigee.v1.GoogleCloudApigeeV1CertificateResponse>;
    /**
     * Type of alias.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a Alias resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AliasArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.environmentId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'environmentId'");
            }
            if ((!args || args.format === undefined) && !opts.urn) {
                throw new Error("Missing required property 'format'");
            }
            if ((!args || args.keystoreId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'keystoreId'");
            }
            if ((!args || args.organizationId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'organizationId'");
            }
            resourceInputs["alias"] = args ? args.alias : undefined;
            resourceInputs["contentType"] = args ? args.contentType : undefined;
            resourceInputs["data"] = args ? args.data : undefined;
            resourceInputs["environmentId"] = args ? args.environmentId : undefined;
            resourceInputs["extensions"] = args ? args.extensions : undefined;
            resourceInputs["format"] = args ? args.format : undefined;
            resourceInputs["ignoreExpiryValidation"] = args ? args.ignoreExpiryValidation : undefined;
            resourceInputs["ignoreNewlineValidation"] = args ? args.ignoreNewlineValidation : undefined;
            resourceInputs["keystoreId"] = args ? args.keystoreId : undefined;
            resourceInputs["organizationId"] = args ? args.organizationId : undefined;
            resourceInputs["certsInfo"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
        } else {
            resourceInputs["alias"] = undefined /*out*/;
            resourceInputs["certsInfo"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Alias.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Alias resource.
 */
export interface AliasArgs {
    alias?: pulumi.Input<string>;
    /**
     * The HTTP Content-Type header value specifying the content type of the body.
     */
    contentType?: pulumi.Input<string>;
    /**
     * The HTTP request/response body as raw binary.
     */
    data?: pulumi.Input<string>;
    environmentId: pulumi.Input<string>;
    /**
     * Application specific response metadata. Must be set in the first response for streaming APIs.
     */
    extensions?: pulumi.Input<pulumi.Input<{[key: string]: pulumi.Input<string>}>[]>;
    format: pulumi.Input<string>;
    ignoreExpiryValidation?: pulumi.Input<string>;
    ignoreNewlineValidation?: pulumi.Input<string>;
    keystoreId: pulumi.Input<string>;
    organizationId: pulumi.Input<string>;
}
