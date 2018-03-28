package middleware

import (
	"reflect"
	"sort"

	"github.com/gin-gonic/gin"
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

type byOrder struct {
	order   int
	handler gin.HandlerFunc
}

// Build middleware from conf
// conf could combine MiddlewarConfig or have fields that implement Config interface
// eg:
// type CustomizedConfig struct {
//		*MiddlewareConfig
//	    YourConfig *YourMiddlewareConfig `toml:"yours"`
// }
// type YourMiddlewareConfig struct {
//      RegistOrder int `toml:"order"`
// }
// func (conf *YourMiddlewareConfig) Order() int {
//	    return conf.RegistOrder
// }
// func (conf *YourMiddlewareConfig) Handler() gin.HandlerFunc{
//		return func(ctx *gin.Context) {
//		todo:
// }
//}
//only those implement Config interface will be registered.
func Build(conf interface{}) []gin.HandlerFunc {
	//recovery should always be the first middleware
	array := []byOrder{byOrder{order: -99, handler: entrance()}}
	if conf != nil {
		array = append(array, build(conf)...)
		sort.Slice(array, func(i, j int) bool {
			return array[i].order < array[j].order
		})
	}
	handlers := make([]gin.HandlerFunc, len(array))
	for i, h := range array {
		handlers[i] = h.handler
	}
	return handlers
}

//internal build by reflection.
func build(conf interface{}) []byOrder {
	var o reflect.Value

	array := make([]byOrder, 0)

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
				ret := c.MethodByName("Order").Call(nil)
				order := ret[0].Interface().(int)
				ret = c.MethodByName("Handler").Call(nil)
				handler := ret[0].Interface().(gin.HandlerFunc)
				if order == 0 {
					order = 99
				}
				array = append(array, byOrder{order: order, handler: handler})
			}
		} else {
			if t.Field(i).Type == middlewareConfigStructType ||
				t.Field(i).Type == middlewareConfigPointerType {
				array = append(array, build(o.Field(i).Interface())...)
			}
		}
	}

	return array
}
