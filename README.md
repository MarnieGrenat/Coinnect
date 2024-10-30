# Coinnect - Trabalho de Sistema Distribuído Cliente/Servidor - RPC/RMI

**Contribuidores**: Bruno Carlan, Gabriela Dellamora, Leonardo Ripes, Luize Iensse 

**Versão Atual**: 1.0

**Data de Atualização**: 30 de outubro de 2024

## Descrição

Este projeto faz parte do trabalho da disciplina de **Fundamentos de Processamento Paralelo e Distribuído**. Ele implementa um sistema distribuído baseado no modelo Cliente/Servidor, com foco em simular operações bancárias, como abertura de contas, transações de depósito, retirada e consulta de saldo, utilizando chamadas de procedimento remoto (RPC) ou chamada de método remoto (RMI).

### Funcionalidades

1. **Processo Administração (servidor):**
   - Realiza abertura e fechamento de contas.
   - Executa operações de saque, depósito e consulta de saldo.
   - Implementa controle de concorrência e garante a execução única (exactly-once) para operações não-idempotentes.

2. **Processo Agência (cliente):**
   - Solicita abertura e fechamento de contas.
   - Pode solicitar depósito, retirada e consulta de saldo em contas existentes.
   - Opera com semântica de execução exactly-once para garantir a consistência.

3. **Processo Caixa Automático (cliente):**
   - Realiza operações de depósito, retirada e consulta de saldo.
   - Garante a execução única mesmo em casos de falhas simuladas para testar a resiliência.

### Estrutura do Repositório

- `src/` - Contém o código fonte dos processos (Administração, Agência e Caixa Automático).
- `docs/` - Inclui o relatório técnico e documentação do projeto.
- `tests/` - Scripts e casos de teste para verificar o controle de concorrência e a semântica exactly-once.
- `README.md` - Informações gerais e instruções de uso.
- `LICENSE` - Informações sobre a licença do projeto.

## Relatório Técnico

O relatório técnico (`docs/relatorio_tecnico.pdf`) documenta:
- O mapeamento do problema para o modelo Cliente/Servidor com RPC/RMI.
- A implementação e testes de controle de concorrência.
- A implementação e testes de controle de idempotência com simulação de falhas.

## Instalação e Uso

1. **Compilação**: Execute o script de build localizado em `src/compile.sh` para compilar o projeto.
2. **Execução**:
   - Inicie o servidor (Processo Administração) usando `./admin_server`.
   - Em seguida, execute as instâncias dos clientes (Processo Agência e Processo Caixa Automático) conforme necessário.

## Licença

Este projeto está licenciado sob a [Licença MIT](LICENSE).
