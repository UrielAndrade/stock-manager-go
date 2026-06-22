# Stock Manager Go

API de gerenciamento de estoque desenvolvida em Go. O projeto foi projetado com foco em simplicidade, escalabilidade e documentação automática, utilizando uma arquitetura modular moderna.

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
