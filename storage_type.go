package gmagick

/*
#include <wand/wand_api.h>
*/
import "C"

type StorageType int

//see http://www.graphicsmagick.org/api/types.html#storagetype
/*
typedef enum
{
  CharPixel,
  ShortPixel,
  IntegerPixel,
  LongPixel,
  FloatPixel,
  DoublePixel
} StorageType;
 */

const (
	STORAGE_CHAR_PIXEL    StorageType = C.CharPixel
	STORAGE_SHORT_PIXEL   StorageType = C.ShortPixel
	STORAGE_INTEGER_PIXEL StorageType = C.IntegerPixel
	STORAGE_LONG_PIXEL    StorageType = C.LongPixel
	STORAGE_FLOAT_PIXEL   StorageType = C.FloatPixel
	STORAGE_DOUBLE_PIXEL  StorageType = C.DoublePixel
)
