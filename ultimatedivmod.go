package piscine

func UltimatedivMod(a *int, b *int) {
	A := *a / *b
	B := *a % *b
	*a = A
	*b = B
}
