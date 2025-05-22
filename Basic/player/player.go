package player

type Player struct{ Health int }

func (p *Player) TakeDamage(amount int) { p.Health -= amount }
