# Shanghai

> Build hierarchies of container images using a simple YAML file

## Usage

Example of the `Shangaifile` and container images could be found in the [example](./example/)
directory.

### Build images

```sh
$ shanghai build -i alps
Building 'example.com/alps:latest'  ✔ 
Building 'example.com/alps-py:latest'  ✔ 
Building 'example.com/alps-node:latest'  ✔ 
```

## Roadmap

- [x] Build images
- [x] Push images
- [x] GitHub Actions
  - [x] Static binaries
  - [x] Completion files
- [ ] YAML file schema

## License

[MIT](./LICENSE)
