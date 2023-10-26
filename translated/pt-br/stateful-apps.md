---
título: Stateful Apps
status: Concluído
categoria: conceito
tags: ["fundamental”, “aplicativo”, “"]
---

## O que é

Quando falamos de aplicativos com estado (e sem estado), 
estado se refere a todos os dados que o aplicativo precisa armazenar para funcionar conforme projetado. 
Qualquer tipo de loja on-line que se lembre do seu carrinho é um aplicativo com estado, por exemplo. 

## Problema que ele resolve

O uso de um aplicativo geralmente exige várias solicitações. 
Por exemplo, ao fazer transações bancárias on-line, você se autenticará por 
digitando sua senha (solicitação #1), 
então você pode transferir dinheiro para um amigo (solicite #2), 
e, finalmente, você desejará ver os detalhes da transferência (solicitação #3). 
Para funcionar corretamente, cada etapa deve lembrar as anteriores, 
e o banco precisa se lembrar do estado das contas de todos. 
Hoje, a maioria dos aplicativos que usamos tem pelo menos parcialmente estado, 
pois eles armazenam itens como preferências e configurações para melhorar a experiência do usuário.

## Como isso ajuda

Há várias maneiras de armazenar o estado de um aplicativo com estado. 
O mais simples é manter o estado na memória e não persisti-lo em nenhum lugar. 
O problema com isso é que, sempre que o aplicativo precisa ser reiniciado, todo o estado é perdido. 
Para evitar isso, o estado persiste localmente (no disco) ou em sistemas de banco de dados. 
