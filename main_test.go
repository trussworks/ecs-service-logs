package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/suite"
)

type cliTestSuite struct {
	suite.Suite
	viper  *viper.Viper
	logger *log.Logger
}

func (suite *cliTestSuite) Setup() {
	// Disable any logging that isn't attached to the logger unless using the verbose flag
	log.SetOutput(ioutil.Discard)
	log.SetFlags(0)

	// Setup logger
	var logger = log.New(os.Stdout, "", log.LstdFlags)

	// Remove the flags for the logger
	logger.SetFlags(0)
	suite.SetLogger(logger)

	// Setup viper
	v := viper.New()
	suite.SetViper(v)
}

func (suite *cliTestSuite) TestCheckRegion() {
	suite.Setup()

	testValues := []string{
		endpoints.UsEast1RegionID,
		endpoints.UsEast2RegionID,
		endpoints.UsWest1RegionID,
		endpoints.UsWest2RegionID,
		endpoints.UsGovWest1RegionID,
	}
	for _, testValue := range testValues {
		suite.viper.Set(flagAWSRegion, testValue)
		suite.NoError(checkConfig(suite.viper))
	}
	testValuesWithErrors := []string{
		"AnyOtherRegionName",
	}
	for _, testValue := range testValuesWithErrors {
		suite.viper.Set(flagAWSRegion, testValue)
		suite.Error(checkConfig(suite.viper))
	}
}

func (suite *cliTestSuite) SetViper(v *viper.Viper) {
	suite.viper = v
}

func (suite *cliTestSuite) SetLogger(logger *log.Logger) {
	suite.logger = logger
}
