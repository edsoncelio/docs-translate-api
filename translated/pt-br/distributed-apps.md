---
título: Aplicativos distribuídos
status: Concluído
categoria: conceito
tags: ["arquitetura”, “”, “"]
---

## O que é

Um aplicativo distribuído é um aplicativo em que a funcionalidade é dividida em várias partes independentes menores. 
Os aplicativos distribuídos geralmente são compostos por [microsserviços] individuais (/microserviços/) 
que lidam com diferentes preocupações dentro de uma aplicação mais ampla. 
Em um ambiente nativo em nuvem, os componentes individuais normalmente são executados como [contêineres] (/container/) em um [cluster] (/cluster/). 

## Problema que ele resolve 

Um aplicativo em execução em um único computador representa um único ponto de falha — se esse computador falhar, o aplicativo ficará indisponível. 
Aplicativos distribuídos geralmente são comparados a [aplicativos monolíticos] (/monolithic-apps/). 
Um aplicativo monolítico pode ser mais difícil de escalar, pois os vários componentes não podem ser dimensionados de forma independente. 
Eles também podem prejudicar a velocidade do desenvolvedor à medida que crescem. 
porque mais desenvolvedores precisam trabalhar em uma base de código compartilhada que não necessariamente tenha limites bem definidos.

## Como isso ajuda

Ao dividir um aplicativo em partes diferentes e executá-lo em vários lugares, o sistema geral pode tolerar mais falhas. 
Também permite que um aplicativo aproveite os recursos de escalabilidade não disponíveis para uma única instância do aplicativo, 
ou seja, a capacidade de [escalar horizontalmente] (/horizontal-scaling/). 
No entanto, isso tem um custo: maior complexidade e sobrecarga operacional 
— agora você está executando vários componentes do aplicativo em vez de um aplicativo.
