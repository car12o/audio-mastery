// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/car12o/audio-mastery/api/generated/models"
	"github.com/car12o/audio-mastery/api/generated/restapi/operations"
	"github.com/car12o/audio-mastery/api/generated/restapi/operations/audios"
	"github.com/car12o/audio-mastery/api/generated/restapi/operations/auth"
	"github.com/car12o/audio-mastery/api/generated/restapi/operations/info"
	"github.com/car12o/audio-mastery/api/generated/restapi/operations/users"
)

//go:generate swagger generate server --target ../../generated --name AudioMastery --spec ../swagger.yaml --principal models.Principal --exclude-main

func configureFlags(api *operations.AudioMasteryAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.AudioMasteryAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	// Applies when the "Authorization" header is set
	if api.BearerAuth == nil {
		api.BearerAuth = func(token string) (*models.Principal, error) {
			return nil, errors.NotImplemented("api key auth (bearer) Authorization from header param [Authorization] has not yet been implemented")
		}
	}

	// Set your custom authorizer if needed. Default one is security.Authorized()
	// Expected interface runtime.Authorizer
	//
	// Example:
	// api.APIAuthorizer = security.Authorized()

	if api.AudiosDeleteAudioHandler == nil {
		api.AudiosDeleteAudioHandler = audios.DeleteAudioHandlerFunc(func(params audios.DeleteAudioParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation audios.DeleteAudio has not yet been implemented")
		})
	}
	if api.AudiosGetAudioHandler == nil {
		api.AudiosGetAudioHandler = audios.GetAudioHandlerFunc(func(params audios.GetAudioParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation audios.GetAudio has not yet been implemented")
		})
	}
	if api.InfoGetInfoHandler == nil {
		api.InfoGetInfoHandler = info.GetInfoHandlerFunc(func(params info.GetInfoParams) middleware.Responder {
			return middleware.NotImplemented("operation info.GetInfo has not yet been implemented")
		})
	}
	if api.AuthGetLogoutHandler == nil {
		api.AuthGetLogoutHandler = auth.GetLogoutHandlerFunc(func(params auth.GetLogoutParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation auth.GetLogout has not yet been implemented")
		})
	}
	if api.AudiosListAudiosHandler == nil {
		api.AudiosListAudiosHandler = audios.ListAudiosHandlerFunc(func(params audios.ListAudiosParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation audios.ListAudios has not yet been implemented")
		})
	}
	if api.AudiosPostAudioHandler == nil {
		api.AudiosPostAudioHandler = audios.PostAudioHandlerFunc(func(params audios.PostAudioParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation audios.PostAudio has not yet been implemented")
		})
	}
	if api.AuthPostLoginHandler == nil {
		api.AuthPostLoginHandler = auth.PostLoginHandlerFunc(func(params auth.PostLoginParams) middleware.Responder {
			return middleware.NotImplemented("operation auth.PostLogin has not yet been implemented")
		})
	}
	if api.UsersPostUserHandler == nil {
		api.UsersPostUserHandler = users.PostUserHandlerFunc(func(params users.PostUserParams) middleware.Responder {
			return middleware.NotImplemented("operation users.PostUser has not yet been implemented")
		})
	}
	if api.AudiosPutAudioHandler == nil {
		api.AudiosPutAudioHandler = audios.PutAudioHandlerFunc(func(params audios.PutAudioParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation audios.PutAudio has not yet been implemented")
		})
	}

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
