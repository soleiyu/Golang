package mandellib                                                                                                     
 
type Complex struct {
  Rn float64
  In float64
}
  
func MkComp(rn , in float64) Complex {
  var c Complex
  c.Rn = rn
  c.In = in 
  return c
}
  
func MandelFunc (zn, c Complex) Complex {
  return MkComp(
    zn.Rn * zn.Rn - zn.In * zn.In + c.Rn, 
    2.0 * zn.In * zn.Rn + c.In)
}
 
func (zn Complex) DistSq () float64 {
  return zn.Rn * zn.Rn + zn.In * zn.In
}

func IsMandel(x, y ,zx, zy float64, lim int) int {
  z := MkComp(zx, zy)
  c := MkComp(x, y)   

  for i := 0; i < lim; i++ {
    z = MandelFunc(z, c)
    if 4 < z.DistSq() {
      return i 
    }
  }

  return lim
}
