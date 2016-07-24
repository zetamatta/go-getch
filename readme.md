getch
=====

`getch` is the library for the programming language Go for Windows,
to get key from console.

	fmt.Print("Hit any key: ")
	rune1, scan1, shift1 := getch.Full()

	fmt.Printf("\n%c %08X %08X %08X\n", rune1, rune1, scan1, shift1)

	fmt.Print("Hit any key: ")
	rune1 = getch.Rune()
	fmt.Printf("\n%c %08X\n", rune1, rune1)