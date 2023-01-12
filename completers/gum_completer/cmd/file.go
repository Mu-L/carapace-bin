package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/gum"
	"github.com/spf13/cobra"
)

var fileCmd = &cobra.Command{
	Use:   "file",
	Short: "Pick a file from a folder",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fileCmd).Standalone()
	fileCmd.Flags().BoolP("all", "a", false, "Show hidden and 'dot' files")
	fileCmd.Flags().StringP("cursor", "c", "", "The cursor character")
	fileCmd.Flags().String("cursor.align", "", "Text Alignment")
	fileCmd.Flags().String("cursor.background", "", "Background Color")
	fileCmd.Flags().Bool("cursor.bold", false, "Bold text")
	fileCmd.Flags().String("cursor.border", "", "Border Style")
	fileCmd.Flags().String("cursor.border-background", "", "Border Background Color")
	fileCmd.Flags().String("cursor.border-foreground", "", "Border Foreground Color")
	fileCmd.Flags().Bool("cursor.faint", false, "Faint text")
	fileCmd.Flags().String("cursor.foreground", "", "Foreground Color")
	fileCmd.Flags().String("cursor.height", "", "Text height")
	fileCmd.Flags().Bool("cursor.italic", false, "Italicize text")
	fileCmd.Flags().String("cursor.margin", "", "Text margin")
	fileCmd.Flags().String("cursor.padding", "", "Text padding")
	fileCmd.Flags().Bool("cursor.strikethrough", false, "Strikethrough text")
	fileCmd.Flags().Bool("cursor.underline", false, "Underline text")
	fileCmd.Flags().String("cursor.width", "", "Text width")
	fileCmd.Flags().Bool("directory", false, "Allow directories selection")
	fileCmd.Flags().String("directory.align", "", "Text Alignment")
	fileCmd.Flags().String("directory.background", "", "Background Color")
	fileCmd.Flags().Bool("directory.bold", false, "Bold text")
	fileCmd.Flags().String("directory.border", "", "Border Style")
	fileCmd.Flags().String("directory.border-background", "", "Border Background Color")
	fileCmd.Flags().String("directory.border-foreground", "", "Border Foreground Color")
	fileCmd.Flags().Bool("directory.faint", false, "Faint text")
	fileCmd.Flags().String("directory.foreground", "", "Foreground Color")
	fileCmd.Flags().String("directory.height", "", "Text height")
	fileCmd.Flags().Bool("directory.italic", false, "Italicize text")
	fileCmd.Flags().String("directory.margin", "", "Text margin")
	fileCmd.Flags().String("directory.padding", "", "Text padding")
	fileCmd.Flags().Bool("directory.strikethrough", false, "Strikethrough text")
	fileCmd.Flags().Bool("directory.underline", false, "Underline text")
	fileCmd.Flags().String("directory.width", "", "Text width")
	fileCmd.Flags().Bool("file", false, "Allow files selection")
	fileCmd.Flags().String("file-size.align", "", "Text Alignment")
	fileCmd.Flags().String("file-size.background", "", "Background Color")
	fileCmd.Flags().Bool("file-size.bold", false, "Bold text")
	fileCmd.Flags().String("file-size.border", "", "Border Style")
	fileCmd.Flags().String("file-size.border-background", "", "Border Background Color")
	fileCmd.Flags().String("file-size.border-foreground", "", "Border Foreground Color")
	fileCmd.Flags().Bool("file-size.faint", false, "Faint text")
	fileCmd.Flags().String("file-size.foreground", "", "Foreground Color")
	fileCmd.Flags().String("file-size.height", "", "Text height")
	fileCmd.Flags().Bool("file-size.italic", false, "Italicize text")
	fileCmd.Flags().String("file-size.margin", "", "Text margin")
	fileCmd.Flags().String("file-size.padding", "", "Text padding")
	fileCmd.Flags().Bool("file-size.strikethrough", false, "Strikethrough text")
	fileCmd.Flags().Bool("file-size.underline", false, "Underline text")
	fileCmd.Flags().String("file-size.width", "", "Text width")
	fileCmd.Flags().String("file.align", "", "Text Alignment")
	fileCmd.Flags().String("file.background", "", "Background Color")
	fileCmd.Flags().Bool("file.bold", false, "Bold text")
	fileCmd.Flags().String("file.border", "", "Border Style")
	fileCmd.Flags().String("file.border-background", "", "Border Background Color")
	fileCmd.Flags().String("file.border-foreground", "", "Border Foreground Color")
	fileCmd.Flags().Bool("file.faint", false, "Faint text")
	fileCmd.Flags().String("file.foreground", "", "Foreground Color")
	fileCmd.Flags().String("file.height", "", "Text height")
	fileCmd.Flags().Bool("file.italic", false, "Italicize text")
	fileCmd.Flags().String("file.margin", "", "Text margin")
	fileCmd.Flags().String("file.padding", "", "Text padding")
	fileCmd.Flags().Bool("file.strikethrough", false, "Strikethrough text")
	fileCmd.Flags().Bool("file.underline", false, "Underline text")
	fileCmd.Flags().String("file.width", "", "Text width")
	fileCmd.Flags().String("height", "", "Maximum number of files to display")
	fileCmd.Flags().String("permissions.align", "", "Text Alignment")
	fileCmd.Flags().String("permissions.background", "", "Background Color")
	fileCmd.Flags().Bool("permissions.bold", false, "Bold text")
	fileCmd.Flags().String("permissions.border", "", "Border Style")
	fileCmd.Flags().String("permissions.border-background", "", "Border Background Color")
	fileCmd.Flags().String("permissions.border-foreground", "", "Border Foreground Color")
	fileCmd.Flags().Bool("permissions.faint", false, "Faint text")
	fileCmd.Flags().String("permissions.foreground", "", "Foreground Color")
	fileCmd.Flags().String("permissions.height", "", "Text height")
	fileCmd.Flags().Bool("permissions.italic", false, "Italicize text")
	fileCmd.Flags().String("permissions.margin", "", "Text margin")
	fileCmd.Flags().String("permissions.padding", "", "Text padding")
	fileCmd.Flags().Bool("permissions.strikethrough", false, "Strikethrough text")
	fileCmd.Flags().Bool("permissions.underline", false, "Underline text")
	fileCmd.Flags().String("permissions.width", "", "Text width")
	fileCmd.Flags().String("selected.align", "", "Text Alignment")
	fileCmd.Flags().String("selected.background", "", "Background Color")
	fileCmd.Flags().Bool("selected.bold", false, "Bold text")
	fileCmd.Flags().String("selected.border", "", "Border Style")
	fileCmd.Flags().String("selected.border-background", "", "Border Background Color")
	fileCmd.Flags().String("selected.border-foreground", "", "Border Foreground Color")
	fileCmd.Flags().Bool("selected.faint", false, "Faint text")
	fileCmd.Flags().String("selected.foreground", "", "Foreground Color")
	fileCmd.Flags().String("selected.height", "", "Text height")
	fileCmd.Flags().Bool("selected.italic", false, "Italicize text")
	fileCmd.Flags().String("selected.margin", "", "Text margin")
	fileCmd.Flags().String("selected.padding", "", "Text padding")
	fileCmd.Flags().Bool("selected.strikethrough", false, "Strikethrough text")
	fileCmd.Flags().Bool("selected.underline", false, "Underline text")
	fileCmd.Flags().String("selected.width", "", "Text width")
	fileCmd.Flags().String("symlink.align", "", "Text Alignment")
	fileCmd.Flags().String("symlink.background", "", "Background Color")
	fileCmd.Flags().Bool("symlink.bold", false, "Bold text")
	fileCmd.Flags().String("symlink.border", "", "Border Style")
	fileCmd.Flags().String("symlink.border-background", "", "Border Background Color")
	fileCmd.Flags().String("symlink.border-foreground", "", "Border Foreground Color")
	fileCmd.Flags().Bool("symlink.faint", false, "Faint text")
	fileCmd.Flags().String("symlink.foreground", "", "Foreground Color")
	fileCmd.Flags().String("symlink.height", "", "Text height")
	fileCmd.Flags().Bool("symlink.italic", false, "Italicize text")
	fileCmd.Flags().String("symlink.margin", "", "Text margin")
	fileCmd.Flags().String("symlink.padding", "", "Text padding")
	fileCmd.Flags().Bool("symlink.strikethrough", false, "Strikethrough text")
	fileCmd.Flags().Bool("symlink.underline", false, "Underline text")
	fileCmd.Flags().String("symlink.width", "", "Text width")
	rootCmd.AddCommand(fileCmd)

	carapace.Gen(fileCmd).FlagCompletion(carapace.ActionMap{
		"cursor.align":                  gum.ActionAlignments(),
		"cursor.background":             gum.ActionColors(),
		"cursor.border":                 gum.ActionBorders(),
		"cursor.border-background":      gum.ActionColors(),
		"cursor.border-foreground":      gum.ActionColors(),
		"cursor.foreground":             gum.ActionColors(),
		"directory.align":               gum.ActionAlignments(),
		"directory.background":          gum.ActionColors(),
		"directory.border":              gum.ActionBorders(),
		"directory.border-background":   gum.ActionColors(),
		"directory.border-foreground":   gum.ActionColors(),
		"directory.foreground":          gum.ActionColors(),
		"file-size.align":               gum.ActionAlignments(),
		"file-size.background":          gum.ActionColors(),
		"file-size.border":              gum.ActionBorders(),
		"file-size.border-background":   gum.ActionColors(),
		"file-size.border-foreground":   gum.ActionColors(),
		"file-size.foreground":          gum.ActionColors(),
		"file.align":                    gum.ActionAlignments(),
		"file.background":               gum.ActionColors(),
		"file.border":                   gum.ActionBorders(),
		"file.border-background":        gum.ActionColors(),
		"file.border-foreground":        gum.ActionColors(),
		"file.foreground":               gum.ActionColors(),
		"permissions.align":             gum.ActionAlignments(),
		"permissions.background":        gum.ActionColors(),
		"permissions.border":            gum.ActionBorders(),
		"permissions.border-background": gum.ActionColors(),
		"permissions.border-foreground": gum.ActionColors(),
		"permissions.foreground":        gum.ActionColors(),
		"selected.align":                gum.ActionAlignments(),
		"selected.background":           gum.ActionColors(),
		"selected.border":               gum.ActionBorders(),
		"selected.border-background":    gum.ActionColors(),
		"selected.border-foreground":    gum.ActionColors(),
		"selected.foreground":           gum.ActionColors(),
		"symlink.align":                 gum.ActionAlignments(),
		"symlink.background":            gum.ActionColors(),
		"symlink.border":                gum.ActionBorders(),
		"symlink.border-background":     gum.ActionColors(),
		"symlink.border-foreground":     gum.ActionColors(),
		"symlink.foreground":            gum.ActionColors(),
	})

	carapace.Gen(fileCmd).PositionalCompletion(
		carapace.ActionDirectories(),
	)
}
