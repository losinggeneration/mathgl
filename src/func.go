package mathgl

const (
	PI         float32 = 3.141592
	PIover180  float32 = 0.017453
	PIunder180 float32 = 57.295779
	epsilon    float32 = 1.0 / 64.0
)

func Fsqr32(s float32) float32 {
	return s * s
}
func Fdeg2rad32(degrees float32) float32 {
	return degrees * PIover180
}

func Frad2deg32(radians float32) float32 {
	return radians * PIunder180
}

func Fmin32(lhs float32, rhs float32) float32 {
	if lhs < rhs {
		return lhs
	}
	return rhs
}

func Fmax32(lhs float32, rhs float32) float32 {
	if lhs > rhs {
		return lhs
	}
	return rhs
}

func FalmostEqual32(lhs float32, rhs float32) bool {
	return (lhs+epsilon > rhs && lhs-epsilon < rhs)
}

