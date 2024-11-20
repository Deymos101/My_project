package ftracker

import (
    "fmt"
    "math"
)

// Основные константы, необходимые для расчетов.
const (
    lenStep   = 0.65  // средняя длина шага.
    mInKm     = 1000  // количество метров в километре.
    minInH    = 60    // количество минут в часе.
    kmhInMsec = 0.27778 // коэффициент для преобразования км/ч в м/с.
    cmInM     = 100   // количество сантиметров в метре.
)

// distance возвращает дистанцию (в километрах), которую преодолел пользователь.
func distance(action int) float64 {
    return float64(action) * lenStep / mInKm
}

// meanSpeed возвращает среднюю скорость движения (км/ч).
func meanSpeed(action int, duration float64) float64 {
    if duration == 0 {
        return 0
    }
    return distance(action) / duration
}

// ShowTrainingInfo возвращает строку с информацией о тренировке.
func ShowTrainingInfo(action int, trainingType string, duration, weight, height float64, lengthPool, countPool int) string {
    switch trainingType {
    case "Бег":
        dist := distance(action)
        speed := meanSpeed(action, duration)
        calories := RunningSpentCalories(action, weight, duration)
        return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", trainingType, duration, dist, speed, calories)
    case "Ходьба":
        dist := distance(action)
        speed := meanSpeed(action, duration)
        calories := WalkingSpentCalories(action, duration, weight, height)
        return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", trainingType, duration, dist, speed, calories)
    case "Плавание":
        dist := float64(lengthPool*countPool) / mInKm
        speed := swimmingMeanSpeed(lengthPool, countPool, duration)
        calories := SwimmingSpentCalories(lengthPool, countPool, duration, weight)
        return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", trainingType, duration, dist, speed, calories)
    default:
        return "Неизвестный тип тренировки"
    }
}

// RunningSpentCalories возвращает калории, потраченные при беге.
func RunningSpentCalories(action int, weight, duration float64) float64 {
    speed := meanSpeed(action, duration)
    return 18*speed*1.79*weight*duration / mInKm
}

// WalkingSpentCalories возвращает калории, потраченные при ходьбе.
func WalkingSpentCalories(action int, duration, weight, height float64) float64 {
    speed := meanSpeed(action, duration) * kmhInMsec // перевод в м/с
    return ((0.035 * weight) + (math.Pow(speed, 2) / (height / cmInM) * 0.029 * weight)) * duration * minInH
}

// swimmingMeanSpeed возвращает среднюю скорость при плавании.
func swimmingMeanSpeed(lengthPool, countPool int, duration float64) float64 {
    if duration == 0 {
        return 0
    }
    return float64(lengthPool*countPool) / mInKm / duration
}

// SwimmingSpentCalories возвращает калории, потраченные при плавании.
func SwimmingSpentCalories(lengthPool, countPool int, duration, weight float64) float64 {
    speed := swimmingMeanSpeed(lengthPool, countPool, duration)
    return (speed + 1.1) * 2 * weight * duration
}