package injector

import "github.com/achmadAlli/go-simple-boilerplate/adapters/api/v1/handler"

func ProvideSwaggerHandler() *handler.SwaggerHandler {
	return handler.NewSwaggerHandler()
}

func ProvideAuthMiddleware() interface{} {
	return nil
}
