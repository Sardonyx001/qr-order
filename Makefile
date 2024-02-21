FRONTEND_PATH = $(PWD)/frontend
BACKEND_PATH = $(PWD)/backend

run: 
	@if [ -d "$(FRONTEND_PATH)" ]; then cd $(FRONTEND_PATH) && pnpm run dev; fi
	@if [ -d "$(BACKEND_PATH)" ]; then cd $(BACKEND_PATH) && $(MAKE) run; fi

# build: 
# 	@if [ -d "$(FRONTEND_PATH)" ]; then cd $(FRONTEND_PATH) && pnpm run build; fi
# 	@if [ -d "$(BACKEND_PATH)" ]; then cd $(BACKEND_PATH) && $(MAKE) build; fi