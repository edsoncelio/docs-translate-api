---
título: Bare Metal Machine
status: Concluído
categoria: tecnologia
tags: ["infraestrutura”, “”, “"]
---

## O que é

Bare metal se refere a um computador físico, mais especificamente um servidor, que tem um e apenas um sistema operacional. 
A distinção é importante na computação moderna porque muitos, se não a maioria, dos servidores são [máquinas virtuais] (/virtual-machine/). 
Um servidor físico geralmente é um computador bastante grande com hardware poderoso embutido. 
Instalando um sistema operacional e executando aplicativos diretamente nesse hardware físico, 
sem [virtualização] (/virtualization/), é chamado de execução em “bare metal”.

## Problema que ele resolve

O emparelhamento de um sistema operacional com um computador físico é o padrão original de computação. 
Todos os recursos do computador físico estão disponíveis diretamente no sistema operacional e sem a presença de nenhuma camada de virtualização, 
não há atraso artificial na tradução das instruções do sistema operacional para o hardware.

## Como isso ajuda

Ao dedicar todos os recursos computacionais de um computador a um único sistema operacional, 
você potencialmente fornece o melhor desempenho possível para o sistema operacional. 
Se você precisar executar uma carga de trabalho que deve ter acesso extremamente rápido aos recursos de hardware, 
metal puro pode ser a solução certa. 

No contexto de [aplicativos nativos da nuvem] (/cloud-native-apps/), 
geralmente pensamos no desempenho em termos de [escalabilidade] (/escalabilidade/) para um grande número de eventos simultâneos, 
que pode ser tratado por [escala horizontal] (/horizontal-scaling/) (adicionando mais máquinas ao seu pool de recursos). 
Mas algumas cargas de trabalho podem exigir [dimensionamento vertical] (/vertical-scaling/) (adicionando mais potência a uma máquina física existente) 
e/ou uma resposta de hardware físico extremamente rápida, caso em que o metal puro é mais adequado. 
O metal puro também permite que você ajuste o hardware físico e possivelmente até mesmo os drivers de hardware para ajudar a realizar sua tarefa.
