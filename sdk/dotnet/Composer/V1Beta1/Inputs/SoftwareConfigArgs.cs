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
    /// Specifies the selection and configuration of software inside the environment.
    /// </summary>
    public sealed class SoftwareConfigArgs : Pulumi.ResourceArgs
    {
        [Input("airflowConfigOverrides")]
        private InputMap<string>? _airflowConfigOverrides;

        /// <summary>
        /// Optional. Apache Airflow configuration properties to override. Property keys contain the section and property names, separated by a hyphen, for example "core-dags_are_paused_at_creation". Section names must not contain hyphens ("-"), opening square brackets ("["), or closing square brackets ("]"). The property name must not be empty and must not contain an equals sign ("=") or semicolon (";"). Section and property names must not contain a period ("."). Apache Airflow configuration property names must be written in [snake_case](https://en.wikipedia.org/wiki/Snake_case). Property values can contain any character, and can be written in any lower/upper case format. Certain Apache Airflow configuration property values are [blocked](/composer/docs/concepts/airflow-configurations), and cannot be overridden.
        /// </summary>
        public InputMap<string> AirflowConfigOverrides
        {
            get => _airflowConfigOverrides ?? (_airflowConfigOverrides = new InputMap<string>());
            set => _airflowConfigOverrides = value;
        }

        [Input("envVariables")]
        private InputMap<string>? _envVariables;

        /// <summary>
        /// Optional. Additional environment variables to provide to the Apache Airflow scheduler, worker, and webserver processes. Environment variable names must match the regular expression `a-zA-Z_*`. They cannot specify Apache Airflow software configuration overrides (they cannot match the regular expression `AIRFLOW__[A-Z0-9_]+__[A-Z0-9_]+`), and they cannot match any of the following reserved names: * `AIRFLOW_HOME` * `C_FORCE_ROOT` * `CONTAINER_NAME` * `DAGS_FOLDER` * `GCP_PROJECT` * `GCS_BUCKET` * `GKE_CLUSTER_NAME` * `SQL_DATABASE` * `SQL_INSTANCE` * `SQL_PASSWORD` * `SQL_PROJECT` * `SQL_REGION` * `SQL_USER`
        /// </summary>
        public InputMap<string> EnvVariables
        {
            get => _envVariables ?? (_envVariables = new InputMap<string>());
            set => _envVariables = value;
        }

        /// <summary>
        /// The version of the software running in the environment. This encapsulates both the version of Cloud Composer functionality and the version of Apache Airflow. It must match the regular expression `composer-([0-9]+(\.[0-9]+\.[0-9]+(-preview\.[0-9]+)?)?|latest)-airflow-([0-9]+\.[0-9]+(\.[0-9]+)?)`. When used as input, the server also checks if the provided version is supported and denies the request for an unsupported version. The Cloud Composer portion of the image version is a full [semantic version](https://semver.org), or an alias in the form of major version number or `latest`. When an alias is provided, the server replaces it with the current Cloud Composer version that satisfies the alias. The Apache Airflow portion of the image version is a full semantic version that points to one of the supported Apache Airflow versions, or an alias in the form of only major and minor versions specified. When an alias is provided, the server replaces it with the latest Apache Airflow version that satisfies the alias and is supported in the given Cloud Composer version. In all cases, the resolved image version is stored in the same field. See also [version list](/composer/docs/concepts/versioning/composer-versions) and [versioning overview](/composer/docs/concepts/versioning/composer-versioning-overview).
        /// </summary>
        [Input("imageVersion")]
        public Input<string>? ImageVersion { get; set; }

        [Input("pypiPackages")]
        private InputMap<string>? _pypiPackages;

        /// <summary>
        /// Optional. Custom Python Package Index (PyPI) packages to be installed in the environment. Keys refer to the lowercase package name such as "numpy" and values are the lowercase extras and version specifier such as "==1.12.0", "[devel,gcp_api]", or "[devel]&gt;=1.8.2, &lt;1.9.2". To specify a package without pinning it to a version specifier, use the empty string as the value.
        /// </summary>
        public InputMap<string> PypiPackages
        {
            get => _pypiPackages ?? (_pypiPackages = new InputMap<string>());
            set => _pypiPackages = value;
        }

        /// <summary>
        /// Optional. The major version of Python used to run the Apache Airflow scheduler, worker, and webserver processes. Can be set to '2' or '3'. If not specified, the default is '3'. Cannot be updated. This field is only supported for Cloud Composer environments in versions composer-1.*.*-airflow-*.*.*. Environments in newer versions always use Python major version 3.
        /// </summary>
        [Input("pythonVersion")]
        public Input<string>? PythonVersion { get; set; }

        /// <summary>
        /// Optional. The number of schedulers for Airflow. This field is supported for Cloud Composer environments in versions composer-1.*.*-airflow-2.*.*.
        /// </summary>
        [Input("schedulerCount")]
        public Input<int>? SchedulerCount { get; set; }

        public SoftwareConfigArgs()
        {
        }
    }
}
