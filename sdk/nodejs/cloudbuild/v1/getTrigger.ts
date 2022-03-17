// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Returns information about a `BuildTrigger`. This API is experimental.
 */
export function getTrigger(args: GetTriggerArgs, opts?: pulumi.InvokeOptions): Promise<GetTriggerResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("google-native:cloudbuild/v1:getTrigger", {
        "location": args.location,
        "project": args.project,
        "projectId": args.projectId,
        "triggerId": args.triggerId,
    }, opts);
}

export interface GetTriggerArgs {
    location: string;
    project?: string;
    projectId: string;
    triggerId: string;
}

export interface GetTriggerResult {
    /**
     * Configuration for manual approval to start a build invocation of this BuildTrigger.
     */
    readonly approvalConfig: outputs.cloudbuild.v1.ApprovalConfigResponse;
    /**
     * Autodetect build configuration. The following precedence is used (case insensitive): 1. cloudbuild.yaml 2. cloudbuild.yml 3. cloudbuild.json 4. Dockerfile Currently only available for GitHub App Triggers.
     */
    readonly autodetect: boolean;
    /**
     * BitbucketServerTriggerConfig describes the configuration of a trigger that creates a build whenever a Bitbucket Server event is received.
     */
    readonly bitbucketServerTriggerConfig: outputs.cloudbuild.v1.BitbucketServerTriggerConfigResponse;
    /**
     * Contents of the build template.
     */
    readonly build: outputs.cloudbuild.v1.BuildResponse;
    /**
     * Time when the trigger was created.
     */
    readonly createTime: string;
    /**
     * Human-readable description of this trigger.
     */
    readonly description: string;
    /**
     * If true, the trigger will never automatically execute a build.
     */
    readonly disabled: boolean;
    /**
     * EventType allows the user to explicitly set the type of event to which this BuildTrigger should respond. This field will be validated against the rest of the configuration if it is set.
     */
    readonly eventType: string;
    /**
     * Path, from the source root, to the build configuration file (i.e. cloudbuild.yaml).
     */
    readonly filename: string;
    /**
     * A Common Expression Language string.
     */
    readonly filter: string;
    /**
     * The file source describing the local or remote Build template.
     */
    readonly gitFileSource: outputs.cloudbuild.v1.GitFileSourceResponse;
    /**
     * GitHubEventsConfig describes the configuration of a trigger that creates a build whenever a GitHub event is received. Mutually exclusive with `trigger_template`.
     */
    readonly github: outputs.cloudbuild.v1.GitHubEventsConfigResponse;
    /**
     * ignored_files and included_files are file glob matches using https://golang.org/pkg/path/filepath/#Match extended with support for "**". If ignored_files and changed files are both empty, then they are not used to determine whether or not to trigger a build. If ignored_files is not empty, then we ignore any files that match any of the ignored_file globs. If the change has no files that are outside of the ignored_files globs, then we do not trigger a build.
     */
    readonly ignoredFiles: string[];
    /**
     * If any of the files altered in the commit pass the ignored_files filter and included_files is empty, then as far as this filter is concerned, we should trigger the build. If any of the files altered in the commit pass the ignored_files filter and included_files is not empty, then we make sure that at least one of those files matches a included_files glob. If not, then we do not trigger a build.
     */
    readonly includedFiles: string[];
    /**
     * User-assigned name of the trigger. Must be unique within the project. Trigger names must meet the following requirements: + They must contain only alphanumeric characters and dashes. + They can be 1-64 characters long. + They must begin and end with an alphanumeric character.
     */
    readonly name: string;
    /**
     * PubsubConfig describes the configuration of a trigger that creates a build whenever a Pub/Sub message is published.
     */
    readonly pubsubConfig: outputs.cloudbuild.v1.PubsubConfigResponse;
    /**
     * The `Trigger` name with format: `projects/{project}/locations/{location}/triggers/{trigger}`, where {trigger} is a unique identifier generated by the service.
     */
    readonly resourceName: string;
    /**
     * The service account used for all user-controlled operations including UpdateBuildTrigger, RunBuildTrigger, CreateBuild, and CancelBuild. If no service account is set, then the standard Cloud Build service account ([PROJECT_NUM]@system.gserviceaccount.com) will be used instead. Format: `projects/{PROJECT_ID}/serviceAccounts/{ACCOUNT_ID_OR_EMAIL}`
     */
    readonly serviceAccount: string;
    /**
     * The repo and ref of the repository from which to build. This field is used only for those triggers that do not respond to SCM events. Triggers that respond to such events build source at whatever commit caused the event. This field is currently only used by Webhook, Pub/Sub, Manual, and Cron triggers.
     */
    readonly sourceToBuild: outputs.cloudbuild.v1.GitRepoSourceResponse;
    /**
     * Substitutions for Build resource. The keys must match the following regular expression: `^_[A-Z0-9_]+$`.
     */
    readonly substitutions: {[key: string]: string};
    /**
     * Tags for annotation of a `BuildTrigger`
     */
    readonly tags: string[];
    /**
     * Template describing the types of source changes to trigger a build. Branch and tag names in trigger templates are interpreted as regular expressions. Any branch or tag change that matches that regular expression will trigger a build. Mutually exclusive with `github`.
     */
    readonly triggerTemplate: outputs.cloudbuild.v1.RepoSourceResponse;
    /**
     * WebhookConfig describes the configuration of a trigger that creates a build whenever a webhook is sent to a trigger's webhook URL.
     */
    readonly webhookConfig: outputs.cloudbuild.v1.WebhookConfigResponse;
}

export function getTriggerOutput(args: GetTriggerOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetTriggerResult> {
    return pulumi.output(args).apply(a => getTrigger(a, opts))
}

export interface GetTriggerOutputArgs {
    location: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    projectId: pulumi.Input<string>;
    triggerId: pulumi.Input<string>;
}
