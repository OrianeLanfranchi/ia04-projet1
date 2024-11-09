IA04 - Oriane LANFRANCHI - Eileen LORENZO
#  Bureau de vote multi-agent

## Général
Ce projet implémente un agent serveur permettant de gérer des scrutins basés sur différentes méthodes de votes. Il permet de créer les scrutins, prendre en compte les votes des agents votants, et traiter les votes selon la méthode de vote choisie pour retourner les résultats (gagnant et classement des candidats).

Le projet implémente également une architecture pour les agents votants, et générer des requêtes.

## Récupérer et exécuter le projet
Récupérer les sources du serveur :

`go get github.com/OrianeLanfranchi/ia04-projet1/cmd/launchServerAgent`

Puis :

`go install github.com/OrianeLanfranchi/ia04-projet1/cmd/launchServerAgent`


Si cela ne fonctionne pas, il est possible de lancer le projet en récupérant ce code source.
* Pour lancer le serveur, il faut alors lancer la commande :
`go run launchServerAgent.go` (dossier cmd/launchServerAgent)
* Pour lancer un agent votant, il faut lancer la commande :
`go run launchVoteAgent.go` (dossier cmd/launchVoteAgent)

## API
Le format des requêtes et des réponses suit le document api de référence, dupliqué dans ce projet : `api.md`

Le seul ajout par rapport au document de référence est le possible retour d'un code HTTP Status Internal Server Error (503) lors de la récupération des résultats d'un scrutin, dans le cas où il n'y a pas de gagnant.

## Méthodes de vote
Les méthodes de vote implémentées (avec correspondance de la nomenclature de l'API) sont :
* Majorité : `majority`
* Borda : `borda`
* STV : `stv`
* Copeland : `copeland`
* Condorcet : `condorcet` (à noter que la règle Condorcet ne retourne pas de classement ; seulement un gagnant si celui-ci existe)
* Approval : `approval`

## Agents implémentés
Ce projet implémente un serveur de vote et des agents votants génériques, dont les paramètres de votes sont à modifier pour respecter les règles du scrutin pour lequel ils doivent voter.

## Améliorations possibles du projet
Il serait possible (et souhaitable) de rendre plus scalable le système de vérification des préférences des votants (en particulier).

L'idée actuelle serait de rajouter un attribut "Verification" pour chaque ballot stocké par le serveur, qui comporterait une fonction de vérification des préférences propre au ballot, et en adéquation avec la règle de vote. Il suffirait alors de l'exécuter pour vérifier que les préférences d'un agent votant sont correctement formatées. En effet, la règle de vote Approval ne nécessite pas les mêmes vérifications que les autres règles de vote (et si on souhaite ajouter de nouvelles règles de vote au serveur, elles-mêmes auront peut-être d'autres vérifications à réaliser).

Egalement, pour la récupération des résultats, il est possible d'obtenir un gagnant, mais pas de classement (ce qui est le cas de la règle Condorcet ici). Tel que le projet est actuellement constitué, il y a une condition dans la récupération des résultats pour déterminer s'il faut récupérer le résultat du classement. Ce n'est pas très esthétique. Cela pourrait être amélioré.