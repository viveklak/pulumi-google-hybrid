// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Create security settings in the specified location.
 */
export class SecuritySetting extends pulumi.CustomResource {
    /**
     * Get an existing SecuritySetting resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): SecuritySetting {
        return new SecuritySetting(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:dialogflow/v3beta1:SecuritySetting';

    /**
     * Returns true if the given object is an instance of SecuritySetting.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SecuritySetting {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SecuritySetting.__pulumiType;
    }

    /**
     * The human-readable name of the security settings, unique within the location.
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * DLP inspect template name. Use this template to define inspect base settings. If empty, we use the default DLP inspect config. The template name will have one of the following formats: `projects/PROJECT_ID/inspectTemplates/TEMPLATE_ID` OR `organizations/ORGANIZATION_ID/inspectTemplates/TEMPLATE_ID`
     */
    public readonly inspectTemplate!: pulumi.Output<string>;
    /**
     * Resource name of the settings. Format: `projects//locations//securitySettings/`.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * List of types of data to remove when retention settings triggers purge.
     */
    public readonly purgeDataTypes!: pulumi.Output<string[]>;
    /**
     * Defines on what data we apply redaction. Note that we don't redact data to which we don't have access, e.g., Stackdriver logs.
     */
    public readonly redactionScope!: pulumi.Output<string>;
    /**
     * Strategy that defines how we do redaction.
     */
    public readonly redactionStrategy!: pulumi.Output<string>;
    /**
     * Retains the data for the specified number of days. User must Set a value lower than Dialogflow's default 30d TTL. Setting a value higher than that has no effect. A missing value or setting to 0 also means we use Dialogflow's default TTL.
     */
    public readonly retentionWindowDays!: pulumi.Output<number>;

    /**
     * Create a SecuritySetting resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SecuritySettingArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.displayName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'displayName'");
            }
            if ((!args || args.location === undefined) && !opts.urn) {
                throw new Error("Missing required property 'location'");
            }
            if ((!args || args.name === undefined) && !opts.urn) {
                throw new Error("Missing required property 'name'");
            }
            if ((!args || args.project === undefined) && !opts.urn) {
                throw new Error("Missing required property 'project'");
            }
            inputs["displayName"] = args ? args.displayName : undefined;
            inputs["inspectTemplate"] = args ? args.inspectTemplate : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["purgeDataTypes"] = args ? args.purgeDataTypes : undefined;
            inputs["redactionScope"] = args ? args.redactionScope : undefined;
            inputs["redactionStrategy"] = args ? args.redactionStrategy : undefined;
            inputs["retentionWindowDays"] = args ? args.retentionWindowDays : undefined;
        } else {
            inputs["displayName"] = undefined /*out*/;
            inputs["inspectTemplate"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["purgeDataTypes"] = undefined /*out*/;
            inputs["redactionScope"] = undefined /*out*/;
            inputs["redactionStrategy"] = undefined /*out*/;
            inputs["retentionWindowDays"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(SecuritySetting.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a SecuritySetting resource.
 */
export interface SecuritySettingArgs {
    /**
     * The human-readable name of the security settings, unique within the location.
     */
    displayName: pulumi.Input<string>;
    /**
     * DLP inspect template name. Use this template to define inspect base settings. If empty, we use the default DLP inspect config. The template name will have one of the following formats: `projects/PROJECT_ID/inspectTemplates/TEMPLATE_ID` OR `organizations/ORGANIZATION_ID/inspectTemplates/TEMPLATE_ID`
     */
    inspectTemplate?: pulumi.Input<string>;
    location: pulumi.Input<string>;
    /**
     * Resource name of the settings. Format: `projects//locations//securitySettings/`.
     */
    name: pulumi.Input<string>;
    project: pulumi.Input<string>;
    /**
     * List of types of data to remove when retention settings triggers purge.
     */
    purgeDataTypes?: pulumi.Input<pulumi.Input<enums.dialogflow.v3beta1.SecuritySettingPurgeDataTypesItem>[]>;
    /**
     * Defines on what data we apply redaction. Note that we don't redact data to which we don't have access, e.g., Stackdriver logs.
     */
    redactionScope?: pulumi.Input<enums.dialogflow.v3beta1.SecuritySettingRedactionScope>;
    /**
     * Strategy that defines how we do redaction.
     */
    redactionStrategy?: pulumi.Input<enums.dialogflow.v3beta1.SecuritySettingRedactionStrategy>;
    /**
     * Retains the data for the specified number of days. User must Set a value lower than Dialogflow's default 30d TTL. Setting a value higher than that has no effect. A missing value or setting to 0 also means we use Dialogflow's default TTL.
     */
    retentionWindowDays?: pulumi.Input<number>;
}
