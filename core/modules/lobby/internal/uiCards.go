// IMPORTANT NOTE
// THIS IS AI GENERATED
// THIS ISN'T THOUGHT THRU
package internal

import (
	"core/modules/lobby"
	"engine/modules/ecs"
	"engine/modules/inputs"
	"engine/modules/layout"
	"engine/modules/render"
	"engine/modules/text"
	"engine/modules/transform"
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/go-gl/mathgl/mgl32"
	"github.com/ogiusek/events"
)

const (
	lobbyCardWidth     float32 = 220
	lobbyControlWidth  float32 = 190
	lobbyControlHeight float32 = 30
	lobbyCardGap       float32 = 16
)

func (s *service) renderPlayerCards(parent ecs.EntityID) {
	cardCount := 1 + len(s.Player().Name().GetEntities())
	totalWidth := float32(cardCount)*lobbyCardWidth + float32(cardCount-1)*lobbyCardGap

	cards := s.World().NewEntity()
	s.Hierarchy().SetParent(cards, parent)
	s.Groups().InheritGroups(cards)
	s.Transform().Inherit().Set(cards, transform.NewInherit(transform.RelativePos))
	s.Transform().Size().Set(cards, transform.NewSize(totalWidth, 1, 1))

	s.Layout().Order().Set(cards, layout.NewOrder(layout.OrderHorizontal))
	s.Layout().Align().Set(cards, layout.NewAlign(.5, .5))
	s.Layout().Gap().Set(cards, layout.NewGap(lobbyCardGap))

	cardEntities := []ecs.EntityID{}
	for _, playerEntity := range s.Player().Name().GetEntities() {
		if card := s.renderPlayerCard(cards, playerEntity); card != 0 {
			cardEntities = append(cardEntities, card)
		}
	}

	var cardHeight float32
	for _, card := range cardEntities {
		cardHeight = max(cardHeight, s.lobbyCardContentHeight(card))
	}
	for _, card := range cardEntities {
		s.Transform().Size().Set(card, transform.NewSize(lobbyCardWidth, cardHeight, 1))
	}
	s.Transform().Size().Set(cards, transform.NewSize(totalWidth, cardHeight, 1))
}

func (s *service) lobbyCardContentHeight(card ecs.EntityID) float32 {
	gap, _ := s.Layout().Gap().Get(card)
	children := s.Hierarchy().Children(card).GetIndices()
	if len(children) == 0 {
		return 2 * gap.Gap
	}

	height := 2 * gap.Gap
	for i, child := range children {
		if i > 0 {
			height += gap.Gap
		}
		size, _ := s.Transform().Size().Get(child)
		height += size.Size[1]
	}
	return height
}

// menu:
// - choose config (map size, units, wallet)

func (s *service) renderPlayerCard(parent ecs.EntityID, playerEntity ecs.EntityID) ecs.EntityID {
	playerUUID, ok := s.UUID().Component().Get(playerEntity)
	if !ok {
		return 0
	}
	name, _ := s.Player().Name().Get(playerEntity)
	color, _ := s.Player().Color().Get(playerEntity)

	card := s.newLobbyCard(parent)

	s.newLobbyInput(card, name.Name, func(raw string) {
		newName := strings.TrimSpace(raw)
		if newName == "" {
			return
		}
		events.Emit(s.Events(), lobby.NewSetPlayerNameEvent(playerUUID.ID, newName))
	})

	swatch := s.newLobbySwatch(card, color.Color)
	s.newLobbyInput(card, formatColor(color.Color), func(raw string) {
		color, err := parseColor(raw)
		if err != nil {
			return
		}
		events.Emit(s.Events(), lobby.NewSetPlayerColorEvent(playerUUID.ID, color))
		if updated, ok := s.Player().Color().Get(playerEntity); ok {
			s.Render().Color().Set(swatch, render.NewColor(updated.Color.Vec4(1)))
		}
	})

	// TODO: show the connection assigned to this player once the lobby
	// knows how connections are mapped onto players.
	s.newLobbyLabel(card, "CONNECTION: -", lobbyControlWidth, 22, 16)

	s.newLobbyButton(card, "REMOVE", lobby.NewRemovePlayerEvent(playerUUID.ID), lobbyControlWidth, lobbyControlHeight)

	return card
}

func (s *service) newLobbyCard(parent ecs.EntityID) ecs.EntityID {
	card := s.World().NewEntity()
	s.Hierarchy().SetParent(card, parent)
	s.Groups().InheritGroups(card)
	s.Transform().Size().Set(card, transform.NewSize(lobbyCardWidth, 1, 1))
	s.Transform().Inherit().Set(card, transform.NewInherit(transform.RelativePos))
	s.Transform().AspectRatio().Set(card, transform.NewAspectRatio(0, 0, 0, 0))

	s.Render().Mesh().Set(card, render.NewMesh(s.Definitions().Assets().SquareMesh))
	s.Render().Texture().Set(card, render.NewTexture(s.Definitions().Assets().Blank))
	s.Render().Color().Set(card, render.NewColor(mgl32.Vec4{.08, .08, .08, 1}))

	s.Layout().Order().Set(card, layout.NewOrder(layout.OrderVectical))
	s.Layout().Align().Set(card, layout.NewAlign(.5, .5))
	s.Layout().Gap().Set(card, layout.NewGap(8))

	return card
}

func (s *service) newLobbyLabel(parent ecs.EntityID, content string, width, height float32, fontSize uint) ecs.EntityID {
	label := s.World().NewEntity()
	s.Hierarchy().SetParent(label, parent)
	s.Groups().InheritGroups(label)
	s.Transform().Size().Set(label, transform.NewSize(width, height, 1))
	s.Transform().Inherit().Set(label, transform.NewInherit(transform.RelativePos))
	s.Transform().AspectRatio().Set(label, transform.NewAspectRatio(0, 0, 0, 0))
	s.Transform().MaxSize().Set(label, transform.NewMaxSize(0, 0, 0))
	s.Transform().MinSize().Set(label, transform.NewMinSize(0, 0, 0))

	s.Text().Content().Set(label, text.NewText(content))
	s.Text().Color().Set(label, text.NewColor(mgl32.Vec4{1, 1, 1, 1}))
	s.Text().FontSize().Set(label, text.NewFontSize(fontSize))
	s.Text().Align().Set(label, text.NewAlign(.5, .5))

	return label
}

func (s *service) newLobbyInput(parent ecs.EntityID, initial string, onChange func(string)) ecs.EntityID {
	input := s.Prototype().Clone(s.Definitions().Hud().Input)
	s.Hierarchy().SetParent(input, parent)
	s.Transform().Size().Set(input, transform.NewSize(lobbyControlWidth, lobbyControlHeight, 1))
	s.Transform().Inherit().Set(input, transform.NewInherit(transform.RelativePos))
	s.Transform().AspectRatio().Set(input, transform.NewAspectRatio(0, 0, 0, 0))
	s.Transform().MaxSize().Set(input, transform.NewMaxSize(0, 0, 0))
	s.Transform().MinSize().Set(input, transform.NewMinSize(0, 0, 0))

	s.Text().Content().Set(input, text.NewText(initial))
	s.Text().FontSize().Set(input, text.NewFontSize(18))

	s.Inputs().TextInput().Set(input, inputs.TextInputComponent{
		Lines:     []string{initial},
		CursorRow: 0,
		CursorCol: len(initial),
	})
	if onChange != nil {
		s.inputChanges[input] = onChange
	}

	return input
}

func (s *service) onTextInputMod(entity ecs.EntityID) {
	onChange, ok := s.inputChanges[entity]
	if !ok {
		return
	}
	input, ok := s.Inputs().TextInput().Get(entity)
	if !ok {
		delete(s.inputChanges, entity)
		return
	}
	onChange(input.Text())
}

func (s *service) newLobbyButton(parent ecs.EntityID, label string, onClick any, width, height float32) ecs.EntityID {
	btn := s.Prototype().Clone(s.Definitions().Hud().Btn)
	s.Hierarchy().SetParent(btn, parent)
	s.Transform().Size().Set(btn, transform.NewSize(width, height, 1))
	s.Transform().Inherit().Set(btn, transform.NewInherit(transform.RelativePos))
	s.Transform().AspectRatio().Set(btn, transform.NewAspectRatio(0, 0, 0, 0))
	s.Transform().MaxSize().Set(btn, transform.NewMaxSize(0, 0, 0))
	s.Transform().MinSize().Set(btn, transform.NewMinSize(0, 0, 0))

	s.Inputs().LeftClick().Set(btn, inputs.NewLeftClick(onClick))
	s.Text().Content().Set(btn, text.NewText(strings.ToUpper(label)))
	s.Text().FontSize().Set(btn, text.NewFontSize(16))

	return btn
}

func (s *service) newLobbySwatch(parent ecs.EntityID, color mgl32.Vec3) ecs.EntityID {
	swatch := s.World().NewEntity()
	s.Hierarchy().SetParent(swatch, parent)
	s.Groups().InheritGroups(swatch)
	s.Transform().Size().Set(swatch, transform.NewSize(lobbyControlWidth, 20, 1))
	s.Transform().Inherit().Set(swatch, transform.NewInherit(transform.RelativePos))
	s.Transform().AspectRatio().Set(swatch, transform.NewAspectRatio(0, 0, 0, 0))

	s.Render().Mesh().Set(swatch, render.NewMesh(s.Definitions().Assets().SquareMesh))
	s.Render().Texture().Set(swatch, render.NewTexture(s.Definitions().Assets().Blank))
	s.Render().Color().Set(swatch, render.NewColor(color.Vec4(1)))

	return swatch
}

func formatColor(color mgl32.Vec3) string {
	toByte := func(value float32) uint8 {
		value = max(0, min(1, value))
		return uint8(math.Round(float64(value) * 255))
	}
	return fmt.Sprintf("%02X%02X%02X", toByte(color.X()), toByte(color.Y()), toByte(color.Z()))
}

func parseColor(raw string) (mgl32.Vec3, error) {
	raw = strings.TrimPrefix(strings.TrimSpace(raw), "#")
	if len(raw) != 6 {
		return mgl32.Vec3{}, fmt.Errorf("color: expected RRGGBB but got %q", raw)
	}
	value, err := strconv.ParseUint(raw, 16, 24)
	if err != nil {
		return mgl32.Vec3{}, err
	}
	return mgl32.Vec3{
		float32(value>>16&0xff) / 255,
		float32(value>>8&0xff) / 255,
		float32(value&0xff) / 255,
	}, nil
}
