// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.CloudResourceManager.V2Beta1
{
    public static class GetFolder
    {
        /// <summary>
        /// Retrieves a Folder identified by the supplied resource name. Valid Folder resource names have the format `folders/{folder_id}` (for example, `folders/1234`). The caller must have `resourcemanager.folders.get` permission on the identified folder.
        /// </summary>
        public static Task<GetFolderResult> InvokeAsync(GetFolderArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetFolderResult>("google-native:cloudresourcemanager/v2beta1:getFolder", args ?? new GetFolderArgs(), options.WithVersion());
    }


    public sealed class GetFolderArgs : Pulumi.InvokeArgs
    {
        [Input("folderId", required: true)]
        public string FolderId { get; set; } = null!;

        public GetFolderArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetFolderResult
    {
        /// <summary>
        /// Timestamp when the Folder was created. Assigned by the server.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// The folder's display name. A folder's display name must be unique amongst its siblings, e.g. no two folders with the same parent can share the same display name. The display name must start and end with a letter or digit, may contain letters, digits, spaces, hyphens and underscores and can be no longer than 30 characters. This is captured by the regular expression: `[\p{L}\p{N}]([\p{L}\p{N}_- ]{0,28}[\p{L}\p{N}])?`.
        /// </summary>
        public readonly string DisplayName;
        /// <summary>
        /// The lifecycle state of the folder. Updates to the lifecycle_state must be performed via DeleteFolder and UndeleteFolder.
        /// </summary>
        public readonly string LifecycleState;
        /// <summary>
        /// The resource name of the Folder. Its format is `folders/{folder_id}`, for example: "folders/1234".
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The Folder's parent's resource name. Updates to the folder's parent must be performed via MoveFolder.
        /// </summary>
        public readonly string Parent;

        [OutputConstructor]
        private GetFolderResult(
            string createTime,

            string displayName,

            string lifecycleState,

            string name,

            string parent)
        {
            CreateTime = createTime;
            DisplayName = displayName;
            LifecycleState = lifecycleState;
            Name = name;
            Parent = parent;
        }
    }
}
