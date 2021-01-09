# todogolist.

A CLI to control your life.

## Installation.

- Clone the repo.
- Add the following yaml file in this path:
```bash
    $HOME/.togo/config.yaml
```
- Fill the config.yaml file, with the db URI and dialect you want. Supported dialects SQLlite, Postgres and MySQL. Use the config.example.yaml as an example.
- You are ready.

## Usage
```bash
./todogolist -h
```

```
The CLI to control your life.
You could easily use it create, update, get, list and delete your daily tasks.

Usage:
  todogolist [command]

Available Commands:
  create      Crete a new task
  delete      Delete all task for a given date or a given label
  get         Retrieve tasks
  help        Help about any command
  list        List all task for a given date or a given label
  setup       Set up/update all the needed DB model

Flags:
  -h, --help   help for todogolist

Use "todogolist [command] --help" for more information about a command.
```