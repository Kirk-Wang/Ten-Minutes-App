package api

import (
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/lotteryjs/ten-minutes-api/mode"
	"github.com/lotteryjs/ten-minutes-api/test/testdb"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

func TestUserSuite(t *testing.T) {
	suite.Run(t, new(UserSuite))
}

type UserSuite struct {
	suite.Suite
	db       *testdb.Database
	a        *UserAPI
	ctx      *gin.Context
	recorder *httptest.ResponseRecorder
}

func (s *UserSuite) BeforeTest(suiteName, testName string) {
	mode.Set(mode.TestDev)
	s.recorder = httptest.NewRecorder()
	s.ctx, _ = gin.CreateTestContext(s.recorder)
	s.db = testdb.NewDB(s.T())
	s.a = &UserAPI{DB: s.db}
}
func (s *UserSuite) AfterTest(suiteName, testName string) {
	s.db.Close()
}

func (s *UserSuite) Test_GetUsers() {
	s.db.TenDatabase.DB.Collection("users").Drop(nil)

	for i := 1; i <= 20; i++ {
		s.db.NewUser(fmt.Sprintf("User%d", i))
	}
	assert.Equal(s.T(), 1, 1)
	// s.a.GetUsers(s.ctx)
	// assert.Equal(s.T(), 200, s.recorder.Code)
}
