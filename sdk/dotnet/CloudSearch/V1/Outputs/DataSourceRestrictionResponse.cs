// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GcpNative.CloudSearch.V1.Outputs
{

    [OutputType]
    public sealed class DataSourceRestrictionResponse
    {
        /// <summary>
        /// Filter options restricting the results. If multiple filters are present, they are grouped by object type before joining. Filters with the same object type are joined conjunctively, then the resulting expressions are joined disjunctively. The maximum number of elements is 20. NOTE: Suggest API supports only few filters at the moment: "objecttype", "type" and "mimetype". For now, schema specific filters cannot be used to filter suggestions.
        /// </summary>
        public readonly ImmutableArray<Outputs.FilterOptionsResponse> FilterOptions;
        /// <summary>
        /// The source of restriction.
        /// </summary>
        public readonly Outputs.SourceResponse Source;

        [OutputConstructor]
        private DataSourceRestrictionResponse(
            ImmutableArray<Outputs.FilterOptionsResponse> filterOptions,

            Outputs.SourceResponse source)
        {
            FilterOptions = filterOptions;
            Source = source;
        }
    }
}