# Stock Manager Go

API de gerenciamento de estoque desenvolvida em Go. O projeto foi projetado com foco em simplicidade, escalabilidade e documentação automática, utilizando uma arquitetura modular moderna inspirada no Clean Architecture e Domain-Driven Design (DDD).

---

## 🛠️ Tecnologias Utilizadas

- **Linguagem:** Go 1.25+
- **Framework Web:** [Fuego](https://github.com/go-fuego/fuego) (Validação integrada via struct tags e documentação OpenAPI/Swagger gerada automaticamente).
- **ORM:** [GORM](https://gorm.io/) (Abstração de banco de dados e migrações integradas).
- **Bancos de Dados Suportados:**
  - **PostgreSQL 15** (para produção e desenvolvimento em container).
  - **SQLite 3** (modo plug-and-play sem dependência de containers).
- **Gerenciador de Tarefas:** [Just](https://github.com/casey/just) (alternativa moderna ao Make).
- **Containerização:** Docker e Docker Compose.

---

## 📂 Estrutura do Projeto

O projeto é dividido em camadas bem delimitadas dentro do diretório `internal/`:

- **[internal/database](file:///home/master/Documentos/DEV/stock-manager-go/internal/database):** Inicialização da conexão com o banco de dados (PostgreSQL/SQLite) com políticas de retry automáticas.
- **[internal/domain](file:///home/master/Documentos/DEV/stock-manager-go/internal/domain):** Entidades de domínio (Core) contendo regras de negócio, invariants e tipos primitivos do negócio.
- **[internal/models](file:///home/master/Documentos/DEV/stock-manager-go/internal/models):** Estruturas de dados de persistência que mapeiam diretamente as tabelas do banco de dados (GORM).
- **[internal/dtos](file:///home/master/Documentos/DEV/stock-manager-go/internal/dtos):** Objetos de Transferência de Dados para validação de entrada (`validate:"required,..."`) e formatação de saídas.
- **[internal/handlers](file:///home/master/Documentos/DEV/stock-manager-go/internal/handlers):** Camada de apresentação e controladores REST utilizando o framework Fuego.
- **[internal/usecase](file:///home/master/Documentos/DEV/stock-manager-go/internal/usecase):** Lógica da aplicação e orquestração de casos de uso (ex.: `OrderService`).
- **[internal/port](file:///home/master/Documentos/DEV/stock-manager-go/internal/port):** Definição de contratos (interfaces) para Repositórios e Serviços de infraestrutura.
- **[internal/infra](file:///home/master/Documentos/DEV/stock-manager-go/internal/infra):** Implementação concreta dos adaptadores (como repositórios utilizando Postgres).

> [!NOTE]
> **Discrepância de Tipos de ID (int vs UUID):**
> O projeto possui uma discrepância intencional ou evolutiva: entidades legadas/simples (`Brand`, `Product`, `User`) utilizam IDs numéricos inteiros auto-incrementados (`int`) em suas tabelas de banco de dados (`internal/models`). Já a nova modelagem de Pedidos (`Order` em `internal/domain`) foi desenvolvida usando identificadores únicos universais (`uuid.UUID`), tanto para os IDs de pedido quanto para as referências de `UserID` e `ProductID`.

---

## 🚀 Como Executar o Projeto

O projeto utiliza um `justfile` com atalhos para todas as tarefas de desenvolvimento.

### Requisitos Mínimos
- [Go](https://go.dev/) instalado localmente.
- [Docker](https://www.docker.com/) e Docker Compose instalados.
- [Just](https://github.com/casey/just) instalado no sistema.

### Comandos Disponíveis (via `just`)

| Comando | Descrição |
| :--- | :--- |
| `just up` | Inicializa a stack completa (API + Banco de Dados PostgreSQL) via Docker Compose. |
| `just down` | Para e remove todos os containers e redes do Docker criados pelo projeto. |
| `just reset` | Destrói o banco de dados (removendo volumes) e reinicia a stack do zero. |
| `just up-sqlite` | Roda a API nativamente no Linux utilizando SQLite local (sem Docker). |
| `just build` | Compila o binário otimizado da API na pasta `bin/`. |
| `just ps` | Lista os containers ativos relacionados ao projeto. |
| `just logs` | Exibe e acompanha os logs em tempo real dos containers em execução. |
| `just tidy` | Executa o formatador de código (`go fmt`) e limpa as dependências do `go.mod`. |

A documentação interativa da API (Swagger UI) fica disponível em `http://localhost:8080/swagger/index.html` assim que o servidor for inicializado.

---

## 📋 Documentação dos Endpoints e CRUD

Todos os endpoints da API estão expostos sob o host padrão `http://localhost:8080`.

### 1. Marcas (`/brands`)
Gerenciamento de marcas associadas aos produtos.

*   **Listar Marcas**
    *   **Método:** `GET`
    *   **Rota:** `/brands/`
    *   **Resposta (200 OK):**
        ```json
        [
          {
            "id": 1,
            "name": "Nike",
            "country": "USA",
            "email": "contact@nike.com",
            "foundation_year": 1964
          }
        ]
        ```

*   **Criar Marca**
    *   **Método:** `POST`
    *   **Rota:** `/brands/`
    *   **Corpo da Requisição (Request Body):**
        ```json
        {
          "name": "Adidas",
          "country": "Germany",
          "email": "contact@adidas.com",
          "foundation_year": 1949
        }
        ```
    *   **Resposta (200 OK / 201 Created):**
        ```json
        {
          "id": 2,
          "name": "Adidas",
          "country": "Germany",
          "email": "contact@adidas.com",
          "foundation_year": 1949
        }
        ```
    *   **Erros Possíveis:** `400 Bad Request` (validação incorreta), `490 Conflict` (e-mail já cadastrado).

*   **Atualizar Marca**
    *   **Método:** `PUT`
    *   **Rota:** `/brands/{id}`
    *   **Corpo da Requisição (Campos opcionais):**
        ```json
        {
          "name": "Adidas Originals"
        }
        ```
    *   **Resposta (200 OK):**
        ```json
        {
          "id": 2,
          "name": "Adidas Originals",
          "country": "Germany",
          "email": "contact@adidas.com",
          "foundation_year": 1949
        }
        ```

*   **Deletar Marca**
    *   **Método:** `DELETE`
    *   **Rota:** `/brands/{id}`
    *   **Resposta (200 OK):**
        ```json
        {
          "message": "Deletado com sucesso"
        }
        ```

---

### 2. Produtos (`/products`)
Gerenciamento do catálogo de produtos em estoque.

*   **Listar Produtos**
    *   **Método:** `GET`
    *   **Rota:** `/products/`
    *   **Resposta (200 OK):**
        ```json
        [
          {
            "id": 1,
            "brand_id": 1,
            "name": "Air Max 90",
            "price": 699.9,
            "quantity": 50
          }
        ]
        ```

*   **Criar Produto**
    *   **Método:** `POST`
    *   **Rota:** `/products/`
    *   **Corpo da Requisição:**
        ```json
        {
          "name": "Ultraboost",
          "price": 899.9,
          "brand_id": 2,
          "quantity": 30
        }
        ```
    *   **Resposta (200 OK / 201 Created):**
        ```json
        {
          "id": 2,
          "brand_id": 2,
          "name": "Ultraboost",
          "price": 899.9,
          "quantity": 30
        }
        ```

*   **Atualizar Produto**
    *   **Método:** `PUT`
    *   **Rota:** `/products/{id}`
    *   **Corpo da Requisição (Campos opcionais):**
        ```json
        {
          "price": 949.9,
          "quantity": 25
        }
        ```
    *   **Resposta (200 OK):**
        ```json
        {
          "id": 2,
          "brand_id": 2,
          "name": "Ultraboost",
          "price": 949.9,
          "quantity": 25
        }
        ```

*   **Deletar Produto**
    *   **Método:** `DELETE`
    *   **Rota:** `/products/{id}`
    *   **Resposta (200 OK):**
        ```json
        {
          "message": "Deletado com sucesso"
        }
        ```

---

### 3. Usuários (`/users`)
Gerenciamento de clientes e usuários do sistema.

*   **Listar Usuários**
    *   **Método:** `GET`
    *   **Rota:** `/users/`
    *   **Resposta (200 OK):**
        ```json
        [
          {
            "id": 1,
            "name": "John Doe",
            "email": "user@example.com",
            "birthday": "2000-01-01",
            "address": "123 Main St, City, Country",
            "cpf": "123.456.789-00"
          }
        ]
        ```

*   **Criar Usuário**
    *   **Método:** `POST`
    *   **Rota:** `/users/`
    *   **Corpo da Requisição:**
        ```json
        {
          "name": "John Doe",
          "email": "user@example.com",
          "password": "supersecretpassword",
          "birthday": "2000-01-01",
          "address": "123 Main St, City, Country",
          "cpf": "123.456.789-00"
        }
        ```
    *   **Resposta (200 OK / 201 Created):**
        ```json
        {
          "id": 1,
          "name": "John Doe",
          "email": "user@example.com",
          "birthday": "2000-01-01",
          "address": "123 Main St, City, Country",
          "cpf": "123.456.789-00"
        }
        ```

*   **Atualizar Usuário**
    *   **Método:** `PUT`
    *   **Rota:** `/users/{id}`
    *   **Corpo da Requisição (Campos opcionais):**
        ```json
        {
          "name": "Johnathan Doe",
          "address": "456 New St, City, Country"
        }
        ```
    *   **Resposta (200 OK):**
        ```json
        {
          "id": 1,
          "name": "Johnathan Doe",
          "email": "user@example.com",
          "birthday": "2000-01-01",
          "address": "456 New St, City, Country",
          "cpf": "123.456.789-00"
        }
        ```

*   **Deletar Usuário**
    *   **Método:** `DELETE`
    *   **Rota:** `/users/{id}`
    *   **Resposta (200 OK):**
        ```json
        {
          "message": "Deletado com sucesso"
        }
        ```

---

### 4. Pedidos (`/orders`)
Gerenciamento do fluxo de compra (`buy`) ou venda (`sell`) de produtos.

*   **Listar Pedidos**
    *   **Método:** `GET`
    *   **Rota:** `/orders/`
    *   **Resposta (200 OK):**
        ```json
        [
          {
            "id": "22e70c53-ad97-4009-8803-b09bb74431e1",
            "user_id": "a90ee9be-cc0e-4361-bca4-d192c73eb64f",
            "product_id": "ee482811-137b-402a-995a-0d8591ef52fa",
            "quantity": 5,
            "price": 699.9,
            "type": "buy",
            "status": "pending",
            "created_at": "2026-06-22T19:54:25Z",
            "updated_at": "2026-06-22T19:54:25Z"
          }
        ]
        ```

*   **Buscar Pedido por ID**
    *   **Método:** `GET`
    *   **Rota:** `/orders/{id}`
    *   **Resposta (200 OK):**
        ```json
        {
          "id": "22e70c53-ad97-4009-8803-b09bb74431e1",
          "user_id": "a90ee9be-cc0e-4361-bca4-d192c73eb64f",
          "product_id": "ee482811-137b-402a-995a-0d8591ef52fa",
          "quantity": 5,
          "price": 699.9,
          "type": "buy",
          "status": "pending",
          "created_at": "2026-06-22T19:54:25Z",
          "updated_at": "2026-06-22T19:54:25Z"
        }
        ```

*   **Criar Pedido**
    *   **Método:** `POST`
    *   **Rota:** `/orders/`
    *   **Corpo da Requisição:**
        ```json
        {
          "user_id": "a90ee9be-cc0e-4361-bca4-d192c73eb64f",
          "product_id": "ee482811-137b-402a-995a-0d8591ef52fa",
          "quantity": 5,
          "price": 699.9,
          "type": "buy"
        }
        ```
    *   **Resposta (200 OK / 201 Created):**
        ```json
        {
          "id": "22e70c53-ad97-4009-8803-b09bb74431e1",
          "user_id": "a90ee9be-cc0e-4361-bca4-d192c73eb64f",
          "product_id": "ee482811-137b-402a-995a-0d8591ef52fa",
          "quantity": 5,
          "price": 699.9,
          "type": "buy",
          "status": "pending",
          "created_at": "2026-06-22T19:54:25Z",
          "updated_at": "2026-06-22T19:54:25Z"
        }
        ```
    *   **Nota:** O tipo (`type`) pode ser `"buy"` ou `"sell"`. O status inicial gerado é sempre `"pending"`.

*   **Executar Pedido**
    *   **Método:** `PUT`
    *   **Rota:** `/orders/{id}/execute`
    *   **Descrição:** Executa a transação pendente, efetivando a movimentação de estoque.
    *   **Resposta (200 OK):**
        ```json
        {
          "id": "22e70c53-ad97-4009-8803-b09bb74431e1",
          "user_id": "a90ee9be-cc0e-4361-bca4-d192c73eb64f",
          "product_id": "ee482811-137b-402a-995a-0d8591ef52fa",
          "quantity": 5,
          "price": 699.9,
          "type": "buy",
          "status": "executed",
          "created_at": "2026-06-22T19:54:25Z",
          "updated_at": "2026-06-22T19:56:00Z"
        }
        ```

*   **Cancelar Pedido**
    *   **Método:** `PUT`
    *   **Rota:** `/orders/{id}/cancel`
    *   **Descrição:** Cancela um pedido pendente.
    *   **Resposta (200 OK):**
        ```json
        {
          "id": "22e70c53-ad97-4009-8803-b09bb74431e1",
          "user_id": "a90ee9be-cc0e-4361-bca4-d192c73eb64f",
          "product_id": "ee482811-137b-402a-995a-0d8591ef52fa",
          "quantity": 5,
          "price": 699.9,
          "type": "buy",
          "status": "canceled",
          "created_at": "2026-06-22T19:54:25Z",
          "updated_at": "2026-06-22T19:56:10Z"
        }
        ```

---

## 🤖 Engenharia de Agentes (`.agents/`)

Este repositório está configurado para colaborar de forma avançada com agentes de Inteligência Artificial através de personas especializadas. Os perfis estão localizados em [.agents/](file:///home/master/Documentos/DEV/stock-manager-go/.agents):

- **[code-reviewer.md](file:///home/master/Documentos/DEV/stock-manager-go/.agents/code-reviewer.md):** Revisor sênior que avalia as alterações propostas em cinco dimensões: Corretude, Legibilidade, Arquitetura, Segurança e Performance.
- **[security-auditor.md](file:///home/master/Documentos/DEV/stock-manager-go/.agents/security-auditor.md):** Auditor de segurança focado em identificar injeções de SQL, vazamento de segredos, problemas de autenticação/autorização e aderência ao OWASP Top 10.
- **[test-engineer.md](file:///home/master/Documentos/DEV/stock-manager-go/.agents/test-engineer.md):** Engenheiro focado em testabilidade, cobertura de código, cenários de exceção e qualidade dos testes unitários/integração.
- **[web-performance-auditor.md](file:///home/master/Documentos/DEV/stock-manager-go/.agents/web-performance-auditor.md):** Especialista em auditoria de latência, consumo de recursos e otimização de rede.

Essas personas podem ser acionadas no chat de desenvolvimento usando comandos rápidos (como `/review` ou `/ship`) para garantir que nenhuma alteração seja mesclada com vulnerabilidades ou problemas estruturais.

---

## 🧠 Habilidades e Boas Práticas (`skills/`)

Para guiar o desenvolvimento colaborativo de forma rigorosa, o repositório traz diretrizes estruturadas no diretório [skills/](file:///home/master/Documentos/DEV/stock-manager-go/skills). Os agentes aplicam essas abordagens sistematicamente ao implementar novas funcionalidades:

1. **Design de APIs e Interfaces:** Criação de endpoints limpos e robustos seguindo as convenções REST e OpenAPI.
2. **Hardening de Segurança:** Validação de entradas nas bordas do sistema e proteção contra vulnerabilidades comuns.
3. **Desenvolvimento Guiado por Testes (TDD) e Especificações:** Escrita de testes e especificações antes da lógica de negócios.
4. **Recuperação de Erros e Resiliência:** Implementação de políticas de retry (como as aplicadas na inicialização da conexão do banco de dados) e tratamento seguro de pânico (panic-recovery).
5. **Doubt-Driven Development:** Resolução ativa de ambiguidades técnicas e de regras de negócio antes de codificar.
