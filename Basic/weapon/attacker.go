package weapon

import "GoLib/Basic/common"

type Weapon struct {
	Damage int
}

// target 接口化，可以传入任何实现了 common.Damageable 接口的对象
func (w *Weapon) Attack(target common.Damageable) {
	target.TakeDamage(w.Damage)
}
