# Sistema de Gerenciamento de Carros - Backend

Este é um projeto backend em Go utilizando o framework **Gin** e o ORM **GORM**, com integração a um banco de dados PostgreSQL. O sistema oferece funcionalidades para listar, buscar e filtrar carros, suas variações e detalhes.

## Funcionalidades

- **Listagem de Carros**: Obtenha uma lista de carros com paginação.
- **Busca por Detalhes de Carros**: Visualize os detalhes dos carros com base em preço, data de referência e outros critérios.
- **Filtragem de Carros**: Busque carros com filtros como tipo, ano, marca, modelo e combustível.
- **Preços de Carros**: Filtre por intervalo de preços e data de referência.

## Tecnologias Utilizadas

- **Linguagem**: Go
- **Framework Web**: Gin
- **ORM**: GORM
- **Banco de Dados**: PostgreSQL
- **Bibliotecas**:
  - `github.com/gin-gonic/gin` - Framework HTTP para Go.
  - `gorm.io/gorm` - ORM para Go.
  - `gorm.io/driver/postgres` - Driver PostgreSQL para GORM.

## Instalação e Execução

### Pré-requisitos

- **Go** (versão 1.18+)
- **PostgreSQL** (configurado e rodando)

### Passos para Instalar

1. Clone este repositório:

   ```bash
   git clone https://github.com/seu-usuario/go-backend-carros.git
   ```

2. Navegue até o diretório do projeto:

   ```bash
   cd go-backend-carros
   ```

3. Instale as dependências do projeto:

   ```bash
   go mod tidy
   ```

4. Configure as credenciais de conexão com o banco de dados PostgreSQL no código (variável `dsn` no arquivo `main.go`).

5. Execute a aplicação:
   ```bash
   go run main.go
   ```

A aplicação estará rodando no endereço `http://localhost:8080`.

## Endpoints

### GET `/carros`

Lista os carros com suporte a paginação.

**Query Parameters**:

- `limit`: Número de resultados por página. Exemplo: `limit=10`
- `page`: Número da página. Exemplo: `page=2`

**Exemplo de Uso**:

```
GET http://localhost:8080/carros?limit=10&page=1
```

### GET `/carros/detalhados`

Lista detalhes dos carros.

**Query Parameters**:

- `limit`: Número de resultados por página. Exemplo: `limit=10`
- `page`: Número da página. Exemplo: `page=2`

**Exemplo de Uso**:

```
GET http://localhost:8080/carros/detalhados?limit=10&page=1
```

### GET `/carros/variacoes`

Lista variações de preços dos carros.

**Query Parameters**:

- `limit`: Número de resultados por página. Exemplo: `limit=10`
- `page`: Número da página. Exemplo: `page=2`

**Exemplo de Uso**:

```
GET http://localhost:8080/carros/variacoes?limit=10&page=1
```

### GET `/carros/search`

Busca carros com base em filtros específicos.

**Query Parameters**:

- `tipo`: Filtra por tipo de carro (ex: "SUV", "Sedan").
- `ano`: Filtra por ano de fabricação.
- `marca`: Filtra por marca (ex: "Toyota", "Ford").
- `modelo`: Filtra por modelo.
- `combustivel`: Filtra por tipo de combustível (ex: "Gasolina", "Diesel").

**Exemplo de Uso**:

```
GET http://localhost:8080/carros/search?tipo=SUV&ano=2020&marca=Toyota
```

### GET `/carros/precos`

Busca preços detalhados dos carros com filtros opcionais.

**Query Parameters**:

- `precoMin`: Filtra carros com preço mínimo.
- `precoMax`: Filtra carros com preço máximo.
- `dataReferencia`: Filtra pela data de referência.

**Exemplo de Uso**:

```
GET http://localhost:8080/carros/precos?precoMin=20000&precoMax=50000&dataReferencia=2024-09-01
```

## Modelos

### Carro

Representa um carro básico com informações como tipo, ano, marca, modelo e combustível.

### CarroDetalhado

Contém informações detalhadas sobre preços de carros, incluindo data de referência e valor.

### CarroVariacao

Representa variações de preços de carros ao longo do tempo.

## Contribuição

Contribuições são bem-vindas! Sinta-se à vontade para abrir issues ou enviar pull requests.

## Licença

Este projeto está licenciado sob a [MIT License](LICENSE).

---

Projeto backend para gerenciamento de carros com funcionalidades de busca e listagem.
