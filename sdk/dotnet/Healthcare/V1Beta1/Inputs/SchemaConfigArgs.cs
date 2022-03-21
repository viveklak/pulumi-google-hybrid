// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Healthcare.V1Beta1.Inputs
{

    /// <summary>
    /// Configuration for the FHIR BigQuery schema. Determines how the server generates the schema.
    /// </summary>
    public sealed class SchemaConfigArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The depth for all recursive structures in the output analytics schema. For example, `concept` in the CodeSystem resource is a recursive structure; when the depth is 2, the CodeSystem table will have a column called `concept.concept` but not `concept.concept.concept`. If not specified or set to 0, the server will use the default value 2. The maximum depth allowed is 5.
        /// </summary>
        [Input("recursiveStructureDepth")]
        public Input<string>? RecursiveStructureDepth { get; set; }

        /// <summary>
        /// Specifies the output schema type. Schema type is required.
        /// </summary>
        [Input("schemaType")]
        public Input<Pulumi.GoogleNative.Healthcare.V1Beta1.SchemaConfigSchemaType>? SchemaType { get; set; }

        public SchemaConfigArgs()
        {
        }
    }
}
