test:
	go test ./...

docs:
	@echo "## Loot\n" > godoc.txt
	@echo "Go SDK for Path of Exile APIs\n" >> godoc.txt
	@echo "## Character\n" >> godoc.txt
	@godoc github.com/willroberts/loot/character >> godoc.txt
	@echo "## Forum\n" >> godoc.txt
	@godoc github.com/willroberts/loot/forum >> godoc.txt
	@echo "## Stash\n" >> godoc.txt
	@godoc github.com/willroberts/loot/stash >> godoc.txt
