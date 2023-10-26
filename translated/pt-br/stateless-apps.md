---
título: Stateless Apps
status: Feedback apreciado
categoria: tecnologia
tags: ["fundamental”, “aplicativo”, “"]
---

## O que é

Um aplicativo sem estado não salva nenhum dado da sessão (estado) do cliente no servidor em que o aplicativo reside. 
Cada sessão é realizada como se fosse a primeira vez e as respostas não dependem dos dados de uma sessão anterior e 
fornece funcionalidade para usar serviços de impressão, CDN (Content Delivery Network) ou servidores Web 
para processar todas as solicitações de curto prazo. 
Por exemplo, alguém está pesquisando uma pergunta no mecanismo de pesquisa e pressionou o botão Enter. 
Caso a operação de pesquisa seja interrompida ou fechada por algum motivo, 
você precisa iniciar um novo, pois não há dados salvos para sua solicitação anterior.

## Problema que ele resolve

Aplicativos sem estado resolvem o problema da resiliência, 
porque diferentes pods em um [cluster] (/cluster/) podem funcionar de forma independente, 
com várias solicitações chegando a eles ao mesmo tempo. 
Se houver algum problema, você pode facilmente reiniciar o aplicativo, 
e retornará ao seu estado inicial com pouco ou nenhum tempo de inatividade. 
Dessa forma, os benefícios dos aplicativos sem estado incluem resiliência, elasticidade e alta disponibilidade. 
No entanto, a maioria dos aplicativos que usamos hoje são pelo menos parcialmente [com estado] (/stateful-apps/), 
pois eles armazenam itens como preferências e configurações para melhorar a experiência do usuário.

## Como isso ajuda

Resumindo tudo, em um aplicativo sem estado, a única coisa pela qual seu cluster é responsável é 
o código e outros conteúdos estáticos que estão hospedados nele. 
Pronto, sem alterar bancos de dados, sem gravações e sem sobrar arquivos quando o pod é excluído. 
[contêineres] sem estado (/container/) são mais fáceis de implantar, 
e você não precisa se preocupar em salvar os dados do contêiner em volumes de armazenamento persistentes. 
Você também não precisa se preocupar em fazer backup dos dados.
