package boat

type Charge int
type Time int
type Distance int
type Velocity int

type Boat struct {
	charge Charge
}

func NewBoat(charge Charge) *Boat {
	return &Boat{charge}
}

func (b *Boat) Distance(time Time) Distance {
	distance := int(time) * int(b.Velocity())
	return Distance(distance)
}

func (b *Boat) Velocity() Velocity {
	return Velocity(b.charge)
}
