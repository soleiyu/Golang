package main                                                                                                          

import ( 
  "fmt" 
  "./mandellib"
  "../pictPr/pictFunc"
  "strconv"
  "math"
)
 
func main() {
//  mkMandelMap(1280, 960, 1.6, -0, 1.2, 0, 0)
  mkJuliaMap(1280, 960, 1.6, 0, 1.2, -0.52, -0.54)
//  mkMandelBurst2(1280, 720, 200, 1.5, -0.7, 0.9)
//    mkJuliaBurst(1280, 960, 200, 1.6, -0.7, 1.2)
}
 
func mkJuliaMap(w, h int, rn, cpr, in, zx, zy float64) {
  p := pictFunc.MkPict(w, h)
  hw := w / 2
  hh := h / 2

  for x := 0; x < w; x++ {
    for y := 0; y < h; y++ {
      pxv := mandellib.IsMandel(zx, zy,
        rn * (float64)(x - hw) / (float64)(hw) +cpr,
        in * (float64)(y - hh) / (float64)(hh), 
        255)

      p.Px[x][y][0] = 255
      p.Px[x][y][1] = (uint8)(254 - pxv)
      p.Px[x][y][2] = (uint8)(254 - pxv)
      p.Px[x][y][3] = (uint8)(254 - pxv)
    }
  }

  p.Save("juliaplot.png")
}
