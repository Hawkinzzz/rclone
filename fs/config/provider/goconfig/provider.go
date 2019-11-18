package goconfig

import (
	"bytes"
	"io"
	"log"

	"github.com/Unknwon/goconfig"

	"github.com/rclone/rclone/fs/config"
)

func Register() {
	config.RegisterConfigProvider(&config.ProviderDefinition{
		NewFunc: NewGoConfigProvider,
		FileTypes: []string{"conf"},
	})
}

func NewGoConfigProvider() config.Provider {
	return &GoConfig{}
}

type GoConfig struct {
	config *goconfig.ConfigFile
}

func (g *GoConfig) String() string {
	buf := bytes.Buffer{}
	err := g.Save(&buf)
	if err != nil {
		log.Fatalf("error stringifying config: %v", err)
		return ""
	}
	return buf.String()
}

func (g *GoConfig) Load(r io.Reader) error {
	c, err := goconfig.LoadFromReader(r)
	if err != nil {
		return err
	}
	g.config = c
	return nil
}

func (g *GoConfig) Save(w io.Writer) error {
	return goconfig.SaveConfigData(g.config, w)
}

func (g *GoConfig) GetRemoteConfig() config.RemoteConfig {
	return g
}

var (
	_ config.Provider = (*GoConfig)(nil)
)
