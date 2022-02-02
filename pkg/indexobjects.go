package indexobjects

import (
	"context"
	"fmt"

	"github.com/operator-framework/operator-registry/alpha/action"
	"github.com/operator-framework/operator-registry/alpha/declcfg"
)

func processRefs(refs []string, filter []string) (*declcfg.DeclarativeConfig, error) {
	var result *declcfg.DeclarativeConfig
	render := action.Render{
		//Refs:           []string{"registry.redhat.io/redhat/redhat-operator-index:v4.9"},
		Refs:           refs,
		Registry:       nil,
		AllowedRefMask: 0,
	}
	config, err := render.Run(context.TODO())
	if err != nil {
		return result, err
	}
	for _, item := range config.Channels {
		fmt.Println(">>>>>>>>>")
		fmt.Printf("%s - %s\n", item.Name, item.Package)
		fmt.Printf("%v\n", item.Entries)
	}
	return result, nil
}
