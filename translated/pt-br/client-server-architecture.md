---
título: Arquitetura cliente-servidor
status: Concluído
categoria: conceito
tags: ["arquitetura”, “fundamental”, “"]
---

## O que é

Em uma arquitetura cliente-servidor, a lógica (ou código) que compõe um aplicativo é dividida entre dois ou mais componentes: 
um cliente que pede que o trabalho seja feito 
(por exemplo, o aplicativo web do Gmail em execução no seu navegador), 
e um ou mais servidores que atendem a essa solicitação 
(por exemplo, o serviço de “enviar e-mail” em execução nos computadores do Google na nuvem). 
Neste exemplo, os e-mails enviados que você escreve são enviados pelo cliente (aplicativo da web em execução no seu navegador da web) 
para um servidor (computadores do Gmail, que encaminham seus e-mails enviados para seus destinatários).

Isso contrasta com aplicativos independentes (como aplicativos de desktop) que fazem todo o trabalho em um só lugar. 
Por exemplo, um programa de processamento de texto como o Microsoft Word pode ser instalado e executado inteiramente em seu computador.

## Problema que ele aborda 

Uma arquitetura cliente-servidor resolve um grande desafio que os aplicativos independentes representam: atualizações regulares. 
Em um aplicativo independente, para cada atualização, os usuários precisariam baixar e instalar a versão mais recente. 
Imagine ter que baixar todo o catálogo de produtos da Amazon em seu próprio computador antes de poder navegar nele!

## Como isso ajuda

Ao implementar a lógica do aplicativo em um servidor ou serviço remoto, 
os operadores podem atualizar isso sem precisar alterar a lógica no lado do cliente. 
Isso significa que as atualizações podem ser feitas com muito mais frequência. 
Armazenar dados no servidor permite que muitos clientes vejam e compartilhem os mesmos dados. 
Considere a diferença entre usar um processador de texto on-line em comparação com um processador de texto offline tradicional. 
No primeiro caso, seus arquivos existem no lado do servidor e 
podem ser compartilhados com outros usuários que simplesmente os baixam do servidor. 
No mundo antigo, os arquivos precisavam ser copiados para uma mídia removível (disquetes!) e compartilhado com indivíduos.
