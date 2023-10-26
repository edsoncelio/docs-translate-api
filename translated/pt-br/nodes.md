---
título: Nodes
status: Concluído
categoria: Conceito
tags: ["infraestrutura”, “fundamental”, “"]
---

## O que é

Um nó é um computador que trabalha em conjunto com outros computadores, ou nós, para realizar uma tarefa comum. 
Pegue seu laptop, modem e impressora, por exemplo. 
Eles estão todos conectados em sua rede wifi, comunicando e colaborando, cada um representando um nó. 
Em [computação em nuvem] (/cloud-computing/), um nó pode ser um computador físico, 
um computador virtual, chamado de [VM] (/virtual-machine/) ou até mesmo um [contêiner] (/container/).

## Problema que ele resolve

Embora um aplicativo possa ser executado (e muitos o fazem) em uma única máquina, há alguns riscos envolvidos nisso. 
Ou seja, a falha do sistema subjacente interromperá o aplicativo. 
Para resolver isso, os desenvolvedores começaram a criar [aplicativos distribuídos] (/distributed-apps/) em que cada processo é executado em seu próprio nó. 
Assim, os nós executam aplicativos ou processos como parte de um grupo que forma um [cluster] (/cluster/), ou grupo, de nós que trabalham juntos para atingir um objetivo comum.

## Como isso ajuda

Um nó fornece uma unidade de computação distinta (memória, CPU, rede) que você pode atribuir a um cluster. 
Em uma plataforma ou aplicativo [nativo da nuvem] (/cloud-native-tech/), um nó representa uma única unidade que pode realizar o trabalho. 
Idealmente, os nós individuais são indiferenciados porque 
qualquer nó de um determinado tipo é indistinguível de qualquer outro nó do mesmo tipo.
