// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates a TargetSslProxy resource in the specified project using the data included in the request.
 */
export class TargetSslProxy extends pulumi.CustomResource {
    /**
     * Get an existing TargetSslProxy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): TargetSslProxy {
        return new TargetSslProxy(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:compute/beta:TargetSslProxy';

    /**
     * Returns true if the given object is an instance of TargetSslProxy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is TargetSslProxy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === TargetSslProxy.__pulumiType;
    }

    /**
     * Creation timestamp in RFC3339 text format.
     */
    public /*out*/ readonly creationTimestamp!: pulumi.Output<string>;
    /**
     * An optional description of this resource. Provide this property when you create the resource.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * Type of the resource. Always compute#targetSslProxy for target SSL proxies.
     */
    public /*out*/ readonly kind!: pulumi.Output<string>;
    /**
     * Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Specifies the type of proxy header to append before sending data to the backend, either NONE or PROXY_V1. The default is NONE.
     */
    public readonly proxyHeader!: pulumi.Output<string>;
    /**
     * Server-defined URL for the resource.
     */
    public /*out*/ readonly selfLink!: pulumi.Output<string>;
    /**
     * URL to the BackendService resource.
     */
    public readonly service!: pulumi.Output<string>;
    /**
     * URLs to SslCertificate resources that are used to authenticate connections to Backends. At least one SSL certificate must be specified. Currently, you may specify up to 15 SSL certificates.
     */
    public readonly sslCertificates!: pulumi.Output<string[]>;
    /**
     * URL of SslPolicy resource that will be associated with the TargetSslProxy resource. If not set, the TargetSslProxy resource will not have any SSL policy configured.
     */
    public readonly sslPolicy!: pulumi.Output<string>;

    /**
     * Create a TargetSslProxy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TargetSslProxyArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.project === undefined) && !opts.urn) {
                throw new Error("Missing required property 'project'");
            }
            inputs["description"] = args ? args.description : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["proxyHeader"] = args ? args.proxyHeader : undefined;
            inputs["requestId"] = args ? args.requestId : undefined;
            inputs["service"] = args ? args.service : undefined;
            inputs["sslCertificates"] = args ? args.sslCertificates : undefined;
            inputs["sslPolicy"] = args ? args.sslPolicy : undefined;
            inputs["creationTimestamp"] = undefined /*out*/;
            inputs["kind"] = undefined /*out*/;
            inputs["selfLink"] = undefined /*out*/;
        } else {
            inputs["creationTimestamp"] = undefined /*out*/;
            inputs["description"] = undefined /*out*/;
            inputs["kind"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["proxyHeader"] = undefined /*out*/;
            inputs["selfLink"] = undefined /*out*/;
            inputs["service"] = undefined /*out*/;
            inputs["sslCertificates"] = undefined /*out*/;
            inputs["sslPolicy"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(TargetSslProxy.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a TargetSslProxy resource.
 */
export interface TargetSslProxyArgs {
    /**
     * An optional description of this resource. Provide this property when you create the resource.
     */
    description?: pulumi.Input<string>;
    /**
     * Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
     */
    name?: pulumi.Input<string>;
    project: pulumi.Input<string>;
    /**
     * Specifies the type of proxy header to append before sending data to the backend, either NONE or PROXY_V1. The default is NONE.
     */
    proxyHeader?: pulumi.Input<enums.compute.beta.TargetSslProxyProxyHeader>;
    requestId?: pulumi.Input<string>;
    /**
     * URL to the BackendService resource.
     */
    service?: pulumi.Input<string>;
    /**
     * URLs to SslCertificate resources that are used to authenticate connections to Backends. At least one SSL certificate must be specified. Currently, you may specify up to 15 SSL certificates.
     */
    sslCertificates?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * URL of SslPolicy resource that will be associated with the TargetSslProxy resource. If not set, the TargetSslProxy resource will not have any SSL policy configured.
     */
    sslPolicy?: pulumi.Input<string>;
}
