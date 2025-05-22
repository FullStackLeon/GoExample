package common

// Attacker 定义公共接口
type Attacker interface {
	Attack(target Damageable)
}

type Damageable interface {
	TakeDamage(amount int)
}
