hello:
	echo "Hello"
# Define the run-server target
run-server:
	cd backend && go run cmd/server/main.go

# Optional: Add a default target
.PHONY: all
all: run-server