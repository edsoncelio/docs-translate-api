---
título: Entrega contínua (CD)
status: Concluído
categoria: conceito
tags: ["metodologia”, “aplicação”, “"]
---

## O que é

A entrega contínua, geralmente abreviada como CD, é um conjunto de práticas 
em que as alterações de código são implantadas automaticamente em um ambiente de aceitação 
(ou, no caso de implantação contínua, na produção). 
O CD inclui procedimentos cruciais para garantir que o software seja testado adequadamente 
antes da implantação e fornece uma maneira de reverter as alterações, se necessário. 
A integração contínua (CI) é o primeiro passo para a entrega contínua 
(ou seja, as mudanças precisam ser mescladas de forma limpa antes de serem testadas e implantadas).

## Problema que ele resolve

A implantação de atualizações [confiáveis] (/confiability/) se torna um problema em grande escala. 
Idealmente, implantaríamos com mais frequência para oferecer melhor valor aos usuários finais. 
No entanto, fazer isso manualmente se traduz em altos custos de transação para cada mudança. 
Historicamente, para evitar esses custos, as organizações lançaram com menos frequência, 
implantando mais mudanças de uma só vez e aumentando o risco de que algo dê errado.

## Como isso ajuda

As estratégias de CD criam um caminho totalmente automatizado para a produção 
que testa e implanta o software usando várias estratégias de implantação 
como as versões [canary] (/canary-deployment/) ou [blue-green] (/blue-green-deployment/). 
Isso permite que os desenvolvedores implantem códigos com frequência, dando a eles a tranquilidade de saber que a nova revisão foi testada. 
Normalmente, o desenvolvimento baseado em troncos é usado em estratégias de CD em vez de ramificação de recursos ou pull requests.

## Termos relacionados

* [Integração contínua] (/integração contínua/)
* [Implantação contínua] (/implantação contínua/)
