package piscine

func UltimateDivMod(a *int, b *int) {
	mod := *a
	*a = *a / *b
	*b = mod % *b
}
