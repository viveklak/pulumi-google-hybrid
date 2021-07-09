// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Retrieves the specified SecuritySettings. The returned settings may be stale by up to 1 minute.
 */
export function getSecuritySetting(args: GetSecuritySettingArgs, opts?: pulumi.InvokeOptions): Promise<GetSecuritySettingResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("google-native:dialogflow/v3beta1:getSecuritySetting", {
        "location": args.location,
        "project": args.project,
        "securitySettingId": args.securitySettingId,
    }, opts);
}

export interface GetSecuritySettingArgs {
    location: string;
    project: string;
    securitySettingId: string;
}

export interface GetSecuritySettingResult {
    /**
     * The human-readable name of the security settings, unique within the location.
     */
    readonly displayName: string;
    /**
     * DLP inspect template name. Use this template to define inspect base settings. If empty, we use the default DLP inspect config. The template name will have one of the following formats: `projects/PROJECT_ID/inspectTemplates/TEMPLATE_ID` OR `organizations/ORGANIZATION_ID/inspectTemplates/TEMPLATE_ID`
     */
    readonly inspectTemplate: string;
    /**
     * Resource name of the settings. Format: `projects//locations//securitySettings/`.
     */
    readonly name: string;
    /**
     * List of types of data to remove when retention settings triggers purge.
     */
    readonly purgeDataTypes: string[];
    /**
     * Defines on what data we apply redaction. Note that we don't redact data to which we don't have access, e.g., Stackdriver logs.
     */
    readonly redactionScope: string;
    /**
     * Strategy that defines how we do redaction.
     */
    readonly redactionStrategy: string;
    /**
     * Retains the data for the specified number of days. User must Set a value lower than Dialogflow's default 30d TTL. Setting a value higher than that has no effect. A missing value or setting to 0 also means we use Dialogflow's default TTL.
     */
    readonly retentionWindowDays: number;
}
