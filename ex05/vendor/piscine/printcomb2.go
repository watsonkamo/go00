package piscine

import "ft"

func PrintComb2() {
	for i := 0; i <= 99; i++ {
		for j := i + 1; j <= 99; j++ {
			if i < 10 {
				ft.PrintRune('0')
			}
			PrintNbr(i)
			ft.PrintRune(' ')
			if j < 10 {
				ft.PrintRune('0')
			}
			PrintNbr(j)
			if i < 98 {
				ft.PrintRune(',')
				ft.PrintRune(' ')
			}
		}
	}
	ft.PrintRune('\n')
}

func PrintNbr(nb int) {
	if nb == 0 {
		ft.PrintRune('0')
		return
	}
	if nb < 0 {
		ft.PrintRune('-')
		nb = -nb
	}
	if nb < 10 {
		ft.PrintRune(rune(nb) + '0')
		return
	}
	PrintNbr(nb / 10)
	PrintNbr(nb % 10)
}
