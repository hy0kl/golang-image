golang-image
============

Adds support of CMYK jpeg files to Golang image library.

You can use it as a drop-in replacement for the Golang default image library.

import "github.com/zhnxin/golang-image/jpeg"

## JFIF support

You can wirte jfif metadata by calling:

    EncodeWithJfif(w io.Writer, m image.Image,jfif *Jfif, o *Options)

And set the thumbnail by calling:

    jfif := jpeg.NewJfif()
    err = jfif.SetThumbnail(thumbnail)
    checkErr(err)

_Notic: The thumbnail should not be lager than 21845 px, which means width*hight <= 21845_
