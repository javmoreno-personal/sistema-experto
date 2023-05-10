package web

import (
	"log"
	"net/http"
	
	"github.com/gin-gonic/gin"
	ctrl "E:/Escritorio/Sistemas-experto/internal/adapters/controllers"

)

const port = ":9001"

// Iniciar el servicio HTTP

func NewHttpServer(sistemaExpertoCtrl ctrl.controller) error {
	router := gin.Default()
	basePath := "/api/v1"
	publicRouter := router.Group(basePath)

	//routes

	//creemos una ruta que nos permita enviar al backend un JSON con los datos que estan en el readme.md
	publicRouter.POST("/sistema-experto", sistemaExpertoCtrl.PostSistemaExperto)
	publicRouter.GET("/sistema-experto", sistemaExpertoCtrl.GetSistemaExperto)


	log.Println("Iniciando el servidor HTTP en el puerto" + port)
	return http.ListenAndServe(port, router)
}
