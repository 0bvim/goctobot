APP_NAME := goctobot
DIR=bin/
APPATH=$(DIR)$(APP_NAME)
MAIN_PATH=cmd/application/main.go

all: dir
	@go build -gcflags='all=-N -l' -o bin/$(APP_NAME) $(MAIN_PATH)
	@echo "\033[1;34mGoctoBot Compiled\033[0m"
	$(call print_usage)

dir:
	@mkdir -p $(DIR)

fclean:
	@rm $(APPATH)
	@echo "Removed App"

re: fclean all

debug:
	@dlv debug $(MAIN_PATH) -- $(ARGS)

run:
	@./$(APPATH) $(ARGS)

tests:
	@go test -v ./...

# Define the function to print the usage guide with colors
define print_usage
	@echo "\033[1;32mUsage:\033[0m"
	@echo "  \033[1;36m./goctobot <command> [username]\033[0m"
	@echo ""
	@echo "\033[1;32mCommands:\033[0m"
	@echo "  \033[1;36mfollow [username]\033[0m    - Follow all followers of the specified user."
	@echo "  \033[1;36munfollow\033[0m            - Unfollow who do not follow back."
	@echo "  \033[1;36mfollowing\033[0m          - Shows count of users you follow."
	@echo "  \033[1;36mfollowers\033[0m          - Show count of your followers."
	@echo "  \033[1;36mstatus\033[0m          - Show both followers and following."
endef

.PHONY: all fclean tests re debug dir
