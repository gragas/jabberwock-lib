package effect

import (
	"github.com/gragas/jabberwock-lib/attributes"
	"time"
)

type Effect struct {
	HealthModifier       attributes.Health
	EnergyModifier       attributes.Energy
	SpiritModifier       attributes.Spirit
	SummoningModifier    attributes.Summoning
	AlterationModifier   attributes.Alteration
	WillpowerModifier    attributes.Willpower
	DivinityModifier     attributes.Divinity
	LifebringingModifier attributes.Lifebringing
}

type ImmediateEffect struct {
	Effect
}

type DurationEffect struct {
	Effect
	Duration time.Duration
}
