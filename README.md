# xy-inc

Projeto para cadastro e recuperação Pontos de Interesse (POI's).

## Como criar o ambiente?

1. Clone o repositório
2. Execute os testes.
3. Instalar as dependencias

```console
git clone https://github.com/n1lux/xy-inc-go.git xyinc
go get -v -u github.com/gorilla/mux
got get -v -u github.com/gorilla/gorm
got get -v -u github.com/mattn/go-sqlite3
cd xyinc
```

## Como executar o servidor

```console
go run main.go
```

## Como testar os serviços
Dados para o envio do json

```bash
{
    "name": <string>,
    "x": <int>,
    "y": <int>,
}

```
### No navegador

#### Para consultar todo os POI's cadastrados
```bash
[GET] http://localhost:8080/api/pois/
```


#### Para buscar os POI's com base nos parametros informados
```bash
[GET] http://localhost:8080/api/pois/search?x=20&y=10&d-max=10
```



### Utilizando curl
#### Para consultar todos os POI's cadastrados:
```console
curl -X GET -H "Accept:application/json" http://localhost:8000/api/pois/
```

#### Para criar um poi:
```console
curl -i -X POST -H "Content-Type:application/json" http://localhost:8080/api/pois -d '{"name":"teste poi curl", "x": 25, "y": 10}'
```

