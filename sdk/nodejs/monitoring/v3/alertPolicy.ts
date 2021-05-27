// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates a new alerting policy.
 */
export class AlertPolicy extends pulumi.CustomResource {
    /**
     * Get an existing AlertPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): AlertPolicy {
        return new AlertPolicy(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:monitoring/v3:AlertPolicy';

    /**
     * Returns true if the given object is an instance of AlertPolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AlertPolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AlertPolicy.__pulumiType;
    }

    /**
     * How to combine the results of multiple conditions to determine if an incident should be opened. If condition_time_series_query_language is present, this must be COMBINE_UNSPECIFIED.
     */
    public readonly combiner!: pulumi.Output<string>;
    /**
     * A list of conditions for the policy. The conditions are combined by AND or OR according to the combiner field. If the combined conditions evaluate to true, then an incident is created. A policy can have from one to six conditions. If condition_time_series_query_language is present, it must be the only condition.
     */
    public readonly conditions!: pulumi.Output<outputs.monitoring.v3.ConditionResponse[]>;
    /**
     * A read-only record of the creation of the alerting policy. If provided in a call to create or update, this field will be ignored.
     */
    public readonly creationRecord!: pulumi.Output<outputs.monitoring.v3.MutationRecordResponse>;
    /**
     * A short name or phrase used to identify the policy in dashboards, notifications, and incidents. To avoid confusion, don't use the same display name for multiple policies in the same project. The name is limited to 512 Unicode characters.
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * Documentation that is included with notifications and incidents related to this policy. Best practice is for the documentation to include information to help responders understand, mitigate, escalate, and correct the underlying problems detected by the alerting policy. Notification channels that have limited capacity might not show this documentation.
     */
    public readonly documentation!: pulumi.Output<outputs.monitoring.v3.DocumentationResponse>;
    /**
     * Whether or not the policy is enabled. On write, the default interpretation if unset is that the policy is enabled. On read, clients should not make any assumption about the state if it has not been populated. The field should always be populated on List and Get operations, unless a field projection has been specified that strips it out.
     */
    public readonly enabled!: pulumi.Output<boolean>;
    /**
     * A read-only record of the most recent change to the alerting policy. If provided in a call to create or update, this field will be ignored.
     */
    public readonly mutationRecord!: pulumi.Output<outputs.monitoring.v3.MutationRecordResponse>;
    /**
     * Required if the policy exists. The resource name for this policy. The format is: projects/[PROJECT_ID_OR_NUMBER]/alertPolicies/[ALERT_POLICY_ID] [ALERT_POLICY_ID] is assigned by Stackdriver Monitoring when the policy is created. When calling the alertPolicies.create method, do not include the name field in the alerting policy passed as part of the request.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Identifies the notification channels to which notifications should be sent when incidents are opened or closed or when new violations occur on an already opened incident. Each element of this array corresponds to the name field in each of the NotificationChannel objects that are returned from the ListNotificationChannels method. The format of the entries in this field is: projects/[PROJECT_ID_OR_NUMBER]/notificationChannels/[CHANNEL_ID] 
     */
    public readonly notificationChannels!: pulumi.Output<string[]>;
    /**
     * User-supplied key/value data to be used for organizing and identifying the AlertPolicy objects.The field can contain up to 64 entries. Each key and value is limited to 63 Unicode characters or 128 bytes, whichever is smaller. Labels and values can contain only lowercase letters, numerals, underscores, and dashes. Keys must begin with a letter.
     */
    public readonly userLabels!: pulumi.Output<{[key: string]: string}>;
    /**
     * Read-only description of how the alert policy is invalid. OK if the alert policy is valid. If not OK, the alert policy will not generate incidents.
     */
    public readonly validity!: pulumi.Output<outputs.monitoring.v3.StatusResponse>;

    /**
     * Create a AlertPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AlertPolicyArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.project === undefined) && !opts.urn) {
                throw new Error("Missing required property 'project'");
            }
            inputs["combiner"] = args ? args.combiner : undefined;
            inputs["conditions"] = args ? args.conditions : undefined;
            inputs["creationRecord"] = args ? args.creationRecord : undefined;
            inputs["displayName"] = args ? args.displayName : undefined;
            inputs["documentation"] = args ? args.documentation : undefined;
            inputs["enabled"] = args ? args.enabled : undefined;
            inputs["mutationRecord"] = args ? args.mutationRecord : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["notificationChannels"] = args ? args.notificationChannels : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["userLabels"] = args ? args.userLabels : undefined;
            inputs["validity"] = args ? args.validity : undefined;
        } else {
            inputs["combiner"] = undefined /*out*/;
            inputs["conditions"] = undefined /*out*/;
            inputs["creationRecord"] = undefined /*out*/;
            inputs["displayName"] = undefined /*out*/;
            inputs["documentation"] = undefined /*out*/;
            inputs["enabled"] = undefined /*out*/;
            inputs["mutationRecord"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["notificationChannels"] = undefined /*out*/;
            inputs["userLabels"] = undefined /*out*/;
            inputs["validity"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(AlertPolicy.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a AlertPolicy resource.
 */
export interface AlertPolicyArgs {
    /**
     * How to combine the results of multiple conditions to determine if an incident should be opened. If condition_time_series_query_language is present, this must be COMBINE_UNSPECIFIED.
     */
    readonly combiner?: pulumi.Input<string>;
    /**
     * A list of conditions for the policy. The conditions are combined by AND or OR according to the combiner field. If the combined conditions evaluate to true, then an incident is created. A policy can have from one to six conditions. If condition_time_series_query_language is present, it must be the only condition.
     */
    readonly conditions?: pulumi.Input<pulumi.Input<inputs.monitoring.v3.ConditionArgs>[]>;
    /**
     * A read-only record of the creation of the alerting policy. If provided in a call to create or update, this field will be ignored.
     */
    readonly creationRecord?: pulumi.Input<inputs.monitoring.v3.MutationRecordArgs>;
    /**
     * A short name or phrase used to identify the policy in dashboards, notifications, and incidents. To avoid confusion, don't use the same display name for multiple policies in the same project. The name is limited to 512 Unicode characters.
     */
    readonly displayName?: pulumi.Input<string>;
    /**
     * Documentation that is included with notifications and incidents related to this policy. Best practice is for the documentation to include information to help responders understand, mitigate, escalate, and correct the underlying problems detected by the alerting policy. Notification channels that have limited capacity might not show this documentation.
     */
    readonly documentation?: pulumi.Input<inputs.monitoring.v3.DocumentationArgs>;
    /**
     * Whether or not the policy is enabled. On write, the default interpretation if unset is that the policy is enabled. On read, clients should not make any assumption about the state if it has not been populated. The field should always be populated on List and Get operations, unless a field projection has been specified that strips it out.
     */
    readonly enabled?: pulumi.Input<boolean>;
    /**
     * A read-only record of the most recent change to the alerting policy. If provided in a call to create or update, this field will be ignored.
     */
    readonly mutationRecord?: pulumi.Input<inputs.monitoring.v3.MutationRecordArgs>;
    /**
     * Required if the policy exists. The resource name for this policy. The format is: projects/[PROJECT_ID_OR_NUMBER]/alertPolicies/[ALERT_POLICY_ID] [ALERT_POLICY_ID] is assigned by Stackdriver Monitoring when the policy is created. When calling the alertPolicies.create method, do not include the name field in the alerting policy passed as part of the request.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Identifies the notification channels to which notifications should be sent when incidents are opened or closed or when new violations occur on an already opened incident. Each element of this array corresponds to the name field in each of the NotificationChannel objects that are returned from the ListNotificationChannels method. The format of the entries in this field is: projects/[PROJECT_ID_OR_NUMBER]/notificationChannels/[CHANNEL_ID] 
     */
    readonly notificationChannels?: pulumi.Input<pulumi.Input<string>[]>;
    readonly project: pulumi.Input<string>;
    /**
     * User-supplied key/value data to be used for organizing and identifying the AlertPolicy objects.The field can contain up to 64 entries. Each key and value is limited to 63 Unicode characters or 128 bytes, whichever is smaller. Labels and values can contain only lowercase letters, numerals, underscores, and dashes. Keys must begin with a letter.
     */
    readonly userLabels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Read-only description of how the alert policy is invalid. OK if the alert policy is valid. If not OK, the alert policy will not generate incidents.
     */
    readonly validity?: pulumi.Input<inputs.monitoring.v3.StatusArgs>;
}
