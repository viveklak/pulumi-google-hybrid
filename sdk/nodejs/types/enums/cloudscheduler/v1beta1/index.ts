// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***


export const AppEngineHttpTargetHttpMethod = {
    /**
     * HTTP method unspecified. Defaults to POST.
     */
    HttpMethodUnspecified: "HTTP_METHOD_UNSPECIFIED",
    /**
     * HTTP POST
     */
    Post: "POST",
    /**
     * HTTP GET
     */
    Get: "GET",
    /**
     * HTTP HEAD
     */
    Head: "HEAD",
    /**
     * HTTP PUT
     */
    Put: "PUT",
    /**
     * HTTP DELETE
     */
    Delete: "DELETE",
    /**
     * HTTP PATCH
     */
    Patch: "PATCH",
    /**
     * HTTP OPTIONS
     */
    Options: "OPTIONS",
} as const;

/**
 * The HTTP method to use for the request. PATCH and OPTIONS are not permitted.
 */
export type AppEngineHttpTargetHttpMethod = (typeof AppEngineHttpTargetHttpMethod)[keyof typeof AppEngineHttpTargetHttpMethod];

export const HttpTargetHttpMethod = {
    /**
     * HTTP method unspecified. Defaults to POST.
     */
    HttpMethodUnspecified: "HTTP_METHOD_UNSPECIFIED",
    /**
     * HTTP POST
     */
    Post: "POST",
    /**
     * HTTP GET
     */
    Get: "GET",
    /**
     * HTTP HEAD
     */
    Head: "HEAD",
    /**
     * HTTP PUT
     */
    Put: "PUT",
    /**
     * HTTP DELETE
     */
    Delete: "DELETE",
    /**
     * HTTP PATCH
     */
    Patch: "PATCH",
    /**
     * HTTP OPTIONS
     */
    Options: "OPTIONS",
} as const;

/**
 * Which HTTP method to use for the request.
 */
export type HttpTargetHttpMethod = (typeof HttpTargetHttpMethod)[keyof typeof HttpTargetHttpMethod];
