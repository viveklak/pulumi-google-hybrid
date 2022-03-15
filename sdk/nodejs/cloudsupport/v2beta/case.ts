// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Create a new case and associate it with the given Cloud resource.
 * Note - this resource's API doesn't support deletion. When deleted, the resource will persist
 * on Google Cloud even though it will be deleted from Pulumi state.
 */
export class Case extends pulumi.CustomResource {
    /**
     * Get an existing Case resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Case {
        return new Case(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:cloudsupport/v2beta:Case';

    /**
     * Returns true if the given object is an instance of Case.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Case {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Case.__pulumiType;
    }

    /**
     * The issue classification applicable to this case.
     */
    public readonly classification!: pulumi.Output<outputs.cloudsupport.v2beta.CaseClassificationResponse>;
    /**
     * The time this case was created.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * The user who created the case. Note: The name and email will be obfuscated if the case was created by Google Support.
     */
    public readonly creator!: pulumi.Output<outputs.cloudsupport.v2beta.ActorResponse>;
    /**
     * A broad description of the issue.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * The short summary of the issue reported in this case.
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * Whether the case is currently escalated.
     */
    public readonly escalated!: pulumi.Output<boolean>;
    /**
     * The resource name for the case.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The priority of this case. If this is set, do not set severity.
     */
    public readonly priority!: pulumi.Output<string>;
    /**
     * The severity of this case. Deprecated. Use priority instead.
     *
     * @deprecated The severity of this case. Deprecated. Use priority instead.
     */
    public readonly severity!: pulumi.Output<string>;
    /**
     * The current status of the support case.
     */
    public /*out*/ readonly state!: pulumi.Output<string>;
    /**
     * The email addresses to receive updates on this case.
     */
    public readonly subscriberEmailAddresses!: pulumi.Output<string[]>;
    /**
     * Whether this case was created for internal API testing and should not be acted on by the support team.
     */
    public readonly testCase!: pulumi.Output<boolean>;
    /**
     * The timezone of the user who created the support case. It should be in a format IANA recognizes: https://www.iana.org/time-zones. There is no additional validation done by the API.
     */
    public readonly timeZone!: pulumi.Output<string>;
    /**
     * The time this case was last updated.
     */
    public /*out*/ readonly updateTime!: pulumi.Output<string>;

    /**
     * Create a Case resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CaseArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.v2betaId1 === undefined) && !opts.urn) {
                throw new Error("Missing required property 'v2betaId1'");
            }
            if ((!args || args.v2betumId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'v2betumId'");
            }
            resourceInputs["classification"] = args ? args.classification : undefined;
            resourceInputs["creator"] = args ? args.creator : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["displayName"] = args ? args.displayName : undefined;
            resourceInputs["escalated"] = args ? args.escalated : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["priority"] = args ? args.priority : undefined;
            resourceInputs["severity"] = args ? args.severity : undefined;
            resourceInputs["subscriberEmailAddresses"] = args ? args.subscriberEmailAddresses : undefined;
            resourceInputs["testCase"] = args ? args.testCase : undefined;
            resourceInputs["timeZone"] = args ? args.timeZone : undefined;
            resourceInputs["v2betaId1"] = args ? args.v2betaId1 : undefined;
            resourceInputs["v2betumId"] = args ? args.v2betumId : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        } else {
            resourceInputs["classification"] = undefined /*out*/;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["creator"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["displayName"] = undefined /*out*/;
            resourceInputs["escalated"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["priority"] = undefined /*out*/;
            resourceInputs["severity"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
            resourceInputs["subscriberEmailAddresses"] = undefined /*out*/;
            resourceInputs["testCase"] = undefined /*out*/;
            resourceInputs["timeZone"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Case.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Case resource.
 */
export interface CaseArgs {
    /**
     * The issue classification applicable to this case.
     */
    classification?: pulumi.Input<inputs.cloudsupport.v2beta.CaseClassificationArgs>;
    /**
     * The user who created the case. Note: The name and email will be obfuscated if the case was created by Google Support.
     */
    creator?: pulumi.Input<inputs.cloudsupport.v2beta.ActorArgs>;
    /**
     * A broad description of the issue.
     */
    description?: pulumi.Input<string>;
    /**
     * The short summary of the issue reported in this case.
     */
    displayName?: pulumi.Input<string>;
    /**
     * Whether the case is currently escalated.
     */
    escalated?: pulumi.Input<boolean>;
    /**
     * The resource name for the case.
     */
    name?: pulumi.Input<string>;
    /**
     * The priority of this case. If this is set, do not set severity.
     */
    priority?: pulumi.Input<enums.cloudsupport.v2beta.CasePriority>;
    /**
     * The severity of this case. Deprecated. Use priority instead.
     *
     * @deprecated The severity of this case. Deprecated. Use priority instead.
     */
    severity?: pulumi.Input<enums.cloudsupport.v2beta.CaseSeverity>;
    /**
     * The email addresses to receive updates on this case.
     */
    subscriberEmailAddresses?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Whether this case was created for internal API testing and should not be acted on by the support team.
     */
    testCase?: pulumi.Input<boolean>;
    /**
     * The timezone of the user who created the support case. It should be in a format IANA recognizes: https://www.iana.org/time-zones. There is no additional validation done by the API.
     */
    timeZone?: pulumi.Input<string>;
    v2betaId1: pulumi.Input<string>;
    v2betumId: pulumi.Input<string>;
}
