package driver

import (
	"fmt"

	"github.com/deislabs/cnab-go/driver"
	"github.com/deislabs/cnab-go/driver/docker"
	"github.com/deislabs/cnab-go/driver/kubernetes"
)

// Lookup takes a driver name and tries to resolve the most pertinent driver.
func Lookup(name string) (driver.Driver, error) {
	switch name {
	case "docker":
		return &docker.Driver{}, nil
	case "kubernetes", "k8s":
		return &kubernetes.Driver{}, nil
	case "debug":
		return &driver.DebugDriver{}, nil
	case "command":
		return &CommandDriver{Name: name}, nil
	default:
		return nil, fmt.Errorf("unsupported driver: %s", name)
	}
}
