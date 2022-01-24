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
    public static readonly __pulumiType = 'google-native:vmmigration/v1alpha1:Source';

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
     * Provides details on the state of the Source in case of an error.
     */
    public /*out*/ readonly error!: pulumi.Output<outputs.vmmigration.v1alpha1.StatusResponse>;
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
    public readonly vmware!: pulumi.Output<outputs.vmmigration.v1alpha1.VmwareSourceDetailsResponse>;

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
            resourceInputs["error"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        } else {
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["error"] = undefined /*out*/;
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
    requestId?: pulumi.Input<string>;
    sourceId: pulumi.Input<string>;
    /**
     * Vmware type source details.
     */
    vmware?: pulumi.Input<inputs.vmmigration.v1alpha1.VmwareSourceDetailsArgs>;
}
