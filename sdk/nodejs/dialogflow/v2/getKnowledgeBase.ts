// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Retrieves the specified knowledge base.
 */
export function getKnowledgeBase(args: GetKnowledgeBaseArgs, opts?: pulumi.InvokeOptions): Promise<GetKnowledgeBaseResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("google-native:dialogflow/v2:getKnowledgeBase", {
        "knowledgeBaseId": args.knowledgeBaseId,
        "location": args.location,
        "project": args.project,
    }, opts);
}

export interface GetKnowledgeBaseArgs {
    knowledgeBaseId: string;
    location: string;
    project: string;
}

export interface GetKnowledgeBaseResult {
    /**
     * The display name of the knowledge base. The name must be 1024 bytes or less; otherwise, the creation request fails.
     */
    readonly displayName: string;
    /**
     * Language which represents the KnowledgeBase. When the KnowledgeBase is created/updated, expect this to be present for non en-us languages. When unspecified, the default language code en-us applies.
     */
    readonly languageCode: string;
    /**
     * The knowledge base resource name. The name must be empty when creating a knowledge base. Format: `projects//locations//knowledgeBases/`.
     */
    readonly name: string;
}
