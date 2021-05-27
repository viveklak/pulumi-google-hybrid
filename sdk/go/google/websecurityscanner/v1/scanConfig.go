// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a new ScanConfig.
type ScanConfig struct {
	pulumi.CustomResourceState

	// The authentication configuration. If specified, service will use the authentication configuration during scanning.
	Authentication AuthenticationResponseOutput `pulumi:"authentication"`
	// The excluded URL patterns as described in https://cloud.google.com/security-command-center/docs/how-to-use-web-security-scanner#excluding_urls
	BlacklistPatterns pulumi.StringArrayOutput `pulumi:"blacklistPatterns"`
	// Required. The user provided display name of the ScanConfig.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// Controls export of scan configurations and results to Security Command Center.
	ExportToSecurityCommandCenter pulumi.StringOutput `pulumi:"exportToSecurityCommandCenter"`
	// Whether to keep scanning even if most requests return HTTP error codes.
	IgnoreHttpStatusErrors pulumi.BoolOutput `pulumi:"ignoreHttpStatusErrors"`
	// Whether the scan config is managed by Web Security Scanner, output only.
	ManagedScan pulumi.BoolOutput `pulumi:"managedScan"`
	// The maximum QPS during scanning. A valid value ranges from 5 to 20 inclusively. If the field is unspecified or its value is set 0, server will default to 15. Other values outside of [5, 20] range will be rejected with INVALID_ARGUMENT error.
	MaxQps pulumi.IntOutput `pulumi:"maxQps"`
	// The resource name of the ScanConfig. The name follows the format of 'projects/{projectId}/scanConfigs/{scanConfigId}'. The ScanConfig IDs are generated by the system.
	Name pulumi.StringOutput `pulumi:"name"`
	// The risk level selected for the scan
	RiskLevel pulumi.StringOutput `pulumi:"riskLevel"`
	// The schedule of the ScanConfig.
	Schedule ScheduleResponseOutput `pulumi:"schedule"`
	// Required. The starting URLs from which the scanner finds site pages.
	StartingUrls pulumi.StringArrayOutput `pulumi:"startingUrls"`
	// Whether the scan configuration has enabled static IP address scan feature. If enabled, the scanner will access applications from static IP addresses.
	StaticIpScan pulumi.BoolOutput `pulumi:"staticIpScan"`
	// The user agent used during scanning.
	UserAgent pulumi.StringOutput `pulumi:"userAgent"`
}

// NewScanConfig registers a new resource with the given unique name, arguments, and options.
func NewScanConfig(ctx *pulumi.Context,
	name string, args *ScanConfigArgs, opts ...pulumi.ResourceOption) (*ScanConfig, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	var resource ScanConfig
	err := ctx.RegisterResource("google-native:websecurityscanner/v1:ScanConfig", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetScanConfig gets an existing ScanConfig resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetScanConfig(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ScanConfigState, opts ...pulumi.ResourceOption) (*ScanConfig, error) {
	var resource ScanConfig
	err := ctx.ReadResource("google-native:websecurityscanner/v1:ScanConfig", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ScanConfig resources.
type scanConfigState struct {
	// The authentication configuration. If specified, service will use the authentication configuration during scanning.
	Authentication *AuthenticationResponse `pulumi:"authentication"`
	// The excluded URL patterns as described in https://cloud.google.com/security-command-center/docs/how-to-use-web-security-scanner#excluding_urls
	BlacklistPatterns []string `pulumi:"blacklistPatterns"`
	// Required. The user provided display name of the ScanConfig.
	DisplayName *string `pulumi:"displayName"`
	// Controls export of scan configurations and results to Security Command Center.
	ExportToSecurityCommandCenter *string `pulumi:"exportToSecurityCommandCenter"`
	// Whether to keep scanning even if most requests return HTTP error codes.
	IgnoreHttpStatusErrors *bool `pulumi:"ignoreHttpStatusErrors"`
	// Whether the scan config is managed by Web Security Scanner, output only.
	ManagedScan *bool `pulumi:"managedScan"`
	// The maximum QPS during scanning. A valid value ranges from 5 to 20 inclusively. If the field is unspecified or its value is set 0, server will default to 15. Other values outside of [5, 20] range will be rejected with INVALID_ARGUMENT error.
	MaxQps *int `pulumi:"maxQps"`
	// The resource name of the ScanConfig. The name follows the format of 'projects/{projectId}/scanConfigs/{scanConfigId}'. The ScanConfig IDs are generated by the system.
	Name *string `pulumi:"name"`
	// The risk level selected for the scan
	RiskLevel *string `pulumi:"riskLevel"`
	// The schedule of the ScanConfig.
	Schedule *ScheduleResponse `pulumi:"schedule"`
	// Required. The starting URLs from which the scanner finds site pages.
	StartingUrls []string `pulumi:"startingUrls"`
	// Whether the scan configuration has enabled static IP address scan feature. If enabled, the scanner will access applications from static IP addresses.
	StaticIpScan *bool `pulumi:"staticIpScan"`
	// The user agent used during scanning.
	UserAgent *string `pulumi:"userAgent"`
}

type ScanConfigState struct {
	// The authentication configuration. If specified, service will use the authentication configuration during scanning.
	Authentication AuthenticationResponsePtrInput
	// The excluded URL patterns as described in https://cloud.google.com/security-command-center/docs/how-to-use-web-security-scanner#excluding_urls
	BlacklistPatterns pulumi.StringArrayInput
	// Required. The user provided display name of the ScanConfig.
	DisplayName pulumi.StringPtrInput
	// Controls export of scan configurations and results to Security Command Center.
	ExportToSecurityCommandCenter pulumi.StringPtrInput
	// Whether to keep scanning even if most requests return HTTP error codes.
	IgnoreHttpStatusErrors pulumi.BoolPtrInput
	// Whether the scan config is managed by Web Security Scanner, output only.
	ManagedScan pulumi.BoolPtrInput
	// The maximum QPS during scanning. A valid value ranges from 5 to 20 inclusively. If the field is unspecified or its value is set 0, server will default to 15. Other values outside of [5, 20] range will be rejected with INVALID_ARGUMENT error.
	MaxQps pulumi.IntPtrInput
	// The resource name of the ScanConfig. The name follows the format of 'projects/{projectId}/scanConfigs/{scanConfigId}'. The ScanConfig IDs are generated by the system.
	Name pulumi.StringPtrInput
	// The risk level selected for the scan
	RiskLevel pulumi.StringPtrInput
	// The schedule of the ScanConfig.
	Schedule ScheduleResponsePtrInput
	// Required. The starting URLs from which the scanner finds site pages.
	StartingUrls pulumi.StringArrayInput
	// Whether the scan configuration has enabled static IP address scan feature. If enabled, the scanner will access applications from static IP addresses.
	StaticIpScan pulumi.BoolPtrInput
	// The user agent used during scanning.
	UserAgent pulumi.StringPtrInput
}

func (ScanConfigState) ElementType() reflect.Type {
	return reflect.TypeOf((*scanConfigState)(nil)).Elem()
}

type scanConfigArgs struct {
	// The authentication configuration. If specified, service will use the authentication configuration during scanning.
	Authentication *Authentication `pulumi:"authentication"`
	// The excluded URL patterns as described in https://cloud.google.com/security-command-center/docs/how-to-use-web-security-scanner#excluding_urls
	BlacklistPatterns []string `pulumi:"blacklistPatterns"`
	// Required. The user provided display name of the ScanConfig.
	DisplayName *string `pulumi:"displayName"`
	// Controls export of scan configurations and results to Security Command Center.
	ExportToSecurityCommandCenter *string `pulumi:"exportToSecurityCommandCenter"`
	// Whether to keep scanning even if most requests return HTTP error codes.
	IgnoreHttpStatusErrors *bool `pulumi:"ignoreHttpStatusErrors"`
	// Whether the scan config is managed by Web Security Scanner, output only.
	ManagedScan *bool `pulumi:"managedScan"`
	// The maximum QPS during scanning. A valid value ranges from 5 to 20 inclusively. If the field is unspecified or its value is set 0, server will default to 15. Other values outside of [5, 20] range will be rejected with INVALID_ARGUMENT error.
	MaxQps *int `pulumi:"maxQps"`
	// The resource name of the ScanConfig. The name follows the format of 'projects/{projectId}/scanConfigs/{scanConfigId}'. The ScanConfig IDs are generated by the system.
	Name    *string `pulumi:"name"`
	Project string  `pulumi:"project"`
	// The risk level selected for the scan
	RiskLevel *string `pulumi:"riskLevel"`
	// The schedule of the ScanConfig.
	Schedule *Schedule `pulumi:"schedule"`
	// Required. The starting URLs from which the scanner finds site pages.
	StartingUrls []string `pulumi:"startingUrls"`
	// Whether the scan configuration has enabled static IP address scan feature. If enabled, the scanner will access applications from static IP addresses.
	StaticIpScan *bool `pulumi:"staticIpScan"`
	// The user agent used during scanning.
	UserAgent *string `pulumi:"userAgent"`
}

// The set of arguments for constructing a ScanConfig resource.
type ScanConfigArgs struct {
	// The authentication configuration. If specified, service will use the authentication configuration during scanning.
	Authentication AuthenticationPtrInput
	// The excluded URL patterns as described in https://cloud.google.com/security-command-center/docs/how-to-use-web-security-scanner#excluding_urls
	BlacklistPatterns pulumi.StringArrayInput
	// Required. The user provided display name of the ScanConfig.
	DisplayName pulumi.StringPtrInput
	// Controls export of scan configurations and results to Security Command Center.
	ExportToSecurityCommandCenter pulumi.StringPtrInput
	// Whether to keep scanning even if most requests return HTTP error codes.
	IgnoreHttpStatusErrors pulumi.BoolPtrInput
	// Whether the scan config is managed by Web Security Scanner, output only.
	ManagedScan pulumi.BoolPtrInput
	// The maximum QPS during scanning. A valid value ranges from 5 to 20 inclusively. If the field is unspecified or its value is set 0, server will default to 15. Other values outside of [5, 20] range will be rejected with INVALID_ARGUMENT error.
	MaxQps pulumi.IntPtrInput
	// The resource name of the ScanConfig. The name follows the format of 'projects/{projectId}/scanConfigs/{scanConfigId}'. The ScanConfig IDs are generated by the system.
	Name    pulumi.StringPtrInput
	Project pulumi.StringInput
	// The risk level selected for the scan
	RiskLevel pulumi.StringPtrInput
	// The schedule of the ScanConfig.
	Schedule SchedulePtrInput
	// Required. The starting URLs from which the scanner finds site pages.
	StartingUrls pulumi.StringArrayInput
	// Whether the scan configuration has enabled static IP address scan feature. If enabled, the scanner will access applications from static IP addresses.
	StaticIpScan pulumi.BoolPtrInput
	// The user agent used during scanning.
	UserAgent pulumi.StringPtrInput
}

func (ScanConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*scanConfigArgs)(nil)).Elem()
}

type ScanConfigInput interface {
	pulumi.Input

	ToScanConfigOutput() ScanConfigOutput
	ToScanConfigOutputWithContext(ctx context.Context) ScanConfigOutput
}

func (*ScanConfig) ElementType() reflect.Type {
	return reflect.TypeOf((*ScanConfig)(nil))
}

func (i *ScanConfig) ToScanConfigOutput() ScanConfigOutput {
	return i.ToScanConfigOutputWithContext(context.Background())
}

func (i *ScanConfig) ToScanConfigOutputWithContext(ctx context.Context) ScanConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScanConfigOutput)
}

type ScanConfigOutput struct {
	*pulumi.OutputState
}

func (ScanConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScanConfig)(nil))
}

func (o ScanConfigOutput) ToScanConfigOutput() ScanConfigOutput {
	return o
}

func (o ScanConfigOutput) ToScanConfigOutputWithContext(ctx context.Context) ScanConfigOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ScanConfigOutput{})
}
