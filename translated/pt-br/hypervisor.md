---
título: Hypervisor
status: Feedback apreciado
categoria: Tecnologia
tags: ["aplicativo”, “”, “"]
---

## O que é

Um hipervisor permite [virtualização] (/virtualization/)
aproveitando as vantagens dos recursos de [máquina de metal nu] (/bare-metal-machine/)
(CPU, memória, rede e armazenamento), dividindo-os em subpartes, 
e alocando recursos adequadamente para criar [máquinas virtuais (VM)] (/virtual-machine/)
até que o host subjacente atinja seus limites de desempenho.

## Problema que ele resolve

Tradicionalmente, um servidor só podia executar aplicativos de um único sistema operacional.
O processo de aquisição de software leva tempo. Requer infraestrutura com um ambiente específico
e uma equipe de engenheiros para gerenciá-los e monitorá-los.
Os servidores foram subutilizados, considerando o poder computacional de um servidor, ele pode executar vários sistemas operacionais e mais aplicativos.
A execução de aplicativos em metal puro não foi suficiente para atender às necessidades de tráfego flutuante.

## Como isso ajuda

No contexto da [computação em nuvem] (/cloud-computing/), o hipervisor se torna uma ferramenta eficaz.
Em contraste com o método tradicional de criar uma máquina virtual, um hipervisor torna o processo muito mais simples e rápido. 
Os recursos de hardware são particionados logicamente e atribuídos às VMs, mantendo-as isoladas como unidades distintas,
garantindo que funcionem de forma independente para que os problemas de um não afetem os outros,
e permitindo que as VMs instalem qualquer sistema operacional necessário.
Um hipervisor é uma abstração do hardware físico, ele cuida das complexidades de baixo nível de gerenciamento e monitoramento das VMs,
tornando as VMs vagamente vinculadas ao hardware, permitindo que as organizações migrem seus aplicativos para os servidores/nuvem remotos 
e autodimensionam seus serviços.
Com o tempo, o uso desse software [multilocatário] (/multilocatário/) reduziu os custos de computação.
