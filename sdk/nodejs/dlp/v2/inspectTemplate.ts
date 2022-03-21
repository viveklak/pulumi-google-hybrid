// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates an InspectTemplate for re-using frequently used configuration for inspecting content, images, and storage. See https://cloud.google.com/dlp/docs/creating-templates to learn more.
 * Auto-naming is currently not supported for this resource.
 */
export class InspectTemplate extends pulumi.CustomResource {
    /**
     * Get an existing InspectTemplate resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): InspectTemplate {
        return new InspectTemplate(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:dlp/v2:InspectTemplate';

    /**
     * Returns true if the given object is an instance of InspectTemplate.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is InspectTemplate {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === InspectTemplate.__pulumiType;
    }

    /**
     * The creation timestamp of an inspectTemplate.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * Short description (max 256 chars).
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * Display name (max 256 chars).
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * The core content of the template. Configuration of the scanning process.
     */
    public readonly inspectConfig!: pulumi.Output<outputs.dlp.v2.GooglePrivacyDlpV2InspectConfigResponse>;
    /**
     * The template name. The template will have one of the following formats: `projects/PROJECT_ID/inspectTemplates/TEMPLATE_ID` OR `organizations/ORGANIZATION_ID/inspectTemplates/TEMPLATE_ID`;
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The last update timestamp of an inspectTemplate.
     */
    public /*out*/ readonly updateTime!: pulumi.Output<string>;

    /**
     * Create a InspectTemplate resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: InspectTemplateArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["displayName"] = args ? args.displayName : undefined;
            resourceInputs["inspectConfig"] = args ? args.inspectConfig : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["templateId"] = args ? args.templateId : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        } else {
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["displayName"] = undefined /*out*/;
            resourceInputs["inspectConfig"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(InspectTemplate.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a InspectTemplate resource.
 */
export interface InspectTemplateArgs {
    /**
     * Short description (max 256 chars).
     */
    description?: pulumi.Input<string>;
    /**
     * Display name (max 256 chars).
     */
    displayName?: pulumi.Input<string>;
    /**
     * The core content of the template. Configuration of the scanning process.
     */
    inspectConfig?: pulumi.Input<inputs.dlp.v2.GooglePrivacyDlpV2InspectConfigArgs>;
    /**
     * Deprecated. This field has no effect.
     *
     * @deprecated Deprecated. This field has no effect.
     */
    location?: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    /**
     * The template id can contain uppercase and lowercase letters, numbers, and hyphens; that is, it must match the regular expression: `[a-zA-Z\d-_]+`. The maximum length is 100 characters. Can be empty to allow the system to generate one.
     */
    templateId?: pulumi.Input<string>;
}
