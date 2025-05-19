package utils

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

const (
	// MinPasswordLength définit la longueur minimale requise pour un mot de passe
	MinPasswordLength = 8
	// DefaultCost définit le coût par défaut pour le hachage bcrypt
	DefaultCost = 14
)

// HashPassword génère un hachage sécurisé du mot de passe en utilisant bcrypt
// Retourne une erreur si le mot de passe est trop court ou si le hachage échoue
func HashPassword(password string) (string, error) {
	if len(password) < MinPasswordLength {
		return "", errors.New("le mot de passe doit contenir au moins 8 caractères")
	}

	bytes, err := bcrypt.GenerateFromPassword([]byte(password), DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// CompareHashedPassword compare un mot de passe en clair avec un hachage bcrypt
// Retourne true si le mot de passe correspond au hachage, false sinon
func CompareHashedPassword(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
