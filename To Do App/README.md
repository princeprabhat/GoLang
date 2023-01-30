# **Todo Application**

A simple todo list application built with Go, It uses cobra library as command line interface.

## **Installation**

1. Clone the repository

```go
git clone https://github.com/rama-kairi/todo.git
```

1. Install the dependencies

```go
go mod download
```

1. You can directly install and use with `todo` command

```go
go install github.com/rama-kairi/todo
```

## **Usage**

The application has four commands:

- **`add`** : Add a new todo to the list
- **`remove`** : Remove a todo from the list
- **`list`** : List all todo's from the list
- **`rmall`** : Remove all todo's from the list

### **Examples**

```go
go run main.go add Learn PostgreSQL
```

```go
go run main.go remove 1
```

```go
go run main.go list
```

```
go run main.go rmall
```

## **Contributing**

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## **License**

**[MIT](https://choosealicense.com/licenses/mit/)**
