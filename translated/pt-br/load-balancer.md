---
título: Load Balancer
status: Feedback apreciado
categoria: conceito
tags: ["infraestrutura”, “rede”, “"]
---

## O que é

Um balanceador de carga é uma ferramenta que distribui com eficiência as solicitações recebidas entre várias instâncias de um aplicativo. 
Veja uma arquitetura de [microsserviço] (/microservices/), por exemplo, em que cada serviço pode ser [dimensionado horizontalmente] (/horizontal-scaling/). 
Um balanceador de carga fica na frente de um microsserviço escalável e garante que nenhuma instância receba a maior parte das solicitações.
Os balanceadores de carga podem ser baseados em software ou hardware.

## Problema que ele resolve

Aplicativos e sites modernos geralmente atendem a centenas de milhares de solicitações simultâneas do usuário final. 
Para lidar com todas essas solicitações, os aplicativos geralmente são dimensionados horizontalmente.
Mas a escala horizontal apresenta um novo desafio. Como você distribui o tráfego de entrada para todos os serviços de forma igual? 
É aqui que entram os balanceadores de carga.

## Como isso ajuda

Os balanceadores de carga distribuem dinamicamente todas as solicitações recebidas entre vários serviços, garantindo que nenhum serviço receba a maior parte, enquanto outros recebem apenas alguns ou nenhum. 
Resumindo, ele distribui a carga por vários serviços, seguindo um esquema definido (ou seja, uniforme ou baseado em porcentagem). 
Os balanceadores de carga são essenciais para o desempenho geral de um aplicativo e, em última análise, para a experiência do usuário.
