golang-image
============

Adds support of CMYK jpeg files to Golang image library.

You can use it as a drop-in replacement for the Golang default image library.

import "github.com/zhnxin/golang-image/jpeg"

## JFIF support

You can wirte jfif metadata by calling:

    EncodeWithJfif(w io.Writer, m image.Image,jfif *Jfif, o *Options)

JFIF is not complete support, such as thumbnail. But the rest is accessable.And if a nil Jfif is passed, the default value would be flush to output.