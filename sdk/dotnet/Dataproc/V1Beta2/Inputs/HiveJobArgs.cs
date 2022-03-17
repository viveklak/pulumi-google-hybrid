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
    /// A Dataproc job for running Apache Hive (https://hive.apache.org/) queries on YARN.
    /// </summary>
    public sealed class HiveJobArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Optional. Whether to continue executing queries if a query fails. The default value is false. Setting to true can be useful when executing independent parallel queries.
        /// </summary>
        [Input("continueOnFailure")]
        public Input<bool>? ContinueOnFailure { get; set; }

        [Input("jarFileUris")]
        private InputList<string>? _jarFileUris;

        /// <summary>
        /// Optional. HCFS URIs of jar files to add to the CLASSPATH of the Hive server and Hadoop MapReduce (MR) tasks. Can contain Hive SerDes and UDFs.
        /// </summary>
        public InputList<string> JarFileUris
        {
            get => _jarFileUris ?? (_jarFileUris = new InputList<string>());
            set => _jarFileUris = value;
        }

        [Input("properties")]
        private InputMap<string>? _properties;

        /// <summary>
        /// Optional. A mapping of property names and values, used to configure Hive. Properties that conflict with values set by the Dataproc API may be overwritten. Can include properties set in /etc/hadoop/conf/*-site.xml, /etc/hive/conf/hive-site.xml, and classes in user code.
        /// </summary>
        public InputMap<string> Properties
        {
            get => _properties ?? (_properties = new InputMap<string>());
            set => _properties = value;
        }

        /// <summary>
        /// The HCFS URI of the script that contains Hive queries.
        /// </summary>
        [Input("queryFileUri")]
        public Input<string>? QueryFileUri { get; set; }

        /// <summary>
        /// A list of queries.
        /// </summary>
        [Input("queryList")]
        public Input<Inputs.QueryListArgs>? QueryList { get; set; }

        [Input("scriptVariables")]
        private InputMap<string>? _scriptVariables;

        /// <summary>
        /// Optional. Mapping of query variable names to values (equivalent to the Hive command: SET name="value";).
        /// </summary>
        public InputMap<string> ScriptVariables
        {
            get => _scriptVariables ?? (_scriptVariables = new InputMap<string>());
            set => _scriptVariables = value;
        }

        public HiveJobArgs()
        {
        }
    }
}
