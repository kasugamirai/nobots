package main

import (
	"fmt"

	_ "entgo.io/ent/dialect"
	_ "entgo.io/ent/dialect/sql"
	_ "entgo.io/ent/dialect/sql/sqlgraph"
	_ "entgo.io/ent/schema/field"
	_ "github.com/go-chi/chi/v5"
	_ "github.com/go-chi/chi/v5/middleware"
	_ "github.com/go-chi/cors"
	_ "github.com/go-chi/httplog"
	_ "github.com/ogen-go/ogen"
	_ "github.com/ogen-go/ogen/gen"
	_ "github.com/rs/zerolog"
	_ "github.com/rs/zerolog/log"
	_ "github.com/urfave/cli/v2"
)

func main() {
	fmt.Println("docker build cache for deps library")
}
