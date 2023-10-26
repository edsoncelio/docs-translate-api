---
título: Sistema distribuído
status: Concluído
categoria: conceito
tags: ["arquitetura”, “”, “"]
---

## O que é

Um sistema distribuído é uma coleção de elementos de computação autônomos 
conectado por meio de uma rede que aparece para os usuários como um único sistema coerente. 
Geralmente chamados de [nós] (/nodes/), esses componentes podem ser dispositivos de hardware (por exemplo, computadores, telefones celulares) ou processos de software. 
Os nós são programados para atingir um objetivo comum e, para colaborar, trocam mensagens pela rede. 

## Problema que ele resolve

Atualmente, vários aplicativos modernos são tão grandes que precisariam de supercomputadores para operar. 
Pense no Gmail ou no Netflix. Nenhum computador é poderoso o suficiente para hospedar o aplicativo inteiro. 
Ao conectar vários computadores, o poder computacional se torna quase ilimitado. 
Sem a computação distribuída, muitos aplicativos nos quais confiamos atualmente não seriam possíveis. 

Tradicionalmente, os sistemas [escalavam] (/escalabilidade/) verticalmente. 
É quando você adiciona mais CPU ou memória a uma máquina individual. 
O dimensionamento vertical é demorado, exige tempo de inatividade e atinge seu limite rapidamente. 

## Como isso ajuda

Os sistemas distribuídos permitem [escala horizontal] (/horizontal-scaling/) (por exemplo, adicionar mais nós ao sistema sempre que necessário). 
Isso pode ser automatizado, permitindo que um sistema gerencie um aumento repentino na carga de trabalho ou no consumo de recursos. 

Um sistema não distribuído se expõe a riscos de falha porque, se uma máquina falhar, todo o sistema falhará. 
Um sistema distribuído pode ser projetado de forma que, 
mesmo que algumas máquinas caiam, o sistema geral ainda pode continuar trabalhando para produzir o mesmo resultado.
