# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities
from . import outputs

__all__ = [
    'GetVersionResult',
    'AwaitableGetVersionResult',
    'get_version',
    'get_version_output',
]

@pulumi.output_type
class GetVersionResult:
    def __init__(__self__, accelerator_config=None, auto_scaling=None, container=None, create_time=None, deployment_uri=None, description=None, error_message=None, etag=None, explanation_config=None, framework=None, is_default=None, labels=None, last_migration_model_id=None, last_migration_time=None, last_use_time=None, machine_type=None, manual_scaling=None, name=None, package_uris=None, prediction_class=None, python_version=None, request_logging_config=None, routes=None, runtime_version=None, service_account=None, state=None):
        if accelerator_config and not isinstance(accelerator_config, dict):
            raise TypeError("Expected argument 'accelerator_config' to be a dict")
        pulumi.set(__self__, "accelerator_config", accelerator_config)
        if auto_scaling and not isinstance(auto_scaling, dict):
            raise TypeError("Expected argument 'auto_scaling' to be a dict")
        pulumi.set(__self__, "auto_scaling", auto_scaling)
        if container and not isinstance(container, dict):
            raise TypeError("Expected argument 'container' to be a dict")
        pulumi.set(__self__, "container", container)
        if create_time and not isinstance(create_time, str):
            raise TypeError("Expected argument 'create_time' to be a str")
        pulumi.set(__self__, "create_time", create_time)
        if deployment_uri and not isinstance(deployment_uri, str):
            raise TypeError("Expected argument 'deployment_uri' to be a str")
        pulumi.set(__self__, "deployment_uri", deployment_uri)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if error_message and not isinstance(error_message, str):
            raise TypeError("Expected argument 'error_message' to be a str")
        pulumi.set(__self__, "error_message", error_message)
        if etag and not isinstance(etag, str):
            raise TypeError("Expected argument 'etag' to be a str")
        pulumi.set(__self__, "etag", etag)
        if explanation_config and not isinstance(explanation_config, dict):
            raise TypeError("Expected argument 'explanation_config' to be a dict")
        pulumi.set(__self__, "explanation_config", explanation_config)
        if framework and not isinstance(framework, str):
            raise TypeError("Expected argument 'framework' to be a str")
        pulumi.set(__self__, "framework", framework)
        if is_default and not isinstance(is_default, bool):
            raise TypeError("Expected argument 'is_default' to be a bool")
        pulumi.set(__self__, "is_default", is_default)
        if labels and not isinstance(labels, dict):
            raise TypeError("Expected argument 'labels' to be a dict")
        pulumi.set(__self__, "labels", labels)
        if last_migration_model_id and not isinstance(last_migration_model_id, str):
            raise TypeError("Expected argument 'last_migration_model_id' to be a str")
        pulumi.set(__self__, "last_migration_model_id", last_migration_model_id)
        if last_migration_time and not isinstance(last_migration_time, str):
            raise TypeError("Expected argument 'last_migration_time' to be a str")
        pulumi.set(__self__, "last_migration_time", last_migration_time)
        if last_use_time and not isinstance(last_use_time, str):
            raise TypeError("Expected argument 'last_use_time' to be a str")
        pulumi.set(__self__, "last_use_time", last_use_time)
        if machine_type and not isinstance(machine_type, str):
            raise TypeError("Expected argument 'machine_type' to be a str")
        pulumi.set(__self__, "machine_type", machine_type)
        if manual_scaling and not isinstance(manual_scaling, dict):
            raise TypeError("Expected argument 'manual_scaling' to be a dict")
        pulumi.set(__self__, "manual_scaling", manual_scaling)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if package_uris and not isinstance(package_uris, list):
            raise TypeError("Expected argument 'package_uris' to be a list")
        pulumi.set(__self__, "package_uris", package_uris)
        if prediction_class and not isinstance(prediction_class, str):
            raise TypeError("Expected argument 'prediction_class' to be a str")
        pulumi.set(__self__, "prediction_class", prediction_class)
        if python_version and not isinstance(python_version, str):
            raise TypeError("Expected argument 'python_version' to be a str")
        pulumi.set(__self__, "python_version", python_version)
        if request_logging_config and not isinstance(request_logging_config, dict):
            raise TypeError("Expected argument 'request_logging_config' to be a dict")
        pulumi.set(__self__, "request_logging_config", request_logging_config)
        if routes and not isinstance(routes, dict):
            raise TypeError("Expected argument 'routes' to be a dict")
        pulumi.set(__self__, "routes", routes)
        if runtime_version and not isinstance(runtime_version, str):
            raise TypeError("Expected argument 'runtime_version' to be a str")
        pulumi.set(__self__, "runtime_version", runtime_version)
        if service_account and not isinstance(service_account, str):
            raise TypeError("Expected argument 'service_account' to be a str")
        pulumi.set(__self__, "service_account", service_account)
        if state and not isinstance(state, str):
            raise TypeError("Expected argument 'state' to be a str")
        pulumi.set(__self__, "state", state)

    @property
    @pulumi.getter(name="acceleratorConfig")
    def accelerator_config(self) -> 'outputs.GoogleCloudMlV1__AcceleratorConfigResponse':
        """
        Optional. Accelerator config for using GPUs for online prediction (beta). Only specify this field if you have specified a Compute Engine (N1) machine type in the `machineType` field. Learn more about [using GPUs for online prediction](/ml-engine/docs/machine-types-online-prediction#gpus).
        """
        return pulumi.get(self, "accelerator_config")

    @property
    @pulumi.getter(name="autoScaling")
    def auto_scaling(self) -> 'outputs.GoogleCloudMlV1__AutoScalingResponse':
        """
        Automatically scale the number of nodes used to serve the model in response to increases and decreases in traffic. Care should be taken to ramp up traffic according to the model's ability to scale or you will start seeing increases in latency and 429 response codes.
        """
        return pulumi.get(self, "auto_scaling")

    @property
    @pulumi.getter
    def container(self) -> 'outputs.GoogleCloudMlV1__ContainerSpecResponse':
        """
        Optional. Specifies a custom container to use for serving predictions. If you specify this field, then `machineType` is required. If you specify this field, then `deploymentUri` is optional. If you specify this field, then you must not specify `runtimeVersion`, `packageUris`, `framework`, `pythonVersion`, or `predictionClass`.
        """
        return pulumi.get(self, "container")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> str:
        """
        The time the version was created.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter(name="deploymentUri")
    def deployment_uri(self) -> str:
        """
        The Cloud Storage URI of a directory containing trained model artifacts to be used to create the model version. See the [guide to deploying models](/ai-platform/prediction/docs/deploying-models) for more information. The total number of files under this directory must not exceed 1000. During projects.models.versions.create, AI Platform Prediction copies all files from the specified directory to a location managed by the service. From then on, AI Platform Prediction uses these copies of the model artifacts to serve predictions, not the original files in Cloud Storage, so this location is useful only as a historical record. If you specify container, then this field is optional. Otherwise, it is required. Learn [how to use this field with a custom container](/ai-platform/prediction/docs/custom-container-requirements#artifacts).
        """
        return pulumi.get(self, "deployment_uri")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        Optional. The description specified for the version when it was created.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="errorMessage")
    def error_message(self) -> str:
        """
        The details of a failure or a cancellation.
        """
        return pulumi.get(self, "error_message")

    @property
    @pulumi.getter
    def etag(self) -> str:
        """
        `etag` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a model from overwriting each other. It is strongly suggested that systems make use of the `etag` in the read-modify-write cycle to perform model updates in order to avoid race conditions: An `etag` is returned in the response to `GetVersion`, and systems are expected to put that etag in the request to `UpdateVersion` to ensure that their change will be applied to the model as intended.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter(name="explanationConfig")
    def explanation_config(self) -> 'outputs.GoogleCloudMlV1__ExplanationConfigResponse':
        """
        Optional. Configures explainability features on the model's version. Some explanation features require additional metadata to be loaded as part of the model payload.
        """
        return pulumi.get(self, "explanation_config")

    @property
    @pulumi.getter
    def framework(self) -> str:
        """
        Optional. The machine learning framework AI Platform uses to train this version of the model. Valid values are `TENSORFLOW`, `SCIKIT_LEARN`, `XGBOOST`. If you do not specify a framework, AI Platform will analyze files in the deployment_uri to determine a framework. If you choose `SCIKIT_LEARN` or `XGBOOST`, you must also set the runtime version of the model to 1.4 or greater. Do **not** specify a framework if you're deploying a [custom prediction routine](/ai-platform/prediction/docs/custom-prediction-routines) or if you're using a [custom container](/ai-platform/prediction/docs/use-custom-container).
        """
        return pulumi.get(self, "framework")

    @property
    @pulumi.getter(name="isDefault")
    def is_default(self) -> bool:
        """
        If true, this version will be used to handle prediction requests that do not specify a version. You can change the default version by calling projects.methods.versions.setDefault.
        """
        return pulumi.get(self, "is_default")

    @property
    @pulumi.getter
    def labels(self) -> Mapping[str, str]:
        """
        Optional. One or more labels that you can add, to organize your model versions. Each label is a key-value pair, where both the key and the value are arbitrary strings that you supply. For more information, see the documentation on using labels.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter(name="lastMigrationModelId")
    def last_migration_model_id(self) -> str:
        """
        The [AI Platform (Unified) `Model`](https://cloud.google.com/ai-platform-unified/docs/reference/rest/v1beta1/projects.locations.models) ID for the last [model migration](https://cloud.google.com/ai-platform-unified/docs/start/migrating-to-ai-platform-unified).
        """
        return pulumi.get(self, "last_migration_model_id")

    @property
    @pulumi.getter(name="lastMigrationTime")
    def last_migration_time(self) -> str:
        """
        The last time this version was successfully [migrated to AI Platform (Unified)](https://cloud.google.com/ai-platform-unified/docs/start/migrating-to-ai-platform-unified).
        """
        return pulumi.get(self, "last_migration_time")

    @property
    @pulumi.getter(name="lastUseTime")
    def last_use_time(self) -> str:
        """
        The time the version was last used for prediction.
        """
        return pulumi.get(self, "last_use_time")

    @property
    @pulumi.getter(name="machineType")
    def machine_type(self) -> str:
        """
        Optional. The type of machine on which to serve the model. Currently only applies to online prediction service. To learn about valid values for this field, read [Choosing a machine type for online prediction](/ai-platform/prediction/docs/machine-types-online-prediction). If this field is not specified and you are using a [regional endpoint](/ai-platform/prediction/docs/regional-endpoints), then the machine type defaults to `n1-standard-2`. If this field is not specified and you are using the global endpoint (`ml.googleapis.com`), then the machine type defaults to `mls1-c1-m2`.
        """
        return pulumi.get(self, "machine_type")

    @property
    @pulumi.getter(name="manualScaling")
    def manual_scaling(self) -> 'outputs.GoogleCloudMlV1__ManualScalingResponse':
        """
        Manually select the number of nodes to use for serving the model. You should generally use `auto_scaling` with an appropriate `min_nodes` instead, but this option is available if you want more predictable billing. Beware that latency and error rates will increase if the traffic exceeds that capability of the system to serve it based on the selected number of nodes.
        """
        return pulumi.get(self, "manual_scaling")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name specified for the version when it was created. The version name must be unique within the model it is created in.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="packageUris")
    def package_uris(self) -> Sequence[str]:
        """
        Optional. Cloud Storage paths (`gs://…`) of packages for [custom prediction routines](/ml-engine/docs/tensorflow/custom-prediction-routines) or [scikit-learn pipelines with custom code](/ml-engine/docs/scikit/exporting-for-prediction#custom-pipeline-code). For a custom prediction routine, one of these packages must contain your Predictor class (see [`predictionClass`](#Version.FIELDS.prediction_class)). Additionally, include any dependencies used by your Predictor or scikit-learn pipeline uses that are not already included in your selected [runtime version](/ml-engine/docs/tensorflow/runtime-version-list). If you specify this field, you must also set [`runtimeVersion`](#Version.FIELDS.runtime_version) to 1.4 or greater.
        """
        return pulumi.get(self, "package_uris")

    @property
    @pulumi.getter(name="predictionClass")
    def prediction_class(self) -> str:
        """
        Optional. The fully qualified name (module_name.class_name) of a class that implements the Predictor interface described in this reference field. The module containing this class should be included in a package provided to the [`packageUris` field](#Version.FIELDS.package_uris). Specify this field if and only if you are deploying a [custom prediction routine (beta)](/ml-engine/docs/tensorflow/custom-prediction-routines). If you specify this field, you must set [`runtimeVersion`](#Version.FIELDS.runtime_version) to 1.4 or greater and you must set `machineType` to a [legacy (MLS1) machine type](/ml-engine/docs/machine-types-online-prediction). The following code sample provides the Predictor interface: class Predictor(object): \"\"\"Interface for constructing custom predictors.\"\"\" def predict(self, instances, **kwargs): \"\"\"Performs custom prediction. Instances are the decoded values from the request. They have already been deserialized from JSON. Args: instances: A list of prediction input instances. **kwargs: A dictionary of keyword args provided as additional fields on the predict request body. Returns: A list of outputs containing the prediction results. This list must be JSON serializable. \"\"\" raise NotImplementedError() @classmethod def from_path(cls, model_dir): \"\"\"Creates an instance of Predictor using the given path. Loading of the predictor should be done in this method. Args: model_dir: The local directory that contains the exported model file along with any additional files uploaded when creating the version resource. Returns: An instance implementing this Predictor class. \"\"\" raise NotImplementedError() Learn more about [the Predictor interface and custom prediction routines](/ml-engine/docs/tensorflow/custom-prediction-routines).
        """
        return pulumi.get(self, "prediction_class")

    @property
    @pulumi.getter(name="pythonVersion")
    def python_version(self) -> str:
        """
        The version of Python used in prediction. The following Python versions are available: * Python '3.7' is available when `runtime_version` is set to '1.15' or later. * Python '3.5' is available when `runtime_version` is set to a version from '1.4' to '1.14'. * Python '2.7' is available when `runtime_version` is set to '1.15' or earlier. Read more about the Python versions available for [each runtime version](/ml-engine/docs/runtime-version-list).
        """
        return pulumi.get(self, "python_version")

    @property
    @pulumi.getter(name="requestLoggingConfig")
    def request_logging_config(self) -> 'outputs.GoogleCloudMlV1__RequestLoggingConfigResponse':
        """
        Optional. *Only* specify this field in a projects.models.versions.patch request. Specifying it in a projects.models.versions.create request has no effect. Configures the request-response pair logging on predictions from this Version.
        """
        return pulumi.get(self, "request_logging_config")

    @property
    @pulumi.getter
    def routes(self) -> 'outputs.GoogleCloudMlV1__RouteMapResponse':
        """
        Optional. Specifies paths on a custom container's HTTP server where AI Platform Prediction sends certain requests. If you specify this field, then you must also specify the `container` field. If you specify the `container` field and do not specify this field, it defaults to the following: ```json { "predict": "/v1/models/MODEL/versions/VERSION:predict", "health": "/v1/models/MODEL/versions/VERSION" } ``` See RouteMap for more details about these default values.
        """
        return pulumi.get(self, "routes")

    @property
    @pulumi.getter(name="runtimeVersion")
    def runtime_version(self) -> str:
        """
        The AI Platform runtime version to use for this deployment. For more information, see the [runtime version list](/ml-engine/docs/runtime-version-list) and [how to manage runtime versions](/ml-engine/docs/versioning).
        """
        return pulumi.get(self, "runtime_version")

    @property
    @pulumi.getter(name="serviceAccount")
    def service_account(self) -> str:
        """
        Optional. Specifies the service account for resource access control. If you specify this field, then you must also specify either the `containerSpec` or the `predictionClass` field. Learn more about [using a custom service account](/ai-platform/prediction/docs/custom-service-account).
        """
        return pulumi.get(self, "service_account")

    @property
    @pulumi.getter
    def state(self) -> str:
        """
        The state of a version.
        """
        return pulumi.get(self, "state")


class AwaitableGetVersionResult(GetVersionResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetVersionResult(
            accelerator_config=self.accelerator_config,
            auto_scaling=self.auto_scaling,
            container=self.container,
            create_time=self.create_time,
            deployment_uri=self.deployment_uri,
            description=self.description,
            error_message=self.error_message,
            etag=self.etag,
            explanation_config=self.explanation_config,
            framework=self.framework,
            is_default=self.is_default,
            labels=self.labels,
            last_migration_model_id=self.last_migration_model_id,
            last_migration_time=self.last_migration_time,
            last_use_time=self.last_use_time,
            machine_type=self.machine_type,
            manual_scaling=self.manual_scaling,
            name=self.name,
            package_uris=self.package_uris,
            prediction_class=self.prediction_class,
            python_version=self.python_version,
            request_logging_config=self.request_logging_config,
            routes=self.routes,
            runtime_version=self.runtime_version,
            service_account=self.service_account,
            state=self.state)


def get_version(model_id: Optional[str] = None,
                project: Optional[str] = None,
                version_id: Optional[str] = None,
                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetVersionResult:
    """
    Gets information about a model version. Models can have multiple versions. You can call projects.models.versions.list to get the same information that this method returns for all of the versions of a model.
    """
    __args__ = dict()
    __args__['modelId'] = model_id
    __args__['project'] = project
    __args__['versionId'] = version_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('google-native:ml/v1:getVersion', __args__, opts=opts, typ=GetVersionResult).value

    return AwaitableGetVersionResult(
        accelerator_config=__ret__.accelerator_config,
        auto_scaling=__ret__.auto_scaling,
        container=__ret__.container,
        create_time=__ret__.create_time,
        deployment_uri=__ret__.deployment_uri,
        description=__ret__.description,
        error_message=__ret__.error_message,
        etag=__ret__.etag,
        explanation_config=__ret__.explanation_config,
        framework=__ret__.framework,
        is_default=__ret__.is_default,
        labels=__ret__.labels,
        last_migration_model_id=__ret__.last_migration_model_id,
        last_migration_time=__ret__.last_migration_time,
        last_use_time=__ret__.last_use_time,
        machine_type=__ret__.machine_type,
        manual_scaling=__ret__.manual_scaling,
        name=__ret__.name,
        package_uris=__ret__.package_uris,
        prediction_class=__ret__.prediction_class,
        python_version=__ret__.python_version,
        request_logging_config=__ret__.request_logging_config,
        routes=__ret__.routes,
        runtime_version=__ret__.runtime_version,
        service_account=__ret__.service_account,
        state=__ret__.state)


@_utilities.lift_output_func(get_version)
def get_version_output(model_id: Optional[pulumi.Input[str]] = None,
                       project: Optional[pulumi.Input[Optional[str]]] = None,
                       version_id: Optional[pulumi.Input[str]] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetVersionResult]:
    """
    Gets information about a model version. Models can have multiple versions. You can call projects.models.versions.list to get the same information that this method returns for all of the versions of a model.
    """
    ...
