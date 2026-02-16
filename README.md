# Groupie Tracker - Potter DB
Un site web en Go qui utilise l'API Potter DB pour explorer l'univers Harry Potter : personnages, livres, films, sorts et potions. Le but de ce projet est de s'entraîner à utiliser les API et à créer des sites web en Go. De plus, il permet de découvrir l'univers magique aux utilisateurs.

## Fonctionnalités
**Personnages** — Afficher des personnages aléatoires sur la page d'accueil  
**Recherche avancée** — Filtrer par nom, maison, espèce et statut (vivant/décédé)  
**Fiche détaillée** — Afficher les détails complets d'un personnage  
**Catalogue de contenu** — Parcourir les livres, films, sorts et potions  
**Pagination** — Navigation par pages avec conservation des filtres  
**Favoris** — Ajouter/supprimer des personnages en favoris  
**Page d'erreur stylisée** — Page d'erreur personnalisée avec codes HTTP  

## Technologies Utilisées
**Go** — Serveur backend (v1.25.0)  
**HTML/CSS** — Templates et mise en page  
**API Potter DB** — Source de données de l'univers Harry Potter  
**JSON** — Stockage local des favoris  

## Structure du projet
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

## Lancement

### Prérequis
- Go installé (v1.25.0 ou supérieur)
- Connexion internet (pour l'API Potter DB)

### Installation et exécution
```bash
# Cloner le projet
git clone <url-du-repo>
cd API_HP

# Lancer le serveur
cd src/cmd
go run main.go
```

Le serveur démarre sur **http://localhost:8080**

## Pages disponibles

| Route | Description |
|-------|-------------|
| `/` | Page d'accueil |
| `/search` | Recherche de personnages |
| `/character/{id}` | Détails d'un personnage |
| `/favorites` | Page des favoris |
| `/favorites/add` | Ajouter aux favoris (POST) |
| `/favorites/remove` | Retirer des favoris (POST) |
| `/books` | Catalogue des livres |
| `/book/{id}` | Détails d'un livre |
| `/movies` | Catalogue des films |
| `/movie/{id}` | Détails d'un film |
| `/spells` | Catalogue des sorts |
| `/spell/{id}` | Détails d'un sort |
| `/potions` | Catalogue des potions |
| `/potion/{id}` | Détails d'une potion |
| `/about` | Page À propos |
| `/error` | Page d'erreur |

## Mon Avis

### Ce que j'ai appris
Franchement, ce projet m'a bien aidé à comprendre **Go**. Au début, j'étais un peu perdu avec la syntaxe, mais finalement c'est assez simple et direct. Ce qui est cool, c'est qu'on peut faire un serveur web complet juste avec la lib standard, sans installer 50 packages comme en Node.js.

Bosser avec l'**API Potter DB** m'a appris à gérer les appels REST et parser du JSON. Le système de favoris m'a fait découvrir comment sauvegarder des données localement et gérer la concurrence avec les mutex (pour pas que tout plante si plusieurs personnes ajoutent des favoris en même temps).

### Les galères
La partie la plus chiante a été la **gestion des filtres** dans la recherche. Fallait construire l'URL de l'API dynamiquement selon les filtres actifs, tout en gardant la pagination. J'ai refait mon code genre 3 fois avant que ça marche bien.

Aussi, le **CSS responsive** c'était un peu relou. J'ai commencé avec un thème dark super chargé en animations, mais au final j'ai tout simplifié pour avoir quelque chose de propre et lisible. Parfois, less is more.

### Ce qui m'a plu
Go, c'est vraiment rapide et les messages d'erreur sont clairs (pas comme Python qui te sort des traceback de 50 lignes 😅). Le fait de coder **tout de A à Z** (back, front, routing, API) m'a donné une vraie vision d'ensemble du dev web.

Et puis bon, faire un site sur Harry Potter, c'est quand même plus motivant que "Gestionnaire de To-Do liste n°476" 🧙‍♂️

### Si je devais continuer
- Mettre une vraie BDD (PostgreSQL) au lieu du JSON
- Ajouter un système de comptes utilisateurs
- Un cache Redis pour que ce soit plus rapide
- Faire des tests (j'avoue, j'en ai pas fait...)
- Rendre le site utilisable hors-ligne (PWA)

### Conclusion
Ce projet m'a bien fait progresser en dev web. Go c'est un langage que je vais sûrement réutiliser, et l'univers Harry Potter rendait le tout plus sympa à développer !



## Auteur
**Hugo P.** — Ynov B1 Informatique 2025-2026

