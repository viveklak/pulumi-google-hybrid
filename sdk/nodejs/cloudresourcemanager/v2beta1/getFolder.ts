// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Retrieves a Folder identified by the supplied resource name. Valid Folder resource names have the format `folders/{folder_id}` (for example, `folders/1234`). The caller must have `resourcemanager.folders.get` permission on the identified folder.
 */
export function getFolder(args: GetFolderArgs, opts?: pulumi.InvokeOptions): Promise<GetFolderResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("google-native:cloudresourcemanager/v2beta1:getFolder", {
        "folderId": args.folderId,
    }, opts);
}

export interface GetFolderArgs {
    folderId: string;
}

export interface GetFolderResult {
    /**
     * Timestamp when the Folder was created. Assigned by the server.
     */
    readonly createTime: string;
    /**
     * The folder's display name. A folder's display name must be unique amongst its siblings, e.g. no two folders with the same parent can share the same display name. The display name must start and end with a letter or digit, may contain letters, digits, spaces, hyphens and underscores and can be no longer than 30 characters. This is captured by the regular expression: `[\p{L}\p{N}]([\p{L}\p{N}_- ]{0,28}[\p{L}\p{N}])?`.
     */
    readonly displayName: string;
    /**
     * The lifecycle state of the folder. Updates to the lifecycle_state must be performed via DeleteFolder and UndeleteFolder.
     */
    readonly lifecycleState: string;
    /**
     * The resource name of the Folder. Its format is `folders/{folder_id}`, for example: "folders/1234".
     */
    readonly name: string;
    /**
     * The Folder's parent's resource name. Updates to the folder's parent must be performed via MoveFolder.
     */
    readonly parent: string;
}
