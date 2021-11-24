// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Uploads the specified SSL certificate.
 * Auto-naming is currently not supported for this resource.
 */
export class AuthorizedCertificate extends pulumi.CustomResource {
    /**
     * Get an existing AuthorizedCertificate resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): AuthorizedCertificate {
        return new AuthorizedCertificate(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:appengine/v1alpha:AuthorizedCertificate';

    /**
     * Returns true if the given object is an instance of AuthorizedCertificate.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AuthorizedCertificate {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AuthorizedCertificate.__pulumiType;
    }

    /**
     * The SSL certificate serving the AuthorizedCertificate resource. This must be obtained independently from a certificate authority.
     */
    public readonly certificateRawData!: pulumi.Output<outputs.appengine.v1alpha.CertificateRawDataResponse>;
    /**
     * The user-specified display name of the certificate. This is not guaranteed to be unique. Example: My Certificate.
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * Aggregate count of the domain mappings with this certificate mapped. This count includes domain mappings on applications for which the user does not have VIEWER permissions.Only returned by GET or LIST requests when specifically requested by the view=FULL_CERTIFICATE option.
     */
    public /*out*/ readonly domainMappingsCount!: pulumi.Output<number>;
    /**
     * Topmost applicable domains of this certificate. This certificate applies to these domains and their subdomains. Example: example.com.
     */
    public /*out*/ readonly domainNames!: pulumi.Output<string[]>;
    /**
     * The time when this certificate expires. To update the renewal time on this certificate, upload an SSL certificate with a different expiration time using AuthorizedCertificates.UpdateAuthorizedCertificate.
     */
    public /*out*/ readonly expireTime!: pulumi.Output<string>;
    /**
     * Only applicable if this certificate is managed by App Engine. Managed certificates are tied to the lifecycle of a DomainMapping and cannot be updated or deleted via the AuthorizedCertificates API. If this certificate is manually administered by the user, this field will be empty.
     */
    public /*out*/ readonly managedCertificate!: pulumi.Output<outputs.appengine.v1alpha.ManagedCertificateResponse>;
    /**
     * Full path to the AuthorizedCertificate resource in the API. Example: apps/myapp/authorizedCertificates/12345.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The full paths to user visible Domain Mapping resources that have this certificate mapped. Example: apps/myapp/domainMappings/example.com.This may not represent the full list of mapped domain mappings if the user does not have VIEWER permissions on all of the applications that have this certificate mapped. See domain_mappings_count for a complete count.Only returned by GET or LIST requests when specifically requested by the view=FULL_CERTIFICATE option.
     */
    public /*out*/ readonly visibleDomainMappings!: pulumi.Output<string[]>;

    /**
     * Create a AuthorizedCertificate resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AuthorizedCertificateArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.appId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'appId'");
            }
            resourceInputs["appId"] = args ? args.appId : undefined;
            resourceInputs["certificateRawData"] = args ? args.certificateRawData : undefined;
            resourceInputs["displayName"] = args ? args.displayName : undefined;
            resourceInputs["domainMappingsCount"] = undefined /*out*/;
            resourceInputs["domainNames"] = undefined /*out*/;
            resourceInputs["expireTime"] = undefined /*out*/;
            resourceInputs["managedCertificate"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["visibleDomainMappings"] = undefined /*out*/;
        } else {
            resourceInputs["certificateRawData"] = undefined /*out*/;
            resourceInputs["displayName"] = undefined /*out*/;
            resourceInputs["domainMappingsCount"] = undefined /*out*/;
            resourceInputs["domainNames"] = undefined /*out*/;
            resourceInputs["expireTime"] = undefined /*out*/;
            resourceInputs["managedCertificate"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["visibleDomainMappings"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(AuthorizedCertificate.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a AuthorizedCertificate resource.
 */
export interface AuthorizedCertificateArgs {
    appId: pulumi.Input<string>;
    /**
     * The SSL certificate serving the AuthorizedCertificate resource. This must be obtained independently from a certificate authority.
     */
    certificateRawData?: pulumi.Input<inputs.appengine.v1alpha.CertificateRawDataArgs>;
    /**
     * The user-specified display name of the certificate. This is not guaranteed to be unique. Example: My Certificate.
     */
    displayName?: pulumi.Input<string>;
}
