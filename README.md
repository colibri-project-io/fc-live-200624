# fc-live-200624
Repositório de fontes da live do dia 20/06/2024 da Full Cycle


## Chamadas ao serviço

### GetAll
```bash
curl 'http://localhost:8080/public/cities?page=1&pageSize=10&name=Salvador&uf=BA'
```

### GetById  
```bash
curl 'http://localhost:8080/public/cities/3d6b712a-2f37-11ef-ad27-53a0959f9cd9'
```

### GetByCep
```bash
curl 'http://localhost:8080/public/cities/cep/42800941'
```

### CreateCity
```bash
curl --request POST 'http://localhost:8080/public/cities' \
--header 'Content-Type: application/json' \
--data '{
    "name": "Salvador",
    "uf": "BA"
}'
```

### DeleteCityById
```bash
curl --request DELETE 'http://localhost:8080/public/cities/3d6b712a-2f37-11ef-ad27-53a0959f9cd9'
```
