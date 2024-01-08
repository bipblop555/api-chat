# Chat Application Backend

## Overview

This repository contains the backend code for a chat application implemented in Golang using the GORM ORM. The application interacts with a MySQL database for data storage. The back of this project is conceived as a REST API. There is no html template. The back only provides JSON responses

## Setup

### Prerequisites

- Golang installed
- MySQL server installed
- Adminer for database administration
- Docker (not mandatory)
- All the lib used (check the go.mod for further information)
  
## Models
Your models will be placed here, define a type (exemple Post) struct which will defined you table and the type of data that'll go to the database
next, define the functions needed. For instance add AddPost() function that you will call to make an INSERT into the database


### Database Configuration

The application uses the following MySQL database configuration:
if you use Docker, DB_HOST will be the IP of the container

### DISCLAIMER : For now, the error while constructing the call to de DB doesn't throw any error. You might be granted by a "hello world" while launchin go run main.go. Make sure that the information to connect to the DataBase are correct.
```go
DB_HOST     = "172.31.0.2"
DB_DRIVER   = "mysql"
DB_USER     = "root"
DB_PASSWORD = "password"
DB_NAME     = "back"
DB_PORT     = "3306"
```
To start the project :

```bash
git clone https://github.com/bipblop555/api-chat.git
cd api-chat && go mod install
cd app && go run main.go
```

## API Routes
/home : to test the project is running.
The routes are basically the url you request and the function to call when the user access the url. When accessing /home you should get something like ""Welcome To This Awesome API"

### Here is an exemple of how the API works

#### Starting off with the controller base.go : 
define a struct that contains the gorm ORM
```go
type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}
```
Initiate your db connection : 
```go
func (server *Server) Initialize(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) {

	var err error

	if Dbdriver == "mysql" {
		DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DbUser, DbPassword, DbHost, DbPort, DbName)
		server.DB, err = gorm.Open(mysql.Open(DBURL), &gorm.Config{})
		if err != nil {
			fmt.Printf("Cannot connect to %s database", Dbdriver)
			log.Fatal("This is the error:", err)
		} else {
			fmt.Printf("We are connected to the %s database", Dbdriver)
		}
	}

	server.DB.Debug().AutoMigrate(&models.User{}) //database migration

	server.Router = mux.NewRouter()

	server.initializeRoutes()
}
```
Note that AutoMigrate will check the correspondiong model, here "User" to create the table according to the way the type is declared (we will come to this file in an instant)

#### The Model
```go
type User struct {
	ID        uint32    `gorm:"primary_key;auto_increment" json:"id"`
	Username  string    `gorm:"size:255;not null;unique" json:"username"`
	Password  string    `gorm:"size:100;not null;" json:"password"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
```
The functions declared here will be the function you link to your router
