package main

import (
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)

// Estructuras para Ingresos y Gastos
type Movimiento struct {
	Monto     float64 `json:"monto"`
	Categoria string  `json:"categoria"`
}

var ingresos []Movimiento
var gastos []Movimiento
var presupuesto float64

func main() {
	// Inicia el servidor Gin
	r := gin.Default()

	// Rutas
	r.LoadHTMLGlob("templates/*")

	// Página principal
	r.GET("/", func(c *gin.Context) {
		totalIngresos := calcularTotal(ingresos)
		totalGastos := calcularTotal(gastos)
		presupuestoRestante := presupuesto - totalGastos

		c.HTML(http.StatusOK, "index.html", gin.H{
			"presupuesto":         presupuesto,
			"ingresos":            totalIngresos,
			"gastos":              totalGastos,
			"presupuestoRestante": presupuestoRestante,
		})
	})

	// Ruta para agregar ingreso
	r.POST("/agregar_ingreso", func(c *gin.Context) {
		monto, _ := strconv.ParseFloat(c.PostForm("monto"), 64)
		categoria := c.PostForm("categoria")

		ingreso := Movimiento{Monto: monto, Categoria: categoria}
		ingresos = append(ingresos, ingreso)

		c.Redirect(http.StatusFound, "/")
	})

	// Ruta para agregar gasto
	r.POST("/agregar_gasto", func(c *gin.Context) {
		monto, _ := strconv.ParseFloat(c.PostForm("monto"), 64)
		categoria := c.PostForm("categoria")

		gasto := Movimiento{Monto: monto, Categoria: categoria}
		gastos = append(gastos, gasto)

		c.Redirect(http.StatusFound, "/")
	})

	// Ruta para establecer presupuesto
	r.POST("/establecer_presupuesto", func(c *gin.Context) {
		presupuesto, _ = strconv.ParseFloat(c.PostForm("presupuesto"), 64)
		c.Redirect(http.StatusFound, "/")
	})

	// Iniciar servidor en el puerto 8080
	r.Run(":8080")
}

// Función para calcular el total de los movimientos (ingresos o gastos)
func calcularTotal(movimientos []Movimiento) float64 {
	total := 0.0
	for _, movimiento := range movimientos {
		total += movimiento.Monto
	}
	return total
}
