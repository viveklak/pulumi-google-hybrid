// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Retrieves the specified document.
 */
export function getDocument(args: GetDocumentArgs, opts?: pulumi.InvokeOptions): Promise<GetDocumentResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("google-native:dialogflow/v2:getDocument", {
        "documentId": args.documentId,
        "knowledgeBaseId": args.knowledgeBaseId,
        "location": args.location,
        "project": args.project,
    }, opts);
}

export interface GetDocumentArgs {
    documentId: string;
    knowledgeBaseId: string;
    location: string;
    project?: string;
}

export interface GetDocumentResult {
    /**
     * The URI where the file content is located. For documents stored in Google Cloud Storage, these URIs must have the form `gs:///`. NOTE: External URLs must correspond to public webpages, i.e., they must be indexed by Google Search. In particular, URLs for showing documents in Google Cloud Storage (i.e. the URL in your browser) are not supported. Instead use the `gs://` format URI described above.
     */
    readonly contentUri: string;
    /**
     * The display name of the document. The name must be 1024 bytes or less; otherwise, the creation request fails.
     */
    readonly displayName: string;
    /**
     * Optional. If true, we try to automatically reload the document every day (at a time picked by the system). If false or unspecified, we don't try to automatically reload the document. Currently you can only enable automatic reload for documents sourced from a public url, see `source` field for the source types. Reload status can be tracked in `latest_reload_status`. If a reload fails, we will keep the document unchanged. If a reload fails with internal errors, the system will try to reload the document on the next day. If a reload fails with non-retriable errors (e.g. PERMISION_DENIED), the system will not try to reload the document anymore. You need to manually reload the document successfully by calling `ReloadDocument` and clear the errors.
     */
    readonly enableAutoReload: boolean;
    /**
     * The knowledge type of document content.
     */
    readonly knowledgeTypes: string[];
    /**
     * The time and status of the latest reload. This reload may have been triggered automatically or manually and may not have succeeded.
     */
    readonly latestReloadStatus: outputs.dialogflow.v2.GoogleCloudDialogflowV2DocumentReloadStatusResponse;
    /**
     * Optional. Metadata for the document. The metadata supports arbitrary key-value pairs. Suggested use cases include storing a document's title, an external URL distinct from the document's content_uri, etc. The max size of a `key` or a `value` of the metadata is 1024 bytes.
     */
    readonly metadata: {[key: string]: string};
    /**
     * The MIME type of this document.
     */
    readonly mimeType: string;
    /**
     * Optional. The document resource name. The name must be empty when creating a document. Format: `projects//locations//knowledgeBases//documents/`.
     */
    readonly name: string;
    /**
     * The raw content of the document. This field is only permitted for EXTRACTIVE_QA and FAQ knowledge types.
     */
    readonly rawContent: string;
    /**
     * The current state of the document.
     */
    readonly state: string;
}

export function getDocumentOutput(args: GetDocumentOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDocumentResult> {
    return pulumi.output(args).apply(a => getDocument(a, opts))
}

export interface GetDocumentOutputArgs {
    documentId: pulumi.Input<string>;
    knowledgeBaseId: pulumi.Input<string>;
    location: pulumi.Input<string>;
    project?: pulumi.Input<string>;
}
