// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dialogflow.V3Beta1.Outputs
{

    /// <summary>
    /// Represents a form parameter.
    /// </summary>
    [OutputType]
    public sealed class GoogleCloudDialogflowCxV3beta1FormParameterResponse
    {
        /// <summary>
        /// The default value of an optional parameter. If the parameter is required, the default value will be ignored.
        /// </summary>
        public readonly object DefaultValue;
        /// <summary>
        /// The human-readable name of the parameter, unique within the form.
        /// </summary>
        public readonly string DisplayName;
        /// <summary>
        /// The entity type of the parameter. Format: `projects/-/locations/-/agents/-/entityTypes/` for system entity types (for example, `projects/-/locations/-/agents/-/entityTypes/sys.date`), or `projects//locations//agents//entityTypes/` for developer entity types.
        /// </summary>
        public readonly string EntityType;
        /// <summary>
        /// Defines fill behavior for the parameter.
        /// </summary>
        public readonly Outputs.GoogleCloudDialogflowCxV3beta1FormParameterFillBehaviorResponse FillBehavior;
        /// <summary>
        /// Indicates whether the parameter represents a list of values.
        /// </summary>
        public readonly bool IsList;
        /// <summary>
        /// Indicates whether the parameter content should be redacted in log. If redaction is enabled, the parameter content will be replaced by parameter name during logging. Note: the parameter content is subject to redaction if either parameter level redaction or entity type level redaction is enabled.
        /// </summary>
        public readonly bool Redact;
        /// <summary>
        /// Indicates whether the parameter is required. Optional parameters will not trigger prompts; however, they are filled if the user specifies them. Required parameters must be filled before form filling concludes.
        /// </summary>
        public readonly bool Required;

        [OutputConstructor]
        private GoogleCloudDialogflowCxV3beta1FormParameterResponse(
            object defaultValue,

            string displayName,

            string entityType,

            Outputs.GoogleCloudDialogflowCxV3beta1FormParameterFillBehaviorResponse fillBehavior,

            bool isList,

            bool redact,

            bool required)
        {
            DefaultValue = defaultValue;
            DisplayName = displayName;
            EntityType = entityType;
            FillBehavior = fillBehavior;
            IsList = isList;
            Redact = redact;
            Required = required;
        }
    }
}
