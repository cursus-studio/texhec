package clicksystems

import (
	"core/modules/definitions"
	"core/modules/deploy"
	"core/modules/tile"
	"core/modules/ui"
	"engine"
	"engine/modules/assets"
	"engine/modules/collider"
	"engine/modules/grid"
	"engine/modules/groups"
	"engine/modules/inputs"
	"engine/modules/render"
	"engine/modules/text"
	"engine/modules/transform"
	"engine/services/ecs"
	"errors"
	"fmt"

	"github.com/ogiusek/events"
	"github.com/ogiusek/ioc/v2"
)

type system struct {
	engine.World `inject:"1"`
	Definitions  definitions.Definitions `inject:"1"`
	Deploy       deploy.Service          `inject:"1"`
	Tile         tile.Service            `inject:"1"`
	Ui           ui.Service              `inject:"1"`
}

func NewSystems(c ioc.Dic) ecs.SystemRegister {
	return ecs.NewSystemRegister(func() error {
		s := ioc.GetServices[*system](c)

		events.Listen(s.EventsBuilder, s.OnClickEntity)

		return nil
	})
}

func (s *system) OnClickEntity(e tile.ClickEntityEvent) {
	pos, ok := s.Tile.Pos().Get(e.Entity)
	if !ok {
		s.Logger.Warn(errors.New("expected entity to have coords component"))
		return
	}
	link, ok := s.Deploy.Link().Get(e.Entity)
	if !ok {
		s.Logger.Warn(errors.New("expected entity to have link component"))
		return
	}
	deployed, ok := s.Deploy.Component().Get(link.Deploy)
	if !ok {
		s.Logger.Warn(errors.New("expected link to have deploy component"))
		return
	}

	type Button struct {
		text  string
		event any
	}
	btns := []Button{
		{fmt.Sprintf("OBJECT POS: %v", pos), nil},
	}
	for _, deployed := range deployed.Deployable {
		btns = append(btns, Button{fmt.Sprintf("DEPLOY: %v", deployed), deploy.NewDeployEvent(deployed, grid.NewCoords(3, 3))})
	}

	btnAsset, err := assets.GetAsset[render.TextureAsset](s.Assets, s.Definitions.Hud.Btn)
	if err != nil {
		s.Logger.Warn(err)
		return
	}
	btnAspectRatio := btnAsset.AspectRatio()

	for _, p := range s.Ui.Show() {
		// i want here to display all actions which can be performed by entity
		// currently implement only building
		for _, btn := range btns {
			btnEntity := s.NewEntity()
			s.Hierarchy.SetParent(btnEntity, p)
			s.Groups.Inherit().Set(btnEntity, groups.InheritGroupsComponent{})

			s.Transform.AspectRatio().Set(btnEntity, transform.NewAspectRatio(float32(btnAspectRatio.Dx()), float32(btnAspectRatio.Dy()), 0, transform.PrimaryAxisX))
			s.Transform.Parent().Set(btnEntity, transform.NewParent(transform.RelativePos|transform.RelativeSizeX))
			s.Transform.MaxSize().Set(btnEntity, transform.NewMaxSize(0, 50, 0))
			s.Transform.Size().Set(btnEntity, transform.NewSize(1, 50, 1))

			s.Render.Mesh().Set(btnEntity, render.NewMesh(s.Definitions.SquareMesh))
			s.Render.Texture().Set(btnEntity, render.NewTexture(s.Definitions.Hud.Btn))

			s.Text.Content().Set(btnEntity, text.TextComponent{Text: btn.text})
			s.Text.FontSize().Set(btnEntity, text.FontSizeComponent{FontSize: 20})
			s.Text.Align().Set(btnEntity, text.TextAlignComponent{Vertical: .5, Horizontal: .5})

			s.Inputs.LeftClick().Set(btnEntity, inputs.NewLeftClick(btn.event))
			s.Inputs.KeepSelected().Set(btnEntity, inputs.KeepSelectedComponent{})
			s.Collider.Component().Set(btnEntity, collider.NewCollider(s.Definitions.SquareCollider))
		}
	}
}
