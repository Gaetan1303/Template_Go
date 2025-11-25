# Template Go MVC

## Description

Template de projet Go professionnel utilisant une architecture MVC, le routeur chi, et respectant les principes SOLID. Idéal pour démarrer une API REST évolutive et maintenable.

## Démarrage rapide

### Prérequis
- Go >= 1.21
- PostgreSQL (optionnel selon vos besoins)

### Installation

```bash
# Cloner le repository
git clone <url-du-repo>
cd Template_Go

# Initialiser les dépendances
go mod tidy

# Configurer les variables d'environnement
cp config/.env.example config/.env

# Lancer l'application
go run ./cmd/main.go
```

## Structure du projet

```
Template_Go/
├── cmd/                # Point d'entrée (main.go)
├── config/             # Configuration (env, etc.)
├── internal/
│   ├── controllers/    # Contrôleurs (gestion des requêtes)
│   ├── factory/        # Patterns créationnels (ex: UserFactory)
│   ├── middlewares/    # Middlewares HTTP
│   ├── models/         # Modèles de données
│   ├── repositories/   # Accès aux données
│   ├── routes/         # Définition des routes
│   ├── services/       # Logique métier
│   └── views/          # Rendu des réponses (JSON, etc.)
├── pkg/                # Utilitaires réutilisables
├── tests/              # Tests unitaires et d'intégration
├── .env.example        # Exemple de configuration d'environnement
├── .gitignore
├── go.mod
├── go.sum
└── README.md
```

## Configuration

### Variables d'environnement (.env)

```env
# Server
PORT=8080
ENV=development

# Database (exemple)
DB_HOST=localhost
DB_PORT=5432
DB_NAME=template_go
DB_USER=postgres
DB_PASSWORD=password
```

## Tests

```bash
# Lancer tous les tests unitaires et d'intégration
go test ./...
```

## Scripts utiles

```bash
# Formatage du code
go fmt ./...
# Vérification des dépendances
go mod tidy
```

## Exemple d'utilisation (API REST)

### Créer un nouvel utilisateur

```http
POST /api/users
Content-Type: application/json

{
  "email": "user@example.com",
  "name": "John Doe"
}
```

### Récupérer un utilisateur

```http
GET /api/users/{id}
```

## Workflow de développement

1. Créer un modèle dans `/internal/models`
2. Créer une interface de repository dans `/internal/repositories`
3. Implémenter le repository
4. Créer une interface de service dans `/internal/services`
5. Implémenter le service
6. Créer un contrôleur dans `/internal/controllers`
7. Définir les routes dans `/internal/routes`
8. Écrire les tests dans `/tests`



## Licence

Ce projet est sous licence El Mimine License.