// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"log"
	"net/http"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"cache/cacheops"
	"cache/models"
	"cache/restapi/operations"
)

var (
	rc *cacheops.RedisConnect
	kc *cacheops.KafkaConsumer
)

//go:generate swagger generate server --target ../../cache --name Cache --spec ../swagger.yaml --principal interface{}

func configureFlags(api *operations.CacheAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.CacheAPI) http.Handler {
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

	//if api.GetGetHandler == nil {
	api.GetGetHandler = operations.GetGetHandlerFunc(func(params operations.GetGetParams) middleware.Responder {
		api.Logger("params received in get request: %v\n", params.EmployeeID)

		result, err := rc.Read(params.EmployeeID)
		if err != nil {
			api.Logger("Not able to find the key: %v\n", err)
			return operations.NewGetGetBadRequest()
		}

		response := &models.Record{
			EmployeeID:   &params.EmployeeID,
			EmployeeName: &result,
		}

		return operations.NewGetGetOK().WithPayload(response)
	})
	//}
	//if api.PostSearchHandler == nil {
	api.PostSearchHandler = operations.PostSearchHandlerFunc(func(params operations.PostSearchParams) middleware.Responder {
		api.Logger("Params received in search request: %v\n", *params.Data.Key)
		api.Logger("Params received in search request: %v\n", *params.Data.Start)
		api.Logger("Params received in search request: %v\n", *params.Data.Stop)
		start, _ := strconv.Atoi(*params.Data.Start)
		stop, _ := strconv.Atoi(*params.Data.Stop)
		key := string(*params.Data.Key)
		results, err := rc.Lrange(key, int64(start), int64(stop))
		if err != nil {
			api.Logger("Not able to Lrange data: %v\n", err)
			return operations.NewPostSearchBadRequest()
		}

		var responses []*models.Record

		// convert []string to []*model.Record
		for empid, empname := range results {
			api.Logger("results received from cache: %v\n", empname)
			var eid string
			eid = strconv.Itoa(start + empid)
			// new memory address will be assigned to response
			// instead of using same memory address
			var response models.Record
			var name string
			name = empname
			// response := &model.Record{} will give the last element only
			response = models.Record{
				EmployeeID:   &eid,
				EmployeeName: &name,
			}

			responses = append(responses, &response)
		}

		for _, response := range responses {
			api.Logger("Response in search request: %v\n", response.EmployeeName)
		}

		return operations.NewPostSearchOK().WithPayload(responses)
	})
	//}
	//if api.PostWriteHandler == nil {
	api.PostWriteHandler = operations.PostWriteHandlerFunc(func(params operations.PostWriteParams) middleware.Responder {
		api.Logger("Params received in write req: %v\n", *params.Data.EmployeeID)
		api.Logger("Params received in write req: %v\n", *params.Data.EmployeeName)

		_, err := rc.Store(*params.Data.EmployeeID, []byte(*params.Data.EmployeeName))
		if err != nil {
			api.Logger("Not able to Push data: %v\n", err)
			return operations.NewPostWriteBadRequest()
		}
		var response models.Validstringresponse
		response = "Successfully set in cache"
		return operations.NewPostWriteOK().WithPayload(response)
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
	rc, kc = cacheops.InitCacheAndConsumer()
	go kc.ReadMessage()
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
