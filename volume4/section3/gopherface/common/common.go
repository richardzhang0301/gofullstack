package common

import (
	"github.com/richardzhang0301/gofullstack/volume4/section3/gopherface/common/datastore"
	"github.com/richardzhang0301/isokit"
)

type Env struct {
	DB          datastore.Datastore
	TemplateSet *isokit.TemplateSet
}
