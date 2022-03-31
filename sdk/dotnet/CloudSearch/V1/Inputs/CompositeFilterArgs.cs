// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.CloudSearch.V1.Inputs
{

    public sealed class CompositeFilterArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The logic operator of the sub filter.
        /// </summary>
        [Input("logicOperator")]
        public Input<Pulumi.GoogleNative.CloudSearch.V1.CompositeFilterLogicOperator>? LogicOperator { get; set; }

        [Input("subFilters")]
        private InputList<Inputs.FilterArgs>? _subFilters;

        /// <summary>
        /// Sub filters.
        /// </summary>
        public InputList<Inputs.FilterArgs> SubFilters
        {
            get => _subFilters ?? (_subFilters = new InputList<Inputs.FilterArgs>());
            set => _subFilters = value;
        }

        public CompositeFilterArgs()
        {
        }
    }
}
