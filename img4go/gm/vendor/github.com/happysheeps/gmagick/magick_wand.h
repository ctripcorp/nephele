#include <wand/wand_api.h>

struct FakeMagickWand
{
  char id[MaxTextExtent];
  ExceptionInfo exception;
  ImageInfo *image_info;
  QuantizeInfo   *quantize_info;
  Image    *image,    *images;
  unsigned int iterator;
  unsigned long signature;
};