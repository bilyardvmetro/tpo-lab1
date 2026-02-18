.PHONY test:
test:
	@echo "Running tests for task1..."
	@cd task1 && go test ./task1/.. -v -cover
	@echo "Running tests for task2..."
	@cd task2 && go test ./task2/.. -v -cover
	@echo "Running tests for task3..."
	@cd task3 && go test ./task3/.. -v -cover