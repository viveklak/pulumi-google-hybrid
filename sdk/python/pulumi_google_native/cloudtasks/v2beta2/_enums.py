# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from enum import Enum

__all__ = [
    'AppEngineHttpRequestHttpMethod',
    'TaskResponseView',
]


class AppEngineHttpRequestHttpMethod(str, Enum):
    """
    The HTTP method to use for the request. The default is POST. The app's request handler for the task's target URL must be able to handle HTTP requests with this http_method, otherwise the task attempt fails with error code 405 (Method Not Allowed). See [Writing a push task request handler](https://cloud.google.com/appengine/docs/java/taskqueue/push/creating-handlers#writing_a_push_task_request_handler) and the App Engine documentation for your runtime on [How Requests are Handled](https://cloud.google.com/appengine/docs/standard/python3/how-requests-are-handled).
    """
    HTTP_METHOD_UNSPECIFIED = "HTTP_METHOD_UNSPECIFIED"
    """
    HTTP method unspecified
    """
    POST = "POST"
    """
    HTTP POST
    """
    GET = "GET"
    """
    HTTP GET
    """
    HEAD = "HEAD"
    """
    HTTP HEAD
    """
    PUT = "PUT"
    """
    HTTP PUT
    """
    DELETE = "DELETE"
    """
    HTTP DELETE
    """


class TaskResponseView(str, Enum):
    """
    The response_view specifies which subset of the Task will be returned. By default response_view is BASIC; not all information is retrieved by default because some data, such as payloads, might be desirable to return only when needed because of its large size or because of the sensitivity of data that it contains. Authorization for FULL requires `cloudtasks.tasks.fullView` [Google IAM](https://cloud.google.com/iam/) permission on the Task resource.
    """
    VIEW_UNSPECIFIED = "VIEW_UNSPECIFIED"
    """
    Unspecified. Defaults to BASIC.
    """
    BASIC = "BASIC"
    """
    The basic view omits fields which can be large or can contain sensitive data. This view does not include the (payload in AppEngineHttpRequest and payload in PullMessage). These payloads are desirable to return only when needed, because they can be large and because of the sensitivity of the data that you choose to store in it.
    """
    FULL = "FULL"
    """
    All information is returned. Authorization for FULL requires `cloudtasks.tasks.fullView` [Google IAM](https://cloud.google.com/iam/) permission on the Queue resource.
    """
