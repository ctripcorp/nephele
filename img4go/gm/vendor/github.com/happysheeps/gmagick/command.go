package gmagick

/*
#include <wand/wand_api.h>
#include "magick_wand.h"
unsigned int dissolveImage(MagickWand *wand, const unsigned int dissolve) {
    int x,y;
    register PixelPacket *q;
    struct FakeMagickWand *fmw = (struct FakeMagickWand *)wand;
    Image *image = fmw->image;
    if (!image->matte)
        SetImageOpacity(image,OpaqueOpacity);
    for(y=0; y< (long)image->rows; y++) {
        q=GetImagePixels(image,0,y,image->columns,1);
        if (q == (PixelPacket *) NULL) {
            CopyException(&fmw->exception, &image->exception);
            return(MagickFalse);
        }
        for (x=0; x < (long) image->columns; x++) {
            if(q->opacity != MaxRGB) {
                q->opacity=(Quantum)(MaxRGB - ((MaxRGB-q->opacity)/100.0*dissolve));
            }
            q++;
        }
        if (!SyncImagePixels(image)) {
            CopyException(&fmw->exception, &image->exception);
            return(MagickFalse);
        }
    }
    return(MagickTrue);
}
*/
import "C"

// Sets current image dissolved percent in composite.
// It multiplies current image opacity by the percent.
//
// percent: range from 0(completely transparent)-100(opacity)
func (mw *MagickWand) Dissolve(percent int) error {
	ok := C.dissolveImage(mw.mw, C.uint(percent))
	return mw.getLastErrorIfFailed(ok)
}