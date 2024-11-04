# Coinnect - Trabalho de Sistema Distribuído Cliente/Servidor - RPC/RMI

**Contribuidores**: Bruno Carlan, Gabriela Dellamora, Leonardo Ripes, Luize Iensse

**Versão Atual**: 2.0

**Data de Atualização**: 04 de novembro de 2024

## Descrição

O projeto **Coinnect-FPPD** é um sistema distribuído baseado no modelo Cliente/Servidor utilizando RPC (Remote Procedure Call). Ele simula a criação de contas de clientes e operações bancárias, como abertura de contas, fechamento de contas, saques, depósitos e consultas de saldo. O sistema é composto por três tipos de processos:

- **Servidor de Administração**: Responsável pela criação e gerenciamento de contas, incluindo controle de transações e idempotência.
- **Cliente - Agência**: Solicita abertura e fechamento de contas e realiza operações bancárias.
- **Cliente - Caixa Automático (ATM)**: Realiza operações bancárias, como saques, depósitos e consultas de saldo.

## Estrutura do Projeto
```plaintext
Coinnect-FPPD/
├── src/
│   ├── Client/
│   │   ├── ATM/ATM.go
│   │   ├── BankBranch/BankBranch.go
│   │   ├── Client.go
│   │   └── Menu/Menu.go
│   ├── Server/
│   │   ├── Bank/
│   │   │   ├── Bank.go
│   │   │   └── Bank_test.go
│   │   └── Server.go
│   └── deps/
│       └── Pygmalion.go
├── tests/
├── settings.yml
├── Makefile
└── README.md
```
### Diagrama de Classes - Client
![image](docs/assets/ClassDiagram%20Client.svg)

### Diagrama de Classes - Server
![image](docs/assets/ClassDiagram%20Server.svg)

### Diagrama de Sequência - Exemplo 1
![image](docs/assets/SequenceDiagram%201.svg)

### Diagrama de Sequência - Exemplo 2
![image](docs/assets/SequenceDiagram%202.svg)
## Funcionalidades
- **Abertura de conta:** Permite que novos usuários criem contas.
- **Fechamento de conta:** Fecha contas existentes após autenticação.
- **Depósito:** Adiciona fundos a uma conta.
- **Saque:** Remove fundos de uma conta, com verificação de saldo.
- **Consulta de saldo:** Retorna o saldo atual de uma conta.
- **Controle de concorrência:** Implementado com sync.RWMutex para operações thread-safe.
- **Idempotência e execução exactly-once:** Implementado com RequestID para evitar duplicação de operações.

## Como Rodar o Projeto
### Requisitos
- *Go* (versão 1.20 ou superior)
- *Make* (para automação de tarefas)

### Instruções de Execução
#### 1. Compilar o projeto
```bash
make a
```

#### 2. Executar os testes
```bash
make t
```

#### 3. Limpar os binários gerados
```bash
make c
```

## Configurações
As configurações do servidor, como endereço e porta, são armazenadas no arquivo `settings.yml`. Certifique-se de configurá-lo de acordo com o seu ambiente antes de iniciar o servidor ou cliente.