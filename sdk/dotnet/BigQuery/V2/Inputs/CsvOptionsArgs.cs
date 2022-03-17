// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.BigQuery.V2.Inputs
{

    public sealed class CsvOptionsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// [Optional] Indicates if BigQuery should accept rows that are missing trailing optional columns. If true, BigQuery treats missing trailing columns as null values. If false, records with missing trailing columns are treated as bad records, and if there are too many bad records, an invalid error is returned in the job result. The default value is false.
        /// </summary>
        [Input("allowJaggedRows")]
        public Input<bool>? AllowJaggedRows { get; set; }

        /// <summary>
        /// [Optional] Indicates if BigQuery should allow quoted data sections that contain newline characters in a CSV file. The default value is false.
        /// </summary>
        [Input("allowQuotedNewlines")]
        public Input<bool>? AllowQuotedNewlines { get; set; }

        /// <summary>
        /// [Optional] The character encoding of the data. The supported values are UTF-8 or ISO-8859-1. The default value is UTF-8. BigQuery decodes the data after the raw, binary data has been split using the values of the quote and fieldDelimiter properties.
        /// </summary>
        [Input("encoding")]
        public Input<string>? Encoding { get; set; }

        /// <summary>
        /// [Optional] The separator for fields in a CSV file. BigQuery converts the string to ISO-8859-1 encoding, and then uses the first byte of the encoded string to split the data in its raw, binary state. BigQuery also supports the escape sequence "\t" to specify a tab separator. The default value is a comma (',').
        /// </summary>
        [Input("fieldDelimiter")]
        public Input<string>? FieldDelimiter { get; set; }

        /// <summary>
        /// [Optional] An custom string that will represent a NULL value in CSV import data.
        /// </summary>
        [Input("nullMarker")]
        public Input<string>? NullMarker { get; set; }

        /// <summary>
        /// [Optional] The value that is used to quote data sections in a CSV file. BigQuery converts the string to ISO-8859-1 encoding, and then uses the first byte of the encoded string to split the data in its raw, binary state. The default value is a double-quote ('"'). If your data does not contain quoted sections, set the property value to an empty string. If your data contains quoted newline characters, you must also set the allowQuotedNewlines property to true.
        /// </summary>
        [Input("quote")]
        public Input<string>? Quote { get; set; }

        /// <summary>
        /// [Optional] The number of rows at the top of a CSV file that BigQuery will skip when reading the data. The default value is 0. This property is useful if you have header rows in the file that should be skipped. When autodetect is on, the behavior is the following: * skipLeadingRows unspecified - Autodetect tries to detect headers in the first row. If they are not detected, the row is read as data. Otherwise data is read starting from the second row. * skipLeadingRows is 0 - Instructs autodetect that there are no headers and data should be read starting from the first row. * skipLeadingRows = N &gt; 0 - Autodetect skips N-1 rows and tries to detect headers in row N. If headers are not detected, row N is just skipped. Otherwise row N is used to extract column names for the detected schema.
        /// </summary>
        [Input("skipLeadingRows")]
        public Input<string>? SkipLeadingRows { get; set; }

        public CsvOptionsArgs()
        {
        }
    }
}
