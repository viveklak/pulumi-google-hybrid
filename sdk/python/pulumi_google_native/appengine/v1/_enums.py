# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from enum import Enum

__all__ = [
    'ApiConfigHandlerAuthFailAction',
    'ApiConfigHandlerLogin',
    'ApiConfigHandlerSecurityLevel',
    'AppDatabaseType',
    'AppServingStatus',
    'EndpointsApiServiceRolloutStrategy',
    'ErrorHandlerErrorCode',
    'IngressRuleAction',
    'SslSettingsSslManagementType',
    'UrlMapAuthFailAction',
    'UrlMapLogin',
    'UrlMapRedirectHttpResponseCode',
    'UrlMapSecurityLevel',
    'VersionInboundServicesItem',
    'VersionServingStatus',
]


class ApiConfigHandlerAuthFailAction(str, Enum):
    """
    Action to take when users access resources that require authentication. Defaults to redirect.
    """
    AUTH_FAIL_ACTION_UNSPECIFIED = "AUTH_FAIL_ACTION_UNSPECIFIED"
    AUTH_FAIL_ACTION_REDIRECT = "AUTH_FAIL_ACTION_REDIRECT"
    AUTH_FAIL_ACTION_UNAUTHORIZED = "AUTH_FAIL_ACTION_UNAUTHORIZED"


class ApiConfigHandlerLogin(str, Enum):
    """
    Level of login required to access this resource. Defaults to optional.
    """
    LOGIN_UNSPECIFIED = "LOGIN_UNSPECIFIED"
    LOGIN_OPTIONAL = "LOGIN_OPTIONAL"
    LOGIN_ADMIN = "LOGIN_ADMIN"
    LOGIN_REQUIRED = "LOGIN_REQUIRED"


class ApiConfigHandlerSecurityLevel(str, Enum):
    """
    Security (HTTPS) enforcement for this URL.
    """
    SECURE_UNSPECIFIED = "SECURE_UNSPECIFIED"
    SECURE_DEFAULT = "SECURE_DEFAULT"
    SECURE_NEVER = "SECURE_NEVER"
    SECURE_OPTIONAL = "SECURE_OPTIONAL"
    SECURE_ALWAYS = "SECURE_ALWAYS"


class AppDatabaseType(str, Enum):
    """
    The type of the Cloud Firestore or Cloud Datastore database associated with this application.
    """
    DATABASE_TYPE_UNSPECIFIED = "DATABASE_TYPE_UNSPECIFIED"
    CLOUD_DATASTORE = "CLOUD_DATASTORE"
    CLOUD_FIRESTORE = "CLOUD_FIRESTORE"
    CLOUD_DATASTORE_COMPATIBILITY = "CLOUD_DATASTORE_COMPATIBILITY"


class AppServingStatus(str, Enum):
    """
    Serving status of this application.
    """
    UNSPECIFIED = "UNSPECIFIED"
    SERVING = "SERVING"
    USER_DISABLED = "USER_DISABLED"
    SYSTEM_DISABLED = "SYSTEM_DISABLED"


class EndpointsApiServiceRolloutStrategy(str, Enum):
    """
    Endpoints rollout strategy. If FIXED, config_id must be specified. If MANAGED, config_id must be omitted.
    """
    UNSPECIFIED_ROLLOUT_STRATEGY = "UNSPECIFIED_ROLLOUT_STRATEGY"
    FIXED = "FIXED"
    MANAGED = "MANAGED"


class ErrorHandlerErrorCode(str, Enum):
    """
    Error condition this handler applies to.
    """
    ERROR_CODE_UNSPECIFIED = "ERROR_CODE_UNSPECIFIED"
    ERROR_CODE_DEFAULT = "ERROR_CODE_DEFAULT"
    ERROR_CODE_OVER_QUOTA = "ERROR_CODE_OVER_QUOTA"
    ERROR_CODE_DOS_API_DENIAL = "ERROR_CODE_DOS_API_DENIAL"
    ERROR_CODE_TIMEOUT = "ERROR_CODE_TIMEOUT"


class IngressRuleAction(str, Enum):
    """
    The action to take on matched requests.
    """
    UNSPECIFIED_ACTION = "UNSPECIFIED_ACTION"
    ALLOW = "ALLOW"
    DENY = "DENY"


class SslSettingsSslManagementType(str, Enum):
    """
    SSL management type for this domain. If AUTOMATIC, a managed certificate is automatically provisioned. If MANUAL, certificate_id must be manually specified in order to configure SSL for this domain.
    """
    SSL_MANAGEMENT_TYPE_UNSPECIFIED = "SSL_MANAGEMENT_TYPE_UNSPECIFIED"
    AUTOMATIC = "AUTOMATIC"
    MANUAL = "MANUAL"


class UrlMapAuthFailAction(str, Enum):
    """
    Action to take when users access resources that require authentication. Defaults to redirect.
    """
    AUTH_FAIL_ACTION_UNSPECIFIED = "AUTH_FAIL_ACTION_UNSPECIFIED"
    AUTH_FAIL_ACTION_REDIRECT = "AUTH_FAIL_ACTION_REDIRECT"
    AUTH_FAIL_ACTION_UNAUTHORIZED = "AUTH_FAIL_ACTION_UNAUTHORIZED"


class UrlMapLogin(str, Enum):
    """
    Level of login required to access this resource. Not supported for Node.js in the App Engine standard environment.
    """
    LOGIN_UNSPECIFIED = "LOGIN_UNSPECIFIED"
    LOGIN_OPTIONAL = "LOGIN_OPTIONAL"
    LOGIN_ADMIN = "LOGIN_ADMIN"
    LOGIN_REQUIRED = "LOGIN_REQUIRED"


class UrlMapRedirectHttpResponseCode(str, Enum):
    """
    30x code to use when performing redirects for the secure field. Defaults to 302.
    """
    REDIRECT_HTTP_RESPONSE_CODE_UNSPECIFIED = "REDIRECT_HTTP_RESPONSE_CODE_UNSPECIFIED"
    REDIRECT_HTTP_RESPONSE_CODE301 = "REDIRECT_HTTP_RESPONSE_CODE_301"
    REDIRECT_HTTP_RESPONSE_CODE302 = "REDIRECT_HTTP_RESPONSE_CODE_302"
    REDIRECT_HTTP_RESPONSE_CODE303 = "REDIRECT_HTTP_RESPONSE_CODE_303"
    REDIRECT_HTTP_RESPONSE_CODE307 = "REDIRECT_HTTP_RESPONSE_CODE_307"


class UrlMapSecurityLevel(str, Enum):
    """
    Security (HTTPS) enforcement for this URL.
    """
    SECURE_UNSPECIFIED = "SECURE_UNSPECIFIED"
    SECURE_DEFAULT = "SECURE_DEFAULT"
    SECURE_NEVER = "SECURE_NEVER"
    SECURE_OPTIONAL = "SECURE_OPTIONAL"
    SECURE_ALWAYS = "SECURE_ALWAYS"


class VersionInboundServicesItem(str, Enum):
    INBOUND_SERVICE_UNSPECIFIED = "INBOUND_SERVICE_UNSPECIFIED"
    INBOUND_SERVICE_MAIL = "INBOUND_SERVICE_MAIL"
    INBOUND_SERVICE_MAIL_BOUNCE = "INBOUND_SERVICE_MAIL_BOUNCE"
    INBOUND_SERVICE_XMPP_ERROR = "INBOUND_SERVICE_XMPP_ERROR"
    INBOUND_SERVICE_XMPP_MESSAGE = "INBOUND_SERVICE_XMPP_MESSAGE"
    INBOUND_SERVICE_XMPP_SUBSCRIBE = "INBOUND_SERVICE_XMPP_SUBSCRIBE"
    INBOUND_SERVICE_XMPP_PRESENCE = "INBOUND_SERVICE_XMPP_PRESENCE"
    INBOUND_SERVICE_CHANNEL_PRESENCE = "INBOUND_SERVICE_CHANNEL_PRESENCE"
    INBOUND_SERVICE_WARMUP = "INBOUND_SERVICE_WARMUP"


class VersionServingStatus(str, Enum):
    """
    Current serving status of this version. Only the versions with a SERVING status create instances and can be billed.SERVING_STATUS_UNSPECIFIED is an invalid value. Defaults to SERVING.
    """
    SERVING_STATUS_UNSPECIFIED = "SERVING_STATUS_UNSPECIFIED"
    SERVING = "SERVING"
    STOPPED = "STOPPED"
