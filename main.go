package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/httpadapter"
)

//go:embed static-files
var staticFiles embed.FS

func main() {
	serverRoot, _ := fs.Sub(staticFiles, "static-files")

	http.Handle("/", http.FileServer(http.FS(serverRoot)))

	if os.Getenv("AWS_LAMBDA_FUNCTION_NAME") == "" {
		// local
		log.Fatal(http.ListenAndServe(":3000", http.DefaultServeMux))
	} else {
		// lambda
		lambda.Start(httpadapter.New(http.DefaultServeMux).ProxyWithContext)
	}
}
