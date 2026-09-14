package stats

import (
	"context"
	"errors"
	"io"
	"time"
)

var ErrEmptyStats = errors.New("empty stats")

type GeneralStats struct {
	TotalReadPages    int
	TotalReadTime     int // in seconds
	AveragePagePerDay int
	AverageTimePerDay int // in seconds
	BookStats         []BookStatsWithTitle
}

type BookStatsWithTitle struct {
	Title string
	BookStats
}

type BookListItem struct {
	Title           string
	TotalPages      int
	TotalReadPages  int
	ProgressPercent int
	TotalReadTime   int // in seconds
	FirstRead       time.Time
	LastRead        time.Time
}

type CurrentlyReadingItem struct {
	Title      string
	Percentage int
	TotalPages int
	LastRead   time.Time
}

type ReadingStats interface {
	GetBookStats(ctx context.Context, fileHash string) (*BookStats, error)
	GetGeneralStats(ctx context.Context, from, to time.Time) (*GeneralStats, error)
	GetDailyStats(ctx context.Context, from, to time.Time) ([]DailyStats, error)
	GetBooksList(ctx context.Context) ([]BookListItem, error)
	GetCurrentlyReading(ctx context.Context) ([]CurrentlyReadingItem, error)
	Write(ctx context.Context, r io.ReadCloser, deviceName string) error
}
