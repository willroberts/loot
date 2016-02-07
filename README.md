# Loot

A Path of Exile forum scraper for items

## Testing

### Running unit tests

```
go test ./...
```

### Showing items from a thread

```
go run main.go -thread 1566069  # or any thread
```

## To Do

* Parse socketed items (there's a link to the structure in items.go)
* Update tests for package forum
* Write tests for package items
