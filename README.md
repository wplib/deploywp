# deploywp
A Golang-based app for deploying a WordPress site to Pantheon


## Basic premise
This package is a fork from JsonToConfig initially and provides a simple framework of "helpers", with all methods visible in both GoLang native and GoLang templates.

## Structure
helper packages are defined within the jsonTemplate/helpers directory. All methods and functions are pulled in automatically on build using `pkgreflect`.

Each helper directory is divided into `frontend` and `backend` methods/functions.


### Frontend
Front end functions/methods are to be used within Templates, (although not exclusively), and have a copy of the backend Type structures. They are all contained with the files with a `helpers` prefix.

Since they are used within templates, they are more "vague" with determining specific types and reflection helpers are used extensively to aide both error correction and recovery.

Front end functions begin with the name "Helper" and will be exported and made available within templates. Any other function will be ignored within templates.

Methods coming from these functions, (of course), are available as normal and will ALWAYS return a `ux.State` structure.

For example:

- `HelperNewGit()` is a function that will appear within templates as `NewGit()`.
- This function returns a `HelperGit` structure, (which is a type copy of `TypeGit`).
- The `HelperGit` structure has a bucket load of methods, (Open(), SetUrl(), GetBranch(), etc.
- The return of this methods will ALWAYS be a `ux.State` structure.


### Backend
Back end functions/methods perform the low level work of a specific helper. There's usually a 1-1 match of backend to frontend, but some frontend functions may perform more work.

All backend functions/methods that start with an `Is` will ALWAYS return a bool type. There are also always a matching `IsNot` to all `Is` functions/methods.


### Building and Testing (native)
The tests directory contains a template and json file for testing the framework.

`make build`

`make test`


### Building and Testing (docker)

`make build-docker`

`make test-docker`

