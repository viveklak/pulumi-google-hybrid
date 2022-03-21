// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Uploads a ZIP-formatted shared flow configuration bundle to an organization. If the shared flow already exists, this creates a new revision of it. If the shared flow does not exist, this creates it. Once imported, the shared flow revision must be deployed before it can be accessed at runtime. The size limit of a shared flow bundle is 15 MB.
 */
export class Sharedflow extends pulumi.CustomResource {
    /**
     * Get an existing Sharedflow resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Sharedflow {
        return new Sharedflow(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:apigee/v1:Sharedflow';

    /**
     * Returns true if the given object is an instance of Sharedflow.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Sharedflow {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Sharedflow.__pulumiType;
    }

    /**
     * The id of the most recently created revision for this shared flow.
     */
    public /*out*/ readonly latestRevisionId!: pulumi.Output<string>;
    /**
     * Metadata describing the shared flow.
     */
    public /*out*/ readonly metaData!: pulumi.Output<outputs.apigee.v1.GoogleCloudApigeeV1EntityMetadataResponse>;
    /**
     * The ID of the shared flow.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * A list of revisions of this shared flow.
     */
    public /*out*/ readonly revision!: pulumi.Output<string[]>;

    /**
     * Create a Sharedflow resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SharedflowArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.action === undefined) && !opts.urn) {
                throw new Error("Missing required property 'action'");
            }
            if ((!args || args.organizationId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'organizationId'");
            }
            resourceInputs["action"] = args ? args.action : undefined;
            resourceInputs["contentType"] = args ? args.contentType : undefined;
            resourceInputs["data"] = args ? args.data : undefined;
            resourceInputs["extensions"] = args ? args.extensions : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["organizationId"] = args ? args.organizationId : undefined;
            resourceInputs["latestRevisionId"] = undefined /*out*/;
            resourceInputs["metaData"] = undefined /*out*/;
            resourceInputs["revision"] = undefined /*out*/;
        } else {
            resourceInputs["latestRevisionId"] = undefined /*out*/;
            resourceInputs["metaData"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["revision"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Sharedflow.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Sharedflow resource.
 */
export interface SharedflowArgs {
    /**
     * Required. Must be set to either `import` or `validate`.
     */
    action: pulumi.Input<string>;
    /**
     * The HTTP Content-Type header value specifying the content type of the body.
     */
    contentType?: pulumi.Input<string>;
    /**
     * The HTTP request/response body as raw binary.
     */
    data?: pulumi.Input<string>;
    /**
     * Application specific response metadata. Must be set in the first response for streaming APIs.
     */
    extensions?: pulumi.Input<pulumi.Input<{[key: string]: pulumi.Input<string>}>[]>;
    /**
     * Required. The name to give the shared flow
     */
    name?: pulumi.Input<string>;
    organizationId: pulumi.Input<string>;
}
