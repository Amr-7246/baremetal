package types

import (
	"log"
	"taskflow/config"
	"taskflow/helper"
	"taskflow/helper/logger"
	"taskflow/repository"
	"taskflow/test"

	_ "github.com/lib/pq" //! we do import the driver in the "Blank Import" mode because of we do not need to implement function from it, we just triger the init() function to link the database/sql with it after it register in self right there
)
func main()  {
	//& load the app config
		cfg, err := config.Load()
		if err != nil {
			log.Fatalf("Failed to load config: %v", err)
		}
	//& initialize the logger
		logger, err := logger.New()
		if err != nil {
			log.Fatalf("Failed to load config: %v", err)
		}
		logger.Infof("Starting TaskFlow API on port %d", cfg.Port)
		logger.Infof("On the %s Environment", cfg.AppEnv)
	//& DB connection
		db, err := repository.NewDBConnection(&cfg.DB)
		if err != nil {
			logger.Fatalf("Failed to connect to database: %v", err)
		}
		defer db.Close()
		logger.Info("Successfully connected to database")
	//& Run migrations (this is a simple way, in production use migration tool)
		if err := helper.RunMigrations(db, logger); err != nil {
			logger.Fatalf("Failed to run migrations: %v", err)
		}
	//& start the user repo
		userRepo := repository.NewPostgresUserRepository(db) //? Why that step
	//& Test the repository
		if err := test.TestUserRepository(userRepo, logger); err != nil {
			logger.Errorf("Repository test failed: %v", err)
		} else {
			logger.Info("User repository tests passed!")
		}

		logger.Info("Task 2 complete! Ready for Task 3")
}