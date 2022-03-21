// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Container.V1.Outputs
{

    /// <summary>
    /// AutoUpgradeOptions defines the set of options for the user to control how the Auto Upgrades will proceed.
    /// </summary>
    [OutputType]
    public sealed class AutoUpgradeOptionsResponse
    {
        /// <summary>
        /// [Output only] This field is set when upgrades are about to commence with the approximate start time for the upgrades, in [RFC3339](https://www.ietf.org/rfc/rfc3339.txt) text format.
        /// </summary>
        public readonly string AutoUpgradeStartTime;
        /// <summary>
        /// [Output only] This field is set when upgrades are about to commence with the description of the upgrade.
        /// </summary>
        public readonly string Description;

        [OutputConstructor]
        private AutoUpgradeOptionsResponse(
            string autoUpgradeStartTime,

            string description)
        {
            AutoUpgradeStartTime = autoUpgradeStartTime;
            Description = description;
        }
    }
}
