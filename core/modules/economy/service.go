package economy

import (
	"engine/modules/ecs"
	"errors"
)

var (
	ErrToExpensive = errors.New("to expensive")
)

type Money uint32

type WalletComponent struct{ Money Money }
type FactoryComponent struct{ MoneyPerSecond Money }
type CostComponent struct{ Cost Money }

func NewWallet(money Money) WalletComponent   { return WalletComponent{money} }
func NewFactory(money Money) FactoryComponent { return FactoryComponent{money} }
func NewCost(money Money) CostComponent       { return CostComponent{money} }

func (comp WalletComponent) Pay(cost CostComponent) WalletComponent {
	return NewWallet(comp.Money - cost.Cost)
}

//

func (WalletComponent) Smooth() {}
func (c1 WalletComponent) Lerp(c2 WalletComponent, mix32 float32) WalletComponent {
	return WalletComponent{Money(float32(c1.Money)*(1-mix32) + float32(c2.Money)*mix32)}
}

//

type Service interface {
	ecs.SystemRegister
	Wallet() ecs.ComponentArray[WalletComponent]
	Factory() ecs.ComponentArray[FactoryComponent]
	Cost() ecs.ComponentArray[CostComponent]
}
