package calculator

// Add adalah fungsi publik yang dapat diakses dari luar package
func Add(a, b int) int {
    return a + b
}

// multiply adalah fungsi privat yang hanya dapat diakses di dalam package
func multiply(a, b int) int {
    return a * b
}
