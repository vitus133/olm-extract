package indexobjects

import (
	"context"

	"github.com/operator-framework/operator-registry/alpha/action"
	"github.com/operator-framework/operator-registry/alpha/declcfg"
	"github.com/operator-framework/operator-registry/alpha/model"
)

func ProcessRefs(refs []string, filter []string) (model.Model, error) {

	render := action.Render{
		//Refs:           []string{"registry.redhat.io/redhat/redhat-operator-index:v4.9"},
		Refs:           refs,
		Registry:       nil,
		AllowedRefMask: 0,
	}
	config, err := render.Run(context.TODO())
	if err != nil {
		return model.Model{}, err
	}

	return filterConfig(config, filter)
}

func filterConfig(config *declcfg.DeclarativeConfig, filter []string) (model.Model, error) {
	var m = model.Model{}
	fullModel, err := declcfg.ConvertToModel(*config)
	if err != nil {
		return m, err
	}
	for _, item := range filter {
		m[item] = fullModel[item]
	}
	return m, nil
}
