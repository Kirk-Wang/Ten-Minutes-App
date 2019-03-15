package database

import (
	"log"
	"os"
	"strconv"

	"github.com/TerrexTech/go-commonutils/commonutil"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pkg/errors"
	mgo "go.mongodb.org/mongo-driver/mongo"
)

var _ = Describe("MongoClient", func() {
	Context("new client is created", func() {
		var (
			connectionTimeout uint32
			clientConfig      ClientConfig
			testDatabase      string
			client            Client
		)

		BeforeEach(func() {
			hosts := os.Getenv("MONGO_TEST_HOSTS")
			username := os.Getenv("MONGO_TEST_USERNAME")
			password := os.Getenv("MONGO_TEST_PASSWORD")
			connectionTimeoutStr := os.Getenv("MONGO_TEST_CONNECTION_TIMEOUT_MS")
			testDatabase = os.Getenv("MONGO_TEST_DATABASE")

			var err error
			// Set Connection Timeout
			connectionTimeoutInt, err := strconv.Atoi(connectionTimeoutStr)
			if err != nil {
				err = errors.Wrap(
					err,
					"error getting CONNECTION_TIMEOUT from env, will use 1000",
				)
				log.Println(err)
				connectionTimeoutInt = 1000
			}
			connectionTimeout = uint32(connectionTimeoutInt)

			// Client Configuration
			clientConfig = ClientConfig{
				Hosts:               *commonutil.ParseHosts(hosts),
				Username:            username,
				Password:            password,
				TimeoutMilliseconds: connectionTimeout,
			}
		})

		AfterEach(func() {
			err := client.Disconnect()
			Expect(err).ToNot(HaveOccurred())
		})

		It("should create a new client", func() {
			client, err := NewClient(clientConfig)
			Expect(err).ToNot(HaveOccurred())

			Expect(client).To(BeAssignableToTypeOf(&Client{}))
			Expect(client.client).ToNot(BeNil())
			Expect(client.client).To(BeAssignableToTypeOf(&mgo.Client{}))
		})

		It("should return any errors that occur", func() {
			clientConfig.Hosts = []string{"@invalid-url^@"}
			_, err := NewClient(clientConfig)
			Expect(err).To(HaveOccurred())
		})

		It("should set the correct timeout as specified", func() {
			clientConfig.TimeoutMilliseconds = 2976
			client, err := NewClient(clientConfig)
			Expect(err).ToNot(HaveOccurred())
			Expect(client.timeout).To(Equal(uint32(2976)))
		})

		It("should set the default timeout if timeout is not specified", func() {
			clientConfig.TimeoutMilliseconds = 0
			client, err := NewClient(clientConfig)
			Expect(err).ToNot(HaveOccurred())
			Expect(client.timeout).To(Equal(uint32(1000)))
		})

		It(
			"should not connect automatically if the NoDefaultConnect is specified",
			func() {
				clientConfig.NoDefaultConnect = true
				client, err := NewClient(clientConfig)
				Expect(err).ToNot(HaveOccurred())
				Expect(client.connected).To(BeFalse())
			},
		)

		It(
			"should connect automatically if the NoDefaultConnect is not specified",
			func() {
				client, err := NewClient(clientConfig)
				Expect(err).ToNot(HaveOccurred())
				Expect(client.connected).To(BeTrue())
			},
		)

		It("should connect to Database when Connect is called", func() {
			clientConfig.NoDefaultConnect = true
			client, err := NewClient(clientConfig)
			Expect(err).ToNot(HaveOccurred())

			err = client.Connect()
			Expect(err).ToNot(HaveOccurred())
			Expect(client.connected).To(BeTrue())
		})

		It("should disconnect from Database when Disconnect is called", func() {
			client, err := NewClient(clientConfig)
			Expect(err).ToNot(HaveOccurred())

			err = client.Disconnect()
			Expect(err).ToNot(HaveOccurred())
			Expect(client.connected).To(BeFalse())
		})

		It("should return the Database instance when Database is requested", func() {
			client, err := NewClient(clientConfig)
			Expect(err).ToNot(HaveOccurred())

			db := client.Database(testDatabase)
			Expect(db).To(BeAssignableToTypeOf(&mgo.Database{}))
			Expect(db.Name()).To(Equal(testDatabase))
		})

		// It("should return the correct DriverClient", func() {
		// 	client, err := NewClient(clientConfig)
		// 	Expect(err).ToNot(HaveOccurred())

		// 	dc := client.DriverClient()
		// 	Expect(dc).To(BeAssignableToTypeOf(&mgo.Client{}))

		// 	expectedConnStr := fmt.Sprintf(
		// 		"mongodb://%s:%s@%s",
		// 		clientConfig.Username,
		// 		clientConfig.Password,
		// 		clientConfig.Hosts[0],
		// 	)
		// 	Expect(dc.).To(Equal(expectedConnStr))
		// })
	})
})
