DB_USER    := root
DB_PASS    := root
DB_NAME    := go_orm_gen_setup_db
CONTAINER  := mysql_gen_query
DDL_FILE   := ./migrations/user.ddl

gen:
	@echo "Generating ORM code..."
	@cd ./tools && go run ./main.go
	@echo "ORM code generation completed."

setup_gen: stop start migrate

start:
	@echo "Starting MySQL Docker container..."
	docker run --name $(CONTAINER) \
		-e MYSQL_ROOT_PASSWORD=$(DB_PASS) \
		-p 3306:3306 \
		-d mysql:8.0
	@echo "Waiting for MySQL to be ready..."
	@sleep 20
	@docker exec -i $(CONTAINER) \
		mysql -u$(DB_USER) -p$(DB_PASS) \
		-e "CREATE DATABASE IF NOT EXISTS $(DB_NAME);"

migrate:
	@echo "Running migration..."
	@docker exec -i $(CONTAINER) \
		mysql -u$(DB_USER) -p$(DB_PASS) $(DB_NAME) < $(DDL_FILE)
	@echo "Migration completed."

stop:
	@echo "Stopping and removing Docker container..."
	@docker stop $(CONTAINER) >/dev/null || true
	@docker rm -v $(CONTAINER) >/dev/null || true
