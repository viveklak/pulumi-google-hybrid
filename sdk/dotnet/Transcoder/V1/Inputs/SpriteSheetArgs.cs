// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Transcoder.V1.Inputs
{

    /// <summary>
    /// Sprite sheet configuration.
    /// </summary>
    public sealed class SpriteSheetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The maximum number of sprites per row in a sprite sheet. The default is 0, which indicates no maximum limit.
        /// </summary>
        [Input("columnCount")]
        public Input<int>? ColumnCount { get; set; }

        /// <summary>
        /// End time in seconds, relative to the output file timeline. When `end_time_offset` is not specified, the sprites are generated until the end of the output file.
        /// </summary>
        [Input("endTimeOffset")]
        public Input<string>? EndTimeOffset { get; set; }

        /// <summary>
        /// File name prefix for the generated sprite sheets. Each sprite sheet has an incremental 10-digit zero-padded suffix starting from 0 before the extension, such as `sprite_sheet0000000123.jpeg`.
        /// </summary>
        [Input("filePrefix", required: true)]
        public Input<string> FilePrefix { get; set; } = null!;

        /// <summary>
        /// Format type. The default is `jpeg`. Supported formats: - `jpeg`
        /// </summary>
        [Input("format")]
        public Input<string>? Format { get; set; }

        /// <summary>
        /// Starting from `0s`, create sprites at regular intervals. Specify the interval value in seconds.
        /// </summary>
        [Input("interval")]
        public Input<string>? Interval { get; set; }

        /// <summary>
        /// The quality of the generated sprite sheet. Enter a value between 1 and 100, where 1 is the lowest quality and 100 is the highest quality. The default is 100. A high quality value corresponds to a low image data compression ratio.
        /// </summary>
        [Input("quality")]
        public Input<int>? Quality { get; set; }

        /// <summary>
        /// The maximum number of rows per sprite sheet. When the sprite sheet is full, a new sprite sheet is created. The default is 0, which indicates no maximum limit.
        /// </summary>
        [Input("rowCount")]
        public Input<int>? RowCount { get; set; }

        /// <summary>
        /// The height of sprite in pixels. Must be an even integer. To preserve the source aspect ratio, set the SpriteSheet.sprite_height_pixels field or the SpriteSheet.sprite_width_pixels field, but not both (the API will automatically calculate the missing field).
        /// </summary>
        [Input("spriteHeightPixels", required: true)]
        public Input<int> SpriteHeightPixels { get; set; } = null!;

        /// <summary>
        /// The width of sprite in pixels. Must be an even integer. To preserve the source aspect ratio, set the SpriteSheet.sprite_width_pixels field or the SpriteSheet.sprite_height_pixels field, but not both (the API will automatically calculate the missing field).
        /// </summary>
        [Input("spriteWidthPixels", required: true)]
        public Input<int> SpriteWidthPixels { get; set; } = null!;

        /// <summary>
        /// Start time in seconds, relative to the output file timeline. Determines the first sprite to pick. The default is `0s`.
        /// </summary>
        [Input("startTimeOffset")]
        public Input<string>? StartTimeOffset { get; set; }

        /// <summary>
        /// Total number of sprites. Create the specified number of sprites distributed evenly across the timeline of the output media. The default is 100.
        /// </summary>
        [Input("totalCount")]
        public Input<int>? TotalCount { get; set; }

        public SpriteSheetArgs()
        {
        }
    }
}
