// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Creates a source.
 * Note - this resource's API doesn't support deletion. When deleted, the resource will persist
 * on Google Cloud even though it will be deleted from Pulumi state.
 */
export class Source extends pulumi.CustomResource {
    /**
     * Get an existing Source resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Source {
        return new Source(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:securitycenter/v1beta1:Source';

    /**
     * Returns true if the given object is an instance of Source.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Source {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Source.__pulumiType;
    }

    /**
     * The description of the source (max of 1024 characters). Example: "Web Security Scanner is a web security scanner for common vulnerabilities in App Engine applications. It can automatically scan and detect four common vulnerabilities, including cross-site-scripting (XSS), Flash injection, mixed content (HTTP in HTTPS), and outdated/insecure libraries."
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * The source's display name. A source's display name must be unique amongst its siblings, for example, two sources with the same parent can't share the same display name. The display name must have a length between 1 and 64 characters (inclusive).
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * The relative resource name of this source. See: https://cloud.google.com/apis/design/resource_names#relative_resource_name Example: "organizations/{organization_id}/sources/{source_id}"
     */
    public readonly name!: pulumi.Output<string>;

    /**
     * Create a Source resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SourceArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.organizationId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'organizationId'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["displayName"] = args ? args.displayName : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["organizationId"] = args ? args.organizationId : undefined;
        } else {
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["displayName"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Source.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Source resource.
 */
export interface SourceArgs {
    /**
     * The description of the source (max of 1024 characters). Example: "Web Security Scanner is a web security scanner for common vulnerabilities in App Engine applications. It can automatically scan and detect four common vulnerabilities, including cross-site-scripting (XSS), Flash injection, mixed content (HTTP in HTTPS), and outdated/insecure libraries."
     */
    description?: pulumi.Input<string>;
    /**
     * The source's display name. A source's display name must be unique amongst its siblings, for example, two sources with the same parent can't share the same display name. The display name must have a length between 1 and 64 characters (inclusive).
     */
    displayName?: pulumi.Input<string>;
    /**
     * The relative resource name of this source. See: https://cloud.google.com/apis/design/resource_names#relative_resource_name Example: "organizations/{organization_id}/sources/{source_id}"
     */
    name?: pulumi.Input<string>;
    organizationId: pulumi.Input<string>;
}
