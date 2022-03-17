// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Storage.V1.Outputs
{

    /// <summary>
    /// The project team associated with the entity, if any.
    /// </summary>
    [OutputType]
    public sealed class ObjectAccessControlProjectTeamResponse
    {
        /// <summary>
        /// The project number.
        /// </summary>
        public readonly string ProjectNumber;
        /// <summary>
        /// The team.
        /// </summary>
        public readonly string Team;

        [OutputConstructor]
        private ObjectAccessControlProjectTeamResponse(
            string projectNumber,

            string team)
        {
            ProjectNumber = projectNumber;
            Team = team;
        }
    }
}
