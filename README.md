# Paquetes instalados

```sh
go install github.com/air-verse/air@latest
```

## Limpieza de cache 
```sh 
go clean -cache
go clean -modcache
``` 

## Docker 

### Build

Para hacer la build con docker 

```sh
docker build --platform linux/amd64 --build-arg VERSION=any --build-arg BUILD=`date +%FT%T%z` -t rps .
```

### run

```sh
docker run --rm -it -p 8080:8080 --name pip rps
```

## Arrancar air 

Para agregar air a nuestro proyecto ejecutamos 

```sh
air init
```

Para arrancar el servicio con air ejecutamos 
```sh
air
```

Para configurar e integrr vscode con `iar`, modificamos el fichero de configuración `.air.toml` con el siguiente contenido 

```toml
[build]

	# add the templ generate when working with templ
	cmd = "templ generate internal/views/ && CGO_ENABLED=0 go build -gcflags=all=\"-N -l\" -o ./tmp/main ./cmd"

	# exclude file generated by templ, otherwise infinite loop will occur
	exclude_regex = ["_test.go", ".*_templ.go"]

	# instead run the binary, we run the dlv
	# we can add flag to our application after the '--'
	full_bin = "dlv exec ./tmp/main --listen=127.0.0.1:8080 --headless=true --api-version=2 --accept-multiclient --continue --log -- "

	# include the .templ file
	include_ext = ["go", "tpl", "tmpl", "html", "templ"]
```

Una vez configurado el archivo, ejecutamos `air` en el terminal para arrancar el servidor y posteriormente `F5` para atachar el vscode al proyecto