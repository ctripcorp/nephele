package gmagick

/*
#include <wand/wand_api.h>
*/
import "C"

type ResourceType int

const (
	RESOURCE_UNDEFINED ResourceType = C.UndefinedResource
	RESOURCE_AREA      ResourceType = C.AreaResource
	RESOURCE_DISK      ResourceType = C.DiskResource
	RESOURCE_FILE      ResourceType = C.FileResource
	RESOURCE_MAP       ResourceType = C.MapResource
	RESOURCE_MEMORY    ResourceType = C.MemoryResource
)