---
título: Escala vertical
status: Concluído
categoria: Conceito
tags: ["infraestrutura”, “”, “"]
---

## O que é

A escala vertical, também conhecida como “escalar para cima e para baixo”, é uma técnica em que 
a capacidade de um sistema aumenta com a adição de CPU e memória a [nós] individuais (/nodes/) à medida que a carga de trabalho aumenta. 
Digamos que você tenha um computador de 4 GB de RAM e queira aumentar sua capacidade para 16 GB de RAM, 
escalá-lo verticalmente significa mudar para um sistema de 16 GB de RAM. 
(Consulte [escala horizontal] (/horizontal-scaling/) para uma abordagem de escala diferente.)

## Problema que ele resolve

À medida que a demanda por um aplicativo cresce além da capacidade atual dessa instância de aplicativo, 
precisamos encontrar uma maneira de escalar (adicionar capacidade) ao sistema. 
Podemos adicionar mais recursos computacionais aos nós existentes (escala vertical) 
ou mais nós no sistema ([escala horizontal] (/horizontal-scaling/)). 
[Escalabilidade] (/escalabilidade/) contribui para competitividade, eficiência, reputação e qualidade.

## Como isso ajuda

O escalonamento vertical permite redimensionar seu servidor sem alterar o código do aplicativo. 
Isso contrasta com a escala horizontal, em que o aplicativo deve ser replicável em escala, potencialmente exigindo atualizações de código. 
O dimensionamento vertical aumenta a capacidade de um aplicativo existente em 
adicionando recursos computacionais, permitindo que o aplicativo processe mais solicitações e trabalhe mais simultaneamente.

## Termos relacionados

* [Escala horizontal] (/escala horizontal/)
* [Escala automática] (/escalonamento automático/)
