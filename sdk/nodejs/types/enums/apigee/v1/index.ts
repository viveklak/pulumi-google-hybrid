// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***


export const DataCollectorType = {
    /**
     * For future compatibility.
     */
    TypeUnspecified: "TYPE_UNSPECIFIED",
    /**
     * For integer values.
     */
    Integer: "INTEGER",
    /**
     * For float values.
     */
    Float: "FLOAT",
    /**
     * For string values.
     */
    String: "STRING",
    /**
     * For boolean values.
     */
    Boolean: "BOOLEAN",
    /**
     * For datetime values.
     */
    Datetime: "DATETIME",
} as const;

/**
 * Immutable. The type of data this data collector will collect.
 */
export type DataCollectorType = (typeof DataCollectorType)[keyof typeof DataCollectorType];

export const GoogleCloudApigeeV1TraceSamplingConfigSampler = {
    /**
     * Sampler unspecified.
     */
    SamplerUnspecified: "SAMPLER_UNSPECIFIED",
    /**
     * OFF means distributed trace is disabled, or the sampling probability is 0.
     */
    Off: "OFF",
    /**
     * PROBABILITY means traces are captured on a probability that defined by sampling_rate. The sampling rate is limited to 0 to 0.5 when this is set.
     */
    Probability: "PROBABILITY",
} as const;

/**
 * Sampler of distributed tracing. OFF is the default value.
 */
export type GoogleCloudApigeeV1TraceSamplingConfigSampler = (typeof GoogleCloudApigeeV1TraceSamplingConfigSampler)[keyof typeof GoogleCloudApigeeV1TraceSamplingConfigSampler];

export const GoogleIamV1AuditLogConfigLogType = {
    /**
     * Default case. Should never be this.
     */
    LogTypeUnspecified: "LOG_TYPE_UNSPECIFIED",
    /**
     * Admin reads. Example: CloudIAM getIamPolicy
     */
    AdminRead: "ADMIN_READ",
    /**
     * Data writes. Example: CloudSQL Users create
     */
    DataWrite: "DATA_WRITE",
    /**
     * Data reads. Example: CloudSQL Users list
     */
    DataRead: "DATA_READ",
} as const;

/**
 * The log type that this config enables.
 */
export type GoogleIamV1AuditLogConfigLogType = (typeof GoogleIamV1AuditLogConfigLogType)[keyof typeof GoogleIamV1AuditLogConfigLogType];

export const InstancePeeringCidrRange = {
    /**
     * Range not specified.
     */
    CidrRangeUnspecified: "CIDR_RANGE_UNSPECIFIED",
    /**
     * `/16` CIDR range.
     */
    Slash16: "SLASH_16",
    /**
     * `/17` CIDR range.
     */
    Slash17: "SLASH_17",
    /**
     * `/18` CIDR range.
     */
    Slash18: "SLASH_18",
    /**
     * `/19` CIDR range.
     */
    Slash19: "SLASH_19",
    /**
     * `/20` CIDR range.
     */
    Slash20: "SLASH_20",
    /**
     * `/23` CIDR range. Supported for evaluation only.
     */
    Slash23: "SLASH_23",
} as const;

/**
 * Optional. Size of the CIDR block range that will be reserved by the instance. PAID organizations support `SLASH_16` to `SLASH_20` and defaults to `SLASH_16`. Evaluation organizations support only `SLASH_23`.
 */
export type InstancePeeringCidrRange = (typeof InstancePeeringCidrRange)[keyof typeof InstancePeeringCidrRange];

export const OrganizationBillingType = {
    /**
     * Billing type not specified.
     */
    BillingTypeUnspecified: "BILLING_TYPE_UNSPECIFIED",
    /**
     * A pre-paid subscription to Apigee.
     */
    Subscription: "SUBSCRIPTION",
    /**
     * Free and limited access to Apigee for evaluation purposes only. only.
     */
    Evaluation: "EVALUATION",
} as const;

/**
 * Billing type of the Apigee organization. See [Apigee pricing](https://cloud.google.com/apigee/pricing).
 */
export type OrganizationBillingType = (typeof OrganizationBillingType)[keyof typeof OrganizationBillingType];

export const OrganizationRuntimeType = {
    /**
     * Runtime type not specified.
     */
    RuntimeTypeUnspecified: "RUNTIME_TYPE_UNSPECIFIED",
    /**
     * Google-managed Apigee runtime.
     */
    Cloud: "CLOUD",
    /**
     * User-managed Apigee hybrid runtime.
     */
    Hybrid: "HYBRID",
} as const;

/**
 * Required. Runtime type of the Apigee organization based on the Apigee subscription purchased.
 */
export type OrganizationRuntimeType = (typeof OrganizationRuntimeType)[keyof typeof OrganizationRuntimeType];

export const OrganizationType = {
    /**
     * Subscription type not specified.
     */
    TypeUnspecified: "TYPE_UNSPECIFIED",
    /**
     * Subscription to Apigee is free, limited, and used for evaluation purposes only.
     */
    TypeTrial: "TYPE_TRIAL",
    /**
     * Full subscription to Apigee has been purchased. See [Apigee pricing](https://cloud.google.com/apigee/pricing/).
     */
    TypePaid: "TYPE_PAID",
    /**
     * For internal users only.
     */
    TypeInternal: "TYPE_INTERNAL",
} as const;

/**
 * Not used by Apigee.
 */
export type OrganizationType = (typeof OrganizationType)[keyof typeof OrganizationType];

export const RatePlanBillingPeriod = {
    /**
     * Billing period not specified.
     */
    BillingPeriodUnspecified: "BILLING_PERIOD_UNSPECIFIED",
    /**
     * Weekly billing period. **Note**: Not supported by Apigee at this time.
     */
    Weekly: "WEEKLY",
    /**
     * Monthly billing period.
     */
    Monthly: "MONTHLY",
} as const;

/**
 * Frequency at which the customer will be billed.
 */
export type RatePlanBillingPeriod = (typeof RatePlanBillingPeriod)[keyof typeof RatePlanBillingPeriod];

export const RatePlanConsumptionPricingType = {
    /**
     * Pricing model not specified. This is the default.
     */
    ConsumptionPricingTypeUnspecified: "CONSUMPTION_PRICING_TYPE_UNSPECIFIED",
    /**
     * Fixed rate charged for each API call.
     */
    FixedPerUnit: "FIXED_PER_UNIT",
    /**
     * Variable rate charged based on the total volume of API calls. Example: * 1-100 calls cost $2 per call * 101-200 calls cost $1.50 per call * 201-300 calls cost $1 per call * Total price for 50 calls: 50 x $2 = $100 * Total price for 150 calls: 150 x $1.5 = $225 * Total price for 250 calls: 250 x $1 = $250. **Note**: Not supported by Apigee at this time.
     */
    Banded: "BANDED",
    /**
     * Variable rate charged for each API call based on price tiers. Example: * 1-100 calls cost $2 per call * 101-200 calls cost $1.50 per call * 201-300 calls cost $1 per call * Total price for 50 calls: 50 x $2 = $100 * Total price for 150 calls: 100 x $2 + 50 x $1.5 = $275 * Total price for 250 calls: 100 x $2 + 100 x $1.5 + 50 x $1 = $400. **Note**: Not supported by Apigee at this time.
     */
    Tiered: "TIERED",
    /**
     * Flat rate charged for a bundle of API calls whether or not the entire bundle is used. Example: * 1-100 calls cost $75 flat fee * 101-200 calls cost $100 flat free * 201-300 calls cost $150 flat fee * Total price for 1 call: $75 * Total price for 50 calls: $75 * Total price for 150 calls: $100 * Total price for 250 calls: $150. **Note**: Not supported by Apigee at this time.
     */
    Stairstep: "STAIRSTEP",
} as const;

/**
 * Pricing model used for consumption-based charges.
 */
export type RatePlanConsumptionPricingType = (typeof RatePlanConsumptionPricingType)[keyof typeof RatePlanConsumptionPricingType];

export const RatePlanPaymentFundingModel = {
    /**
     * Billing account type not specified.
     */
    PaymentFundingModelUnspecified: "PAYMENT_FUNDING_MODEL_UNSPECIFIED",
    /**
     * Prepaid billing account type. Developer pays in advance for the use of your API products. Funds are deducted from their prepaid account balance. **Note**: Not supported by Apigee at this time.
     */
    Prepaid: "PREPAID",
    /**
     * Postpaid billing account type. Developer is billed through an invoice after using your API products.
     */
    Postpaid: "POSTPAID",
} as const;

/**
 * Flag that specifies the billing account type, prepaid or postpaid.
 */
export type RatePlanPaymentFundingModel = (typeof RatePlanPaymentFundingModel)[keyof typeof RatePlanPaymentFundingModel];

export const RatePlanRevenueShareType = {
    /**
     * Revenue share type is not specified.
     */
    RevenueShareTypeUnspecified: "REVENUE_SHARE_TYPE_UNSPECIFIED",
    /**
     * Fixed percentage of the total revenue will be shared. The percentage to be shared can be configured by the API provider.
     */
    Fixed: "FIXED",
    /**
     * Amount of revenue shared depends on the number of API calls. The API call volume ranges and the revenue share percentage for each volume can be configured by the API provider. **Note**: Not supported by Apigee at this time.
     */
    VolumeBanded: "VOLUME_BANDED",
} as const;

/**
 * Method used to calculate the revenue that is shared with developers.
 */
export type RatePlanRevenueShareType = (typeof RatePlanRevenueShareType)[keyof typeof RatePlanRevenueShareType];

export const RatePlanState = {
    /**
     * State of the rate plan is not specified.
     */
    StateUnspecified: "STATE_UNSPECIFIED",
    /**
     * Rate plan is in draft mode and only visible to API providers.
     */
    Draft: "DRAFT",
    /**
     * Rate plan is published and will become visible to developers for the configured duration (between `startTime` and `endTime`).
     */
    Published: "PUBLISHED",
} as const;

/**
 * Current state of the rate plan (draft or published).
 */
export type RatePlanState = (typeof RatePlanState)[keyof typeof RatePlanState];