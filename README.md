# API Go com PostgreSQL e Docker 🚀

Esta é uma API REST desenvolvida em Go como projeto de estudo para portfólio. A aplicação utiliza o framework Gin, PostgreSQL como banco de dados e é totalmente containerizada com Docker.

## 🛠️ Tecnologias Utilizadas

- **Go (Golang)**: Linguagem principal.
- **Gin Gonic**: Framework web para roteamento e middleware.
- **PostgreSQL**: Banco de dados relacional.
- **Docker & Docker Compose**: Gerenciamento de containers.
- **sqlx / lib/pq**: Driver para conexão com o banco de dados.

## 🏗️ Arquitetura do Projeto

O projeto segue uma estrutura organizada em camadas para facilitar a manutenção e escalabilidade:

- `cmd/`: Ponto de entrada da aplicação (`main.go`).
- `controller/`: Camada de controle (HTTP handlers).
- `usecase/`: Camada de regras de negócio.
- `repository/`: Camada de persistência/acesso ao banco de dados.
- `model/`: Definição das estruturas de dados (Entities).
- `db/`: Configurações de conexão e scripts SQL.

## 🚀 Como Executar

### Pré-requisitos
- Docker e Docker Compose instalados.

### Passo a Passo

1. **Clonar o repositório:**
   ```bash
   git clone https://github.com/alisonwillf/api-go-postgres.git
   cd api-go-postgres
   ```

2. **Subir os containers:**
   ```bash
   docker compose up -d
   ```
   *Este comando irá baixar as imagens necessárias, configurar o banco de dados PostgreSQL e iniciar a aplicação Go na porta 8000.*

3. **Verificar se a aplicação está rodando:**
   Acesse no seu navegador ou Postman: `http://localhost:8000/ping`

## 🔌 Endpoints da API

- **GET `/ping`**: Verifica a saúde da aplicação.
- **GET `/products`**: Lista todos os produtos cadastrados.
- **GET `/product/:id`**: Busca um produto específico pelo ID.
- **POST `/product`**: Cria um novo produto.
  - Exemplo de Body (JSON):
    ```json
    {
      "product_name": "Teclado Mecânico",
      "price": 250.50
    }
    ```

## 📝 Banco de Dados

O banco de dados é inicializado automaticamente com a tabela `product` através do arquivo `db/init.sql` montado no container do PostgreSQL.

---
Desenvolvido por [Alison Ferreira](https://github.com/alisonwillf) 👨‍💻
