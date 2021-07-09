// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets a ScanConfig.
func LookupScanConfig(ctx *pulumi.Context, args *LookupScanConfigArgs, opts ...pulumi.InvokeOption) (*LookupScanConfigResult, error) {
	var rv LookupScanConfigResult
	err := ctx.Invoke("google-native:websecurityscanner/v1beta:getScanConfig", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupScanConfigArgs struct {
	Project      string `pulumi:"project"`
	ScanConfigId string `pulumi:"scanConfigId"`
}

type LookupScanConfigResult struct {
	// The authentication configuration. If specified, service will use the authentication configuration during scanning.
	Authentication AuthenticationResponse `pulumi:"authentication"`
	// The excluded URL patterns as described in https://cloud.google.com/security-command-center/docs/how-to-use-web-security-scanner#excluding_urls
	BlacklistPatterns []string `pulumi:"blacklistPatterns"`
	// The user provided display name of the ScanConfig.
	DisplayName string `pulumi:"displayName"`
	// Controls export of scan configurations and results to Security Command Center.
	ExportToSecurityCommandCenter string `pulumi:"exportToSecurityCommandCenter"`
	// Whether to keep scanning even if most requests return HTTP error codes.
	IgnoreHttpStatusErrors bool `pulumi:"ignoreHttpStatusErrors"`
	// Latest ScanRun if available.
	LatestRun ScanRunResponse `pulumi:"latestRun"`
	// Whether the scan config is managed by Web Security Scanner, output only.
	ManagedScan bool `pulumi:"managedScan"`
	// The maximum QPS during scanning. A valid value ranges from 5 to 20 inclusively. If the field is unspecified or its value is set 0, server will default to 15. Other values outside of [5, 20] range will be rejected with INVALID_ARGUMENT error.
	MaxQps int `pulumi:"maxQps"`
	// The resource name of the ScanConfig. The name follows the format of 'projects/{projectId}/scanConfigs/{scanConfigId}'. The ScanConfig IDs are generated by the system.
	Name string `pulumi:"name"`
	// The risk level selected for the scan
	RiskLevel string `pulumi:"riskLevel"`
	// The schedule of the ScanConfig.
	Schedule ScheduleResponse `pulumi:"schedule"`
	// The starting URLs from which the scanner finds site pages.
	StartingUrls []string `pulumi:"startingUrls"`
	// Whether the scan configuration has enabled static IP address scan feature. If enabled, the scanner will access applications from static IP addresses.
	StaticIpScan bool `pulumi:"staticIpScan"`
	// Set of Google Cloud platforms targeted by the scan. If empty, APP_ENGINE will be used as a default.
	TargetPlatforms []string `pulumi:"targetPlatforms"`
	// The user agent used during scanning.
	UserAgent string `pulumi:"userAgent"`
}
