---
título: Arquitetura orientada por eventos
status: Concluído
categoria: conceito
tags: ["arquitetura”, “”, “"]
---

## O que é

A arquitetura orientada a eventos é uma arquitetura de software que promove a criação, o processamento e o consumo de eventos.
Um evento é qualquer alteração no estado de um aplicativo.
Por exemplo, chamar uma carona em um aplicativo de compartilhamento de caronas representa um evento.
Essa arquitetura cria a estrutura na qual os eventos podem ser encaminhados adequadamente de sua origem (o aplicativo solicitando uma carona) para os receptores desejados (os aplicativos dos motoristas disponíveis nas proximidades).

## Problema que ele resolve

À medida que mais dados se tornam em tempo real, encontrar maneiras confiáveis de garantir que os eventos sejam capturados e encaminhados para o [serviço] (/service/) apropriado que deve processar solicitações de eventos fica cada vez mais desafiador.
Os métodos tradicionais de tratamento de eventos geralmente não têm como garantir que as mensagens sejam roteadas adequadamente ou que tenham sido realmente enviadas ou recebidas.
À medida que os aplicativos começam a crescer, fica mais difícil orquestrar eventos.

## Como isso ajuda

As arquiteturas orientadas por eventos estabelecem um hub central para todos os eventos (por exemplo, Kafka).
Em seguida, você define os produtores do evento (fonte) e os consumidores (destinatários), e o hub central de eventos garante o fluxo dos eventos.
Essa arquitetura garante que os serviços permaneçam desacoplados e que os eventos sejam encaminhados adequadamente do produtor para o consumidor.
O produtor pegará o evento recebido, geralmente pelo protocolo HTTP, e encaminhará as informações do evento.
