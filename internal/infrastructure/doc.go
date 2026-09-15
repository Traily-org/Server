// Package infrastructure regroupe les adapters secondaires (driven) : tout ce qui
// implémente une interface définie par le domaine pour parler à l'extérieur
// (base de données, cache, API tierces, ...). Le domaine ne dépend jamais de ce
// package, c'est ce package qui dépend du domaine.
package infrastructure
