---
título: Containerização
status: Concluído
categoria: Tecnologia
tags: ["aplicativo”, “”, “"]
---

## O que é

A conteinerização é o processo de agrupar um aplicativo e suas dependências em uma imagem de contêiner. 
O processo de criação do contêiner exige a adesão ao padrão [Open Container Initiative] (https://opencontainers.org) (OCI). 
Desde que a saída seja uma imagem de contêiner que siga esse padrão, qual ferramenta de conteinerização é usada não importa.

## Problema que ele resolve 

Antes que os contêineres se tornassem predominantes, as organizações dependiam de máquinas virtuais (VMs) para 
orquestre vários aplicativos em uma única [máquina de metal puro] (/bare-metal-machine/). 
As VMs são significativamente maiores que os contêineres e precisam de um hipervisor para serem executadas. 
Devido ao armazenamento, backup e transferência desses modelos maiores de VM, a criação dos modelos de VM também é lenta. 
Além disso, as VMs podem sofrer alterações na configuração, o que viola o princípio de [imutabilidade] (/immutable-infrastructure/).

## Como isso ajuda

As imagens de contêiner são leves (diferentemente das VMs tradicionais) e 
o processo de conteinerização requer um arquivo com uma lista de dependências. 
Esse arquivo pode ser controlado por versão e o processo de construção automatizado, 
permitindo que uma organização se concentre em outras prioridades 
enquanto os processos automatizados cuidam da construção. 
Uma imagem de contêiner é armazenada por um identificador exclusivo 
que está vinculado ao seu conteúdo e configuração exatos. 
À medida que os contêineres são programados e reprogramados, 
eles são sempre redefinidos para seu estado inicial, o que elimina o desvio de configuração.
