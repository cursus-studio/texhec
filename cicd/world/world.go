package world

import (
	"cicd/modules/docs"
	"cicd/modules/git"
	"cicd/modules/hooks"

	"github.com/ogiusek/ioc/v2"
)

type CICDWorld struct {
	Git   ioc.Lazy[git.Service]   `inject:""`
	Docs  ioc.Lazy[docs.Service]  `inject:""`
	Hooks ioc.Lazy[hooks.Service] `inject:""`
}
