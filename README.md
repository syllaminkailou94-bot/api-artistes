# API Artistes

API REST en Go (Fiber v3 + GORM + SQLite) pour gérer des artistes.

## Lancer le projet

```bash
go mod tidy
go run main.go
```

Le serveur démarre sur http://localhost:3001.
Au premier lancement, la base `artists.db` est créée avec 3 artistes.

## Routes

| Méthode | Route | Rôle |
|---|---|---|
| GET | /health | état du serveur |
| GET | /artists | liste des artistes |
| GET | /artists/:id | un artiste |
| POST | /artists | créer un artiste |
| PUT | /artists/:id | remplacer un artiste |
| PATCH | /artists/:id | modifier certains champs |
| DELETE | /artists/:id | supprimer un artiste |

## Exemple

```bash
curl -X POST http://localhost:3001/artists \
  -H "Content-Type: application/json" \
  -d '{"name":"Angele","genre":"Pop"}'
```