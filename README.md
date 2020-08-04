# Introduction

go-crud-starter helps to kickstart your go app in style.

It aims to use very less external dependancy(only 1 for routing, which was chosen carefully but can easily be replaced) but at the same time give a robust code organisation for any go app

It includes

1. bootstrap 4.
2. go-chi (you will need to `go get` )
3. Some middleware from go-chi/chi/middleware - logger, recoverer etc.
4. sqlite

# Usage

1. `git clone git@github.com:kgthegreat/go-crud-starter.git app-name`
2. `cd app-name`
3. `git remote rm origin`
4. Replace all occurence of go-crud-starter with app-name in all files. This is how you do this in emacs https://stackoverflow.com/questions/270930/using-emacs-to-recursively-find-and-replace-in-text-files-not-already-open
5. Populate app.json with all configs such as db name. Add the absolute path to project directory to others
6. Without the above, the app will fail. You can see by doing `go build`
7. `go get` any missing dependencies 
7. `go build`
8. `./app-name`

