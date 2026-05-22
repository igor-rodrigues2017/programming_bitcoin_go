package criptography

func Mod(a, b int64) int64 {
	return ((a % b) + b) % b
}
