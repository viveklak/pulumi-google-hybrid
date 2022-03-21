// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Composer.V1Beta1.Inputs
{

    /// <summary>
    /// Configuration for controlling how IPs are allocated in the GKE cluster.
    /// </summary>
    public sealed class IPAllocationPolicyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Optional. The IP address range used to allocate IP addresses to pods in the cluster. For Cloud Composer environments in versions composer-1.*.*-airflow-*.*.*, this field is applicable only when `use_ip_aliases` is true. Set to blank to have GKE choose a range with the default size. Set to /netmask (e.g. `/14`) to have GKE choose a range with a specific netmask. Set to a [CIDR](https://en.wikipedia.org/wiki/Classless_Inter-Domain_Routing) notation (e.g. `10.96.0.0/14`) from the RFC-1918 private networks (e.g. `10.0.0.0/8`, `172.16.0.0/12`, `192.168.0.0/16`) to pick a specific range to use. Specify `cluster_secondary_range_name` or `cluster_ipv4_cidr_block` but not both.
        /// </summary>
        [Input("clusterIpv4CidrBlock")]
        public Input<string>? ClusterIpv4CidrBlock { get; set; }

        /// <summary>
        /// Optional. The name of the cluster's secondary range used to allocate IP addresses to pods. Specify either `cluster_secondary_range_name` or `cluster_ipv4_cidr_block` but not both. For Cloud Composer environments in versions composer-1.*.*-airflow-*.*.*, this field is applicable only when `use_ip_aliases` is true.
        /// </summary>
        [Input("clusterSecondaryRangeName")]
        public Input<string>? ClusterSecondaryRangeName { get; set; }

        /// <summary>
        /// Optional. The IP address range of the services IP addresses in this cluster. For Cloud Composer environments in versions composer-1.*.*-airflow-*.*.*, this field is applicable only when `use_ip_aliases` is true. Set to blank to have GKE choose a range with the default size. Set to /netmask (e.g. `/14`) to have GKE choose a range with a specific netmask. Set to a [CIDR](https://en.wikipedia.org/wiki/Classless_Inter-Domain_Routing) notation (e.g. `10.96.0.0/14`) from the RFC-1918 private networks (e.g. `10.0.0.0/8`, `172.16.0.0/12`, `192.168.0.0/16`) to pick a specific range to use. Specify `services_secondary_range_name` or `services_ipv4_cidr_block` but not both.
        /// </summary>
        [Input("servicesIpv4CidrBlock")]
        public Input<string>? ServicesIpv4CidrBlock { get; set; }

        /// <summary>
        /// Optional. The name of the services' secondary range used to allocate IP addresses to the cluster. Specify either `services_secondary_range_name` or `services_ipv4_cidr_block` but not both. For Cloud Composer environments in versions composer-1.*.*-airflow-*.*.*, this field is applicable only when `use_ip_aliases` is true.
        /// </summary>
        [Input("servicesSecondaryRangeName")]
        public Input<string>? ServicesSecondaryRangeName { get; set; }

        /// <summary>
        /// Optional. Whether or not to enable Alias IPs in the GKE cluster. If `true`, a VPC-native cluster is created. This field is only supported for Cloud Composer environments in versions composer-1.*.*-airflow-*.*.*. Environments in newer versions always use VPC-native GKE clusters.
        /// </summary>
        [Input("useIpAliases")]
        public Input<bool>? UseIpAliases { get; set; }

        public IPAllocationPolicyArgs()
        {
        }
    }
}
