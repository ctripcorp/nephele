package process

import (
	"fmt"
	"strconv"
	"strings"
)

// Process uri process info
type Process struct {
	Name  Cmd
	Param map[string]string
}

//BuildProcesses parse uri and create process array
func BuildProcesses(uri string) ([]Process, error) {
	procs, e := parseURI(uri)
	if e != nil {
		return nil, e
	}
	return checkParam(procs)
}

func parseURI(uri string) ([]Process, error) {
	procs := strings.Split(uri, "?")
	prefix := "x-nephele-process=image/"
	processes := []Process{}
	for _, proc := range procs {
		if strings.HasPrefix(proc, prefix) {
			cmds := strings.Split(strings.TrimPrefix(proc, prefix), "/")
			for _, cmd := range cmds {
				if cmd == "" {
					continue
				}
				arr := strings.Split(cmd, ",")
				paramMap := make(map[string]string)
				for index := 1; index < len(arr); index++ {
					kv := strings.Split(arr[index], "_")
					if len(kv) != 2 {
						continue
					}
					paramMap[kv[0]] = kv[1]
				}
				processes = append(processes, Process{Name: Cmd(arr[0]), Param: paramMap})
			}
		}
	}
	return processes, nil
}

func checkParam(procs []Process) ([]Process, error) {
	processes := []Process{}
	for _, proc := range procs {
		f, isExists := CmdCheckMap[proc.Name]
		if !isExists {
			continue
		}
		if e := f(proc.Param); e != nil {
			return nil, e
		}
		processes = append(processes, proc)
	}
	return processes, nil
}

//CmdCheckMap Cmd list and check func
var CmdCheckMap = map[Cmd]func(map[string]string) error{
	Resize:     checkResizeParam,
	Crop:       checkCropParam,
	Rotate:     checkRotateParam,
	AutoOrient: nil,
	Format:     checkFormatParam,
	Quality:    checkQualityParam,
	Watermark:  checkWatermarkParam,
	Sharpen:    nil,
	Style:      checkStyleParam,
	Panorama:   nil,
}
var infoFormat = "The value: %s of parameter: %s is invalid."

//Cmd cmd
type Cmd string

//cmd name
const (
	Resize     Cmd = "resize"
	Crop       Cmd = "crop"
	Rotate     Cmd = "rotate"
	AutoOrient Cmd = "autoorient"
	Format     Cmd = "format"
	Quality    Cmd = "quality"
	Watermark  Cmd = "watermark"
	Sharpen    Cmd = "sharpen"
	Style      Cmd = "style"
	Panorama   Cmd = "panorama"
)

//Param cmd param
type Param string

//param key
const (
	ParamValue    Param = "v"
	ParamWidth    Param = "w"
	ParamHeight   Param = "h"
	ParamMethod   Param = "m"
	ParamL        Param = "l"
	ParamPercent  Param = "p"
	ParamX        Param = "x"
	ParamY        Param = "y"
	ParamAIO      Param = "aio"
	ParamName     Param = "n"
	ParamDissolve Param = "d"
	ParamMW       Param = "mw"
	ParamMH       Param = "mh"
	ParamR        Param = "r"
	ParamS        Param = "s"
)

//resize m value
const (
	ResizeMethodFixed = "fixed"
	ResizeMethodLfit  = "lfit"
)

//crop m value
const (
	CropMethodCenter = "c"
	CropMethodTop    = "t"
	CropMethodBottom = "b"
	CropMethodLeft   = "l"
	CropMethodRight  = "r"
	CropMethodWC     = "wc"
	CropMethodHC     = "hc"
	CropMethodResize = "resize"
	CropMethodCrop   = "crop"
)

//format value
const (
	FormatJPG  = "jpg"
	FormatPNG  = "png"
	FormatGIF  = "gif"
	FormatWEBP = "webp"
)

//watermark l value
const (
	WaterMarkLocationNW     = "nw"
	WaterMarkLocationNorth  = "north"
	WaterMarkLocationNE     = "ne"
	WaterMarkLocationWest   = "west"
	WaterMarkLocationCenter = "center"
	WaterMarkLocationEast   = "east"
	WaterMarkLocationSW     = "sw"
	WaterMarkLocationSouth  = "south"
	WaterMarkLocationSE     = "se"
)

//style value
const (
	StyleLOMO = "lomo"
	StyleOIL  = "oil"
)

//procs = append(procs, Process{Name:  })
func checkResizeParam(m map[string]string) error {
	for k, v := range m {
		p := Param(k)
		if p == ParamWidth || p == ParamHeight {
			if _, e := strconv.Atoi(v); e != nil {
				return fmt.Errorf(infoFormat, v, k)
			}
		}
		if p == ParamMethod && v != ResizeMethodFixed && v != ResizeMethodLfit {
			return fmt.Errorf(infoFormat, v, k)
		}
		if p == ParamL && v != "0" && v != "1" {
			return fmt.Errorf(infoFormat, v, k)
		}
		if p == ParamPercent {
			i, e := strconv.Atoi(v)
			if e != nil || i < 0 || i > 10000 {
				return fmt.Errorf(infoFormat, v, k)
			}
		}
	}
	return nil
}

func checkCropParam(m map[string]string) error {
	for k, v := range m {
		p := Param(k)
		if p == ParamMethod &&
			v != CropMethodBottom &&
			v != CropMethodCenter &&
			v != CropMethodCrop &&
			v != CropMethodHC &&
			v != CropMethodLeft &&
			v != CropMethodResize &&
			v != CropMethodRight &&
			v != CropMethodTop &&
			v != CropMethodWC {
			return fmt.Errorf(infoFormat, v, k)
		}

		if p == ParamL && v != "0" && v != "1" {
			return fmt.Errorf(infoFormat, v, k)
		}
		if p == ParamPercent {
			i, e := strconv.Atoi(v)
			if e != nil || i < 0 || i > 10000 {
				return fmt.Errorf(infoFormat, v, k)
			}
		}
		if p == ParamHeight || p == ParamWidth || p == ParamX || p == ParamY {
			if _, e := strconv.Atoi(v); e != nil {
				return fmt.Errorf(infoFormat, v, k)
			}
		}
	}
	return nil
}

func checkRotateParam(m map[string]string) error {
	for k, v := range m {
		p := Param(k)
		if p == ParamValue {
			value, e := strconv.Atoi(v)
			if e != nil || value < 1 || value > 360 {
				return fmt.Errorf(infoFormat, v, k)
			}
		}
	}
	return nil
}

func checkFormatParam(m map[string]string) error {
	for k, v := range m {
		p := Param(k)
		if p == ParamValue &&
			v != FormatGIF &&
			v != FormatJPG &&
			v != FormatPNG &&
			v != FormatWEBP {
			return fmt.Errorf(infoFormat, v, k)
		}
	}
	return nil
}

func checkQualityParam(m map[string]string) error {
	for k, v := range m {
		p := Param(k)
		if p == ParamValue {
			value, e := strconv.Atoi(v)
			if e != nil || value < 1 || value > 100 {
				return fmt.Errorf(infoFormat, v, k)
			}
		}
		if p == ParamAIO && v != "1" && v != "0" {
			return fmt.Errorf(infoFormat, v, k)
		}
	}
	return nil
}

func checkWatermarkParam(m map[string]string) error {
	for k, v := range m {
		p := Param(k)
		if p == ParamDissolve {
			value, e := strconv.Atoi(v)
			if e != nil || value < 1 || value > 100 {
				return fmt.Errorf(infoFormat, v, k)
			}
		}
		if p == ParamL &&
			v != WaterMarkLocationCenter &&
			v != WaterMarkLocationEast &&
			v != WaterMarkLocationNE &&
			v != WaterMarkLocationNorth &&
			v != WaterMarkLocationNW &&
			v != WaterMarkLocationSE &&
			v != WaterMarkLocationSW &&
			v != WaterMarkLocationSouth &&
			v != WaterMarkLocationWest {
			return fmt.Errorf(infoFormat, v, k)
		}
		if p == ParamX || p == ParamY || p == ParamMW || p == ParamMH {
			if _, e := strconv.Atoi(v); e != nil {
				return fmt.Errorf(infoFormat, v, k)
			}
		}
	}
	return nil
}

func checkStyleParam(m map[string]string) error {
	for k, v := range m {
		p := Param(k)
		if p == ParamValue &&
			v != StyleLOMO &&
			v != StyleOIL {
			return fmt.Errorf(infoFormat, v, k)
		}
	}
	return nil
}
