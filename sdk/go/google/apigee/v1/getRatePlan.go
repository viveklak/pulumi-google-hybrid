// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the details of a rate plan.
func LookupRatePlan(ctx *pulumi.Context, args *LookupRatePlanArgs, opts ...pulumi.InvokeOption) (*LookupRatePlanResult, error) {
	var rv LookupRatePlanResult
	err := ctx.Invoke("google-native:apigee/v1:getRatePlan", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupRatePlanArgs struct {
	ApiproductId   string `pulumi:"apiproductId"`
	OrganizationId string `pulumi:"organizationId"`
	RateplanId     string `pulumi:"rateplanId"`
}

type LookupRatePlanResult struct {
	// Name of the API product that the rate plan is associated with.
	Apiproduct string `pulumi:"apiproduct"`
	// Frequency at which the customer will be billed.
	BillingPeriod string `pulumi:"billingPeriod"`
	// API call volume ranges and the fees charged when the total number of API calls is within a given range. The method used to calculate the final fee depends on the selected pricing model. For example, if the pricing model is `STAIRSTEP` and the ranges are defined as follows: ```{ "start": 1, "end": 100, "fee": 75 }, { "start": 101, "end": 200, "fee": 100 }, }``` Then the following fees would be charged based on the total number of API calls (assuming the currency selected is `USD`): * 1 call costs $75 * 50 calls cost $75 * 150 calls cost $100 The number of API calls cannot exceed 200.
	ConsumptionPricingRates []GoogleCloudApigeeV1RateRangeResponse `pulumi:"consumptionPricingRates"`
	// Pricing model used for consumption-based charges.
	ConsumptionPricingType string `pulumi:"consumptionPricingType"`
	// Time that the rate plan was created in milliseconds since epoch.
	CreatedAt string `pulumi:"createdAt"`
	// Currency to be used for billing. Consists of a three-letter code as defined by the [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) standard.
	CurrencyCode string `pulumi:"currencyCode"`
	// Description of the rate plan.
	Description string `pulumi:"description"`
	// Display name of the rate plan.
	DisplayName string `pulumi:"displayName"`
	// Time when the rate plan will expire in milliseconds since epoch. Set to 0 or `null` to indicate that the rate plan should never expire.
	EndTime string `pulumi:"endTime"`
	// Frequency at which the fixed fee is charged.
	FixedFeeFrequency int `pulumi:"fixedFeeFrequency"`
	// Fixed amount that is charged at a defined interval and billed in advance of use of the API product. The fee will be prorated for the first billing period.
	FixedRecurringFee GoogleTypeMoneyResponse `pulumi:"fixedRecurringFee"`
	// Time the rate plan was last modified in milliseconds since epoch.
	LastModifiedAt string `pulumi:"lastModifiedAt"`
	// Name of the rate plan.
	Name string `pulumi:"name"`
	// Flag that specifies the billing account type, prepaid or postpaid.
	PaymentFundingModel string `pulumi:"paymentFundingModel"`
	// Details of the revenue sharing model.
	RevenueShareRates []GoogleCloudApigeeV1RevenueShareRangeResponse `pulumi:"revenueShareRates"`
	// Method used to calculate the revenue that is shared with developers.
	RevenueShareType string `pulumi:"revenueShareType"`
	// Initial, one-time fee paid when purchasing the API product.
	SetupFee GoogleTypeMoneyResponse `pulumi:"setupFee"`
	// Time when the rate plan becomes active in milliseconds since epoch.
	StartTime string `pulumi:"startTime"`
	// Current state of the rate plan (draft or published).
	State string `pulumi:"state"`
}