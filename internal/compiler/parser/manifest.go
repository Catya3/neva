package parser

import (
	"fmt"

	yaml "gopkg.in/yaml.v3"

	src "github.com/nevalang/neva/pkg/sourcecode"
)

func (p Parser) ParseManifest(raw []byte) (src.ModuleManifest, error) {
	var manifest src.ModuleManifest
	if err := yaml.Unmarshal(raw, &manifest); err != nil {
		return src.ModuleManifest{}, fmt.Errorf("yaml unmarshal: %w", err)
	}
	return processParsedManifest(manifest), nil
}

func processParsedManifest(manifest src.ModuleManifest) src.ModuleManifest {
	deps := make(map[string]src.ModuleRef, len(manifest.Deps))

	for alias, dep := range manifest.Deps {
		path := dep.Path
		if path == "" {
			path = alias
		}
		deps[alias] = src.ModuleRef{
			Path:    path,
			Version: dep.Version,
		}
	}

	return src.ModuleManifest{
		LanguageVersion: manifest.LanguageVersion,
		Deps:            deps,
	}
}
