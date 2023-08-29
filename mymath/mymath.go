//Необходимо создать внешний пакет с именем mymath,
//который будет содержать функции для выполнения математических операций.
//
//Внешний пакет:

// Пример кода для иллюстрации задачи
import "math"

func Sqrt(x float64) float64 {
return math.Sqrt(x)
}
func Floor(x float64) float64 {
return math.Floor(x)
}
func Max(x, y float64) float64 {
return math.Max(x, y)
}
func Min(x, y float64) float64 {
return math.Min(x, y)
}
func Pow(x, y float64) float64 {
return math.Pow(x, y)
}
func Ceil(x float64) float64 {
return math.Ceil(x)
}