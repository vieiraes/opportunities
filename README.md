# 🚀 Opportunities API

> 📚 Repositório de estudos - API em Go para aprendizado

## 📝 Sobre o Projeto

Este é um projeto de **estudos e aprendizado** da linguagem Go (Golang). O objetivo é desenvolver uma API REST para gestão de oportunidades, aplicando conceitos fundamentais da linguagem.

## ⚠️ Status

🔨 **Em Desenvolvimento** - Projeto em construção para fins educacionais

## 🛠️ Tecnologias

- **Go** - Linguagem de programação
- **PostgreSQL** - Banco de dados
- **godotenv** - Gerenciamento de variáveis de ambiente
- **lib/pq** - Driver PostgreSQL para Go

## 📦 Estrutura do Projeto

```
opportunities/
├── config/
│   └── db.go          # Configuração do banco de dados
├── main.go            # Ponto de entrada da aplicação
├── .env               # Variáveis de ambiente (não versionado)
├── .gitignore
├── go.mod
└── README.md
```

## 🚀 Como Executar

### Pré-requisitos

- Go 1.18 ou superior
- PostgreSQL instalado e rodando
- Git

### Instalação

1. Clone o repositório:
```bash
git clone https://github.com/vieiraes/opportunities.git
cd opportunities
```

2. Instale as dependências:
```bash
go mod download
```

3. Configure as variáveis de ambiente:
```bash
cp .env.example .env
```

Edite o arquivo `.env` com suas configurações:
```env
DB_HOST=localhost
DB_PORT=5432
DB_USERNAME=seu_usuario
DB_PASSWORD=sua_senha
DB_DATABASE_NAME=opportunities
PORT=:3344
```

4. Execute a aplicação:
```bash
go run main.go
```

O servidor estará rodando em `http://localhost:3344`

## 📚 Aprendizados

Durante o desenvolvimento deste projeto, estou estudando:

- ✅ Sintaxe e estrutura básica do Go
- ✅ Gerenciamento de pacotes com Go Modules
- ✅ Conexão com banco de dados PostgreSQL
- ✅ Tratamento de erros idiomático em Go
- ✅ Variáveis de ambiente
- ✅ HTTP Server básico
- 🔨 Rotas e handlers
- 🔨 CRUD operations
- 🔨 Middleware
- 🔨 Validações

## 🎯 Próximos Passos

- [ ] Implementar rotas CRUD
- [ ] Adicionar validações
- [ ] Criar migrations do banco
- [ ] Implementar testes
- [ ] Adicionar documentação da API (Swagger)
- [ ] Implementar autenticação JWT

## 👨‍💻 Autor

**Bruno Vieira**
- GitHub: [@vieiraes](https://github.com/vieiraes)

## 📄 Licença

Este projeto é para fins educacionais e está disponível sob a licença MIT.

---