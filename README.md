# Loot

Loot is a Go SDK for the legacy forum-based Path of Exile APIs. Before the trade
API was introduced, items were found by scraping forum pages. This method is no
longer advisable.

## DEPRECATED

For a modern, feature-complete client for the Path of Exile API, please use
[github.com/willroberts/poeapi](https://github.com/willroberts/poeapi) instead.

## Contents

* `character`: A client for retrieving character info and items from the PoE website.
* `forum`: Tools for parsing items directly from forums threads.
* `stash`: A client for the official public stash tab API.

## Documentation

Generated documentation can be found in `godoc.txt`. Documentation
can be regenerated by running `make docs`.

## Running tests

```bash
make test
```

## To do

- [ ] Improve item models for stash API
- [ ] Build something with the poller (live search?)
- [ ] Build an indexer
- [ ] Implement the character-window API
