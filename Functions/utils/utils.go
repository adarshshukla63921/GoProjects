package utils


// func functionName() returnTye {}
func Increment(x *int) int {
	*x = *x +1;
	return *x;
}
func ExportedFunction() string {
	return "This is an exported function"
}

 var ExportedVariable = "This is an exported variable"

