# 🧙 Harry Potter Wiki API

Application web en Go permettant d'explorer l'univers de Harry Potter avec une interface de recherche de personnages.

## 📋 Fonctionnalités

- ✅ Recherche de personnages par nom
- ✅ Filtrage par maison (Gryffondor, Serpentard, Serdaigle, Poufsouffle)
- ✅ Filtrage par espèce (Humain, Demi-Géant, Loup-garou, etc.)
- ✅ Pagination des résultats
- ✅ Page de détails pour chaque personnage
- ✅ Gestion des erreurs (400, 404, 500)
- ✅ Design responsive
- ✅ Tests unitaires

## 🚀 Technologies utilisées

- **Go 1.25.0** - Langage de programmation
- **net/http** - Serveur HTTP natif
- **html/template** - Moteur de templates
- **Potter DB API** - Source des données

## 📁 Structure du projet

```
API_HP/
├── assets/              # Fichiers statiques
│   ├── css/
│   │   └── style.css    # Styles CSS
│   └── img/             # Images
├── src/
│   ├── cmd/
│   │   └── main.go      # Point d'entrée de l'application
│   ├── controllers/     # Contrôleurs HTTP
│   │   ├── characters.controller.go
│   │   ├── errors.controller.go
│   │   ├── suggestions.controller.go
│   │   └── *_test.go    # Tests unitaires
│   ├── helpers/         # Fonctions utilitaires
│   │   ├── errors.helper.go
│   │   └── *_test.go
│   ├── middlewares/     # Middlewares HTTP
│   │   └── errors.middleware.go
│   ├── models/          # Structures de données
│   │   ├── characters.model.go
│   │   └── errors.model.go
│   ├── routers/         # Configuration des routes
│   │   ├── main.router.go
│   │   ├── characters.router.go
│   │   └── errors.router.go
│   ├── services/        # Logique métier et appels API
│   │   ├── characters.service.go
│   │   └── *_test.go
│   └── templates/       # Gestion des templates
│       └── templates.go
├── templates/           # Templates HTML
│   ├── home.html        # Page d'accueil
│   ├── detail.html      # Page de détails
│   ├── error.html       # Page d'erreur générique
│   ├── 404.html         # Page 404
│   ├── 400.html         # Page 400
│   └── 500.html         # Page 500
├── go.mod               # Dépendances Go
└── README.md            # Ce fichier
```

## 🛠️ Installation

### Prérequis

- Go 1.25.0 ou supérieur
- Connexion internet (pour accéder à l'API Potter DB)

### Étapes d'installation

1. **Cloner le projet**
```bash
git clone <votre-repo>
cd API_HP
```

2. **Vérifier l'installation de Go**
```bash
go version
```

3. **Initialiser les dépendances**
```bash
go mod tidy
```

## 🎯 Utilisation

### Lancer le serveur

Depuis la racine du projet (`API_HP/`), exécutez :

```bash
go run ./src/cmd
```

Le serveur démarre sur **http://localhost:8080**

### Accéder à l'application

Ouvrez votre navigateur et accédez à :
- **Page d'accueil** : http://localhost:8080
- **Recherche** : http://localhost:8080/?search=Harry
- **Filtres** : http://localhost:8080/?house=Gryffindor&species=Human
- **Détail personnage** : http://localhost:8080/character/[id]

### Exemples d'utilisation

**Rechercher "Harry"**
```
GET http://localhost:8080/?search=Harry
```

**Filtrer par maison Gryffondor**
```
GET http://localhost:8080/?house=Gryffindor
```

**Combiner recherche et filtres**
```
GET http://localhost:8080/?search=Hermione&house=Gryffindor&species=Human
```

## 🧪 Tests

### Exécuter tous les tests

```bash
go test ./src/... -v
```

### Exécuter les tests d'un package spécifique

```bash
go test ./src/controllers -v
go test ./src/services -v
go test ./src/helpers -v
```

### Couverture de tests

```bash
go test ./src/... -cover
```

## 🏗️ Architecture

### Pattern MVC

Le projet suit une architecture inspirée du pattern MVC :

- **Models** : Structures de données (personnages, pagination, erreurs)
- **Controllers** : Handlers HTTP gérant les requêtes
- **Services** : Logique métier et communication avec l'API externe
- **Routers** : Configuration des routes
- **Helpers** : Fonctions utilitaires réutilisables
- **Middlewares** : Gestion des erreurs HTTP
- **Templates** : Vues HTML

### Flux de données

```
Client → Router → Controller → Service → API Potter DB
                      ↓
                   Template
                      ↓
                   Client
```

## 🎨 Personnalisation

### Modifier le CSS

Les styles se trouvent dans `assets/css/style.css`

### Modifier les templates

Les templates HTML sont dans le dossier `templates/`

### Ajouter de nouveaux filtres

1. Modifier `characters.controller.go` pour récupérer le paramètre
2. Modifier `characters.service.go` pour construire l'URL de l'API
3. Modifier `home.html` pour ajouter le filtre dans le formulaire

## 🐛 Gestion des erreurs

L'application gère les erreurs HTTP suivantes :

- **400 Bad Request** : Paramètres invalides
- **404 Not Found** : Page ou ressource introuvable
- **500 Internal Server Error** : Erreur serveur ou API

Chaque erreur affiche une page dédiée avec un message explicatif.

## 📚 API utilisée

**Potter DB API** : https://api.potterdb.com/v1

Endpoints utilisés :
- `GET /characters` - Liste des personnages
- `GET /characters/{id}` - Détails d'un personnage

## 🤝 Contribution

Ce projet est un projet pédagogique. Pour contribuer :

1. Fork le projet
2. Créer une branche (`git checkout -b feature/AmazingFeature`)
3. Commit vos changements (`git commit -m 'Add some AmazingFeature'`)
4. Push vers la branche (`git push origin feature/AmazingFeature`)
5. Ouvrir une Pull Request

## 📝 Auteur

Projet réalisé dans le cadre du cours Ynov 2025-2026

## 📄 Licence

Ce projet est à des fins éducatives uniquement.
# ProjGO.2-Groupie-tracker
