package goconfig

import (
	"log"

	"github.com/rclone/rclone/fs/config"
)

func (g *GoConfig) ListRemotes() []string {
	return g.config.GetSectionList()
}

func (g *GoConfig) HasRemote(remote string) bool {
	for _, v := range g.ListRemotes() {
		if v == remote {
			return true
		}
	}
	return false
}

func (g *GoConfig) GetRemote(remote string) config.Section {
	return newSection(g.config, remote)
}

func (g *GoConfig) CreateRemote(remote string) config.Section {
	g.config.SetValue(remote, "", "")
	return g.GetRemote(remote)
}

func (g *GoConfig) DeleteRemote(name string) {
	g.config.DeleteSection(name)
}

func (g *GoConfig) RenameRemote(oldName string, newName string) {
	g.CopyRemote(oldName, newName)
	g.config.DeleteSection(oldName)
}

func (g *GoConfig) CopyRemote(source string, destination string) {
	data, err := g.config.GetSection(source)
	if err != nil {
		log.Fatalf("couldnt load section: %s", err)
	}

	for k, v := range data {
		g.config.SetValue(destination, k, v)
	}
}

var (
	_ config.RemoteConfig = (*GoConfig)(nil)
)
