package ingredient

import (
	"fmt"
	"github.com/gragas/jabberwock-lib/effect"
	"testing"
	"time"
)

func TestString(t *testing.T) {
	e := Ingredient{
		Name:             "Philosophus Tremens",
		Description:      "Philosophus tremens is a rare mushroom. It is known to grant extraordinary willpower to those who consume it.",
		Rarity:           Rare,
		DurationEffect:  effect.DurationEffect{Effect: effect.Effect{
			HealthModifier: 0,
			EnergyModifier: 0,
			SpiritModifier: 5,
			SummoningModifier: 0,
			AlterationModifier: 0,
			WillpowerModifier: 15.0,
			DivinityModifier: 0,
			LifebringingModifier: 0,}, Duration: time.Duration(3*60e9)},
		ImmediateEffect: effect.ImmediateEffect{
			Effect: effect.Effect{
				HealthModifier: 0,
				EnergyModifier: 0,
				SpiritModifier: 0,
				SummoningModifier: 0,
				AlterationModifier: 0,
				WillpowerModifier: 0,
				DivinityModifier: 0,
				LifebringingModifier: 0}}}
	fmt.Println(e)
}
