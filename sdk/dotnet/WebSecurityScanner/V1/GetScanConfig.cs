// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.WebSecurityScanner.V1
{
    public static class GetScanConfig
    {
        /// <summary>
        /// Gets a ScanConfig.
        /// </summary>
        public static Task<GetScanConfigResult> InvokeAsync(GetScanConfigArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetScanConfigResult>("google-native:websecurityscanner/v1:getScanConfig", args ?? new GetScanConfigArgs(), options.WithVersion());
    }


    public sealed class GetScanConfigArgs : Pulumi.InvokeArgs
    {
        [Input("project", required: true)]
        public string Project { get; set; } = null!;

        [Input("scanConfigId", required: true)]
        public string ScanConfigId { get; set; } = null!;

        public GetScanConfigArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetScanConfigResult
    {
        /// <summary>
        /// The authentication configuration. If specified, service will use the authentication configuration during scanning.
        /// </summary>
        public readonly Outputs.AuthenticationResponse Authentication;
        /// <summary>
        /// The excluded URL patterns as described in https://cloud.google.com/security-command-center/docs/how-to-use-web-security-scanner#excluding_urls
        /// </summary>
        public readonly ImmutableArray<string> BlacklistPatterns;
        /// <summary>
        /// The user provided display name of the ScanConfig.
        /// </summary>
        public readonly string DisplayName;
        /// <summary>
        /// Controls export of scan configurations and results to Security Command Center.
        /// </summary>
        public readonly string ExportToSecurityCommandCenter;
        /// <summary>
        /// Whether to keep scanning even if most requests return HTTP error codes.
        /// </summary>
        public readonly bool IgnoreHttpStatusErrors;
        /// <summary>
        /// Whether the scan config is managed by Web Security Scanner, output only.
        /// </summary>
        public readonly bool ManagedScan;
        /// <summary>
        /// The maximum QPS during scanning. A valid value ranges from 5 to 20 inclusively. If the field is unspecified or its value is set 0, server will default to 15. Other values outside of [5, 20] range will be rejected with INVALID_ARGUMENT error.
        /// </summary>
        public readonly int MaxQps;
        /// <summary>
        /// The resource name of the ScanConfig. The name follows the format of 'projects/{projectId}/scanConfigs/{scanConfigId}'. The ScanConfig IDs are generated by the system.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The risk level selected for the scan
        /// </summary>
        public readonly string RiskLevel;
        /// <summary>
        /// The schedule of the ScanConfig.
        /// </summary>
        public readonly Outputs.ScheduleResponse Schedule;
        /// <summary>
        /// The starting URLs from which the scanner finds site pages.
        /// </summary>
        public readonly ImmutableArray<string> StartingUrls;
        /// <summary>
        /// Whether the scan configuration has enabled static IP address scan feature. If enabled, the scanner will access applications from static IP addresses.
        /// </summary>
        public readonly bool StaticIpScan;
        /// <summary>
        /// The user agent used during scanning.
        /// </summary>
        public readonly string UserAgent;

        [OutputConstructor]
        private GetScanConfigResult(
            Outputs.AuthenticationResponse authentication,

            ImmutableArray<string> blacklistPatterns,

            string displayName,

            string exportToSecurityCommandCenter,

            bool ignoreHttpStatusErrors,

            bool managedScan,

            int maxQps,

            string name,

            string riskLevel,

            Outputs.ScheduleResponse schedule,

            ImmutableArray<string> startingUrls,

            bool staticIpScan,

            string userAgent)
        {
            Authentication = authentication;
            BlacklistPatterns = blacklistPatterns;
            DisplayName = displayName;
            ExportToSecurityCommandCenter = exportToSecurityCommandCenter;
            IgnoreHttpStatusErrors = ignoreHttpStatusErrors;
            ManagedScan = managedScan;
            MaxQps = maxQps;
            Name = name;
            RiskLevel = riskLevel;
            Schedule = schedule;
            StartingUrls = startingUrls;
            StaticIpScan = staticIpScan;
            UserAgent = userAgent;
        }
    }
}
