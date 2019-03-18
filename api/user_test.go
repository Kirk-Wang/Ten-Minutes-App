package api

import (
	// "github.com/lotteryjs/ten-minutes-api/auth/password"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/lotteryjs/ten-minutes-api/mode"
	// "github.com/lotteryjs/ten-minutes-api/model"
	// "github.com/lotteryjs/ten-minutes-api/test"
	"github.com/lotteryjs/ten-minutes-api/test/testdb"
	// "github.com/stretchr/testify/assert"
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

func (s *UserSuite) Test_CreateUser() {
	s.ctx.Request = httptest.NewRequest("POST", "/user", strings.NewReader(`{"name": "tom", "pass": "mylittlepony", "admin": true}`))
	s.ctx.Request.Header.Set("Content-Type", "application/json")

	s.a.CreateUser(s.ctx)

	// user := &model.UserExternal{ID: 1, Name: "tom", Admin: true}
	// test.BodyEquals(s.T(), user, s.recorder)
	// assert.Equal(s.T(), 200, s.recorder.Code)

	// created := s.db.GetUserByName("tom")
	// assert.NotNil(s.T(), created)
	// assert.True(s.T(), password.ComparePassword(created.Pass, []byte("mylittlepony")))
}
