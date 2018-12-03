package main

import (
	"errors"
	"os"
	"path/filepath"
	"plugin"
	"strings"
	"sync"
)

// RunHandler - the run function within a plugin
type RunHandler = func(*sync.Map)

// BootPlugins - initialize the plugins
func BootPlugins() error {
	for _, name := range strings.Split(os.Getenv("SERVER_PLUGINS_DIR"), ",") {
		name = strings.ToLower(strings.TrimSpace(name))
		filename := filepath.Join(os.Getenv("SERVER_PLUGINS_DIR"), name)
		plgn, err := plugin.Open(filename)
		if err != nil {
			return errors.New("plugin " + name + " " + err.Error())
		}
		runner, err := plgn.Lookup("Run")
		if err != nil {
			return errors.New("plugin " + name + " " + err.Error())
		}
		fn, ok := runner.(RunHandler)
		if !ok {
			return errors.New("plugin " + name + " doesn't implements the RunHandler(*sync.Map)")
		}
		fn(globalContext)
	}
	return nil
}
