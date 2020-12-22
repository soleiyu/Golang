package pictFunc

import (
	"math"
	"github.com/pkg/profile"
)

func GetPx(inp Pict, xq, yq float32) []uint8 {
	xi := int(xq)
	yi := int(yq)
	xf := xq - float32(xi)
	yf := yq - float32(yi)

	res := make([]uint8, 4)
	if xi < inp.Width - 1 && yi < inp.Height - 1 {
		for i := 0; i < 4; i ++ {
			res[i] = uint8(
				float32(inp.Px[xi][yi][i]) * (float32(1) - xf) * (float32(1) - yf) +
				float32(inp.Px[xi + 1][yi][i]) * xf * (float32(1) - yf) +
				float32(inp.Px[xi][yi + 1][i]) * (float32(1) - xf) * yf +
				float32(inp.Px[xi + 1][yi + 1][i]) * xf * yf)
		}
	} else {
		res[0] = 0
		res[1] = 0
		res[2] = 0
		res[3] = 0
	}

	return res
}

func Xlens(inp Pict, ratio float32) Pict {
	res := MkPict(inp.Width, inp.Height)
	ar := float32(1) / ratio

	cx := float32(inp.Width) / float32(2)
	vm := ar * cx

	for x := 0; x < inp.Width; x++ {
		for y := 0; y < inp.Height; y++ {
			vx := float32(x) - cx
			vr :=  (vm - float32(math.Sqrt(float64(vx * vx)))) / (vm)

			res.Px[x][y] = GetPx(inp, cx + vr * vx, float32(y))
		}
	}
	return res
}

func XlensR(inp Pict, ratio float32) Pict {
	res := MkPict(inp.Width, inp.Height)
	ar := float32(1) / ratio

	cx := float32(inp.Width) / float32(2)
	vm := ar * cx

	for x := 0; x < inp.Width; x++ {
		for y := 0; y < inp.Height; y++ {
			vx := float32(x) - cx
			vr :=  (vm - float32(math.Sqrt(float64(vx * vx)))) / (vm)

			c := GetPx(inp, cx + vr * vx, float32(y))
			res.Px[x][y][0] = inp.Px[x][y][0]
			res.Px[x][y][1] = c[1]
			res.Px[x][y][2] = inp.Px[x][y][2]
			res.Px[x][y][3] = inp.Px[x][y][3]
		}
	}
	return res
}

func Clens(inp Pict, ratio float32) Pict {
	res := MkPict(inp.Width, inp.Height)
	ar := float32(1) / ratio

	cx := float32(inp.Width) / float32(2)
	cy := float32(inp.Height) / float32(2)
	vm := ar * float32(math.Sqrt(float64(cx * cx + cy * cy)))

	for x := 0; x < inp.Width; x++ {
		for y := 0; y < inp.Height; y++ {
			vx := float32(x) - cx
			vy := float32(y) - cy
			vs := float32(math.Sqrt(float64(vx * vx + vy * vy)))
			vr :=  (vm - vs) / (vm)

			res.Px[x][y] = GetPx(inp, cx + vr * vx, cy + vr * vy)
		}
	}
	return res
}

func ClensR(inp Pict, ratio float32) Pict {
	defer profile.Start(profile.ProfilePath(".")).Stop()

	res := MkPict(inp.Width, inp.Height)
	ar := float32(1) / ratio

	cx := float32(inp.Width) / float32(2)
	cy := float32(inp.Height) / float32(2)
	vm := ar * float32(math.Sqrt(float64(cx * cx + cy * cy)))

	for x := 0; x < inp.Width; x++ {
		for y := 0; y < inp.Height; y++ {
			vx := float32(x) - cx
			vy := float32(y) - cy
			vs := float32(math.Sqrt(float64(vx * vx + vy * vy)))
			vr :=  (vm - vs) / (vm)

			c := GetPx(inp, cx + vr * vx, cy + vr * vy)

			res.Px[x][y][0] = inp.Px[x][y][0]
			res.Px[x][y][1] = c[1]
			res.Px[x][y][2] = inp.Px[x][y][2]
			res.Px[x][y][3] = inp.Px[x][y][3]
		}
	}
	return res
}

func VlensR(inp Pict, ratio float32) Pict {
	res := MkPict(inp.Width, inp.Height)
	ar :=  float32(1) - ratio

	cx := float32(inp.Width) / float32(2)
	cy := float32(inp.Height) / float32(2)

	for x := 0; x < inp.Width; x++ {
		for y := 0; y < inp.Height; y++ {
			vx := float32(x) - cx
			vy := float32(y) - cy

			c := GetPx(inp, cx + ar * vx, cy + ar * vy)

			res.Px[x][y][0] = inp.Px[x][y][0]
			res.Px[x][y][1] = c[1]
			res.Px[x][y][2] = inp.Px[x][y][2]
			res.Px[x][y][3] = inp.Px[x][y][3]
		}
	}
	return res
}

