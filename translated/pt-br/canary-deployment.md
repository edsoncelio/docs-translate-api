---
título: Canary Deployment
status: Concluído
categoria: conceito
tags: ["metodologia”, “aplicação”, “"]
---

## O que é

As implantações do Canary são uma estratégia de implantação que começa com dois ambientes: 
um com tráfego ao vivo e outro contendo o código atualizado sem tráfego ao vivo. 
O tráfego é transferido gradualmente da versão original do aplicativo para a versão atualizada. 
Ele pode começar movendo 1% do tráfego ao vivo, depois 10%, 25% e assim por diante, 
até que todo o tráfego esteja passando pela versão atualizada. 
As organizações podem testar a nova versão do software em produção, obter feedback, 
diagnostique erros e volte rapidamente para a versão estável, se necessário. 

O termo “canário” se refere à prática do “canário em uma mina de carvão” 
onde pássaros canários eram levados para minas de carvão para manter os mineiros seguros. 
Se gases nocivos inodoros estivessem presentes, o pássaro morreria e os mineiros sabiam que precisavam evacuar rapidamente. 
Da mesma forma, se algo der errado com o código atualizado, o tráfego ativo será “evacuado” de volta para a versão original. 

## Problema que ele resolve

Não importa o quão completa seja a estratégia de teste, sempre há alguns bugs descobertos na produção. 
Transferir 100% do tráfego de uma versão de um aplicativo para outra pode levar a falhas mais impactantes.

## Como isso ajuda

As implantações do Canary permitem que as organizações vejam como o novo software se comporta em cenários do mundo real 
antes de transferir tráfego significativo para a nova versão. 
Essa estratégia permite que as organizações minimizem o tempo de inatividade e recuem rapidamente em caso de problemas com a nova implantação. 
Ele também permite testes mais aprofundados de aplicativos de produção sem um impacto significativo na experiência geral do usuário.
