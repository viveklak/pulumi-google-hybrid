// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.WebSecurityScanner.V1Alpha
{
    public static class GetScanConfig
    {
        /// <summary>
        /// Gets a ScanConfig.
        /// </summary>
        public static Task<GetScanConfigResult> InvokeAsync(GetScanConfigArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetScanConfigResult>("google-native:websecurityscanner/v1alpha:getScanConfig", args ?? new GetScanConfigArgs(), options.WithVersion());
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
        /// Latest ScanRun if available.
        /// </summary>
        public readonly Outputs.ScanRunResponse LatestRun;
        /// <summary>
        /// The maximum QPS during scanning. A valid value ranges from 5 to 20 inclusively. If the field is unspecified or its value is set 0, server will default to 15. Other values outside of [5, 20] range will be rejected with INVALID_ARGUMENT error.
        /// </summary>
        public readonly int MaxQps;
        /// <summary>
        /// The resource name of the ScanConfig. The name follows the format of 'projects/{projectId}/scanConfigs/{scanConfigId}'. The ScanConfig IDs are generated by the system.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The schedule of the ScanConfig.
        /// </summary>
        public readonly Outputs.ScheduleResponse Schedule;
        /// <summary>
        /// The starting URLs from which the scanner finds site pages.
        /// </summary>
        public readonly ImmutableArray<string> StartingUrls;
        /// <summary>
        /// Set of Google Cloud platforms targeted by the scan. If empty, APP_ENGINE will be used as a default.
        /// </summary>
        public readonly ImmutableArray<string> TargetPlatforms;
        /// <summary>
        /// The user agent used during scanning.
        /// </summary>
        public readonly string UserAgent;

        [OutputConstructor]
        private GetScanConfigResult(
            Outputs.AuthenticationResponse authentication,

            ImmutableArray<string> blacklistPatterns,

            string displayName,

            Outputs.ScanRunResponse latestRun,

            int maxQps,

            string name,

            Outputs.ScheduleResponse schedule,

            ImmutableArray<string> startingUrls,

            ImmutableArray<string> targetPlatforms,

            string userAgent)
        {
            Authentication = authentication;
            BlacklistPatterns = blacklistPatterns;
            DisplayName = displayName;
            LatestRun = latestRun;
            MaxQps = maxQps;
            Name = name;
            Schedule = schedule;
            StartingUrls = startingUrls;
            TargetPlatforms = targetPlatforms;
            UserAgent = userAgent;
        }
    }
}
