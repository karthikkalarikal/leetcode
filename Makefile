MODULES = \
	./question-367 \
	./question-441 \
	./question-744 \
	./question-1539 \
	./question-35 \
	./question-852 \
	./question-1351 \
	./question-349 \
	./question-350

.PHONY: test

test:
	@for module in $(MODULES); do \
		echo "Running test for $$module"; \
		go test $$module; \
	done