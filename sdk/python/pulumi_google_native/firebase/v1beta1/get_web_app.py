# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities

__all__ = [
    'GetWebAppResult',
    'AwaitableGetWebAppResult',
    'get_web_app',
    'get_web_app_output',
]

@pulumi.output_type
class GetWebAppResult:
    def __init__(__self__, api_key_id=None, app_id=None, app_urls=None, display_name=None, name=None, project=None, web_id=None):
        if api_key_id and not isinstance(api_key_id, str):
            raise TypeError("Expected argument 'api_key_id' to be a str")
        pulumi.set(__self__, "api_key_id", api_key_id)
        if app_id and not isinstance(app_id, str):
            raise TypeError("Expected argument 'app_id' to be a str")
        pulumi.set(__self__, "app_id", app_id)
        if app_urls and not isinstance(app_urls, list):
            raise TypeError("Expected argument 'app_urls' to be a list")
        pulumi.set(__self__, "app_urls", app_urls)
        if display_name and not isinstance(display_name, str):
            raise TypeError("Expected argument 'display_name' to be a str")
        pulumi.set(__self__, "display_name", display_name)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if project and not isinstance(project, str):
            raise TypeError("Expected argument 'project' to be a str")
        pulumi.set(__self__, "project", project)
        if web_id and not isinstance(web_id, str):
            raise TypeError("Expected argument 'web_id' to be a str")
        pulumi.set(__self__, "web_id", web_id)

    @property
    @pulumi.getter(name="apiKeyId")
    def api_key_id(self) -> str:
        """
        The key_id of the GCP ApiKey associated with this App. If set must have no restrictions, or only have restrictions that are valid for the associated Firebase App. Cannot be set in create requests, instead an existing valid API Key will be chosen, or if no valid API Keys exist, one will be provisioned for you. Cannot be set to an empty value in update requests.
        """
        return pulumi.get(self, "api_key_id")

    @property
    @pulumi.getter(name="appId")
    def app_id(self) -> str:
        """
        Immutable. The globally unique, Firebase-assigned identifier for the `WebApp`. This identifier should be treated as an opaque token, as the data format is not specified.
        """
        return pulumi.get(self, "app_id")

    @property
    @pulumi.getter(name="appUrls")
    def app_urls(self) -> Sequence[str]:
        """
        The URLs where the `WebApp` is hosted.
        """
        return pulumi.get(self, "app_urls")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> str:
        """
        The user-assigned display name for the `WebApp`.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The resource name of the WebApp, in the format: projects/PROJECT_IDENTIFIER /webApps/APP_ID * PROJECT_IDENTIFIER: the parent Project's [`ProjectNumber`](../projects#FirebaseProject.FIELDS.project_number) ***(recommended)*** or its [`ProjectId`](../projects#FirebaseProject.FIELDS.project_id). Learn more about using project identifiers in Google's [AIP 2510 standard](https://google.aip.dev/cloud/2510). Note that the value for PROJECT_IDENTIFIER in any response body will be the `ProjectId`. * APP_ID: the globally unique, Firebase-assigned identifier for the App (see [`appId`](../projects.webApps#WebApp.FIELDS.app_id)).
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def project(self) -> str:
        """
        Immutable. A user-assigned unique identifier of the parent FirebaseProject for the `WebApp`.
        """
        return pulumi.get(self, "project")

    @property
    @pulumi.getter(name="webId")
    def web_id(self) -> str:
        """
        Immutable. A unique, Firebase-assigned identifier for the `WebApp`. This identifier is only used to populate the `namespace` value for the `WebApp`. For most use cases, use `appId` to identify or reference the App. The `webId` value is only unique within a `FirebaseProject` and its associated Apps.
        """
        return pulumi.get(self, "web_id")


class AwaitableGetWebAppResult(GetWebAppResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetWebAppResult(
            api_key_id=self.api_key_id,
            app_id=self.app_id,
            app_urls=self.app_urls,
            display_name=self.display_name,
            name=self.name,
            project=self.project,
            web_id=self.web_id)


def get_web_app(project: Optional[str] = None,
                web_app_id: Optional[str] = None,
                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetWebAppResult:
    """
    Gets the specified WebApp.
    """
    __args__ = dict()
    __args__['project'] = project
    __args__['webAppId'] = web_app_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('google-native:firebase/v1beta1:getWebApp', __args__, opts=opts, typ=GetWebAppResult).value

    return AwaitableGetWebAppResult(
        api_key_id=__ret__.api_key_id,
        app_id=__ret__.app_id,
        app_urls=__ret__.app_urls,
        display_name=__ret__.display_name,
        name=__ret__.name,
        project=__ret__.project,
        web_id=__ret__.web_id)


@_utilities.lift_output_func(get_web_app)
def get_web_app_output(project: Optional[pulumi.Input[Optional[str]]] = None,
                       web_app_id: Optional[pulumi.Input[str]] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetWebAppResult]:
    """
    Gets the specified WebApp.
    """
    ...
