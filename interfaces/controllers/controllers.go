package interfaces

type Context interface {
	Param(key string) string
    JSON(code int, obj interface{})
}
