package funcs

//GetSqrtfunc returns func val/pointer to Sqrtfunc / Function closures
func GetSqrtfunc() func(x, y float64) string {
	return func(x, y float64) string {
		return Sqrtfunc(x, y)
	}
}

//SqrtfuncInvoker is a invoker for Sqrtfunc
func SqrtfuncInvoker(a, b float64, fn func(x, y float64) string) string {
	return fn(a, b)
}
