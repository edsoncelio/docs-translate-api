---
título: Integração contínua (CI)
status: Concluído 
categoria: conceito
tags: ["aplicativo”, “metodologia”, “"]
---

## O que é 

A integração contínua, geralmente abreviada como CI, é a prática de integrar alterações de código com a maior regularidade possível. 
A CI é um pré-requisito para [entrega contínua] (/entrega contínua/) (CD). 
Tradicionalmente, o processo de CI começa quando as alterações de código são comprometidas com um sistema de controle de origem (Git, Mercurial ou Subversion) 
e termina com um artefato testado pronto para ser consumido por um sistema de CD. 

## Problema que ele aborda

Os sistemas de software geralmente são grandes e complexos, com vários desenvolvedores os mantendo e atualizando. 
Trabalhando em paralelo em diferentes partes do sistema, 
esses desenvolvedores podem fazer mudanças conflitantes e, inadvertidamente, interromper o trabalho uns dos outros. 
Além disso, com vários desenvolvedores trabalhando no mesmo projeto, 
qualquer tarefa diária, como testar e calcular a qualidade do código, precisaria ser repetida por cada desenvolvedor, desperdiçando tempo.

## Como isso ajuda

O software de CI verifica automaticamente se as alterações de código são mescladas de forma limpa sempre que um desenvolvedor comete uma alteração. 
É uma prática quase onipresente usar o servidor de CI para executar verificações, testes e até implantações de qualidade de código. 
Como tal, torna-se uma implementação concreta do controle de qualidade dentro das equipes. 
A CI permite que as equipes de software transformem cada confirmação de código em uma falha concreta ou em um candidato a lançamento viável.

## Termos relacionados

* [Entrega contínua] (/entrega contínua/)
* [Implantação contínua] (/implantação contínua/)
