// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***


export const ScanConfigTargetPlatformsItem = {
    /**
     * The target platform is unknown. Requests with this enum value will be rejected with INVALID_ARGUMENT error.
     */
    TargetPlatformUnspecified: "TARGET_PLATFORM_UNSPECIFIED",
    /**
     * Google App Engine service.
     */
    AppEngine: "APP_ENGINE",
    /**
     * Google Compute Engine service.
     */
    Compute: "COMPUTE",
} as const;

export type ScanConfigTargetPlatformsItem = (typeof ScanConfigTargetPlatformsItem)[keyof typeof ScanConfigTargetPlatformsItem];

export const ScanConfigUserAgent = {
    /**
     * The user agent is unknown. Service will default to CHROME_LINUX.
     */
    UserAgentUnspecified: "USER_AGENT_UNSPECIFIED",
    /**
     * Chrome on Linux. This is the service default if unspecified.
     */
    ChromeLinux: "CHROME_LINUX",
    /**
     * Chrome on Android.
     */
    ChromeAndroid: "CHROME_ANDROID",
    /**
     * Safari on IPhone.
     */
    SafariIphone: "SAFARI_IPHONE",
} as const;

/**
 * The user agent used during scanning.
 */
export type ScanConfigUserAgent = (typeof ScanConfigUserAgent)[keyof typeof ScanConfigUserAgent];

export const ScanRunExecutionState = {
    /**
     * Represents an invalid state caused by internal server error. This value should never be returned.
     */
    ExecutionStateUnspecified: "EXECUTION_STATE_UNSPECIFIED",
    /**
     * The scan is waiting in the queue.
     */
    Queued: "QUEUED",
    /**
     * The scan is in progress.
     */
    Scanning: "SCANNING",
    /**
     * The scan is either finished or stopped by user.
     */
    Finished: "FINISHED",
} as const;

/**
 * The execution state of the ScanRun.
 */
export type ScanRunExecutionState = (typeof ScanRunExecutionState)[keyof typeof ScanRunExecutionState];

export const ScanRunResultState = {
    /**
     * Default value. This value is returned when the ScanRun is not yet finished.
     */
    ResultStateUnspecified: "RESULT_STATE_UNSPECIFIED",
    /**
     * The scan finished without errors.
     */
    Success: "SUCCESS",
    /**
     * The scan finished with errors.
     */
    Error: "ERROR",
    /**
     * The scan was terminated by user.
     */
    Killed: "KILLED",
} as const;

/**
 * The result state of the ScanRun. This field is only available after the execution state reaches "FINISHED".
 */
export type ScanRunResultState = (typeof ScanRunResultState)[keyof typeof ScanRunResultState];
