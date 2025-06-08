package common

// Attacker 定义公共接口
type Attacker interface {
	Attack(target Damageable)
}

// Damageable 定义受伤接口
type Damageable interface {
	TakeDamage(amount int)
}
