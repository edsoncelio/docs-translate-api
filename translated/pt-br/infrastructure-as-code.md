---
título: Infraestrutura como código (IaC)
status: Concluído
categoria: conceito
tags: ["infraestrutura”, “metodologia”, “"]
---

## O que é

Infraestrutura como código é a prática de armazenar a definição de infraestrutura como um ou mais arquivos. 
Isso substitui o modelo tradicional em que a infraestrutura como serviço é provisionada manualmente, 
geralmente por meio de scripts de shell ou outras ferramentas de configuração.

## Problema que ele resolve

A criação de aplicativos de forma nativa na nuvem exige que a infraestrutura seja descartável e reproduzível. 
Ele também precisa [escalar] (/escalabilidade/) sob demanda de forma automatizada e repetível, potencialmente sem intervenção humana. 
O provisionamento manual não pode atender aos requisitos de capacidade de resposta e escala de [aplicativos nativos da nuvem] (/cloud-native-apps/). 
As alterações manuais na infraestrutura não são reproduzíveis, atingem rapidamente os limites de escala e introduzem erros de configuração.

## Como isso ajuda

Ao representar os recursos do data center, como servidores, balanceadores de carga e sub-redes, como código, 
ele permite que as equipes de infraestrutura tenham uma única fonte confiável para todas as configurações e 
também permite que eles gerenciem seu data center em um pipeline de [CI] (/continuous-integration/)/[CD] (/continuous-delivery/), 
implementação de estratégias de controle de versão e implantação.
