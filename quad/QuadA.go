package piscine

import "github.com/01-edu/z01"

func QuadA(x, y int) {
	if x > 0 && y > 0 {
		for j := 1; j <= y; j++ {
			for i := 1; i <= x; i++ {
				if (i == 1 && j == 1) || (i == x && j == y) ||
					(i == x && j == 1) || (i == 1 && j == y) {
					z01.PrintRune('o')
				} else {
					if i == 1 || i == x && (j != 1 || j != y) {
						z01.PrintRune('|')
					} else {
						if (j == 1 || j == y) && !(i == 1 || i == x) {
							z01.PrintRune('-')
						} else {
							z01.PrintRune(' ')
						}
					}
				}
			}
			z01.PrintRune('\n')
		}
	}
}
