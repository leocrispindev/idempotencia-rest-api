# Idempotencia + Golang API

Esta API exemplifica a prevenção de reprocessamento de requisições, assegurando a idempotência da aplicação e evitando a criação e processamento redundante de recursos. Ao garantir a correta utilização dos métodos HTTP, o projeto prioriza o tratamento do método não idempotente, ou seja, o POST.

## Explicação

O que é Idempotência?
Idempotência refere-se à propriedade de uma operação onde repetir a operação várias vezes produz o mesmo resultado que executá-la uma vez. Em termos simples, se uma operação é idempotente, fazer algo mais de uma vez não tem efeitos colaterais indesejados.

### Idempotência x Métodos HTTP

| Método HTTP | Idempotência | Descrição                                       |
|-------------|--------------|--------------------------------------------------|
| GET         | Sim          | Recupera informações, não modifica o estado.     |
| POST        | Não          | Cria um novo recurso, pode ter efeitos colaterais.|
| PUT         | Sim          | Atualiza um recurso ou cria se não existir.       |
| DELETE      | Sim          | Remove um recurso.                               |
| PATCH       | Não          | Modifica parcialmente um recurso.                 |

## Problema

A API foi projetada para gerenciar e armazenar mensagens trocadas entre usuários. Contudo, em determinadas circunstâncias, a mesma mensagem pode ser indevidamente associada a uma conversa mais de uma vez, comprometendo a integridade e a precisão dos registros.

## Solução 

A solução foi o envio de um header específico denominado "idempotencia-key", para garantir a singularidade, cada requisição o header recebe um valor diferente. Ao incluir este header em uma requisição, o sistema realiza um processo de validação dupla.

### Primeira validação - Cache:
Na primeira etapa, a API efetua uma consulta em um cache específico, verificando o status associado à chave fornecida no header. Esta consulta inicial permite uma resposta rápida e eficiente, determinando se a chave já foi previamente processada.

#### Lista de Status
* IN_PROCESS: a mensagem está em fase de processamento
* PROCESSED: a mensagem já foi processada
* ERROR_ON_PROCESS: ocorreu um erro durante o processamento


### Segunda validação - Banco de Dados:
Caso a chave não seja encontrada ou haja, ocorre a segunda etapa. 
É realizada uma consulta ao banco de dados para garantir que a chave não foi processada, caso não seja encontrada, a mensagem é salva no banco de dados, garantindo assim a integridade e singularidade dos dados.