package db

import (
	"context"
	"testing"
	"time"

	"github.com/emiliogozo/panahon-api-go/util"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"
)

func TestCreateStationHealth(t *testing.T) {
	station := createRandomStation(t)
	health := createRandomStationHealth(t, station.ID)

	// cleanup
	_deleteStationHealth(t, DeleteStationHealthParams{
		StationID: health.StationID,
		ID:        health.ID,
	})
}

func TestGetStationHealth(t *testing.T) {
	station := createRandomStation(t)
	health := createRandomStationHealth(t, station.ID)

	arg := GetStationHealthParams{
		StationID: health.StationID,
		ID:        health.ID,
	}
	gotHealth, err := testStore.GetStationHealth(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, gotHealth)

	require.Equal(t, health, gotHealth)

	// cleanup
	_deleteStationHealth(t, DeleteStationHealthParams{
		StationID: health.StationID,
		ID:        health.ID,
	})
}

func TestListStationHealths(t *testing.T) {
	station := createRandomStation(t)
	n := 10
	healths := make([]ObservationsStationhealth, n)
	for i := 0; i < n; i++ {
		healths[i] = createRandomStationHealth(t, station.ID)
	}

	arg := ListStationHealthsParams{
		StationID: station.ID,
		Limit:     5,
		Offset:    5,
	}

	gotHealths, err := testStore.ListStationHealths(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, gotHealths, 5)

	for _, h := range gotHealths {
		require.NotEmpty(t, h)
	}

	// cleanup
	for _, h := range healths {
		_deleteStationHealth(t, DeleteStationHealthParams{
			StationID: h.StationID,
			ID:        h.ID,
		})
	}
}

func TestUpdateStationHealth(t *testing.T) {
	var (
		oldHealth ObservationsStationhealth
		newBp1    util.NullFloat4
	)

	station := createRandomStation(t)

	testCases := []struct {
		name        string
		buildArg    func() UpdateStationHealthParams
		checkResult func(updatedHealth ObservationsStationhealth, err error)
	}{
		{
			name: "SomeValues",
			buildArg: func() UpdateStationHealthParams {
				oldHealth = createRandomStationHealth(t, station.ID)
				newBp1 = util.RandomNullFloat4(0.0, 15.0)

				return UpdateStationHealthParams{
					StationID: station.ID,
					ID:        oldHealth.ID,
					Bp1:       newBp1,
				}
			},
			checkResult: func(updatedHealth ObservationsStationhealth, err error) {
				require.NoError(t, err)
				require.NotEqual(t, oldHealth.Bp1, updatedHealth.Bp1)
				require.Equal(t, oldHealth.Vb1, updatedHealth.Vb1)
				require.Equal(t, oldHealth.Curr, updatedHealth.Curr)

				require.True(t, updatedHealth.UpdatedAt.Valid)
				require.NotZero(t, updatedHealth.UpdatedAt.Time)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			updatedHealth, err := testStore.UpdateStationHealth(context.Background(), tc.buildArg())
			tc.checkResult(updatedHealth, err)

			// cleanup
			_deleteStationHealth(t, DeleteStationHealthParams{
				StationID: updatedHealth.StationID,
				ID:        updatedHealth.ID,
			})
		})
	}
}

func TestDeleteStationHealth(t *testing.T) {
	station := createRandomStation(t)
	health := createRandomStationHealth(t, station.ID)

	// cleanup
	_deleteStationHealth(t, DeleteStationHealthParams{
		StationID: health.StationID,
		ID:        health.ID,
	})
}

func createRandomStationHealth(t *testing.T, stationID int64) ObservationsStationhealth {
	arg := CreateStationHealthParams{
		Vb1:  util.RandomNullFloat4(15.0, 30.0),
		Curr: util.RandomNullFloat4(0, 5.0),
		Timestamp: pgtype.Timestamptz{
			Time:  time.Now(),
			Valid: true,
		},
		StationID: stationID,
	}

	health, err := testStore.CreateStationHealth(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, health)

	require.Equal(t, arg.Vb1, health.Vb1)
	require.Equal(t, arg.Curr, health.Curr)
	require.Equal(t, arg.StationID, health.StationID)

	require.True(t, health.UpdatedAt.Time.IsZero())
	require.True(t, health.CreatedAt.Valid)
	require.NotZero(t, health.CreatedAt.Time)

	return health
}

func _deleteStationHealth(t *testing.T, arg DeleteStationHealthParams) {
	err := testStore.DeleteStationHealth(context.Background(), arg)
	require.NoError(t, err)

	gotHealth, err := testStore.GetStationHealth(context.Background(), GetStationHealthParams(arg))
	require.Error(t, err)
	require.Empty(t, gotHealth)
}
