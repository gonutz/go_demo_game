package ease

import "math"

func Linear(x float64) float64 {
	return x
}

func InSine(x float64) float64 {
	return 1 - math.Cos((x*math.Pi)/2)
}

func OutSine(x float64) float64 {
	return math.Sin((x * math.Pi) / 2)
}

func InOutSine(x float64) float64 {
	return -(math.Cos(math.Pi*x) - 1) / 2
}

func InQuad(x float64) float64 {
	return x * x
}

func OutQuad(x float64) float64 {
	return 1 - (1-x)*(1-x)
}

func InOutQuad(x float64) float64 {
	if x < 0.5 {
		return 2 * x * x
	}
	return 1 - math.Pow(-2*x+2, 2)/2
}

func InCubic(x float64) float64 {
	return x * x * x
}

func OutCubic(x float64) float64 {
	return 1 - math.Pow(1-x, 3)
}

func InOutCubic(x float64) float64 {
	if x < 0.5 {
		return 4 * x * x * x
	}
	return 1 - math.Pow(-2*x+2, 3)/2
}

func InQuart(x float64) float64 {
	return x * x * x * x
}

func OutQuart(x float64) float64 {
	return 1 - math.Pow(1-x, 4)
}

func InOutQuart(x float64) float64 {
	if x < 0.5 {
		return 8 * x * x * x * x
	}
	return 1 - math.Pow(-2*x+2, 4)/2
}

func InQuint(x float64) float64 {
	return x * x * x * x * x
}

func OutQuint(x float64) float64 {
	return 1 - math.Pow(1-x, 5)
}

func InOutQuint(x float64) float64 {
	if x < 0.5 {
		return 16 * x * x * x * x * x
	}
	return 1 - math.Pow(-2*x+2, 5)/2
}

func InExpo(x float64) float64 {
	if x == 0 {
		return 0
	}
	return math.Pow(2, 10*x-10)
}

func OutExpo(x float64) float64 {
	if x == 1 {
		return 1
	}
	return 1 - math.Pow(2, -10*x)
}

func InOutExpo(x float64) float64 {
	if x == 0 {
		return 0
	}
	if x == 1 {
		return 1
	}
	if x < 0.5 {
		return math.Pow(2, 20*x-10) / 2
	}
	return (2 - math.Pow(2, -20*x+10)) / 2
}

func InCirc(x float64) float64 {
	return 1 - math.Sqrt(1-math.Pow(x, 2))
}

func OutCirc(x float64) float64 {
	return math.Sqrt(1 - math.Pow(x-1, 2))
}

func InOutCirc(x float64) float64 {
	if x < 0.5 {
		return (1 - math.Sqrt(1-math.Pow(2*x, 2))) / 2
	}
	return (math.Sqrt(1-math.Pow(-2*x+2, 2)) + 1) / 2
}

func InBack(x float64) float64 {
	const c1 = 1.70158
	const c3 = c1 + 1

	return c3*x*x*x - c1*x*x
}

func OutBack(x float64) float64 {
	const c1 = 1.70158
	const c3 = c1 + 1

	return 1 + c3*math.Pow(x-1, 3) + c1*math.Pow(x-1, 2)
}

func InOutBack(x float64) float64 {
	const c1 = 1.70158
	const c2 = c1 * 1.525

	if x < 0.5 {
		return (math.Pow(2*x, 2) * ((c2+1)*2*x - c2)) / 2
	}
	return (math.Pow(2*x-2, 2)*((c2+1)*(x*2-2)+c2) + 2) / 2
}

func InElastic(x float64) float64 {
	const c4 = (2 * math.Pi) / 3

	if x == 0 {
		return 0
	}
	if x == 1 {
		return 1
	}
	return -math.Pow(2, 10*x-10) * math.Sin((x*10-10.75)*c4)
}

func OutElastic(x float64) float64 {
	const c4 = (2 * math.Pi) / 3

	if x == 0 {
		return 0
	}
	if x == 1 {
		return 1
	}
	return math.Pow(2, -10*x)*math.Sin((x*10-0.75)*c4) + 1
}

func InOutElastic(x float64) float64 {
	const c5 = (2 * math.Pi) / 4.5

	if x == 0 {
		return 0
	}
	if x == 1 {
		return 1
	}
	if x < 0.5 {
		return -(math.Pow(2, 20*x-10) * math.Sin((20*x-11.125)*c5)) / 2
	}
	return (math.Pow(2, -20*x+10)*math.Sin((20*x-11.125)*c5))/2 + 1
}

func InBounce(x float64) float64 {
	return 1 - OutBounce(1-x)
}

func OutBounce(x float64) float64 {
	const n1 = 7.5625
	const d1 = 2.75

	if x < 1/d1 {
		return n1 * x * x
	} else if x < 2/d1 {
		x -= 1.5 / d1
		return n1*x*x + 0.75
	} else if x < 2.5/d1 {
		x -= 2.25 / d1
		return n1*x*x + 0.9375
	}
	x -= 2.625 / d1
	return n1*x*x + 0.984375
}

func InOutBounce(x float64) float64 {
	if x < 0.5 {
		return (1 - OutBounce(1-2*x)) / 2
	}
	return (1 + OutBounce(2*x-1)) / 2
}
