---
título: Cluster
status: Concluído
categoria: Conceito
tags: ["infraestrutura”, “fundamental”, “"]
---

## O que é

Um cluster é um grupo de computadores ou aplicativos que trabalham juntos em prol de um objetivo comum. 
No contexto da computação nativa em nuvem, o termo é mais frequentemente aplicado a [Kubernetes] (/kubernetes/). 
Um cluster Kubernetes é um conjunto de serviços (ou cargas de trabalho) que são executados em seus próprios contêineres, geralmente em máquinas diferentes. 
A coleção de todos esses serviços [conteinerizados] (/containerization/), conectados em uma rede, representa um cluster.

## Problema que ele resolve 

O software executado em um único computador apresenta um único ponto de falha 
— se o computador falhar ou alguém acidentalmente desconectar o cabo de alimentação, 
então, algum sistema essencial para os negócios pode ser colocado off-line. 
É por isso que o software moderno geralmente é criado como [aplicativos distribuídos] (/distributed-apps/), agrupados como clusters. 

## Como isso ajuda

Aplicativos distribuídos em cluster são executados em várias máquinas, eliminando um único ponto de falha. 
Mas construir sistemas distribuídos é muito difícil. 
Na verdade, é uma disciplina de ciência da computação por si só. 
A necessidade de sistemas globais e anos de tentativa e erro levaram ao desenvolvimento de um novo tipo de pilha de tecnologia: 
[tecnologias nativas da nuvem] (/cloud-native-tech/). 
Essas novas tecnologias são os alicerces que facilitam a operação e a criação de sistemas distribuídos.
