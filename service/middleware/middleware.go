package middleware

import (
	"github.com/gin-gonic/gin"
	"reflect"
)

type Config interface {
	Order() int
	Handler() gin.HandlerFunc
}

type MiddlewareConfig struct {
	CORS *CORSConfig `toml:"cors"`
}

var configInterface = reflect.TypeOf((*Config)(nil)).Elem()
var middlewareConfigStructType = reflect.TypeOf(MiddlewareConfig{})
var middlewareConfigPointerType = reflect.TypeOf(&MiddlewareConfig{})

// Build middleware from conf
// conf could combine MiddlewarConfig or have field that implements Config interface
// eg:
// type CustomizedConfig struct {
//		*MiddlewareConfig
//	    YourConfig *YourMiddlewareConfig
// }
// type YourMiddlewareConfig struct {
// }
// func (conf *YourMiddlewareConfig) Handler() gin.HandlerFunc{
//		return func(ctx *gin.Context) {
//		todo:
// }
//}
//only those implement Config interface will be registered.
func Build(conf interface{}) []gin.HandlerFunc {
	fn := []gin.HandlerFunc{recovery()}
	if conf != nil {
		fn = append(fn, build(conf)...)
	}
	return fn
}

//internal build by reflection.
func build(conf interface{}) []gin.HandlerFunc {
	var o reflect.Value

	fn := make([]gin.HandlerFunc, 0)

	v := reflect.ValueOf(conf)
	switch v.Kind() {
	case reflect.Ptr:
		fallthrough
	case reflect.Interface:
		o = v.Elem()
	case reflect.Struct:
		o = v
	}

	t := o.Type()
	for i := 0; i < t.NumField(); i++ {
		if t.Field(i).Type.Implements(configInterface) {
			c := reflect.ValueOf(o.Field(i).Interface())
			if !c.IsNil() {
				ret := c.MethodByName("Handler").Call(nil)
				handler := ret[0].Interface().(gin.HandlerFunc)
				fn = append(fn, handler)
			}
		} else {
			if t.Field(i).Type == middlewareConfigStructType ||
				t.Field(i).Type == middlewareConfigPointerType {
				fn = append(fn, build(o.Field(i).Interface())...)
			}
		}
	}

	return fn
}
