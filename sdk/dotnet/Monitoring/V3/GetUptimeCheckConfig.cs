// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Monitoring.V3
{
    public static class GetUptimeCheckConfig
    {
        /// <summary>
        /// Gets a single Uptime check configuration.
        /// </summary>
        public static Task<GetUptimeCheckConfigResult> InvokeAsync(GetUptimeCheckConfigArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetUptimeCheckConfigResult>("google-native:monitoring/v3:getUptimeCheckConfig", args ?? new GetUptimeCheckConfigArgs(), options.WithDefaults());

        /// <summary>
        /// Gets a single Uptime check configuration.
        /// </summary>
        public static Output<GetUptimeCheckConfigResult> Invoke(GetUptimeCheckConfigInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetUptimeCheckConfigResult>("google-native:monitoring/v3:getUptimeCheckConfig", args ?? new GetUptimeCheckConfigInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetUptimeCheckConfigArgs : Pulumi.InvokeArgs
    {
        [Input("project")]
        public string? Project { get; set; }

        [Input("uptimeCheckConfigId", required: true)]
        public string UptimeCheckConfigId { get; set; } = null!;

        public GetUptimeCheckConfigArgs()
        {
        }
    }

    public sealed class GetUptimeCheckConfigInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("uptimeCheckConfigId", required: true)]
        public Input<string> UptimeCheckConfigId { get; set; } = null!;

        public GetUptimeCheckConfigInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetUptimeCheckConfigResult
    {
        /// <summary>
        /// The type of checkers to use to execute the Uptime check.
        /// </summary>
        public readonly string CheckerType;
        /// <summary>
        /// The content that is expected to appear in the data returned by the target server against which the check is run. Currently, only the first entry in the content_matchers list is supported, and additional entries will be ignored. This field is optional and should only be specified if a content match is required as part of the/ Uptime check.
        /// </summary>
        public readonly ImmutableArray<Outputs.ContentMatcherResponse> ContentMatchers;
        /// <summary>
        /// A human-friendly name for the Uptime check configuration. The display name should be unique within a Stackdriver Workspace in order to make it easier to identify; however, uniqueness is not enforced. Required.
        /// </summary>
        public readonly string DisplayName;
        /// <summary>
        /// Contains information needed to make an HTTP or HTTPS check.
        /// </summary>
        public readonly Outputs.HttpCheckResponse HttpCheck;
        /// <summary>
        /// The internal checkers that this check will egress from. If is_internal is true and this list is empty, the check will egress from all the InternalCheckers configured for the project that owns this UptimeCheckConfig.
        /// </summary>
        public readonly ImmutableArray<Outputs.InternalCheckerResponse> InternalCheckers;
        /// <summary>
        /// If this is true, then checks are made only from the 'internal_checkers'. If it is false, then checks are made only from the 'selected_regions'. It is an error to provide 'selected_regions' when is_internal is true, or to provide 'internal_checkers' when is_internal is false.
        /// </summary>
        public readonly bool IsInternal;
        /// <summary>
        /// The monitored resource (https://cloud.google.com/monitoring/api/resources) associated with the configuration. The following monitored resource types are valid for this field: uptime_url, gce_instance, gae_app, aws_ec2_instance, aws_elb_load_balancer k8s_service servicedirectory_service
        /// </summary>
        public readonly Outputs.MonitoredResourceResponse MonitoredResource;
        /// <summary>
        /// A unique resource name for this Uptime check configuration. The format is: projects/[PROJECT_ID_OR_NUMBER]/uptimeCheckConfigs/[UPTIME_CHECK_ID] [PROJECT_ID_OR_NUMBER] is the Workspace host project associated with the Uptime check.This field should be omitted when creating the Uptime check configuration; on create, the resource name is assigned by the server and included in the response.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// How often, in seconds, the Uptime check is performed. Currently, the only supported values are 60s (1 minute), 300s (5 minutes), 600s (10 minutes), and 900s (15 minutes). Optional, defaults to 60s.
        /// </summary>
        public readonly string Period;
        /// <summary>
        /// The group resource associated with the configuration.
        /// </summary>
        public readonly Outputs.ResourceGroupResponse ResourceGroup;
        /// <summary>
        /// The list of regions from which the check will be run. Some regions contain one location, and others contain more than one. If this field is specified, enough regions must be provided to include a minimum of 3 locations. Not specifying this field will result in Uptime checks running from all available regions.
        /// </summary>
        public readonly ImmutableArray<string> SelectedRegions;
        /// <summary>
        /// Contains information needed to make a TCP check.
        /// </summary>
        public readonly Outputs.TcpCheckResponse TcpCheck;
        /// <summary>
        /// The maximum amount of time to wait for the request to complete (must be between 1 and 60 seconds). Required.
        /// </summary>
        public readonly string Timeout;

        [OutputConstructor]
        private GetUptimeCheckConfigResult(
            string checkerType,

            ImmutableArray<Outputs.ContentMatcherResponse> contentMatchers,

            string displayName,

            Outputs.HttpCheckResponse httpCheck,

            ImmutableArray<Outputs.InternalCheckerResponse> internalCheckers,

            bool isInternal,

            Outputs.MonitoredResourceResponse monitoredResource,

            string name,

            string period,

            Outputs.ResourceGroupResponse resourceGroup,

            ImmutableArray<string> selectedRegions,

            Outputs.TcpCheckResponse tcpCheck,

            string timeout)
        {
            CheckerType = checkerType;
            ContentMatchers = contentMatchers;
            DisplayName = displayName;
            HttpCheck = httpCheck;
            InternalCheckers = internalCheckers;
            IsInternal = isInternal;
            MonitoredResource = monitoredResource;
            Name = name;
            Period = period;
            ResourceGroup = resourceGroup;
            SelectedRegions = selectedRegions;
            TcpCheck = tcpCheck;
            Timeout = timeout;
        }
    }
}
