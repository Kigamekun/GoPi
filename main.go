package main

import "github.com/gin-gonic/gin"
import (
	// _"encoding/json"
	// "log"
	// _"net/http"
	"fmt"
	"goApi/models"
	"goApi/config"
)




func getPegawai(ctx *gin.Context) {	
	var users models.Users
	var arr_user []models.Users

	db := config.Connect()
	defer db.Close()

	result, err := db.Query("SELECT * from pegawai")
	
	if err != nil {
		panic(err.Error())
		}  
		defer result.Close()  
		for result.Next() {
			err := result.Scan(&users.Id,&users.Nip,&users.Nama,&users.Password,&users.Role)
			if err != nil {
				panic(err.Error())
			} else {
				arr_user = append(arr_user, users)
			}
	}  

	fmt.Print(arr_user)

	fmt.Println("Love You :)")

	ctx.JSON(200, gin.H{
		"statusCode": 200,
		"message": "Success get Data",
		"data": arr_user,
	})
}



func getSpp(ctx *gin.Context) {	
	var users models.Users
	var arr_user []models.Users

	db := config.Connect()
	defer db.Close()

	result, err := db.Query("SELECT * from pegawai")
	
	if err != nil {
		panic(err.Error())
		}  
		defer result.Close()  
		for result.Next() {
			err := result.Scan(&users.Id,&users.Nip,&users.Nama,&users.Password,&users.Role)
			if err != nil {
				panic(err.Error())
			} else {
				arr_user = append(arr_user, users)
			}
	}  

	ctx.JSON(200, gin.H{
		"statusCode": 200,
		"message": "Success get Data",
		"data": arr_user,
	})
}




func getSiswa(ctx *gin.Context) {	
	var users models.Users
	var arr_user []models.Users

	db := config.Connect()
	defer db.Close()

	result, err := db.Query("SELECT * from pegawai")
	
	if err != nil {
		panic(err.Error())
		}  
		defer result.Close()  
		for result.Next() {
			err := result.Scan(&users.Id,&users.Nip,&users.Nama,&users.Password,&users.Role)
			if err != nil {
				panic(err.Error())
			} else {
				arr_user = append(arr_user, users)
			}
	}  
	
	ctx.JSON(200, gin.H{
		"statusCode": 200,
		"message": "Success get Data",
		"data": arr_user,
	})
}


func getKelas(ctx *gin.Context) {	
	var users models.Users
	var arr_user []models.Users

	db := config.Connect()
	defer db.Close()

	result, err := db.Query("SELECT * from pegawai")
	
	if err != nil {
		panic(err.Error())
		}  
		defer result.Close()  
		for result.Next() {
			err := result.Scan(&users.Id,&users.Nip,&users.Nama,&users.Password,&users.Role)
			if err != nil {
				panic(err.Error())
			} else {
				arr_user = append(arr_user, users)
			}
	}  

	ctx.JSON(200, gin.H{
		"statusCode": 200,
		"message": "Success get Data",
		"data": arr_user,
	})
}


func getJurusan(ctx *gin.Context) {	
	ctx.JSON(200, gin.H{
		"message": "success",
	})
}


func main() {
	router := gin.Default()
	router.GET("/getPegawai", getPegawai)
	router.POST("/createPegawai", createPerson)
	router.PUT("/updatePegawai", createPerson)
	router.DELETE("/deletePegawai", createPerson)
	router.POST("/getSpp", createPerson)
	router.POST("/createSpp", createPerson)
	router.PUT("/updateSpp", createPerson)
	router.DELETE("/deleteSpp", createPerson)
	router.Run("127.0.0.1:8000")
}
