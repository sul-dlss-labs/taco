// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	errors "github.com/go-openapi/errors"
	loads "github.com/go-openapi/loads"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	security "github.com/go-openapi/runtime/security"
	spec "github.com/go-openapi/spec"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/sul-dlss-labs/taco/authorization"
)

// NewTacoAPI creates a new Taco instance
func NewTacoAPI(spec *loads.Document) *TacoAPI {
	return &TacoAPI{
		handlers:              make(map[string]map[string]http.Handler),
		formats:               strfmt.Default,
		defaultConsumes:       "application/json",
		defaultProduces:       "application/json",
		ServerShutdown:        func() {},
		spec:                  spec,
		ServeError:            errors.ServeError,
		BasicAuthenticator:    security.BasicAuth,
		APIKeyAuthenticator:   security.APIKeyAuth,
		BearerAuthenticator:   security.BearerAuth,
		JSONConsumer:          runtime.JSONConsumer(),
		MultipartformConsumer: runtime.DiscardConsumer,
		JSONProducer:          runtime.JSONProducer(),
		DeleteResourceHandler: DeleteResourceHandlerFunc(func(params DeleteResourceParams) middleware.Responder {
			return middleware.NotImplemented("operation DeleteResource has not yet been implemented")
		}),
		DepositFileHandler: DepositFileHandlerFunc(func(params DepositFileParams, principal *authorization.Agent) middleware.Responder {
			return middleware.NotImplemented("operation DepositFile has not yet been implemented")
		}),
		DepositResourceHandler: DepositResourceHandlerFunc(func(params DepositResourceParams, principal *authorization.Agent) middleware.Responder {
			return middleware.NotImplemented("operation DepositResource has not yet been implemented")
		}),
		GetProcessStatusHandler: GetProcessStatusHandlerFunc(func(params GetProcessStatusParams, principal *authorization.Agent) middleware.Responder {
			return middleware.NotImplemented("operation GetProcessStatus has not yet been implemented")
		}),
		HealthCheckHandler: HealthCheckHandlerFunc(func(params HealthCheckParams) middleware.Responder {
			return middleware.NotImplemented("operation HealthCheck has not yet been implemented")
		}),
		RetrieveFileHandler: RetrieveFileHandlerFunc(func(params RetrieveFileParams, principal *authorization.Agent) middleware.Responder {
			return middleware.NotImplemented("operation RetrieveFile has not yet been implemented")
		}),
		RetrieveResourceHandler: RetrieveResourceHandlerFunc(func(params RetrieveResourceParams, principal *authorization.Agent) middleware.Responder {
			return middleware.NotImplemented("operation RetrieveResource has not yet been implemented")
		}),
		UpdateResourceHandler: UpdateResourceHandlerFunc(func(params UpdateResourceParams) middleware.Responder {
			return middleware.NotImplemented("operation UpdateResource has not yet been implemented")
		}),

		// Applies when the "On-Behalf-Of" header is set
		RemoteUserAuth: func(token string) (*authorization.Agent, error) {
			return nil, errors.NotImplemented("api key auth (RemoteUser) On-Behalf-Of from header param [On-Behalf-Of] has not yet been implemented")
		},

		// default authorizer is authorized meaning no requests are blocked
		APIAuthorizer: security.Authorized(),
	}
}

/*TacoAPI TACO, the Stanford Digital Repository (SDR) Management Layer API */
type TacoAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler

	// BasicAuthenticator generates a runtime.Authenticator from the supplied basic auth function.
	// It has a default implemention in the security package, however you can replace it for your particular usage.
	BasicAuthenticator func(security.UserPassAuthentication) runtime.Authenticator
	// APIKeyAuthenticator generates a runtime.Authenticator from the supplied token auth function.
	// It has a default implemention in the security package, however you can replace it for your particular usage.
	APIKeyAuthenticator func(string, string, security.TokenAuthentication) runtime.Authenticator
	// BearerAuthenticator generates a runtime.Authenticator from the supplied bearer token auth function.
	// It has a default implemention in the security package, however you can replace it for your particular usage.
	BearerAuthenticator func(string, security.ScopedTokenAuthentication) runtime.Authenticator

	// JSONConsumer registers a consumer for a "application/json+ld" mime type
	JSONConsumer runtime.Consumer
	// MultipartformConsumer registers a consumer for a "multipart/form-data" mime type
	MultipartformConsumer runtime.Consumer

	// JSONProducer registers a producer for a "application/json" mime type
	JSONProducer runtime.Producer

	// RemoteUserAuth registers a function that takes a token and returns a principal
	// it performs authentication based on an api key On-Behalf-Of provided in the header
	RemoteUserAuth func(string) (*authorization.Agent, error)

	// APIAuthorizer provides access control (ACL/RBAC/ABAC) by providing access to the request and authenticated principal
	APIAuthorizer runtime.Authorizer

	// DeleteResourceHandler sets the operation handler for the delete resource operation
	DeleteResourceHandler DeleteResourceHandler
	// DepositFileHandler sets the operation handler for the deposit file operation
	DepositFileHandler DepositFileHandler
	// DepositResourceHandler sets the operation handler for the deposit resource operation
	DepositResourceHandler DepositResourceHandler
	// GetProcessStatusHandler sets the operation handler for the get process status operation
	GetProcessStatusHandler GetProcessStatusHandler
	// HealthCheckHandler sets the operation handler for the health check operation
	HealthCheckHandler HealthCheckHandler
	// RetrieveFileHandler sets the operation handler for the retrieve file operation
	RetrieveFileHandler RetrieveFileHandler
	// RetrieveResourceHandler sets the operation handler for the retrieve resource operation
	RetrieveResourceHandler RetrieveResourceHandler
	// UpdateResourceHandler sets the operation handler for the update resource operation
	UpdateResourceHandler UpdateResourceHandler

	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// SetDefaultProduces sets the default produces media type
func (o *TacoAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *TacoAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *TacoAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *TacoAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *TacoAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *TacoAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *TacoAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the TacoAPI
func (o *TacoAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.MultipartformConsumer == nil {
		unregistered = append(unregistered, "MultipartformConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.RemoteUserAuth == nil {
		unregistered = append(unregistered, "OnBehalfOfAuth")
	}

	if o.DeleteResourceHandler == nil {
		unregistered = append(unregistered, "DeleteResourceHandler")
	}

	if o.DepositFileHandler == nil {
		unregistered = append(unregistered, "DepositFileHandler")
	}

	if o.DepositResourceHandler == nil {
		unregistered = append(unregistered, "DepositResourceHandler")
	}

	if o.GetProcessStatusHandler == nil {
		unregistered = append(unregistered, "GetProcessStatusHandler")
	}

	if o.HealthCheckHandler == nil {
		unregistered = append(unregistered, "HealthCheckHandler")
	}

	if o.RetrieveFileHandler == nil {
		unregistered = append(unregistered, "RetrieveFileHandler")
	}

	if o.RetrieveResourceHandler == nil {
		unregistered = append(unregistered, "RetrieveResourceHandler")
	}

	if o.UpdateResourceHandler == nil {
		unregistered = append(unregistered, "UpdateResourceHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *TacoAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *TacoAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {

	result := make(map[string]runtime.Authenticator)
	for name, scheme := range schemes {
		switch name {

		case "RemoteUser":

			result[name] = o.APIKeyAuthenticator(scheme.Name, scheme.In, func(token string) (interface{}, error) {
				return o.RemoteUserAuth(token)
			})

		}
	}
	return result

}

// Authorizer returns the registered authorizer
func (o *TacoAPI) Authorizer() runtime.Authorizer {

	return o.APIAuthorizer

}

// ConsumersFor gets the consumers for the specified media types
func (o *TacoAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {

	result := make(map[string]runtime.Consumer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/json":
			result["application/json"] = o.JSONConsumer

		case "application/json+ld":
			result["application/json+ld"] = o.JSONConsumer

		case "multipart/form-data":
			result["multipart/form-data"] = o.MultipartformConsumer

		}
	}
	return result

}

// ProducersFor gets the producers for the specified media types
func (o *TacoAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {

	result := make(map[string]runtime.Producer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/json":
			result["application/json"] = o.JSONProducer

		}
	}
	return result

}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *TacoAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	if path == "/" {
		path = ""
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

// Context returns the middleware context for the taco API
func (o *TacoAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *TacoAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened

	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/resource/{ID}"] = NewDeleteResource(o.context, o.DeleteResourceHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/resource/{FilesetID}/file"] = NewDepositFile(o.context, o.DepositFileHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/resource"] = NewDepositResource(o.context, o.DepositResourceHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/status/{ID}"] = NewGetProcessStatus(o.context, o.GetProcessStatusHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/healthcheck"] = NewHealthCheck(o.context, o.HealthCheckHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/file/{ID}"] = NewRetrieveFile(o.context, o.RetrieveFileHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/resource/{ID}"] = NewRetrieveResource(o.context, o.RetrieveResourceHandler)

	if o.handlers["PATCH"] == nil {
		o.handlers["PATCH"] = make(map[string]http.Handler)
	}
	o.handlers["PATCH"]["/resource/{ID}"] = NewUpdateResource(o.context, o.UpdateResourceHandler)

}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *TacoAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middelware as you see fit
func (o *TacoAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}
