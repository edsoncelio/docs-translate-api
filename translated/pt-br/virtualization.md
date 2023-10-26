---
título: Virtualização
status: concluído
categoria: tecnologia
tags: ["fundamental”, “infraestrutura”, “metodologia"]
---

## O que é

Virtualização, no contexto da computação nativa em nuvem, 
refere-se ao processo de pegar um computador físico, às vezes chamado de servidor, 
e permitindo que ele execute vários sistemas operacionais isolados. 
Esses sistemas operacionais isolados e seus recursos de computação dedicados (CPU, memória e rede) são 
chamadas de máquinas virtuais ou VMs. 
Quando falamos sobre uma [máquina virtual] (/virtual-machine/), estamos falando sobre um computador definido por software. 
Algo que parece e funciona como um computador real, mas está compartilhando hardware com outras máquinas virtuais.
[Computação em nuvem] (/cloud-computing/) é alimentada principalmente pela tecnologia de virtualização.
Como exemplo, você pode alugar um “computador” da AWS — esse computador é na verdade uma VM.

## Problema que ele resolve

A virtualização aborda vários problemas, incluindo a melhoria do uso do hardware físico 
permitindo que mais aplicativos sejam executados na mesma máquina física 
enquanto ainda estão isolados um do outro por motivos de segurança.

## Como isso ajuda

Os aplicativos executados em máquinas virtuais não sabem que estão compartilhando um computador físico. 
A virtualização também permite que os usuários do datacenter criem um novo “computador” (também conhecido como VM) em minutos 
sem se preocupar com as restrições físicas de adicionar um novo computador a um data center. 
