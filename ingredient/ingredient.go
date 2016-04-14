package ingredient

import (
	"encoding/json"
	"github.com/gragas/jabberwock-lib/attributes"
	"github.com/gragas/jabberwock-lib/effect"
	"io/ioutil"
)

type Rarity int

const (
	Common = iota
	Uncommon
	Rare
	UltraRare
	Legendary
)

type Ingredient struct {
	Name            string
	Description     string
	Rarity          Rarity
	DurationEffect  effect.DurationEffect
	ImmediateEffect effect.ImmediateEffect
}

func (i Ingredient) Bytes() []byte {
	bytes, err := json.MarshalIndent(i, "", "\t")
	if err != nil {
		panic(err)
	}
	return bytes
}

func (i Ingredient) String() string {
	return string(i.Bytes())
}

func FromFile(filename string) *Ingredient {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return FromBytes(bytes)
}

func FromBytes(bytes []byte) *Ingredient {
	i := new(Ingredient)
	err := json.Unmarshal(bytes, i)
	if err != nil {
		panic(err)
	}
	return i
}

type PrettyIngredient struct {
	Name             string
	Description      string
	Rarity           string
	DurationEffects  []string
	ImmediateEffects []string
	Duration         string
}

func (i Ingredient) PrettyIngredient() *PrettyIngredient {
	p := new(PrettyIngredient)
	p.Name = i.Name
	p.Description = i.Description
	p.Rarity = i.Rarity.String()
	p.Duration = i.DurationEffect.Duration.String()
	if i.DurationEffect.HealthModifier != 0 {
		p.DurationEffects = append(p.DurationEffects,
			"Health: "+attributes.AttributeString(i.DurationEffect.HealthModifier))
	}
	if i.DurationEffect.EnergyModifier != 0 {
		p.DurationEffects = append(p.DurationEffects,
			"Energy: "+attributes.AttributeString(i.DurationEffect.EnergyModifier))
	}
	if i.DurationEffect.SpiritModifier != 0 {
		p.DurationEffects = append(p.DurationEffects,
			"Spirit: "+attributes.AttributeString(i.DurationEffect.SpiritModifier))
	}
	if i.DurationEffect.SummoningModifier != 0 {
		p.DurationEffects = append(p.DurationEffects,
			"Summoning: "+attributes.AttributeString(i.DurationEffect.SummoningModifier))
	}
	if i.DurationEffect.AlterationModifier != 0 {
		p.DurationEffects = append(p.DurationEffects,
			"Alteration: "+attributes.AttributeString(i.DurationEffect.AlterationModifier))
	}
	if i.DurationEffect.WillpowerModifier != 0 {
		p.DurationEffects = append(p.DurationEffects,
			"Willpower: "+attributes.AttributeString(i.DurationEffect.WillpowerModifier))
	}
	if i.DurationEffect.DivinityModifier != 0 {
		p.DurationEffects = append(p.DurationEffects,
			"Divinity: "+attributes.AttributeString(i.DurationEffect.DivinityModifier))
	}
	if i.DurationEffect.LifebringingModifier != 0 {
		p.DurationEffects = append(p.DurationEffects,
			"Lifebringing: "+attributes.AttributeString(i.DurationEffect.LifebringingModifier))
	}
	if i.ImmediateEffect.HealthModifier != 0 {
		p.ImmediateEffects = append(p.ImmediateEffects,
			"Health: "+attributes.AttributeString(i.ImmediateEffect.HealthModifier))
	}
	if i.ImmediateEffect.EnergyModifier != 0 {
		p.ImmediateEffects = append(p.ImmediateEffects,
			"Energy: "+attributes.AttributeString(i.ImmediateEffect.EnergyModifier))
	}
	if i.ImmediateEffect.SpiritModifier != 0 {
		p.ImmediateEffects = append(p.ImmediateEffects,
			"Spirit: "+attributes.AttributeString(i.ImmediateEffect.SpiritModifier))
	}
	if i.ImmediateEffect.SummoningModifier != 0 {
		p.ImmediateEffects = append(p.ImmediateEffects,
			"Summoning: "+attributes.AttributeString(i.ImmediateEffect.SummoningModifier))
	}
	if i.ImmediateEffect.AlterationModifier != 0 {
		p.ImmediateEffects = append(p.ImmediateEffects,
			"Alteration: "+attributes.AttributeString(i.ImmediateEffect.AlterationModifier))
	}
	if i.ImmediateEffect.WillpowerModifier != 0 {
		p.ImmediateEffects = append(p.ImmediateEffects,
			"Willpower: "+attributes.AttributeString(i.ImmediateEffect.WillpowerModifier))
	}
	if i.ImmediateEffect.DivinityModifier != 0 {
		p.ImmediateEffects = append(p.ImmediateEffects,
			"Divinity: "+attributes.AttributeString(i.ImmediateEffect.DivinityModifier))
	}
	if i.ImmediateEffect.LifebringingModifier != 0 {
		p.ImmediateEffects = append(p.ImmediateEffects,
			"Lifebringing: "+attributes.AttributeString(i.ImmediateEffect.LifebringingModifier))
	}
	return p
}

func (r Rarity) String() string {
	switch r {
	case Common:
		return "Common"
	case Uncommon:
		return "Uncommon"
	case Rare:
		return "Rare"
	case UltraRare:
		return "Ultra Rare"
	case Legendary:
		return "Legendary"
	default:
		return "No such rarity."
	}
}
