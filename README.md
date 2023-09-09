## Go Tri

This is a basic CLI application created to get a bit more time spent with Go. The application represents a very basic todo app on the CLI that was developed using [Cobra](https://cobra.dev).

This application has the following features:

* Add Multiple Items to your todo list:
```
go_tri add "Item 1" "Item 2"
```

* List items in your todo list:
```
go_tri list
```

* Mark Item as done
```
go_tri done 1
```

* Manage Multiple lists with a flag
```
go_tri add "Alternative Item 1" --datalist ./alternatelist.json
```
* And more

You can see all the commands and help with all the commands by just running:

```
go_tri help
```
