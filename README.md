# soat23-gp14-serverless

## O que a Plataforma é capaz de fazer?
- Criar usuários no Cognito;
- Autenticar usuários no Cognito;

## Pré-requisitos e como rodar a aplicação
- AWS SAM CLI
- Go 1.19

## Estrutura do Projeto
.
├── src
| ├── domain
| ├── handlers
| ├── infra
|   ├── adapters
|   └── settings
| ├── models
| └── services
├── template.yaml
└── test
| └── env-example.json

## Instalação
Faça o download do repositório através do arquivo zip ou do terminal usando o git clone;

```bash
git clone https://github.com/fiap-soat2023-gp14/soat23-gp14-serverless
```

Ao finalizar o download, acesse o diretório do projeto pelo seu terminal;

```bash
cd soat23-gp14-serverless
```

## Build
```bash
sam build
```

## Local Testing
```bash
sam local start-api --env-vars test/env.json
```

## Running Tests
```bash
make tests
```