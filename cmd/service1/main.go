package main

import (
	"fmt"

	_ "github.com/lxhanghub/go-mfish/api/service1/docs" // swagger 一定要有这行,指向你的文档地址
	"github.com/lxhanghub/go-mfish/internal/service1/grpcapi/hello"
	"github.com/lxhanghub/go-mfish/internal/service1/webapi"
	"github.com/lxhanghub/go-mfish/pkg/host"
	"go.uber.org/zap"
)

func main() {

	builder := host.NewWebHostBuilder()

	builder.ConfigureAppConfiguration(func(build host.ConfigBuilder) {
		build.AddYamlFile("./config.yaml")
	})

	app, err := builder.Build()

	if err != nil {
		fmt.Printf("Failed to build application: %v\n", err)
		return
	}

	if app.Env.IsDevelopment {
		app.UseSwagger()
	}

	app.MapRoutes(webapi.Hello)

	app.MapGrpcServices(hello.NewHelloService)

	if err := app.Run(); err != nil {
		app.Logger().Error("Error running application", zap.Error(err))
	}
}
