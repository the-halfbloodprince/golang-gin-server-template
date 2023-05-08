# Basic Server Template with Gin

this is a basic server template to start developing a golang server setup with Gin.

- implements a modular architecture
- support for code division into
  - modules: modules are added as directories inside the modules folder
  - controller: controllers are housed in a separate `controller.go` file in the module directory
  - services: services which are to be injected into the application through controllers go inside the `services.go` file inside the module's directory

To create a new module:
- Add a new directory corresponding to your module
- Add a new `module.go` file
- Add a new Module struct with all the fields provided
- Register it in the `config/modules.go` file
- Put all the controllers and services in their own files
- Done ✅