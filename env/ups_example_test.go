package env_test

import (
	"fmt"

	"github.com/cloud-gov/go-cfenv"
	"github.com/cloud-gov/cf-common/v2/env"
)

func ExampleWithUPSLookup() {
	app, err := cfenv.Current()
	if err != nil {
		// ...
	}

	opts := []env.VarSetOpt{
		env.WithOSLookup(), // Always look in the OS env first.
		env.WithUPSLookup(app, "service-1"),
	}

	vs := env.NewVarSet(opts...)

	v := vs.MustString("FOO")

	fmt.Println(v)
}
