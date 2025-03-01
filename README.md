# Simple Back-end for Simple Bank

If you're reading this README.md, it's likely that you're reading what Ive learned from a project.
</br>

I've been learning Kubernetes and Docker in theory, this project is to make my hands dirty with Cloud computing and DevOps.
</br>

Despite that, along the way, I also picked up and used `CI/CD` with `Github Actions`, `gRPC` vs `HTTP`, `Go` with `Gin`, `AWS`, `PostgreSQL`, `JWT` vs `PASETO`, etc.
Overall, this project was a good learning experience.
</br>

Shout out to [`Tech School`](https://bit.ly/m/techschool) for making the guide.
</br>

## Architecture

<img src="https://github.com/matoanbach/simple-bank/images/architecture.png"/>

## API documentation using Swagger UI

## Database

Alright, I'm not going to write an interface for PostgreSQL from scratch in Go (it's fun, but painful). So `sqlc` is here to help us, minimizing mistakes along the way.

</br>

So, I need a yaml for that ... like below:

```yaml
version: "2"
sql:
  - schema: "./db/migration" # Path to the schema directory
    queries: "./db/query" # Path to the queries directory
    engine: "postgresql" # Database engine
    gen:
      go:
        package: "db"
        out: "./db/sqlc" # Directory for generated Go code
        emit_json_tags: true
        emit_empty_slices: true
        emit_interface: true
    rules:
      - sqlc/db-prepare
```

</br>
Run this to generate go code, so that we can talk with PostgresSQL

```bash
sqlc generate
# or
make sqlc # it's the Makefile, which is basically the same
```

## RESTful with Gin

`Gin` is a popular web frameworks for Go. `JWT` and/or `PASETO` are used to enhance security by creating and verifying tokens. Later on, we can use this to make token-based sessions.

</br>

```go
func runGinServer(config util.Config, store db.Store) {
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal().Msg("cannot create server")
	}
	err = server.Start(config.HTTPServerAddress)
	if err != nil {
		log.Fatal().Msg("cannot start the server")
	}
}
```

```go
// routers down below
func (server *Server) setupServer() {
	router := gin.Default()

	router.POST("/users", server.createUser)
	router.POST("/users/login", server.loginUser)
	router.POST("/tokens/renew_access", server.renewAccessToken)

	authRoutes := router.Group("/").Use(authMiddleWare(server.tokenMaker))
	authRoutes.POST("/accounts", server.createAccount)
	authRoutes.GET("/accounts/:id", server.getAccount)
	authRoutes.GET("/accounts", server.listAccount)
	authRoutes.POST("/transfers", server.createTransfer)
	server.router = router
}
```

## Containerize it with k8s and docker

Here we go, this phase is what I've been waiting for. This is where we set `Github Actions` for `CI/CD`, `Docker` and `Kubernetes` for containerization and orchestration.
</br>

For cloud provider, we work with

<ul>
<li>ECR (Elastic Container Registy) to store docker images</li>
<li>Secret Managers to retrieve env variables</li>
<li>EC2 instances to run worker nodes</li>
<li>EKS (Elastic Kubernetes Service) to run kubernetes</li>
<li>RDS (Managed relational database service) to run PostgreSQL virtually, but most of the time, we use local PostgreSQL</li>
<li>Of course, IAM to manage authentication and authorization in AWS</li>
</ul>

## gRPC
