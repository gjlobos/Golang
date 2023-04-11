package main

import "github.com/gin-gonic/gin"

func main() {
	// Setea el router como el default que viene con Gin
	r := gin.Default()

	// Vamos a crear una aplicación web con el framework Gin
	// que tenga un endpoint /hola-mundo que responda con un mensaje.
	// Tener en cuenta que:
	//	- El endpoint deberá ser de método GET.
	//	- La respuesta deberá ser recibida en formato JSON.
	// Ejemplo:
	// {
	// 	"mensaje": "¡Hola Mundo!"
	// }

	r.GET("/hola-mundo", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "¡Hola Mundo!",
		})
	})

	r.Run() // Escucha y expone el códiog en el puerto x default 8080
}
