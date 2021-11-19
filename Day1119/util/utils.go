package utils

func Cal(x int,y int)func(int)int{
	return func(z int)int{
		return x+y+z
	}
}