package db

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/emiliogozo/panahon-api-go/internal/util"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type APITokenTestSuite struct {
	suite.Suite
}

func TestAPITokenTestSuite(t *testing.T) {
	suite.Run(t, new(APITokenTestSuite))
}

func (ts *APITokenTestSuite) SetupTest() {
	err := testMigration.Up()
	require.NoError(ts.T(), err, "db migration problem")
}

func (ts *APITokenTestSuite) TearDownTest() {
	err := testMigration.Down()
	require.NoError(ts.T(), err, "reverse db migration problem")
}

func (ts *APITokenTestSuite) TestCreateAPIToken() {
	t := ts.T()
	createRandomAPIToken(t, nil)
}

func (ts *APITokenTestSuite) TestGetAPIToken() {
	t := ts.T()
	tok := createRandomAPIToken(t, nil)

	gotTok, err := testStore.GetAPIToken(context.Background(), tok.ID)
	require.NoError(t, err)
	require.NotEmpty(t, gotTok)

	require.Equal(t, tok.UserID, gotTok.UserID)
	require.Equal(t, tok.Name, gotTok.Name)
	require.Equal(t, tok.TokenHash, gotTok.TokenHash)
}

func (ts *APITokenTestSuite) TestGetAPITokenByHash() {
	t := ts.T()
	tok := createRandomAPIToken(t, nil)

	gotTok, err := testStore.GetAPITokenByHash(context.Background(), tok.TokenHash)
	require.NoError(t, err)
	require.NotEmpty(t, gotTok)

	require.Equal(t, tok.UserID, gotTok.UserID)
	require.Equal(t, tok.Name, gotTok.Name)
	require.Equal(t, tok.TokenHash, gotTok.TokenHash)
}

func (ts *APITokenTestSuite) TestListAPITokensByUser() {
	t := ts.T()
	n, m := 10, 5
	ctx := context.Background()
	toks := make([]APIToken, n)
	user := createRandomUser(t)
	for i := range n {
		if i < m {
			toks[i] = createRandomAPIToken(t, &user)
		} else {
			toks[i] = createRandomAPIToken(t, nil)
		}
	}

	gotToks, err := testStore.ListAPITokensByUser(ctx, user.ID)
	require.NoError(t, err)
	require.Len(t, gotToks, m)

	for _, tok := range gotToks {
		require.NotEmpty(t, tok)
	}
}

func (ts *APITokenTestSuite) TestCountAPITokensByUser() {
	t := ts.T()
	n, m := 10, 3
	ctx := context.Background()
	toks := make([]APIToken, n)
	user := createRandomUser(t)
	for i := range n {
		if i < m {
			toks[i] = createRandomAPIToken(t, &user)
		} else {
			toks[i] = createRandomAPIToken(t, nil)
		}
	}

	nToks, err := testStore.CountAPITokensByUser(ctx, user.ID)
	require.NoError(t, err)
	require.Equal(t, nToks, int64(m))
}

func (ts *APITokenTestSuite) TestDeleteAPIToken() {
	t := ts.T()
	tok := createRandomAPIToken(t, nil)

	err := testStore.DeleteAPIToken(context.Background(), tok.ID)
	require.NoError(t, err)

	gotTok, err := testStore.GetAPIToken(context.Background(), tok.ID)
	require.Error(t, err)
	require.Empty(t, gotTok)
}

func createRandomAPIToken(t *testing.T, user *User) APIToken {
	if user == nil {
		u := createRandomUser(t)
		user = &u
	}

	tokPrefix := "mo-"
	rawTok := fmt.Sprintf("%s-%s", tokPrefix, util.RandomString(32))
	arg := CreateAPITokenParams{
		UserID:      user.ID,
		Name:        util.RandomString(12),
		TokenPrefix: tokPrefix,
		TokenHash:   getHash(rawTok),
		Permissions: createStubAPITokenPermissions(t),
		ExpiresAt: pgtype.Timestamptz{
			Time:  time.Now().Add(time.Hour * 24),
			Valid: true,
		},
	}

	tok, err := testStore.CreateAPIToken(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, tok)

	require.Equal(t, arg.UserID, tok.UserID)
	require.Equal(t, arg.Name, tok.Name)
	require.Equal(t, arg.TokenPrefix, tok.TokenPrefix)
	require.Equal(t, arg.TokenHash, tok.TokenHash)
	require.True(t, tok.CreatedAt.Valid)
	require.NotZero(t, tok.CreatedAt.Time)

	return tok
}

func getHash(tok string) string {
	hash := sha256.Sum256([]byte(tok))
	return hex.EncodeToString(hash[:])
}

type APITokenPermissions struct {
	StationIDs     []int64  `json:"station_ids"`      // null means all stations
	Variables      []string `json:"variables"`        // null means all variables
	MaxDaysHistory int      `json:"max_days_history"` // 0 means unlimited
}

func createStubAPITokenPermissions(t *testing.T) []byte {
	perm := APITokenPermissions{
		StationIDs: []int64{util.RandomInt[int64](10000, 99999), util.RandomInt[int64](10000, 99999)},
		Variables:  []string{"temp", "rain"},
	}

	data, err := json.Marshal(perm)
	require.NoError(t, err)
	require.NotEmpty(t, data)
	return data
}
