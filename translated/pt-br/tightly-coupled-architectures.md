---
título: Arquiteturas fortemente acopladas
status: Concluído
categoria: Propriedade
tags: ["fundamental”, “arquitetura”, “propriedade"]
---

A arquitetura fortemente acoplada é um estilo arquitetônico em que vários componentes do aplicativo são interdependentes 
(o paradigma oposto de [arquiteturas fracamente acopladas] (/loosely-coupled-architecture/)). 
Isso significa que uma alteração em um componente provavelmente afetará outros componentes. 
Geralmente, é mais fácil de implementar do que estilos arquitetônicos mais fracamente acoplados, 
mas pode deixar um sistema mais vulnerável a falhas em cascata. 
Eles também tendem a exigir lançamentos coordenados de componentes 
o que pode prejudicar a produtividade dos desenvolvedores.

Arquiteturas de aplicativos fortemente acopladas são uma forma bastante tradicional de criar aplicativos. 
Embora não seja necessariamente consistente com todas as melhores práticas de desenvolvimento de [microsserviços] (/microserviços/) 
eles podem ser úteis em circunstâncias específicas. 
Eles tendem a ser mais rápidos e simples de implementar e 
assim como [aplicativos monolíticos] (/monolithic-apps/), eles podem acelerar o ciclo de desenvolvimento inicial.
