package appProps

import (
	"sync"

	"github.com/magiconair/properties"

	"github.com/pksingh/gin-curd-demo/config"
)

var (
	// onceInit guarantees initialization of properties only once
	onceInit      = new(sync.Once)
	concreteImpls = make(map[string]interface{})
)

const (
	propsImplKey = "propsImpl"
)

// Load is an exported method that loads props depending on environment
func Load(dir string) error {
	var appErr error
	onceInit.Do(func() {
		appErr = loadImpls(dir)
	})
	return appErr
}

func loadImpls(dir string) error {
	if concreteImpls[propsImplKey] == nil {
		allProps := config.GetProps(dir)
		concreteImpls[propsImplKey] = allProps
	}
	return nil
}

// GetAll gives all the application properties loaded depending on the env.
// This can be used from anywhere in the app to get configuration data.
func GetAll() *properties.Properties {
	x, _ := concreteImpls[propsImplKey].(*properties.Properties)
	return x
}
