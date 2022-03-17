// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.SQLAdmin.V1Beta4.Outputs
{

    /// <summary>
    /// Database instance local user password validation policy
    /// </summary>
    [OutputType]
    public sealed class PasswordValidationPolicyResponse
    {
        /// <summary>
        /// The complexity of the password.
        /// </summary>
        public readonly string Complexity;
        /// <summary>
        /// Disallow username as a part of the password.
        /// </summary>
        public readonly bool DisallowUsernameSubstring;
        /// <summary>
        /// Minimum number of characters allowed.
        /// </summary>
        public readonly int MinLength;
        /// <summary>
        /// Minimum interval after which the password can be changed. This flag is only supported for PostgresSQL.
        /// </summary>
        public readonly string PasswordChangeInterval;
        /// <summary>
        /// Number of previous passwords that cannot be reused.
        /// </summary>
        public readonly int ReuseInterval;

        [OutputConstructor]
        private PasswordValidationPolicyResponse(
            string complexity,

            bool disallowUsernameSubstring,

            int minLength,

            string passwordChangeInterval,

            int reuseInterval)
        {
            Complexity = complexity;
            DisallowUsernameSubstring = disallowUsernameSubstring;
            MinLength = minLength;
            PasswordChangeInterval = passwordChangeInterval;
            ReuseInterval = reuseInterval;
        }
    }
}
