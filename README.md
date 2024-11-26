# go-clima-cep:
Projeto de Laboratório Cloud Run do curso Go Expert.

### Comandos:

#### Antes de executar os comandos para rodar o projeto, garanta que sua APIKEY do serviço weatherapi esteja configurada no arquivo .env.

#### Executa a aplicação no container
```
make up
```

#### Parar o container
```
make stop
 ```

#### Executa testes unitarios
 ```
make test
```

#### Executa a aplicação localmente.
```
make run
```

#### Exemplo de uso local

```
curl http://localhost:8080/weather/89201001

Retorno esperado: 
{ 
    "temp_C":"26.10",
    "temp_F":"78.98",
    "temp_K":"299.10"
}
```
#### Exemplo de uso Cloud Run:

```
curl https://go-clima-cep-892172836551.southamerica-east1.run.app/weather/89201001

Retorno esperado: 
{ 
    "temp_C":"26.10",
    "temp_F":"78.98",
    "temp_K":"299.10"
}

```


#### Endereço ativo do Cloud Run:

- https://go-clima-cep-892172836551.southamerica-east1.run.app