# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities
from . import outputs
from ._inputs import *

__all__ = ['AppArgs', 'App']

@pulumi.input_type
class AppArgs:
    def __init__(__self__, *,
                 apps_id: pulumi.Input[str],
                 auth_domain: Optional[pulumi.Input[str]] = None,
                 code_bucket: Optional[pulumi.Input[str]] = None,
                 database_type: Optional[pulumi.Input[str]] = None,
                 default_bucket: Optional[pulumi.Input[str]] = None,
                 default_cookie_expiration: Optional[pulumi.Input[str]] = None,
                 default_hostname: Optional[pulumi.Input[str]] = None,
                 dispatch_rules: Optional[pulumi.Input[Sequence[pulumi.Input['UrlDispatchRuleArgs']]]] = None,
                 feature_settings: Optional[pulumi.Input['FeatureSettingsArgs']] = None,
                 gcr_domain: Optional[pulumi.Input[str]] = None,
                 iap: Optional[pulumi.Input['IdentityAwareProxyArgs']] = None,
                 id: Optional[pulumi.Input[str]] = None,
                 location_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 serving_status: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a App resource.
        :param pulumi.Input[str] auth_domain: Google Apps authentication domain that controls which users can access this application.Defaults to open access for any Google Account.
        :param pulumi.Input[str] code_bucket: Google Cloud Storage bucket that can be used for storing files associated with this application. This bucket is associated with the application and can be used by the gcloud deployment commands.@OutputOnly
        :param pulumi.Input[str] database_type: The type of the Cloud Firestore or Cloud Datastore database associated with this application.
        :param pulumi.Input[str] default_bucket: Google Cloud Storage bucket that can be used by this application to store content.@OutputOnly
        :param pulumi.Input[str] default_cookie_expiration: Cookie expiration policy for this application.
        :param pulumi.Input[str] default_hostname: Hostname used to reach this application, as resolved by App Engine.@OutputOnly
        :param pulumi.Input[Sequence[pulumi.Input['UrlDispatchRuleArgs']]] dispatch_rules: HTTP path dispatch rules for requests to the application that do not explicitly target a service or version. Rules are order-dependent. Up to 20 dispatch rules can be supported.
        :param pulumi.Input['FeatureSettingsArgs'] feature_settings: The feature specific settings to be used in the application.
        :param pulumi.Input[str] gcr_domain: The Google Container Registry domain used for storing managed build docker images for this application.
        :param pulumi.Input[str] id: Identifier of the Application resource. This identifier is equivalent to the project ID of the Google Cloud Platform project where you want to deploy your application. Example: myapp.
        :param pulumi.Input[str] location_id: Location from which this application runs. Application instances run out of the data centers in the specified location, which is also where all of the application's end user content is stored.Defaults to us-central.View the list of supported locations (https://cloud.google.com/appengine/docs/locations).
        :param pulumi.Input[str] name: Full path to the Application resource in the API. Example: apps/myapp.@OutputOnly
        :param pulumi.Input[str] serving_status: Serving status of this application.
        """
        pulumi.set(__self__, "apps_id", apps_id)
        if auth_domain is not None:
            pulumi.set(__self__, "auth_domain", auth_domain)
        if code_bucket is not None:
            pulumi.set(__self__, "code_bucket", code_bucket)
        if database_type is not None:
            pulumi.set(__self__, "database_type", database_type)
        if default_bucket is not None:
            pulumi.set(__self__, "default_bucket", default_bucket)
        if default_cookie_expiration is not None:
            pulumi.set(__self__, "default_cookie_expiration", default_cookie_expiration)
        if default_hostname is not None:
            pulumi.set(__self__, "default_hostname", default_hostname)
        if dispatch_rules is not None:
            pulumi.set(__self__, "dispatch_rules", dispatch_rules)
        if feature_settings is not None:
            pulumi.set(__self__, "feature_settings", feature_settings)
        if gcr_domain is not None:
            pulumi.set(__self__, "gcr_domain", gcr_domain)
        if iap is not None:
            pulumi.set(__self__, "iap", iap)
        if id is not None:
            pulumi.set(__self__, "id", id)
        if location_id is not None:
            pulumi.set(__self__, "location_id", location_id)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if serving_status is not None:
            pulumi.set(__self__, "serving_status", serving_status)

    @property
    @pulumi.getter(name="appsId")
    def apps_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "apps_id")

    @apps_id.setter
    def apps_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "apps_id", value)

    @property
    @pulumi.getter(name="authDomain")
    def auth_domain(self) -> Optional[pulumi.Input[str]]:
        """
        Google Apps authentication domain that controls which users can access this application.Defaults to open access for any Google Account.
        """
        return pulumi.get(self, "auth_domain")

    @auth_domain.setter
    def auth_domain(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "auth_domain", value)

    @property
    @pulumi.getter(name="codeBucket")
    def code_bucket(self) -> Optional[pulumi.Input[str]]:
        """
        Google Cloud Storage bucket that can be used for storing files associated with this application. This bucket is associated with the application and can be used by the gcloud deployment commands.@OutputOnly
        """
        return pulumi.get(self, "code_bucket")

    @code_bucket.setter
    def code_bucket(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "code_bucket", value)

    @property
    @pulumi.getter(name="databaseType")
    def database_type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of the Cloud Firestore or Cloud Datastore database associated with this application.
        """
        return pulumi.get(self, "database_type")

    @database_type.setter
    def database_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "database_type", value)

    @property
    @pulumi.getter(name="defaultBucket")
    def default_bucket(self) -> Optional[pulumi.Input[str]]:
        """
        Google Cloud Storage bucket that can be used by this application to store content.@OutputOnly
        """
        return pulumi.get(self, "default_bucket")

    @default_bucket.setter
    def default_bucket(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "default_bucket", value)

    @property
    @pulumi.getter(name="defaultCookieExpiration")
    def default_cookie_expiration(self) -> Optional[pulumi.Input[str]]:
        """
        Cookie expiration policy for this application.
        """
        return pulumi.get(self, "default_cookie_expiration")

    @default_cookie_expiration.setter
    def default_cookie_expiration(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "default_cookie_expiration", value)

    @property
    @pulumi.getter(name="defaultHostname")
    def default_hostname(self) -> Optional[pulumi.Input[str]]:
        """
        Hostname used to reach this application, as resolved by App Engine.@OutputOnly
        """
        return pulumi.get(self, "default_hostname")

    @default_hostname.setter
    def default_hostname(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "default_hostname", value)

    @property
    @pulumi.getter(name="dispatchRules")
    def dispatch_rules(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['UrlDispatchRuleArgs']]]]:
        """
        HTTP path dispatch rules for requests to the application that do not explicitly target a service or version. Rules are order-dependent. Up to 20 dispatch rules can be supported.
        """
        return pulumi.get(self, "dispatch_rules")

    @dispatch_rules.setter
    def dispatch_rules(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['UrlDispatchRuleArgs']]]]):
        pulumi.set(self, "dispatch_rules", value)

    @property
    @pulumi.getter(name="featureSettings")
    def feature_settings(self) -> Optional[pulumi.Input['FeatureSettingsArgs']]:
        """
        The feature specific settings to be used in the application.
        """
        return pulumi.get(self, "feature_settings")

    @feature_settings.setter
    def feature_settings(self, value: Optional[pulumi.Input['FeatureSettingsArgs']]):
        pulumi.set(self, "feature_settings", value)

    @property
    @pulumi.getter(name="gcrDomain")
    def gcr_domain(self) -> Optional[pulumi.Input[str]]:
        """
        The Google Container Registry domain used for storing managed build docker images for this application.
        """
        return pulumi.get(self, "gcr_domain")

    @gcr_domain.setter
    def gcr_domain(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "gcr_domain", value)

    @property
    @pulumi.getter
    def iap(self) -> Optional[pulumi.Input['IdentityAwareProxyArgs']]:
        return pulumi.get(self, "iap")

    @iap.setter
    def iap(self, value: Optional[pulumi.Input['IdentityAwareProxyArgs']]):
        pulumi.set(self, "iap", value)

    @property
    @pulumi.getter
    def id(self) -> Optional[pulumi.Input[str]]:
        """
        Identifier of the Application resource. This identifier is equivalent to the project ID of the Google Cloud Platform project where you want to deploy your application. Example: myapp.
        """
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "id", value)

    @property
    @pulumi.getter(name="locationId")
    def location_id(self) -> Optional[pulumi.Input[str]]:
        """
        Location from which this application runs. Application instances run out of the data centers in the specified location, which is also where all of the application's end user content is stored.Defaults to us-central.View the list of supported locations (https://cloud.google.com/appengine/docs/locations).
        """
        return pulumi.get(self, "location_id")

    @location_id.setter
    def location_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "location_id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Full path to the Application resource in the API. Example: apps/myapp.@OutputOnly
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="servingStatus")
    def serving_status(self) -> Optional[pulumi.Input[str]]:
        """
        Serving status of this application.
        """
        return pulumi.get(self, "serving_status")

    @serving_status.setter
    def serving_status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "serving_status", value)


class App(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 apps_id: Optional[pulumi.Input[str]] = None,
                 auth_domain: Optional[pulumi.Input[str]] = None,
                 code_bucket: Optional[pulumi.Input[str]] = None,
                 database_type: Optional[pulumi.Input[str]] = None,
                 default_bucket: Optional[pulumi.Input[str]] = None,
                 default_cookie_expiration: Optional[pulumi.Input[str]] = None,
                 default_hostname: Optional[pulumi.Input[str]] = None,
                 dispatch_rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['UrlDispatchRuleArgs']]]]] = None,
                 feature_settings: Optional[pulumi.Input[pulumi.InputType['FeatureSettingsArgs']]] = None,
                 gcr_domain: Optional[pulumi.Input[str]] = None,
                 iap: Optional[pulumi.Input[pulumi.InputType['IdentityAwareProxyArgs']]] = None,
                 id: Optional[pulumi.Input[str]] = None,
                 location_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 serving_status: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Creates an App Engine application for a Google Cloud Platform project. Required fields: id - The ID of the target Cloud Platform project. location - The region (https://cloud.google.com/appengine/docs/locations) where you want the App Engine application located.For more information about App Engine applications, see Managing Projects, Applications, and Billing (https://cloud.google.com/appengine/docs/standard/python/console/).

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] auth_domain: Google Apps authentication domain that controls which users can access this application.Defaults to open access for any Google Account.
        :param pulumi.Input[str] code_bucket: Google Cloud Storage bucket that can be used for storing files associated with this application. This bucket is associated with the application and can be used by the gcloud deployment commands.@OutputOnly
        :param pulumi.Input[str] database_type: The type of the Cloud Firestore or Cloud Datastore database associated with this application.
        :param pulumi.Input[str] default_bucket: Google Cloud Storage bucket that can be used by this application to store content.@OutputOnly
        :param pulumi.Input[str] default_cookie_expiration: Cookie expiration policy for this application.
        :param pulumi.Input[str] default_hostname: Hostname used to reach this application, as resolved by App Engine.@OutputOnly
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['UrlDispatchRuleArgs']]]] dispatch_rules: HTTP path dispatch rules for requests to the application that do not explicitly target a service or version. Rules are order-dependent. Up to 20 dispatch rules can be supported.
        :param pulumi.Input[pulumi.InputType['FeatureSettingsArgs']] feature_settings: The feature specific settings to be used in the application.
        :param pulumi.Input[str] gcr_domain: The Google Container Registry domain used for storing managed build docker images for this application.
        :param pulumi.Input[str] id: Identifier of the Application resource. This identifier is equivalent to the project ID of the Google Cloud Platform project where you want to deploy your application. Example: myapp.
        :param pulumi.Input[str] location_id: Location from which this application runs. Application instances run out of the data centers in the specified location, which is also where all of the application's end user content is stored.Defaults to us-central.View the list of supported locations (https://cloud.google.com/appengine/docs/locations).
        :param pulumi.Input[str] name: Full path to the Application resource in the API. Example: apps/myapp.@OutputOnly
        :param pulumi.Input[str] serving_status: Serving status of this application.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AppArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates an App Engine application for a Google Cloud Platform project. Required fields: id - The ID of the target Cloud Platform project. location - The region (https://cloud.google.com/appengine/docs/locations) where you want the App Engine application located.For more information about App Engine applications, see Managing Projects, Applications, and Billing (https://cloud.google.com/appengine/docs/standard/python/console/).

        :param str resource_name: The name of the resource.
        :param AppArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AppArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 apps_id: Optional[pulumi.Input[str]] = None,
                 auth_domain: Optional[pulumi.Input[str]] = None,
                 code_bucket: Optional[pulumi.Input[str]] = None,
                 database_type: Optional[pulumi.Input[str]] = None,
                 default_bucket: Optional[pulumi.Input[str]] = None,
                 default_cookie_expiration: Optional[pulumi.Input[str]] = None,
                 default_hostname: Optional[pulumi.Input[str]] = None,
                 dispatch_rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['UrlDispatchRuleArgs']]]]] = None,
                 feature_settings: Optional[pulumi.Input[pulumi.InputType['FeatureSettingsArgs']]] = None,
                 gcr_domain: Optional[pulumi.Input[str]] = None,
                 iap: Optional[pulumi.Input[pulumi.InputType['IdentityAwareProxyArgs']]] = None,
                 id: Optional[pulumi.Input[str]] = None,
                 location_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 serving_status: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AppArgs.__new__(AppArgs)

            if apps_id is None and not opts.urn:
                raise TypeError("Missing required property 'apps_id'")
            __props__.__dict__["apps_id"] = apps_id
            __props__.__dict__["auth_domain"] = auth_domain
            __props__.__dict__["code_bucket"] = code_bucket
            __props__.__dict__["database_type"] = database_type
            __props__.__dict__["default_bucket"] = default_bucket
            __props__.__dict__["default_cookie_expiration"] = default_cookie_expiration
            __props__.__dict__["default_hostname"] = default_hostname
            __props__.__dict__["dispatch_rules"] = dispatch_rules
            __props__.__dict__["feature_settings"] = feature_settings
            __props__.__dict__["gcr_domain"] = gcr_domain
            __props__.__dict__["iap"] = iap
            __props__.__dict__["id"] = id
            __props__.__dict__["location_id"] = location_id
            __props__.__dict__["name"] = name
            __props__.__dict__["serving_status"] = serving_status
        super(App, __self__).__init__(
            'gcp-native:appengine/v1beta:App',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'App':
        """
        Get an existing App resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = AppArgs.__new__(AppArgs)

        __props__.__dict__["auth_domain"] = None
        __props__.__dict__["code_bucket"] = None
        __props__.__dict__["database_type"] = None
        __props__.__dict__["default_bucket"] = None
        __props__.__dict__["default_cookie_expiration"] = None
        __props__.__dict__["default_hostname"] = None
        __props__.__dict__["dispatch_rules"] = None
        __props__.__dict__["feature_settings"] = None
        __props__.__dict__["gcr_domain"] = None
        __props__.__dict__["iap"] = None
        __props__.__dict__["location_id"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["serving_status"] = None
        return App(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="authDomain")
    def auth_domain(self) -> pulumi.Output[str]:
        """
        Google Apps authentication domain that controls which users can access this application.Defaults to open access for any Google Account.
        """
        return pulumi.get(self, "auth_domain")

    @property
    @pulumi.getter(name="codeBucket")
    def code_bucket(self) -> pulumi.Output[str]:
        """
        Google Cloud Storage bucket that can be used for storing files associated with this application. This bucket is associated with the application and can be used by the gcloud deployment commands.@OutputOnly
        """
        return pulumi.get(self, "code_bucket")

    @property
    @pulumi.getter(name="databaseType")
    def database_type(self) -> pulumi.Output[str]:
        """
        The type of the Cloud Firestore or Cloud Datastore database associated with this application.
        """
        return pulumi.get(self, "database_type")

    @property
    @pulumi.getter(name="defaultBucket")
    def default_bucket(self) -> pulumi.Output[str]:
        """
        Google Cloud Storage bucket that can be used by this application to store content.@OutputOnly
        """
        return pulumi.get(self, "default_bucket")

    @property
    @pulumi.getter(name="defaultCookieExpiration")
    def default_cookie_expiration(self) -> pulumi.Output[str]:
        """
        Cookie expiration policy for this application.
        """
        return pulumi.get(self, "default_cookie_expiration")

    @property
    @pulumi.getter(name="defaultHostname")
    def default_hostname(self) -> pulumi.Output[str]:
        """
        Hostname used to reach this application, as resolved by App Engine.@OutputOnly
        """
        return pulumi.get(self, "default_hostname")

    @property
    @pulumi.getter(name="dispatchRules")
    def dispatch_rules(self) -> pulumi.Output[Sequence['outputs.UrlDispatchRuleResponse']]:
        """
        HTTP path dispatch rules for requests to the application that do not explicitly target a service or version. Rules are order-dependent. Up to 20 dispatch rules can be supported.
        """
        return pulumi.get(self, "dispatch_rules")

    @property
    @pulumi.getter(name="featureSettings")
    def feature_settings(self) -> pulumi.Output['outputs.FeatureSettingsResponse']:
        """
        The feature specific settings to be used in the application.
        """
        return pulumi.get(self, "feature_settings")

    @property
    @pulumi.getter(name="gcrDomain")
    def gcr_domain(self) -> pulumi.Output[str]:
        """
        The Google Container Registry domain used for storing managed build docker images for this application.
        """
        return pulumi.get(self, "gcr_domain")

    @property
    @pulumi.getter
    def iap(self) -> pulumi.Output['outputs.IdentityAwareProxyResponse']:
        return pulumi.get(self, "iap")

    @property
    @pulumi.getter(name="locationId")
    def location_id(self) -> pulumi.Output[str]:
        """
        Location from which this application runs. Application instances run out of the data centers in the specified location, which is also where all of the application's end user content is stored.Defaults to us-central.View the list of supported locations (https://cloud.google.com/appengine/docs/locations).
        """
        return pulumi.get(self, "location_id")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Full path to the Application resource in the API. Example: apps/myapp.@OutputOnly
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="servingStatus")
    def serving_status(self) -> pulumi.Output[str]:
        """
        Serving status of this application.
        """
        return pulumi.get(self, "serving_status")

