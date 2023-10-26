---
título: Escala horizontal
status: Concluído
categoria: Conceito
tags: ["infraestrutura”, “”, “"]
---

## O que é

A escala horizontal é uma técnica em que a capacidade de um sistema é aumentada com a adição de mais [nós] (/nodes/) 
em vez de adicionar mais recursos de computação a nós individuais (o último sendo conhecido como [escala vertical] (/vertical-scaling/)). 
Digamos que temos um sistema de 4 GB de RAM e queremos aumentar sua capacidade para 16 GB de RAM, 
escalá-lo horizontalmente significa fazer isso adicionando 4 x 4 GB de RAM em vez de mudar para um sistema de 16 GB de RAM.

Essa abordagem aprimora o desempenho de um aplicativo adicionando novas instâncias, ou [nós] (/nodes/), 
para melhor distribuir a carga de trabalho. 
Em palavras simples, ele visa diminuir a carga do servidor 
em vez de expandir a capacidade do servidor individual.

## Problema que ele resolve

À medida que a demanda por um aplicativo cresce além da capacidade atual dessa instância de aplicativo, 
precisamos encontrar uma maneira de [escalar] (/escalabilidade/) (adicionar capacidade ao) sistema. 
Podemos adicionar mais nós ao sistema (escala horizontal) 
ou mais recursos computacionais para nós existentes (escala vertical).

## Como isso ajuda

O escalonamento horizontal permite que os aplicativos sejam dimensionados de acordo com os limites fornecidos pelo cluster subjacente. 
Ao adicionar mais instâncias ao sistema, o aplicativo pode processar um número maior de solicitações. 
Se um único nó puder lidar com 1.000 solicitações por segundo, 
cada nó adicional deve aumentar o número total de solicitações em cerca de 1.000 solicitações por segundo. 
Isso permite que o aplicativo trabalhe mais simultaneamente. 
sem precisar aumentar a capacidade de qualquer nó em particular.

## Termos relacionados

* [Escala vertical] (/escala vertical/)
* [Escala automática] (/escalonamento automático/)
