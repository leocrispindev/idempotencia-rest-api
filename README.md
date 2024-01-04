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
