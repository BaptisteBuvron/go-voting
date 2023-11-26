# GO-Voting

[![CI: Test](https://github.com/BaptisteBuvron/go-voting/actions/workflows/test.yml/badge.svg)](https://github.com/BaptisteBuvron/go-voting/actions/workflows/test.yml)

Subject:

- [TP3](https://gitlab.utc.fr/lagruesy/ia04/-/blob/main/docs/sujets/td3/sujet.md)
- [Serveur de vote](https://gitlab.utc.fr/lagruesy/ia04/-/blob/main/docs/sujets/activit%C3%A9s/serveur-vote/api.md)

## Installation

Install [Go](https://golang.org/doc/install).

Clone the repository:

```bash
git clone https://github.com/BaptisteBuvron/go-voting
```

Run the server:

```bash
go run ia04/cmd/server
```

Example commands for client:

```bash
go run ia04/cmd/client v1 new_ballot majority '2023-11-26T16:27:11+00:00' 'v1,v2,v3' 5 '0,1,2,3,4'
go run ia04/cmd/client v1 vote majority-18c0c24a3245e '0,1,2,3,4'
go run ia04/cmd/client v1 result majority-18c0c24a3245e
```

Run tests:

```bash
go test '-coverprofile=coverage.txt' -v ./...
go tool cover '-html=coverage.txt'
```

## API

### Erreurs communes

Pendant l'utilisation de l'API certaines erreures peuvent se produire. Voici la liste des erreurs possibles :

| Code retour | Signification                               |
|-------------|---------------------------------------------|
| `500`       | Une erreur interne au serveur est survenue  |
| `400`       | La requête est mal formée ou JSON incomplet |
| `405`       | La ressource n'existe pas                   |

### Commande `/new_ballot`

- Requête : `POST`
- Objet `JSON` envoyé

| propriété   | type           | exemple de valeurs possibles                                |
|-------------|----------------|-------------------------------------------------------------|
| `rule`      | `string`       | `"majority"`,`"borda"`, `"approval"`, `"stv"`, `"copeland"` |
| `deadline`  | `string`       | `"2023-10-09T23:05:08+02:00"`  (format RFC 3339)            |
| `voter-ids` | `[string,...]` | `["ag_id1", "ag_id2", "ag_id3"]`                            |
| `#alts`     | `int`          | `12`                                                        |
| `tie-break` | `[int,...]`.   | `[4, 2, 3, 5, 9, 8, 7, 1, 6, 11, 12, 10]`                   |

*Remarques :* la deadline représente la date de fin de vote. Pour celle-ci, utiliser la bibliothèque standard de `Go`, en particulier le package `time`. La propriété `#alts` représente le nombre d'alternatives, numérotées de 1 à `#alts`.

- Code retour

| Code retour | Signification                                                                                                                                                                                                                        |
|-------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| `201`       | vote créé                                                                                                                                                                                                                            |
| `400`       | bad request (il manque une règle ou il n'y a pas au moins un voter d'enregistré ou il n'y a pas au moins deux alternatives, les alternatives données ne sont pas valides, le tie-break ne correspond pas aux nombres d'alternatives) |
| `501`       | not implemented (la règle de vote choisie n'est pas implémenter dans le système)                                                                                                                                                     |
| `500`       | L'ID aléatoire n'a pas pu être généré                                                                                                                                                                                                |

- Objet `JSON` renvoyé (si `201`)

| propriété   | type     | exemple de valeurs possibles |
|-------------|----------|------------------------------|
| `ballot-id` | `string` | `"majority-18c0c24a3245e"`   |

### Commande `/vote`

- Requête : `POST`
- Objet `JSON` envoyé

| propriété   | type        | exemple de valeurs possibles |
|-------------|-------------|------------------------------|
| `agent-id`  | `string`    | `"ag_id1"`                   |
| `ballot-id` | `string`    | `"majority-18c0c24a3245e"`   |
| `prefs`     | `[int,...]` | `[1, 2, 4, 3]`               |
| `options`   | `[int,...]` | `[3]`                        |

*Remarque :*`options` est facultatif et permet de passer des renseignements supplémentaires (par exemple le seuil d'acceptation en approval)

- code retour

| Code retour | Signification                                           |
|-------------|---------------------------------------------------------|
| `200`       | vote pris en compte                                     |
| `400`       | bad request (erreur dans les préférences de vote)       |
| `403`       | vote déjà effectué ou la personne n'a pas accès au vote |
| `404`       | le ballot de vote n'existe pas                          |
| `503`       | la deadline est dépassée                                |

Nous n'avons pas utilisé le code `501` Not Implemented.
Nous avons ajouté le code `404` Not Found quand le ballot n'existe pas.

### Commande `/result`

- Requête : `POST`
- Objet `JSON` envoyé

| propriété   | type     | exemple de valeurs possibles |
|-------------|----------|------------------------------|
| `ballot-id` | `string` | `"majority-18c0c24a3245e"`   |

- code retour

| Code retour | Signification |
|-------------|------------|
| `200`       | OK         |
| `425`       | Too early  |
| `404`       | Not Found  |

- Objet `JSON` renvoyé (si `200`)

| propriété | type        | exemple de valeurs possibles |
|-----------|-------------|------------------------------|
| `winner`  | `int`       | `4`                          |
| `ranking` | `[int,...]` | `[2, 1, 4, 3]`               |

*Remarque :* la propriété `ranking` est facultative.
