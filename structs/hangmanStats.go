package structs

func HangmanStats(life int) (int, int, bool) {
	var start int
	var end int
	status := true
	switch life {
	case 10:
		status = true
	case 9:
		start = 0
		end = 69
	case 8:
		start = 71
		end = 140
	case 7:
		start = 142
		end = 211
	case 6:
		start = 213
		end = 282
	case 5:
		start = 284
		end = 353
	case 4:
		start = 355
		end = 424
	case 3:
		start = 426
		end = 495
	case 2:
		start = 497
		end = 566
	case 1:
		start = 568
		end = 637
	case 0:
		start = 639
		end = 708
	}
	return start, end, status
}
