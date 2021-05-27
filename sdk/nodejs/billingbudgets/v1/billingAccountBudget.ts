// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates a new budget. See [Quotas and limits](https://cloud.google.com/billing/quotas) for more information on the limits of the number of budgets you can create.
 */
export class BillingAccountBudget extends pulumi.CustomResource {
    /**
     * Get an existing BillingAccountBudget resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): BillingAccountBudget {
        return new BillingAccountBudget(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:billingbudgets/v1:BillingAccountBudget';

    /**
     * Returns true if the given object is an instance of BillingAccountBudget.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is BillingAccountBudget {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === BillingAccountBudget.__pulumiType;
    }

    /**
     * Required. Budgeted amount.
     */
    public readonly amount!: pulumi.Output<outputs.billingbudgets.v1.GoogleCloudBillingBudgetsV1BudgetAmountResponse>;
    /**
     * Optional. Filters that define which resources are used to compute the actual spend against the budget amount, such as projects, services, and the budget's time period, as well as other filters.
     */
    public readonly budgetFilter!: pulumi.Output<outputs.billingbudgets.v1.GoogleCloudBillingBudgetsV1FilterResponse>;
    /**
     * User data for display name in UI. The name must be less than or equal to 60 characters.
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * Optional. Etag to validate that the object is unchanged for a read-modify-write operation. An empty etag will cause an update to overwrite other changes.
     */
    public readonly etag!: pulumi.Output<string>;
    /**
     * Resource name of the budget. The resource name implies the scope of a budget. Values are of the form `billingAccounts/{billingAccountId}/budgets/{budgetId}`.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Optional. Rules to apply to notifications sent based on budget spend and thresholds.
     */
    public readonly notificationsRule!: pulumi.Output<outputs.billingbudgets.v1.GoogleCloudBillingBudgetsV1NotificationsRuleResponse>;
    /**
     * Optional. Rules that trigger alerts (notifications of thresholds being crossed) when spend exceeds the specified percentages of the budget.
     */
    public readonly thresholdRules!: pulumi.Output<outputs.billingbudgets.v1.GoogleCloudBillingBudgetsV1ThresholdRuleResponse[]>;

    /**
     * Create a BillingAccountBudget resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: BillingAccountBudgetArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.billingAccountId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'billingAccountId'");
            }
            inputs["amount"] = args ? args.amount : undefined;
            inputs["billingAccountId"] = args ? args.billingAccountId : undefined;
            inputs["budgetFilter"] = args ? args.budgetFilter : undefined;
            inputs["displayName"] = args ? args.displayName : undefined;
            inputs["etag"] = args ? args.etag : undefined;
            inputs["notificationsRule"] = args ? args.notificationsRule : undefined;
            inputs["thresholdRules"] = args ? args.thresholdRules : undefined;
            inputs["name"] = undefined /*out*/;
        } else {
            inputs["amount"] = undefined /*out*/;
            inputs["budgetFilter"] = undefined /*out*/;
            inputs["displayName"] = undefined /*out*/;
            inputs["etag"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["notificationsRule"] = undefined /*out*/;
            inputs["thresholdRules"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(BillingAccountBudget.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a BillingAccountBudget resource.
 */
export interface BillingAccountBudgetArgs {
    /**
     * Required. Budgeted amount.
     */
    readonly amount?: pulumi.Input<inputs.billingbudgets.v1.GoogleCloudBillingBudgetsV1BudgetAmountArgs>;
    readonly billingAccountId: pulumi.Input<string>;
    /**
     * Optional. Filters that define which resources are used to compute the actual spend against the budget amount, such as projects, services, and the budget's time period, as well as other filters.
     */
    readonly budgetFilter?: pulumi.Input<inputs.billingbudgets.v1.GoogleCloudBillingBudgetsV1FilterArgs>;
    /**
     * User data for display name in UI. The name must be less than or equal to 60 characters.
     */
    readonly displayName?: pulumi.Input<string>;
    /**
     * Optional. Etag to validate that the object is unchanged for a read-modify-write operation. An empty etag will cause an update to overwrite other changes.
     */
    readonly etag?: pulumi.Input<string>;
    /**
     * Optional. Rules to apply to notifications sent based on budget spend and thresholds.
     */
    readonly notificationsRule?: pulumi.Input<inputs.billingbudgets.v1.GoogleCloudBillingBudgetsV1NotificationsRuleArgs>;
    /**
     * Optional. Rules that trigger alerts (notifications of thresholds being crossed) when spend exceeds the specified percentages of the budget.
     */
    readonly thresholdRules?: pulumi.Input<pulumi.Input<inputs.billingbudgets.v1.GoogleCloudBillingBudgetsV1ThresholdRuleArgs>[]>;
}
