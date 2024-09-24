# Projeto de Gerenciamento de Carros

Este projeto é uma API para gerenciamento de carros, com suporte a consultas detalhadas e variações de preço, utilizando o framework Gin e o ORM GORM em Go. A API conecta-se a um banco de dados PostgreSQL para armazenar e recuperar informações sobre carros, suas variações de preço e outros detalhes.

## Requisitos

- [Go](https://golang.org/dl/) instalado
- Conexão com um banco de dados PostgreSQL (configurável no arquivo `main.go`)

## Como Executar o Projeto

Para rodar a aplicação, siga os seguintes passos:

1. Clone o repositório:

   ```bash
   git clone <url-do-repositorio>
   cd <nome-do-diretorio>
   ```

2. Instale as dependências:

   ```bash
   go mod tidy
   ```

3. Configure as credenciais do banco de dados PostgreSQL no arquivo `main.go`:

   ```go
   dsn := "host=seu-host user=seu-usuario password=sua-senha dbname=seu-banco port=sua-porta sslmode=require"
   ```

4. Execute a aplicação:

   ```bash
   go run main.go
   ```

5. A API estará disponível em:
   ```
   http://localhost:8080
   ```

## Endpoints Principais

### 1. Listar Carros

- **GET** `/carros`
- Retorna uma lista de carros cadastrados no banco de dados.

### 2. Listar Variações de Preço de Carros (Com Paginação)

- **GET** `/carros/variacoes?page=1&limit=100`
- Exibe as variações de preço dos carros, com suporte a paginação.
- **Parâmetros:**

  - `page`: número da página (padrão: 1)
  - `limit`: número de itens por página (padrão: 10, máximo sugerido: 100)

    Exemplo de requisição:

    ```
    http://localhost:8080/carros/variacoes?page=1&limit=100
    ```

### 3. Listar Detalhes dos Carros

- **GET** `/carros/detalhados`
- Retorna informações detalhadas de preços e datas de referência dos carros.

## Observações

- A API utiliza paginação para otimizar o carregamento de dados, especialmente em grandes volumes de registros.
- O limite padrão de itens por página é 10, mas você pode ajustá-lo passando o parâmetro `limit` na URL, até o valor sugerido de 100 itens por página.
