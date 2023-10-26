---
título: Escalabilidade
status: Concluído
categoria: propriedade
tags: ["fundamental”, “propriedade”, “"]
---

A escalabilidade se refere ao quão bem um sistema pode crescer. 
Isso está aumentando a capacidade de fazer o que o sistema deveria fazer. 
Por exemplo, um [Kubernetes] (/kubernetes/) [cluster] (/cluster/) é dimensionado por 
aumentar ou reduzir o número de aplicativos [em contêineres] (/containerization/), 
mas essa escalabilidade depende de vários fatores. 
Quantos [nós] (/nodes/) ele tem, quantos [contêineres] (/container/) cada nó pode manipular, 
e quantos registros e operações o plano de controle pode suportar?

Um sistema escalável facilita a adição de mais capacidade. 
Nós diferenciamos entre duas abordagens de escalabilidade. 
Por um lado, há [escala horizontal] (/horizontal-scaling/) que adiciona mais nós para lidar com o aumento da carga. 
Em contraste, em [escala vertical] (/vertical-scaling/), os nós individuais se tornam mais poderosos para realizar mais transações 
(por exemplo, adicionando mais memória ou CPU a uma máquina individual). 
Um sistema escalável é capaz de mudar facilmente e atender às necessidades do usuário.
