# Euro2016 CLI

## Installation

`go get github.com/xlucas/euro2016`

## Configuration

The underlying API has a rate limit of 50 calls per day for a given IP.
To overcome this limitation, register a free account [here](http://api.football-data.org/register).

Register the token you received by email in `$HOME/.euro2016.json` :

```json
{"token": "<token_value>"}
```

## Usage

### Display full competition schedule

```
euro2016 schedule full
```

### Display games of the day

```
euro2016 schedule today
```

### Display live games

```
euro2016 schedule status IN_PLAY
```

### Show rankings

```
euro2016 rankings full
```

### Show ranking for a specific group

```
euro2016 rankings group A
```

### Print teams

```
euro2016 teams
```


## License

MIT
