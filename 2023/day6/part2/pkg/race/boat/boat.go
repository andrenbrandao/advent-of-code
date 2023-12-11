package boat

import "day6/pkg/race/types"

type Boat struct {
	charge types.Charge
}

func NewBoat(charge types.Charge) *Boat {
	return &Boat{charge}
}

func (b *Boat) Distance(time types.Time) types.Distance {
	distance := int(time) * int(b.Velocity())
	return types.Distance(distance)
}

func (b *Boat) Velocity() types.Velocity {
	return types.Velocity(b.charge)
}
