# Go Hexagonal Architecture

I create this template for try a new folder organization grouped by domain with sub folders emulating hexagonal architecture folder for every domain.

The objective is make a project easier to maintain, if you want to change a controller of the user, you can easily go to :
```internal/users/infrastructure/controllers```

Inside this folder you will find one file for every action:
```create.go```
```delete.go```
```update.go```
```get_by_id.go```
```get_all.go```