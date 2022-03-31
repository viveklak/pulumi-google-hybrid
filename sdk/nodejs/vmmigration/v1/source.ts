// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates a new Source in a given project and location.
 * Auto-naming is currently not supported for this resource.
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
    public static readonly __pulumiType = 'google-native:vmmigration/v1:Source';

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
     * The create time timestamp.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * User-provided description of the source.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * The labels of the source.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string}>;
    /**
     * The Source name.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The update time timestamp.
     */
    public /*out*/ readonly updateTime!: pulumi.Output<string>;
    /**
     * Vmware type source details.
     */
    public readonly vmware!: pulumi.Output<outputs.vmmigration.v1.VmwareSourceDetailsResponse>;

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
            if ((!args || args.sourceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'sourceId'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["requestId"] = args ? args.requestId : undefined;
            resourceInputs["sourceId"] = args ? args.sourceId : undefined;
            resourceInputs["vmware"] = args ? args.vmware : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        } else {
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["labels"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
            resourceInputs["vmware"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Source.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Source resource.
 */
export interface SourceArgs {
    /**
     * User-provided description of the source.
     */
    description?: pulumi.Input<string>;
    /**
     * The labels of the source.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    location?: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    /**
     * A request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. The server will guarantee that for at least 60 minutes since the first request. For example, consider a situation where you make an initial request and t he request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported (00000000-0000-0000-0000-000000000000).
     */
    requestId?: pulumi.Input<string>;
    /**
     * Required. The source identifier.
     */
    sourceId: pulumi.Input<string>;
    /**
     * Vmware type source details.
     */
    vmware?: pulumi.Input<inputs.vmmigration.v1.VmwareSourceDetailsArgs>;
}
