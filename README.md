# traily - backend

## Architecture

Architecture hexagonale minimaliste : le domaine ne dépend d'aucun framework, les
détails techniques (HTTP, config, ...) vivent en périphérie et dépendent du domaine,
jamais l'inverse.

```
cmd/                        # point d'entrée, wiring uniquement
internal/
  domain/                   # logique métier pure (entités + services), zéro dépendance externe
    greeting/
  adapters/                 # adapters primaires (driving) : déclenchent le domaine
    http/                   # routes Gin, handlers, DTOs de réponse
  infrastructure/           # adapters secondaires (driven) : implémentent les ports du domaine
    (db, cache, API tierces, ...)
  config/                   # chargement de la configuration (variables d'env, defaults)
```

- `domain/` définit les entités et la logique métier. Il ignore Gin, JSON, HTTP, et
  toute dépendance technique.
- `adapters/http/` (primaire, "driving") traduit les requêtes HTTP en appels au
  domaine, et le résultat du domaine en réponse JSON (DTOs séparés des entités).
- `infrastructure/` (secondaire, "driven") contiendra les implémentations concrètes
  des interfaces définies par le domaine pour parler à l'extérieur : base de
  données, cache, API tierces, etc. Le domaine dépend d'une interface, jamais de
  ce package directement — c'est `infrastructure/` qui dépend du domaine. Vide
  pour l'instant, à peupler dès qu'une vraie source de données arrive.
- `cmd/main.go` assemble tout : config, implémentations d'infrastructure, service
  de domaine, handler, serveur.

Pour ajouter une nouvelle ressource : une entité + un service dans `domain/`, un
handler + ses routes dans `adapters/http/`, rien d'autre à toucher.
