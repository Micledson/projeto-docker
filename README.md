# Projeto Docker - API em Go

Este projeto é uma API desenvolvida em Go que utiliza o PostgreSQL como banco de dados.

## Pré-requisitos
Antes de executar o projeto, é necessário configurar o arquivo `.env`. Para isso, faça uma cópia do arquivo 
`.env.example` e renomeie-a para `.env`, colocando-o na raiz do projeto.

## Como executar
Após configurar o `.env`, você pode rodar o projeto utilizando o Docker. Para subir os containers da API e do 
banco de dados PostgreSQL, execute o comando:

```bash
docker compose up --build
```

## API
A API tem como finalidade o gerenciamento de ToDos, nela temos os seguintes métodos disponíveis:
- List [GET]: Lista todos os ToDos ativos
- GetByID [GET]: Retorna o ToDo com o ID especificado
- Create [POST]: Cria um ToDo
- Update [PUT]: Atualiza a descrição de um ToDo
- Delete [PUT]: Deleta um ToDo
- Restore [PUT]: Restaura um ToDo deletado

Para mais detalhes sobre os endpoints existe um swagger, nele também é possível fazer as requisições.
Para acessa-lo use o endpoint: http://localhost:8000/api/docs