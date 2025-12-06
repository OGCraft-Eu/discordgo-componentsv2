package v2

import (
	"encoding/json"

	"github.com/bwmarrin/discordgo"
)

const (
	ComponentTypeContainer    discordgo.ComponentType = 17
	ComponentTypeMediaGallery discordgo.ComponentType = 12
	ComponentTypeTextDisplay  discordgo.ComponentType = 10
	ComponentTypeSection      discordgo.ComponentType = 9
)

type TextDisplay struct {
	Content string `json:"content"`
}

func (t TextDisplay) Type() discordgo.ComponentType {
	return ComponentTypeTextDisplay
}

func (t TextDisplay) MarshalJSON() ([]byte, error) {
	type textDisplay TextDisplay
	return json.Marshal(struct {
		textDisplay
		Type discordgo.ComponentType `json:"type"`
	}{
		textDisplay: textDisplay(t),
		Type:        t.Type(),
	})
}

func NewTextDisplayBuilder() *TextDisplay {
	return &TextDisplay{}
}

func (t *TextDisplay) SetContent(content string) *TextDisplay {
	t.Content = content
	return t
}

func (t *TextDisplay) Build() discordgo.MessageComponent {
	return *t
}

type MediaGallery struct {
	ID    int                `json:"id,omitempty"`
	Items []MediaGalleryItem `json:"items"`
}

type MediaGalleryItem struct {
	Media       UnfurledMediaItem `json:"media"`
	Description *string           `json:"description,omitempty"`
	Spoiler     bool              `json:"spoiler"`
}

type UnfurledMediaItem struct {
	URL string `json:"url"`
}

func (m MediaGallery) Type() discordgo.ComponentType {
	return ComponentTypeMediaGallery
}

func (m MediaGallery) MarshalJSON() ([]byte, error) {
	type mediaGallery MediaGallery
	return json.Marshal(struct {
		mediaGallery
		Type discordgo.ComponentType `json:"type"`
	}{
		mediaGallery: mediaGallery(m),
		Type:         m.Type(),
	})
}

func NewMediaGalleryBuilder() *MediaGallery {
	return &MediaGallery{
		Items: []MediaGalleryItem{},
	}
}

func (m *MediaGallery) AddImageURL(url string) *MediaGallery {
	m.Items = append(m.Items, MediaGalleryItem{
		Media: UnfurledMediaItem{
			URL: url,
		},
		Spoiler: false,
	})
	return m
}

func (m *MediaGallery) Build() discordgo.MessageComponent {
	return *m
}

type Container struct {
	ID          int                          `json:"id,omitempty"`
	AccentColor *int                         `json:"accent_color,omitempty"`
	Spoiler     bool                         `json:"spoiler"`
	Components  []discordgo.MessageComponent `json:"components"`
}

func (c Container) Type() discordgo.ComponentType {
	return ComponentTypeContainer
}

func (c Container) MarshalJSON() ([]byte, error) {
	type container Container
	return json.Marshal(struct {
		container
		Type discordgo.ComponentType `json:"type"`
	}{
		container: container(c),
		Type:      c.Type(),
	})
}

func NewContainerBuilder() *Container {
	return &Container{
		Components: []discordgo.MessageComponent{},
	}
}

func (c *Container) AddComponent(comp discordgo.MessageComponent) *Container {
	c.Components = append(c.Components, comp)
	return c
}

func (c *Container) SetAccentColor(color int) *Container {
	c.AccentColor = &color
	return c
}

func (c *Container) Build() discordgo.MessageComponent {
	return *c
}
