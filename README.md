# Injector

This tool injects .env values to an executables

# Usage

- Write .env

```
FOO_BAR=TEST
```

- Run

```
> go run injector.go cmd /c echo %FOO_BAR%
```
