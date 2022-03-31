// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.ContainerAnalysis.V1.Inputs
{

    /// <summary>
    /// An Upgrade Occurrence represents that a specific resource_url could install a specific upgrade. This presence is supplied via local sources (i.e. it is present in the mirror and the running system has noticed its availability). For Windows, both distribution and windows_update contain information for the Windows update.
    /// </summary>
    public sealed class UpgradeOccurrenceArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Metadata about the upgrade for available for the specific operating system for the resource_url. This allows efficient filtering, as well as making it easier to use the occurrence.
        /// </summary>
        [Input("distribution")]
        public Input<Inputs.UpgradeDistributionArgs>? Distribution { get; set; }

        /// <summary>
        /// Required for non-Windows OS. The package this Upgrade is for.
        /// </summary>
        [Input("package")]
        public Input<string>? Package { get; set; }

        /// <summary>
        /// Required for non-Windows OS. The version of the package in a machine + human readable form.
        /// </summary>
        [Input("parsedVersion")]
        public Input<Inputs.VersionArgs>? ParsedVersion { get; set; }

        /// <summary>
        /// Required for Windows OS. Represents the metadata about the Windows update.
        /// </summary>
        [Input("windowsUpdate")]
        public Input<Inputs.WindowsUpdateArgs>? WindowsUpdate { get; set; }

        public UpgradeOccurrenceArgs()
        {
        }
    }
}
