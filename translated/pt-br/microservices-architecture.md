---
título: Arquitetura de microsserviços
status: Concluído
tags: ["arquitetura”, “fundamental”, “"]
---

## O que é

Uma arquitetura de microsserviços é uma abordagem arquitetônica que divide os aplicativos em (micro) [serviços] independentes (/service/) individuais, com cada serviço focado em uma funcionalidade específica.
Esses serviços funcionam em conjunto, aparecendo para o usuário final como uma única entidade. 
Veja a Netflix como exemplo. 
Sua interface permite acessar, pesquisar e visualizar vídeos. 
Esses recursos provavelmente são alimentados por serviços menores, cada um gerenciando uma funcionalidade, por exemplo, autenticação, pesquisa e execução de visualizações prévias em seu navegador.

Essa abordagem arquitetônica permite que os desenvolvedores desenvolvam novos recursos ou atualizem funcionalidades com muito mais rapidez do que se estivessem todos fortemente acoplados, como em um [aplicativo monolítico] (/monolithic-apps/) (mais sobre isso abaixo).

## Problema que ele resolve

Os aplicativos são compostos por partes diferentes, cada uma responsável por uma capacidade específica. 
A demanda por uma funcionalidade específica não necessariamente aumentará ou diminuirá com a demanda por outras partes do aplicativo. 
Voltando ao nosso exemplo da Netflix. 
Digamos que depois de uma grande campanha de marketing, a Netflix tenha um grande aumento nas inscrições, mas o streaming permaneceu mais ou menos estável nas primeiras horas do dia. 
O aumento nas inscrições exige mais capacidade de inscrição. 
Tradicionalmente (abordagem monolítica), todo o aplicativo precisaria ser [escalado] (/escalabilidade/) para acomodar o aumento — um uso muito ineficiente dos recursos. 

As arquiteturas monolíticas também facilitam que os desenvolvedores sucumbem às armadilhas do design. 
Como todo o código está em um só lugar, é mais fácil criar esse código [fortemente acoplado] (/tightly-coupled-architectures/) e mais difícil aplicar o princípio da separação de interesses. 
Os monólitos também costumam exigir que os desenvolvedores entendam toda a base de código antes de implantar qualquer chance. 
A arquitetura de microsserviços é uma resposta a esses desafios. 


## Como isso ajuda

Separar a funcionalidade em diferentes microsserviços facilita a implantação, a atualização e o dimensionamento de forma independente. 
Ele também permite que equipes diferentes trabalhem simultaneamente em uma pequena parte de um aplicativo maior sem impactar negativamente o resto do aplicativo. 
Embora uma arquitetura de microsserviços resolva muitos problemas, ela também cria sobrecarga operacional — as coisas que você precisa implantar e acompanhar aumentam em ordem de magnitude. 
Muitas [tecnologias nativas da nuvem] (/cloud-native-tech/) visam tornar os microsserviços mais fáceis de implantar e gerenciar.
