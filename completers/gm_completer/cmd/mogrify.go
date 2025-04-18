package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/gm_completer/cmd/action"
	"github.com/spf13/cobra"
)

var mogrifyCmd = &cobra.Command{
	Use:   "mogrify",
	Short: "transform an image or sequence of images",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mogrifyCmd).Standalone()

	mogrifyCmd.Flags().StringSliceS("affine", "affine", []string{}, "affine transform matrix")
	mogrifyCmd.Flags().CountS("antialias", "antialias", "remove pixel-aliasing")
	mogrifyCmd.Flags().StringSliceS("asc-cdl", "asc-cdl", []string{}, "apply ASC CDL transform")
	mogrifyCmd.Flags().StringSliceS("authenticate", "authenticate", []string{}, "decrypt image with this password")
	mogrifyCmd.Flags().CountS("auto-orient", "auto-orient", "orient (rotate) image so it is upright")
	mogrifyCmd.Flags().StringSliceS("background", "background", []string{}, "background color")
	mogrifyCmd.Flags().StringSliceS("black-threshold", "black-threshold", []string{}, "pixels below the threshold become black")
	mogrifyCmd.Flags().StringSliceS("blue-primary", "blue-primary", []string{}, "chomaticity blue primary point")
	mogrifyCmd.Flags().StringSliceS("blur", "blur", []string{}, "blur the image")
	mogrifyCmd.Flags().StringSliceS("border", "border", []string{}, "surround image with a border of color")
	mogrifyCmd.Flags().StringSliceS("bordercolor", "bordercolor", []string{}, "border color")
	mogrifyCmd.Flags().StringSliceS("box", "box", []string{}, "set the color of the annotation bounding box")
	mogrifyCmd.Flags().StringSliceS("channel", "channel", []string{}, "extract a particular color channel from image")
	mogrifyCmd.Flags().StringSliceS("charcoal", "charcoal", []string{}, "simulate a charcoal drawing")
	mogrifyCmd.Flags().StringSliceS("chop", "chop", []string{}, "remove pixels from the image interior")
	mogrifyCmd.Flags().StringSliceS("colorize", "colorize", []string{}, "colorize the image with the fill color")
	mogrifyCmd.Flags().CountS("colors", "colors", "preferred number of colors in the image")
	mogrifyCmd.Flags().StringSliceS("colorspace", "colorspace", []string{}, "alternate image colorspace")
	mogrifyCmd.Flags().StringSliceS("comment", "comment", []string{}, "annotate image with comment")
	mogrifyCmd.Flags().StringSliceS("compose", "compose", []string{}, "composite operator")
	mogrifyCmd.Flags().StringSliceS("compress", "compress", []string{}, "image compression type")
	mogrifyCmd.Flags().CountS("contrast", "contrast", "enhance or reduce the image contrast")
	mogrifyCmd.Flags().StringSliceS("convolve", "convolve", []string{}, "convolve image with the specified convolution kernel")
	mogrifyCmd.Flags().CountS("create-directories", "create-directories", "create output directories if required")
	mogrifyCmd.Flags().StringSliceS("crop", "crop", []string{}, "preferred size and location of the cropped image")
	mogrifyCmd.Flags().StringSliceS("cycle", "cycle", []string{}, "cycle the image colormap")
	mogrifyCmd.Flags().StringSliceS("debug", "debug", []string{}, "display copious debugging information")
	mogrifyCmd.Flags().StringSliceS("define", "define", []string{}, "Coder/decoder specific options")
	mogrifyCmd.Flags().StringSliceS("delay", "delay", []string{}, "display the next image after pausing")
	mogrifyCmd.Flags().StringSliceS("density", "density", []string{}, "horizontal and vertical density of the image")
	mogrifyCmd.Flags().StringSliceS("depth", "depth", []string{}, "image depth")
	mogrifyCmd.Flags().CountS("despeckle", "despeckle", "reduce the speckles within an image")
	mogrifyCmd.Flags().StringSliceS("display", "display", []string{}, "get image or font from this X server")
	mogrifyCmd.Flags().StringSliceS("dispose", "dispose", []string{}, "Undefined, None, Background, Previous")
	mogrifyCmd.Flags().CountS("dither", "dither", "apply Floyd/Steinberg error diffusion to image")
	mogrifyCmd.Flags().StringSliceS("draw", "draw", []string{}, "annotate the image with a graphic primitive")
	mogrifyCmd.Flags().StringSliceS("edge", "edge", []string{}, "apply a filter to detect edges in the image")
	mogrifyCmd.Flags().StringSliceS("emboss", "emboss", []string{}, "emboss an image")
	mogrifyCmd.Flags().StringSliceS("encoding", "encoding", []string{}, "text encoding type")
	mogrifyCmd.Flags().StringSliceS("endian", "endian", []string{}, "multibyte word order (LSB, MSB, or Native)")
	mogrifyCmd.Flags().CountS("enhance", "enhance", "apply a digital filter to enhance a noisy image")
	mogrifyCmd.Flags().StringSliceS("equalize", "equalize", []string{}, "perform histogram equalization to an image")
	mogrifyCmd.Flags().CountS("extent", "extent", "composite image on background color canvas image")
	mogrifyCmd.Flags().StringSliceS("fill", "fill", []string{}, "color to use when filling a graphic primitive")
	mogrifyCmd.Flags().StringSliceS("filter", "filter", []string{}, "use this filter when resizing an image")
	mogrifyCmd.Flags().CountS("flip", "flip", "flip image in the vertical direction")
	mogrifyCmd.Flags().CountS("flop", "flop", "flop image in the horizontal direction")
	mogrifyCmd.Flags().StringSliceS("font", "font", []string{}, "render text with this font")
	mogrifyCmd.Flags().StringSliceS("format", "format", []string{}, "image format type")
	mogrifyCmd.Flags().StringSliceS("frame", "frame", []string{}, "surround image with an ornamental border")
	mogrifyCmd.Flags().StringSliceS("fuzz", "fuzz", []string{}, "colors within this distance are considered equal")
	mogrifyCmd.Flags().StringSliceS("gamma", "gamma", []string{}, "level of gamma correction")
	mogrifyCmd.Flags().StringSliceS("gaussian", "gaussian", []string{}, "gaussian blur an image")
	mogrifyCmd.Flags().StringSliceS("geometry", "geometry", []string{}, "preferred size or location of the image")
	mogrifyCmd.Flags().StringSliceS("gravity", "gravity", []string{}, "horizontal and vertical text/object placement")
	mogrifyCmd.Flags().StringSliceS("green-primary", "green-primary", []string{}, "chomaticity green primary point")
	mogrifyCmd.Flags().StringSliceS("hald-clut", "hald-clut", []string{}, "apply a Hald CLUT to the image")
	mogrifyCmd.Flags().BoolS("help", "help", false, "print program options")
	mogrifyCmd.Flags().StringSliceS("implode", "implode", []string{}, "implode image pixels about the center")
	mogrifyCmd.Flags().StringSliceS("interlace", "interlace", []string{}, "None, Line, Plane, or Partition")
	mogrifyCmd.Flags().StringSliceS("label", "labellabel", []string{}, "assign a label to an image")
	mogrifyCmd.Flags().StringSliceS("lat", "lat", []string{}, "local adaptive thresholding")
	mogrifyCmd.Flags().StringSliceS("level", "level", []string{}, "adjust the level of image contrast")
	mogrifyCmd.Flags().StringSliceS("limit", "limit", []string{}, "Disk, File, Map, Memory, Pixels, Width, Height or Threads resource limit")
	mogrifyCmd.Flags().StringSliceS("linewidth", "linewidth", []string{}, "the line width for subsequent draw operations")
	mogrifyCmd.Flags().StringSliceS("list", "list", []string{}, "Color, Delegate, Format, Magic, Module, Resource, or Type")
	mogrifyCmd.Flags().StringSliceS("log", "log", []string{}, "format of debugging information")
	mogrifyCmd.Flags().StringSliceS("loop", "loop", []string{}, "add Netscape loop extension to your GIF animation")
	mogrifyCmd.Flags().CountS("magnify", "magnify", "interpolate image to double size")
	mogrifyCmd.Flags().StringSliceS("map", "map", []string{}, "transform image colors to match this set of colors")
	mogrifyCmd.Flags().StringSliceS("mask", "mask", []string{}, "set the image clip mask")
	mogrifyCmd.Flags().CountS("matte", "matte", "store matte channel if the image has one")
	mogrifyCmd.Flags().StringSliceS("mattecolor", "mattecolor", []string{}, "specify the color to be used with the -frame option")
	mogrifyCmd.Flags().StringSliceS("median", "median", []string{}, "apply a median filter to the image")
	mogrifyCmd.Flags().CountS("minify", "minify", "interpolate the image to half size")
	mogrifyCmd.Flags().StringSliceS("modulate", "modulate", []string{}, "vary the brightness, saturation, and hue")
	mogrifyCmd.Flags().CountS("monitor", "monitor", "show progress indication")
	mogrifyCmd.Flags().CountS("monochrome", "monochrome", "transform image to black and white")
	mogrifyCmd.Flags().StringSliceS("motion-blur", "motion-blur", []string{}, "simulate motion blur")
	mogrifyCmd.Flags().CountS("negate", "negate", "replace every pixel with its complementary color")
	mogrifyCmd.Flags().StringSliceS("noise", "noise", []string{}, "add or reduce noise in an image")
	mogrifyCmd.Flags().CountS("noop", "noop", "do not apply options to image")
	mogrifyCmd.Flags().CountS("normalize", "normalize", "transform image to span the full range of colors")
	mogrifyCmd.Flags().StringSliceS("opaque", "opaque", []string{}, "change this color to the fill color")
	mogrifyCmd.Flags().StringSliceS("operator", "operator", []string{}, "apply a mathematical or bitwise operator to channel")
	mogrifyCmd.Flags().StringSliceS("ordered-dither", "ordered-dither", []string{}, "ordered dither the image")
	mogrifyCmd.Flags().StringSliceS("orient", "orient", []string{}, "set image orientation attribute")
	mogrifyCmd.Flags().StringSliceS("output-directory", "output-directory", []string{}, "write output files to directory")
	mogrifyCmd.Flags().StringSliceS("page", "page", []string{}, "size and location of an image canvas")
	mogrifyCmd.Flags().StringSliceS("paint", "paint", []string{}, "simulate an oil painting")
	mogrifyCmd.Flags().StringSliceS("pointsize", "pointsize", []string{}, "font point size")
	mogrifyCmd.Flags().CountS("preserve-timestamp", "preserve-timestamp", "preserve original timestamps of the file")
	mogrifyCmd.Flags().StringSliceS("profile", "profile", []string{}, "add ICM or IPTC information profile to image")
	mogrifyCmd.Flags().StringSliceS("quality", "quality", []string{}, "JPEG/MIFF/PNG compression level")
	mogrifyCmd.Flags().StringSliceS("raise", "raise", []string{}, "lighten/darken image edges to create a 3-D effect")
	mogrifyCmd.Flags().StringSliceS("random-threshold", "random-threshold", []string{}, "random threshold the image")
	mogrifyCmd.Flags().StringSliceS("recolor", "recolor", []string{}, "apply a color translation matrix to image channels")
	mogrifyCmd.Flags().StringSliceS("red-primary", "red-primary", []string{}, "chomaticity red primary point")
	mogrifyCmd.Flags().StringSliceS("region", "region", []string{}, "apply options to a portion of the image")
	mogrifyCmd.Flags().CountS("render", "render", "render vector graphics")
	mogrifyCmd.Flags().StringSliceS("repage", "repage", []string{}, "adjust current page offsets by geometry")
	mogrifyCmd.Flags().StringSliceS("resample", "resample", []string{}, "resample to horizontal and vertical resolution")
	mogrifyCmd.Flags().StringSliceS("resize", "resize", []string{}, "preferred size or location of the image")
	mogrifyCmd.Flags().StringSliceS("roll", "roll", []string{}, "roll an image vertically or horizontally")
	mogrifyCmd.Flags().StringSliceS("rotate", "rotate", []string{}, "apply Paeth rotation to the image")
	mogrifyCmd.Flags().StringSliceS("sample", "sample", []string{}, "scale image with pixel sampling")
	mogrifyCmd.Flags().StringSliceS("sampling-factor", "sampling-factor", []string{}, "horizontal and vertical sampling factors")
	mogrifyCmd.Flags().StringSliceS("scale", "scale", []string{}, "scale the image")
	mogrifyCmd.Flags().StringSliceS("scene", "scene", []string{}, "image scene number")
	mogrifyCmd.Flags().StringSliceS("seed", "seed", []string{}, "pseudo-random number generator seed value")
	mogrifyCmd.Flags().StringSliceS("segment", "segment", []string{}, "segment an image")
	mogrifyCmd.Flags().StringSliceS("set", "set", []string{}, "set image attribute")
	mogrifyCmd.Flags().StringSliceS("shade", "shade", []string{}, "shade the image using a distant light source")
	mogrifyCmd.Flags().StringSliceS("sharpen", "sharpen", []string{}, "sharpen the image")
	mogrifyCmd.Flags().StringSliceS("shave", "shave", []string{}, "shave pixels from the image edges")
	mogrifyCmd.Flags().StringSliceS("shear", "shear", []string{}, "slide one edge of the image along the X or Y axis")
	mogrifyCmd.Flags().StringSliceS("size", "size", []string{}, "width and height of image")
	mogrifyCmd.Flags().StringSliceS("solarize", "solarizesolarize", []string{}, "negate all pixels above the threshold level")
	mogrifyCmd.Flags().StringSliceS("spread", "spreadspread", []string{}, "displace image pixels by a random amount")
	mogrifyCmd.Flags().CountS("strip", "stripstrip", "strip all profiles and text attributes from image")
	mogrifyCmd.Flags().StringSliceS("stroke", "strokestroke", []string{}, "graphic primitive stroke color")
	mogrifyCmd.Flags().StringSliceS("strokewidth", "strokewidthstrokewidth", []string{}, "graphic primitive stroke width")
	mogrifyCmd.Flags().StringSliceS("swirl", "swirl", []string{}, "swirl image pixels about the center")
	mogrifyCmd.Flags().StringSliceS("texture", "texture", []string{}, "name of texture to tile onto the image background")
	mogrifyCmd.Flags().StringSliceS("threshold", "threshold", []string{}, "threshold the image")
	mogrifyCmd.Flags().StringSliceS("thumbnail", "thumbnail", []string{}, "resize the image (optimized for thumbnails)")
	mogrifyCmd.Flags().StringSliceS("tile", "tile", []string{}, "tile image when filling a graphic primitive")
	mogrifyCmd.Flags().CountS("transform", "transform", "affine transform image")
	mogrifyCmd.Flags().StringSliceS("transparent", "transparent", []string{}, "make this color transparent within the image")
	mogrifyCmd.Flags().StringSliceS("treedepth", "treedepth", []string{}, "color tree depth")
	mogrifyCmd.Flags().CountS("trim", "trim", "trim image edges")
	mogrifyCmd.Flags().StringSliceS("type", "type", []string{}, "image type")
	mogrifyCmd.Flags().StringSliceS("undercolor", "undercolor", []string{}, "annotation bounding box color")
	mogrifyCmd.Flags().StringSliceS("units", "units", []string{}, "PixelsPerInch, PixelsPerCentimeter, or Undefined")
	mogrifyCmd.Flags().StringSliceS("unsharp", "unsharp", []string{}, "sharpen the image")
	mogrifyCmd.Flags().BoolS("verbose", "verbose", false, "print detailed information about the image")
	mogrifyCmd.Flags().BoolS("version", "version", false, "print version information")
	mogrifyCmd.Flags().CountS("view", "view", "FlashPix viewing transforms")
	mogrifyCmd.Flags().StringSliceS("virtual-pixel", "virtual-pixel", []string{}, "Constant, Edge, Mirror, or Tile")
	mogrifyCmd.Flags().StringSliceS("wave", "wave", []string{}, "alter an image along a sine wave")
	mogrifyCmd.Flags().StringSliceS("white-point", "white-point", []string{}, "chomaticity white point")
	mogrifyCmd.Flags().StringSliceS("white-threshold", "white-threshold", []string{}, "pixels above the threshold become white")
	rootCmd.AddCommand(mogrifyCmd)

	carapace.Gen(mogrifyCmd).FlagCompletion(carapace.ActionMap{
		"fill": action.ActionColor(),
	})

	carapace.Gen(mogrifyCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
