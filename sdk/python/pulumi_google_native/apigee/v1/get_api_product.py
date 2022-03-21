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
    'GetApiProductResult',
    'AwaitableGetApiProductResult',
    'get_api_product',
    'get_api_product_output',
]

@pulumi.output_type
class GetApiProductResult:
    def __init__(__self__, api_resources=None, approval_type=None, attributes=None, created_at=None, description=None, display_name=None, environments=None, graphql_operation_group=None, last_modified_at=None, name=None, operation_group=None, proxies=None, quota=None, quota_interval=None, quota_time_unit=None, scopes=None):
        if api_resources and not isinstance(api_resources, list):
            raise TypeError("Expected argument 'api_resources' to be a list")
        pulumi.set(__self__, "api_resources", api_resources)
        if approval_type and not isinstance(approval_type, str):
            raise TypeError("Expected argument 'approval_type' to be a str")
        pulumi.set(__self__, "approval_type", approval_type)
        if attributes and not isinstance(attributes, list):
            raise TypeError("Expected argument 'attributes' to be a list")
        pulumi.set(__self__, "attributes", attributes)
        if created_at and not isinstance(created_at, str):
            raise TypeError("Expected argument 'created_at' to be a str")
        pulumi.set(__self__, "created_at", created_at)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if display_name and not isinstance(display_name, str):
            raise TypeError("Expected argument 'display_name' to be a str")
        pulumi.set(__self__, "display_name", display_name)
        if environments and not isinstance(environments, list):
            raise TypeError("Expected argument 'environments' to be a list")
        pulumi.set(__self__, "environments", environments)
        if graphql_operation_group and not isinstance(graphql_operation_group, dict):
            raise TypeError("Expected argument 'graphql_operation_group' to be a dict")
        pulumi.set(__self__, "graphql_operation_group", graphql_operation_group)
        if last_modified_at and not isinstance(last_modified_at, str):
            raise TypeError("Expected argument 'last_modified_at' to be a str")
        pulumi.set(__self__, "last_modified_at", last_modified_at)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if operation_group and not isinstance(operation_group, dict):
            raise TypeError("Expected argument 'operation_group' to be a dict")
        pulumi.set(__self__, "operation_group", operation_group)
        if proxies and not isinstance(proxies, list):
            raise TypeError("Expected argument 'proxies' to be a list")
        pulumi.set(__self__, "proxies", proxies)
        if quota and not isinstance(quota, str):
            raise TypeError("Expected argument 'quota' to be a str")
        pulumi.set(__self__, "quota", quota)
        if quota_interval and not isinstance(quota_interval, str):
            raise TypeError("Expected argument 'quota_interval' to be a str")
        pulumi.set(__self__, "quota_interval", quota_interval)
        if quota_time_unit and not isinstance(quota_time_unit, str):
            raise TypeError("Expected argument 'quota_time_unit' to be a str")
        pulumi.set(__self__, "quota_time_unit", quota_time_unit)
        if scopes and not isinstance(scopes, list):
            raise TypeError("Expected argument 'scopes' to be a list")
        pulumi.set(__self__, "scopes", scopes)

    @property
    @pulumi.getter(name="apiResources")
    def api_resources(self) -> Sequence[str]:
        return pulumi.get(self, "api_resources")

    @property
    @pulumi.getter(name="approvalType")
    def approval_type(self) -> str:
        """
        Flag that specifies how API keys are approved to access the APIs defined by the API product. If set to `manual`, the consumer key is generated and returned in "pending" state. In this case, the API keys won't work until they have been explicitly approved. If set to `auto`, the consumer key is generated and returned in "approved" state and can be used immediately. **Note:** Typically, `auto` is used to provide access to free or trial API products that provide limited quota or capabilities.
        """
        return pulumi.get(self, "approval_type")

    @property
    @pulumi.getter
    def attributes(self) -> Sequence['outputs.GoogleCloudApigeeV1AttributeResponse']:
        """
        Array of attributes that may be used to extend the default API product profile with customer-specific metadata. You can specify a maximum of 18 attributes. Use this property to specify the access level of the API product as either `public`, `private`, or `internal`. Only products marked `public` are available to developers in the Apigee developer portal. For example, you can set a product to `internal` while it is in development and then change access to `public` when it is ready to release on the portal. API products marked as `private` do not appear on the portal, but can be accessed by external developers.
        """
        return pulumi.get(self, "attributes")

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> str:
        """
        Response only. Creation time of this environment as milliseconds since epoch.
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        Description of the API product. Include key information about the API product that is not captured by other fields. Comma-separated list of API resources to be bundled in the API product. By default, the resource paths are mapped from the `proxy.pathsuffix` variable. The proxy path suffix is defined as the URI fragment following the ProxyEndpoint base path. For example, if the `apiResources` element is defined to be `/forecastrss` and the base path defined for the API proxy is `/weather`, then only requests to `/weather/forecastrss` are permitted by the API product. You can select a specific path, or you can select all subpaths with the following wildcard: - `/**`: Indicates that all sub-URIs are included. - `/*` : Indicates that only URIs one level down are included. By default, / supports the same resources as /** as well as the base path defined by the API proxy. For example, if the base path of the API proxy is `/v1/weatherapikey`, then the API product supports requests to `/v1/weatherapikey` and to any sub-URIs, such as `/v1/weatherapikey/forecastrss`, `/v1/weatherapikey/region/CA`, and so on. For more information, see Managing API products.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> str:
        """
        Name displayed in the UI or developer portal to developers registering for API access.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter
    def environments(self) -> Sequence[str]:
        """
        Comma-separated list of environment names to which the API product is bound. Requests to environments that are not listed are rejected. By specifying one or more environments, you can bind the resources listed in the API product to a specific environment, preventing developers from accessing those resources through API proxies deployed in another environment. This setting is used, for example, to prevent resources associated with API proxies in `prod` from being accessed by API proxies deployed in `test`.
        """
        return pulumi.get(self, "environments")

    @property
    @pulumi.getter(name="graphqlOperationGroup")
    def graphql_operation_group(self) -> 'outputs.GoogleCloudApigeeV1GraphQLOperationGroupResponse':
        """
        Configuration used to group Apigee proxies or remote services with graphQL operation name, graphQL operation type and quotas. This grouping allows us to precisely set quota for a particular combination of graphQL name and operation type for a particular proxy request. If graphQL name is not set, this would imply quota will be applied on all graphQL requests matching the operation type.
        """
        return pulumi.get(self, "graphql_operation_group")

    @property
    @pulumi.getter(name="lastModifiedAt")
    def last_modified_at(self) -> str:
        """
        Response only. Modified time of this environment as milliseconds since epoch.
        """
        return pulumi.get(self, "last_modified_at")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Internal name of the API product. Characters you can use in the name are restricted to: `A-Z0-9._\-$ %`. **Note:** The internal name cannot be edited when updating the API product.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="operationGroup")
    def operation_group(self) -> 'outputs.GoogleCloudApigeeV1OperationGroupResponse':
        """
        Configuration used to group Apigee proxies or remote services with resources, method types, and quotas. The resource refers to the resource URI (excluding the base path). With this grouping, the API product creator is able to fine-tune and give precise control over which REST methods have access to specific resources and how many calls can be made (using the `quota` setting). **Note:** The `api_resources` setting cannot be specified for both the API product and operation group; otherwise the call will fail.
        """
        return pulumi.get(self, "operation_group")

    @property
    @pulumi.getter
    def proxies(self) -> Sequence[str]:
        """
        Comma-separated list of API proxy names to which this API product is bound. By specifying API proxies, you can associate resources in the API product with specific API proxies, preventing developers from accessing those resources through other API proxies. Apigee rejects requests to API proxies that are not listed. **Note:** The API proxy names must already exist in the specified environment as they will be validated upon creation.
        """
        return pulumi.get(self, "proxies")

    @property
    @pulumi.getter
    def quota(self) -> str:
        """
        Number of request messages permitted per app by this API product for the specified `quotaInterval` and `quotaTimeUnit`. For example, a `quota` of 50, for a `quotaInterval` of 12 and a `quotaTimeUnit` of hours means 50 requests are allowed every 12 hours.
        """
        return pulumi.get(self, "quota")

    @property
    @pulumi.getter(name="quotaInterval")
    def quota_interval(self) -> str:
        """
        Time interval over which the number of request messages is calculated.
        """
        return pulumi.get(self, "quota_interval")

    @property
    @pulumi.getter(name="quotaTimeUnit")
    def quota_time_unit(self) -> str:
        """
        Time unit defined for the `quotaInterval`. Valid values include `minute`, `hour`, `day`, or `month`.
        """
        return pulumi.get(self, "quota_time_unit")

    @property
    @pulumi.getter
    def scopes(self) -> Sequence[str]:
        """
        Comma-separated list of OAuth scopes that are validated at runtime. Apigee validates that the scopes in any access token presented match the scopes defined in the OAuth policy associated with the API product.
        """
        return pulumi.get(self, "scopes")


class AwaitableGetApiProductResult(GetApiProductResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetApiProductResult(
            api_resources=self.api_resources,
            approval_type=self.approval_type,
            attributes=self.attributes,
            created_at=self.created_at,
            description=self.description,
            display_name=self.display_name,
            environments=self.environments,
            graphql_operation_group=self.graphql_operation_group,
            last_modified_at=self.last_modified_at,
            name=self.name,
            operation_group=self.operation_group,
            proxies=self.proxies,
            quota=self.quota,
            quota_interval=self.quota_interval,
            quota_time_unit=self.quota_time_unit,
            scopes=self.scopes)


def get_api_product(apiproduct_id: Optional[str] = None,
                    organization_id: Optional[str] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetApiProductResult:
    """
    Gets configuration details for an API product. The API product name required in the request URL is the internal name of the product, not the display name. While they may be the same, it depends on whether the API product was created via the UI or the API. View the list of API products to verify the internal name.
    """
    __args__ = dict()
    __args__['apiproductId'] = apiproduct_id
    __args__['organizationId'] = organization_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('google-native:apigee/v1:getApiProduct', __args__, opts=opts, typ=GetApiProductResult).value

    return AwaitableGetApiProductResult(
        api_resources=__ret__.api_resources,
        approval_type=__ret__.approval_type,
        attributes=__ret__.attributes,
        created_at=__ret__.created_at,
        description=__ret__.description,
        display_name=__ret__.display_name,
        environments=__ret__.environments,
        graphql_operation_group=__ret__.graphql_operation_group,
        last_modified_at=__ret__.last_modified_at,
        name=__ret__.name,
        operation_group=__ret__.operation_group,
        proxies=__ret__.proxies,
        quota=__ret__.quota,
        quota_interval=__ret__.quota_interval,
        quota_time_unit=__ret__.quota_time_unit,
        scopes=__ret__.scopes)


@_utilities.lift_output_func(get_api_product)
def get_api_product_output(apiproduct_id: Optional[pulumi.Input[str]] = None,
                           organization_id: Optional[pulumi.Input[str]] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetApiProductResult]:
    """
    Gets configuration details for an API product. The API product name required in the request URL is the internal name of the product, not the display name. While they may be the same, it depends on whether the API product was created via the UI or the API. View the list of API products to verify the internal name.
    """
    ...
