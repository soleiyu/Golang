package pictFunc

func SliceRGBY(inp Pict) Pict {
	res := MkPict(inp.Width, inp.Height * 4)

	for x := 0; x < res.Width; x++ {
		for y := 0; y < inp.Height; y++ {
			// TOP
			res.Px[x][y][0] = inp.Px[x][y][0]
			res.Px[x][y][1] = inp.Px[x][y][1]
			res.Px[x][y][2] = inp.Px[x][y][2]
			res.Px[x][y][3] = inp.Px[x][y][3]

			// R
			res.Px[x][inp.Height + y][0] = inp.Px[x][y][0]
			res.Px[x][inp.Height + y][1] = inp.Px[x][y][1]
			res.Px[x][inp.Height + y][2] = 0
			res.Px[x][inp.Height + y][3] = 0

			// G
			res.Px[x][2 * inp.Height + y][0] = inp.Px[x][y][0]
			res.Px[x][2 * inp.Height + y][2] = inp.Px[x][y][2]
			res.Px[x][2 * inp.Height + y][1] = 0
			res.Px[x][2 * inp.Height + y][3] = 0

			// B
			res.Px[x][3 * inp.Height + y][0] = inp.Px[x][y][0]
			res.Px[x][3 * inp.Height + y][3] = inp.Px[x][y][3]
			res.Px[x][3 * inp.Height + y][2] = 0
			res.Px[x][3 * inp.Height + y][1] = 0
		}
	}
	return res
}

func SliceCMYY(inp Pict) Pict {
	res := MkPict(inp.Width, inp.Height * 4)
	imp := NGPS(inp)

	for x := 0; x < res.Width; x++ {
		for y := 0; y < inp.Height; y++ {
			// TOP
			res.Px[x][y][0] = inp.Px[x][y][0]
			res.Px[x][y][1] = inp.Px[x][y][1]
			res.Px[x][y][2] = inp.Px[x][y][2]
			res.Px[x][y][3] = inp.Px[x][y][3]

			// C
			res.Px[x][inp.Height + y][0] = imp.Px[x][y][0]
			res.Px[x][inp.Height + y][1] = 0
			res.Px[x][inp.Height + y][2] = 255 - imp.Px[x][y][1]
			res.Px[x][inp.Height + y][3] = 255 - imp.Px[x][y][1]

			// M
			res.Px[x][2 * inp.Height + y][0] = inp.Px[x][y][0]
			res.Px[x][2 * inp.Height + y][1] = 255 - imp.Px[x][y][2]
			res.Px[x][2 * inp.Height + y][2] = 0
			res.Px[x][2 * inp.Height + y][3] = 255 - imp.Px[x][y][2]

			// Y
			res.Px[x][3 * inp.Height + y][0] = inp.Px[x][y][0]
			res.Px[x][3 * inp.Height + y][1] = 255 - imp.Px[x][y][3] 
			res.Px[x][3 * inp.Height + y][2] = 255 - imp.Px[x][y][3]
			res.Px[x][3 * inp.Height + y][3] = 0
		}
	}
	return res
}
