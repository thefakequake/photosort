# Photosort
A program written in Go that sorts images based on EXIF metadata into folders.

# Use
To use the program, compile on Go 1.18 and run once to generate config file.

After applying config, run the program again.

# Config
`inputDir` - input photo directory
`outputDir` - output photo directory, make sure this isn't inside your `inputDir`
`supportedFormats` - array of supported image formats
`namePrefix` - prefix to name files with
`nextImage` - next photo ID, do not change unless you want to reset the count
