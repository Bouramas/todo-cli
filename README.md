# Todo CLI Tool

This is a simple CLI application written in GO. It's a simple Task Manager, with the purpose of 
exploring the CLI capabilities of the language, rather than creating a CLI application. It is my take on Jonathan Calhoun's [Exercise #7: CLI Task Manager](https://courses.calhoun.io/courses/cor_gophercises).

## Dependencies
This application relies on the following external packages:

    github.com/boltdb/bolt v1.3.1
    github.com/mitchellh/go-homedir v1.1.0
    github.com/spf13/cobra v1.8.0

## Usage

Build and install the go executable:

```bash
 go build .
 go install .
```
Run todo-cli
```bash
 todo-cli help
```

Available Commands:
```bash
  add         Adds a task to your task list.
  do          Marks a task as complete
  help        Help about any command
  list        Lists all of your tasks.
  rm          Remove a task from the list
```