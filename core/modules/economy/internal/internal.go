package internal

import (
	"core/game"
	"core/modules/definitions"
	"core/modules/economy"
	"engine/modules/ecs"
	"engine/modules/groups"
	"engine/modules/loop"
	"engine/modules/text"
	"engine/modules/transform"
	"fmt"

	"github.com/ogiusek/events"
	"github.com/ogiusek/ioc/v2"
)

type walletPreview struct{}

type service struct {
	game.GameWorld `inject:""`
	wallet         ecs.ComponentArray[economy.WalletComponent]
	factory        ecs.ComponentArray[economy.FactoryComponent]
	cost           ecs.ComponentArray[economy.CostComponent]

	walletPreview ecs.ComponentArray[walletPreview]
}

func NewService(c ioc.Dic) economy.Service {
	s := ioc.GetServices[*service](c)
	s.wallet = ecs.GetComponentArray[economy.WalletComponent](s.World())
	s.factory = ecs.GetComponentArray[economy.FactoryComponent](s.World())
	s.cost = ecs.GetComponentArray[economy.CostComponent](s.World())

	s.walletPreview = ecs.GetComponentArray[walletPreview](s.World())
	return s
}

func (s *service) Wallet() ecs.ComponentArray[economy.WalletComponent]   { return s.wallet }
func (s *service) Factory() ecs.ComponentArray[economy.FactoryComponent] { return s.factory }
func (s *service) Cost() ecs.ComponentArray[economy.CostComponent]       { return s.cost }

func (s *service) WalletPreview() ecs.EntityID {
	entities := s.walletPreview.GetEntities()
	if len(entities) == 0 {
		entity := s.World().NewEntity()
		s.walletPreview.Set(entity, walletPreview{})
		return entity
	}
	for _, entity := range entities[1:] {
		s.World().RemoveEntity(entity)
	}
	return entities[0]
}

func (s *service) Register() error {
	events.Listen(s.EventsBuilder(), s.OnFrame)
	events.Listen(s.EventsBuilder(), s.OnTick)
	return nil
}

func (s *service) OnFrame(loop.FrameEvent) {
	entity := s.WalletPreview()

	uiCameraEntities := s.Ui().UiCamera().GetEntities()
	if len(uiCameraEntities) == 0 {
		s.World().RemoveEntity(entity)
		return
	}
	uiCamera := uiCameraEntities[0]
	var wallet economy.WalletComponent
	var ok bool
	for _, entity := range s.Player().ActingPlayer().GetEntities() {
		wallet, ok = s.Wallet().Get(entity)
		if ok {
			break
		}
	}
	if !ok {
		s.World().RemoveEntity(entity)
		return
	}

	s.Hierarchy().SetParent(entity, uiCamera)
	s.Transform().Pos().Set(entity, transform.NewPos(-10, 0, 0))
	s.Transform().PivotPoint().Set(entity, transform.NewPivotPoint(1, 1, .5))
	s.Transform().Inherit().Set(entity, transform.NewInherit(transform.RelativePos))
	s.Transform().ParentPivotPoint().Set(entity, transform.NewParentPivotPoint(1, 1, .5))
	s.Groups().Component().Set(entity, groups.EmptyGroups().Enable(definitions.UiGroup))

	s.Text().Content().Set(entity, text.NewText(fmt.Sprintf("%d$", wallet.Money)))
	s.Text().Break().Set(entity, text.NewBreak(text.BreakNone))
	s.Text().Align().Set(entity, text.NewAlign(1, 0))
	s.Text().FontSize().Set(entity, text.NewFontSize(30))
}

func (s *service) OnTick(tick loop.TickEvent) {
	for _, factoryEntity := range s.factory.GetEntities() {
		factoryComp, ok := s.Economy().Factory().Get(factoryEntity)
		if !ok {
			continue
		}
		factoryOwner, ok := s.Player().Owner().Get(factoryEntity)
		if !ok {
			continue
		}
		wallet, _ := s.Economy().Wallet().Get(factoryOwner.Owner)

		moneyGain := float64(factoryComp.MoneyPerSecond) * tick.Delta.Seconds()
		wallet.Money += economy.Money(moneyGain)
		s.Economy().Wallet().Set(factoryOwner.Owner, wallet)
	}
}
