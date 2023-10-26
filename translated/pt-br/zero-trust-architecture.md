---
título: Arquitetura Zero Trust
status: Concluído
categoria: Conceito
tags: ["segurança”, “”, “"]
---

## O que é

A arquitetura Zero Trust prescreve uma abordagem para o projeto e a implementação de sistemas de TI 
onde a confiança é completamente removida. 
O princípio fundamental é “nunca confie, sempre verifique”, dispositivos ou sistemas em si, 
ao se comunicar com outros componentes de um sistema, sempre verifique a si mesmos antes de fazer isso. 
Atualmente, em muitas redes, dentro da rede corporativa, os sistemas e dispositivos internos podem se comunicar livremente uns com os outros 
pois estão dentro do limite confiável do perímetro da rede corporativa. 
A arquitetura Zero Trust adota a abordagem oposta, onde, embora dentro do perímetro da rede, 
os componentes do sistema precisam primeiro passar pela verificação antes que qualquer comunicação seja feita.

## Problema que ele resolve

Com a abordagem tradicional baseada em confiança, em que sistemas e dispositivos que existem dentro de um perímetro de rede corporativa, 
a suposição é que, como há confiança, não há problema. 
No entanto, a arquitetura Zero Trust reconhece que a confiança é uma vulnerabilidade. 
No caso de um atacante ter acesso a um dispositivo confiável, 
dependendo do nível de confiança e acesso que foi concedido a esse dispositivo, 
o sistema agora está vulnerável a ataques 
pois o atacante está dentro do perímetro de rede “confiável” e é capaz de se mover lateralmente por todo o sistema. 
Em uma arquitetura de confiança zero, a confiança é removida, reduzindo, portanto, a superfície de ataque 
já que um atacante agora é forçado a verificar antes de prosseguir pelo sistema.

## Como isso ajuda

A adoção de uma arquitetura de confiança zero traz o principal benefício do aumento da segurança 
com uma redução na superfície de ataque. 
A remoção da confiança do seu sistema corporativo agora aumenta o número e a força das portas de segurança 
pelas quais um invasor precisa passar para obter acesso a outras áreas do sistema.
