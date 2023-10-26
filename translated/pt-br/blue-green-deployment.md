---
título: Implantação em azul e verde
status: Concluído
categoria: conceito
tags: ["metodologia”, “aplicação”, “"]
---

## O que é

A implantação azul-esverdeada é uma estratégia para atualizar sistemas de computador em execução com o mínimo de tempo de inatividade. 
A operadora mantém dois ambientes, apelidados de “azul” e “verde”. 
Um atende ao tráfego de produção (a versão que todos os usuários estão usando atualmente), enquanto o outro é atualizado. 
Uma vez concluído o teste no ambiente não ativo (verde), 
o tráfego de produção é comutado (geralmente por meio do uso de um balanceador de carga). 
Observe que a implantação azul-esverdeada geralmente significa alternar todos os ambientes, incluindo muitos [serviços] (/service/), de uma só vez. 
De forma confusa, às vezes o termo é usado com relação a serviços individuais dentro de um sistema. 
Para evitar essa ambigüidade, o termo “implantação sem tempo de inatividade” é preferido quando se refere a componentes individuais.

## Problema que ele resolve

As implantações azul-esverdeadas permitem um tempo mínimo de inatividade ao atualizar o software, que deve ser alterado em “sincronia” devido à falta de compatibilidade com versões anteriores. 
Por exemplo, a implantação azul-verde seria apropriada para uma loja on-line 
consistindo em um site e um banco de dados que precisam ser atualizados, 
mas a nova versão do banco de dados não funciona com a versão antiga do site e vice-versa. 
Nesse caso, ambos precisam ser alterados ao mesmo tempo. 
Se isso fosse feito no sistema de produção, os clientes notariam o tempo de inatividade.

## Como isso ajuda

A implantação azul-esverdeada é uma estratégia apropriada para software não nativo da nuvem que precisa ser atualizado com o mínimo de tempo de inatividade. 
No entanto, seu uso normalmente é um “cheiro” de que o software legado precisa ser reprojetado para que os componentes possam ser atualizados individualmente.
