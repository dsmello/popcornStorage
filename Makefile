env:
	@echo "Setting up environment"
	@docker-compose -f docker-compose.yml up -d
	@echo