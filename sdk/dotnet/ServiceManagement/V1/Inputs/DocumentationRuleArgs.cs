// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.ServiceManagement.V1.Inputs
{

    /// <summary>
    /// A documentation rule provides information about individual API elements.
    /// </summary>
    public sealed class DocumentationRuleArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Deprecation description of the selected element(s). It can be provided if an element is marked as `deprecated`.
        /// </summary>
        [Input("deprecationDescription")]
        public Input<string>? DeprecationDescription { get; set; }

        /// <summary>
        /// Description of the selected proto element (e.g. a message, a method, a 'service' definition, or a field). Defaults to leading &amp; trailing comments taken from the proto source definition of the proto element.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The selector is a comma-separated list of patterns for any element such as a method, a field, an enum value. Each pattern is a qualified name of the element which may end in "*", indicating a wildcard. Wildcards are only allowed at the end and for a whole component of the qualified name, i.e. "foo.*" is ok, but not "foo.b*" or "foo.*.bar". A wildcard will match one or more components. To specify a default for all applicable elements, the whole pattern "*" is used.
        /// </summary>
        [Input("selector")]
        public Input<string>? Selector { get; set; }

        public DocumentationRuleArgs()
        {
        }
    }
}
