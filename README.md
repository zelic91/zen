# zen-erating boilerplates

I always need to Google everytime I want to have configs for Docker, Kubernetes, ... It's time-consuming and to be honest, I couldn't remember their syntax very well. That's why zen was born.

## Install

Make sure you already installed Go, then install zen binary

```
go install github.com/zelic91/zen
```

In order to run all Go binaries globally, you need to include `/go/bin` in the PATH as (in ~/.zshrc for example):

```
export PATH="$GOPATH/bin:$PATH"
```

## Usage
1. Create a new project with `zen new <project-name>`. This command will create a new folder and a sample YAML config file.
2. Edit the `zen.yaml` file.
3. Generate configs and stubs with `zen run`.
4. Edit the according configs in the generated files where appropriates and follow the instructions in generated `README.md` file.
5. Profit!!!

## TODO
- [ ] Add the ability to check for existing package/file and only generate if missing.
- [ ] Unit tests
- [ ] Redis client
- [ ] Remote logging
- [ ] Stubbing

## Notes

This is a very limited package which is built for my own need, you might find it very silly and clumpsy. If so, ignore it. Thanks.







