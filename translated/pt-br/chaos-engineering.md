---
título: Chaos Engineering
status: Concluído
categoria: conceito
tags: ["metodologia”, “”, “"]
---

## O que é

Chaos Engineering ou CE é a disciplina de experimentar em um [sistema distribuído] (/distributed-systems/) na produção 
para criar confiança na capacidade do sistema de suportar condições turbulentas e inesperadas.

## Problema que ele aborda

As práticas [SRE] (/site-reliability-engineering/) e [DevOps] (/devops/) se concentram em 
técnicas para aumentar a resiliência do produto e [confiabilidade] (/confiabilidade/). 
A capacidade de um sistema de tolerar falhas e, ao mesmo tempo, garantir a qualidade adequada do serviço é 
normalmente é um requisito de desenvolvimento de software. 
Há vários aspectos envolvidos que podem levar à interrupção de um aplicativo, 
como infraestrutura, plataforma ou outras partes móveis de um aplicativo (baseado em [microsserviço] (/microserviços/)). 
A implantação de alta frequência de novos recursos no ambiente de produção pode 
resultam em uma alta probabilidade de tempo de inatividade e um incidente crítico 
— com consequências consideráveis para o negócio.

## Como isso ajuda

A engenharia do caos é uma técnica para atender aos requisitos de resiliência. 
Ele é usado para obter resiliência contra falhas de infraestrutura, plataforma e aplicativos. 
Os engenheiros do caos usam experimentos de caos para injetar proativamente falhas aleatórias 
para verificar se um aplicativo, infraestrutura ou plataforma pode se recuperar automaticamente e se a falha não pode afetar os clientes de forma perceptível. 
Experimentos de caos visam descobrir pontos cegos 
(por exemplo, técnicas de monitoramento ou escalonamento automático) e para melhorar a comunicação entre as equipes durante incidentes críticos. 
Essa abordagem ajuda a aumentar a resiliência e a confiança da equipe em sistemas complexos, especialmente na produção.
