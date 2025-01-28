package models

import (
	"fmt"
	"time"

	"github.com/gemdivk/Crowdfunding-system/internal/db"
)

type UserPoints struct {
	UserID    string    `json:"user_id"`
	Points    int       `json:"points"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Achievement struct {
	ID              int       `json:"id"`
	UserID          string    `json:"user_id"`
	AchievementName string    `json:"achievement_name"`
	AchievedAt      time.Time `json:"achieved_at"`
}

// --- User Points CRUD ---

func GetAllUserPoints() ([]UserPoints, error) {
	var points []UserPoints
	query := `SELECT user_id, points, created_at, updated_at FROM "UserPoints"`
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("Failed to fetch user points: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var userPoint UserPoints
		if err := rows.Scan(&userPoint.UserID, &userPoint.Points, &userPoint.CreatedAt, &userPoint.UpdatedAt); err != nil {
			return nil, fmt.Errorf("Failed to scan user points: %v", err)
		}
		points = append(points, userPoint)
	}
	return points, nil
}

func AddUserPoints(userPoints UserPoints) error {
	query := `
		INSERT INTO "UserPoints" (user_id, points, created_at, updated_at)
		VALUES ($1, $2, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
	`
	_, err := db.DB.Exec(query, userPoints.UserID, userPoints.Points)
	if err != nil {
		return fmt.Errorf("Failed to add user points: %v", err)
	}
	return nil
}

func UpdateUserPoints(userID string, points int) error {
	query := `
		UPDATE "UserPoints"
		SET points = $1, updated_at = CURRENT_TIMESTAMP
		WHERE user_id = $2
	`
	_, err := db.DB.Exec(query, points, userID)
	if err != nil {
		return fmt.Errorf("Failed to update user points: %v", err)
	}
	return nil
}

func DeleteUserPoints(userID string) error {
	query := `DELETE FROM "UserPoints" WHERE user_id = $1`
	_, err := db.DB.Exec(query, userID)
	if err != nil {
		return fmt.Errorf("Failed to delete user points: %v", err)
	}
	return nil
}

// --- Achievements CRUD ---

func GetAllAchievements() ([]Achievement, error) {
	var achievements []Achievement
	query := `SELECT id, user_id, achievement_name, achieved_at FROM "Achievements"`
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("Failed to fetch achievements: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var achievement Achievement
		if err := rows.Scan(&achievement.ID, &achievement.UserID, &achievement.AchievementName, &achievement.AchievedAt); err != nil {
			return nil, fmt.Errorf("Failed to scan achievement: %v", err)
		}
		achievements = append(achievements, achievement)
	}
	return achievements, nil
}

func AddAchievement(achievement Achievement) error {
	query := `
		INSERT INTO "Achievements" (user_id, achievement_name, achieved_at)
		VALUES ($1, $2, CURRENT_TIMESTAMP)
	`
	_, err := db.DB.Exec(query, achievement.UserID, achievement.AchievementName)
	if err != nil {
		return fmt.Errorf("Failed to add achievement: %v", err)
	}
	return nil
}

func DeleteAchievement(id int) error {
	query := `DELETE FROM "Achievements" WHERE id = $1`
	_, err := db.DB.Exec(query, id)
	if err != nil {
		return fmt.Errorf("Failed to delete achievement: %v", err)
	}
	return nil
}
