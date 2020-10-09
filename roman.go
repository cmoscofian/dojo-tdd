package main

type converter struct {
	Remaining int
	Roman     string
}

func ConvertToRoman(num int) string {
	c := &converter{
		Remaining: num,
		Roman:     "",
	}

	buildPattern(1000, "M", c)
	buildPattern(900, "CM", c)
	buildPattern(500, "D", c)
	buildPattern(400, "CD", c)
	buildPattern(100, "C", c)
	buildPattern(90, "XC", c)
	buildPattern(50, "L", c)
	buildPattern(40, "XL", c)
	buildPattern(10, "X", c)
	buildPattern(9, "IX", c)
	buildPattern(5, "V", c)
	buildPattern(4, "IV", c)
	buildPattern(1, "I", c)

	return c.Roman
}

func buildPattern(num int, pattern string, c *converter) {
	for c.Remaining >= num {
		c.Roman += pattern
		c.Remaining -= num
	}
}
