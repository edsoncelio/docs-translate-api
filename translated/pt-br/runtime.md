---
título: Runtime
status: Concluído
categoria: conceito
tags: ["aplicativo”, “”, “"]
---

## O que é

Um tempo de execução, em geral, executa um software.
É uma [abstração] (/abstração/) do sistema operacional subjacente que traduz os comandos do programa em ações respectivas para o sistema operacional. 

No contexto de [cloud native] (/cloud-native-apps/), _runtime_ geralmente se refere ao tempo de execução do contêiner. 
Um tempo de execução de contêiner implementa especificamente a especificação [Open Container Initiative] (https://opencontainers.org/) para garantir o tratamento consistente de diferentes tecnologias de orquestração de contêineres. 

## Problema que ele resolve

Sem a abstração do tempo de execução de um contêiner, o aplicativo teria que lidar com toda a mecânica de cada sistema operacional, aumentando a complexidade da execução do aplicativo. 

## Como isso ajuda
Os tempos de execução de contêineres são um componente necessário de orquestradores de contêineres, como o Kubernetes. 
Eles lidam com o ciclo de vida do contêiner, que faz principalmente três coisas.
Primeiro, ele define como as imagens do contêiner são especificadas e como o tempo de execução pode recuperá-las. 
Em segundo lugar, eles lidam com a forma como essas imagens são descompactadas, colocadas em camadas, montadas e executadas.
Além disso, os tempos de execução gerenciam os recursos de hardware cuidando de todas essas ações no nível do sistema operacional. 
Isso inclui alocação e isolamento de recursos. 
Com o tempo, diferentes produtos de tempo de execução de contêineres evoluíram, levando à Especificação OCI, 
que se tornou o padrão para tempos de execução de contêineres. 

A introdução desse padrão permite que os usuários finais combinem qualquer tempo de execução compatível com OCI com qualquer orquestrador de contêineres compatível com OCI (como o Kubernetes). 

## Termos relacionados

- [Nativo da nuvem] (https://glossary.cncf.io/cloud-native-apps/)
- [Containerização] (https://glossary.cncf.io/containerization/)
- [Orquestração de contêineres] (https://glossary.cncf.io/container-orchestration/)
- [Arquitetura de microsserviços] (https://glossary.cncf.io/microservices-architecture/)
