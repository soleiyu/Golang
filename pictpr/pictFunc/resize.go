package pictFunc

func Xhsample(a Pict) Pict {
	res := MkPict(a.Width / 2, a.Height / 2)

	for x := 0; x < res.Width; x++ {
		for y := 0; y < res.Height; y++ {
			res.Px[x][y][0] = 255

			for i := 1; i < 4; i++ {
					res.Px[x][y][i] = 
						a.Px[2 * x][2 * y][i] / 2 + a.Px[2 * x + 1][2 * y + 1][i] / 2
				}
			}
		}

	return res
}

func X2sample(a Pict) Pict {
	res := MkPict(a.Width * 2, a.Height * 2)

	for x := 0; x < res.Width; x++ {
		for y := 0; y < res.Height; y++ {
			res.Px[x][y][0] = 255

			for i := 1; i < 4; i++ {
					res.Px[x][y][i] = a.Px[x / 2][y / 2][i]
			}	
		}
	}

	return res
}

