package main

import (
	"github.com/KubeOperator/KubeOperator/pkg/server"
	_ "golang.org/x/text/message"
	_ "golang.org/x/text/message/catalog"
	"log"
)

// @title KubeOperator Restful API
// @version.go 1.0
// @termsOfService http://kubeoperator.io
// @contact.name Fit2cloud Support
// @contact.url https://www.fit2cloud.com
// @contact.email support@fit2cloud.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @ho- st localhost:8080
// @BasePath /api/v1
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}
