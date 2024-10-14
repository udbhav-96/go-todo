package main

import "github.com/udbhav-96/go-todo/internal/models"

type templateData struct {
    Task *models.Tasks
    Tasks []*models.Tasks
}