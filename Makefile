APP_NAME := goctobot
MAIN_PATH=main.go

all: dir
	@go build -gcflags='all=-N -l' -o $(APP_NAME) $(MAIN_PATH)
	@echo "\033[1;34mGoctoBot Compiled\033[0m"
	@./$(APP_NAME) --help

fclean:
	@rm $(APP_NAME)
	@echo "Removed App"

re: fclean all

debug:
	@dlv debug $(MAIN_PATH) -- $(ARGS)

run:
	@./$(APP_NAME) $(ARGS)

tests:
	@go test -v ./...

# Define the function to print the usage guide with colors
# Deprecated, cobra is used instead
define print_usage
	@echo "\033[1;32mUsage:\033[0m"
	@echo "  \033[1;36m./goctobot <command> [username]\033[0m"
	@echo ""
	@echo "\033[1;32mCommands:\033[0m"
	@echo "  \033[1;36mfollow [username]\033[0m    - Follow all followers of the specified user."
	@echo "  \033[1;36munfollow\033[0m             - Unfollow who do not follow back."
	@echo "  \033[1;36mfollowing\033[0m            - Shows count of users you follow."
	@echo "  \033[1;36mfollowers\033[0m            - Show count of your followers."
	@echo "  \033[1;36mstatus\033[0m               - Show both followers and following."
endef

.PHONY: all fclean tests re debug dir
