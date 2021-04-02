// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleCloud.Ml.V1.Inputs
{

    /// <summary>
    /// Represents the spec to match discrete values from parent parameter.
    /// </summary>
    public sealed class GoogleCloudMlV1_StudyConfigParameterSpec_MatchingParentDiscreteValueSpecArgs : Pulumi.ResourceArgs
    {
        [Input("values")]
        private InputList<double>? _values;

        /// <summary>
        /// Matches values of the parent parameter with type 'DISCRETE'. All values must exist in `discrete_value_spec` of parent parameter.
        /// </summary>
        public InputList<double> Values
        {
            get => _values ?? (_values = new InputList<double>());
            set => _values = value;
        }

        public GoogleCloudMlV1_StudyConfigParameterSpec_MatchingParentDiscreteValueSpecArgs()
        {
        }
    }
}