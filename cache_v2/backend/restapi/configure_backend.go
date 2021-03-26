// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"log"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	bops "backend/backendoperations"
	"backend/restapi/operations"
)

//go:generate swagger generate server --target ../../backend --name Backend --spec ../swagger.yaml --principal interface{}

func configureFlags(api *operations.BackendAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.BackendAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	//if api.GetSearchHandler == nil {
	api.GetSearchHandler = operations.GetSearchHandlerFunc(func(params operations.GetSearchParams) middleware.Responder {
		api.Logger("Offset received: %d\n", params.Offset)
		api.Logger("Empolyee Name received: %s\n", params.EmpName)
		res, err := bops.SearchResult(int(params.Offset), params.EmpName)
		if err != nil {
			return operations.NewGetSearchBadRequest()
		}
		api.Logger("Value received from SearchResults: %s", res)
		response := &operations.GetSearchOKBody{
			Value: res,
		}
		return operations.NewGetSearchOK().WithPayload(response)
	})
	//}

	//if api.GetSearchbyIDHandler == nil {
	api.GetSearchbyIDHandler = operations.GetSearchbyIDHandlerFunc(func(params operations.GetSearchbyIDParams) middleware.Responder {
		api.Logger("Params received in searchbyID: %s", params.Key)
		res, err := bops.ReadID(params.Key)
		api.Logger("Data received from cache: %s", res)
		if err != nil {
			return operations.NewGetSearchbyIDBadRequest()
		}

		if res == "No record found" {
			return operations.NewGetSearchbyIDBadRequest()
		}
		response := &operations.GetSearchbyIDOKBody{
			Value: res,
		}

		return operations.NewGetSearchbyIDOK().WithPayload(response)
	})
	//}

	//if api.PostSubmitHandler == nil {
	api.PostSubmitHandler = operations.PostSubmitHandlerFunc(func(params operations.PostSubmitParams) middleware.Responder {
		api.Logger("key received: %d", *params.Data.Key)
		api.Logger("Value received: %d", *params.Data.Value)

		data := make(map[string]string)
		data["EmployeeID"] = *params.Data.Key
		data["EmployeeName"] = *params.Data.Value
		api.Logger("data received: %v", data)

		_, dbStatus := bops.Submit(data)

		if dbStatus {
			return operations.NewPostSubmitOK()
		}

		return operations.NewPostSubmitBadRequest()
	})
	//}

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
