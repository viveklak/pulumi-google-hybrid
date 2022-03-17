// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***


export const CapacityCommitmentPlan = {
    /**
     * Invalid plan value. Requests with this value will be rejected with error code `google.rpc.Code.INVALID_ARGUMENT`.
     */
    CommitmentPlanUnspecified: "COMMITMENT_PLAN_UNSPECIFIED",
    /**
     * Flex commitments have committed period of 1 minute after becoming ACTIVE. After that, they are not in a committed period anymore and can be removed any time.
     */
    Flex: "FLEX",
    /**
     * Trial commitments have a committed period of 182 days after becoming ACTIVE. After that, they are converted to a new commitment based on the `renewal_plan`. Default `renewal_plan` for Trial commitment is Flex so that it can be deleted right after committed period ends.
     */
    Trial: "TRIAL",
    /**
     * Monthly commitments have a committed period of 30 days after becoming ACTIVE. After that, they are not in a committed period anymore and can be removed any time.
     */
    Monthly: "MONTHLY",
    /**
     * Annual commitments have a committed period of 365 days after becoming ACTIVE. After that they are converted to a new commitment based on the renewal_plan.
     */
    Annual: "ANNUAL",
} as const;

/**
 * Capacity commitment commitment plan.
 */
export type CapacityCommitmentPlan = (typeof CapacityCommitmentPlan)[keyof typeof CapacityCommitmentPlan];

export const CapacityCommitmentRenewalPlan = {
    /**
     * Invalid plan value. Requests with this value will be rejected with error code `google.rpc.Code.INVALID_ARGUMENT`.
     */
    CommitmentPlanUnspecified: "COMMITMENT_PLAN_UNSPECIFIED",
    /**
     * Flex commitments have committed period of 1 minute after becoming ACTIVE. After that, they are not in a committed period anymore and can be removed any time.
     */
    Flex: "FLEX",
    /**
     * Trial commitments have a committed period of 182 days after becoming ACTIVE. After that, they are converted to a new commitment based on the `renewal_plan`. Default `renewal_plan` for Trial commitment is Flex so that it can be deleted right after committed period ends.
     */
    Trial: "TRIAL",
    /**
     * Monthly commitments have a committed period of 30 days after becoming ACTIVE. After that, they are not in a committed period anymore and can be removed any time.
     */
    Monthly: "MONTHLY",
    /**
     * Annual commitments have a committed period of 365 days after becoming ACTIVE. After that they are converted to a new commitment based on the renewal_plan.
     */
    Annual: "ANNUAL",
} as const;

/**
 * The plan this capacity commitment is converted to after commitment_end_time passes. Once the plan is changed, committed period is extended according to commitment plan. Only applicable for ANNUAL and TRIAL commitments.
 */
export type CapacityCommitmentRenewalPlan = (typeof CapacityCommitmentRenewalPlan)[keyof typeof CapacityCommitmentRenewalPlan];
