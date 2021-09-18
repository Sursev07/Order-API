package main


import(
	"Order-API/database"
	"Order-API/router"
)

var PORT = ":3000"

func main()  {
	database.InitDB()
	r := router.StartApp()
	r.Run(PORT)
}
