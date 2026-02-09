# Groupie_Tracker
Groupie Tracker project worked on during Ynov cursus

## Description

**Groupie Tracker** est une application web développée en **Go** qui permet de visualiser et manipuler des données sur des groupes de musique et des artistes. Le projet consomme une API RESTful externe pour récupérer des informations complexes (membres, dates de création, discographie, concerts) et les affiche via une interface utilisateur moderne et ergonomique.

Ce projet met en avant l'utilisation du langage Go pour le backend, la gestion de templates HTML dynamiques, et l'intégration de fonctionnalités de cartographie interactive.

## Fonctionnalités

* **Catalogue d'Artistes** : Affichage d'une liste complète d'artistes avec pagination ou défilement, incluant leurs images et informations principales.
* **Recherche Intelligente** : Système de barre de recherche permettant de trouver un groupe par son nom ou par le nom de ses membres.
* **Fiche Détaillée** : Page dédiée pour chaque artiste affichant :
    * Les membres du groupe.
    * La date de création et le premier album.
    * Les dates et lieux de concerts (Relations).
* **Carte Interactive (Live)** : Visualisation géographique des lieux de concerts à venir sur une carte mondiale (utilisant Leaflet & OpenStreetMap).
* **Interface Responsive** : Design soigné utilisant le style "Glassmorphism" (Tailwind CSS) adapté aux mobiles et aux ordinateurs.
* **Gestion d'Erreurs** : Pages personnalisées pour les erreurs 404 (Page non trouvée) et 500 (Erreur serveur).

## Stack Technique

### Backend
* **Langage** : Go (Golang)
* **Serveur Web** : Librairie standard `net/http` (Utilisation de `http.ServeMux`)
* **Templating** : `html/template`
* **Données** : Consommation de l'API [Groupie Trackers](https://groupietrackers.herokuapp.com/api)

### Frontend
* **Structure** : HTML5 sémantique
* **Style** : Tailwind CSS (via CDN)
* **Scripting** : JavaScript (Vanilla)
* **Cartographie** : Leaflet.js & API Nominatim (OpenStreetMap)

## Structure du Projet

L'architecture du projet respecte les bonnes pratiques de Go :

.
├── main.go                # Point d'entrée de l'application
├── go.mod                 # Gestion des dépendances
├── internal/
│   ├── api/               # Client HTTP pour les appels à l'API externe
│   ├── handlers/          # Contrôleurs des routes (Home, Artists, Live...)
│   ├── models/            # Structures de données (Structs Go)
│   └── render/            # Gestionnaire de rendu des templates HTML
├── server/
│   └── server.go          # Configuration du routeur (Mux) et démarrage serveur
└── web/
├── static/            # Fichiers statiques (JS, CSS, Images)
│   └── js/            # Scripts (ex: map.js)
└── templates/         # Fichiers HTML (index.html, artist.html, live.html...)

Installation et Lancement
Prérequis
Avoir Go installé sur votre machine (version 1.18+ recommandée).

Instructions
Cloner le dépôt :

Bash
git clone [https://github.com/VOTRE-USERNAME/groupie-tracker.git](https://github.com/VOTRE-USERNAME/groupie-tracker.git)
cd groupie-tracker
Lancer le serveur :
Exécutez la commande suivante à la racine du projet :

Bash
go run main.go
Accéder à l'application :
Ouvrez votre navigateur web et rendez-vous à l'adresse indiquée dans le terminal (généralement) :

Plaintext
http://localhost:8080

Utilisation de l'API
L'application agit comme un intermédiaire entre le client et l'API Groupie Trackers. Elle récupère les données JSON brutes, les traite via des structures Go (internal/models), et les injecte dans les templates HTML.

Les endpoints principaux gérés sont :

/ : Accueil

/artists : Liste et recherche

/artist?id={id} : Détails d'un artiste spécifique

/live : Carte des concerts

/filter : Filtres avancés

Auteurs
Ce projet a été réalisé par : JEANNOT Louis et COURTY Joan



https://github.com/HaxTwoWater/Groupie_Tracker