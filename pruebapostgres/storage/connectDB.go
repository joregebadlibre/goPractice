package storage 

import ( 
	"fmt" 
	"log" 
	"github.com/jinzhu/gorm"  
	//Postgres Driver imported 
	_ "github.com/lib/pq"
  )

  //ConnectDB connect to Postgres DB
func ConnectDB() *gorm.DB {
	var (
	 host = "localhost"
	 user = "postgres"
	 port = 5432
	 password = "password"
	 name = "db"
   )
   //Connect to DB
   var DB *gorm.DB  
   DB, err = gorm.Open("postgres", fmt.Sprintf("host=%s port=%d 
   user=%s password=%s dbname=%s sslmode=disable", host, portInt, 
   user, password, name))

    //Check for Errors in DB 
	 if err != nil {  
		log.Fatalf("Error in connect the DB %v", err)  
		return nil 
	   }  
	   if err := DB.DB().Ping(); err != nil { 
		 log.Fatalln("Error in make ping the DB " + err.Error())  
		 return nil 
	   }  
	   if DB.Error != nil { 
		log.Fatalln("Any Error in connect the DB " + err.Error()) 
		return nil 
	   }  
	   log.Println("DB connected")  
	   return DB

  }