package rate

import "fmt"

type Limiter struct {
	Accounts map[string]Rate
}

func (l *Limiter) Acquire(account string, credits int) (int, error) {
	if rate, ok := l.Accounts[account]; ok {
		return rate.Acquire(credits), nil
	}
	return 0, fmt.Errorf("no account named %s", account)
}
