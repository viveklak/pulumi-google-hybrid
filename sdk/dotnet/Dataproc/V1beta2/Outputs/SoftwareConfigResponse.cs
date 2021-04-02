// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleCloud.Dataproc.V1beta2.Outputs
{

    [OutputType]
    public sealed class SoftwareConfigResponse
    {
        /// <summary>
        /// Optional. The version of software inside the cluster. It must be one of the supported Dataproc Versions (https://cloud.google.com/dataproc/docs/concepts/versioning/dataproc-versions#supported_dataproc_versions), such as "1.2" (including a subminor version, such as "1.2.29"), or the "preview" version (https://cloud.google.com/dataproc/docs/concepts/versioning/dataproc-versions#other_versions). If unspecified, it defaults to the latest Debian version.
        /// </summary>
        public readonly string ImageVersion;
        /// <summary>
        /// The set of optional components to activate on the cluster.
        /// </summary>
        public readonly ImmutableArray<string> OptionalComponents;
        /// <summary>
        /// Optional. The properties to set on daemon config files.Property keys are specified in prefix:property format, for example core:hadoop.tmp.dir. The following are supported prefixes and their mappings: capacity-scheduler: capacity-scheduler.xml core: core-site.xml distcp: distcp-default.xml hdfs: hdfs-site.xml hive: hive-site.xml mapred: mapred-site.xml pig: pig.properties spark: spark-defaults.conf yarn: yarn-site.xmlFor more information, see Cluster properties (https://cloud.google.com/dataproc/docs/concepts/cluster-properties).
        /// </summary>
        public readonly ImmutableDictionary<string, string> Properties;

        [OutputConstructor]
        private SoftwareConfigResponse(
            string imageVersion,

            ImmutableArray<string> optionalComponents,

            ImmutableDictionary<string, string> properties)
        {
            ImageVersion = imageVersion;
            OptionalComponents = optionalComponents;
            Properties = properties;
        }
    }
}