// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the profile for an Apigee organization. See [Understanding organizations](https://cloud.google.com/apigee/docs/api-platform/fundamentals/organization-structure).
func LookupOrganization(ctx *pulumi.Context, args *LookupOrganizationArgs, opts ...pulumi.InvokeOption) (*LookupOrganizationResult, error) {
	var rv LookupOrganizationResult
	err := ctx.Invoke("google-native:apigee/v1:getOrganization", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupOrganizationArgs struct {
	OrganizationId string `pulumi:"organizationId"`
}

type LookupOrganizationResult struct {
	// Addon configurations of the Apigee organization.
	AddonsConfig GoogleCloudApigeeV1AddonsConfigResponse `pulumi:"addonsConfig"`
	// Primary GCP region for analytics data storage. For valid values, see [Create an Apigee organization](https://cloud.google.com/apigee/docs/api-platform/get-started/create-org).
	AnalyticsRegion string `pulumi:"analyticsRegion"`
	// Not used by Apigee.
	Attributes []string `pulumi:"attributes"`
	// Compute Engine network used for Service Networking to be peered with Apigee runtime instances. See [Getting started with the Service Networking API](https://cloud.google.com/service-infrastructure/docs/service-networking/getting-started). Valid only when [RuntimeType](#RuntimeType) is set to `CLOUD`. The value must be set before the creation of a runtime instance and can be updated only when there are no runtime instances. For example: `default`. Apigee also supports shared VPC (that is, the host network project is not the same as the one that is peering with Apigee). See [Shared VPC overview](https://cloud.google.com/vpc/docs/shared-vpc). To use a shared VPC network, use the following format: `projects/{host-project-id}/{region}/networks/{network-name}`. For example: `projects/my-sharedvpc-host/global/networks/mynetwork` **Note:** Not supported for Apigee hybrid.
	AuthorizedNetwork string `pulumi:"authorizedNetwork"`
	// Billing type of the Apigee organization. See [Apigee pricing](https://cloud.google.com/apigee/pricing).
	BillingType string `pulumi:"billingType"`
	// Base64-encoded public certificate for the root CA of the Apigee organization. Valid only when [RuntimeType](#RuntimeType) is `CLOUD`.
	CaCertificate string `pulumi:"caCertificate"`
	// Time that the Apigee organization was created in milliseconds since epoch.
	CreatedAt string `pulumi:"createdAt"`
	// Not used by Apigee.
	CustomerName string `pulumi:"customerName"`
	// Description of the Apigee organization.
	Description string `pulumi:"description"`
	DisplayName string `pulumi:"displayName"`
	// List of environments in the Apigee organization.
	Environments []string `pulumi:"environments"`
	// Time that the Apigee organization is scheduled for deletion.
	ExpiresAt string `pulumi:"expiresAt"`
	// Time that the Apigee organization was last modified in milliseconds since epoch.
	LastModifiedAt string `pulumi:"lastModifiedAt"`
	// Name of the Apigee organization.
	Name string `pulumi:"name"`
	// Project ID associated with the Apigee organization.
	Project string `pulumi:"project"`
	// Properties defined in the Apigee organization profile.
	Properties GoogleCloudApigeeV1PropertiesResponse `pulumi:"properties"`
	// Cloud KMS key name used for encrypting the data that is stored and replicated across runtime instances. Update is not allowed after the organization is created. Required when [RuntimeType](#RuntimeType) is `CLOUD`. If not specified when [RuntimeType](#RuntimeType) is `TRIAL`, a Google-Managed encryption key will be used. For example: "projects/foo/locations/us/keyRings/bar/cryptoKeys/baz". **Note:** Not supported for Apigee hybrid.
	RuntimeDatabaseEncryptionKeyName string `pulumi:"runtimeDatabaseEncryptionKeyName"`
	// Runtime type of the Apigee organization based on the Apigee subscription purchased.
	RuntimeType string `pulumi:"runtimeType"`
	// State of the organization. Values other than ACTIVE means the resource is not ready to use.
	State string `pulumi:"state"`
	// Not used by Apigee.
	Type string `pulumi:"type"`
}
