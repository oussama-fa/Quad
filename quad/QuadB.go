package piscine

import "github.com/01-edu/z01"

func QuadB(x, y int) {
	if x > 0 && y > 0 {
		for j := 1; j <= y; j++ {
			for i := 1; i <= x; i++ {
				if (i == 1 && j == 1) || ((i == x && j == y) {//&& (x != 1 && y != 1)) {
					z01.PrintRune('/')
				} else if (i == x && j == 1) || (i == 1 && j == y) {
					z01.PrintRune('\\')
				} else {
					if (i == 1 || i == x) || (j == 1 || j == y) {
						z01.PrintRune('*')
					} else {
						z01.PrintRune(' ')
					}
				}
			}
			z01.PrintRune('\n')
		}
	}
}

