# Define the Go compiler and linker
GO := go
GOLDFLAGS := -ldflags="-s -w"


# Build go app
build:
   $(GO) build $(GOLDFLAGS) -o sqlc_openapi


run:
   ./sqlc_openapi


# Define a "clean" target to remove generated files
clean:
   rm -f sqlc_openapi

   
