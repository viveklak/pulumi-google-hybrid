// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates an instruction for how data should be labeled.
 * Auto-naming is currently not supported for this resource.
 */
export class Instruction extends pulumi.CustomResource {
    /**
     * Get an existing Instruction resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Instruction {
        return new Instruction(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:datalabeling/v1beta1:Instruction';

    /**
     * Returns true if the given object is an instance of Instruction.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Instruction {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Instruction.__pulumiType;
    }

    /**
     * The names of any related resources that are blocking changes to the instruction.
     */
    public /*out*/ readonly blockingResources!: pulumi.Output<string[]>;
    /**
     * Creation time of instruction.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * Deprecated: this instruction format is not supported any more. Instruction from a CSV file, such as for classification task. The CSV file should have exact two columns, in the following format: * The first column is labeled data, such as an image reference, text. * The second column is comma separated labels associated with data.
     *
     * @deprecated Deprecated: this instruction format is not supported any more. Instruction from a CSV file, such as for classification task. The CSV file should have exact two columns, in the following format: * The first column is labeled data, such as an image reference, text. * The second column is comma separated labels associated with data.
     */
    public readonly csvInstruction!: pulumi.Output<outputs.datalabeling.v1beta1.GoogleCloudDatalabelingV1beta1CsvInstructionResponse>;
    /**
     * The data type of this instruction.
     */
    public readonly dataType!: pulumi.Output<string>;
    /**
     * Optional. User-provided description of the instruction. The description can be up to 10000 characters long.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * The display name of the instruction. Maximum of 64 characters.
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * Instruction resource name, format: projects/{project_id}/instructions/{instruction_id}
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Instruction from a PDF document. The PDF should be in a Cloud Storage bucket.
     */
    public readonly pdfInstruction!: pulumi.Output<outputs.datalabeling.v1beta1.GoogleCloudDatalabelingV1beta1PdfInstructionResponse>;
    /**
     * Last update time of instruction.
     */
    public /*out*/ readonly updateTime!: pulumi.Output<string>;

    /**
     * Create a Instruction resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: InstructionArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.dataType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'dataType'");
            }
            if ((!args || args.displayName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'displayName'");
            }
            resourceInputs["csvInstruction"] = args ? args.csvInstruction : undefined;
            resourceInputs["dataType"] = args ? args.dataType : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["displayName"] = args ? args.displayName : undefined;
            resourceInputs["pdfInstruction"] = args ? args.pdfInstruction : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["blockingResources"] = undefined /*out*/;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        } else {
            resourceInputs["blockingResources"] = undefined /*out*/;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["csvInstruction"] = undefined /*out*/;
            resourceInputs["dataType"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["displayName"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["pdfInstruction"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Instruction.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Instruction resource.
 */
export interface InstructionArgs {
    /**
     * Deprecated: this instruction format is not supported any more. Instruction from a CSV file, such as for classification task. The CSV file should have exact two columns, in the following format: * The first column is labeled data, such as an image reference, text. * The second column is comma separated labels associated with data.
     *
     * @deprecated Deprecated: this instruction format is not supported any more. Instruction from a CSV file, such as for classification task. The CSV file should have exact two columns, in the following format: * The first column is labeled data, such as an image reference, text. * The second column is comma separated labels associated with data.
     */
    csvInstruction?: pulumi.Input<inputs.datalabeling.v1beta1.GoogleCloudDatalabelingV1beta1CsvInstructionArgs>;
    /**
     * The data type of this instruction.
     */
    dataType: pulumi.Input<enums.datalabeling.v1beta1.InstructionDataType>;
    /**
     * Optional. User-provided description of the instruction. The description can be up to 10000 characters long.
     */
    description?: pulumi.Input<string>;
    /**
     * The display name of the instruction. Maximum of 64 characters.
     */
    displayName: pulumi.Input<string>;
    /**
     * Instruction from a PDF document. The PDF should be in a Cloud Storage bucket.
     */
    pdfInstruction?: pulumi.Input<inputs.datalabeling.v1beta1.GoogleCloudDatalabelingV1beta1PdfInstructionArgs>;
    project?: pulumi.Input<string>;
}
