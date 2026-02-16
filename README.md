# Harry Potter Wiki - Application Web Go

Un site web en Go qui utilise l'API Potter DB pour explorer l'univers Harry Potter : personnages, livres, films, sorts et potions.

## Fonctionnalités

- **Page d'accueil** : Affichage de personnages aléatoires
- **Recherche avancée** : Filtrer par nom, maison, espèce et statut
- **Fiches détaillées** : Informations complètes sur chaque élément
- **Catalogues** : Livres, films, sorts et potions
- **Pagination** : Navigation par pages avec conservation des filtres
- **Favoris** : Ajout et suppression de personnages favoris
- **Gestion d'erreurs** : Pages d'erreur personnalisées

## Technologies Utilisées

- **Go 1.25.0** : Langage de programmation pour le serveur backend
- **HTML/CSS** : Templates et mise en page
- **API Potter DB** : Source de données
- **JSON** : Stockage local des favoris

## Structure du Projet

```
API_HP/
├── src/
│   ├── cmd/main.go           
│   ├── controllers/           
│   ├── models/                
│   ├── routers/               
│   ├── services/              
│   ├── helpers/               
│   └── templates/             
├── templates/                 
├── assets/
│   ├── css/                   
│   └── img/                   
├── favorites.json             
└── go.mod                     
```

## Installation et Lancement

### Prérequis

- Go 1.25.0 ou version supérieure
- Connexion Internet

### Commandes

```bash
# Cloner le projet
git clone <url-du-repo>
cd API_HP

# Lancer le serveur
cd src/cmd
go run main.go
```

Le serveur démarre sur **http://localhost:8080**

## Routes Disponibles

| Route                  | Méthode | Description                      |
|------------------------|---------|----------------------------------|
| `/`                    | GET     | Page d'accueil                   |
| `/search`              | GET     | Recherche de personnages         |
| `/character/{id}`      | GET     | Détails d'un personnage          |
| `/favorites`           | GET     | Liste des favoris                |
| `/favorites/add`       | POST    | Ajouter un favori                |
| `/favorites/remove`    | POST    | Retirer un favori                |
| `/books`               | GET     | Catalogue des livres             |
| `/book/{id}`           | GET     | Détails d'un livre               |
| `/movies`              | GET     | Catalogue des films              |
| `/movie/{id}`          | GET     | Détails d'un film                |
| `/spells`              | GET     | Catalogue des sorts              |
| `/spell/{id}`          | GET     | Détails d'un sort                |
| `/potions`             | GET     | Catalogue des potions            |
| `/potion/{id}`         | GET     | Détails d'une potion             |
| `/about`               | GET     | Page À propos                    |
| `/error`               | GET     | Page d'erreur                    |

## Ressources Utilisées

### Documentation Go - Packages utilisés

- [Cours de Go de Cyril Rodrigues]
- [Documentation officielle Go](https://go.dev/doc/) - Guide complet du langage Go
- [Package net/http](https://pkg.go.dev/net/http) - Serveur web, requêtes HTTP, routage
  - [http.HandleFunc](https://pkg.go.dev/net/http#HandleFunc) - Enregistrement des routes
  - [http.ListenAndServe](https://pkg.go.dev/net/http#ListenAndServe) - Démarrage du serveur
  - [http.Client](https://pkg.go.dev/net/http#Client) - Requêtes vers l'API externe
  - [http.ResponseWriter](https://pkg.go.dev/net/http#ResponseWriter) - Écriture des réponses HTTP
- [Package html/template](https://pkg.go.dev/html/template) - Rendu des templates HTML
  - [template.ParseFiles](https://pkg.go.dev/html/template#ParseFiles) - Chargement des fichiers HTML
  - [template.Execute](https://pkg.go.dev/html/template#Template.Execute) - Rendu avec données
- [Package encoding/json](https://pkg.go.dev/encoding/json) - Manipulation JSON
  - [json.NewDecoder](https://pkg.go.dev/encoding/json#NewDecoder) - Décodage des réponses API
  - [json.Marshal/Unmarshal](https://pkg.go.dev/encoding/json#Marshal) - Conversion structures Go ↔ JSON
- [Package net/url](https://pkg.go.dev/net/url) - Manipulation des URLs
  - [url.QueryEscape](https://pkg.go.dev/net/url#QueryEscape) - Encodage des paramètres de recherche
- [Package fmt](https://pkg.go.dev/fmt) - Formatage de strings (Sprintf pour URLs)
- [Package time](https://pkg.go.dev/time) - Timeout des requêtes HTTP
- [Package sync](https://pkg.go.dev/sync) - Mutex pour gestion concurrence (favoris)

### API Externe

- [Potter DB API](https://potterdb.com/) - Base de données Harry Potter
- [Potter DB Documentation](https://docs.potterdb.com/) - Documentation complète de l'API REST
  - [Characters Endpoint](https://docs.potterdb.com/apis/rest#get-characters) - Liste et recherche de personnages
  - [Books Endpoint](https://docs.potterdb.com/apis/rest#get-books) - Récupération des livres
  - [Movies Endpoint](https://docs.potterdb.com/apis/rest#get-movies) - Récupération des films
  - [Spells Endpoint](https://docs.potterdb.com/apis/rest#get-spells) - Récupération des sorts
  - [Potions Endpoint](https://docs.potterdb.com/apis/rest#get-potions) - Récupération des potions
  - [Pagination](https://docs.potterdb.com/apis/rest#pagination) - Système de pagination de l'API
  - [Filtering](https://docs.potterdb.com/apis/rest#filtering) - Filtres de recherche

### Tutoriels et Guides

- [Go by Example](https://gobyexample.com/) - Exemples pratiques de code Go
  - [HTTP Servers](https://gobyexample.com/http-servers) - Création de serveurs web
  - [HTTP Clients](https://gobyexample.com/http-clients) - Requêtes HTTP
  - [JSON](https://gobyexample.com/json) - Manipulation JSON en Go
- [Building Web Applications with Go](https://go.dev/doc/articles/wiki/) - Tutoriel officiel
- [Making HTTP requests in Go](https://go.dev/doc/tutorial/web-service-gin) - Guide requêtes HTTP
- [Working with Templates](https://gowebexamples.com/templates/) - Utilisation des templates HTML


### Concepts Utilisés

- **Architecture MVC** (Model-View-Controller) - Séparation de la logique
- **Appels API REST** - Requêtes HTTP vers API externe
- **Gestion des erreurs HTTP** - Codes de statut et pages d'erreur
- **Routage HTTP** - Configuration des routes avec `net/http`
- **Templates HTML dynamiques** - Génération de pages avec données
- **Stockage de données JSON** - Sauvegarde locale des favoris

## Mon Avis

### Ce que j'ai appris

Ce projet m'a bien aidé à comprendre **Go**. Au début, j'étais un peu perdu avec la syntaxe, mais finalement c'est assez simple et direct. Ce qui est cool, c'est qu'on peut faire un serveur web complet juste avec la bibliothèque standard, sans installer beaucoup de packages.

Travailler avec l'**API Potter DB** m'a appris à gérer les appels REST et parser du JSON. Le système de favoris m'a fait découvrir comment sauvegarder des données localement et gérer la concurrence avec les mutex.

### Les défis

La partie la plus difficile a été la **gestion des filtres** dans la recherche. Il fallait construire l'URL de l'API dynamiquement selon les filtres actifs, tout en gardant la pagination. J'ai refait mon code plusieurs fois avant que ça marche bien.

Le **CSS responsive** était aussi compliqué. J'ai commencé avec un design très chargé en animations, mais au final j'ai tout simplifié pour avoir quelque chose de propre et lisible.

### Ce qui m'a plu

Go est vraiment rapide et les messages d'erreur sont clairs. Le fait de coder **tout de A à Z** (backend, frontend, routing, API) m'a donné une vraie vision d'ensemble du développement web.

Et puis faire un site sur Harry Potter, c'est quand même plus motivant qu'un projet classique !

## Auteur

**Hugo P.** — Ynov B1 Informatique 2025-2026