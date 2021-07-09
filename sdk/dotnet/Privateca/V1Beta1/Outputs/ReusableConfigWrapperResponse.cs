// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Privateca.V1Beta1.Outputs
{

    [OutputType]
    public sealed class ReusableConfigWrapperResponse
    {
        /// <summary>
        /// A resource path to a ReusableConfig in the format `projects/*/locations/*/reusableConfigs/*`.
        /// </summary>
        public readonly string ReusableConfig;
        /// <summary>
        /// A user-specified inline ReusableConfigValues.
        /// </summary>
        public readonly Outputs.ReusableConfigValuesResponse ReusableConfigValues;

        [OutputConstructor]
        private ReusableConfigWrapperResponse(
            string reusableConfig,

            Outputs.ReusableConfigValuesResponse reusableConfigValues)
        {
            ReusableConfig = reusableConfig;
            ReusableConfigValues = reusableConfigValues;
        }
    }
}
