// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"encoding/json"
	"log"
	"net/http"

	"database/restapi/operations"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"database/dbops"
	m "database/models"
)

var (
	dbConn *dbops.Dbconnection
	kp     *dbops.KafkaProducer
)

//go:generate swagger generate server --target ../../database --name Database --spec ../swagger.yaml --principal interface{}

func configureFlags(api *operations.DatabaseAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.DatabaseAPI) http.Handler {
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

	//if api.PostLoadHandler == nil {
	api.PostLoadHandler = operations.PostLoadHandlerFunc(func(params operations.PostLoadParams) middleware.Responder {
		api.Logger("Params received is: %v\n", params.Data.EmployeeID)

		data := dbConn.FetchMultipleRecord(params.Data.EmployeeName)
		api.Logger("data is: %v\n", data)
		dataInByte, _ := json.Marshal(data)
		api.Logger("dataInByte is: %v\n", string(dataInByte))
		err := kp.SendMessage(string(dataInByte))
		if err != nil {
			return operations.NewPostLoadBadRequest()
		}
		var response m.Validstringresponse
		response = "loaded successfully"
		return operations.NewPostLoadOK().WithPayload(response)
	})
	//}
	//if api.GetReadHandler == nil {
	api.GetReadHandler = operations.GetReadHandlerFunc(func(params operations.GetReadParams) middleware.Responder {
		api.Logger("Employee ID received in request is: %v", params.EmployeeID)

		record := dbConn.FetchOneRecord(params.EmployeeID)
		api.Logger("response received is %v %v\n", *record.EmployeeID, *record.EmployeeName)
		var response *m.Record
		response = &m.Record{
			EmployeeID:   record.EmployeeID,
			EmployeeName: record.EmployeeName,
		}

		return operations.NewGetReadOK().WithPayload(response)
	})
	//}

	//if api.PostWriteHandler == nil {
	api.PostWriteHandler = operations.PostWriteHandlerFunc(func(params operations.PostWriteParams) middleware.Responder {
		api.Logger("Employee ID received in request is: %v", *params.Data.EmployeeID)
		api.Logger("Employee Name received in request is: %v", *params.Data.EmployeeName)
		var response m.Validstringresponse
		response = "created successfully"
		status := dbConn.CreateRecord(params.Data.EmployeeID, params.Data.EmployeeName)
		if status {
			return operations.NewPostWriteOK().WithPayload(response)
		}
		return operations.NewPostWriteBadRequest()
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
	dbConn, kp = dbops.Init()
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
