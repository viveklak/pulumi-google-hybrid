// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates a composite type.
 */
export class CompositeType extends pulumi.CustomResource {
    /**
     * Get an existing CompositeType resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): CompositeType {
        return new CompositeType(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:deploymentmanager/v2beta:CompositeType';

    /**
     * Returns true if the given object is an instance of CompositeType.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is CompositeType {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === CompositeType.__pulumiType;
    }

    /**
     * An optional textual description of the resource; provided by the client when the resource is created.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * Creation timestamp in RFC3339 text format.
     */
    public /*out*/ readonly insertTime!: pulumi.Output<string>;
    /**
     * Map of labels; provided by the client when the resource is created or updated. Specifically: Label keys must be between 1 and 63 characters long and must conform to the following regular expression: `[a-z]([-a-z0-9]*[a-z0-9])?` Label values must be between 0 and 63 characters long and must conform to the regular expression `([a-z]([-a-z0-9]*[a-z0-9])?)?`.
     */
    public readonly labels!: pulumi.Output<outputs.deploymentmanager.v2beta.CompositeTypeLabelEntryResponse[]>;
    /**
     * Name of the composite type, must follow the expression: `[a-z]([-a-z0-9_.]{0,61}[a-z0-9])?`.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The Operation that most recently ran, or is currently running, on this composite type.
     */
    public /*out*/ readonly operation!: pulumi.Output<outputs.deploymentmanager.v2beta.OperationResponse>;
    /**
     * Server defined URL for the resource.
     */
    public /*out*/ readonly selfLink!: pulumi.Output<string>;
    public readonly status!: pulumi.Output<string>;
    /**
     * Files for the template type.
     */
    public readonly templateContents!: pulumi.Output<outputs.deploymentmanager.v2beta.TemplateContentsResponse>;

    /**
     * Create a CompositeType resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: CompositeTypeArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["id"] = args ? args.id : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["status"] = args ? args.status : undefined;
            resourceInputs["templateContents"] = args ? args.templateContents : undefined;
            resourceInputs["insertTime"] = undefined /*out*/;
            resourceInputs["operation"] = undefined /*out*/;
            resourceInputs["selfLink"] = undefined /*out*/;
        } else {
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["insertTime"] = undefined /*out*/;
            resourceInputs["labels"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["operation"] = undefined /*out*/;
            resourceInputs["selfLink"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["templateContents"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(CompositeType.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a CompositeType resource.
 */
export interface CompositeTypeArgs {
    /**
     * An optional textual description of the resource; provided by the client when the resource is created.
     */
    description?: pulumi.Input<string>;
    id?: pulumi.Input<string>;
    /**
     * Map of labels; provided by the client when the resource is created or updated. Specifically: Label keys must be between 1 and 63 characters long and must conform to the following regular expression: `[a-z]([-a-z0-9]*[a-z0-9])?` Label values must be between 0 and 63 characters long and must conform to the regular expression `([a-z]([-a-z0-9]*[a-z0-9])?)?`.
     */
    labels?: pulumi.Input<pulumi.Input<inputs.deploymentmanager.v2beta.CompositeTypeLabelEntryArgs>[]>;
    /**
     * Name of the composite type, must follow the expression: `[a-z]([-a-z0-9_.]{0,61}[a-z0-9])?`.
     */
    name?: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    status?: pulumi.Input<enums.deploymentmanager.v2beta.CompositeTypeStatus>;
    /**
     * Files for the template type.
     */
    templateContents?: pulumi.Input<inputs.deploymentmanager.v2beta.TemplateContentsArgs>;
}
