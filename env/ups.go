package env

import (
	"fmt"

	"github.com/cloud-gov/go-cfenv"
)

// WithUPSLookup configures the VarSet to use the CloudFoundry user-provided
// service with the given name as a lookup source.
func WithUPSLookup(app *cfenv.App, name string) VarSetOpt {
	return func(v *VarSet) {
		v.AppendSource(NewLookupFromUPS(app, name))
	}
}

// NewLookupFromUPS looks for a CloudFoundry bound service with the given name.
// This allows sourcing environment variables from a user-provided service.
// If no service is found, a passthrough lookup is used that always returns
// false.
func NewLookupFromUPS(app *cfenv.App, name string) Lookup {
	if app == nil {
		return NoopLookup
	}
	service, err := app.Services.WithName(name)
	// The only error WithName will return is that the service was not found,
	// so it is OK that we ignore the error here.
	if err != nil {
		return NoopLookup
	}
	return func(k string) (string, bool) {
		v, ok := service.Credentials[k]
		if !ok {
			return "", false
		}
		// The lookup mechanism expects strings to be returned. However UPS
		// Credentials allow any JSON value. To get around this, we always just
		// format the value as a string. This means that nested values like
		// arrays and objects will lose their meaning and cannot really be used
		// by this library.
		return fmt.Sprintf("%v", v), true
	}
}
