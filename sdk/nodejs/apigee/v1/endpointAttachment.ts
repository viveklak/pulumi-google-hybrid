// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Creates an endpoint attachment. **Note:** Not supported for Apigee hybrid.
 * Auto-naming is currently not supported for this resource.
 */
export class EndpointAttachment extends pulumi.CustomResource {
    /**
     * Get an existing EndpointAttachment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): EndpointAttachment {
        return new EndpointAttachment(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:apigee/v1:EndpointAttachment';

    /**
     * Returns true if the given object is an instance of EndpointAttachment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is EndpointAttachment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === EndpointAttachment.__pulumiType;
    }

    /**
     * Host that can be used in either the HTTP target endpoint directly or as the host in target server.
     */
    public /*out*/ readonly host!: pulumi.Output<string>;
    /**
     * Location of the endpoint attachment.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * Name of the endpoint attachment. Use the following structure in your request: `organizations/{org}/endpointAttachments/{endpoint_attachment}`
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Format: projects/*&#47;regions/*&#47;serviceAttachments/*
     */
    public readonly serviceAttachment!: pulumi.Output<string>;

    /**
     * Create a EndpointAttachment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: EndpointAttachmentArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.organizationId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'organizationId'");
            }
            resourceInputs["endpointAttachmentId"] = args ? args.endpointAttachmentId : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["organizationId"] = args ? args.organizationId : undefined;
            resourceInputs["serviceAttachment"] = args ? args.serviceAttachment : undefined;
            resourceInputs["host"] = undefined /*out*/;
        } else {
            resourceInputs["host"] = undefined /*out*/;
            resourceInputs["location"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["serviceAttachment"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(EndpointAttachment.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a EndpointAttachment resource.
 */
export interface EndpointAttachmentArgs {
    /**
     * ID to use for the endpoint attachment. The ID can contain lowercase letters and numbers, must start with a letter, and must be 1-20 characters in length.
     */
    endpointAttachmentId?: pulumi.Input<string>;
    /**
     * Location of the endpoint attachment.
     */
    location?: pulumi.Input<string>;
    /**
     * Name of the endpoint attachment. Use the following structure in your request: `organizations/{org}/endpointAttachments/{endpoint_attachment}`
     */
    name?: pulumi.Input<string>;
    organizationId: pulumi.Input<string>;
    /**
     * Format: projects/*&#47;regions/*&#47;serviceAttachments/*
     */
    serviceAttachment?: pulumi.Input<string>;
}
