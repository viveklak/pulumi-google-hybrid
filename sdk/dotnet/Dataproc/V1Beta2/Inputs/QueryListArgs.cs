// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dataproc.V1Beta2.Inputs
{

    /// <summary>
    /// A list of queries to run on a cluster.
    /// </summary>
    public sealed class QueryListArgs : Pulumi.ResourceArgs
    {
        [Input("queries", required: true)]
        private InputList<string>? _queries;

        /// <summary>
        /// The queries to execute. You do not need to end a query expression with a semicolon. Multiple queries can be specified in one string by separating each with a semicolon. Here is an example of a Dataproc API snippet that uses a QueryList to specify a HiveJob: "hiveJob": { "queryList": { "queries": [ "query1", "query2", "query3;query4", ] } } 
        /// </summary>
        public InputList<string> Queries
        {
            get => _queries ?? (_queries = new InputList<string>());
            set => _queries = value;
        }

        public QueryListArgs()
        {
        }
    }
}
