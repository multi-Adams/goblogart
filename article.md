APIs, or Application Programming Interfaces, are a way for different software applications to communicate with each other. They are used to share data and functionality between applications and are a vital part of the modern web. Creating production-ready APIs is crucial for developers looking to build robust and scalable applications. In the world of backend development, Go has gained significant popularity for its simplicity, performance, and strong community support. With frameworks like Gin and Gorm, Go becomes a powerful tool for developing RESTful APIs.

In this step-by-step guide, you will learn how to create production-ready blog API using Go, along with the Gin web framework for routing and middleware, and Gorm as an ORM (Object-Relational Mapping) tool for working with databases.

Let's take a look at the prerequisites in the next section.

## Prerequisites

To follow along with this tutorial, you'll need the folowing:

- Working Go installation
- Basic knowledge of Go. You can check out my Go Beginners Series to get started.
- Postman

## Introduction to Go and the Libraries

In this section, I'll go over a brief introduction of Go, Gorm, and the other libraries that will be used to build the blog API that we will be creating.

### Go

Golang, popularly known as Go, is an open-source, compiled, high-level, and statically typed language created in 2007 by Google developers; Robert Griesemer, Rob Pike, and Ken Thompson. It was open-sourced in 2009 and has since become popular as big companies have started using it to build significant applications.

Go focuses on improving developers' productivity by providing run-time efficiency, simple syntax, static typing, and other very useful features. A lot of companies have testified that Go has saved them a lot of development time with its incredible features. You can get started with Go using my Go Beginners Series, a series of hands-on, short, step-by-step articles explaining the basics of Go.

### Gin Web Framework

Gin is the most popular web framework written in Go. It has over 68 thousand stars on Github and so many Go developers (Gophers) use and recommend it for fast development and performance.

It boasts of so many incredible features, such as blazing fast performance, JSON validation, routes grouping, error management, and so much more.

Let's look at some of the pros and cons for Gin in the next sections.

#### Pros of Gin Framework

#### Simplicity and Lightweight Nature

Gin is widely appreciated for its simplicity and minimalistic design. It offers a straightforward API and follows the principles of the Go language, making it easy to learn and work with. The lightweight framework results in faster startup times and reduced memory footprint, which can be advantageous in high-performance applications.

#### Fast and High Performance

One of the key advantages of using Gin is its impressive speed and high-performance characteristics. The framework is built to prioritize speed, leveraging the inherent efficiency of Go. Gin achieves this by utilizing fast HTTP router and middleware components, allowing for rapid development and efficient request handling.

#### Robust Middleware Support

Gin provides robust support for middleware, enabling developers to add additional functionality to their applications. Middleware functions can be used for tasks like logging, authentication, error handling, rate limiting, etc. The middleware ecosystem in Gin is extensive, offering a wide range of pre-built middleware packages that can be easily integrated into your project.

#### Scalability and Concurrency

Gin's design is well-suited for building scalable and concurrent applications. Go's native support for concurrency, combined with Gin's lightweight nature, makes it ideal for handling a large number of concurrent requests without sacrificing performance. This scalability is particularly valuable for applications that experience high traffic or require real-time updates.

#### Cons of Gin Framework

#### Lack of Built-in Features

While Gin's simplicity is a strength, it can also be considered a limitation. The framework does not provide an extensive set of built-in features compared to some other web frameworks. Developers may need to rely on third-party libraries or build custom solutions for certain functionalities, which can result in additional configuration and maintenance efforts.

#### Steeper Learning Curve for Beginners

Although Gin's simplicity is advantageous for experienced developers, it may pose a challenge for beginners or those new to Go programming. Understanding the underlying concepts of Go and the idiomatic ways of using Gin may require a certain level of familiarity with the language. However, newcomers can still become proficient in using Gin with dedication and resources like documentation and community support.

#### Limited Flexibility

Gin follows a specific design philosophy and structure, which may limit the flexibility of developers who prefer more freedom in their project architecture. While this structure can benefit small to medium-sized applications, complex or highly customized projects may require more flexibility that may not align with Gin's opinionated approach.

Ultimately, the decision to use Gin should be based on your project requirements, the level of flexibility needed, and your familiarity with the Go language. However, if simplicity, performance, and middleware support align with your application goals, Gin can be an excellent choice.

### Gorm

Gorm is a popular and fully featured object-relational mapper (ORM) written in Go that has over 32 thousand stars on GitHub. It provides convenient abstractions that simplify database interactions and enhance productivity.

It boasts of a lot of fantastic features, including associations, hooks, transactions, SQL builder, and so much more.

#### Pros of Gorm

#### Simplified Database Interactions

Gorm simplifies working with databases by abstracting away much of the underlying SQL complexity. It provides a high-level API that allows you to interact with databases using simple and intuitive Go code. Gorm handles tasks such as table creation, querying, data manipulation, and relationships, reducing boilerplate code, and accelerating development.

#### Database Agnostic

Gorm is designed to be database agnostic, meaning it supports multiple database systems such as PostgreSQL, MySQL, SQLite, and more. This flexibility allows you to switch between databases without significant code changes. Gorm's database agnosticism is beneficial for projects that require database portability or may need to support multiple database systems in the future.

#### Automatic CRUD Operations

Gorm automates typical Create, Read, Update, and Delete (CRUD) operations, reducing the need for manual SQL queries. It provides convenient methods like Create, Find, Update, and Delete that handle the underlying SQL generation, making database interactions more efficient and less error-prone. This abstraction simplifies data management and improves productivity.

#### Association and Relationship Management

Gorm excels at handling database relationships, such as one-to-one, one-to-many, and many-to-many associations. It provides intuitive methods for defining and managing relationships between entities, allowing you to work with related data effortlessly. Gorm's association management capabilities streamline complex data modeling and retrieval, eliminating the need for manual JOIN statements.

#### Middleware Support

Gorm offers middleware support, enabling you to add additional functionality and behaviors during database operations. Middleware can be used for tasks such as data validation, timestamp management, logging, and more. This extensibility allows you to customize Gorm's behavior to suit their specific project requirements and implement cross-cutting concerns effectively.

#### Cons of Gorm

#### Learning Curve

Gorm's feature-rich nature and abstraction come with a learning curve. While Gorm simplifies database operations, understanding its concepts, conventions, and API may require initial investment and familiarization. In general, developers new to Gorm or ORM concepts may face a learning curve to leverage their capabilities thoroughly. However, the Gorm documentation and community resources can help mitigate this challenge.

#### Performance Overhead

Like any ORM, Gorm introduces an additional abstraction layer between the application and the database. This abstraction can result in some performance overhead compared to writing raw SQL queries. While Gorm's performance is generally acceptable for most applications, projects with highly performance-sensitive requirements or complex questions may benefit from using basic SQL queries instead.

#### Limited Customization

Gorm's primary goal is to provide a simplified API for everyday database operations. While it offers various customization options, there may be cases where more intricate and specialized database interactions are required. In such scenarios, Gorm's abstraction may limit the developer's ability to optimize or fine-tune queries, potentially requiring workarounds or direct SQL execution.

#### Database-specific Features

While Gorm supports multiple databases, it may only expose some of the advanced features specific to each database system. Database-specific functionalities, such as advanced indexing options or specialized query features, may not be directly accessible through Gorm's API. In such cases, developers may need to resort to raw SQL or database-specific extensions to leverage those features.

#### Lack of Maturity

Compared to other well-established ORM libraries, Gorm is relatively newer and may be considered less mature. This means it may have occasional bugs or limitations yet to be addressed. However, the Gorm community actively maintains and improves the library, releasing updates and addressing issues regularly, which helps mitigate this concern.

Gorm offers significant advantages in simplifying database interactions; however, before adopting Gorm, carefully evaluate your project's requirements, consider the trade-offs, and assess whether Gorm's strengths align with your needs.

Now that you know what Go and Gorm are, their pros and cons, and how to decide whether to use them or not, let's set up the development environment in the next section.

## Environment and Folders Setup

In this section, we will set up the development environment for the blog API project.

### Initialize Project

The first step is to initialize a new Go project. Open a new terminal and run the following command to create a new folder; You can name it whatever you like, I'll call mine `goblogart`:

```sh
mkdir goblogart
```

Next, move into that folder with the following command:

```sh
cd goblogart
```

Then initialize the new Go project with the with the `go mod` command:

```sh
go mod init goblogart
```

### Create Folders and Main file

Run the following command to create all the folders that will hold the project files:

```sh
mkdir inits controllers middlewares migrations models
```

The command above will create five folders:

1. inits - will contain all the initializer files.
2. controllers - will contain controller files that will handle route functions.
3. middlewares - will contain middleware files
4. migrations - will contain migration files
5. models - will contain model files

### Install Libraries

In the section, you will install all the necessary frameworks and libraries for your project.

Run the following command to install the CompileDaemon package for automatic builds:

```sh
go get github.com/githubnemo/CompileDaemon
```

Then you need to install it using the `go install` command like so:

```sh
go install github.com/githubnemo/CompileDaemon
```

Once that is done, run the following command to install the `godotenv` package for securing your application's secrets:

```sh
go get github.com/joho/godotenv
```

Next, install the Gin framemork with the following command:

```sh
go get -u github.com/gin-gonic/gin
```

Once Gin is installed, run the following command to install Gorm:

```sh
go get -u gorm.io/gorm
```

You'll need a database driver to work with Gorm. Run the following command to install a database driver which will be used to connect with your database:

```sh
go get -u gorm.io/driver/mysql
```

The command above will install a database driver for `mysql` which is the database system that XAMPP provides me. Visit the [XAMPP Downloads Page](https://www.apachefriends.org/) to install XAMPP for your machine.

> Note: You can use and install drivers for other database systems like PostgreSQL, MySQL, and more.

Once Gorm and the database driver is installed, run the following command to install the `bcrypt` package which will be used to hash passwords:

```sh
go get -u golang.org/x/crypto/bcrypt
```

Next, run the following command to install the `jwt-go` package which will be used to authenticate users:

```sh
go get -u github.com/golang-jwt/jwt/v5
```

Once that's done, create a `main.go` file in the root of your project and add the following code inside it:

```sh
package main

import "fmt"

func main() {
 fmt.Println("Hello, World!")
}
```

Next, you need to run the ComplileDaemon command so the project will build automatically every time you save a file:

```sh
CompileDaemon -command="./goblogart"
```

Once you run that command, you should see a message that ends with a blinking cursor in the command line:

![Image description](https://dev-to-uploads.s3.amazonaws.com/uploads/articles/25557ei67hawhe9pxshz.png)

And that's it! You have installed the required libraries and frameworks that's needed for your project. Let's set up the database in the next section.

> Note: After installing all the frameworks, libraries, and building your application, two files named `go.sum` and `gobogart.exe` will be automatically created in the root of your application. Do **not** delete these files because they are crucial for the application to work.

## Database Setup

In this section, you will learn how to connect your application to a database. First, you need to store your application's credentials and secrets in a safe place.

### Set up Environment Variables

Create two new files, `.env` file in the root folder of your application and `envLoader.go` inside the `inits` folder.

You need to set a `PORT` variable, which is a port where you can access your application. Add the following code inside the `.env` file:

```sh
PORT=3000
```

The `PORT` can be any port number that is available on your machine. However, it must be defined in uppercase. Next, you need to load the `.env` file when the project is being initialized. Add the following code inside the `envLoader.go` file:

```go
package inits

import (
 "log"

 "github.com/joho/godotenv"
)

func LoadEnv() {
 err := godotenv.Load()
 if err != nil {
  log.Fatal("Error loading .env file")
 }
}
```

Next, you need to call the `LoadEnv` function and create a `/` route inside the `main.go` file. Replace the code inside `main.go` file with the following:

```sh
package main

import (
 "goblogart/inits"

 "github.com/gin-gonic/gin"
)

func init() {

 inits.LoadEnv()
}

func main() {
 r := gin.Default()

 r.GET("/", func(c *gin.Context) {
  c.JSON(200, gin.H{
   "message": "Hello World!",
  })
 })

 r.Run()
}
```

Now, go to `localhost:3000` in your browser and you will see the following in the browser:

```sh
{
"message": "Hello World!"
}
```

![Image description](https://dev-to-uploads.s3.amazonaws.com/uploads/articles/ofsyi9fivvirggwfajpz.png)

> Note: The `init` function will be called automatically and cannot have any parameters.

### Create and Connect Database

Start the Apache and MySQL services on XAMPP, go to `localhost/phpmyadmin` and click New on the sidebar. You should see a form like the one below, enter your preferred database name and click the Create button.

![Image description](https://dev-to-uploads.s3.amazonaws.com/uploads/articles/tu1r42w1a67z8ctefsf2.png)

![Image description](https://dev-to-uploads.s3.amazonaws.com/uploads/articles/cfttih8b75py21b8iwzh.png)

Next, you need to add a database connection string `DB_URL` variable inside the `.env` file like so:

```sh
DB_URL="username:password@tcp(host)/databasename?charset=utf8mb4&parseTime=True&loc=Local"
```

> Note: You need to replace the `username`, `password`, `host`, and `databasename` with the corresponding values in your database credentials.

Next, you need to create a function that will connect to the database using the `DB_URL` and return a exported `DB` variable that can be used in other parts of the application. Create a new file `db.go` inside the `inits` folder and add the following code inside it:

```go
package inits

import (
 "os"

 "gorm.io/driver/mysql"

 "gorm.io/gorm"
)

var DB *gorm.DB

func DBInit() {
 dsn := os.Getenv("DB_URL")
 db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
 if err != nil {
  panic("failed to connect database")
 }

 DB = db
}
```

> Note: This connection function will not work for other database systems like Postgres and others. Visit the [Connecting to a Database](https://gorm.io/docs/connecting_to_the_database.html) page on the Gorm documentation.

Next, add the following code inside the `init` function inside the `main.go` file:

```go
func init() {
// same as before...
 inits.DBInit()
}
```

Save all your files and make sure there's no error in the terminal before continuing to the next step.

Finally, open Postman and send a `GET` request to `http://localhost:3000`, and you should get the same response like this:

```sh
{
"message": "Hello World!"
}
```

That's it for the database connection. Let's work on the Create, Read, Update, and Delete (CRUD) operations in the next section.

## Models and Migrations

Now that the environment and database connection are all set, let's create the models and migrations. Models and migrations are used to interact with the database from your code, you can create, delete, and update tables with them.

Create a `postModel.go` file inside your `models` folder and paste the following code inside it:

```go
package models

import "gorm.io/gorm"

type Post struct {
 gorm.Model
 Title  string
 Body   string
 Likes  int
 Draft  bool
 Author string
}
```

The code above declared the file as a `models` package, imports the `gorm` package, and defines a `Post` struct that has the fields, `Title`, `Body`, `Likes`, `Draft`, and `Author`. This struct is recognized as a model with the `gorm.Models` line that precedes the fields.

Next, create a `migrations.go` file inside the `migrations` folder and add the following code inside it:

```go
package main

import (
 "goblogart/inits"
 "goblogart/models"
)

func init() {
 inits.LoadEnv()
 inits.DBInit()
}

func main() {
 inits.DB.AutoMigrate(&models.Post{})
}
```

The code above declared the file as a `main` package, imports the `inits` and `models` packages, and uses the `init` function to run the `LoadEnv` and `DBInit` functions. Finally, it defines a `main` function that calls the `AutoMigrate` method on the `DB` variable with the `Post` model as an argument.

Finally, you need to run the migrations with the terminal. Run the following command to do so:

```sh
go run migrations/migrations.go
```

You should now have a `posts` table inside your database like so:

![posts table](https://dev-to-uploads.s3.amazonaws.com/uploads/articles/4soxc17j8ac29q9ffnjt.png)

The Gorm library will create the `id`, `created_at`, `updated_at`, and `deleted_at` columns for you automatically.

## The CRUD Operations

Now that the models, migration, and database table are set, let's create the controller functions to create, read, update, and delete posts.

Create a `postController` function inside the `controllers` folder and add the following code inside it:

```go
package controllers

import (
 "goblogart/inits"
 "goblogart/models"
 "fmt"

 "github.com/gin-gonic/gin"
)
```

The code above declares the file as part of the `controllers` package and imports the required packages, `inits`, `models`, and `gin`.

### The `CreatePost` Function

The `CreatePost` function will be responsible for adding a post to the database. Add the following code to the `postController` file:

```go
func CreatePost(ctx *gin.Context) {

 var body struct {
  Title  string
  Body   string
  Likes  int
  Draft  bool
  Author string
 }

 ctx.BindJSON(&body)

 post := models.Post{Title: body.Title, Body: body.Body, Likes: body.Likes, Draft: body.Draft, Author: body.Author}

 fmt.Println(post)
 result := inits.DB.Create(&post)
 if result.Error != nil {
  ctx.JSON(500, gin.H{"error": result.Error})
  return
 }

 ctx.JSON(200, gin.H{"data": post})

}
```

The code above defines a `body` struct that represents the request data, binds it to JSON, defines a post variable with type `models.Post` that binds `body` to the `Post` model, and defines a `result` variable that will contain the result of the `Create` method on `DB`. It then checks if there's any error, returns the error if there's one, if not, it returns a 200 response with the new post.

### The `GetPosts` Function

The `GetPosts` function will be responsible for getting all posts from the database. Add the following code to the `postController` file:

```go
func GetPosts(ctx *gin.Context) {

 var posts []models.Post

 result := inits.DB.Find(&posts)
 if result.Error != nil {
  ctx.JSON(500, gin.H{"error": result.Error})
  return
 }

 ctx.JSON(200, gin.H{"data": posts})

}
```

The code above defines a `posts` slice with the type of `model.Post`, creates a `result` variable that uses the `Find` method on `DB` and stores the posts inside the `posts` slice, checks for errors, and returns a 200 response with the returned posts.

### The `GetPost` Function

The `GetPost` function will be responsible for getting a single post with the `id` from the database. Add the following code to the `postController` file:

```go
func GetPost(ctx *gin.Context) {

 var post models.Post

 result := inits.DB.First(&post, ctx.Param("id"))
 if result.Error != nil {
  ctx.JSON(500, gin.H{"error": result.Error})
  return
 }

 ctx.JSON(200, gin.H{"data": post})

}
```

The code above defines a `post` variable with the type `models.Post`, creates a `result` variable that uses the `First` method on `DB` with the given `id` and stores the post inside the `post` slice, checks for errors, and then returns a 200 response with the returned post.

### The `UpdatePost` Function

The `UpdatePost` function will be responsible for updating a post in the database. Add the following code to the `postController` file:

```go
func UpdatePost(ctx *gin.Context) {

 var body struct {
  Title  string
  Body   string
  Likes  int
  Draft  bool
  Author string
 }

 ctx.BindJSON(&body)

 var post models.Post

 result := inits.DB.First(&post, ctx.Param("id"))
 if result.Error != nil {
  ctx.JSON(500, gin.H{"error": result.Error})
  return
 }

 inits.DB.Model(&post).Updates(models.Post{Title: body.Title, Body: body.Body, Likes: body.Likes, Draft: body.Draft, Author: body.Author})

 ctx.JSON(200, gin.H{"data": post})

}
```

The code above will defines a `body` struct that represents the request data, binds it to JSON, defines a post variable with type `models.Post` that binds `body` to the `Post` model, and defines a `result` function that contains the result of the `First` method with the given `id` on the DB variable. It then checks if there's any error, returns the error if there's one, if not, it returns a 200 response with the updated post.

### The `DeletePost` Function

The `DeletePost` function will be responsible for deleting a post from the database. Add the following code to the `postController` file:

```go
func DeletePost(ctx *gin.Context) {

 id := ctx.Param("id")

 inits.DB.Delete(&models.Post{}, id)

 ctx.JSON(200, gin.H{"data": "post has been deleted successfully"})

}
```

The code above defines a variable, `id` that extracts the id of the post from the URL, uses the `Delete` method with the `models.Post` type and the `id` on the DB variable, then returns a 200 response with a success message.

## Routes Setup

Now that the create, read, update, and delete functions for posts are in place, you need to set up each function with its corresponding route in the `main.go` file.

### The Create Route

To connect the `CreatePost` to a route, replace the code inside the `main` function of the `main.go` file with the following code:

```go
 r := gin.Default()

 r.POST("/", controllers.CreatePost)

 r.Run()
```

The code above defines a variable, `r` which is an engine instance of Gin, defines a route for the `CreatePost` function in the `postController` file, and uses the `Run` method to start the server. You also need to import the `controllers` package at the top of the `main.go` file like so:

```go
import (
        // same as before...
 "goblogart/controllers"
)
```

And that's it! You can now test the route in Postman. Send a `POST` request to the `localhost:3000/` with the following JSON data:

```json
{
    "title":"How to use Type Hinting in PHP",
    "body":"Errors, also known as bugs, are developers' biggest nightmare. Bugs are so ubiquitous that everything you do after building a website or app involves finding and fixing them. What if there was a way to reduce the number of errors you must fix after pushing a project to production? In this article, I will explain everything you need to know about type-hinting in PHP, which will help you achieve this goal. We’ll begin by exploring what type-hinting in PHP is, and then I’ll show you how to start using it in your applications. However, before we begin, let's examine what you need to know to get the most out of this article.",
    "likes":0,
    "draft": true,
    "author": "Adebayo Adams"
}
```

The request should return a new post like so:

![create posts result](https://dev-to-uploads.s3.amazonaws.com/uploads/articles/54f39689icbkbrg1mwn9.png)
The purple box - type of request.
The yellow box - URL of the route.
The red box - JSON data sent to the server.
The green box - result from the server.

### The `GetPosts` Route

To connect the `GetPosts` to a route, add the following code below the previous route inside the `main` function of the `main.go`:

```go
r.GET("/", controllers.GetPosts)
```

This request doesn't send any data to the server. Send a `GET` request to the `localhost:3000/` and you should get a response like this:

![get posts result](https://dev-to-uploads.s3.amazonaws.com/uploads/articles/160tp8brscvn5c2b6eij.png)

The purple box - type of request.
The yellow box - URL of the route.
The green box - result from the server.

### The GetPost Route

To connect the `GetPost` to a route, add the following code below the previous route inside the `main` function of the `main.go`:

```go
r.GET("/:id", controllers.GetPost)
```

This request needs a post id at the end of the URL to work correctly. Send a `GET` request to the `localhost:3000/2` and you should get a response like this:

![get single post result](https://dev-to-uploads.s3.amazonaws.com/uploads/articles/q4x9lhusk07b50pub862.png)
The purple box - type of request.
The yellow box - URL of the route.
The green box - result from the server.

### The UpdatePost Route

To connect the `UpdatePost` to a route, add the following code below the previous route inside the `main` function of the `main.go`:

```go
r.PUT("/:id", controllers.UpdatePost)
```

This request needs both post id at the end of the URL and the data that needs to be updated to work correctly. Send a `PUT` request with the field you want to update as JSON to the `localhost:3000/2` and you should get a response like this:

![update post result](https://dev-to-uploads.s3.amazonaws.com/uploads/articles/qgz0idtnb52iaaz9g7wa.png)

The purple box - type of request.
The yellow box - URL of the route.
The red box - JSON data sent to the server.
The green box - result from the server.

### The DeletePost Route

To connect the `DeletePost` to a route, add the following code below the previous route inside the `main` function of the `main.go`:

```go
r.DELETE("/:id", controllers.DeletePost)
```

This request needs a post id at the end of the URL to work correctly. Send a `DELETE` request to the `localhost:3000/2` and you should get a response like this:

![delete post request](https://dev-to-uploads.s3.amazonaws.com/uploads/articles/qrp46nl7hrdwwda2l2o4.png)

The purple box - type of request.
The yellow box - URL of the route.
The green box - result from the server.

Now that the posts controller functions and routes are done and working correctly, let's set up the user controller and authentication in the next section.

## The User Model, Controller, and Authentication

In this section we will explore table relationships, user authentication with JWT, and middlewares.

### User Table and Has Many Relationship

Create a `userModel.go` file inside the `models` folder, and paste the following code inside it:

```go
package models

import "gorm.io/gorm"

type User struct {
 gorm.Model
 Name     string
 Email    string `gorm:"unique"`
 Password string
 Posts    []Post
}
```

The code above declared the file as a `models` package, imports the `gorm` package, and defines a `User` struct that has the fields, `Name`, `Email`, `Password`, and `Posts`. This struct is recognized as a model with the `gorm.Models` line that precedes the fields.

The `Posts` field in the above code is the array that will hold the posts of each user. To connect this to the user, you need to add a `UserID` field to the `Post` model like so:

```go
// same as before...
type Post struct {
 gorm.Model
// same as before...
 UserID uint `gorm:"user_id"`
}
```

This new line of code will add a new `user_id` column to the `posts` table in the database. Next, add the following code below the `ctx.BindJSON(&body)` line inside the `CreatePost` function:

```go
 user, exists := ctx.Get("user")

 if !exists {
  ctx.JSON(500, gin.H{"error": "user not found"})
  return
 }

 body.UserID = user.(models.User).ID
```

Next, update the `post variable inside the function to look like this:

```go
post := models.Post{Title: body.Title, Body: body.Body, Likes: body.Likes, Draft: body.Draft, Author: body.Author, UserID: body.UserID}
```

These additions will enable the function to get the logged in user ID and attached it to the new post.

Next, add the new `User` model to the migrations file like so:

```go
inits.DB.AutoMigrate(&models.User{})
```

Next, run the migrations to update your database with the following command:

```sh
go run migrations/migrations.go
```

A new `users` table should be in your database and the `posts` table should now have a `user_id` column. The `Post` model and the `User` model are now connected with the `has many` relationship, which means a user can have as many posts as possible.

> Note: The posts that have been in the database before the relatioship will have the value of `NULL` in their `user_id` columns.

### User Signup

Create a `userController` file inside the `controllers` folder and add the following code inside it:

```go
package controllers

import (
 "goblogart/inits"
 "goblogart/models"

 "github.com/gin-gonic/gin"
 "golang.org/x/crypto/bcrypt"
)
```

The code above declares the file as part of the `controllers` package and imports the `init`, `models`, `gin`, and `bcrypt` packages.

Next, add the following code to the file to create the `Signup` function for users:

```go
func Signup(ctx *gin.Context) {
 var body struct {
  Name     string
  Email    string
  Password string
 }

 if ctx.BindJSON(&body) != nil {
  ctx.JSON(400, gin.H{"error": "bad request"})
  return
 }

 hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

 if err != nil {
  ctx.JSON(500, gin.H{"error": err})
  return
 }

 user := models.User{Name: body.Name, Email: body.Email, Password: string(hash)}

 result := inits.DB.Create(&user)

 if result.Error != nil {
  ctx.JSON(500, gin.H{"error": result.Error})
  return
 }

 ctx.JSON(200, gin.H{"data": user})

}
```

The code above will validate the request data, hash the password, binds the data to the `user` variable, adds the `user` to the database, and returns the new user data.

### User Login

The `Login` function will be responsible for validating users and generating authorization cookies.

Add a `SECRET` variable inside the `.env` folder like so:

```env
SECRET="FGNMfvbnmVBNM23456789$%^&*(*&^%$%^&*&^%$%^&^%$%^&^%$#)"
```

The `SECRET` key will be used to encode and decode the token that will be generated after logging in.

> Note: The `SECRET` value can be any value that cannot be guessed easily.

Next, add the following code below the `Signup` function:

```go
func Login(ctx *gin.Context) {
 var body struct {
  Email    string
  Password string
 }

 if ctx.BindJSON(&body) != nil {
  ctx.JSON(400, gin.H{"error": "bad request"})
  return
 }

 var user models.User

 result := inits.DB.Where("email = ?", body.Email).First(&user)

 if result.Error != nil {
  ctx.JSON(500, gin.H{"error": "User not found"})
  return
 }

 err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

 if err != nil {
  ctx.JSON(401, gin.H{"error": "unauthorized"})
  return
 }

 // generate jwt token
 token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
  "id":  user.ID,
  "exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
 })

 tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

 if err != nil {
  ctx.JSON(500, gin.H{"error": "error signing token"})
  return
 }

 ctx.SetSameSite(http.SameSiteLaxMode)
 ctx.SetCookie("Authorization", tokenString, 3600*24*30, "", "localhost", false, true)
}
```

The code above will validate the user by finding the user in the database with the email, hashing and comparing the password, generating a `jwt` token, and setting a cookie that will expire in 30 days.

### The GetUsers Function

The `GetUsers` function will be responsible for returning all the users and their posts from the database. Add the following code to the `Login` function:

```go
func GetUsers(ctx *gin.Context) {
 var users []models.User

 err := inits.DB.Model(&models.User{}).Preload("Posts").Find(&users).Error

 if err != nil {
  fmt.Println(err)
  ctx.JSON(500, gin.H{"error": "error getting users"})
  return
 }

 ctx.JSON(200, gin.H{"data": users})

}
```

The code above will get all the users and their posts from the database.

### User Validation Middleware

After generating the token that will be used for authorization, you need to create a `validate` function that will be used to validate the token and expiration. Add the following code below the `GetUsers` function:

```go
func Validate(ctx *gin.Context) {
 user, err := ctx.Get("user")
 if err != false {
  ctx.JSON(500, gin.H{"error": err})
  return
 }
 ctx.JSON(200, gin.H{"data": "You are logged in!", "user": user})
}
```

The code above will use the middleware that you will create next to authenticate the user. Create a `authMiddleware.go` file inside the `middlewares` folder and add the following code inside it:

```go
package middlewares

import (
 "fmt"
 "goblogart/inits"
 "goblogart/models"
 "net/http"
 "os"
 "time"

 "github.com/gin-gonic/gin"
 "github.com/golang-jwt/jwt/v5"
)

func RequireAuth(ctx *gin.Context) {
 tokenString, err := ctx.Cookie("Authorization")

 if err != nil {
  ctx.JSON(401, gin.H{"error": "unauthorized"})
  ctx.AbortWithStatus(http.StatusUnauthorized)
  return
 }
 token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
  if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
   return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
  }

  return []byte(os.Getenv("SECRET")), nil
 })

 if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
  if float64(time.Now().Unix()) > claims["exp"].(float64) {
   ctx.JSON(401, gin.H{"error": "unauthorized"})
   ctx.AbortWithStatus(http.StatusUnauthorized)
   return
  }

  var user models.User
  inits.DB.First(&user, int(claims["id"].(float64)))
  if user.ID == 0 {
   ctx.JSON(401, gin.H{"error": "unauthorized"})
   ctx.AbortWithStatus(http.StatusUnauthorized)
   return
  }

  ctx.Set("user", user)
  fmt.Println(claims["foo"], claims["nbf"])
 } else {
  ctx.AbortWithStatus(http.StatusUnauthorized)
 }
 ctx.Next()
}
The code above will get the cookie from the request, decode and validate it, check if it is still valid, find the user in the database, and set the user variable.

### The `Logout` Function
The `Logout` function will be responsible for returning all the users and their posts from the database. Add the following code to the `useController` file below the `Validate` function:
```go
func Logout(ctx *gin.Context) {
 ctx.SetSameSite(http.SameSiteLaxMode)
 ctx.SetCookie("Authorization", "", -1, "", "localhost", false, true)
 ctx.JSON(200, gin.H{"data": "You are logged out!"})
}

```

The code above will delete the cookie that is being used by the logged in user and send back a message.

## User Routes

Now that the user controller and middleware functions are done, let's connect each of them to a route and test them.

### The Signup Route

Add the following code to the `main.go` file below the posts' routes:

```go
r.POST("/user", controllers.Signup)
```

Send a `POST` request to `locahost:3000/users` with the following JSON data:

```json
{
    "name": "Adams Adebayo",
    "email":"adea@ml.com",
    "password": "12345"
}
```

You should get a response that looks like this:

![sign up result](https://dev-to-uploads.s3.amazonaws.com/uploads/articles/0591ylc2cdcfdgdac8y2.png)
The purple box - type of request.
The yellow box - URL of the route.
The red box - JSON data sent to the server.
The green box - result from the server.

### The Login Route

Add the following code to the `main.go` file below the previous route:

```go
r.POST("/login", controllers.Login)
```

Send a `POST` request to `locahost:3000/users` with the following JSON data:

```json
{
    "email":"adea@ml.com",
    "password": "12345"
}
```

You should get a result with a cookie like so:

![login result](https://dev-to-uploads.s3.amazonaws.com/uploads/articles/4yncyx2mrlpd3edxkfm0.png)

The purple box - type of request.
The yellow box - URL of the route.
The red box - JSON data sent to the server.
The green box - result from the server.

### The Validate Route

Add the following code to the `main.go` file below the previous' route:

```go
r.POST("/auth", controllers.Validate)
```

This request doesn't send any data to the server. Send a `GET` request to `locahost:3000/auth` and you should get a response that looks like this:

![validate result](https://dev-to-uploads.s3.amazonaws.com/uploads/articles/nxdh3j0vd7dyemzdw6kb.png)

### The GetUsers Route

Add the following code to the `main.go` file below the previous route:

```go
r.GET("/users", controllers.GetUsers)
```

This request doesn't send any data to the server. Send a `GET` request to `locahost:3000/users` and you should get a response that looks like this:

![get users result](https://dev-to-uploads.s3.amazonaws.com/uploads/articles/pjuj7f2lv3p6o43iu0e2.png)

### The Logout Route

Add the following code to the `main.go` file below the previous route:

```go
r.GET("/logout", controllers.Logout)
```

This request doesn't send any data to the server. Send a `GET` request to `locahost:3000/logout` and you should get a response that looks like this:

![logout result](https://dev-to-uploads.s3.amazonaws.com/uploads/articles/307glsxbz8serpqhr5r9.png)

### Relationship Test

Now that all the user functions and authentications are working, let's test the `has many` relationship. Log in with the user that you created earlier and create a post
