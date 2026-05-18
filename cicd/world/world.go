package world

import (
	"cicd/modules/docs"
	"cicd/modules/git"
	"cicd/modules/pipe"
	"cicd/modules/projectfs"

	"github.com/ogiusek/ioc/v2"
)

type CICDWorld struct {
	Docs      ioc.Lazy[docs.Service]      `inject:""`
	Git       ioc.Lazy[git.Service]       `inject:""`
	ProjectFS ioc.Lazy[projectfs.Service] `inject:""`
	Pipe      ioc.Lazy[pipe.Service]      `inject:""`
}
